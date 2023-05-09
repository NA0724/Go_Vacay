# Go_Vacay

## Description

Vacation Booking Application : Bed and Breakfast Booking with Go Vacay! 
This is the repository for Go Vacay application,which is a bookings and reservations website.

### Tech Stack:

- Built in Go version 1.19  
- Uses [chi router](https://github.com/go-chi/chi) v5.0.8, 
- [alex edwards scs session management](https://github.com/alexedwards/scs) v2.5.0, 
- [nosurf](https://github.com/justinas/nosurf) v1.1.1


To run the application, you need to do the following
1. go to config.go
2. inside `GetPath()` method, update your directory path
3. open terminal and type `./run.sh` command (MacOS/Linux only) or run `run.bat` for windows

To run test file
- `go test`
- `go test -v`
- `go test -cover`
- `go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

To make run.sh executable
- chmod +x run.sh

PostgreSQL Configuration using soda cli "https://github.com/gobuffalo/fizz/blob/main/README.md"
- database.yml
- `soda generate fizz CreateTable`
- `soda migrate` to run script in up migration file
- `soda migrate down` to run script in down migration file



