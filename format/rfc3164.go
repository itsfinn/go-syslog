package format

import (
	"bufio"
	"time"

	"github.com/jeromer/syslogparser/parsercommon"
	"github.com/jeromer/syslogparser/rfc3164"
)

type RFC3164 struct {
	pri                   *parsercommon.Priority
	location              *time.Location
	hostname              *string
	customTag             *string
	customTimestampFormat *string
}

func NewFormatRFC3164() *RFC3164 {
	return &RFC3164{}
}

func (f *RFC3164) GetParser(line []byte) LogParser {
	parser := rfc3164.NewParser(line)
	if f.pri != nil {
		parser.WithPriority(f.pri)
	}
	if f.location != nil {
		parser.WithLocation(f.location)
	}
	if f.hostname != nil {
		parser.WithHostname(*f.hostname)
	}
	if f.customTag != nil {
		parser.WithTag(*f.customTag)
	}
	if f.customTimestampFormat != nil {
		parser.WithTimestampFormat(*f.customTimestampFormat)
	}
	return &parserWrapper{parser}
}

func (f *RFC3164) GetSplitFunc() bufio.SplitFunc {
	return nil
}
func (p *RFC3164) WithLocation(l *time.Location) {
	p.location = l
}

func (p *RFC3164) WithHostname(h string) {
	p.hostname = &h
}

func (p *RFC3164) WithTag(t string) {
	p.customTag = &t
}

func (p *RFC3164) WithTimestampFormat(s string) {
	p.customTimestampFormat = &s
}
