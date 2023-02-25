package formatters

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

type creationSuccessResponse struct {
	ID int `json:"id"`
}

func FormatErrorResponse(err string) errorResponse {
	return errorResponse{
		ErrorMessage: err,
	}
}

func FormatCreationSuccessResponse(id int) creationSuccessResponse {
	return creationSuccessResponse{
		ID: id,
	}
}
