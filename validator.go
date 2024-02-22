package gpc

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validation request from struct field
func Validator(s interface{}) (*FormatError, error) {
	var (
		translatorEnglish    locales.Translator      = en.New()
		universalTranslator  *ut.UniversalTranslator = ut.New(translatorEnglish, translatorEnglish)
		getTranslator, found                         = universalTranslator.GetTranslator("en")
	)

	if !found {
		return nil, errors.New("translator not found")
	}

	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return nil, fmt.Errorf("validator value not supported, because %v is not struct", reflect.TypeOf(s).Kind().String())
	}

	if res, err := keyExist(s); err != nil || res == 0 {
		return nil, fmt.Errorf("validator value can't be empty struct %v", s)
	}

	val := validator.New()
	err := val.Struct(s)

	if err == nil {
		return nil, nil
	}

	if err := en_translations.RegisterDefaultTranslations(val, getTranslator); err != nil {
		return nil, err
	}

	return formatError(err, getTranslator, s)

}

// Core module validator from https://github.com/go-playground/validator
func GoValidator() *validator.Validate {
	return validator.New()
}
