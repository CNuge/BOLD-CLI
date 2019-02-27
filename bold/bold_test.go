package bold

import (
	"errors"
	"reflect"
	"testing"
)

func TestUrlParamValidate(t *testing.T) {

	test_summary_good := "ids"
	test_summary_bad := "marker"

	test_sequence_good := "bin"
	test_sequence_bad := "format"

	t_sum_good_out := validateParam(test_summary_good, "summary")
	t_sum_bad_out := validateParam(test_summary_bad, "summary")
	t_sum_bad_expected := errors.New("Error! \"marker\" is not a valid paramater for BOLD query of type: summary")

	t_seq_good_out := validateParam(test_sequence_good, "sequence")
	t_seq_bad_out := validateParam(test_sequence_bad, "sequence")
	t_seq_bad_expected := errors.New("Error! \"format\" is not a valid paramater for BOLD query of type: sequence")

	if reflect.DeepEqual(t_sum_good_out, nil) != true {
		t.Errorf("URL param validation of combo: summary, %v incorrectly returned an error:\n %v ", test_summary_good, t_sum_good_out)
	}

	if reflect.DeepEqual(t_sum_bad_out, t_sum_bad_expected) != true {
		t.Errorf("Incorrect error message for combo: summary, marker\nobserved: %v\nexpected: %v", t_sum_bad_out, t_sum_bad_expected)
	}

	if reflect.DeepEqual(t_seq_good_out, nil) != true {
		t.Errorf("URL param validation of combo: sequence, %v incorrectly returned an error:\n %v ", test_sequence_good, t_seq_good_out)
	}

	if reflect.DeepEqual(t_seq_bad_out, t_seq_bad_expected) != true {
		t.Errorf("Incorrect error message for combo: sequence, format\nobserved: %v\nexpected: %v", t_seq_bad_out, t_seq_bad_expected)
	}
}

func TestBoldUrl(t *testing.T) {

	var URL1_components = map[string][]string{
		"taxon":        []string{"Chordata"},
		"geo":          []string{"Florida"},
		"institutions": []string{"Smithsonian Institution"},
	}

	expected_URL1 := "http://www.boldsystems.org/index.php/API_Public/sequence?geo=Florida&institutions=Smithsonian%20Institution&taxon=Chordata"
	built_URL1 := BoldURL("sequence", URL1_components)

	if reflect.DeepEqual(expected_URL1, built_URL1) != true {
		t.Errorf("URL 1 did not match expected!\nobserved: %v\nexpected: %v", built_URL1, expected_URL1)
	}

	var URL2_components = map[string][]string{
		"taxon":  []string{"Aves"},
		"geo":    []string{"Costa Rica"},
		"format": []string{"tsv"},
	}

	expected_URL2 := "http://www.boldsystems.org/index.php/API_Public/specimen?format=tsv&geo=Costa%20Rica&taxon=Aves"
	built_URL2 := BoldURL("specimen", URL2_components)

	if reflect.DeepEqual(expected_URL2, built_URL2) != true {
		t.Errorf("URL 2 did not match expected!\nobserved: %v\nexpected: %v", built_URL2, expected_URL2)
	}

	var URL3_components = map[string][]string{
		"taxon":        []string{"Aves", "Reptilia"},
		"geo":          []string{"Florida"},
		"institutions": []string{"Smithsonian Institution"},
	}

	expected_URL3 := "http://www.boldsystems.org/index.php/API_Public/sequence?geo=Florida&institutions=Smithsonian%20Institution&taxon=Aves|Reptilia"
	built_URL3 := BoldURL("sequence", URL3_components)

	if reflect.DeepEqual(expected_URL3, built_URL3) != true {
		t.Errorf("URL 3 did not match expected!\nobserved: %v\nexpected: %v", built_URL3, expected_URL3)
	}

}

// need a test for QueryToFile
