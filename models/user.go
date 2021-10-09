package models

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Dob              string `json:"dob" bson:"dob"`
	Address          string `json:"address" bson:"address"`
	Description      string `json:"description" bson:"description"`
}

func NewUser(name string, dob string, address string, description string) *User {
	return &User{
		Name:        name,
		Dob:         dob,
		Address:     address,
		Description: description,
	}
}
