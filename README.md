# RESTKit
RESTKit is a simple CLI to generate REST APIs.

https://github.com/rwirdemann/restkit/assets/28768/bd227566-582d-4c83-a8fb-fad464837994

## Usage
```
restkit create bookstore  // generates the project 
restkit add resource book // adds the book resource
restkit --help            // prints help message
```

## Example
```bash
# Create project 
restkit create bookstore
cd bookstore

# Add first resource
restkit add resource book

# Update dependencies
go mod tidy

# Start server
go run main.go
```

## Build

```text
go get github.com/gobuffalo/packr/packr
make
```

## Configuration

### Environment variables
```
# Root directory where new projects are generated
RESTKIT_ROOT   

# RESTKit template directory
RESTKIT_TEMPLATES
```

## Understand the structure of the generated API
The structure of the generated REST API follows the idea of hexagonal architectures according to the
following directory layout:

```text
application
  domain
    book.go
  services
    books.go
ports
  in
    books_service.go
  out
    books_repository.go
context
  http
    books_handler.go
  sql
    books_repository.go    
```

All application specific code lives in the application package. Its subpackages `application.domain`
and `application.services` contain the core domain model and its related services. The `ports.in`
package provides interfaces that publish the services of the API. The `ports.out` package is used by
the application itself to access external components like databases.

The `context` package contains adapter classes that adapt the external world (e.g. HTTP, SQL) to the
language of the internal world. The subpackage name indicates the type of the external world and its
direction. Thus, the http.books_handler implements a set of HTTP endpoints like `GET /books` and
maps incoming HTTP requests to internal services. Incoming adapter classes use the interfaces
provided by the `ports.in`` package to delegate incoming requests to the suitable application
service.

