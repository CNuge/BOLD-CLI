# BOLD-CLI
## ** UNDER CONSTRUCTION **
## A command line interface for data retrieval from the barcode of life database
## http://www.boldsystems.org
[![Build Status](https://travis-ci.org/CNuge/BOLD-CLI.svg?branch=master)](https://travis-ci.org/CNuge/BOLD-CLI)	[![GoDoc](https://godoc.org/github.com/CNuge/BOLD-CLI/bold?status.svg)](https://godoc.org/github.com/CNuge/BOLD-CLI/bold)

Uses the BOLD API as a starting point:

http://www.boldsystems.org/index.php/resources/api?type=webservices

This program lets you download sequence and summary data from the barcode of life database, directly from the command line.

## Installation

Install via go:
```
go get github.com/CNuge/BOLD-CLI
go install github.com/CNuge/BOLD-CLI
```


example usage:
- using command line arguments
```
bold-cli -q specimen -o test.tsv -taxon Aves -geo Florida -format tsv
```
- using multiple arguments for one paramater from command line
```
bold-cli -q specimen -o test2.tsv -taxon Aves,Reptilia -geo Florida -format tsv
```
- using multiple arguments for one paramater via a text file
```
bold-cli -q specimen -o test3.tsv -taxon ./example_data/taxon_test.txt -geo Florida -format tsv
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
