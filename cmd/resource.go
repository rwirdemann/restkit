package cmd

import (
	"fmt"
	"github.com/rwirdemann/restkit/textx"
	"log"
	"os"
	"unicode"

	"github.com/rwirdemann/restkit/ports"

	"github.com/rwirdemann/restkit/gotools"
	"github.com/spf13/cobra"
)

func init() {
	addCmd.AddCommand(resourceCmd)
}

var resourceCmd = &cobra.Command{
	Use:   "resource name",
	Short: "creates a resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := add(args[0]); err != nil {
			return err
		}
		return nil
	},
}

func add(resourceName string) error {
	// Check if current directory is a RESTkit's project root
	if !fileSystem.Exists(".restkit.yml") {
		return fmt.Errorf("current directory contains no .restkit.yml")
	}

	config, err := yml.ReadConfig()
	if err != nil {
		return err
	}

	if err := createHttpAdapter(resourceName, config); err != nil {
		return err
	}

	if err := createPostgresAdapter(resourceName, config); err != nil {
		return err
	}

	if err := createDomainObject(resourceName); err != nil {
		return err
	}

	if err := createService(resourceName, config); err != nil {
		return err
	}

	if err := createPorts(resourceName, config); err != nil {
		return err
	}

	if err := updateMain(resourceName, config); err != nil {
		return err
	}

	if err := gotools.Fmt(); err != nil {
		return err
	}

	return nil
}

func createPostgresAdapter(resourceName string, config ports.Config) error {
	// Create context dir if not exists
	if err := createDirIfNotExists("context"); err != nil {
		return err
	}

	// Create postgres dir if not exist
	postgresDir := fmt.Sprintf("%s%c%s", "context", os.PathSeparator, "postgres")
	if err := createDirIfNotExists(postgresDir); err != nil {
		return err
	}

	// Create resource handler file
	data := struct {
		Resource          string
		ResourceLowerCaps string
		Module            string
	}{
		Resource:          capitalize(resourceName),
		ResourceLowerCaps: resourceName,
		Module:            config.Module,
	}
	if err := createFromTemplate(fmt.Sprintf("%s_repository.go", pluralize(resourceName)), postgresDir, "postgres_repository.go.txt", data); err != nil {
		return err
	}

	return nil
}

func updateMain(resourceName string, config ports.Config) error {
	// Import postgres adatper
	if err := insertImportStatement(fmt.Sprintf("postgres \"%s/context/postgres\"", config.Module)); err != nil {
		return err
	}

	// Import http adapter
	if err := insertImportStatement(fmt.Sprintf("http2 \"%s/context/http\"", config.Module)); err != nil {
		return err
	}

	// Import services
	if err := insertImportStatement(fmt.Sprintf("\"%s/application/services\"", config.Module)); err != nil {
		return err
	}

	// Insert create adapter into main file
	builder := textx.FragmentBuilder{}
	builder.Append("%rsRepository := postgres.New%RsRepository(db)")
	check := builder.Build(resourceName)
	if contains, _ := template.Contains("main.go", check); contains {
		log.Printf("insert: %s...already there\n", "http handler")
	} else {
		log.Printf("insert: %s...ok\n", "http handler")
		builder := textx.FragmentBuilder{}
		builder.Append("%rsRepository := postgres.New%RsRepository(db)")
		builder.Append("%rsService := services.New%RsService(%rsRepository)")
		builder.Append("%rsAdapter := http2.New%RsHandler(*%rsService)")
		builder.Append("\trouter.HandleFunc(\"/%rs\", %rsAdapter.GetAll()).Methods(\"GET\")")
		f := builder.Build(resourceName)
		if err := template.Insert("main.go", "log.Printf(\"starting http service on port %d...\", c.Port)", f); err != nil {
			return err
		}
	}

	return nil
}

func createHttpAdapter(resourceName string, config ports.Config) error {
	// Create context dir if not exists
	if err := createDirIfNotExists("context"); err != nil {
		return err
	}

	// Create http dir if not exist
	httpDir := fmt.Sprintf("%s%c%s", "context", os.PathSeparator, "http")
	if err := createDirIfNotExists(httpDir); err != nil {
		return err
	}

	// Create http adapter file
	data := struct {
		Resource          string
		ResourceLowerCaps string
		Module            string
	}{
		Resource:          capitalize(resourceName),
		ResourceLowerCaps: resourceName,
		Module:            config.Module,
	}
	if err := createFromTemplate(fmt.Sprintf("%s_handler.go", pluralize(resourceName)), httpDir, "resource_handler.go.txt", data); err != nil {
		return err
	}

	return nil
}

func insertImportStatement(stmt string) error {
	if contains, _ := template.Contains("main.go", stmt); contains {
		log.Printf("insert: %s...already there\n", stmt)
	} else {
		log.Printf("insert: %s...ok\n", "import")
		if err := template.Insert("main.go", "\"net/http\"", stmt); err != nil {
			return err
		}
	}
	return nil
}

func createDomainObject(resourceName string) error {
	// Create application dir if not exist
	if err := createDirIfNotExists("application"); err != nil {
		return err
	}

	// Create domain dir if not exist
	appDir := fmt.Sprintf("%s%c%s", "application", os.PathSeparator, "domain")
	if err := createDirIfNotExists(appDir); err != nil {
		return err
	}

	// Create domain object for resource representation
	data := struct {
		Resource string
	}{
		Resource: capitalize(resourceName),
	}
	if err := createFromTemplate(fmt.Sprintf("%s.go", resourceName), appDir, "resource.go.txt", data); err != nil {
		return err
	}

	return nil
}

func createService(resourceName string, config ports.Config) error {
	// Create application dir if not exist
	if err := createDirIfNotExists("application"); err != nil {
		return err
	}

	// Create services dir if not exist
	appDir := fmt.Sprintf("%s%c%s", "application", os.PathSeparator, "services")
	if err := createDirIfNotExists(appDir); err != nil {
		return err
	}

	// Create service object for resource
	data := struct {
		Resource          string
		ResourceLowerCaps string
		Module            string
	}{
		Resource:          capitalize(resourceName),
		ResourceLowerCaps: resourceName,
		Module:            config.Module,
	}
	if err := createFromTemplate(fmt.Sprintf("%ss.go", resourceName), appDir, "service.go.txt", data); err != nil {
		return err
	}

	return nil
}

func createPorts(resourceName string, config ports.Config) error {
	// Create "ports" dir if not exists
	if err := createDirIfNotExists("ports"); err != nil {
		return err
	}

	// Create "in" dir if not exist
	inDir := fmt.Sprintf("%s%c%s", "ports", os.PathSeparator, "in")
	if err := createDirIfNotExists(inDir); err != nil {
		return err
	}

	data := struct {
		Resource string
		Module   string
	}{
		Resource: capitalize(resourceName),
		Module:   config.Module,
	}
	if err := createFromTemplate(fmt.Sprintf("%s_service.go", pluralize(resourceName)), inDir, "in_port.go.txt", data); err != nil {
		return err
	}

	// Create "out" dir if not exist
	outDir := fmt.Sprintf("%s%c%s", "ports", os.PathSeparator, "out")
	if err := createDirIfNotExists(outDir); err != nil {
		return err
	}

	if err := createFromTemplate(fmt.Sprintf("%s_repository.go", pluralize(resourceName)), outDir, "repository_out_port.go.txt", data); err != nil {
		return err
	}

	return nil
}

func createFromTemplate(filename, path, tmpl string, data interface{}) error {
	fn := fmt.Sprintf("%s/%s", path, filename)
	if !fileSystem.Exists(fn) {
		if err := template.Create(tmpl, filename, path, data); err != nil {
			return err
		}
		log.Printf("create: %s...ok\n", fn)
		return nil
	}

	log.Printf("create: %s...exists\n", fn)
	if force {
		if err := fileSystem.Remove(fn); err != nil {
			return err
		}
		log.Printf("remove: %s...ok\n", fn)
		if err := template.Create(tmpl, filename, path, data); err != nil {
			return err
		}
		log.Printf("create: %s...ok\n", fn)
	}
	return nil
}

func createDirIfNotExists(dir string) error {
	if !fileSystem.Exists(dir) {
		if err := fileSystem.CreateDir(dir); err != nil {
			return err
		}
	}
	return nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func pluralize(str string) string {
	return str + "s"
}
