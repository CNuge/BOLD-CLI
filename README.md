# BOLD-CLI
## A command line interface for data retrieval from the barcode of life database
## http://www.boldsystems.org
[![Build Status](https://travis-ci.org/CNuge/BOLD-CLI.svg?branch=master)](https://travis-ci.org/CNuge/BOLD-CLI)	[![GoDoc](https://godoc.org/github.com/CNuge/BOLD-CLI/bold?status.svg)](https://godoc.org/github.com/CNuge/BOLD-CLI/bold)

## About

BOLD-CLI lets you download specimen and sequence data from the barcode of life database (BOLD) directly from the command line. BOLD-CLI interfaces with the database via the [BOLD API](http://www.boldsystems.org/index.php/resources/api?type=webservices) and allows you to obtain local copies of data without needing to utilize [the public data portal](http://www.boldsystems.org/index.php/Public_BINSearch?searchtype=records). The command line interface provides an efficient means of querying the database and also allows returned data to be directly piped to other UNIX shell commands. 

## Installation

To use BOLD-CLI, download the executable file in this repository for your operating system (Windows, Mac and Linux provided). Make sure the executable is located on your [PATH](https://en.wikipedia.org/wiki/PATH_(variable)) or within your working directory and then you should be ready to roll. Test that it is working by running the command: `bold-cli -h` from the command line/command prompt.

Or if you have go installed on your computer and your GOPATH is properly configured, you can install the program directly from the command line:
```
go get github.com/CNuge/BOLD-CLI
go install github.com/CNuge/BOLD-CLI
```
## Command line arguments
BOLD-CLI uses the following command line flags to control which data is retrieved:

- `-query` [The Bold Query type](http://www.boldsystems.org/index.php/resources/api?type=webservices). Options are: summary, specimen, sequence, combined, trace. If no option specified the default is to return the combined (specimen and sequence) format.  
- `-output` The file name that the output data should be saved to.
- `-format` The flag gives the output file format. Different options are available for different query types, and the first listed option is the default for that type. 
	- summary: json, xml
	- specimen: tsv, xml json, dwc
	- sequence: fasta
	- combined: tsv, xml json, dwc
	- trace: tar
- `-bin` [Barcode index number](http://www.boldsystems.org/index.php/Public_BarcodeIndexNumber_Home). Returns all records in the BIN.
- `-taxon` Returns all records from matching designation. Valid taxonomic designations: phylum, class, order, family, subfamily, genus, and species.
- `-marker` Returns all specimen records that contain data for the specified barcode marker.
- `-ids` BOLD IDs. Valid IDs include: Sample IDs, Process IDs, Museum IDs and Field IDs.
- `-container` Return all records from a given BOLD container. Containers include project codes and dataset codes.
- `-researcher` Return all records containing a matching researcher name.
- `-geo` Geographic location. Possible geographic inputs include countries and provinces/states.
- `-dataType` Returns all records in one of the specified formats. Options are either overview or drill_down (default), which will respectively return:
		- drill_down: provides record counts by [BINs, Country, Storing Institution, Species]
		- overview: provides the total counts of [BINs, Countries, Storing Institutions, Orders, Families, Genus, Species] found by the query.

- `-print` Return data to standard output instead of to an output file.

For most flags, multiple values can be specified in a comma delimited list or by passing the flag a text file containing the values (one value per line).

## Examples
```
bold-cli -output salp_barcodes.fasta -query sequence -taxon salvelinus alpinus

bold-cli -query specimen -output test.tsv -taxon Aves -geo Florida -format tsv
```
- Default output is the combined summary and sequence data in tsv format.
```
bold-cli -taxon Aves -geo Florida
```
- Can use multiple arguments for one paramater from command line, just comma delimit them.
```
bold-cli -query specimen -output test2.tsv -taxon Aves,Reptilia -geo Florida -format tsv
```
- Or use multiple arguments for one paramater by passing in a text file, with each option listed on a separate line.
```
bold-cli -query sequence -output test3.fasta -taxon ./example_data/taxon_test.txt -geo Ontario -marker COI-5P
```
- Send results to standard output as opposed to files. Can then be piped to other things. Note if you have paramaters with spaces in them (i.e. drosophila melanogaster) put this last in the command query, or use a text file format to input the paramaters.
```
bold-cli -query specimen -print -taxon drosophila melanogaster  
```
- The `-print` flag can be added to a query to send data to standard output instead of to a file. Below is an example use of pipe functionality that counts the number of COI drosophila sequences in the bold database.
```
bold-cli -query sequence -taxon drosophila -marker COI-5P -print | grep -c "^>"
```

