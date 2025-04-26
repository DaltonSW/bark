// Package bark provides a colorful and stylish logging interface
// built on top of Charmbracelet's log and lipgloss packages.
// It supports Info, Warn, Error, Debug, and Fatal levels, with custom colors and formats.
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

// BarkOptions specifies configuration for colors and time formatting.
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

	if opts.DebugHex != "" {
		merge.DebugHex = opts.DebugHex
	} else {
		merge.DebugHex = defaultOptions.DebugHex
	}

	if opts.TimeFormat != "" {
		merge.TimeFormat = opts.TimeFormat
	} else {
		merge.TimeFormat = defaultOptions.TimeFormat
	}

	return merge
}

// Init initializes the logging system with the provided BarkOptions.
// If any fields are omitted, defaults are used.
// This must be called before using the other logging functions.
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
	stdLogger.SetTimeFormat(mergedOpts.TimeFormat)
	stdLogger.SetReportTimestamp(true)

	loggers = append(loggers, stdLogger)
}

// SetDebugLevel sets the log verbosity.
// When v is true, debug messages are shown. Otherwise, only Info and above are logged.
func SetDebugLevel(v bool) {
	var level log.Level
	if v {
		level = log.DebugLevel
	} else {
		level = log.InfoLevel
	}

	for _, logger := range loggers {
		logger.SetLevel(level)
	}
}

// Info logs a message at Info level.
func Info(msg string) {
	for _, logger := range loggers {
		logger.Info(msg)
	}
}

// Info logs a formatted message at Info level.
func Infof(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Infof(formatMsg, vals...)
	}
}

// Warn logs a message at Warn level.
func Warn(msg string) {
	for _, logger := range loggers {
		logger.Warn(msg)
	}
}

// Warnf logs a formatted message at Warn level.
func Warnf(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Warnf(formatMsg, vals...)
	}
}

// Error logs a message at Error level.
func Error(msg string) {
	for _, logger := range loggers {
		logger.Error(msg)
	}
}

// Errorf logs a formatted message at Error level.
func Errorf(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Errorf(formatMsg, vals...)
	}
}

// Fatal logs a message at Fatal level and terminates the program.
func Fatal(msg string) {
	for _, logger := range loggers {
		logger.Fatal(msg)
	}
}

// Fatalf logs a formatted message at Fatal level and terminates the program.
func Fatalf(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Fatalf(formatMsg, vals...)
	}
}

// Debug logs a message at Debug level.
func Debug(msg string) {
	for _, logger := range loggers {
		logger.Debug(msg)
	}
}

// Debugf logs a formatted message at Debug level.
func Debugf(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Debugf(formatMsg, vals...)
	}
}

// DebugAndWait logs a Debug message and waits for the user to press Enter.
// Useful for debugging program flow.
func DebugAndWait(msg string) {
	for _, logger := range loggers {
		logger.Debug(fmt.Sprintf("%v (󰌑)", msg))
	}

	fmt.Scanln()
}

// DebugfAndWait logs a formatted Debug message and waits for the user to press Enter.
// Useful for debugging program flow.
func DebugfAndWait(formatMsg string, vals ...any) {
	for _, logger := range loggers {
		logger.Debugf(fmt.Sprintf("%v (󰌑)", formatMsg), vals...)
	}

	fmt.Scanln()
}

// NewError creates a new error with the provided message.
func NewError(msg string) error {
	return errors.New(msg)
}

// NewErrorf creates a new formatted error using the format string and arguments.
func NewErrorf(formatMsg string, vals ...any) error {
	return fmt.Errorf(formatMsg, vals...)
}
