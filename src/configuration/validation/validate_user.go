package validation

import (
	"encoding/json"
	"errors"
	"github.com/Guilermiz/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"

	enTranslator "github.com/go-playground/validator/v10/translations/en"
)

var (
	//Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		english := en.New()
		unTransl := ut.New(english, english)
		transl, _ = unTransl.GetTranslator("en")
		err := enTranslator.RegisterDefaultTranslations(val, transl)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid Field type", nil)
	} else if errors.As(validationErr, &jsonValidationErr) {
		var errorCauses []rest_err.Cause

		for _, er := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: er.Translate(transl),
				Field: er.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestError("There are some invalid fields ", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error Trying convert fields", nil)
	}
}
