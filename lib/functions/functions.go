package functions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/gouthamve/dredd"
	"github.com/gouthamve/dredd/judge"
	"github.com/juju/errors"
	"github.com/spf13/viper"
)

// ExecuteSubmission executes the submission against the test-cases and returns
// the output
// TODO: Fix the fucked up API
// TODO: Fix the hardcoded URLs
func ExecuteSubmission(ra judge.RunnerArgs) ([]dredd.Result, error) {
	endpoint := viper.GetString("functions-endpoint")
	epURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Annotatef(err, "IronFunctions Endpoint is baaad: %s", endpoint)
	}

	epURL.Path = path.Join(epURL.Path, "/r/dredd/judge/go")
	byt := new(bytes.Buffer)

	if err := json.NewEncoder(byt).Encode(ra); err != nil {
		return nil, errors.Annotate(err, "Cannot marshal args")
	}

	resp, err := http.Post(epURL.String(), "application/json", byt)
	if err != nil {
		return nil, errors.Annotate(err, "Request to Functions failed")
	}
	defer resp.Body.Close()

	results := make([]dredd.Result, 0)
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, errors.Annotate(err, "Invalid Response")
	}

	return results, nil
}
