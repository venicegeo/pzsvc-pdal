/*
Copyright 2015, RadiantBlue Technologies, Inc.

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

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/venicegeo/pzsvc-pdal/functions"
	"github.com/venicegeo/pzsvc-pdal/objects"
	"github.com/venicegeo/pzsvc-pdal/utils"
)

type functionFunc func(http.ResponseWriter, *http.Request,
	*objects.JobOutput, objects.JobInput)

func makeFunction(fn func(http.ResponseWriter, *http.Request,
	*objects.JobOutput, objects.JobInput, string)) functionFunc {
	return func(w http.ResponseWriter, r *http.Request, res *objects.JobOutput,
		msg objects.JobInput) {
		file, err := os.Create("download_file.laz")
		if err != nil {
			utils.InternalError(w, r, *res, err.Error())
			return
		}
		defer file.Close()

		err = utils.S3Download(file, msg.Source.Bucket, msg.Source.Key)
		if err != nil {
			utils.InternalError(w, r, *res, err.Error())
			return
		}
		fn(w, r, res, msg, file.Name())
	}
}

// PdalHandler handles PDAL jobs.
func PdalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var res objects.JobOutput
	res.StartedAt = time.Now()

	if r.Body == nil {
		utils.BadRequest(w, r, res, "No JSON")
		return
	}

	// Parse the incoming JSON body, and unmarshal as events.NewData struct.
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.InternalError(w, r, res, err.Error())
		return
	}

	var msg objects.JobInput
	if err := json.Unmarshal(b, &msg); err != nil {
		utils.BadRequest(w, r, res, err.Error())
		return
	}
	if msg.Function == nil {
		utils.BadRequest(w, r, res, "Must provide a function")
		return
	}

	res.Input = msg
	utils.UpdateJobManager(objects.Running, r)

	switch *msg.Function {
	case "info":
		makeFunction(functions.InfoFunction)(w, r, &res, msg)

	case "pipeline":
		fmt.Println("pipeline not implemented yet")

	case "ground":
		makeFunction(functions.GroundFunction)(w, r, &res, msg)

	case "height":
		makeFunction(functions.HeightFunction)(w, r, &res, msg)

	case "groundopts":
		_, err := exec.Command("pdal",
			"--options=filters.ground").CombinedOutput()

		if err != nil {
			utils.InternalError(w, r, res, err.Error())
			return
		}

	case "dtm":
		makeFunction(functions.DtmFunction)(w, r, &res, msg)

	/*
		I get a bad_alloc here, but only via go test. The same command run natively
		works fine.
		case "drivers":
			out, err := exec.Command("pdal",
				"--drivers").CombinedOutput()

			fmt.Println(string(out))
			if err != nil {
				utils.InternalError(w, r, res, err.Error())
				return
			}
	*/

	default:
		utils.BadRequest(w, r, res,
			"Only the info and pipeline functions are supported at this time")
		return
	}

	res.FinishedAt = time.Now()
	utils.Okay(w, r, res, "Success!")
}
