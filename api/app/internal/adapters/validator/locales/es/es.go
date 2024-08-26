package es

import (
	"log"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomTranslations(v *validator.Validate, ut ut.Translator) (err error) {

	translations := []struct {
		tag         string
		translation string
		override    bool
	}{
		{
			tag:         "password",
			translation: "La contraseña debe tener entre 8 y 25 caracteres, al menos un número y al menos un carácter especial.",
			override:    true,
		},
		{
			tag:         "password_confirmation",
			translation: "La confirmación de contraseña no coincide con la contraseña",
			override:    true,
		},
	}
	for _, t := range translations {

		err = v.RegisterTranslation(t.tag, ut, registrationFunc(t.tag, t.translation, t.override), translateFunc)

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}
	return t
}
