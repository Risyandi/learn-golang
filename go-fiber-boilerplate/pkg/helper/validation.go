package helper

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

type (
	XValidator struct {
		validator *validator.Validate
	}
	ErrorValidate struct {
		Error        bool
		FailedField  string
		Tag          string
		Param        string
		Value        interface{}
		TranslateKey *i18n.LocalizeConfig
	}
)

var validate = validator.New()

func NewValidator() *XValidator {

	// Custom validation function
	validate.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		// Example User.Age needs to fit our needs, 12-18 years old.
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
	validate.RegisterValidation("cleanmoji", func(f1 validator.FieldLevel) bool {
		return !utils.EmojiRx.MatchString(f1.Field().String())
	})

	validate.RegisterValidation("gtdate", gtDate)
	validate.RegisterValidation("lteparent", lteParent)

	return &XValidator{
		validator: validate,
	}
}

func (v XValidator) ValidateRequest(ctx *fiber.Ctx, data interface{}) error {
	if errs := v.validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			err.msgValidation()
			message := fiberi18n.MustLocalize(ctx, err.TranslateKey)

			errMsgs = append(errMsgs, message)
		}

		return errors.New(strings.Join(errMsgs, " , "))
	}
	return nil
}

func (v XValidator) validate(data interface{}) []ErrorValidate {
	validationErrors := []ErrorValidate{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorValidate

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Param = err.Param()       // Export validation rule parameter (if any)
			elem.TranslateKey = &i18n.LocalizeConfig{
				MessageID: constants.ResponseErrorMessage,
			}
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (err *ErrorValidate) msgValidation() {
	field := err.FailedField
	// Change Field Name
	switch err.Tag {
	case "required":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorRequired,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	case "email":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorEmail,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	case "number":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorNumber,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	case "unique":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorUnique,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	case "gte":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorGte,
			TemplateData: map[string]string{
				"field": field,
				"gte":   err.Param,
			},
		}
	case "gtefield":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorGte,
			TemplateData: map[string]string{
				"field": field,
				"gte":   err.Param,
			},
		}
	case "gtfield":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorGt,
			TemplateData: map[string]string{
				"field": field,
				"gte":   err.Param,
			},
		}
	case "gtdate":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorGt,
			TemplateData: map[string]string{
				"field": field,
				"gte":   err.Param,
			},
		}
	case "lte":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLte,
			TemplateData: map[string]string{
				"field": field,
				"lte":   err.Param,
			},
		}
	case "ltefield":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLte,
			TemplateData: map[string]string{
				"field": field,
				"lte":   err.Param,
			},
		}
	case "lteparent":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLte,
			TemplateData: map[string]string{
				"field": field,
				"lte":   err.Param,
			},
		}
	case "ltfield":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLt,
			TemplateData: map[string]string{
				"field": field,
				"lte":   err.Param,
			},
		}
	case "ltdate":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLt,
			TemplateData: map[string]string{
				"field": field,
				"lte":   err.Param,
			},
		}
	case "min":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorMin,
			TemplateData: map[string]string{
				"field":  field,
				"length": err.Param,
			},
		}
	case "max":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorMax,
			TemplateData: map[string]string{
				"field":  field,
				"length": err.Param,
			},
		}
	case "startswith":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorStartswith,
			TemplateData: map[string]string{
				"field": field,
				"start": err.Param,
			},
		}
	case "len":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorLen,
			TemplateData: map[string]string{
				"field":  field,
				"length": err.Param,
			},
		}
	case "oneof":
		choices := utils.ReplaceAllChar(err.Param, " ", "/ ")
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorOneof,
			TemplateData: map[string]string{
				"field": field,
				"in":    choices,
			},
		}

	case "uuid4":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorUUID,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	case "cleanmoji":
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrorCleanEmoji,
			TemplateData: map[string]string{
				"field": field,
			},
		}
	default:
		err.TranslateKey = &i18n.LocalizeConfig{
			MessageID: constants.ValidationErrors,
			TemplateData: map[string]string{
				"field": err.FailedField,
				"tag":   err.Tag,
			},
		}
	}
	// break
	return
}

var gtDate validator.Func = func(fl validator.FieldLevel) bool {
	Indonesia, _ := time.LoadLocation("Asia/Jakarta")
	structValue := fl.Parent()
	param := fl.Param()

	dateToTield := structValue.FieldByName(param)
	if !dateToTield.IsValid() || dateToTield.Kind() != reflect.Int {
		return false
	}

	dateTo := dateToTield.Int()
	value := fl.Field().Int()
	if value != 0 {
		date := time.Unix(value, 0).In(Indonesia)
		today := utils.GetLocalDateTime()
		if dateTo != 0 {
			today = time.Unix(dateTo, 0).In(Indonesia)
			return !date.After(today)
		}

		if date.Day() < today.Day() {
			return false
		}
	}
	return true
}

var lteParent validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	param := fl.Param()

	parent := fl.Parent()
	parentValue := parent.FieldByName(param).Int()

	return value <= parentValue
}
