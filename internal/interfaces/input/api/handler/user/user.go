package user

import (
	"net/http"
	"user-service-grpc/internal/core/user"
	"user-service-grpc/pkg/response"
)

type UserHandlerImpl struct {
	uc user.UserUsecase
}

func (u *UserHandlerImpl) FindUserById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.URL.Query().Get("user_id")

		userDet, err := u.uc.FindUserById(r.Context(), userID)
		if err != nil {
			u.handleError(w, r, err)
			return
		}

		apiResponse := response.NewAPIResponse("success", http.StatusOK)
		apiResponse.Message = "user details fetched successfully"
		apiResponse.Data = userDet
		response.ResponseJSON(w, r, apiResponse)

	}
}
