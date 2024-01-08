
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data interface{}) error {
	if err := c.BindJSON(data); err != nil {
		return err
	}

	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}

func jsonResponse(status int, message string, data interface{}) gin.H {
	return gin.H{
		"Status":  status,
		"Message": message,
		"Data":    data,
	}
}

func (app *Config) writeJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, jsonResponse(status, "success", data))
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	c.JSON(statusCode, jsonResponse(statusCode, err.Error(), nil))
}
