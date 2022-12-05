package usecase

import (
	"github.com/kbiits/chat-app/repo"
)

type Usecase struct {
	repo *repo.Repo
}

func NewUsecase(r *repo.Repo) *Usecase {
	return &Usecase{
		repo: r,
	}
}
