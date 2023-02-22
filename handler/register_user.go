package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sugimoto-ne/go_sample_app.git/entity"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required,min=8,custom-low,custom-upp,custom-num,custom-symbol"`
		Role     string `json:"role" validate:"required"`
	}

	// jsonのデコード
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	err := ru.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
