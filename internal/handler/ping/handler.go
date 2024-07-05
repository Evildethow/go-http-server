package ping

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, "pong\n")
}
