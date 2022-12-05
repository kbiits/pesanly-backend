package usecase

import (
	"strings"

	"github.com/kbiits/chat-app/entity"
	"github.com/kbiits/chat-app/utils"
)

func (uc *Usecase) RegisterUser(user entity.User) (err error) {
	user.PhoneNumber = "0" + strings.TrimLeft(utils.RegexPrefixPhoneNumber.ReplaceAllString(user.PhoneNumber, ""), "0")
	err = uc.repo.RegisterUser(user)
	return
}
