package bold


import (
	"fmt"
	"reflect"
	"testing"
)


func TestUrlParamValidate (t * testing.T) {

	test_summary_good := "ids"
	test_summary_bad := "marker"

	test_sequence_good := "bin"	
	test_sequence_bad := "format"


	t_sum_good_out := validateParam("summary", test_summary_good)
	t_sum_bad_out := validateParam("summary", test_summary_bad)
	t_sum_bad_expected := "Error! \"marker\" is not a valid paramater for BOLD query of type: summary"

	t_seq_good_out := validateParam("sequence", test_sequence_good)
	t_seq_bad_out := validateParam("sequence", test_sequence_bad)
	t_seq_bad_expected := "Error! \"format\" is not a valid paramater for BOLD query of type: sequence"

	if reflect.DeepEqual(t_sum_good_out, nil)!= true {
		t.Errorf("URL param validation of combo: summary, %v incorrectly returned an error:\n %v ", test_summary_good , t_sum_good_out)
	}
	
	if reflect.DeepEqual(t_sum_bad_out, t_sum_bad_expected)!= true {
		t.Errorf("Incorrect error message for combo: summary, marker\nobserved: %v\nexpected: %v",t_sum_bad_out , t_sum_bad_expected)
	}
	
	if reflect.DeepEqual(t_seq_good_out, nil) != true {
		t.Errorf("URL param validation of combo: sequence, %v incorrectly returned an error:\n %v ", test_sequence_good , t_seq_good_out)
	}
	
	if reflect.DeepEqual(t_seq_bad_out, t_seq_bad_expected)!= true {
		t.Errorf("Incorrect error message for combo: sequence, format\nobserved: %v\nexpected: %v",t_seq_bad_out , t_seq_bad_expected)
	}
}


func TestBoldUrl (t * testing.T) {

	URL1_components := make(map[string]string){
		"taxon" : "Chordata",
		"geo" : "Florida",
		"institutions" : "Smithsonian Institution",
	}

	expected_URL1 := "http://www.boldsystems.org/index.php/API_Public/sequence?taxon=Chordata&geo=Florida&institutions=Smithsonian%20Institution"
	built_URL1 := BoldURL("sequence", URL1_components)

	if reflect.DeepEqual(expected_URL1, built_URL1) != true {
		t.Errorf("URL 1 did not match expected!\nobserved: %v\nexpected: %v" , expected_URL1, built_URL1)
	}


	URL2_components := make(map[string]string){
		"taxon" : "Aves",
		"geo" : "Costa Rica",
		"format" : "tsv",
	}

	expected_URL2 := "http://www.boldsystems.org/index.php/API_Public/specimen?taxon=Aves&geo=Costa%20Rica&format=tsv"
	built_URL2 := BoldURL("specimen", URL2_components)

	if reflect.DeepEqual(expected_URL2, built_URL2) != true {
		t.Errorf("URL 2 did not match expected!\nobserved: %v\nexpected: %v" , expected_URL2, built_URL2)
	}

}