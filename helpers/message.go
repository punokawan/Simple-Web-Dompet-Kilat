package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Code    int         `json:"code,omitempty"`    // status code HTTP
	Succes  bool        `json:"succes,omitempty"`  // is endpoint error? (T/F)
	Message string      `json:"message,omitempty"` // Message from handler
	Data    interface{} `json:"data,omitempty"`    // Data from resource
}

func ResponseMsg(c *fiber.Ctx, code int, success bool, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    code,
		Succes:  success,
		Message: msg,
		Data:    data,
	}
	return c.Status(code).JSON(resPonse)
}
