package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"km5go/configs"
	"km5go/middleware"
	"km5go/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func RegisterController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

	userRequest.Password, _ = generateSHA256Hash(userRequest.Password)

	result := configs.DB.Create(&userRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed to insert user",
			Data: nil,
		})
	}

	var userResponse = models.UserResponse {
		Id: userRequest.Id,
		PhotoUrl: userRequest.PhotoUrl,
		UserName: userRequest.UserName,
		FullName: userRequest.FullName,
	}

	// generate token jwt
	userResponse.Token, _ = middleware.GenerateJwt(
		userResponse.FullName,
		userResponse.Id,
	)

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status: true,
		Message: "Success add user",
		Data: userResponse,
	})
}

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func generateSHA256Hash(input string) (string, error) {
    // Membuat sebuah objek hasher SHA-256
    hasher := sha256.New()

    // Menambahkan data string ke hasher
    _, err := hasher.Write([]byte(input))
    if err != nil {
        return "", err
    }

    // Mendapatkan hasil hash sebagai byte
    hashBytes := hasher.Sum(nil)

    // Mengonversi hasil hash menjadi string hexadecimal
    hashString := hex.EncodeToString(hashBytes)

    return hashString, nil
}

func GetUsersController(c echo.Context) error {
	var users []models.User

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil get data",
		Data: users,
	})
}