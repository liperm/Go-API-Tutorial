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
	CustomerID int               `json:"customer_id"`
	Items      []buyingListItems `json:"items"`
	TotalValue float32           `json:"total_value"`
}

type buyingListItems struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Value float32 `json:"value"`
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
	items := customer.Items
	var itemsList []buyingListItems
	for _, item := range items {
		itemsList = append(itemsList, buyingListItems{
			ID:    item.ID,
			Name:  item.Name,
			Code:  item.Code,
			Value: item.Value,
		})
	}

	return buyingListResponse{
		CustomerID: customer.ID,
		Items:      itemsList,
		TotalValue: totalValue,
	}
}
