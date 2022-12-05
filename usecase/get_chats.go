package usecase

import (
	"time"

	"github.com/kbiits/chat-app/entity"
	"github.com/kbiits/chat-app/utils"
)

func (uc *Usecase) GetChats(forUser string, toUser string, after time.Time) ([]entity.Chat, error) {
	toUser = utils.CleanPhoneNumber(toUser)
	forUser = utils.CleanPhoneNumber(forUser)
	return uc.repo.GetChats(forUser, toUser, after)
}
