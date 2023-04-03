package organization

import (
	"errors"
	"fmt"
	"strings"
)

// type aliases
// type TwitterHandler = string

// TwitterHandler defines a new type which is just string
// type declaration
type TwitterHandler string

func (th *TwitterHandler) RedirectUrl() string {
	cleanedHandler := strings.TrimPrefix(string(*th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanedHandler)
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}

type Identifiable interface {
	ID() string
}

type Person struct {
	// embedded struct, so p.first, p.last and p.FullName() are accessible. So there is no need to add Name
	Name
	Age            int
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

func (p *Person) ID() string {
	return p.Identifiable.ID()
}

func (p *Person) SetHandler(handler TwitterHandler) error {
	if !strings.HasPrefix(string(handler), "@") {
		return errors.New("handler should start with '@'")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
