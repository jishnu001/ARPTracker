package main

import (
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

		for _, val := range arpCache {
			found := false
			for _, val1 := range arpCacheTemp {
				if val == val1 {
					found = true
					//fmt.Println("second map")
					//fmt.Println("key:", key1, "val:", val1)
					break
				}
			}
			if !found {
				arpCache = arpCacheTemp
			}
		}
		time.Sleep(time.Second * 10)
	}

}
