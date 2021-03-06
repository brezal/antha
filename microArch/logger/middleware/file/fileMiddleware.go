// microArch/logger/middleware/file/fileMiddleware.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
//
// For more information relating to the software or licensing issues please
// contact license@antha-lang.org or write to the Antha team c/o
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

package file

import (
	"fmt"
	"os"
	"sync"

	"github.com/antha-lang/antha/microArch/logger"
)

type fileMiddleware struct {
	fileName string
	minLevel logger.LogLevel
	sync.Mutex
}

var _ logger.LoggerMiddleware = &fileMiddleware{}

func NewFileMiddleware(fileName string, minLevel logger.LogLevel) *fileMiddleware {
	ret := new(fileMiddleware)
	ret.fileName = fileName
	ret.minLevel = minLevel
	return ret
}

func (m *fileMiddleware) Log(level logger.LogLevel, ts int64, source string, message string, extra ...interface{}) {
	m.Lock()
	defer m.Unlock()
	if level < m.minLevel {
		return
	}
	f, err := os.OpenFile(m.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprint(level, " ", message, " | ", extra))
}

func (m *fileMiddleware) Measure(ts int64, source string, message string, extra ...interface{}) {
	m.Lock()
	defer m.Unlock()
	f, err := os.OpenFile(m.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprint("[Telemetry] ", message, " | ", extra))
}

func (m *fileMiddleware) Sensor(ts int64, source string, message string, extra ...interface{}) {
	m.Lock()
	defer m.Unlock()
	f, err := os.OpenFile(m.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprint("[Sensor] ", message, " | ", extra))
}

func (m *fileMiddleware) Data(ts int64, data interface{}, extra ...interface{}) {
	m.Lock()
	defer m.Unlock()

	f, err := os.OpenFile(m.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprintf("[Data] $d | %+v | %+v\n", data, extra))
}
