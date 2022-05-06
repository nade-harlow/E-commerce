package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"log"
)

var validate *validator.Validate
var Translate ut.Translator

func init() {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	Translate, _ = uni.GetTranslator("en")
	err := enTranslations.RegisterDefaultTranslations(validate, Translate)
	if err != nil {
		log.Println(err)
	}
}

func ValidateStruct(data interface{}) []string {
	err := validate.Struct(data)
	log.Println("struct validation error: ", TranslateError(err, Translate))
	return TranslateError(err, Translate)
}

func ValidateVariable(variable interface{}, tag string) []string {
	err := validate.Var(variable, tag)
	if err != nil {
		log.Println(err)
		return TranslateError(err, Translate)
	}
	return nil
}

func TranslateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr.Error())
	}
	return errs
}
