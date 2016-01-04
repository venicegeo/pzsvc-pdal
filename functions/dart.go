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
	"fmt"
	"net/http"
	"os/exec"

	"github.com/venicegeo/pzsvc-pdal/objects"
)

// DartFunction implements pdal height.
func DartFunction(w http.ResponseWriter, r *http.Request,
	res *objects.JobOutput, msg objects.JobInput, i, o string) {
	out, err := exec.Command("pdal", "translate", i, o,
		"dartsample", "-v10", "--debug").CombinedOutput()

	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err.Error())
	}
}