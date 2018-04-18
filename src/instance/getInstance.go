package instance

import "github.com/labstack/echo"

var e *echo.Echo

func GetInstance() *echo.Echo {
	if e != nil {
		return e
	}

	e = echo.New()
	return e
}
