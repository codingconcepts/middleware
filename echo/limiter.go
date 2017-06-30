package echo

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/config"
	"github.com/labstack/echo"
)

// LimitMiddleware wraps the tollbooth Limiter's functionality
// and aborts or calls next for subsequent handlers accordingly
func LimitMiddleware(limiter *config.Limiter) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tollbooth.SetResponseHeaders(limiter, c.Response().Writer)

			if err := tollbooth.LimitByRequest(limiter, c.Request()); err != nil {
				return echo.NewHTTPError(http.StatusTooManyRequests, err)
			}
			
			return next(c)
		}
	}
}
