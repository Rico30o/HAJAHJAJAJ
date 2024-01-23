package tracealert

import (
	"instapay/db"
	model "instapay/model/response"

	"github.com/gofiber/fiber/v2"
)

func TransactionIDmatches(c *fiber.Ctx) error {
	IDMatches := &model.TransacIDmatches{}
	if parsErr := c.BodyParser(IDMatches); parsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": parsErr.Error(),
		})
	}

	ID := &model.ResponseIDmatches{}
	if err := db.DB.Debug().Raw(`SELECT * FROM public.martchingid`).Scan(&ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"data": err.Error()})
	}

	return c.JSON(ID)
}
