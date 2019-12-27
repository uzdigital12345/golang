package models

import (
	pt "github.com/uzdigital12345/golang/models_grpc/proto"
)

type ContactManagerInterface interface {
	Add(c *pt.Contact) error
	Update(i int, c pt.Contact) error
	Delete(i int) error
	ListAll() error
	GetAll() ([]pt.Contact, error)
}
