package service

import (
	"context"
	"errors"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
)

type LumiereService struct {
	repo ports.LumiereRepository
}

func NewLumiereService(repo ports.LumiereRepository) *LumiereService {
	return &LumiereService{repo: repo}
}

func (s *LumiereService) Create(ctx context.Context, l *model.Lumiere) error {
	if l.Dim < 0 || l.Dim > 100 {
		return errors.New("dim must be between 0 and 100")
	}
	return s.repo.Create(ctx, l)
}

func (s *LumiereService) UpdateState(ctx context.Context, id uint, power bool, dim int) error {
	if dim < 0 || dim > 100 {
		return errors.New("dim must be between 0 and 100")
	}
	return s.repo.UpdatePowerAndDim(ctx, id, power, dim)
}

func (s *LumiereService) ListByArea(ctx context.Context, areaID uint) ([]model.Lumiere, error) {
	return s.repo.ListByAreaID(ctx, areaID)
}

func (s *LumiereService) GetByID(ctx context.Context, id uint) (*model.Lumiere, error) {
	return s.repo.GetByID(ctx, id)
}
