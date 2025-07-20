package ports

import (
	"context"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/model"
)

type AreaRepository interface {
	Create(ctx context.Context, area *model.Area) error
	GetByID(ctx context.Context, id uint) (*model.Area, error)
	ListCities(ctx context.Context) ([]model.Area, error)
	ListZonesByCityID(ctx context.Context, cityID uint) ([]model.Area, error)
	Delete(ctx context.Context, id uint) error
}

type LumiereRepository interface {
	Create(ctx context.Context, lumiere *model.Lumiere) error
	GetByID(ctx context.Context, id uint) (*model.Lumiere, error)
	ListByAreaID(ctx context.Context, areaID uint) ([]model.Lumiere, error)
	UpdatePowerAndDim(ctx context.Context, id uint, power bool, dim int) error
	Delete(ctx context.Context, id uint) error
}

type CommandRepository interface {
	Create(ctx context.Context, cmd *model.Command) error
	ListPending(ctx context.Context, currentTime time.Time) ([]model.Command, error) // ✅ FIXED
	MarkExecuted(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Command, error)
	List(ctx context.Context) ([]model.Command, error)
}
