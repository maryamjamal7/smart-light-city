package storage

import (
	"context"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
	"gorm.io/gorm"
)

type commandRepo struct {
	db *gorm.DB
}

// Constructor
func NewCommandRepository(db *gorm.DB) ports.CommandRepository {
	return &commandRepo{db}
}

// MarkExecuted sets the ExecutedAt timestamp to NOW
func (r *commandRepo) MarkExecuted(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&model.Command{}).
		Where("id = ?", id).
		Update("executed_at", time.Now()).Error
}

// Create a new command
func (r *commandRepo) Create(ctx context.Context, cmd *model.Command) error {
	return r.db.WithContext(ctx).Create(cmd).Error
}

// Get command by ID
func (r *commandRepo) GetByID(ctx context.Context, id uint) (*model.Command, error) {
	var cmd model.Command
	err := r.db.WithContext(ctx).First(&cmd, id).Error
	return &cmd, err
}

// List all commands
func (r *commandRepo) List(ctx context.Context) ([]model.Command, error) {
	var cmds []model.Command
	err := r.db.WithContext(ctx).
		Order("scheduled_for DESC").
		Find(&cmds).Error
	return cmds, err
}

// List commands that are scheduled and not yet executed
func (r *commandRepo) ListPending(ctx context.Context, now time.Time) ([]model.Command, error) {
	var cmds []model.Command
	err := r.db.WithContext(ctx).
		Where("scheduled_for <= ? AND executed_at IS NULL", now).
		Find(&cmds).Error
	return cmds, err
}
