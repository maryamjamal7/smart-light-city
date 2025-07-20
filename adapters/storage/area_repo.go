// package storage

// import (
// 	"context"

// 	"github.com/maryamjamal7/smart-light-city/domain/model"
// 	"github.com/maryamjamal7/smart-light-city/domain/ports"
// 	"gorm.io/gorm"
// )

// type areaRepo struct {
// 	db *gorm.DB
// }

// // NewAreaRepository creates a new instance of AreaRepository
// func NewAreaRepository(db *gorm.DB) ports.AreaRepository {
// 	return &areaRepo{db: db}
// }

// // Create inserts a new Area (city or zone)
// func (r *areaRepo) Create(ctx context.Context, area *model.Area) error {
// 	return r.db.WithContext(ctx).Create(area).Error
// }

// // GetByID returns an area by its ID
// func (r *areaRepo) GetByID(ctx context.Context, id uint) (*model.Area, error) {
// 	var area model.Area
// 	err := r.db.WithContext(ctx).First(&area, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &area, nil
// }

// // ListCities returns all areas with type = 'city'
// func (r *areaRepo) ListCities(ctx context.Context) ([]model.Area, error) {
// 	var cities []model.Area
// 	err := r.db.WithContext(ctx).Where("type = ?", "city").Find(&cities).Error
// 	return cities, err
// }

// // ListZonesByCityID returns all zones that belong to a specific city
// func (r *areaRepo) ListZonesByCityID(ctx context.Context, cityID uint) ([]model.Area, error) {
// 	var zones []model.Area
// 	err := r.db.WithContext(ctx).Where("city_id = ?", cityID).Find(&zones).Error
// 	return zones, err
// }

// // Delete removes an area by ID
//
//	func (r *areaRepo) Delete(ctx context.Context, id uint) error {
//		return r.db.WithContext(ctx).Delete(&model.Area{}, id).Error
//	}
package storage

import (
	"context"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
	"gorm.io/gorm"
)

type areaRepo struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) ports.AreaRepository {
	return &areaRepo{db: db}
}

func (r *areaRepo) Create(ctx context.Context, area *model.Area) error {
	return r.db.WithContext(ctx).Create(area).Error
}

func (r *areaRepo) GetByID(ctx context.Context, id uint) (*model.Area, error) {
	var area model.Area
	err := r.db.WithContext(ctx).First(&area, id).Error
	return &area, err
}

func (r *areaRepo) ListCities(ctx context.Context) ([]model.Area, error) {
	var cities []model.Area
	err := r.db.WithContext(ctx).Where("type = ?", "city").Find(&cities).Error
	return cities, err
}

func (r *areaRepo) ListZonesByCityID(ctx context.Context, cityID uint) ([]model.Area, error) {
	var zones []model.Area
	err := r.db.WithContext(ctx).Where("city_id = ?", cityID).Find(&zones).Error
	return zones, err
}

func (r *areaRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Area{}, id).Error
}
