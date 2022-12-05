package handler

import "github.com/kbiits/chat-app/usecase"

type Handler struct {
	uc *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
