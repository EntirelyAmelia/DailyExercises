package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	counts := []string{
		"900,google.com",
		"60,mail.yahoo.com",
		"10,mobile.sports.yahoo.com",
		"40,sports.yahoo.com",
		"300,yahoo.com",
		"10,stackoverflow.com",
		"2,en.wikipedia.org",
		"1,es.wikipedia.org",
	}
	countMap := getCounts(counts)

	fmt.Printf("Result: %+v", countMap)
}

func getCounts(counts []string) map[string]int64 {
	countMap := make(map[string]int64, 0)

	for _, link := range counts {
		s := strings.Split(link, ",")
		domains := getAllDomains(s[1])

		count, _ := strconv.ParseInt(s[0], 10, 64)

		for _, d := range domains {
			if _, ok := countMap[d]; ok {
				countMap[d] += count
			} else {
				countMap[d] = count
			}
		}

	}

	return countMap
}

func getAllDomains(domain string) []string {
	var domains []string

	domains = append(domains, domain)

	indexOfDot := strings.Index(domain, ".")
	if indexOfDot == -1 {
		return domains
	}

	subDomain := domain[indexOfDot+1:]
	domains2 := getAllDomains(subDomain)
	domains = append(domains, domains2...)

	return domains
}
