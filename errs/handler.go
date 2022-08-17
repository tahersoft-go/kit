package errs

import (
	"database/sql"
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/mhosseintaher/kit/dtp"
	"gorm.io/gorm"
)

func CustomHandler(err error, c echo.Context) {
	response := buildErrorResponse(err)
	if response.StatusCode() == http.StatusInternalServerError {
		c.Logger().Errorf("encountered internal server error: %v", err)
	}

	if response.Details == nil {
		response.Details = []dtp.H{}
	}
	if err := c.JSON(response.StatusCode(), dtp.H{
		"message": response.Message,
		"errors":  response.Details,
	}); err != nil {
		c.Logger().Errorf("failed writing error response: %v", err)
	}
}

// buildErrorResponse builds an error response from an error.
func buildErrorResponse(err error) ErrorResponse {
	switch err := err.(type) {
	case ErrorResponse:
		return err
	case validation.Errors:
		return InvalidInput(err)
	case *echo.HTTPError:
		switch err.Code {
		case http.StatusNotFound:
			return NotFound()
		default:
			return ErrorResponse{
				Status:  err.Code,
				Message: err.Error(),
			}
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound()
	}

	if errors.Is(err, gorm.ErrInvalidValue) {
		return NotFound("داده های وارد شده اشتباه می‌باشد")
	}

	if errors.Is(err, sql.ErrNoRows) {
		return NotFound()
	}

	return InternalServerError()
}
