// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package helpers

import (
	"encoding/json"
	"net/http"
)

// SendSuccessHeader is a generic method for setting and sending headers to
// our front end
func SendSuccessHeader(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// SendErrorHeader is a generic method for setting and sending errors to
// the front end
func SendErrorHeader(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
