package format

import (
	"bufio"
	"time"

	"github.com/jeromer/syslogparser"
)

type LogParts map[string]interface{}

type LogParser interface {
	Parse() error
	Dump() LogParts
	WithTimestampFormat(string)
	WithLocation(*time.Location)
	WithHostname(string)
	WithTag(string)
}

type Format interface {
	GetParser([]byte) LogParser
	GetSplitFunc() bufio.SplitFunc
}

type parserWrapper struct {
	syslogparser.LogParser
}

func (w *parserWrapper) Dump() LogParts {
	return LogParts(w.LogParser.Dump())
}

func (w *parserWrapper) Parse() error {
	return w.LogParser.Parse()
}
func (w *parserWrapper) WithTimestampFormat(s string) {
	w.LogParser.WithTimestampFormat(s)
}
func (w *parserWrapper) WithLocation(l *time.Location) {
	w.LogParser.WithLocation(l)
}
func (w *parserWrapper) WithHostname(hostname string) {
	w.LogParser.WithHostname(hostname)
}
func (w *parserWrapper) WithTag(t string) {
	w.LogParser.WithTag(t)
}
