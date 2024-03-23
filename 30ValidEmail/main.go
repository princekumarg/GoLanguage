package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkEmail(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error could not read from input", err)
	}
}
func checkEmail(email string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfrecord, dmarcRecord string

	mxRecords, err := net.LookupMX(email)
	if err != nil {
		log.Printf("Error %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(email)
	if err != nil {
		log.Printf("Error %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfrecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + email)
	if err != nil {
		log.Printf("Error %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v", email, hasMX, hasSPF, spfrecord, hasDMARC, dmarcRecord)
}
