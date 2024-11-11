package handlers

import (
	"app-example/internal/dto"
	"app-example/internal/entity"
	"app-example/internal/infra/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	userDb *database.User
}

func NewUserHandler(userDb *database.User) *UserHandler {
	return &UserHandler{
		userDb: userDb,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.CreateUserInput true "User data"
// @Success 201 {object} entity.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.userDb.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+user.ID.String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetJwtToken godoc
// @Summary Get JWT token
// @Description Get JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.GetJwtTokenInput true "User data"
// @Success 200 {object} dto.GetJwtTokenOutput
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /users/token [post]
func (h *UserHandler) GetJwtToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var jwtDto dto.GetJwtTokenInput
	err := json.NewDecoder(r.Body).Decode(&jwtDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userDb.FindByEmail(jwtDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = user.CheckPassword(jwtDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.GetJwtTokenOutput{Token: tokenString})
}
