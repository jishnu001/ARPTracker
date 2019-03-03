package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jishnu001/ARPTracker/linuxarp"
	"github.com/jishnu001/ARPTracker/osxarp"
)

func getARPTable(arpCache map[string]string) {
	if runtime.GOOS == "linux" {
		linuxarp.ReadARPCache(arpCache)
	} else if runtime.GOOS == "darwin" {
		osxarp.ReadARPCache(arpCache)
	}
}

func main() {

	arpCache := make(map[string]string)
	getARPTable(arpCache)

	for {

		arpCacheTemp := make(map[string]string)
		getARPTable(arpCacheTemp)

		/*for ip, mac := range arpCache {
			fmt.Println(ip, mac)
		}*/

		for key, val := range arpCache {
			valTemp, found := arpCacheTemp[key]
			if found {
				if val != valTemp {
					fmt.Println(key, "/", val, " became -->", key, "/", valTemp)
				}
			}

		}

		for key, val := range arpCacheTemp {
			_, found := arpCache[key]
			if !found {
				fmt.Println("New IP/MAC in ARP cache: ", key, "/", val)
			}

		}

		arpCache = arpCacheTemp
		time.Sleep(time.Second * 10)
	}

}
