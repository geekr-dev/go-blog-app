package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Key     string
	Message string
}

type ValidationErrors []*ValidationError

func (v *ValidationError) Error() string {
	return v.Message
}

func (v ValidationErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValidate(c *gin.Context, v interface{}) (bool, ValidationErrors) {
	var errs ValidationErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidationError{
				Key:     key,
				Message: value,
			})
		}
		return true, errs
	}
	return false, nil
}
