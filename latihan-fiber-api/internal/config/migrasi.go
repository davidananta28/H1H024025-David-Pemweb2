package config

import (
	"log"

	"latihan-fiber-api/internal/model"
)

func Migrasi() {
	err := DB.AutoMigrate(&model.ProgramStudi{}, &model.Mahasiswa{},
		&model.Pengguna{}, &model.DaftarHitamToken{})
	if err != nil {
		log.Fatal("Gagal menjalankan migrasi: ", err)
	}

	log.Println("Migrasi selesai")
}

func DataAwal() {
	var jumlahProdi int64
	DB.Model(&model.ProgramStudi{}).Count(&jumlahProdi)

	if jumlahProdi == 0 {
		daftarProdi := []model.ProgramStudi{
			{Kode: "TK", Nama: "Teknik Komputer", Jenjang: "S1"},
			{Kode: "IF", Nama: "Informatika", Jenjang: "S1"},
			{Kode: "TE", Nama: "Teknik Elektro", Jenjang: "S1"},
		}

		DB.Create(&daftarProdi)
		log.Println("Data awal program studi dibuat")
	}

	var jumlahMatakuliah int64
	DB.Model(&model.Matakuliah{}).Count(&jumlahMatakuliah)

	if jumlahMatakuliah == 0 {
		daftarMatakuliah := []model.Matakuliah{
			{Kode: "IF101", Nama: "Algoritma dan Pemrograman", Sks: 3, Semester: 1},
			{Kode: "IF102", Nama: "Struktur Data", Sks: 3, Semester: 2},
			{Kode: "IF201", Nama: "Basis Data", Sks: 3, Semester: 3},
			{Kode: "IF202", Nama: "Pemrograman Web", Sks: 3, Semester: 4},
			{Kode: "IF301", Nama: "Kecerdasan Buatan", Sks: 3, Semester: 5},
		}

		DB.Create(&daftarMatakuliah)
		log.Println("Data awal matakuliah dibuat")
	}
}
