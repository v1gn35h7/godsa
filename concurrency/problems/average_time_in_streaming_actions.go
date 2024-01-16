package problems

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Problem Statement
//   You need to be able to accept "actions" from a json string and store them to calculate the average time to complete.
//   E.g. "{'jump': 100'}". As you get more actions in, calculate the running average.
//   Add an additional method to return the average time of all the actions in a json string.
//   Everything needs to be handle concurrent accesses.

// Action represents an action with a name and time.
type Action struct {
	Name string `json:"name"`
	Time int    `json:"time"`
}

// ActionStats represents the statistics for an action.
type ActionStats struct {
	TotalTime int
	Count     int
}

// ActionsTracker tracks actions and calculates running averages.
type ActionsTracker struct {
	mu          sync.Mutex
	actionStats map[string]ActionStats
}

// NewActionsTracker creates a new ActionsTracker.
func NewActionsTracker() *ActionsTracker {
	return &ActionsTracker{
		actionStats: make(map[string]ActionStats),
	}
}

// AddAction adds an action and updates the running average.
func (at *ActionsTracker) AddAction(action Action) {
	at.mu.Lock()
	defer at.mu.Unlock()

	stats := at.actionStats[action.Name]
	stats.TotalTime += action.Time
	stats.Count++
	at.actionStats[action.Name] = stats
}

// GetAverage returns the running average for a given action.
func (at *ActionsTracker) GetAverage(actionName string) float64 {
	at.mu.Lock()
	defer at.mu.Unlock()

	stats := at.actionStats[actionName]
	if stats.Count == 0 {
		return 0
	}

	return float64(stats.TotalTime) / float64(stats.Count)
}

// HandleJSONAction handles incoming JSON strings representing actions.
func (at *ActionsTracker) HandleJSONAction(jsonStr string) error {
	var action Action
	err := json.Unmarshal([]byte(jsonStr), &action)
	if err != nil {
		return err
	}

	at.AddAction(action)
	return nil
}

// GetTotalAverage returns the overall average of all actions.
func (at *ActionsTracker) GetTotalAverage() float64 {
	at.mu.Lock()
	defer at.mu.Unlock()

	var totalSum, totalCount int
	for _, stats := range at.actionStats {
		totalSum += stats.TotalTime
		totalCount += stats.Count
	}

	if totalCount == 0 {
		return 0
	}

	return float64(totalSum) / float64(totalCount)
}

func AverageInStream() {
	// Create a new ActionsTracker
	tracker := NewActionsTracker()

	// Simulate concurrent JSON action handling
	for i := 0; i < 1000; i++ {
		go func() {
			// Simulate an action every second
			time.Sleep(time.Second)
			action := Action{Name: "jump", Time: 100}
			jsonStr, _ := json.Marshal(action)
			tracker.HandleJSONAction(string(jsonStr))
		}()
	}

	// Simulate concurrent JSON action handling for another action
	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(time.Second)
			action := Action{Name: "run", Time: 50}
			jsonStr, _ := json.Marshal(action)
			tracker.HandleJSONAction(string(jsonStr))
		}()
	}

	// Simulate some time passing
	time.Sleep(5 * time.Second)

	// Get and print the average for "jump"
	jumpAvg := tracker.GetAverage("jump")
	fmt.Printf("Average time for 'jump': %.2f\n", jumpAvg)

	// Get and print the overall average
	totalAvg := tracker.GetTotalAverage()
	fmt.Printf("Overall average: %.2f\n", totalAvg)
}
