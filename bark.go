package bark

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var defaultOptions BarkOptions = BarkOptions{
	InfoHex:  "#1982c4",
	WarnHex:  "#ffca3a",
	ErrorHex: "#ff595e",
	DebugHex: "#ca7df9",

	TimeFormat: "01/02 03:04:05PM",
}

var loggers []*log.Logger

type BarkOptions struct {
	InfoHex  string
	WarnHex  string
	ErrorHex string
	DebugHex string

	TimeFormat string
}

func mergeOpts(opts BarkOptions) BarkOptions {
	merge := BarkOptions{}

	if opts.InfoHex != "" {
		merge.InfoHex = opts.InfoHex
	} else {
		merge.InfoHex = defaultOptions.InfoHex
	}

	if opts.WarnHex != "" {
		merge.WarnHex = opts.WarnHex
	} else {
		merge.WarnHex = defaultOptions.WarnHex
	}

	if opts.ErrorHex != "" {
		merge.ErrorHex = opts.ErrorHex
	} else {
		merge.ErrorHex = defaultOptions.ErrorHex
	}

	if opts.InfoHex != "" {
		merge.InfoHex = opts.InfoHex
	} else {
		merge.InfoHex = defaultOptions.InfoHex
	}

	if opts.TimeFormat != "" {
		merge.TimeFormat = opts.TimeFormat
	} else {
		merge.TimeFormat = defaultOptions.TimeFormat
	}

	return merge
}

func Init(opts BarkOptions) {
	mergedOpts := mergeOpts(opts)

	loggers = make([]*log.Logger, 0)

	stdLogger := log.New(os.Stderr)
	styles := log.DefaultStyles()

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().SetString(" INFO ").Padding(0, 1).Foreground(lipgloss.Color(mergedOpts.InfoHex)).Bold(true)
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().SetString(" WARN ").Padding(0, 1).Foreground(lipgloss.Color(mergedOpts.WarnHex)).Bold(true)
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().SetString("ERROR ").Padding(0, 1).Foreground(lipgloss.Color(mergedOpts.ErrorHex)).Bold(true)
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().SetString("FATAL ").Padding(0, 1).Foreground(lipgloss.Color(mergedOpts.ErrorHex)).Bold(true)
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().SetString("DEBUG ").Padding(0, 1).Foreground(lipgloss.Color(mergedOpts.DebugHex)).Bold(true)

	stdLogger.SetStyles(styles)

	loggers = append(loggers, stdLogger)
}

func SetDebugLevel(v bool) {
	for _, logger := range loggers {
		logger.SetLevel(log.DebugLevel)
	}
}

func Info(msg string) {
	for _, logger := range loggers {
		logger.Info(msg)
	}
}

func Warn(msg string) {
	for _, logger := range loggers {
		logger.Warn(msg)
	}
}

func Error(msg string) {
	for _, logger := range loggers {
		logger.Error(msg)
	}
}

func Fatal(msg string) {
	for _, logger := range loggers {
		logger.Fatal(msg)
	}
}

func Debug(msg string) {
	for _, logger := range loggers {
		logger.Debug(msg)
	}
}

func DebugAndWait(msg string) {
	for _, logger := range loggers {
		logger.Debug(fmt.Sprintf("%v (Awaiting Return)", msg))
	}

	fmt.Scanln()
}

func NewError(msg string) error {
	return errors.New(msg)
}
