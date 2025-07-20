package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/maryamjamal7/smart-light-city/domain/service"
)

type DkronScheduler struct {
	cmdService *service.CommandService
	ticker     *time.Ticker
	stopChan   chan struct{}
}

// Constructor
func NewDkronScheduler(cmdService *service.CommandService) *DkronScheduler {
	return &DkronScheduler{
		cmdService: cmdService,
		ticker:     time.NewTicker(10 * time.Second), // every 10 seconds
		stopChan:   make(chan struct{}),
	}
}

// Start begins the periodic execution loop
func (s *DkronScheduler) Start() {
	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.runScheduledJobs()
			case <-s.stopChan:
				log.Println("🛑 Dkron scheduler stopped")
				s.ticker.Stop()
				return
			}
		}
	}()
	log.Println("✅ Dkron scheduler started")
}

// Stop ends the periodic job checking
func (s *DkronScheduler) Stop() {
	close(s.stopChan)
}

// Actual execution logic
func (s *DkronScheduler) runScheduledJobs() {
	ctx := context.Background()
	now := time.Now()

	commands, err := s.cmdService.GetPendingCommands(ctx, now)
	if err != nil {
		log.Printf("⚠️ Failed to fetch pending commands: %v", err)
		return
	}

	for _, cmd := range commands {
		log.Printf("⏰ Executing command ID %d: %s", cmd.ID, string(cmd.CommandData))

		// Call the service logic to apply the command
		if err := s.cmdService.ExecuteCommand(ctx, &cmd); err != nil {
			log.Printf("❌ Failed to execute command ID %d: %v", cmd.ID, err)
			continue
		}

		log.Printf("✅ Executed command ID %d", cmd.ID)
	}
}
