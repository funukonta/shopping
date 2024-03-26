package services

import (
	"strconv"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type PaymentService interface {
	DoPayment(string, *models.PaymentModel) (*models.PaymentModel, error)
}

type paymentService struct {
	repo repositories.PaymentRepo
}

func NewPaymentService(repo repositories.PaymentRepo) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) DoPayment(id_invoice string, payment *models.PaymentModel) (*models.PaymentModel, error) {
	invId, err := strconv.Atoi(id_invoice)
	if err != nil {
		return nil, err
	}

	paymentRes, err := s.repo.DoPayment(invId, payment)
	return paymentRes, err
}
