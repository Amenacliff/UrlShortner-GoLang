package controller

import (
	"github.com/gofiber/fiber/v2"
	"url_shortner/services"
	"url_shortner/util"
)

type ShortURLController struct {
	ShortUrlMapService services.ShortUrlsService
}

func (shortUrlMapService *ShortURLController) Create(ctx *fiber.Ctx) error {
	return util.GenerateResponse(ctx, "", true, "")
}
