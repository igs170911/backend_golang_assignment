package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var translator ut.Translator

func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		var messages string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err2 := range validationErrors {
				messages += err2.Translate(translator) + ", "
			}

			// 過濾最後字尾 ", "
			messages = messages[:len(messages)-2]
		}

		return errors.New(messages)
	}

	return nil
}
