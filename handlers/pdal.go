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

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/venicegeo/pzsvc-sdk-go/objects"
	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/venicegeo/pzsvc-sdk-go/utils"
	"github.com/venicegeo/pzsvc-pdal/functions"
)

// PdalHandler handles PDAL jobs.
func PdalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Create the job output message. No matter what happens, we should always be
	// able to populate the StartedAt field.
	var res objects.JobOutput
	res.StartedAt = time.Now()

	msg := utils.GetJobInput(w, r, res)

	// Throw 400 if the JobInput does not specify a function.
	if msg.Function == nil {
		utils.BadRequest(w, r, res, "Must provide a function")
		return
	}

	// If everything is okay up to this point, we will echo the JobInput in the
	// JobOutput and mark the job as Running.
	res.Input = msg
	utils.UpdateJobManager(objects.Running, r)

	// Make/execute the requested function.
	switch *msg.Function {
	case "dart":
		utils.MakeFunction(functions.DartFunction)(w, r, &res, msg)

	case "dtm":
		utils.MakeFunction(functions.DtmFunction)(w, r, &res, msg)

	case "ground":
		utils.MakeFunction(functions.GroundFunction)(w, r, &res, msg)

	case "height":
		utils.MakeFunction(functions.HeightFunction)(w, r, &res, msg)

	case "info":
		utils.MakeFunction(functions.InfoFunction)(w, r, &res, msg)

	case "list":
		out := []byte(`{"functions":["info","ground","height","dtm","dart","list","translate"]}`)

		if err := json.Unmarshal(out, &res.Response); err != nil {
			log.Fatal(err)
		}

	case "options":
		type AllOptions struct {
			Ground *functions.GroundOptions `json:"ground,omitempty"`
			Info   *functions.InfoOptions   `json:"info,omitempty"`
			Dtm    *functions.DtmOptions    `json:"dtm,omitempty"`
		}
		var a AllOptions
		a.Ground = functions.NewGroundOptions()
		a.Info = functions.NewInfoOptions()
		a.Dtm = functions.NewDtmOptions()
		bar, _ := json.Marshal(a)
		if err := json.Unmarshal(bar, &res.Response); err != nil {
			log.Fatal(err)
		}

	case "translate":
		utils.MakeFunction(functions.TranslateFunction)(w, r, &res, msg)

	// An unrecognized function will result in 400 error, with message explaining
	// how to list available functions.
	default:
		utils.BadRequest(w, r, res, "")
		return
	}

	// If we made it here, we can record the FinishedAt time, notify the job
	// manager of success, and return 200.
	res.FinishedAt = time.Now()
	utils.Okay(w, r, res, "Success!")
}
