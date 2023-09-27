package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ServeHTTP(t *testing.T) {
	t.Run("When response is from SuccessResponse", func(t *testing.T) {
		data := "Hello"
		h := Handler{
			Fn: func(w http.ResponseWriter, r *http.Request) Response {
				return Success(data)
			},
		}
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		body := w.Result().Body
		defer body.Close()
		bytes, err := io.ReadAll(body)
		if err != nil {
			t.Error(err)
		}
		bodyData := map[string]interface{}{}
		err = json.Unmarshal(bytes, &bodyData)
		if err != nil {
			t.Error(err)
		}
		dataField, ok := bodyData["data"]
		if !ok {
			t.Error("Response should have data field")
		}
		statusCodeField, ok := bodyData["statusCode"]
		if !ok {
			t.Error("Response should have statusCode field")
		}
		if dataField != data {
			t.Error("Response data is invalid")
		}
		if statusCodeField.(float64) != http.StatusOK {
			t.Error("Status code should be equal to 200")
		}
	})

	t.Run("When response is from ErrorResponse", func(t *testing.T) {
		data := "Hello"
		h := Handler{
			Fn: func(w http.ResponseWriter, r *http.Request) Response {
				return Error(data, http.StatusBadRequest)
			},
		}
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		body := w.Result().Body
		defer body.Close()
		bytes, err := io.ReadAll(body)
		if err != nil {
			t.Error(err)
		}
		bodyData := map[string]interface{}{}
		err = json.Unmarshal(bytes, &bodyData)
		if err != nil {
			t.Error(err)
		}
		dataField, ok := bodyData["data"]
		if !ok {
			t.Error("Response should have data field")
		}
		statusCodeField, ok := bodyData["statusCode"]
		if !ok {
			t.Error("Response should have statusCode field")
		}
		if dataField != data {
			t.Error("Response data is invalid")
		}
		if statusCodeField.(float64) != http.StatusBadRequest {
			t.Error("Status code should be equal to 400")
		}
	})
}

func TestHandle(t *testing.T) {
	h := Handle(func(w http.ResponseWriter, r *http.Request) Response {
		return Success("Hello")
	})
	dataType := fmt.Sprintf("%T", h)
	if dataType != "response.Handler" {
		t.Error("Result is invalid")
	}
}
