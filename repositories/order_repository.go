package repositories

import (
	"github.com/Frientia/my-firebase-backend/config"
	"github.com/Frientia/my-firebase-backend/models"
)

func CreateOrder(
	order *models.Order,
) error {

	return config.DB.Create(order).Error
}

func CreateOrderItem(
	item *models.OrderItem,
) error {

	return config.DB.Create(item).Error
}

func GetOrdersByUserID(
	userID uint,
) ([]models.Order, error) {

	var orders []models.Order

	err := config.DB.
		Preload("Items").
		Where("user_id = ?", userID).
		Order("id DESC").
		Find(&orders).Error

	return orders, err
}

func GetOrderByID(
	id uint,
) (*models.Order, error) {

	var order models.Order

	err := config.DB.
		Preload("Items").
		Preload("Items.Product").
		First(&order, id).Error

	return &order, err
}

func UpdateOrderStatus(orderID int, status string) error {
	err := config.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
	return err
}