package bold

//Description of package
//

import (
	"fmt"
	"strings"
	"log"
	"errors"
)

var valid_dict = map[string][]string{
	"summary":  []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "dataType", "format"},
	"specimen": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format"},
	"sequence": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo"},
	"combined": []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format", "marker"},
	"trace":    []string{"taxon", "ids", "bin", "container", "institutions", "researchers", "geo", "format", "marker"},
}

// take a paramater specified and check that it is a valid paramater
// for the given data type
func validateParam(param string, data_type string) error {

	for _, i := range valid_dict[data_type] {
		if i == param {
			return nil
		}
	}

	err := fmt.Sprintf("Error! \"%v\" is not a valid paramater for BOLD query of type: %v", param, data_type)

	return errors.New(err)
}

func urlString(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

// Take the data type and a map of all the url component paramaters
// validates that the data type and the map paramaters are allowed
// in combination for the bold data retrieval type
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

	if url_dtype == "not_specified"{
		err := "You must specify the BOLD query type. Options: summary, specimen, sequence, combined, trace"
		log.Fatal(err)
	}

	url_params := []string{}

	// iterate through the alloted params, make sure they are valid,
	// if so then build the components of the url
	for k, v := range params {

		err := validateParam(k, data_type)
		if err != nil {
			log.Fatal(err)
		}

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
