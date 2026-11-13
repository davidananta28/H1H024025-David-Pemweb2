package validasi

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var polaNim = regexp.MustCompile(`^H[0-9][A-Za-z0-9]*$`)

func init() {
	err := pemeriksa.RegisterValidation("nim", func(fl validator.FieldLevel) bool {
		return polaNim.MatchString(fl.Field().String())
	})
	if err != nil {
		panic("gagal mendaftarkan aturan validasi nim: " + err.Error())
	}
}
