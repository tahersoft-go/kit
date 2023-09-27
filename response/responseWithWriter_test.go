package response

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonWithWriter(t *testing.T) {
	t.Run("When w is ready", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]interface{}{
			"statusCode": http.StatusOK,
			"data":       "data",
		}
		statusCode := http.StatusOK
		err := JsonWithWriter(w, data, statusCode)
		if err != nil {
			t.Error("Error should be nil")
		}
		contentType := w.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Error("Content-Type should be application/json")
		}
		if w.Code != http.StatusOK {
			t.Error("Status code is invalid")
		}
		body := w.Result().Body
		defer body.Close()
		bytes, err := io.ReadAll(body)
		if err != nil {
			t.Error(err)
		}
		var bodyData map[string]interface{}
		err = json.Unmarshal(bytes, &bodyData)
		if err != nil {
			t.Error(err)
		}
		statusCodeField, ok := bodyData["statusCode"]
		if !ok {
			t.Error("Status code field not found in response")
		}
		dataField, ok := bodyData["data"]
		if !ok {
			t.Error("Data field not found in response data")
		}
		if statusCodeField.(float64) != float64(statusCode) {
			t.Error("Status code value is invalid")
		}
		if dataField.(string) != "data" {
			t.Error("Data is invalid")
		}
	})

	t.Run("When there is error in json encoding", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]interface{}{
			"statusCode": http.StatusOK,
			"data":       make(chan int),
		}
		statusCode := http.StatusOK
		err := JsonWithWriter(w, data, statusCode)
		if err == nil {
			t.Error("Error can not be nil")
		}
	})
}

func TestSuccessWithWriter(t *testing.T) {
	w := httptest.NewRecorder()
	data := SuccessResponse{http.StatusOK, "data"}
	err := SuccessWithWriter(w, data)
	if err != nil {
		t.Error("Error should be nil")
	}
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Error("Content-Type should be application/json")
	}
	if w.Code != http.StatusOK {
		t.Error("Status code is invalid")
	}
	body := w.Result().Body
	defer body.Close()
	bytes, err := io.ReadAll(body)
	if err != nil {
		t.Error(err)
	}
	var bodyData map[string]interface{}
	err = json.Unmarshal(bytes, &bodyData)
	if err != nil {
		t.Error(err)
	}
	statusCodeField, ok := bodyData["statusCode"]
	if !ok {
		t.Error("Status code field not found in response")
	}
	dataField, ok := bodyData["data"]
	if !ok {
		t.Error("Data field not found in response data")
	}
	if statusCodeField.(float64) != http.StatusOK {
		t.Error("Status code value is invalid")
	}
	if dataField.(string) != "data" {
		t.Error("Data is invalid")
	}
}

func TestErrorWithWriter(t *testing.T) {
	t.Run("When custom message is not present", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := ErrorResponse{http.StatusNotFound, "data", "Sample Error"}
		err := ErrorWithWriter(w, data, http.StatusNotFound)
		if err != nil {
			t.Error("Error should be nil")
		}
		contentType := w.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Error("Content-Type should be application/json")
		}
		if w.Code != http.StatusNotFound {
			t.Error("Status code is invalid")
		}
		body := w.Result().Body
		defer body.Close()
		bytes, err := io.ReadAll(body)
		if err != nil {
			t.Error(err)
		}
		var bodyData map[string]interface{}
		err = json.Unmarshal(bytes, &bodyData)
		if err != nil {
			t.Error(err)
		}
		statusCodeField, ok := bodyData["statusCode"]
		if !ok {
			t.Error("Status code field not found in response")
		}
		dataField, ok := bodyData["data"]
		if !ok {
			t.Error("Data field not found in response data")
		}
		messageField, ok := bodyData["message"]
		if !ok {
			t.Error("Message field not found in response data")
		}
		if statusCodeField.(float64) != http.StatusNotFound {
			t.Error("Status code value is invalid")
		}
		if dataField.(string) != "data" {
			t.Error("Data is invalid")
		}
		if messageField.(string) != "Sample Error" {
			t.Error("Message is invalid")
		}
	})

	t.Run("When custom message is prepare", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := ErrorResponse{http.StatusNotFound, "data", "Sample Error"}
		err := ErrorWithWriter(w, data, http.StatusNotFound, "Custom Message")
		if err != nil {
			t.Error("Error should be nil")
		}
		contentType := w.Header().Get("Content-Type")
		if contentType != "application/json" {
			t.Error("Content-Type should be application/json")
		}
		if w.Code != http.StatusNotFound {
			t.Error("Status code is invalid")
		}
		body := w.Result().Body
		defer body.Close()
		bytes, err := io.ReadAll(body)
		if err != nil {
			t.Error(err)
		}
		var bodyData map[string]interface{}
		err = json.Unmarshal(bytes, &bodyData)
		if err != nil {
			t.Error(err)
		}
		statusCodeField, ok := bodyData["statusCode"]
		if !ok {
			t.Error("Status code field not found in response")
		}
		dataField, ok := bodyData["data"]
		if !ok {
			t.Error("Data field not found in response data")
		}
		messageField, ok := bodyData["message"]
		if !ok {
			t.Error("Message field not found in response data")
		}
		if statusCodeField.(float64) != http.StatusNotFound {
			t.Error("Status code value is invalid")
		}
		if dataField.(string) != "data" {
			t.Error("Data is invalid")
		}
		if messageField.(string) != "Custom Message" {
			t.Error("Message is invalid")
		}
	})
}
