package helper

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ParseID(r *http.Request, prefix string) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}

func DecodeJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
