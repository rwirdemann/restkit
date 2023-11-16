# RESTKit
RESTKit is a simple CLI to generate REST APIs.

https://github.com/rwirdemann/restkit/assets/28768/bd227566-582d-4c83-a8fb-fad464837994

## Usage
```
restkit create bookstore // generates the project 
restkit add book         // adds the book resource
```

## Example
```bash
# Create project 
restkit create bookstore
cd bookstore

# Add first resource
restkit add book

# Update dependencies
go mod tidy

# Start server
go run main.go
```

## Configuration

### Enviroment variables
```
# Root directory where new projects are generated
RESTKIT_ROOT   

# RESTKit template directory
RESTKIT_TEMPLATES
```
