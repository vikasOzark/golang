package handlers

import (
	model "echo/database/models"
	"fmt"
	"net/http"
	"strings"

	database "echo/database/connection"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"

	helpers "echo/api"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// @Summary Get a greeting
// @Description Returns a greeting message
// @Tags Example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/login [post]
func AuthHandler(c echo.Context) error {
	var payload LoginPayload
	// payload := c.Bind()
	// if email == "" || password == "" {
	// 	response := helpers.NewResponse[any, any](false, "Email and password are required", http.StatusBadRequest, nil, nil)
	// 	return helpers.ResponseProvider(response, c)
	// }
	if err := c.Bind(&payload); err != nil {
		response := helpers.NewResponse[any, error](false, "Invalid request payload", http.StatusBadRequest, nil, err)
		return helpers.ResponseProvider(response, c)
	}

	fmt.Println("== > ", payload)
	val := database.Connect().Find(payload.Email)

	// Validate user credentials (this is just a placeholder, implement your own logic)
	if payload.Password != "" || *payload.Email != "" {
		response := helpers.NewResponse[any, any](false, "Unauthorized", http.StatusUnauthorized, nil, nil)
		return helpers.ResponseProvider(response, c)
	}

	claims := &jwtCustomClaims{
		Name:  "Jon Snow",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		response := helpers.NewResponse[any, error](false, "Failed to generate token", http.StatusInternalServerError, nil, err)
		return helpers.ResponseProvider(response, c)
	}

	response := helpers.NewResponse[echo.Map, any](true, "Token generated successfully", http.StatusOK, echo.Map{"token": t}, nil)
	return helpers.ResponseProvider(response, c)
}

// @Summary Get a greeting
// @Description Returns a greeting message
// @Tags Example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/registration [post]
func Registration(c echo.Context) error {
	var payload model.User

	if err := c.Bind(&payload); err != nil {
		response := helpers.NewResponse[any, error](false, "Invalid request payload", http.StatusBadRequest, nil, err)
		return helpers.ResponseProvider(response, c)
	}

	if err := payload.Validate(); err != nil {
		response := helpers.NewResponse[any, string](false, "Validation failed", http.StatusUnprocessableEntity, nil, err.Error())
		return helpers.ResponseProvider(response, c)
	}

	payload.EncyPassword()
	result := database.Connect().Create(&payload)

	if result.Error != nil {
		response := helpers.NewResponse[any, string](false, "Failed to create user", http.StatusBadRequest, nil, strings.Split(result.Error.Error(), ":")[1])
		return helpers.ResponseProvider(response, c)
	}

	response := helpers.NewResponse[*gorm.DB, any](true, "User created successfully", http.StatusCreated, result, nil)
	return helpers.ResponseProvider(response, c)
}
