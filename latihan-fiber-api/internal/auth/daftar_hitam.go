package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func HashToken(token string) string {
	ringkasan := sha256.Sum256([]byte(token))

	return hex.EncodeToString(ringkasan[:])
}

func CabutToken(token string, kedaluwarsa time.Time) error {
	entri := model.DaftarHitamToken{
		TokenHash:       HashToken(token),
		KedaluwarsaPada: kedaluwarsa,
	}

	return config.DB.Create(&entri).Error
}

func TokenDicabut(token string) bool {
	var jumlah int64

	err := config.DB.Model(&model.DaftarHitamToken{}).
		Where("token_hash = ?", HashToken(token)).
		Count(&jumlah).Error
	if err != nil {
		return true // gagal memeriksa: tolak demi keamanan
	}

	return jumlah > 0
}

func BersihkanDaftarHitam() {
	config.DB.Where("kedaluwarsa_pada < ?", time.Now()).Delete(&model.DaftarHitamToken{})
}
