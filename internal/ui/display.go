package ui

import (
	"fmt"
	"time"
	"typeracer_cli/internal/config"

	"github.com/fatih/color"
)

type Display struct {
	green  *color.Color
	red    *color.Color
	yellow *color.Color
	blue   *color.Color
	cyan   *color.Color
}

func NewDisplay() *Display {
	return &Display{
		green:  color.New(color.FgGreen),
		red:    color.New(color.FgRed),
		yellow: color.New(color.FgYellow),
		blue:   color.New(color.FgBlue),
		cyan:   color.New(color.FgCyan),
	}
}

type Result struct {
	WPM          float64
	Accuracy     float64
	TotalWords   int
	CorrectWords int
	Errors       int
	Duration     time.Duration
	Timestamp    time.Time
}

type StatsHistory struct {
	Tests []Result `json:"tests"`
}

func (d *Display) ShowWelcome(cfg *config.Config) {
	d.cyan.Println("═══════════════════════════════════════")
	d.cyan.Println("         TypeRacer CLI v1.0")
	d.cyan.Println("═══════════════════════════════════════")
	fmt.Printf("Configuration:\n")
	fmt.Printf("  Words: %d\n", cfg.WordCount)
	fmt.Printf("  Difficulty: %s\n", cfg.Difficulty)
	if cfg.TimeLimit > 0 {
		fmt.Printf("  Time Limit: %d seconds\n", cfg.TimeLimit)
	} else {
		fmt.Printf("  Time Limit: No limit\n")
	}
	fmt.Println()
}

func (d *Display) ShowInstructions() {
	d.yellow.Println("Instructions:")
	fmt.Println("• Type the words as they appear")
	fmt.Println("• Press Enter to finish early")
	fmt.Println("• Backspace is not supported (realistic typing)")
	fmt.Println("• Focus on accuracy over speed")
}

func (d *Display) UpdateProgress(target, typed string, errors int) {
	fmt.Print("\033[H\033[2J")

	fmt.Printf("Progress: %d/%d characters | Errors: %d\n\n", len(typed), len(target), errors)
	for i, char := range target {
		if i < len(typed) {
			typedChar := rune(typed[i])
			if typedChar == char {
				d.green.Printf("%c", char)
			} else {
				d.red.Printf("%c", char)
			}
		} else if i == len(typed) {
			d.yellow.Printf("%c", char)
		} else {
			fmt.Printf("%c", char)
		}
	}
	fmt.Println()
}

func (d *Display) ShowResults(result *Result) {
	fmt.Println()
	d.cyan.Println("═══════════════════════════════════════")
	d.cyan.Println("              RESULTS")
	d.cyan.Println("═══════════════════════════════════════")

	d.green.Printf("WPM: %.1f\n", result.WPM)
	d.blue.Printf("Accuracy: %.1f%%\n", result.Accuracy)
	fmt.Printf("Correct Words: %d/%d\n", result.CorrectWords, result.TotalWords)
	fmt.Printf("Errors: %d\n", result.Errors)
	fmt.Printf("Duration: %.1f seconds\n", result.Duration.Seconds())

	fmt.Println()
	d.yellow.Println("Performance:")
	if result.WPM >= 60 {
		d.green.Println("🚀 Excellent! You're a fast typist!")
	} else if result.WPM >= 40 {
		d.blue.Println("👍 Good job! Keep practicing!")
	} else {
		d.yellow.Println("📚 Keep practicing to improve your speed!")
	}

	if result.Accuracy >= 95 {
		d.green.Println("🎯 Outstanding accuracy!")
	} else if result.Accuracy >= 90 {
		d.blue.Println("✅ Good accuracy!")
	} else {
		d.yellow.Println("🎯 Focus on accuracy for better results!")
	}
}

func (d *Display) ShowStats(history *StatsHistory) {
	if len(history.Tests) == 0 {
		d.yellow.Println("No typing tests completed yet. Run a test first!")
		return
	}

	d.cyan.Println("═══════════════════════════════════════")
	d.cyan.Println("           STATISTICS")
	d.cyan.Println("═══════════════════════════════════════")

	var totalWPM, totalAccuracy float64
	var bestWPM, bestAccuracy float64

	for i, test := range history.Tests {
		totalWPM += test.WPM
		totalAccuracy += test.Accuracy

		if i == 0 || test.WPM > bestWPM {
			bestWPM = test.WPM
		}
		if i == 0 || test.Accuracy > bestAccuracy {
			bestAccuracy = test.Accuracy
		}
	}

	avgWPM := totalWPM / float64(len(history.Tests))
	avgAccuracy := totalAccuracy / float64(len(history.Tests))

	fmt.Printf("Total Tests: %d\n", len(history.Tests))
	fmt.Printf("Average WPM: %.1f\n", avgWPM)
	fmt.Printf("Best WPM: %.1f\n", bestWPM)
	fmt.Printf("Average Accuracy: %.1f%%\n", avgAccuracy)
	fmt.Printf("Best Accuracy: %.1f%%\n", bestAccuracy)

	fmt.Println("\nRecent Tests:")
	recent := history.Tests
	if len(recent) > 5 {
		recent = recent[len(recent)-5:]
	}

	for i, test := range recent {
		fmt.Printf("%d. %.1f WPM, %.1f%% accuracy (%s)\n",
			len(recent)-i, test.WPM, test.Accuracy,
			test.Timestamp.Format("2006-01-02 15:04"))
	}
}

func (d *Display) ShowConfig(cfg *config.Config) {
	d.cyan.Println("═══════════════════════════════════════")
	d.cyan.Println("          CONFIGURATION")
	d.cyan.Println("═══════════════════════════════════════")

	fmt.Printf("Word Count: %d\n", cfg.WordCount)
	fmt.Printf("Difficulty: %s\n", cfg.Difficulty)
	fmt.Printf("Time Limit: %d seconds\n", cfg.TimeLimit)
	fmt.Printf("Show WPM: %t\n", cfg.ShowWPM)
	fmt.Printf("Show Errors: %t\n", cfg.ShowErrors)
}
