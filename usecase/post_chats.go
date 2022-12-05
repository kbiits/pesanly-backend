package usecase

import (
	"github.com/kbiits/chat-app/entity"
	"github.com/kbiits/chat-app/utils"
)

func (uc *Usecase) PostChat(req entity.Chat) (entity.Chat, error) {
	req.From = utils.CleanPhoneNumber(req.From)
	req.To = utils.CleanPhoneNumber(req.To)

	chat := entity.Chat{
		Content: req.Content,
		From:    req.From,
		To:      req.To,
	}
	return uc.repo.PostChat(chat)
}
