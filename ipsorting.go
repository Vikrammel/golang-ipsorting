package ipsorting

import (
	"math/rand"
	"strings"
)

// OrderIPPair returns an array of two IP:PORTs in ascending order
func OrderIPPair(firstIP string, secondIP string) [2]string {
	//extract numbers out of IP strings
	firstIPPortArr := strings.Split(firstIP, ":")
	firstIPArr := strings.Split(firstIPPortArr[0], ".")
	secondIPPortArr := strings.Split(secondIP, ":")
	secondIPArr := strings.Split(secondIPPortArr[0], ".")

	//compare and return ordered string literal array
	for index, num := range firstIPArr {
		if num == secondIPArr[index] {
			continue
		} else if num > secondIPArr[index] {
			return [2]string{secondIP, firstIP}
		} else {
			return [2]string{firstIP, secondIP}
		}
	} //now check ports if nums were the same and return ordered IP:Ports
	if firstIPPortArr[1] < secondIPPortArr[1] {
		return [2]string{firstIP, secondIP}
	}
	return [2]string{secondIP, firstIP}
}

// SortIPs sorts IP addresses of an array in asc. order (quicksort)
func SortIPs(addrs []string) []string {
	//recursion base case
	if len(addrs) < 2 {
		return addrs
	}
	pivot := addrs[rand.Intn(len(addrs))] //random ip in addrs as pivot

	var left []string   //for IPs<pivot
	var middle []string //for IPs=pivot
	var right []string  //for IPs>pivot

	for _, ip := range addrs {
		if OrderIPPair(ip, pivot)[0] == ip {
			left = append(left, ip)
		} else if ip == pivot {
			middle = append(middle, ip)
		} else {
			right = append(right, ip)
		}
	}

	//combine and return
	left, right = SortIPs(left), SortIPs(right)
	sortedIPs := append(left, middle...)
	sortedIPs = append(sortedIPs, right...)
	return sortedIPs
}

// InsortIP inserts IP into an array of IPs (using binary search)
func InsortIP(sortedIPs []string, newIP string) []string {
	first := 0 //initially first index, finally becomes the index to insert
	last := len(sortedIPs) - 1

	//while the list is still being searched for the insert point
	for first <= last {
		middle := (first + last) / 2
		newMiddleOrdered := OrderIPPair(newIP, sortedIPs[middle])
		if (newMiddleOrdered[0] == newIP) || (newMiddleOrdered[0] == newMiddleOrdered[1]) {
			//search on left half
			last = middle - 1
		} else {
			//search on right half
			first = middle + 1
		}
	}

	//insert into place and return new slice with newIP
	var newIPSlice[]string
	newIPSlice = append(newIPSlice, sortedIPs[:first]...)
	newIPSlice = append(newIPSlice, newIP)
	newIPSlice = append(newIPSlice, sortedIPs[first:]...)

	return newIPSlice
}
