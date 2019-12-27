package postgres

import (
	pt "github.com/uzdigital12345/golang/models_grpc/proto"
)

type ContactManagerInterface interface {
	Add(c *pt.Contact) error
	Update(i int64, c *pt.Contact) error
	Delete(i int64) error
	GetAll() ([]*pt.Contact, error)
}
