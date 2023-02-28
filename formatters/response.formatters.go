package formatters

import "fmt"

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

type successResponse struct {
	ID int `json:"id"`
}

type notFoundResponse struct {
	Message string `json:"message"`
}

func FormatErrorResponse(err string) errorResponse {
	return errorResponse{
		ErrorMessage: err,
	}
}

func FormatSuccessResponse(id int) successResponse {
	return successResponse{
		ID: id,
	}
}

func FormatNotFoundResponse(id int) notFoundResponse {
	return notFoundResponse{
		Message: fmt.Sprint("id: ", id, " not found"),
	}
}
