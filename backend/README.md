# Backend

## Swagger documentation generation
```bash
swag init
```

## Coverage report
These commands will run all the tests in the `test` directory and generate a coverage report in the `cover.html` file. (note that this script only test on linux, for other OS, you may need to change the command accordingly)
``` bash
# Make sure you are in the backend directory
go test ./test -coverprofile=./cover.out -coverpkg=./src/...

go tool cover -html=cover.out -o cover.html
```