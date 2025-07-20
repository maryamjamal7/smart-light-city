// package storage

// import (
// 	"context"

// 	"github.com/maryamjamal7/smart-light-city/domain/model"
// 	"github.com/maryamjamal7/smart-light-city/domain/ports"
// 	"gorm.io/gorm"
// )

// type lumiereRepo struct {
// 	db *gorm.DB
// }

// func NewLumiereRepository(db *gorm.DB) ports.LumiereRepository {
// 	return &lumiereRepo{db: db}
// }

// func (r *lumiereRepo) Create(ctx context.Context, l *model.Lumiere) error {
// 	return r.db.WithContext(ctx).Create(l).Error
// }

// func (r *lumiereRepo) GetByID(ctx context.Context, id uint) (*model.Lumiere, error) {
// 	var l model.Lumiere
// 	err := r.db.WithContext(ctx).Preload("Area").First(&l, id).Error
// 	return &l, err
// }

// func (r *lumiereRepo) ListByAreaID(ctx context.Context, areaID uint) ([]model.Lumiere, error) {
// 	var list []model.Lumiere
// 	err := r.db.WithContext(ctx).Where("area_id = ?", areaID).Find(&list).Error
// 	return list, err
// }

// func (r *lumiereRepo) UpdatePowerAndDim(ctx context.Context, id uint, power bool, dim int) error {
// 	return r.db.WithContext(ctx).
// 		Model(&model.Lumiere{}).
// 		Where("id = ?", id).
// 		Updates(map[string]interface{}{"power": power, "dim": dim}).Error
// }

//	func (r *lumiereRepo) Delete(ctx context.Context, id uint) error {
//		return r.db.WithContext(ctx).Delete(&model.Lumiere{}, id).Error
//	}
package storage

import (
	"context"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
	"gorm.io/gorm"
)

type lumiereRepo struct {
	db *gorm.DB
}

func NewLumiereRepository(db *gorm.DB) ports.LumiereRepository {
	return &lumiereRepo{db}
}

func (r *lumiereRepo) Create(ctx context.Context, l *model.Lumiere) error {
	return r.db.WithContext(ctx).Create(l).Error
}

func (r *lumiereRepo) GetByID(ctx context.Context, id uint) (*model.Lumiere, error) {
	var l model.Lumiere
	err := r.db.WithContext(ctx).First(&l, id).Error
	return &l, err
}

func (r *lumiereRepo) ListByAreaID(ctx context.Context, areaID uint) ([]model.Lumiere, error) {
	var lights []model.Lumiere
	err := r.db.WithContext(ctx).Where("area_id = ?", areaID).Find(&lights).Error
	return lights, err
}

func (r *lumiereRepo) UpdatePowerAndDim(ctx context.Context, id uint, power bool, dim int) error {
	return r.db.WithContext(ctx).Model(&model.Lumiere{}).Where("id = ?", id).
		Updates(map[string]interface{}{"power": power, "dim": dim}).Error
}

func (r *lumiereRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Lumiere{}, id).Error
}
