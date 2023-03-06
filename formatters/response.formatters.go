package formatters

import (
	"fmt"
	"reflect"

	"github.com/liperm/go-api/models"
)

type errorResponse struct {
	ErrorMessage string `json:"error"`
}

type successResponse struct {
	ID int `json:"id"`
}

type notFoundResponse struct {
	Message string `json:"message"`
}

type buyingListResponse struct {
	CustomerID int           `json:"customer_id"`
	Items      []models.Item `json:"items"`
	TotalValue float32       `json:"total_value"`
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

func FormatNotFoundResponse[T string | int](value T) notFoundResponse {
	if reflect.TypeOf(value).Kind() == reflect.Int {
		return notFoundResponse{
			Message: fmt.Sprint("id: ", value, " not found"),
		}
	}

	return notFoundResponse{
		Message: fmt.Sprint("code: ", value, " not found"),
	}
}

func FormatBuyingListResponse(customer models.Customer, totalValue float32) buyingListResponse {
	return buyingListResponse{
		CustomerID: customer.ID,
		Items:      customer.Items,
		TotalValue: totalValue,
	}
}
