package errorhandling

import (
	"instapay/model"

	"github.com/gofiber/fiber/v2"
)

// 400
func Bad_Request(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"Source":      "ALERT_FINANCIAL_CRIME",
					"ReasonCode":  "BAD_REQUEST",
					"Description": "We could not handle your request",
					"Recoverable": false,
					"Details":     details,
				},
			},
		},
	})
}

// 403
func Permision_Denied(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"Errors": map[string]interface{}{
			"Error": []map[string]interface{}{
				{
					"Source":      "Gateway",
					"ReasonCode":  "PERMISSION_DENIED",
					"Description": "Invalid customer for third party",
					"Recoverable": false,
					"Details":     nil,
				},
			},
		},
	})
}

// 429
func Rate_Limit_Exceeded(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(model.ErrorResponses{
		Errors: struct {
			Error []model.ErrorDetail `json:"Error"`
		}{
			Error: []model.ErrorDetail{{
				Source:      "Gateway",
				ReasonCode:  "RATE_LIMIT_EXCEEDED",
				Description: "You have exceeded the service rate limit. Maximum allowed: ${rate_limit.output} TPS",
				Recoverable: true,
				Details:     nil,
			}},
		},
	})

}

// 409
func Conflict(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusConflict).JSON(model.ErrorResponses{
		Errors: struct {
			Error []model.ErrorDetail `json:"Error"`
		}{
			Error: []model.ErrorDetail{{
				Source:      "FEEDBACK_FINANCIAL_CRIME",
				ReasonCode:  "CONFLICT",
				Description: "Alert ID does not match the specified entity",
				Recoverable: false,
			}},
		},
	})

}

// 422
func Unprocessable_Entity(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(model.ErrorResponses{
		Errors: struct {
			Error []model.ErrorDetail `json:"Error"`
		}{
			Error: []model.ErrorDetail{{
				Source:      "FEEDBACK_FINANCIAL_CRIME",
				ReasonCode:  "UNPROCESSABLE_ENTITY",
				Description: "Expects a single JSON object and not an array",
				Recoverable: false,
			}},
		},
	})
}

// 415
func Unsupported_Media_Type(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusUnsupportedMediaType).JSON(model.ErrorResponses{
		Errors: struct {
			Error []model.ErrorDetail `json:"Error"`
		}{
			Error: []model.ErrorDetail{{
				Source:      "FEEDBACK_FINANCIAL_CRIME",
				ReasonCode:  "UNSUPPORTED_MEDIA_TYPE",
				Description: "Unsupported media type",
				Recoverable: false,
				Details:     "The request media type 'application/x-www-form-urlencoded' is not supported by this resource",
			}},
		},
	})
}

// 405
func Method_Not_Allowed(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"Source":      "TRACE_FINANCIAL_CRIME",
					"ReasonCode":  "METHOD_NOT_ALLOWED",
					"Description": "Only POST method allowed",
					"Recoverable": false,
					"Details":     nil,
				},
			},
		},
	})
}

// 404
func Url_Not_Found(c *fiber.Ctx, details string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"Errors": fiber.Map{
			"Error": []fiber.Map{
				{
					"Source":      "Gateway",
					"ReasonCode":  "NOT_FOUND",
					"Description": "URL not found",
					"Recoverable": false,
					"Details":     nil,
				},
			},
		},
	})
}
