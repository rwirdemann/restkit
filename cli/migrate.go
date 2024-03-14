package cli

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Migrate(filename string) error {
	log.Printf("psql -d postgres < %s", filename)
	cmd := fmt.Sprintf("psql -d postgres < %s", filename)
	if out, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatalln(err)
		return err
	} else {
		logOutput(string(out))
	}

	return nil
}

func logOutput(out string) {
	lines := strings.Split(out, "\n")
	for _, l := range lines {
		if len(strings.Trim(l, "")) > 0 {
			log.Println(l)
		}
	}
}
