package endpoint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/o1vanova/calculator/cmd/calculator/internal/service"
)

type EndPoint struct {
	svc service.SimpleOperationService
}

func New(svc service.SimpleOperationService) *EndPoint {
	return &EndPoint{
		svc: svc,
	}
}

func (ep *EndPoint) Info(c echo.Context) error {
	err := c.HTML(http.StatusOK, `<b>Information</b>&nbsp;<ul>
	<li>Plus example <a href="/plus?first=2&second=3">2 + 3</a></li>
	<li>Minus example <a href="/minus?first=2&second=3">2 - 3</a></li>
	<li>Multify example <a href="/multify?first=2&second=3">2 * 3</a></li></ul>`)
	if err != nil {
		return err
	}
	return nil
}

func (ep *EndPoint) Plus(c echo.Context) error {
	first, second, err := ep.getParams(c)
	if err != nil {
		return nil
	}

	data := ep.svc.Plus(first, second)

	err = c.String(http.StatusOK, fmt.Sprintf("Plus result: %d", data))
	if err != nil {
		return err
	}
	return nil
}

func (ep *EndPoint) Minus(c echo.Context) error {
	first, second, err := ep.getParams(c)
	if err != nil {
		return nil
	}

	data := ep.svc.Minus(first, second)

	err = c.String(http.StatusOK, fmt.Sprintf("Minus result: %d", data))
	if err != nil {
		return err
	}
	return nil
}

func (ep *EndPoint) Multify(c echo.Context) error {
	first, second, err := ep.getParams(c)
	if err != nil {
		return nil
	}

	data := ep.svc.Multify(first, second)

	err = c.String(http.StatusOK, fmt.Sprintf("Multify result: %d", data))
	if err != nil {
		return err
	}
	return nil
}

func (ep *EndPoint) getParams(c echo.Context) (int64, int64, error) {
	param := c.QueryParam("first")
	first, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	param = c.QueryParam("second")
	second, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return first, second, nil
}
