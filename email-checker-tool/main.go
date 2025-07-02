// // package main

// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"log"
// // 	"net"
// // 	"os"
// // 	"strings"
// // )

// // func main() {

// // 	scanner := bufio.NewScanner(os.Stdin)
// // 	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

// // 	for scanner.Scan() {
// // 		checkDomain(scanner.Text())
// // 	}

// // 	if err := scanner.Err(); err != nil {
// // 		log.Fatal("Error: could not read from input: %v\n", err)
// // 	}
// // }

// // func checkDomain(domain string) {

// // 	var hasMX, hasSPF, hasDMARC bool
// // 	var spfRecord, dmarcRecord string

// // 	mxRecords, err := net.LookupMX(domain)

// // 	if err != nil {
// // 		log.Printf("Error: %v\n", err)
// // 	}

// // 	if len(mxRecords) > 0 {
// // 		hasMX = true
// // 	}

// // 	txtRecords, err := net.LookupTXT(domain)

// // 	if err != nil {
// // 		log.Printf("Error:%v\n", err)
// // 	}

// // 	for _, record := range txtRecords {
// // 		if strings.HasPrefix(record, "v=spf1") {
// // 			hasSPF = true
// // 			spfRecord = record
// // 			break
// // 		}
// // 	}

// // 	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
// // 	if err != nil {
// // 		log.Printf("Error:%v\n", err)
// // 	}

// // 	for _, record := range dmarcRecords {
// // 		if strings.HasPrefix(record, "v=DMARC1") {
// // 			hasDMARC = true
// // 			dmarcRecord = record
// // 			break
// // 		}
// // 	}

// // 	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
// // }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord\n")

// 	for scanner.Scan() {
// 		email := scanner.Text()
// 		parts := strings.Split(email, "@")
// 		if len(parts) != 2 {
// 			log.Printf("Invalid email: %s\n", email)
// 			continue
// 		}
// 		domain := parts[1]
// 		checkDomain(domain)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error: could not read from input: %v\n", err)
// 	}
// }

// func checkDomain(domain string) {
// 	var hasMX, hasSPF, hasDMARC bool
// 	var spfRecord, dmarcRecord string

// 	mxRecords, err := net.LookupMX(domain)
// 	if err != nil {
// 		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
// 	} else if len(mxRecords) > 0 {
// 		hasMX = true
// 	}

// 	txtRecords, err := net.LookupTXT(domain)
// 	if err != nil {
// 		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
// 	}
// 	for _, record := range txtRecords {
// 		if strings.HasPrefix(record, "v=spf1") {
// 			hasSPF = true
// 			spfRecord = record
// 			break
// 		}
// 	}

// 	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
// 	if err != nil {
// 		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
// 	}
// 	for _, record := range dmarcRecords {
// 		if strings.HasPrefix(record, "v=DMARC1") {
// 			hasDMARC = true
// 			dmarcRecord = record
// 			break
// 		}
// 	}

// 	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
// }

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
		email := scanner.Text()
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			log.Printf("❌ Invalid email format: %s\n", email)
			continue
		}
		domain := parts[1]
		fmt.Printf("\n📬 Checking domain: %s\n", domain)
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
}

func checkDomain(domain string) {
	hasMX, hasSPF, hasDMARC := false, false, false
	var spfRecord, dmarcRecord string

	// MX check
	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	}

	// SPF check
	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// DMARC check
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err == nil {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	// Output
	if hasMX {
		fmt.Println("✔️  Has MX (Mail Server): YES")
	} else {
		fmt.Println("❌ Has MX (Mail Server): NO")
	}

	if hasSPF {
		fmt.Println("✔️  Has SPF (Anti-Spoofing): YES")
		fmt.Printf("    ↪️ SPF Record: %s\n", spfRecord)
	} else {
		fmt.Println("❌ Has SPF (Anti-Spoofing): NO")
	}

	if hasDMARC {
		fmt.Println("✔️  Has DMARC: YES")
		fmt.Printf("    ↪️ DMARC Record: %s\n", dmarcRecord)
	} else {
		fmt.Println("❌ Has DMARC: NO")
	}
}
