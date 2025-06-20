package middleware

import (
	"errors"
	"microblog/internal/core/lib/customerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler captures errors and returns a consistent JSON error response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a custom message
			var customErr customerror.CustomError
			if errors.As(err, &customErr) {
				c.JSON(customErr.StatusCode, gin.H{
					"message": customErr.Error(),
					"status":  customErr.StatusCode,
				})
				return
			}
			// Step5: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		}
	}
}
