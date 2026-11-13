package handler

import (
	"time"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func Segarkan(c fiber.Ctx) error {
	idPengguna, cocok := c.Locals("id_pengguna").(uint)
	if cocok == false {
		return fiber.NewError(fiber.StatusUnauthorized, "Identitas pengguna tidak ditemukan pada token")
	}

	var pengguna model.Pengguna

	err := config.DB.First(&pengguna, idPengguna).Error
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Pengguna pada token sudah tidak ada")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(pengguna.ID, pengguna.Peran, konf.JwtSecret, konf.JwtTtlMenit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menerbitkan token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil disegarkan",
		"data": fiber.Map{
			"token": token,
			"tipe":  "Bearer",
		},
	})
}
func Logout(c fiber.Ctx) error {
	token, cocokToken := c.Locals("token").(string)
	kedaluwarsa, cocokWaktu := c.Locals("token_kedaluwarsa").(time.Time)

	if cocokToken == false || cocokWaktu == false {
		return fiber.NewError(fiber.StatusUnauthorized, "Token tidak ditemukan pada permintaan")
	}

	err := auth.CabutToken(token, kedaluwarsa)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mencabut token")
	}

	auth.BersihkanDaftarHitam()

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil, token telah dicabut",
	})
}
