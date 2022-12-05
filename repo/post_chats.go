package repo

import (
	"github.com/kamva/mgm/v3"
	"github.com/kbiits/chat-app/entity"
)

func (r *Repo) PostChat(chat entity.Chat) (entity.Chat, error) {

	coll := mgm.Coll(&chat)
	err := coll.Create(&chat)
	if err != nil {
		return chat, err
	}

	return chat, nil
}
