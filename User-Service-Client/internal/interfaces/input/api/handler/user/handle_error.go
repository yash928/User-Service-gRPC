package user

import (
	"net/http"
	"user-service-client/internal/core/user"
	"user-service-client/pkg/response"
)

func (u *UserHandlerImpl) handleError(w http.ResponseWriter, r *http.Request, err error) {
	var apiResponse *response.APIResponse
	switch err {
	case user.ErrUserNotFound:
		apiResponse = response.NewAPIResponse("failed", http.StatusBadRequest)
		apiResponse.Error = "could not fetch the user details"
	case user.ErrInvalidID:
		apiResponse = response.NewAPIResponse("failed", http.StatusBadRequest)
		apiResponse.Error = "invalid id"
	case user.ErrInvalidMaritalStatus:
		apiResponse = response.NewAPIResponse("failed", http.StatusBadRequest)
		apiResponse.Error = "invalid marital status"
	default:
		apiResponse = response.NewAPIResponse("failed", http.StatusInternalServerError)
		apiResponse.Error = "something went wrong"
	}
	response.ResponseJSON(w, r, apiResponse)
}
