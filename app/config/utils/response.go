package utils

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	Success = "success"
	Failed  = "failed"
)

func NewResponse(status, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
