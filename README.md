# Go_Vacay

## Description

Vacation Booking Application : Bed and Breakfast Booking with Go Vacay! 
This is the repository for Go Vacay application,which is a bookings and reservations website.

### Tech Stack:

- Built in Go version 1.19  
- Uses [chi router](https://github.com/go-chi/chi) v5.0.8, 
- [alex edwards scs session management](https://github.com/alexedwards/scs) v2.5.0, 
- [nosurf](https://github.com/justinas/nosurf) v1.1.1


To change diretory path
1. go to config.go
2. inside `GetPath()` method, update your directory path

To run test file
go test
go test -v
go test -cover
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
