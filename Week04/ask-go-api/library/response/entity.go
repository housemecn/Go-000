package response

// @Project: yuerso-go-api
// @Author: houseme
// @Description:
// @File: entity
// @Version: 1.0.0
// @Date: 2020/11/15 11:10
// @Package response

// Response result
type Response struct {
	Err     error
	Data    interface{}
	Code    int
	Message string
}

// NewDefaultResponse new default Object
func NewDefaultResponse() *Response {
	return &Response{
		Err:     nil,
		Data:    nil,
		Code:    200,
		Message: "success",
	}
}
