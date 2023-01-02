package stdapi

import "github.com/gofiber/fiber/v2"

type Response struct {
	code int
	info string
	data interface{}
}

func NewResponse(code int, info string, data interface{}) *Response {
	return &Response{
		code: code,
		info: info,
		data: data,
	}
}

func (r *Response) ToFiberMap() fiber.Map {
	res := fiber.Map{
		"code": r.code,
		"info": r.info,
	}

	if r.data != nil {
		res["data"] = r.data
	}

	return res
}
