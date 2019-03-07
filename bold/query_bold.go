package bold

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Takes a valid BOLD URL (i.e. built by the BoldURL function) and filename
// The URL is used to make an http request. The response data from the request 
// is written to the specified output file.
func QueryToFile(url string, output string) error {

	//make the file
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	// make the http request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// write the url response data to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Takes a valid BOLD URL (i.e. built by the BoldURL function) and makes an http request.
// The response data is sent to standard output.
func QueryToIO(url string) {

	// make the http request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// io.Copy(dst io.Writer, src io.Reader), copies from the Body to Stdout
	io.Copy(os.Stdout, resp.Body)

}
