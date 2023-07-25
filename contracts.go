package httpsms

type ApiResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
