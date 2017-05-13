package debug

import (
	"io"
	"net/http"
	"os"
)

// PrintResponse prints the http response to stdout
func PrintResponse(response http.Response) {
	io.Copy(os.Stdout, response.Body)
}
