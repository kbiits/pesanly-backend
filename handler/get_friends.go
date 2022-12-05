package handler

import "github.com/gofiber/fiber/v2"

type GetFriendsRequest struct {
	Phones []string `json:"phones"`
}

func (h *Handler) GetFriends(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}()

	loggedInUser := c.Query("myphone")
	if loggedInUser == "" {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": []string{},
		})
		return nil
	}

	var req GetFriendsRequest
	err = c.BodyParser(&req)
	if err != nil {
		return nil
	}

	users, err := h.uc.GetFriends(loggedInUser, req.Phones)
	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
	return nil
}
