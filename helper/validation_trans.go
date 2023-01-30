package helper

import (
	"fmt"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en "github.com/go-playground/validator/v10/translations/en"
)

func TranslateValidationEnglish(validate *validator.Validate) ut.Translator {

	english := english.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en.RegisterDefaultTranslations(validate, trans)

	return trans

}

func TranslateError(err error, ut ut.Translator) (erss []error) {

	if err == nil {
		return nil
	}

	validationErrs := err.(validator.ValidationErrors)
	for _, fieldError := range validationErrs {
		translatedError := fmt.Errorf(fieldError.Translate(ut))
		erss = append(erss, translatedError)
	}

	return erss

}
