package controller

import (
	"github.com/gofiber/fiber/v2"
	"url_shortner/services"
)

type UserController struct {
	UserService *services.UserService
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	return nil
}
