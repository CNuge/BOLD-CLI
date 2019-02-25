package bold

import (
	"fmt"
	"strings"
)

// take a paramater specified and check that it is a valid paramater
// for the given 
func validateParam(param string, data_type string) (bool, error) {

	// check if the param is in the valid inputs for the specified data_type
	// if so, return true, nil
	// if not return an error saying that:
	//  'param' not supported for BOLD query of type: 'data_type'


}


// take the data type and a map of all the url component paramaters
// validates that the data type and the map paramaters are allowed 
// in combination for the bold data retrieval type
func BoldURL(data_type string, params map[string]string) string {

	base := "http://www.boldsystems.org/index.php/API_Public/"
	url_dtype := "not_specified"

	if data_type == "summary"{
		url_dtype = "stats?"
	}else if data_type == "specimen"{
		url_dtype = "specimen?"
	}else if data_type == "sequence"{
		url_dtype = "sequence?"
	}else if data_type == "combined"{
		url_dtype = "combined?"		
	}else if data_type == "trace"{
		url_dtype = "trace?"	
	}

	url_params := []string

	for k , v := range params {

		_, err:= validateParams(k, data_type)
		if err != nil{
			log.Fatal(err)
		}

		param_str := fmt.Sprintf("%v=%v", k , v) 

		url_params = append(url_params, param_str)
	}
	// join all the params together
	strings.Join(url_params, "&")
	// patch the three components together
	url := []string{base, url_dtype, url_params}
	// return the url string
	return strings.Join(url, "")

}