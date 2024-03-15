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
	fmt.Printf("domain, hasMX,hasSPF, sprRecord, hasDMARC, dmrcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}

	// // .env 파일에서 환경 변수를 가져옵니다.
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")

	// // 가져온 환경 변수를 출력합니다.
	// fmt.Println("DB_USERNAME:", username)
	// fmt.Println("DB_PASSWORD:", password)
}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmrcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmrcRecords, err := net.LookupTXT("_dmrc." + domain)
	if err != nil {
		log.Printf("Error%v", err)
	}

	for _, record := range dmrcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmrcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmrcRecord)

}
