package handler

import (
	"gin_gorm_tutorial/testutil"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func SendRequest(t *testing.T, recorder *httptest.ResponseRecorder, httpMethod string, path string, body io.Reader) {
	t.Helper()

	gin.SetMode(gin.ReleaseMode)
	router := SetupRouter()

	req, _ := http.NewRequest(httpMethod, path, body)
	router.ServeHTTP(recorder, req)
}

func TestUserIndex(t *testing.T) {
	t.Run("collect_return", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		SendRequest(t, recorder, http.MethodGet, testutil.ApiV1Url("/users"), nil)

		testutil.AssertResponseHeader200(t, recorder)
		bytes, _ := os.ReadFile(testutil.GoldenPath(t))
		expected := string(bytes)
		actual := recorder.Body.String()

		assert.JSONEq(t, expected, actual)
	})
}
