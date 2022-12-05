package repo

import (
	"context"

	"github.com/kamva/mgm/v3"
	"github.com/kbiits/chat-app/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repo) GetFriends(userPhone string, searchPhoneNumbers []string) ([]entity.User, error) {
	var currentUser entity.User
	coll := mgm.Coll(&currentUser)
	err := coll.First(bson.M{
		"phone_number": userPhone,
	}, &currentUser)
	if err != nil {
		return nil, err
	}

	arr := make(bson.A, 0)
	for _, v := range searchPhoneNumbers {
		arr = append(arr, v)
	}

	cursor, err := coll.Find(context.Background(), bson.M{
		"phone_number": bson.M{
			"$in": arr,
		},
	})
	if err != nil {
		return nil, err
	}

	result := make([]entity.User, 0)
	for cursor.Next(context.Background()) {
		var u entity.User
		err := cursor.Decode(&u)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}

	return result, nil
}
