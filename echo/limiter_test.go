package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/didip/tollbooth"
	"github.com/labstack/echo"
)

func ExampleLimitMiddleware() {
	router := echo.New()
	router.Use(LimitMiddleware(tollbooth.NewLimiter(1, time.Minute)))
	router.GET("/", func(c echo.Context) (err error) {
		c.String(http.StatusOK, "Hello, Gin Limiter!")
		return
	})

	var req *http.Request
	var resp *httptest.ResponseRecorder
	for i := 0; i < 2; i++ {
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		resp = httptest.NewRecorder()
		router.ServeHTTP(resp, req)
	}

	fmt.Println(resp.Result().StatusCode)
	// OUTPUT: 429
}
