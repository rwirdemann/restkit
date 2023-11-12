# RESTKit
RESTKit is a simple CLI to generate REST APIs.

## Usage
```
rk create bookstore // generates the project 
rk add book         // adds the book resource
```

## Example
```bash
rk create bookstore
cd bookstore
go mod tidy
go run main.go
```

## Configuration

### Enviroment variables
```
# Root directory where new projects are generated
RESTKIT_ROOT   
```