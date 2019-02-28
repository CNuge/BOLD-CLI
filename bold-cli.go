package main

/*
BOLD-CLI: a command line interface for data retrieval from http://www.boldsystems.org
*/

import (
	"github.com/CNuge/BOLD-CLI/bold" // switch this to the github import
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

// read in a file with a list of paramater values, each value should
// be on a separate line.
func ReadValues(filename string) []string {
	//Opening a file
	file, err := ioutil.ReadFile(filename)
	// check if that caused an error
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")

	// remove leading and trailing strings if they exist
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}

	if data[0] == "" {
		data = data[1:]
	}
	return data
}

func main() {

	typePtr := flag.String("query", "combined", "Query type. One of: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("output", "bold_data.txt", "Output file name.")

	taxonPtr := flag.String("taxon", "none", "")

	binPtr := flag.String("bin", "none", "")

	containerPtr := flag.String("container", "none", "")

	researchersPtr := flag.String("researchers", "none", "")

	geoPtr := flag.String("geo", "none", "")

	dataTypePtr := flag.String("dataType", "none", "")

	formatPtr := flag.String("format", "none", "")

	markerPtr := flag.String("marker", "none", "")

	// parse the command line arguments
	flag.Parse()

	passed_params := make(map[string][]string)

	//all of the paramaters of the argument parser
	var all_params = map[string]string{"taxon": *taxonPtr,
		"bin":         *binPtr,
		"container":   *containerPtr,
		"researchers": *researchersPtr,
		"geo":         *geoPtr,
		"dataType":    *dataTypePtr,
		"format":      *formatPtr,
		"marker":      *markerPtr}

	// iterate through the passed params
	for k, v := range all_params {
		if v != "none" {
			if len(strings.Split(v, ".")) > 1 {
				// read the data in from a file
				passed_vals := ReadValues(v)
				passed_params[k] = passed_vals
			} else {
				passed_vals := strings.Split(v, ",")
				passed_params[k] = passed_vals
			}
		}
	}

	// build the query url
	url := bold.BoldURL(*typePtr, passed_params)

	// retrieve the data, write to file
	bold.QueryToFile(url, *outputPtr)

}
