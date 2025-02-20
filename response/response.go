package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error,omitempty"`
}

type ErrorDetail struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Success helper function
func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Data: data,
	})
}

// Error helper function
func Error(c *fiber.Ctx, statusCode, code int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Error: ErrorDetail{
			Message: message,
			Code:    code,
		},
	})
}
