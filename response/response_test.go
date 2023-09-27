package response

import (
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

func TestCreated(t *testing.T) {
	data := "Hello"
	result := Created(data)
	if result.StatusCode != http.StatusCreated {
		t.Error("Status code is invalid")
	}
	switch result.Data.(type) {
	case string:
	default:
		t.Error("Data type is invalid")
	}
	if result.Data.(string) != data {
		t.Error("Data is invalid")
	}
}

func TestAccepted(t *testing.T) {
	data := "Hello"
	result := Accepted(data)
	if result.StatusCode != http.StatusAccepted {
		t.Error("Status code is invalid")
	}
	switch result.Data.(type) {
	case string:
	default:
		t.Error("Data type is invalid")
	}
	if result.Data.(string) != data {
		t.Error("Data is invalid")
	}
}

func TestNoContent(t *testing.T) {
	data := ""
	result := NoContent()
	if result.StatusCode != http.StatusNoContent {
		t.Error("Status code is invalid")
	}
	switch result.Data.(type) {
	case string:
	default:
		t.Error("Data type is invalid")
	}
	if result.Data.(string) != data {
		t.Error("Data is invalid")
	}
}

func TestAlreadyReported(t *testing.T) {
	data := "Hello"
	result := AlreadyReported(data)
	if result.StatusCode != http.StatusAlreadyReported {
		t.Error("Status code is invalid")
	}
	switch result.Data.(type) {
	case string:
	default:
		t.Error("Data type is invalid")
	}
	if result.Data.(string) != data {
		t.Error("Data is invalid")
	}
}

func TestError(t *testing.T) {
	t.Run("When message is not ready", func(t *testing.T) {
		data := "Hello"
		result := Error(data, http.StatusBadRequest)
		if result.StatusCode != http.StatusBadRequest {
			t.Error("Status code is invalid")
		}
		switch result.Data.(type) {
		case string:
		default:
			t.Error("Data type is invalid")
		}
		if result.Data.(string) != data {
			t.Error("Data is invalid")
		}
	})

	t.Run("When message is prepare", func(t *testing.T) {
		data := "Hello"
		message := "value"
		result := Error(data, http.StatusBadRequest, message)
		if result.StatusCode != http.StatusBadRequest {
			t.Error("Status code is invalid")
		}
		switch result.Data.(type) {
		case string:
		default:
			t.Error("Data type is invalid")
		}
		if result.Data.(string) != data {
			t.Error("Data is invalid")
		}
		if result.Message != message {
			t.Error("Message is not valid")
		}
	})
}

type ErrorTestCase struct {
	Fn       func(data interface{}, messages ...string) ErrorResponse
	Data     interface{}
	Message  *string
	Expected ErrorResponse
}

var sampleErrorMessage = "message"

var errorTestCases = []ErrorTestCase{
	{Fn: ErrorBadRequest, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusBadRequest, "data", "Bad Request"}},
	{Fn: ErrorBadRequest, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusBadRequest, "data", sampleErrorMessage}},
	{Fn: ErrorUnAuthorized, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusUnauthorized, "data", "Not Logged In"}},
	{Fn: ErrorUnAuthorized, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusUnauthorized, "data", sampleErrorMessage}},
	{Fn: ErrorForbidden, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusForbidden, "data", "Access Denied"}},
	{Fn: ErrorForbidden, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusForbidden, "data", sampleErrorMessage}},
	{Fn: ErrorNotFound, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusNotFound, "data", "Not Found"}},
	{Fn: ErrorNotFound, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusNotFound, "data", sampleErrorMessage}},
	{Fn: ErrorMethodNotAllowed, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusMethodNotAllowed, "data", "Method Not Allowed"}},
	{Fn: ErrorMethodNotAllowed, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusMethodNotAllowed, "data", sampleErrorMessage}},
	{Fn: ErrorUnprocessableEntity, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusUnprocessableEntity, "data", "Request Temporary Blocked"}},
	{Fn: ErrorUnprocessableEntity, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusUnprocessableEntity, "data", sampleErrorMessage}},
	{Fn: ErrorLocked, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusLocked, "data", "Request Temporary Blocked"}},
	{Fn: ErrorLocked, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusLocked, "data", sampleErrorMessage}},
	{Fn: ErrorTooManyRequests, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusTooManyRequests, "data", "Too Many Requests"}},
	{Fn: ErrorTooManyRequests, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusTooManyRequests, "data", sampleErrorMessage}},
	{Fn: ErrorInternalServerError, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusInternalServerError, "data", "Something Went Wrong"}},
	{Fn: ErrorInternalServerError, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusInternalServerError, "data", sampleErrorMessage}},
	{Fn: ErrorRequestEntityTooLarge, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusRequestEntityTooLarge, "data", "Request Entity Too Large"}},
	{Fn: ErrorRequestEntityTooLarge, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusRequestEntityTooLarge, "data", sampleErrorMessage}},
	{Fn: ErrorUpgradeRequired, Data: "data", Message: nil, Expected: ErrorResponse{http.StatusUpgradeRequired, "data", "Upgrade Your Client"}},
	{Fn: ErrorUpgradeRequired, Data: "data", Message: &sampleErrorMessage, Expected: ErrorResponse{http.StatusUpgradeRequired, "data", sampleErrorMessage}},
}

func TestErrorBadRequest(t *testing.T) {
	for _, errorTestCase := range errorTestCases {
		functionName := runtime.FuncForPC(reflect.ValueOf(errorTestCase.Fn).Pointer()).Name()
		var result ErrorResponse
		if errorTestCase.Message == nil {
			result = errorTestCase.Fn(errorTestCase.Data)
		} else {
			result = errorTestCase.Fn(errorTestCase.Data, *errorTestCase.Message)
		}
		if result.StatusCode != errorTestCase.Expected.StatusCode {
			t.Errorf("%s: Status code mismatch", functionName)
		}
		if result.Data != errorTestCase.Expected.Data {
			t.Errorf("%s: Data mismatch", functionName)
		}
		if result.Message != errorTestCase.Expected.Message {
			t.Errorf("%s: Message mismatch", functionName)
		}
	}
}
