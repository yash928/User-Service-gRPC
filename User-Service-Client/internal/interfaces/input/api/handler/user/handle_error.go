package user

import (
	"net/http"
	"user-service-client/internal/core/user"
	"user-service-client/pkg/response"
)

func (u *UserHandlerImpl) handleError(w http.ResponseWriter, r *http.Request, err error) {
	var apiResponse *response.APIResponse
	switch err {
	case user.ErrInvalidID:
		apiResponse = response.NewAPIResponse("failed", http.StatusBadRequest)
		apiResponse.Error = "invalid id"
	default:
		apiResponse = response.NewAPIResponse("failed", http.StatusInternalServerError)
		apiResponse.Error = "something went wrong"
	}
	response.ResponseJSON(w, r, apiResponse)
}
