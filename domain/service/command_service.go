package service

import (
	"context"
	"errors"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/ports"
)

type CommandService struct {
	repo ports.CommandRepository
	mqtt ports.MQTTPublisher // <-- Add this field
}

func NewCommandService(repo ports.CommandRepository, mqtt ports.MQTTPublisher) *CommandService {
	return &CommandService{
		repo: repo,
		mqtt: mqtt,
	}
}

// ScheduleCommand adds a new command to the database
func (s *CommandService) ScheduleCommand(ctx context.Context, cmd *model.Command) error {
	if cmd.CommandData == nil {
		return errors.New("command data is required")
	}
	if cmd.ScheduledAt == nil {
		return errors.New("scheduled time is required")
	}
	return s.repo.Create(ctx, cmd)
}

// GetDueCommands fetches commands that should be executed
func (s *CommandService) GetDueCommands(ctx context.Context) ([]model.Command, error) {
	return s.repo.ListPending(ctx, time.Now())
}

// MarkCommandExecuted sets executed_at to now
func (s *CommandService) MarkCommandExecuted(ctx context.Context, id uint) error {
	return s.repo.MarkExecuted(ctx, id)
}
func (s *CommandService) ListCommands(ctx context.Context) ([]model.Command, error) {
	return s.repo.List(ctx)

}
func (s *CommandService) GetPendingCommands(ctx context.Context, before time.Time) ([]model.Command, error) {
	return s.repo.ListPending(ctx, before)
}
func (s *CommandService) ExecuteCommand(ctx context.Context, cmd *model.Command) error {
	// Business logic can go here if needed in future
	return s.repo.MarkExecuted(ctx, cmd.ID)
}
