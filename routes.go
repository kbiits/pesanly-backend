package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/chat-app/handler"
)

func initializeRoutes(router *fiber.App, h *handler.Handler)  {
	
	router.Post("/register", h.RegisterUser)
	router.Post("/friends", h.GetFriends)
	router.Get("/get-chats", h.GetChats)
	router.Post("/post-chats", h.PostChat)
}