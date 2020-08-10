package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/JohnnyS318/RoyalAfgInGo/account/pkg/account/models"
)

// Register registers a new user
// swagger:route POST /account/register authentication
//
//	Register a new user
//
// This will register a new user with the provided details
//
//	Consumes:
//	-	application/json
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
// 	Responses:
// 		default: ErrorResponse
//		400: ErrorResponse
//		422: ValidationErrorResponse
//		500: ErrorResponse
//		200: UserResponse
//
func (h *User) Register(rw http.ResponseWriter, r *http.Request) {
	h.l.Info("Register route called")

	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	dto := &RegisterUser{}
	err := FromJSON(dto, r.Body)
	if err != nil {
		h.l.Error("Decoding JSON", "error", err)
		JSONError(rw, &ErrorResponse{Error: "user could not be decoded"}, http.StatusBadRequest)
		return
	}

	h.l.Debug("Decoded user")

	if err := dto.Validate(); err != nil {
		h.l.Error("Validation", "error", err)
		JSONError(rw, &ValidationError{Errors: err}, http.StatusUnprocessableEntity)
		return
	}

	user, err := dto.ToObject()

	if err != nil {
		h.l.Error(err)
		JSONError(rw, &ErrorResponse{Error: "User could not be created"}, http.StatusInternalServerError)
		return
	}

	h.l.Debug("User validated")

	if err = h.db.CreateUser(user); err != nil {
		h.l.Error(err)
		JSONError(rw, &ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	h.l.Debug("User saved")

	token, err := getJwt(user)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		h.l.Error(err)
		return
	}

	cookie := generateCookie(token, dto.RememberMe)

	http.SetCookie(rw, cookie)
	ToJSON(NewUserDTO(user), rw)
}

// RegisterUser defines the dto for the user account registration
type RegisterUser struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FullName   string `json:"fullName"`
	RememberMe bool   `json:"rememberme"`
}

// Validate validates if the RegisterUser dto matches all the user requirements
func (dto RegisterUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.FullName, validation.Required, validation.Length(1, 100)),
		validation.Field(&dto.Email, validation.Required, is.Email),
	)
}

// ToObject converts the RegisterUser dto to the internal user object
func (dto RegisterUser) ToObject() (*models.User, error) {
	user := &models.User{
		Username: dto.Username,
		Email:    dto.Email,
		FullName: dto.FullName,
	}
	err := user.SetPassword(dto.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
