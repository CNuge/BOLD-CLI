# goBOLDly
## possible good names: BOLD-query
## command line tool for interfacing with the BOLD data API

Using this as a starting point:

http://www.boldsystems.org/index.php/resources/api?type=webservices

Write a go program that can build the requsite URL for a cmd line input of a taxonomic designation
or an ID or BIN and then write the information to a fasta file.
- this is basically an extension of go-Fasta... so use that as a starting point for designing
the input/output, download and command line interfaces.
	-take either a list of names/ids as input, or read them from a file like with go-fasta
- include the necessary unit tests, document it and then put it up on github in full.
- extend the functionality to download the other data types as well
	- summary
	- specimen 
- have ability to name the output file 
-interface with the fasta package of go-fasta in order to add summary/

- with tested verion, add a travis.yml file
- add a godoc page, make sure the functions are well documented.
- set it up locally via go get/go install.
- use it to download sequences for all of the different kingdoms, these can be utilized as 
inputs for DAPR transition matrix population.
 