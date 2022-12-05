package entity

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	PhoneNumber      string `json:"phone_number" bson:"phone_number"`
}

func (model *User) CollectionName() string {
	return "users"
}
