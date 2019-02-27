# BOLD-CLI
## ** UNDER CONSTRUCTION **
## A command line interface for data retrieval of from the barcode of life database
## http://www.boldsystems.org


Uses the BOLD API as a starting point:

http://www.boldsystems.org/index.php/resources/api?type=webservices

This program lets you download sequence and summary data from the barcode of life database, directly from the command line.

## Installation
via go:
```
go get github.com/CNuge/BOLD-CLI
go install github.com/CNuge/BOLD-CLI
```


example usage:
./bold-cli -q specimen -o test2.tsv -taxon ./example_data/taxon_test.txt -geo Florida -format tsv


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
- add godocs
- add travis.yml file