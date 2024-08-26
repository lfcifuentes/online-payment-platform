package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	custom_en "github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator/locales/en"
	custom_es "github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator/locales/es"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/validator/rules"
)

var CurrentRequestClient *string

type ValidatorAdapter struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

// NewValidatorAdapter return a new instance of the validator adapter
func NewValidatorAdapter() *ValidatorAdapter {
	esTranslator := es.New()
	uni := ut.New(esTranslator, esTranslator)
	trans, _ := uni.GetTranslator("es")

	validate := validator.New()
	err := es_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil
	}
	// register custom translations
	err = RegisterTranslations("es", validate, trans)
	if err != nil {
		return nil
	}
	// Register the tag name to be used in the error messages
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if fld.Tag.Get("documentation") != "" {
			return fld.Tag.Get("documentation")
		}
		return fld.Name
	})

	if err = validate.RegisterValidation("password", rules.CustomPasswordValidator); err != nil {
		fmt.Printf("Ups, an error occurred while registering the custom validation: %s", err)
		return nil
	}

	if err = validate.RegisterValidation("password_confirmation", rules.CustomPasswordConfirmationValidator); err != nil {
		fmt.Printf("Ups, an error occurred while registering the custom validation: %s", err)
		return nil
	}

	return &ValidatorAdapter{
		Validate:   validate,
		Translator: trans,
	}
}

func RegisterTranslations(language string, v *validator.Validate, trans ut.Translator) error {
	var err error
	if language == "es" {
		err = custom_es.RegisterCustomTranslations(v, trans)
	}
	if language == "en" {
		err = custom_en.RegisterCustomTranslations(v, trans)
	}

	return err
}

func (v *ValidatorAdapter) ValidateStruct(s interface{}) validator.ValidationErrorsTranslations {

	err := v.Validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errs.Translate(v.Translator)
	}

	return nil
}
