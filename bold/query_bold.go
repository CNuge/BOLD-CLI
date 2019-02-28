package bold

import (
	"io"
	"log"
	"net/http"
	"os"
)

// change the input to a url built by the url_build functions...
// this way its general so we can do fasta or the other data types
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
