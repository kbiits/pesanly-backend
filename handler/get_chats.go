package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/chat-app/entity"
)

func (h *Handler) GetChats(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}()

	forUser := c.Query("for_user")
	if forUser == "" {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": []entity.Chat{},
		})
		return nil
	}

	toUser := c.Query("to_user")
	if toUser == "" {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": []entity.Chat{},
		})
		return nil
	}

	var afterTime time.Time
	after := c.Query("after")
	layout := "2006-01-02T15:04:05-07:00"
	if after != "" {
		afterTime, err = time.Parse(layout, after)
		if err != nil {
			return nil
		}
	}

	chats, err := h.uc.GetChats(forUser, toUser, afterTime)
	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": chats,
	})
	return nil
}
