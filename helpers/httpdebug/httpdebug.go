package httpdebug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
)

// PrintResponse prints the http response to stdout
func PrintResponse(response *http.Response) {
	io.Copy(os.Stdout, response.Body)
}

func PrettyJson(reader io.ReadCloser) {
	var out bytes.Buffer
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	json.Indent(&out, input, "", "  ")
	out.WriteTo(os.Stdout)
}

func PrettyResponse(resp *http.Response) {

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseDump))
}
