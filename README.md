# BOLD-CLI
## ** UNDER CONSTRUCTION **
## A command line interface for data retrieval from the barcode of life database
## http://www.boldsystems.org
[![Build Status](https://travis-ci.org/CNuge/BOLD-CLI.svg?branch=master)](https://travis-ci.org/CNuge/BOLD-CLI)	[![GoDoc](https://godoc.org/github.com/CNuge/BOLD-CLI/bold?status.svg)](https://godoc.org/github.com/CNuge/BOLD-CLI/bold)

Uses the BOLD API as a starting point:

http://www.boldsystems.org/index.php/resources/api?type=webservices

This program lets you download sequence and summary data from the barcode of life database, directly from the command line.

## Installation

// add compiled versions to repo, explain usage

If you have go installed on your computer as your GOPATH configured you can install from the command line:
```
go get github.com/CNuge/BOLD-CLI
go install github.com/CNuge/BOLD-CLI
```


## Example usage:
```
bold-cli -o salp_barcodes.fasta -query sequence -taxon salvelinus alpinus

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
bold-cli -query sequence -output test3.fasta -taxon ./example_data/taxon_test.txt -geo Florida -format tsv
```
- Send results to standard output as opposed to files. Can then be piped to other things.
```
bold-cli -query specimen -taxon drosophila melanogaster  -print
```
-Example use of pip functionality: Count number of drosophila sequences in the bold database
```
bold-cli -query sequence -taxon drosophila -print | grep -c "^>"
```

TODO:
- need to catch erroneous combinations of paramaters on the input to aid the user.
	i.e.
	this works:
	./bold-cli -q specimen -o test.tsv -taxon Aves -geo Florida -format tsv
	but this doesn't:
	./bold-cli -q summary -o test.tsv -taxon Aves -geo Florida -format tsv

	not because the program is wrong... but because this isn't a valid data format
	for the summary option.
- add documentation for the command line flags
- write the readme file and add examples
- add godocs detail
