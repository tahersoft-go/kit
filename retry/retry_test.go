package retry

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("When error is nil", func(t *testing.T) {
		result, err := Get(func() (string, error) {
			return "test", nil
		}, 1, 1)
		if err != nil {
			t.Error("Error is not nil")
		}
		if result != "test" {
			t.Error("Result is not test")
		}
	})

	t.Run("When error is not nil and retry is successful", func(t *testing.T) {
		result, err := Get(func() (string, error) {
			return "", errors.New("some error")
		}, 1, 2)
		if err == nil {
			t.Error("Error is not nil")
		}
		if result != "" {
			t.Error("Result is not empty")
		}
	})
}
