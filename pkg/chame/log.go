// Copyright 2017 Kohei YOSHIDA <https://yosida95.com/>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chame

import (
	"golang.org/x/net/context"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type loggerKey struct{}

func NewContextWithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, log)
}

func LoggerFromContext(ctx context.Context) Logger {
	if log, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return log
	}
	return DiscardLogger
}

type discardLogger struct{}

var DiscardLogger = &discardLogger{}

func (discardLogger) Debugf(format string, v ...interface{})   {}
func (discardLogger) Infof(format string, v ...interface{})    {}
func (discardLogger) Warningf(format string, v ...interface{}) {}
func (discardLogger) Errorf(format string, v ...interface{})   {}