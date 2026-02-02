package dto

// ErrorBody - Định nghĩa cấu trúc body của error
type ErrorBody struct {
	Body []string `json:"body"`
}

// ErrorResponse - Định nghĩa cấu trúc error response theo RealWorld API spec
type ErrorResponse struct {
	Errors ErrorBody `json:"errors"`
}
