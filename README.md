# BOLD-CLI
## ** UNDER CONSTRUCTION **
## A command line interface for data retrieval from the barcode of life database
## http://www.boldsystems.org
[![Build Status](https://travis-ci.org/CNuge/BOLD-CLI.svg?branch=master)](https://travis-ci.org/CNuge/BOLD-CLI)	[![GoDoc](https://godoc.org/github.com/CNuge/BOLD-CLI/bold?status.svg)](https://godoc.org/github.com/CNuge/BOLD-CLI/bold)

Uses the BOLD API as a starting point:

http://www.boldsystems.org/index.php/resources/api?type=webservices

This program lets you download sequence and summary data from the barcode of life database, directly from the command line.

## Installation

To use BOLD-CLI, download the executable file *COMING SOON* found in this repository corresponding to your operating system (Windows, Mac and Linux versions provided). Make sure the executable is located on your [PATH](https://en.wikipedia.org/wiki/PATH_(variable) or within your working directory and then you should be ready to roll. Test its working by running `bold-cli -h` from the command line/command prompt.

Or if you have go installed on your computer and your GOPATH is properly configured, you can install the program directly from the command line:
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
- add documentation for the command line flags
- write the readme file and add examples
- add godocs detail

bugs to fix:
- if not params are passed, then the call of the program with no flags builds
an empty file.
	- have it check the passed params... if len(not_file_or_query) < 1
	then throw an error saying you have to specify the data to get from bold. then list the options.

- Compiling on windows and linux as well, as last step upload a compiled version of each, zipped and labelled.
	- a version compiled on windows will work without go installed... so compile and ship
	- or possibly do this: https://golang.org/doc/install/source#environment