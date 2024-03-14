# RESTKit
RESTKit is a simple CLI to generate REST APIs.

https://github.com/rwirdemann/restkit/assets/28768/7ce5a4a9-66bf-49d5-b4a7-e1733aa2454d

## Usage
```
rk create module     // creates the project 
rk add resource book // adds the book resource
rk --help            // prints help message
rk add --help        // prints help a specific command
```

## Example
```bash
# Create project 
rk create github.com/rwirdemann/bookstore
cd $GOPATH/src/github.com/rwirdemann/bookstore

# Add first resource
rk add resource book title:string author:string

# Update dependencies
go mod tidy

# Start server
go run main.go
```

## Build

```text
make
```

## Understanding the structure of the generated API
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

# Database scripts

```
CREATE TABLE books
(
    id    serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);
```

