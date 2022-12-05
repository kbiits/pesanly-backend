package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/chat-app/entity"
)

func (h *Handler) RegisterUser(ctx *fiber.Ctx) (error) {
	var err error
	defer func() {
		if err != nil {
			ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}()

	var user entity.User
	err = ctx.BodyParser(&user)
	if err != nil {
		return nil
	}

	if user.Name == "" || user.PhoneNumber == "" {
		err = fmt.Errorf("name or phone number can't be empty")
		return nil
	}

	err = h.uc.RegisterUser(user)
	if err != nil {
		return nil
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
	return nil
}
