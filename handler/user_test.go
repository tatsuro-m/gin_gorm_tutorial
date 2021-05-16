package handler

import (
	"gin_gorm_tutorial/testutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserIndex(t *testing.T) {
	t.Run("collect_return", func(t *testing.T) {
		router := SetupRouter()
		res := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, testutil.ApiV1Url("users"), nil)
		router.ServeHTTP(res, req)

		testutil.AssertResponseHeader200(t, res)
		b, _ := os.ReadFile(testutil.GoldenPath(t))
		expected := string(b)
		actual := res.Body.String()

		assert.JSONEq(t, expected, actual)
	})
}
