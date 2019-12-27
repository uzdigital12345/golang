package models

import (
	"fmt"
	"testing"
)


func TestSqlx_NewContactManagerInterface(t *testing.T) {

	_,err = NewContactManagerInterface()
	if err != nil {
		fmt.Println("Can't connect to database")
	}
}

func TestSqlx_Add(t *testing.T) {
	c ,err:= NewContactManagerInterface()
	t1 := Contact{1,19,"Enver","Male",991562536}
	t2 := Contact{2,15,"Abdurahmon","Male",941563245}
	t3 := Contact{3,85,"Boboy","Male",9415625315}
	t4 := Contact{4,25,"Sardor","Male",941563515}
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

