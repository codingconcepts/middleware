package gin

import (
	"fmt"
	"net/http"
	"time"

	"net/http/httptest"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func ExampleLimitMiddleware() {
	router := gin.New()
	router.Use(LimitMiddleware(tollbooth.NewLimiter(1, time.Minute)))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin Limiter!")
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
