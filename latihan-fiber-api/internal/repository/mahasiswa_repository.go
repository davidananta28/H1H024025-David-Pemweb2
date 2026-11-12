package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

type FilterMahasiswa struct {
	Cari          string
	Angkatan      int
	Halaman       int
	PerHalaman    int
	HanyaTerhapus bool
}

func Ambil(filter FilterMahasiswa) ([]model.Mahasiswa, int64, error) {
	var daftar []model.Mahasiswa
	var total int64

	kueri := config.DB.Model(&model.Mahasiswa{})
	if filter.HanyaTerhapus {
		kueri = config.DB.Unscoped().Model(&model.Mahasiswa{}).Where("deleted_at IS NOT NULL")
	}

	if filter.Cari != "" {
		pola := "%" + filter.Cari + "%"
		kueri = kueri.Where("nama LIKE ? OR nim LIKE ?", pola, pola)
	}

	if filter.Angkatan > 0 {
		kueri = kueri.Where("angkatan = ?", filter.Angkatan)
	}

	err := kueri.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	lompat := (filter.Halaman - 1) * filter.PerHalaman

	err = kueri.Preload("ProgramStudi").
		Order("nama ASC").
		Limit(filter.PerHalaman).
		Offset(lompat).
		Find(&daftar).Error

	if err != nil {
		return nil, 0, err
	}

	return daftar, total, nil
}

func AmbilPerId(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.Preload("ProgramStudi").First(&mahasiswa, id).Error

	return mahasiswa, err
}

func AmbilPerIdUnscoped(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.Unscoped().Preload("ProgramStudi").First(&mahasiswa, id).Error

	return mahasiswa, err
}

func Simpan(mahasiswa *model.Mahasiswa) error {
	return config.DB.Create(mahasiswa).Error
}

func Perbarui(mahasiswa *model.Mahasiswa) error {
	return config.DB.Save(mahasiswa).Error
}

func Hapus(mahasiswa *model.Mahasiswa) error {
	return config.DB.Delete(mahasiswa).Error
}

func Pulihkan(id uint) error {
	return config.DB.Unscoped().Model(&model.Mahasiswa{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

// Alias untuk konsistensi penamaan
var (
	AmbilMahasiswa      = Ambil
	AmbilMahasiswaPerId = AmbilPerId
	SimpanMahasiswa     = Simpan
	PerbaruiMahasiswa   = Perbarui
	HapusMahasiswa      = Hapus
)
