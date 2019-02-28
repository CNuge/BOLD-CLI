package main

/*
BOLD-CLI: a command line interface for data retrieval from http://www.boldsystems.org
*/

import (
	"flag"
	"github.com/CNuge/BOLD-CLI/bold" // "./bold"
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

	// remove leading and trailing spaces if they exist
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}

	if data[0] == "" {
		data = data[1:]
	}

	return data
}

func main() {

	typePtr := flag.String("query", "combined", "BOLD query type: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("output", "bold_data.txt", "Output file name.")

	taxonPtr := flag.String("taxon", "none", "")

	idsPtr := flag.String("taxon", "none", "BOLD ID. Valid IDs are of types: Sample IDs, Process IDs, Museum IDs and Field IDs"+
			"Multiple IDs can be specified in a comma delimited list, or by passing a text file (with one ID per line)")

	binPtr := flag.String("bin", "none", "Barcode index number. Returns all records in the BINs"+
			"Multiple markers can be specified in a comma delimited list, or by passing a text file (with one bin per line)")

	containerPtr := flag.String("container", "none", "")

	researchersPtr := flag.String("researchers", "none", "")

	geoPtr := flag.String("geo", "none", "")

	dataTypePtr := flag.String("dataType", "none", "")

	formatPtr := flag.String("format", "query_dependent", "The output file format. Different options available for different query types listed below. First listed option is the default.\n"+
		"summary: json, xml\n"+
		"specimen: tsv, xml json, dwc\n"+
		"sequence: fasta\n"+
		"combined: tsv, xml json, dwc\n"+
		"trace: tar\n")

	markerPtr := flag.String("marker", "none", "Barcode marker: returns all specimen records that contain data for the specified barcode marker.\n"+
		"Options include (but are not limited to): COI-5P, matK, rbcL"+
		"Multiple markers can be specified in a comma delimited list.")

	ioPtr := flag.Bool("print", false, "If this flag is passed, instead of data being output to a file, the query will be returned to standard output.")

	// parse the command line arguments
	flag.Parse()

	passed_params := make(map[string][]string)

	if *formatPtr == "query_dependent" {
		if *typePtr == "summary" {
			*formatPtr = "json"
		} else if *typePtr == "specimen" {
			*formatPtr = "tsv"
		} else if *typePtr == "sequence" {
			*formatPtr = "none"
		} else if *typePtr == "combined" {
			*formatPtr = "tsv"
		} else if *typePtr == "trace" {
			*formatPtr = "none"
		}
	}

	//all of the paramaters of the argument parser
	var all_params = map[string]string{"taxon": *taxonPtr,
		"ids":         *idsPtr,
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

	if *ioPtr == true {
		//pipe to stdout if the -print flag was passed
		bold.QueryToIO(url)

	} else {
		// retrieve the data, write to file
		bold.QueryToFile(url, *outputPtr)
	}

}
