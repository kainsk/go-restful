package helpers

import "sqlc-rest-api/responses"

func SuccessResponse(message string, data any) responses.ApiResponse {
	return responses.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}
