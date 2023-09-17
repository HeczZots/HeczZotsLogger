package logger

import "context"

type Hook interface {
	// Run runs the hook with the event.
	Run(e *Event, level Level, message string)
}
type HookFunc func(e *Event, level Level, message string)
type Event struct {
	buf       []byte
	w         LevelWriter
	level     Level
	done      func(msg string)
	stack     bool            // enable error stack trace
	ch        []Hook          // hooks from context
	skipFrame int             // The number of additional frames to skip when printing the caller.
	ctx       context.Context // Optional Go context for event
}
type Level int8

func (h HookFunc) Run(e *Event, level Level, message string) {
	h(e, level, message)
}

type LevelHook struct {
	NoLevelHook, TraceHook, DebugHook, InfoHook, WarnHook, ErrorHook, FatalHook, PanicHook Hook
}

func (h LevelHook) Run(e *Event, level Level, message string) {
	switch level {
	case TraceLevel:
		if h.TraceHook != nil {
			h.TraceHook.Run(e, level, message)
		}
	case DebugLevel:
		if h.DebugHook != nil {
			h.DebugHook.Run(e, level, message)
		}
	case InfoLevel:
		if h.InfoHook != nil {
			h.InfoHook.Run(e, level, message)
		}
	case WarnLevel:
		if h.WarnHook != nil {
			h.WarnHook.Run(e, level, message)
		}
	case ErrorLevel:
		if h.ErrorHook != nil {
			h.ErrorHook.Run(e, level, message)
		}
	case FatalLevel:
		if h.FatalHook != nil {
			h.FatalHook.Run(e, level, message)
		}
	case PanicLevel:
		if h.PanicHook != nil {
			h.PanicHook.Run(e, level, message)
		}
	case NoLevel:
		if h.NoLevelHook != nil {
			h.NoLevelHook.Run(e, level, message)
		}
	}
}

// NewLevelHook returns a new LevelHook.
func NewLevelHook() LevelHook {
	return LevelHook{}
}
