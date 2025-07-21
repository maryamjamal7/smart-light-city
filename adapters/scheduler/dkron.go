// package scheduler

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/maryamjamal7/smart-light-city/domain/service"
// )

// type DkronScheduler struct {
// 	cmdService *service.CommandService
// 	ticker     *time.Ticker
// 	stopChan   chan struct{}
// 	scheduleJobFn func(uint, time.Time) error
// }
// type DkronJob struct {
// 	Name           string            `json:"name"`
// 	Schedule       string            `json:"schedule"` // e.g., "R0/2025-07-21T12:00:00Z/PT1M"
// 	Executor       string            `json:"executor"`
// 	Owner          string            `json:"owner"`
// 	OwnerEmail     string            `json:"owner_email"`
// 	Disabled       bool              `json:"disabled"`
// 	Tags           map[string]string `json:"tags,omitempty"`
// 	Metadata       map[string]string `json:"metadata,omitempty"`
// 	ExecutorConfig map[string]string `json:"executor_config"`
// }

// // CreateDkronJob creates a Dkron job to call your internal endpoint
// func CreateDkronJob(commandID uint, runAt time.Time) error {
// 	job := DkronJob{
// 		Name:       fmt.Sprintf("cmd-%d", commandID),
// 		Schedule:   fmt.Sprintf("R1/%s/PT1M", runAt.UTC().Format(time.RFC3339)),
// 		Executor:   "http",
// 		Owner:      "system",
// 		OwnerEmail: "system@example.com",
// 		Disabled:   false,
// 		ExecutorConfig: map[string]string{
// 			"method": "POST",
// 			"url":    fmt.Sprintf("http://localhost:8081/internal/execute/%d", commandID),
// 		},
// 	}

// 	body, _ := json.Marshal(job)
// 	resp, err := http.Post("http://localhost:8080/v1/jobs", "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return fmt.Errorf("failed to create dkron job: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode >= 300 {
// 		return fmt.Errorf("failed to create job, status: %s", resp.Status)
// 	}

// 	return nil
// }

// // Constructor
// func NewDkronScheduler(cmdService *service.CommandService) *DkronScheduler {
// 	return &DkronScheduler{
// 		cmdService: cmdService,
// 		ticker:     time.NewTicker(10 * time.Second), // every 10 seconds
// 		stopChan:   make(chan struct{}),
// 	}
// }

// // Start begins the periodic execution loop
// func (s *DkronScheduler) Start() {
// 	go func() {
// 		for {
// 			select {
// 			case <-s.ticker.C:
// 				s.runScheduledJobs()
// 			case <-s.stopChan:
// 				log.Println("🛑 Dkron scheduler stopped")
// 				s.ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// 	log.Println(" Dkron scheduler started")
// }

// // Stop ends the periodic job checking
// func (s *DkronScheduler) Stop() {
// 	close(s.stopChan)
// }

// // Actual execution logic
// func (s *DkronScheduler) runScheduledJobs() {
// 	ctx := context.Background()
// 	now := time.Now()

// 	commands, err := s.cmdService.GetPendingCommands(ctx, now)
// 	if err != nil {
// 		log.Printf("⚠️ Failed to fetch pending commands: %v", err)
// 		return
// 	}

// 	for _, cmd := range commands {
// 		log.Printf("⏰ Executing command ID %d: %s", cmd.ID, string(cmd.CommandData))

// 		// Call the service logic to apply the command
// 		if err := s.cmdService.ExecuteCommand(ctx, &cmd); err != nil {
// 			log.Printf("  Failed to execute command ID %d: %v", cmd.ID, err)
// 			continue
// 		}

//			log.Printf(" Executed command ID %d", cmd.ID)
//		}
//	}
package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DkronJob struct {
	Name           string            `json:"name"`
	Schedule       string            `json:"schedule"` // e.g. R1/2025-07-21T12:00:00Z/PT1M
	Executor       string            `json:"executor"`
	Owner          string            `json:"owner"`
	OwnerEmail     string            `json:"owner_email"`
	Disabled       bool              `json:"disabled"`
	ExecutorConfig map[string]string `json:"executor_config"`
}

// ✅ Called by CommandService
func CreateDkronJob(commandID uint, runAt time.Time) error {
	job := DkronJob{
		Name:       fmt.Sprintf("cmd-%d", commandID),
		Schedule:   fmt.Sprintf("R1/%s/PT1M", runAt.UTC().Format(time.RFC3339)),
		Executor:   "http",
		Owner:      "system",
		OwnerEmail: "system@example.com",
		Disabled:   false,
		ExecutorConfig: map[string]string{
			"method": "POST",
			"url":    fmt.Sprintf("http://localhost:8081/internal/execute/%d", commandID),
		},
	}

	body, _ := json.Marshal(job)
	resp, err := http.Post("http://localhost:8080/v1/jobs", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create dkron job: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("Dkron job creation failed: %s", resp.Status)
	}
	return nil
}
