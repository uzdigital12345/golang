package contacts

import "fmt"

type ContactManager struct {
	contacts []Contact
}

func New() *ContactManager {
	cm := ContactManager{}
	cm.contacts = []Contact{}
	return &cm
}

type Contact struct {
	Id int
	Name string
	Gender string
	Phone int
	Mail string
}
func (c *ContactManager) Add(ct Contact) {
	c.contacts = append(c.contacts, ct)
}


func (c *ContactManager) Update(ct Contact) {
	c.contacts[ct.Id].Name = ct.Name
	c.contacts[ct.Id].Mail = ct.Mail
	c.contacts[ct.Id].Gender = ct.Gender
	c.contacts[ct.Id].Mail = ct.Mail
}

func (c *ContactManager) Delete(id int) {
	c.contacts = append(c.contacts	[:id], c.contacts[id+1:]...)
}

func (c *ContactManager) ListContact(id int) {
	fmt.Println(c.contacts[id].Id)
	fmt.Println(c.contacts[id].Name)
	fmt.Println(c.contacts[id].Gender)
	fmt.Println(c.contacts[id].Mail)
	fmt.Println(c.contacts[id].Phone)
	fmt.Println("-------------")
}

func (c *ContactManager) ListAll() []Contact{
	for i, _ := range c.contacts{
		c.ListContact(i)
	}
	return c.contacts
}
