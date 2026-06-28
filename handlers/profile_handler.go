package handlers

import (
	"net/http"

	"github.com/Frientia/my-firebase-backend/repositories"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	userRepo *repositories.UserRepository
}

// Constructor untuk ProfileHandler
func NewProfileHandler(repo *repositories.UserRepository) *ProfileHandler {
	return &ProfileHandler{userRepo: repo}
}

// GetProfile menangani request GET /v1/profile
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	// 1. Ambil "firebase_uid" yang sudah di-set oleh AuthMiddleware
	uid, exists := c.Get("firebase_uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Sesi tidak valid atau tidak ditemukan",
		})
		return
	}

	// 2. Cari data user di database (casting tipe datanya ke string)
	user, err := h.userRepo.FindByFirebaseUID(uid.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data profil tidak ditemukan",
		})
		return
	}

	// 3. Kembalikan data user ke Flutter
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}