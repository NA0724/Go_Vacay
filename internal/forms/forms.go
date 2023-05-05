package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct and embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}

// MinLength check for minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks for a valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// MaxLength check for maximum length
func (f *Form) MaxLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) > length {
		f.Errors.Add(field, fmt.Sprintf("This field must be less than %d characters long", length))
		return false
	}
	return true
}

// check if input is a numeric value only
func (f *Form) Is_numeric(word string, r *http.Request) bool {
	return govalidator.IsNumeric(word)
}

// check if input is a numeric value only
func (f *Form) IsAlphaNumeric(word string, r *http.Request) bool {
	return govalidator.IsAlphanumeric(word)
}

// check if input is a numeric value only
func (f *Form) IsSame(word1, word2 string) bool {
	return strings.Compare(word1, word2) == 0
}
