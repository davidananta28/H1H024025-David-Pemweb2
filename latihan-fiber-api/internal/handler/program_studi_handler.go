package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/repository"
)

func DaftarProgramStudi(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaProgramStudi()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data program studi")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

func DetailProgramStudi(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	prodi, err := repository.AmbilProgramStudiPerId(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Program studi tidak ditemukan")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data program studi")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   prodi,
	})
}
