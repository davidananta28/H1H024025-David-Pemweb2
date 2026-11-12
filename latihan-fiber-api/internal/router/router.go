package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
)

func Daftarkan(app *fiber.App, kunciApi string) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Endpoint Mahasiswa
	mahasiswa := api.Group("/mahasiswa")
	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/terhapus", handler.DaftarMahasiswaTerhapus)
	mahasiswa.Post("/", handler.BuatMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)
	mahasiswa.Put("/:id/pulihkan", handler.PulihkanMahasiswa)
	mahasiswa.Post("/:id/restore", handler.PulihkanMahasiswa)
	mahasiswa.Put("/:id/restore", handler.PulihkanMahasiswa)

	// Endpoint Matakuliah
	matakuliah := api.Group("/matakuliah")
	matakuliah.Get("/", handler.DaftarMatakuliah)
	matakuliah.Post("/", handler.BuatMatakuliah)
	matakuliah.Get("/:id", handler.DetailMatakuliah)
	matakuliah.Put("/:id", handler.PerbaruiMatakuliah)
	matakuliah.Delete("/:id", handler.HapusMatakuliah)

	// Endpoint Program Studi
	programStudi := api.Group("/program-studi")
	programStudi.Get("/", handler.DaftarProgramStudi)
	programStudi.Get("/:id", handler.DetailProgramStudi)
}
