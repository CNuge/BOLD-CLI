package main



import (
	"bold" // switch this to the github import
)




func main(){

	typePtr := flag.String("q", "__none__", "Query type. One of: summary, specimen, sequence, combined, trace")

	outputPtr := flag.String("o", "__none__", "Output file name, specify the file extension here as well")

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

	/*
	These are the other accepted args:

	// Take in a comma delimited list of arguments for each or a single filename
	// that directs to a \n delimited list of values for the param

	"taxon", 
	"ids", 
	"bin", 
	"container", 
	"institutions", 
	"researchers", 
	"geo", 
	"dataType", 
	"format"

	need cli flags for each of them that can accept a string, file or pipe sep pair of values.


	*/
}





