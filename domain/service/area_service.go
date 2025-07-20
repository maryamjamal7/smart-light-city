package service

import (
	"context"
	"errors"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
)

type AreaService struct {
	repo ports.AreaRepository
}

func NewAreaService(repo ports.AreaRepository) *AreaService {
	return &AreaService{repo: repo}
}

func (s *AreaService) CreateArea(ctx context.Context, area *model.Area) error {
	if area.Name == "" {
		return errors.New("area name is required")
	}
	if area.Type == "" {
		area.Type = "zone"
	}
	if area.Type != "city" && area.Type != "zone" {
		return errors.New("invalid area type")
	}
	return s.repo.Create(ctx, area)
}

func (s *AreaService) ListCities(ctx context.Context) ([]model.Area, error) {
	return s.repo.ListCities(ctx)
}

func (s *AreaService) ListZonesByCityID(ctx context.Context, cityID uint) ([]model.Area, error) {
	return s.repo.ListZonesByCityID(ctx, cityID)
}
