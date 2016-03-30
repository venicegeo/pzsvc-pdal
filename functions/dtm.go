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
	"os"
	"os/exec"
	"strconv"
)

// DtmOptions defines options for the Dtm function.
type DtmOptions struct {
	GridSize float64 `json:"grid_size"` // size of grid cell in XY dimensions
}

// NewDtmOptions constructs DtmOptions with default values.
func NewDtmOptions() *DtmOptions {
	return &DtmOptions{GridSize: 1.0}
}

// Dtm implements pdal dtm.
func Dtm(i, o string, options *json.RawMessage) ([]byte, error) {
	opts := NewDtmOptions()
	if options != nil {
		if err := json.Unmarshal(*options, &opts); err != nil {
			return nil, err
		}
	}

	var args []string
	args = append(args, "translate")
	args = append(args, i)
	args = append(args, "output")
	args = append(args, "ground")
	args = append(args, "--filters.ground.extract=true")
	args = append(args, "--filters.ground.classify=false")
	args = append(args, "-w", "writers.p2g")
	args = append(args, "--writers.p2g.output_type=min")
	args = append(args, "--writers.p2g.output_format=tif")
	args = append(args, "--writers.p2g.grid_dist_x="+
		strconv.FormatFloat(opts.GridSize, 'f', -1, 64))
	args = append(args, "--writers.p2g.grid_dist_y="+
		strconv.FormatFloat(opts.GridSize, 'f', -1, 64))
	args = append(args, "-v", "10", "--debug")

	out, err := exec.Command("pdal", args...).CombinedOutput()

	fmt.Println(string(out))
	if err != nil {
		return nil, err
	}

	err = os.Rename("output.min.tif", o)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
