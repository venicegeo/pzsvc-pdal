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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/venicegeo/pzsvc-pdal/objects"
	"github.com/venicegeo/pzsvc-pdal/utils"
)

// InfoOptions defines options for Info.
type InfoOptions struct {
	Boundary bool `json:"boundary"`
	Metadata bool `json:"metadata"`
	Schema   bool `json:"schema"`
}

// NewInfoOptions constructs infoOptions with default values.
func NewInfoOptions() *InfoOptions {
	return &InfoOptions{Boundary: false, Metadata: false, Schema: false}
}

// InfoFunction implements pdal info.
func InfoFunction(w http.ResponseWriter, r *http.Request,
	res *objects.JobOutput, msg objects.JobInput, i, o string) {
	var args []string
	args = append(args, *msg.Function)
	args = append(args, i)

	if msg.Options != nil {
		opts := NewInfoOptions()
		if err := json.Unmarshal(*msg.Options, &opts); err != nil {
			utils.BadRequest(w, r, *res, err.Error())
			return
		}
		if opts.Boundary {
			args = append(args, "--boundary")
		}
		if opts.Metadata {
			args = append(args, "--metadata")
		}
		if opts.Schema {
			args = append(args, "--schema")
		}
	}

	out, _ := exec.Command("pdal", args...).CombinedOutput()

	// Trim whitespace
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, out); err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(buffer.Bytes(), &res.Response); err != nil {
		log.Fatal(err)
	}
}
