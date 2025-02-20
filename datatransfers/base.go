package datatransfers

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error,omitempty"`
}

type ErrorDetail struct {
	Remark string `json:"remark"`
	Code   int    `json:"code"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	response := Response{Data: data}
	return c.Status(fiber.StatusOK).JSON(response)
}

func Error(c *fiber.Ctx, statusCode, code int, remark string) error {
	errorResponse := Response{
		Error: ErrorDetail{
			Remark: remark,
			Code:   code,
		},
	}

	return c.Status(statusCode).JSON(errorResponse)
}

type ListQueryParams struct {
	Limit  int
	Cursor int

	Nama         string
	NoRekening   string
	NIK          string
	NoHP         string
	IsCreateUser bool
}
