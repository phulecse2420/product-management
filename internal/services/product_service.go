package services

import (
	"errors"
	"pm/internal/models"
	"pm/internal/repositories"
)

var (
	ErrNotFound         = errors.New("product not found")
	ErrNameRequired     = errors.New("name is required")
	ErrNameTooShort     = errors.New("name must be at least 3 characters")
	ErrPriceRequired    = errors.New("price must be greater than 0")
	ErrQuantityNegative = errors.New("quantity must be greater than or equal to 0")
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func validate(name string, price float64, qty int) error {
	if name == "" {
		return ErrNameRequired
	}
	if len(name) < 3 {
		return ErrNameTooShort
	}
	if price <= 0 {
		return ErrPriceRequired
	}
	if qty < 0 {
		return ErrQuantityNegative
	}
	return nil
}

func (s *ProductService) Create(inp models.CreateProductInput) (*models.Product, error) {
	if err := validate(inp.Name, inp.Price, inp.Quantity); err != nil {
		return nil, err
	}
	return s.repo.Create(inp)
}

func (s *ProductService) List(keyword string) ([]models.Product, error) {
	return s.repo.List(keyword)
}

func (s *ProductService) GetByID(id int64) (*models.Product, error) {
	p, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *ProductService) Update(id int64, inp models.UpdateProductInput) (*models.Product, error) {
	if err := validate(inp.Name, inp.Price, inp.Quantity); err != nil {
		return nil, err
	}
	p, err := s.repo.Update(id, inp)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, ErrNotFound
	}
	return p, nil
}

func (s *ProductService) Delete(id int64) error {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotFound
	}
	return nil
}
