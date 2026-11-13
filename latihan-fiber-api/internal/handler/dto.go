package handler

type PermintaanMahasiswa struct {
	ProgramStudiID uint    `json:"program_studi_id" validate:"required,gte=1"`
	Nim            string  `json:"nim" validate:"required,min=5,max=20,nim"` // BARU: ,nim
	Nama           string  `json:"nama" validate:"required,min=3,max=100"`
	Email          string  `json:"email" validate:"required,email,max=100"`
	Angkatan       int     `json:"angkatan" validate:"required,gte=2000,lte=2100"`
	Ipk            float64 `json:"ipk" validate:"gte=0,lte=4"`
	Aktif          *bool   `json:"aktif"`
}

type PermintaanRegistrasi struct {
	Nama      string `json:"nama" validate:"required,min=3,max=100"`
	Email     string `json:"email" validate:"required,email,max=100"`
	KataSandi string `json:"kata_sandi" validate:"required,min=8"`
}

type PermintaanLogin struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type PermintaanMatakuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Sks      int    `json:"sks"`
	Semester int    `json:"semester"`
}
