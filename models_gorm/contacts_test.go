package models

import (
	"testing"
)

var (
	c,_ = NewContactManagerInterface()
)

func TestNewContactManagerInterface(t *testing.T) {

	_,err = NewContactManagerInterface()
	if err != nil {
		t.Error("Can't connect to database")
	}
}

func TestGormDb_Add(t *testing.T) {
	t1 := Contact{1,19,"Enver","Male","991562536"}
	t2 := Contact{2,15,"Abdurahmon","Male","941563245"}
	t3 := Contact{3,85,"Boboy","Male","9415625315"}
	t4 := Contact{4,25,"Sardor","Male","941563515"}
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


