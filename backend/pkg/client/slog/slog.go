package slog

import "github.com/gookit/slog"

func init() {
	slog.Configure(func(l *slog.SugaredLogger) {
		f := l.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
}
