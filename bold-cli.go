package main
/*
BOLD-CLI: a command line interface for data retrieval from http://www.boldsystems.org
*/

import (
	"bold" // switch this to the github import
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
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

func main(){

	typePtr := flag.String("q", "__none__", "Query type. One of: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("o", "__none__", "Output file name.")

	taxonPtr := flag.String("taxon", "__none__", "")

	binPtr := flag.String("bin", "__none__", "")

	containerPtr := flag.String("container", "__none__", "")

	researchersPtr := flag.String("researchers", "__none__", "")

	geoPtr := flag.String("geo", "__none__", "")

	dataTypePtr := flag.String("dataType", "__none__", "")

	formatPtr := flag.String("format", "__none__", "")

	markerPtr := flag.String("marker", "__none__", "")


	// parse the command line arguments
	flag.Parse()	

	//parse the flags for multiple arguments
	if *typePtr == "__none__" {
		err := "You must specify the BOLD query type. Options: summary, specimen, sequence, combined, trace"
		log.Fatal(err)
	}

	if *outputPtr == "__none__" {
		err := "You must specify an output file name."
		log.Fatal(err)
	}

	passed_params := make(map[string][]string)

	//all of the paramaters of the argument parser
	var all_params = map[string]string{"taxon" : *taxonPtr, 
										"bin" : *binPtr, 
										"container" : *containerPtr, 
										"researchers": *researchersPtr, 
										"geo" : *geoPtr, 
										"dataType": *dataTypePtr, 
										"format" : *formatPtr, 
										"marker" : *markerPtr}

	// iterate through the passed params 
	for k, v := range all_params {
		if v != "__none__" {
			if len(strings.Split(v, ".")) > 1 {
				// read the data in from a file
				passed_vals := ReadValues(v)
				passed_params[k] = 
			} else{
				passed_vals := strings.Split(v, ",")
				passed_params[k] = passed_vals				
			}
		}
	}

	// build the query url
	url := bold.BoldURL(*typePtr, )

	// retrieve the data, write to file
	bold.QueryToFile(url, *outputPtr)

}





