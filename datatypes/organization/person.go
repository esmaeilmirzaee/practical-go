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

type Name struct {
	first string
	last  string
}

type Identifiable interface {
	ID() string
}

type Person struct {
	Name           Name
	Age            int
	twitterHandler TwitterHandler
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

func (p *Person) ID() string {
	return p.Name.FullName()
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
