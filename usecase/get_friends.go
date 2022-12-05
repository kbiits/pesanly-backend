package usecase

import (
	"strings"

	"github.com/kbiits/chat-app/entity"
	"github.com/kbiits/chat-app/utils"
)

func (uc *Usecase) GetFriends(userPhone string, searchPhoneNumbers []string) ([]entity.User, error) {
	var newSearchPhoneNumbers []string
	for _, v := range searchPhoneNumbers {
		newPhoneNumber := "0" + strings.TrimLeft(utils.RegexPrefixPhoneNumber.ReplaceAllString(v, ""), "0")
		newSearchPhoneNumbers = append(newSearchPhoneNumbers, newPhoneNumber)
	}

	return uc.repo.GetFriends(userPhone, newSearchPhoneNumbers)
}
