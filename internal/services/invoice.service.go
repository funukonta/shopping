package services

import (
	"strconv"

	"github.com/funukonta/shopping/internal/models"
	"github.com/funukonta/shopping/internal/repositories"
)

type InvoiceService interface {
	Create(id_cust string, product []models.InvoiceDetail) (*models.ResJoinInvoice, error)
	GetInvoice(id_cust string) (*models.ResJoinInvoice, error)
	GetInvoiceById(id_cust, id_invoice string) (*models.ResJoinInvoice, error)
}

type invoiceService struct {
	repo repositories.InvoiceRepo
}

func NewInvoiceService(repo repositories.InvoiceRepo) InvoiceService {
	return &invoiceService{repo: repo}
}

func (s *invoiceService) Create(id_cust string, products []models.InvoiceDetail) (*models.ResJoinInvoice, error) {
	custId, err := strconv.Atoi(id_cust)
	if err != nil {
		return nil, err
	}

	inv, invdtl, err := s.repo.Create(custId, products)
	if err != nil {
		return nil, err
	}

	resInv := &models.ResJoinInvoice{
		Inv:    *inv,
		InvDtl: invdtl,
	}

	return resInv, err
}

func (s *invoiceService) GetInvoice(id_cust string) (*models.ResJoinInvoice, error) {
	return nil, nil
}
func (s *invoiceService) GetInvoiceById(id_cust, id_invoice string) (*models.ResJoinInvoice, error) {
	return nil, nil
}
