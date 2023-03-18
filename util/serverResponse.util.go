package util

import "github.com/gofiber/fiber/v2"

type Response[D any] struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Data    D      `json:"data"`
}

func GenerateResponse[D any](context *fiber.Ctx, data D, success bool, reason string) error {
	responseData := Response[D]{
		Success: success,
		Reason:  reason,
		Data:    data,
	}

	return context.Status(fiber.StatusOK).JSON(responseData)
}
