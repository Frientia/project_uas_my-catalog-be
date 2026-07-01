package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Frientia/my-firebase-backend/services"
)

type CheckoutRequest struct {
	ShippingAddress string `json:"shipping_address"`
	Notes           string `json:"notes"`
	PaymentMethod   string `json:"payment_method"`
}

func Checkout(c *gin.Context) {
	var req CheckoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := uint(1)

	order, err := services.Checkout(
		userID,
		req.ShippingAddress,
		req.Notes,
		req.PaymentMethod,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Checkout success",
		"data":    order,
	})
}

func GetMyOrders(c *gin.Context) {
	userID := uint(1)

	orders, err := services.GetMyOrders(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func GetOrderDetail(c *gin.Context) {
	idParam := c.Param("id")

	var id uint

	fmt.Sscanf(idParam, "%d", &id)

	order, err := services.GetOrderDetail(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func PaymentCallback(c *gin.Context) {
	idParam := c.Param("id")

	var id uint
	fmt.Sscanf(idParam, "%d", &id)

	err := services.MarkOrderAsPaid(int(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Status pesanan berhasil diubah menjadi paid",
	})
}