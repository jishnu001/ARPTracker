package main

import (
	"fmt"
	"os/exec"
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

func notifyUser(message string) {

	title := "ARPTracker"

	if runtime.GOOS == "linux" {

		exec.Command("notify-send", "-i", "", title, message).Run()

	} else if runtime.GOOS == "darwin" {
		
		exec.Command("osascript", "-e", "display notification \""
						+message+"\" with title \""+title+"\"").Run()

	}

}

func main() {

	arpCache := make(map[string]string)
	getARPTable(arpCache)

	for {

		arpCacheTemp := make(map[string]string)
		getARPTable(arpCacheTemp)

		for key, val := range arpCache {
			valTemp, found := arpCacheTemp[key]
			if found {
				if val != valTemp {
					notifyUser(key +  "/"  + val +  " became -->" +  key +  "/" +  valTemp)
				}
			}

		}

		for key, val := range arpCacheTemp {
			_, found := arpCache[key]
			if !found {
				notifyUser("New IP/MAC in ARP cache: " + key + "/" + val)
			}

		}

		arpCache = arpCacheTemp
		time.Sleep(time.Second * 10)
	}

}
