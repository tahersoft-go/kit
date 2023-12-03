package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tahersoft-go/kit/response"
)

func CustomHandler(err error, c echo.Context) {
	r := buildErrorResponse(err)

	if err := c.JSON(r.StatusCode, err); err != nil {
		c.Logger().Errorf("failed writing error response: %v", err)
	}
}

// buildErrorResponse builds an error response from an error.
func buildErrorResponse(err error) response.ErrorResponse {
	switch err := err.(type) {
	case *echo.HTTPError:
		switch err.Code {
		case http.StatusNotFound:
			return response.ErrorNotFound(nil, "درخواست مورد نظر یافت نشد")
		default:
			return response.ErrorResponse{
				StatusCode: err.Code,
				Data:       nil,
				Message:    "خطایی در سرور رخ داد",
			}
		}
	}

	return response.ErrorInternalServerError(nil, "خطایی در سرور رخ داد")
}
