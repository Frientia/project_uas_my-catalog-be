package services

import (
	"errors"

	"github.com/Frientia/my-firebase-backend/models"
	"github.com/Frientia/my-firebase-backend/repositories"
)

func Checkout(
	userID uint,
	shippingAddress string,
	notes string,
	paymentMethod string,
) (*models.Order, error) {

	cartItems, err :=
		repositories.GetCartByUserID(userID)

	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	var total float64 = 0

	for _, item := range cartItems {
		subtotal :=
			item.Product.Price *
				float64(item.Quantity)

		total += subtotal
	}

	order := &models.Order{
		UserID: userID,

		TotalAmount: total,

		Status: "pending",

		ShippingAddress: shippingAddress,

		Notes: notes,

		PaymentMethod: paymentMethod,
	}

	err = repositories.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {

		subtotal :=
			cartItem.Product.Price *
				float64(cartItem.Quantity)

		orderItem := &models.OrderItem{
			OrderID: order.ID,

			ProductID: cartItem.ProductID,

			Quantity: cartItem.Quantity,

			Price: cartItem.Product.Price,

			Subtotal: subtotal,
		}

		err := repositories.CreateOrderItem(orderItem)

		if err != nil {
			return nil, err
		}
	}

	err = repositories.ClearCartByUserID(userID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func GetMyOrders(
	userID uint,
) ([]models.Order, error) {

	return repositories.GetOrdersByUserID(userID)
}

func GetOrderDetail(
	id uint,
) (*models.Order, error) {

	return repositories.GetOrderByID(id)
}

func MarkOrderAsPaid(orderID int) error {
	return repositories.UpdateOrderStatus(orderID, "paid")
}