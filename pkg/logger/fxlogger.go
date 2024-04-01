package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx/fxevent"
	"io"
)

type fxLogger struct {
	l zerolog.Logger
}

var _ io.Writer = (*fxLogger)(nil)

func (l fxLogger) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		p = p[0 : n-1]
	}
	l.l.Debug().Msg(string(p))
	return
}

func FxLogger() fxevent.Logger {
	logger := fxLogger{
		l: log.Logger.
			With().
			Str("event", "fx.init").
			Logger(),
	}
	return &fxevent.ConsoleLogger{
		W: logger,
	}
}
