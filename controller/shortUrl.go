package controller

import (
	"github.com/gofiber/fiber/v2"
	"url_shortner/services"
)

type ShortURLController struct {
	ShortUrlMapService services.ShortUrlMapService
}

func (shortUrlMapService *ShortURLController) Create(ctx *fiber.Ctx) error {
	return nil
}
