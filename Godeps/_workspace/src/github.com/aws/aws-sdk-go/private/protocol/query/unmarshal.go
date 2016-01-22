package query

//go:generate go run ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/output/query.json unmarshal_test.go

import (
	"encoding/xml"

	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/venicegeo/pzsvc-pdal/Godeps/_workspace/src/github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"
)

// Unmarshal unmarshals a response for an AWS Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		decoder := xml.NewDecoder(r.HTTPResponse.Body)
		err := xmlutil.UnmarshalXML(r.Data, decoder, r.Operation.Name+"Result")
		if err != nil {
			r.Error = awserr.New("SerializationError", "failed decoding Query response", err)
			return
		}
	}
}

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {
	r.RequestID = r.HTTPResponse.Header.Get("X-Amzn-Requestid")
}
