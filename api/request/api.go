package request

import (
	"io/ioutil"
	"net/http"

	json "github.com/json-iterator/go"

	"github.com/go-playground/form/v4"
	"github.com/pkg/errors"
)

func ParseHTTPRequest(v interface{}, r *http.Request) (interface{}, error) {
	// make sure the http request is not empty
	if r.ContentLength == 0 {
		return nil, errors.New("empty request")
	}

	// try to parse the http request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}

	if len(body) > 0 {
		if err = json.ConfigFastest.Unmarshal(body, &v); err != nil {
			return nil, errors.Wrap(err, "unmarshal body")
		}
	} else {
		// try to parse the http request form
		if err = r.ParseForm(); err != nil {
			return nil, errors.Wrap(err, "parse form")
		}
		if len(r.Form) > 0 {
			if err := form.NewDecoder().Decode(v, r.Form); err != nil {
				return nil, errors.Wrap(err, "decode form")
			}
		}
	}

	return v, nil
}
