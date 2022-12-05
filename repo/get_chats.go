package repo

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/kbiits/chat-app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repo) GetChats(forUser string, toUser string, after time.Time) ([]entity.Chat, error) {
	coll := mgm.Coll(&entity.Chat{})

	filter := bson.M{
		"from": bson.M{
			"$in": bson.A{forUser, toUser},
		},
		"to": bson.M{
			"$in": bson.A{forUser, toUser},
		},
	}

	if (after != time.Time{}) {
		filter["created_at"] = bson.M{
			"$gt": after,
		}
	}

	cursor, err := coll.Find(context.Background(), filter, &options.FindOptions{
		Sort: bson.M{
			"created_at": -1,
		},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return make([]entity.Chat, 0), nil
		}
		return nil, err
	}

	var result []entity.Chat
	for cursor.Next(context.Background()) {
		var chat entity.Chat
		err := cursor.Decode(&chat)
		if err != nil {
			return nil, err
		}

		result = append(result, chat)
	}

	return result, nil
}
