package format

import (
	"bufio"
	"time"

	"github.com/jeromer/syslogparser/parsercommon"
	"github.com/jeromer/syslogparser/rfc5424"
)

type RFC5424 struct {
	pri                   *parsercommon.Priority
	location              *time.Location
	hostname              *string
	customTag             *string
	customTimestampFormat *string
}

func NewFormatRFC5424() *RFC5424 {
	return &RFC5424{}
}

func (f *RFC5424) GetParser(line []byte) LogParser {

	parser := rfc5424.NewParser(line)
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

func (f *RFC5424) GetSplitFunc() bufio.SplitFunc {
	return nil
}

func (p *RFC5424) WithLocation(l *time.Location) {
	p.location = l
}

func (p *RFC5424) WithHostname(h string) {
	p.hostname = &h
}

func (p *RFC5424) WithTag(t string) {
	p.customTag = &t
}

func (p *RFC5424) WithTimestampFormat(s string) {
	p.customTimestampFormat = &s
}
