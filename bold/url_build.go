// The bold package provides a set of functions for building a URL to interating with the BOLD API (http://www.boldsystems.org/index.php/resources/api?type=webservices),
// querying BOLD using the constructed URL and sending the output of the query to a file or to standard output.
package bold

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
)

var valid_dict = map[string][]string{
	"summary":  []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "dataType", "format"},
	"specimen": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format"},
	"sequence": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "marker"},
	"combined": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format", "marker"},
	"trace":    []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format", "marker"},
}

// Takes a paramater specified in the paramater map and checks that it is a valid paramater
// for the given data type of the BOLD query.
func validateParam(param string, data_type string) error {

	for _, i := range valid_dict[data_type] {
		if i == param {
			return nil
		}
	}

	err := fmt.Sprintf("Error! \"%v\" is not a valid paramater for BOLD query of type: %v\n"+
		"This flag should be omitted.", param, data_type)

	return errors.New(err)
}

// Takes an input string and replaces all spaces with %20
func urlString(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

// Takes the data type and a map of all the url component paramaters.
// Returns a URL string that can be used to query BOLD.
// Sorts the paramaters alphabetically then validates that all map 
// paramaters are allowed for the data type. If valid then any spacesin
// the paramaters are filled, multiple values for a paramater are combined 
// into a single string and finally the component of the URL is constructed.
func BoldURL(data_type string, params map[string][]string) string {

	base := "http://www.boldsystems.org/index.php/API_Public/"
	url_dtype := "not_specified"

	if data_type == "summary" {
		url_dtype = "stats?"
	} else if data_type == "specimen" {
		url_dtype = "specimen?"
	} else if data_type == "sequence" {
		url_dtype = "sequence?"
	} else if data_type == "combined" {
		url_dtype = "combined?"
	} else if data_type == "trace" {
		url_dtype = "trace?"
	}

	if url_dtype == "not_specified" {
		err := errors.New("You must specify the BOLD query type. Options: summary, specimen, sequence, combined, trace")
		log.Fatal(err)
	}

	url_params := []string{}

	// iterate through the alloted params, make sure they are valid,
	// if so then build the components of the url

	// get a slice of the params, sort in alphabetical order
	sorted_k := []string{}
	for k, _ := range params {
		sorted_k = append(sorted_k, k)
	}
	sort.Strings(sorted_k)

	// iterate through the sorted keys, validate them and call the corresponding
	// values prior to constructing the part of the url
	for _, k := range sorted_k {

		err := validateParam(k, data_type)
		if err != nil {
			log.Fatal(err)
		}

		// call the value
		v := params[k]

		// if multiple values passed for the param, join them together with a "|"
		joined_vals := strings.Join(v, "|")

		// here sub any spaces in v with %20 using urlString
		param_str := fmt.Sprintf("%v=%v", k, urlString(joined_vals))

		url_params = append(url_params, param_str)
	}
	// join all the params together
	joined_params := strings.Join(url_params, "&")
	// patch the three components together
	url := []string{base, url_dtype, joined_params}
	// return the url string
	return strings.Join(url, "")

}
