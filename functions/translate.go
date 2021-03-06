/*
Copyright 2015-2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package functions

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

/*
TranslateOptions defines options for the Translate function.

Assume the following command:

	$ pdal translate <input> <output> [args]

The Translate function takes a single option "args" - a string that completes the command-line invocation of the PDAL CLI.

See http://www.pdal.io/apps.html#translate-command for more on the PDAL translate command.
*/
type TranslateOptions struct {
	Args string `json:"args"`
}

// NewTranslateOptions constructs TranslateOptions with default values.
func NewTranslateOptions() *TranslateOptions {
	return &TranslateOptions{Args: ""}
}

// Translate implements pdal translate.
func Translate(i, o string, options *json.RawMessage) ([]byte, error) {
	opts := NewTranslateOptions()
	if options != nil {
		if err := json.Unmarshal(*options, &opts); err != nil {
			return nil, err
		}
	}

	optArgs := strings.Split(opts.Args, " ")

	var args []string
	args = append(args, "translate")
	args = append(args, i)
	args = append(args, o)
	args = append(args, optArgs...)
	args = append(args, "-v", "10", "--debug")

	out, err := exec.Command("pdal", args...).CombinedOutput()

	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
