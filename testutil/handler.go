package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ApiV1Url(path string) string {
	return "/api/v1" + path
}

func GoldenPath(t *testing.T) string {
	t.Helper()

	return "./testdata/" + t.Name() + ".golden"
}

func AssertContentType(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	expected := "application/json; charset=utf-8"
	actual := res.Header().Get("Content-Type")
	assert.Equal(t, expected, actual)
}

func AssertResponseHeader200(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()

	assert.Equal(t, http.StatusOK, res.Code)
	AssertContentType(t, res)
}
