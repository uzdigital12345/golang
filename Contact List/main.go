package main

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
	c.contacts = append(c.contacts[:id], c.contacts[id+1:]...)
}

func (c *ContactManager) ListContact(id int) {
	fmt.Println(c.contacts[id].Id)
	fmt.Println(c.contacts[id].Name)
	fmt.Println(c.contacts[id].Gender)
	fmt.Println(c.contacts[id].Mail)
	fmt.Println(c.contacts[id].Phone)
	fmt.Println("-------------")
}

func (c *ContactManager) ListAll() {
	for i, _ := range c.contacts{
		c.ListContact(i)
	}
}


func main() {
	c := New()
	ct:=Contact{0,"Khamidullokh","male",901233323,"kholikov.x@gmail.com"}
	ct1:=Contact{1,"Temur","male",951251515,"temur@gmail.com"}
	ct2:=Contact{2,"Akbar","male",912555225,"Urazbaev@gmail.com"}
	ct3:=Contact{3,"Hamidulloh","male",1561556,"Hamidulloh@gmail.com"}
	c.Add(ct)
	c.Add(ct1)
	c.Add(ct2)
	c.Add(ct3)
	c.ListAll()
	ct4:=Contact{2,"Dilmurod","male",941562156,"dilmurod@gmail.com"}
	c.Update(ct4)
	c.ListAll()
	c.ListContact(0)
}
