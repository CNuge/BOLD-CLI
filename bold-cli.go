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
	"errors"
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

// make sure that at least one of the required paramaters was passed in by the user
func PassedRequired(all_params map[string]string) error	{
	required_params := []string{"ids", "bin", "container", "researchers", "geo", "marker"}

	for _, p := range required_params{
		if all_params[p] != "none" {
			return nil
		}
	}

	err := errors.New("Not enough information. You must specify at least one of the following paramaters:\n"+
						"-ids -bin -container -researchers -geo -marker")

	return err
}

func main() {

	typePtr := flag.String("query", "combined", "BOLD query type: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("output", "bold_data.txt", "Output file name.")

	taxonPtr := flag.String("taxon", "none", "Taxonomic designation. Returns all records from matching designation."+
		"Valid taxonomic designations: phylum, class, order, family, subfamily, genus, and species"+
		"Multiple taxa can be specified in a comma delimited list, or by passing a text file (with one taxon per line)")

	idsPtr := flag.String("ids", "none", "BOLD ID. Valid IDs include: Sample IDs, Process IDs, Museum IDs and Field IDs."+
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

	// check to make sure at least one of the required params was passed
	err := PassedRequired(all_params)
	if err != nil {
		log.Fatal(err)
	}

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
