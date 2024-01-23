package routes

import (
	"instapay/model/errors"
	model "instapay/model/response"

	"github.com/gofiber/fiber/v2"
)

func checkHealth() model.ResponseModel {
	return model.ResponseModel{
		RetCode: "100",
		Message: "Request success!",
		Data: errors.ErrorModel{
			Message:   "Service is available!",
			IsSuccess: true,
			Error:     nil,
		},
	}
}

func CheckServiceHealth(c *fiber.Ctx) error {
	health := checkHealth()
	response := errors.ErrorModel{}
	response = health.Data.(errors.ErrorModel)
	if !response.IsSuccess {
		return c.JSON(health)
	}
	return c.JSON(health)
}
