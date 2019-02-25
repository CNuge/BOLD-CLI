package bold


import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)


// change the input to a url built by the url_build functions...
// this way its general so we can do fasta or the other data types
func QueryToFile(accession []string, output string) error {
	// construct the url
	query_url := buildURL(accession)

	//make the file
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	// make the http request
	resp, err := http.Get(query_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Write data direct to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}