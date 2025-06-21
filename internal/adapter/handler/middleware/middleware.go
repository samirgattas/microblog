package middleware

import (
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
			msg := gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			}
			statusCode := http.StatusInternalServerError
			switch errHand := err.(type) {
			case customerror.NotFoundError:
				msg = gin.H{
					"message": errHand.Error(),
					"status":  errHand.StatusCode,
				}
				statusCode = errHand.StatusCode
			case customerror.BadRequestError:
				msg = gin.H{
					"message": errHand.Error(),
					"status":  errHand.StatusCode,
				}
				statusCode = errHand.StatusCode
			}
			// Step5: Respond with a generic error message
			c.JSON(statusCode, msg)
		}
	}
}
