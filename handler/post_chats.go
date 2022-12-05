package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/chat-app/entity"
)

type PostChatRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func (h *Handler) PostChat(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}()

	var req PostChatRequest
	err = c.BodyParser(&req)
	if err != nil {
		return nil
	}

	chat, err := h.uc.PostChat(entity.Chat{
		Content: req.Content,
		From:    req.From,
		To:      req.To,
	})
	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": chat,
	})
	return nil
}
