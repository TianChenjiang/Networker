package log

import (
	"context"
	"io"
	"os"
	"time"
)

var _defaultStdout = NewStdout()

// StdoutHandler stdout log middleware
type StdoutHandler struct {
	out    io.Writer
	err    io.Writer
	render Render
}

// NewStdout create a stdout log middleware
func NewStdout() *StdoutHandler {
	return &StdoutHandler{
		out:    os.Stdout,
		err:    os.Stderr,
		render: newPatternRender("[%D %T] [%s] %M"),
	}
}

// Log stdout loging, only for developing env.
func (h *StdoutHandler) Log(ctx context.Context, lv Level, args ...D) {
	d := toMap(args...)
	d[_time] = time.Now().Format(_timeFormat)
	if lv <= _infoLevel {
		h.render.Render(h.out, d)
	} else {
		h.render.Render(h.err, d)
	}
	h.out.Write([]byte("\n"))
}

// Close stdout loging
func (h *StdoutHandler) Close() error {
	return nil
}

// SetFormat set stdout log output format
// %T time format at "15:04:05.999"
// %t time format at "15:04:05"
// %D data format at "2006/01/02"
// %d data format at "01/02"
// %L log level e.g. INFO WARN ERROR
// %f function name and line number e.g. model.Get:121
// %i instance id
// %e deploy env e.g. dev uat fat prod
// %z zone
// %S full file name and line number: /a/b/c/d.go:23
// %s final file name element and line number: d.go:23
// %M log message and additional fields: key=value this is log message
func (h *StdoutHandler) SetFormat(format string) {
	h.render = newPatternRender(format)
}
