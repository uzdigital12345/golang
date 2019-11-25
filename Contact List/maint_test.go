package main

import (
	"testing"
)

var cm *ContactManager

func TestContactManager_Add(t *testing.T) {
	cm = New()

	cm.Add(Contact{1,"Hamidulloh","Male",901233326,"kholikov.x@gmail.com"})

	if (len(cm.contacts)==0) {
		t.Error("Contact not added")
	}
}

func TestContactManager_Update(t *testing.T) {
	cm.Update(Contact{1,"Temur","Male",913525116,"temur@gmail.com"})

	if cm.contacts[1].Name != "Temur" {
			t.Error("Contact not updated")
	}
}

func TestContactManager_Delete(t *testing.T) {
	a:=len(cm.contacts)
	cm.Delete(1)
	if (len(cm.contacts)!=a-1) {
		t.Error("Contact not deleted")
	}
}