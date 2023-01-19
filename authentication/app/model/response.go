package model

type (
	Response struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Results    interface{} `json:"results"`
	}

	PaginatedResponse struct {
		Count    int         `json:"count"`
		Next     string      `json:"next"`
		Previous string      `json:"previous"`
		Results  interface{} `json:"results"`
	}

	ValidatedResponse struct {
		StatusCode int                 `json:"statusCode"`
		Message    string              `json:"message"`
		Errors     map[string][]string `json:"errors"`
	}
)

func NewResponse(statusCode int, message string, results interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Results:    results,
	}
}

func NewPaginatedResponse(statusCode, count int, message, next, prev string, results interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Results: &PaginatedResponse{
			Count:    count,
			Next:     next,
			Previous: prev,
			Results:  results,
		},
	}
}

func NewValidatedResponse(statusCode int, message string, errors map[string][]string) *ValidatedResponse {
	return &ValidatedResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}
}
