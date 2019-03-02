package osxarp

import (
	"log"
	"os/exec"
	"strings"
)

// ReadARPCache for OS X
func ReadARPCache(arpCache map[string]string) {
	out, err := exec.Command("bash", "-c", "arp -a").Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("The output is %s\n", out)
	str := string(out)
	list := strings.SplitAfter(str, "\n")
	for _, line := range list {
		//var ip string
		//var mac string
		if len(line) > 0 {
			index1 := strings.Index(line, "(")
			index2 := strings.Index(line, ")")

			ip := line[index1+1 : index2]

			//fmt.Println(ip)

			index1 = strings.Index(line, "at ")
			index2 = strings.Index(line, " on")

			mac := line[index1+2 : index2]
			mac = strings.TrimSpace(mac)
			//fmt.Println(mac)

			arpCache[ip] = mac
		}

	}
}
