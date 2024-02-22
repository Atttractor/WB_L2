package render

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func RenderResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	b := bytes.Buffer{}
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(b.Bytes()); err != nil {
		return
	}
}

func ErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error, details string) {
	RenderResponse(w, r, httpStatusCode, map[string]interface{}{"error": err.Error(), "details": details})
}

func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
