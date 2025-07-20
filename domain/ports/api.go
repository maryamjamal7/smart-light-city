package ports

import (
	"context"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/model"
)

type AreaAPI interface {
	CreateArea(ctx context.Context, area *model.Area) error
	GetAreaByID(ctx context.Context, id uint) (*model.Area, error)
	ListCities(ctx context.Context) ([]model.Area, error)
	ListZonesByCityID(ctx context.Context, cityID uint) ([]model.Area, error)
	DeleteArea(ctx context.Context, id uint) error
}

type LumiereAPI interface {
	CreateLumiere(ctx context.Context, lumiere *model.Lumiere) error
	GetLumiereByID(ctx context.Context, id uint) (*model.Lumiere, error)
	ListByAreaID(ctx context.Context, areaID uint) ([]model.Lumiere, error)
	UpdateState(ctx context.Context, id uint, power bool, dim int) error
	DeleteLumiere(ctx context.Context, id uint) error
}

type CommandAPI interface {
	CreateCommand(ctx context.Context, cmd *model.Command) error
	GetCommandByID(ctx context.Context, id uint) (*model.Command, error)
	ListPendingCommands(ctx context.Context, now time.Time) ([]model.Command, error)
	MarkCommandExecuted(ctx context.Context, id uint) error
	ListAllCommands(ctx context.Context) ([]model.Command, error)
}

type CityManagerAPI interface {
	PowerOffAll(ctx context.Context) error
	ScheduleDimForAll(ctx context.Context, dim int, at time.Time) error
}
