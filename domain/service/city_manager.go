package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/model"
)

var ErrInvalidDim = errors.New("dim must be between 0 and 100")

type CityManager struct {
	areaService    *AreaService
	lumiereService *LumiereService
	commandService *CommandService
}

func NewCityManager(areaSvc *AreaService, lumiereSvc *LumiereService, cmdSvc *CommandService) *CityManager {
	return &CityManager{
		areaService:    areaSvc,
		lumiereService: lumiereSvc,
		commandService: cmdSvc,
	}
}

// PowerOffAll turns off all lumières in every area and city
func (m *CityManager) PowerOffAll(ctx context.Context) error {
	cities, err := m.areaService.ListCities(ctx)
	if err != nil {
		return err
	}

	for _, city := range cities {
		zones, err := m.areaService.ListZonesByCityID(ctx, city.ID)
		if err != nil {
			return err
		}

		for _, zone := range zones {
			lights, err := m.lumiereService.ListByArea(ctx, zone.ID)
			if err != nil {
				return err
			}

			for _, light := range lights {
				err := m.lumiereService.UpdateState(ctx, light.ID, false, 0)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ScheduleDimForAll dims all lights to a given value at a future time
func (m *CityManager) ScheduleDimForAll(ctx context.Context, dim int, at time.Time) error {
	if dim < 0 || dim > 100 {
		return ErrInvalidDim
	}

	cities, err := m.areaService.ListCities(ctx)
	if err != nil {
		return err
	}

	for _, city := range cities {
		zones, err := m.areaService.ListZonesByCityID(ctx, city.ID)
		if err != nil {
			return err
		}

		for _, zone := range zones {
			lights, err := m.lumiereService.ListByArea(ctx, zone.ID)
			if err != nil {
				return err
			}

			for _, light := range lights {
				cmd := &model.Command{
					AreaID:      &zone.ID,
					LumiereID:   &light.ID,
					CommandData: buildCommandJSON(false, dim),
					ScheduledAt: &at,
				}
				err := m.commandService.ScheduleCommand(ctx, cmd)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// buildCommandJSON creates JSONB command data
func buildCommandJSON(power bool, dim int) []byte {
	obj := map[string]interface{}{
		"power": power,
		"dim":   dim,
	}
	jsonData, _ := json.Marshal(obj)
	return jsonData
}
