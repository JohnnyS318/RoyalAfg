package handlers

import (
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
)

// Session verifies the session and extends the jwt token if valid.
// swagger:route GET /api/auth/session session
//
// Session verifies the session and extends the jwt token if valid.
//
// After verification the extended jwt will be passed as a cookie and the user id and username will be returned
//	Consumes:
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
//	Responses:
//	default: ErrorResponse
//	401: ErrorResponse
//	200: UserResponse
func (h *Auth) Session(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(viper.GetString(config.SessionCookieName))
	if err != nil || cookie.Value == "" {
		h.l.Errorw("Cookie Extraction", "error", err, "cookieName")
		responses.Unauthorized(rw)
		return
	}

	parsed, token, err := services.ExtendToken(cookie.Value)

	if err != nil {
		switch e := err.(type) {
		case *errors.InvalidTokenError:
			cookie = services.GenerateCookie("", false)
			cookie.Expires = time.Unix(0, 0)
			http.SetCookie(rw, cookie)
			h.l.Errorw("Error token invalid", "error", err, "wrapped", e.Err)
			responses.Error(rw, "token could not be parsed. This removes the session", http.StatusUnauthorized)
			http.Error(rw, "token could not be parsed", http.StatusUnauthorized)
			return
		default:
			h.l.Errorw("Error token parsing", "error", err)
			responses.Unauthorized(rw)
			return
		}
	}

	user := mw.FromUserTokenContext(parsed)

	cookie = services.GenerateCookie(token, false)

	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
	err = ToJSON(dtos.SessionResponse{User: &dtos.SessionUser{Username: user.Username, Name: user.Name, Id: user.ID}}, rw)

	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
