package auth

import (
	network "go-auth-v1/pkg/http"
	"go-auth-v1/pkg/validator"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewAuthHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var input UserStoreSchema

	// Parse request body
	err := network.ParseRequest(r, &input)
	if err != nil {
		network.SendErrorResponse(w, "BAD_REQUEST", "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	validationErrors := validator.ValidateStruct(input)
	if validationErrors != nil {
		network.SendJSONResponse(w, network.Response{
			Code: "VALIDATION_ERROR",
			Msg:  validationErrors.Msg,
			Data: validationErrors.Details,
		})
		return
	}

	// Call AuthService to create user
	user, err := h.service.CreateUser(input)
	if err != nil {
		network.SendErrorResponse(w, "INTERNAL_ERROR", "Could not create user", http.StatusInternalServerError)
		return
	}

	network.SendJSONResponse(w, network.Response{
		Code: "SUCCESS",
		Msg:  "User created successfully",
		Data: user,
	})
}
