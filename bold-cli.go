package main

/*
BOLD-CLI: a command line interface for data retrieval from http://www.boldsystems.org
*/

import (
	"errors"
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

// make sure that at least one of the required paramaters was passed in by the user
func PassedRequired(all_params map[string]string) error {
	required_params := []string{"ids", "bin", "container", "researchers", "geo", "marker", "taxon"}

	for _, p := range required_params {
		if all_params[p] != "none" {
			return nil
		}
	}

	err := errors.New("Not enough information. You must specify at least one of the following paramaters:\n" +
		"-taxon -ids -bin -container -researchers -geo -marker")

	return err
}

func main() {

	typePtr := flag.String("query", "combined", "BOLD query type: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("output", "bold_data.txt", "Output file name.")

	taxonPtr := flag.String("taxon", "none", "Taxonomic designation. Returns all records from matching designation.\n"+
		"Valid taxonomic designations: phylum, class, order, family, subfamily, genus, and species\n"+
		"Multiple taxa can be specified in a comma delimited list, or by passing a text file (with one taxon per line).\n")

	idsPtr := flag.String("ids", "none", "BOLD ID. Valid IDs include: Sample IDs, Process IDs, Museum IDs and Field IDs.\n"+
		"Multiple IDs can be specified in a comma delimited list, or by passing a text file (with one ID per line).\n")

	binPtr := flag.String("bin", "none", "Barcode index number. Returns all records in the BINs\n"+
		"Multiple markers can be specified in a comma delimited list, or by passing a text file (with one bin per line).\n")

	containerPtr := flag.String("container", "none", "Return all records from a given BOLD container. Containers include project codes and dataset codes\n"+
		"Multiple containers can be specified in a comma delimited list, or by passing a text file (with one container per line).\n")

	researchersPtr := flag.String("researchers", "none", "Return all records containing a matching researcher names\n"+
		"Multiple researchers can be specified in a comma delimited list, or by passing a text file (with one researcher per line).\n")

	geoPtr := flag.String("geo", "none", "Possible geographic inputs include countries and provinces/states.\n"+
		"Multiple researchers can be specified in a comma delimited list, or by passing a text file (with one location per line).\n")

	dataTypePtr := flag.String("dataType", "none", "Returns all records in one of the specified formats. Options are either overview or drill_down (default), which will respectively return:\n"+
		"drill_down: provides record counts by [BINs, Country, Storing Institution, Species]\n"+
		"overview: provides the total counts of [BINs, Countries, Storing Institutions, Orders, Families, Genus, Species] found by the query.\n")

	formatPtr := flag.String("format", "query_dependent", "The output file format. Different options are available for different query types listed below.\n"+ 
		"The first listed option is the default.\n"+
		"summary: json, xml\n"+
		"specimen: tsv, xml json, dwc\n"+
		"sequence: fasta\n"+
		"combined: tsv, xml json, dwc\n"+
		"trace: tar\n")

	markerPtr := flag.String("marker", "none", "Barcode marker: returns all specimen records that contain data for the specified barcode marker.\n"+
		"Options include (but are not limited to): COI-5P, matK, rbcL\n"+
		"Multiple markers can be specified in a comma delimited list.\n")

	ioPtr := flag.Bool("print", false, "If this flag is passed, instead of the data being sent to a file, the query will be returned to standard output.\n")

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
