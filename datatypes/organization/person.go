package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	Age            int
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "1234."
}

func (p *Person) SetHandler(handler string) error {
	if !strings.HasPrefix(handler, "@") {
		return errors.New("handler should start with '@'")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() string {
	return p.twitterHandler
}
