// @file:        person.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Method test.

package person

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
