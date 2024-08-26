package en

import (
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
			translation: "The password must be between 12 and 25 characters long, contain at least one number, and at least one special character.",
			override:    true,
		},
		{
			tag:         "password_confirmation",
			translation: "The password confirmation does not match the password.",
			override:    true,
		},
	}
	for _, t := range translations {

		if err = ut.Add(t.tag, t.translation, t.override); err != nil {
			return err
		}

		if err != nil {
			return
		}
	}
	return
}
