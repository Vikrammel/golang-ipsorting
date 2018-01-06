# Sorting <IP:Port>s in Golang

### by Vikram Melkote

## Setup

* git clone into `~/go/src` or download `.go` files and place in `~/go/src/ipsorting`:

	`git clone https://github.com/Vikrammel/golang-ipsorting.git`

* if vendoring, place `.go` files in `<projectroot>/app/vendor/ipsorting`

* import package in code:
	
	```go
	import (
		"ipsorting"
	)
	```

## Usage in code

use on string or an array/slice of strings of ip:ports (ipsorting.go can be easily modified to work with just IPs instead of IP:Port)
	
* compare two IP:Ports
	
	```go
	sortedPair := ipsorting.OrderIPPair("10.0.0.14:8080", "10.0.0.14:5060")
	//sortedPair is now a string array literal of size 2 of the two ip:port strings passed in
	//in ascending numerical order
	log.Println(sortedPair[0])
	//10.0.0.14:5060
	```

* sort a slice of IP:Ports
	
	```go
	//unsortedList = ["10.0.0.24:8080", "10.0.0.21:8080", "10.0.0.22:8080", "10.0.0.23:8080"]
	sortedList := ipsorting.SortIPs(unsortedList)
	log.Println(sortedList)
	//["10.0.0.21:8080", "10.0.0.22:8080", "10.0.0.23:8080", "10.0.0.24:8080"]
	//ports do not need to be the same, they will also be compared
	//adapted quickSort, avg complexity = O(nlog(n)), worst case = O(n^2)
	```

* insert IP:Port into slice in correct position
	
	```go
	//sortedList = ["10.0.0.21:8080", "10.0.0.22:8080", "10.0.0.23:8080", "10.0.0.24:8080"]
	sortedList = ipsorting.InsortIP(sortedList, "10.0.0.23:8070")
	log.Println(sortedList)
	//["10.0.0.21:8080", "10.0.0.22:8080", "10.0.0.23:8070", "10.0.0.23:8080", "10.0.0.24:8080"]
	//adapted binarySearch, complexity = O(log(n))
	```

## Reference

[Quicksort in python](https://inst.eecs.berkeley.edu/~cs188/sp09/projects/tutorial/docs/quickSort.html)

[Binary Search [for InsortIP()]](https://rosettacode.org/wiki/Binary_search)


Feel free to use this in any of your projects and please let me know if you encounter any bugs. Thank you.
