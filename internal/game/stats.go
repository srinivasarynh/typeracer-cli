package game

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Stats struct {
	filepath string
}

type StatsHistory struct {
	Tests []Result `json:"tests"`
}

func NewStats() *Stats {
	home, _ := os.UserHomeDir()
	filepath := filepath.Join(home, ".typeracer_stats.json")

	return &Stats{
		filepath: filepath,
	}
}

func (s *Stats) Save(result *Result) error {
	history := s.Load()
	history.Tests = append(history.Tests, *result)

	if len(history.Tests) > 100 {
		history.Tests = history.Tests[len(history.Tests)-100:]
	}

	data, err := json.MarshalIndent(history, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filepath, data, 0644)
}

func (s *Stats) Load() *StatsHistory {
	history := &StatsHistory{
		Tests: []Result{},
	}

	data, err := os.ReadFile(s.filepath)
	if err != nil {
		return history
	}

	json.Unmarshal(data, history)
	return history
}

func (s *Stats) GetAverageWPM() float64 {
	history := s.Load()
	if len(history.Tests) == 0 {
		return 0
	}
	total := 0.0
	for _, test := range history.Tests {
		total += test.WPM
	}

	return total / float64(len(history.Tests))
}

func (s *Stats) GetBestWPM() float64 {
	history := s.Load()
	if len(history.Tests) == 0 {
		return 0
	}

	best := history.Tests[0].WPM
	for _, test := range history.Tests {
		if test.WPM > best {
			best = test.WPM
		}
	}

	return best
}

func (s *Stats) GetRecentTests(days int) []Result {
	history := s.Load()
	cutoff := time.Now().AddDate(0, 0, -days)

	var recent []Result
	for _, test := range history.Tests {
		if test.Timestamp.After(cutoff) {
			recent = append(recent, test)
		}
	}

	return recent
}
