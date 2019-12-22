package models

import (
	"fmt"
	"testing"
)

var (
	c ContactManagerInterface
)

func TestNewContactManagerInterface(t *testing.T) {

	c, err = NewContactManagerInterface()
	if err != nil {
		t.Error("Can't connect to database")
	}
}

func TestGormDb_Add(t *testing.T) {
	t1 := Contact {
		Age: 19,
		Name: "Enver",
		Gender: "Male",
		PhoneNumber:"991562536",
	}
	t2 := Contact {
		Age: 85,
		Name: "Boboy",
		Gender: "Male",
		PhoneNumber:"9415625315",
	}
	t3 := Contact {
		Age: 15,
		Name: "Abdurahmon",
		Gender: "Male",
		PhoneNumber:"941563245",
	}
	t4 := Contact {
		Age: 25,
		Name: "Sardor",
		Gender: "Male",
		PhoneNumber:"941563515",
	}
	if err != nil {
		t.Error("Can't Add", err)
	}
	err = c.Add(t1)
	err = c.Add(t2)
	err = c.Add(t3)
	err = c.Add(t4)

	if err != nil {
		t.Error("Can't Add", err)
	}

}

func TestGormDb_UpdateName(t *testing.T) {

   err = c.UpdateName(1,"jahongir")

	if err != nil {
		t.Error("Can't Update", err)
	}
}

func TestGormDb_Delete(t *testing.T) {
   err = c.Delete(2)
	if err != nil {
		t.Error("Can't Update", err)
	}
}

func TestGormDb_ListAll(t *testing.T) {
	err = c.ListAll()
	if err != nil {
		t.Error("Can't List All", err)
	}
}

func TestGormDb_GetAll(t *testing.T) {
	_,err = c.GetAll()
	if err != nil {
		t.Error("Can't Get All", err)
	}
}

func TestGormDb_GetPaging(t *testing.T) {
	var contact []Contact
	fmt.Println("Paging print")
	contact,err = c.GetPaging(2,5)
	fmt.Println(contact)
	if err != nil {
		t.Error("Can't get paging")
	}
}


