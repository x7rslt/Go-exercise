package complicated_test


//Validating a Simple Value
import (
	"fmt"
	"regexp"
	"testing"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func TestSimpleValue(t *testing.T) {
	data := "example"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		is.URL,                    // is a valid URL
	)
	fmt.Println(err)
	// Output:
	// must be a valid URL
}

//Validating a Struct
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func TestValidatStruct(t *testing.T){
	a := Address{
		Street: "123",
		City:   "Unknown",
		State:  "Virginia",
		Zip:    "12345",
	}

	err := a.Validate()
	fmt.Println(err)


}

