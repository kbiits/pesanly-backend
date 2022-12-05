package repo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kbiits/chat-app/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) RegisterUser(user entity.User) (err error) {
	coll := mgm.Coll(&user)

	err = coll.First(bson.M{
		"phone_number": user.PhoneNumber,
	}, &user)
	if err == nil {
		return nil
	}

	err = coll.Create(&user)
	return
}
