package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslations "github.com/go-playground/validator/v10/translations/en"
)

var verifier = validator.New(validator.WithRequiredStructEnabled())
var translator ut.Translator

func init() {
	english := en.New()
	uni := ut.New(english, english)

	var found bool
	translator, found = uni.GetTranslator("en")

	if !found {
		panic("no translator found for english language.")
	}

	err := englishTranslations.RegisterDefaultTranslations(verifier, translator)
	if err != nil {
		panic("can not register detail english translations to the validator.")
	}

	// register validations
	registerValidations()
}

type ErrValidation struct {
	Msg     string
	Details map[string]string
	Raw     error
}

func (ev *ErrValidation) Error() string {
	return ev.Msg
}

func ValidateStruct(data any) *ErrValidation {
	err := verifier.Struct(data)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return &ErrValidation{
				Msg: "invalid request body",
				Raw: err,
			}
		}

		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if ok {
			var validationErrors = make(map[string]string, len(errs))
			for _, e := range errs {
				fmt.Println(e.Field(), e.Error())
				validationErrors[e.Field()] = e.Translate(translator)
			}

			return &ErrValidation{
				Msg:     "invalid request body",
				Details: validationErrors,
			}
		}

		return &ErrValidation{
			Msg: "invalid request body",
			Raw: err,
		}
	}

	return nil
}
