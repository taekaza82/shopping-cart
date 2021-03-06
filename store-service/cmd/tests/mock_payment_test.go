package tests

import (
	"store-service/internal/payment"

	"github.com/stretchr/testify/mock"
)

type mockPaymentService struct {
	mock.Mock
}

func (paymentService *mockPaymentService) ConfirmPayment(orderID int, paymentDetail payment.PaymentDetail) string {
	argument := paymentService.Called(orderID, paymentDetail)
	return argument.String(0)
}
