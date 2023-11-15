# RESTKit
RESTKit is a simple CLI to generate REST APIs.

https://github.com/rwirdemann/restkit/assets/28768/bd227566-582d-4c83-a8fb-fad464837994

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
