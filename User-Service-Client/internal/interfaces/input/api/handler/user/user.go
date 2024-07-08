package user

import (
	"log"
	"net/http"
	"user-service-client/internal/core/user"
	"user-service-client/pkg/response"

	"github.com/go-chi/chi"
)

type UserHandlerImpl struct {
	uc user.UserUsecase
}

func NewUserHand(userUC user.UserUsecase) *UserHandlerImpl {
	return &UserHandlerImpl{
		uc: userUC,
	}
}

func (u *UserHandlerImpl) FindUserById() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("GET:/{user_id}")
		userID := chi.URLParam(r, "user_id")

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
