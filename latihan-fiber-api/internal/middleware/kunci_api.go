package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// KunciApi memeriksa header X-Api-Key dan menolak permintaan
// dengan status 401 apabila nilainya tidak sesuai dengan kunci yang diberikan.
func KunciApi(kunci string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if c.Get("X-Api-Key") != kunci {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Kunci API tidak valid atau tidak diberikan",
			)
		}
		return c.Next()
	}
}

