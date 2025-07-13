package cmd

import (
	"fmt"
	"os"
	"typeracer_cli/internal/config"
	"typeracer_cli/internal/game"
	"typeracer_cli/internal/ui"
	"typeracer_cli/pkg/words"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	wordCount  int
	difficulty string
	timeLimit  int
)

var rootCmd = &cobra.Command{
	Use:   "typeracer",
	Short: "A CLI typing speed test application",
	Long:  "TypeRacer CLI is a production-grade command-line typing speed test application. Test your typing speed and accuracy with various difficulty levels and customizations",
	Run:   runTypingTest,
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "view typing statistics",
	Long:  "Display your typing performance statistics and history",
	Run:   showStats,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure application settings",
	Long:  "View and modify application configuration",
	Run:   showConfig,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.typeracer.yaml)")
	rootCmd.Flags().IntVarP(&wordCount, "words", "w", 20, "number of words to type")
	rootCmd.Flags().StringVarP(&difficulty, "difficulty", "d", "medium", "difficulty level (easy, medium, hard)")
	rootCmd.Flags().IntVarP(&timeLimit, "time", "t", 60, "time limit in seconds(0 for no limit)")

	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(configCmd)

	viper.BindPFlag("words", rootCmd.Flags().Lookup("words"))
	viper.BindPFlag("difficulty", rootCmd.Flags().Lookup("difficulty"))
	viper.BindPFlag("time", rootCmd.Flags().Lookup("time"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".typeracer")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "using config file: ", viper.ConfigFileUsed())
	}
}

func runTypingTest(cmd *cobra.Command, args []string) {
	cfg := config.Load()

	if cmd.Flags().Changed("words") {
		cfg.WordCount = wordCount
	}
	if cmd.Flags().Changed("difficulty") {
		cfg.Difficulty = difficulty
	}
	if cmd.Flags().Changed("time") {
		cfg.TimeLimit = timeLimit
	}

	generator := words.NewGenerator(cfg.Difficulty)
	testWords := generator.Generate(cfg.WordCount)

	engine := game.NewEngine(cfg, testWords)

	display := ui.NewDisplay()
	inputHandler := ui.NewInputHandler()

	display.ShowWelcome(cfg)
	display.ShowInstructions()

	fmt.Println("\nPress Enter to start...")
	inputHandler.WaitForEnter()

	result := engine.Start(display, inputHandler)

	uiResult := &ui.Result{
		WPM:          result.WPM,
		Accuracy:     result.Accuracy,
		TotalWords:   result.TotalWords,
		CorrectWords: result.CorrectWords,
		Errors:       result.Errors,
		Duration:     result.Duration,
		Timestamp:    result.Timestamp,
	}

	display.ShowResults(uiResult)

	stats := game.NewStats()
	stats.Save(result)
}

func showStats(cmd *cobra.Command, args []string) {
	stats := game.NewStats()
	history := stats.Load()
	uiHistory := &ui.StatsHistory{
		Tests: make([]ui.Result, len(history.Tests)),
	}

	for i, test := range history.Tests {
		uiHistory.Tests[i] = ui.Result{
			WPM:          test.WPM,
			Accuracy:     test.Accuracy,
			TotalWords:   test.TotalWords,
			CorrectWords: test.CorrectWords,
			Errors:       test.Errors,
			Duration:     test.Duration,
			Timestamp:    test.Timestamp,
		}
	}

	display := ui.NewDisplay()
	display.ShowStats(uiHistory)
}

func showConfig(cmd *cobra.Command, args []string) {
	cfg := config.Load()
	display := ui.NewDisplay()
	display.ShowConfig(cfg)
}
