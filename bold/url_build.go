package bold

import (
	"fmt"
	"strings"
)


func


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

	




}