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

package stdlogger

import (
	"log"

	"github.com/yosida95/chame/pkg/chame"
)

type stdlogger struct{}

var Logger chame.Logger = stdlogger{}

func (stdlogger) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG]"+format, v...)
}
func (stdlogger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO]"+format, v...)
}

func (stdlogger) Warningf(format string, v ...interface{}) {
	log.Printf("[WARN]"+format, v...)
}

func (stdlogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR]"+format, v...)
}
