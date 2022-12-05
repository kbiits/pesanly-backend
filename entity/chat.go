package entity

import "github.com/kamva/mgm/v3"

type Chat struct {
	mgm.DefaultModel `bson:",inline"`
	Content          string `json:"content" bson:"content"`
	From             string `json:"from" bson:"from"`
	To               string `json:"to" bson:"to"`
}
