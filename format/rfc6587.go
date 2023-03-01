package format

import (
	"bufio"
	"bytes"
	"strconv"
	"time"

	"github.com/jeromer/syslogparser/parsercommon"
	"github.com/jeromer/syslogparser/rfc5424"
)

type RFC6587 struct {
	pri                   *parsercommon.Priority
	location              *time.Location
	hostname              *string
	customTag             *string
	customTimestampFormat *string
}

func NewFormatRFC6587() *RFC6587 {
	return &RFC6587{}
}

func (f *RFC6587) GetParser(line []byte) LogParser {

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

func (f *RFC6587) GetSplitFunc() bufio.SplitFunc {
	return rfc6587ScannerSplit
}

func (p *RFC6587) WithLocation(l *time.Location) {
	p.location = l
}

func (p *RFC6587) WithHostname(h string) {
	p.hostname = &h
}

func (p *RFC6587) WithTag(t string) {
	p.customTag = &t
}

func (p *RFC6587) WithTimestampFormat(s string) {
	p.customTimestampFormat = &s
}

func rfc6587ScannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ' '); i > 0 {
		pLength := data[0:i]
		length, err := strconv.Atoi(string(pLength))
		if err != nil {
			if string(data[0:1]) == "<" {
				// Assume this frame uses non-transparent-framing
				return len(data), data, nil
			}
			return 0, nil, err
		}
		end := length + i + 1
		if len(data) >= end {
			// Return the frame with the length removed
			return end, data[i+1 : end], nil
		}
	}

	// Request more data
	return 0, nil, nil
}
