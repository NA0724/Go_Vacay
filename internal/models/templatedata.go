package models

import "Go_Vacay/internal/forms"

// holds data set from handlets to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Checkin   string
	Checkout  string
	Hotel     Hotel
}

// registeration form
type Registration struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
}

type Login struct {
	Email    string
	Password string
}

type Traveller struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Gender    string
}

type User struct {
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Phone      string
	Bookings   map[string]string
	Gender     string
	Travellers []Traveller
}

type Booking struct {
	Hotel    Hotel
	Room     Room
	Checkin  string
	Checkout string
	Details  string
}

type Hotel struct {
	Rooms    []Room
	Checkin  string
	Checkout string
	Details  string
	Address  string
}

type Room struct {
	Name        string
	Type        string
	Price       float64
	Description string
}
