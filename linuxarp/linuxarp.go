package linuxarp

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
	"strings"
)

// ReadARPCache for linux OS
func ReadARPCache(arpCache map[string]string) {

	if runtime.GOOS != "linux" {
		fmt.Println("Error not Linux OS")
		return
	}

	file, err := ioutil.ReadFile("/proc/net/arp")

	if err != nil {
		fmt.Println("Error cannot find APR cache")
	}

	str := string(file)
	list := strings.SplitAfter(str, "\n")

	for i, line := range list {
		var ip string
		var mac string

		if i > 0 {
			tokens := strings.SplitAfter(line, " ")
			for k, token := range tokens {

				if k == 0 {
					ip = token
				}
				matched, err := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){2}([0-9A-Fa-f]{2})", token)

				if err != nil {
					fmt.Println("Not a MAC address")
				}

				if matched {
					mac = token
				}

			}
			if len(ip) != 0 && len(mac) != 0 {
				arpCache[ip] = mac
			}
		}

	}

}
