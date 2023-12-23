package response

// HttpResponse . . .
type HttpResponse struct {
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}
