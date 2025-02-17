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
	user, err := h.service.Register(input)
	if err != nil {
		// Return specific error message from service
		network.SendErrorResponse(w, "USER_CREATION_ERROR", err.Error(), http.StatusConflict)
		return
	}

	// Successful response
	network.SendJSONResponse(w, network.Response{
		Code: "SUCCESS",
		Msg:  "User created successfully",
		Data: user,
	})
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var input UserLoginSchema

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

	// Call AuthService to login user
	token, err := h.service.Login(input)
	if err != nil {
		// Return specific error message from service
		network.SendErrorResponse(w, "USER_LOGIN_ERROR", err.Error(), http.StatusConflict)
		return
	}

	// set token in header cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,                    // Prevent JavaScript access (XSS protection)
		Secure:   true,                    // Set to `true` in production (HTTPS only)
		SameSite: http.SameSiteStrictMode, // Helps mitigate CSRF
		MaxAge:   3600,                    // 1 hour expiry
	})

	// Successful login, return the token
	network.SendJSONResponse(w, network.Response{
		Code: "SUCCESS",
		Msg:  "Login successful",
		Data: map[string]string{
			"token": token,
		},
	})

}
