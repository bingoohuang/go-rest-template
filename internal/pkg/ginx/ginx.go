package ginx

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Status int
	Data   interface{}
}

func StatusJSON(status int, v interface{}) JSONResponse {
	return JSONResponse{
		Status: status,
		Data:   v,
	}
}

func JSON(v interface{}) JSONResponse {
	return StatusJSON(http.StatusOK, v)
}

func (j JSONResponse) Render(c *gin.Context) {
	c.JSON(j.Status, j.Data)
}

type Render interface {
	Render(c *gin.Context)
}

func ParamBindJSON(f func(*gin.Context, string, interface{}) Render, name string, bind interface{}) func(c *gin.Context) {
	typ := reflect.TypeOf(bind)

	return func(c *gin.Context) {
		b := reflect.New(typ)
		if err := c.BindJSON(b.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})

			return
		}

		p := c.Param(name)
		ret := f(c, p, b.Elem().Interface())
		ret.Render(c)
	}
}

func BindJSON(f func(*gin.Context, interface{}) Render, bind interface{}) func(c *gin.Context) {
	typ := reflect.TypeOf(bind)

	return func(c *gin.Context) {
		b := reflect.New(typ)
		if err := c.BindJSON(b.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})

			return
		}

		ret := f(c, b.Elem().Interface())
		ret.Render(c)
	}
}

func Param(f func(*gin.Context, string) Render, name string) func(c *gin.Context) {
	return func(c *gin.Context) {
		ret := f(c, c.Param(name))
		ret.Render(c)
	}
}

func Bind(f func(*gin.Context, interface{}) Render, bind interface{}) func(c *gin.Context) {
	typ := reflect.TypeOf(bind)

	return func(c *gin.Context) {
		b := reflect.New(typ)
		if err := c.Bind(b.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})

			return
		}

		ret := f(c, b.Elem().Interface())
		ret.Render(c)
	}
}

func Wrap(f func(*gin.Context) Render) func(c *gin.Context) {
	return func(c *gin.Context) {
		ret := f(c)
		ret.Render(c)
	}
}

type StatusError struct {
	Msg    string
	Err    error
	Status int
}

func (n StatusError) Render(c *gin.Context) {
	c.JSON(n.Status, HTTPError{
		Code:    n.Status,
		Message: n.Msg,
	})

	log.Println(n.Err)
}

func NewBadRequestError(msg string, err error) StatusError {
	return StatusError{Msg: msg, Err: err, Status: http.StatusBadRequest}
}

func NewForbiddenError(msg string, err error) StatusError {
	return StatusError{Msg: msg, Err: err, Status: http.StatusForbidden}
}

func NewNotFoundError(msg string, err error) StatusError {
	return StatusError{Msg: msg, Err: err, Status: http.StatusNotFound}
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
