package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bool64/dev/version"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	target := flag.Float64("target", 20.0, "Target total CPU usage (%)")
	ver := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *ver {
		fmt.Println(version.Info().Version)
		return
	}

	for {
		perc, err := cpu.Percent(0, false)

		if err != nil || len(perc) != 1 {
			log.Fatal(err)
		}

		cu := perc[0]

		if cu < *target {
			log.Printf("😡 CPU Usage: %.2f%%, punch(%d, %s)\n", cu, int(iterations), time.Duration(delay).String())
			punch()
		} else {
			log.Printf("😴 CPU Usage: %.2f%%\n", cu)
			if punches == 0 {
				// Longer sleep on stable load.
				time.Sleep(5 * time.Second)
			} else {
				// Punch too strong, increase delay.
				if punches < 3 {
					delay *= 1.1
				}

				// Shorter sleep after punching to avoid CPU usage drop.
				punches = 0
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

var (
	iterations = 10000.0
	delay      = float64(50 * time.Microsecond)
	punches    = 0
)

func punch() {
	start := time.Now()

	wg := sync.WaitGroup{}
	for c := 0; c < runtime.NumCPU(); c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var m map[string]string
			for i := 0; i <= int(iterations); i++ {
				_ = json.Unmarshal([]byte(`{"foo":"bar"}`), &m)
				time.Sleep(time.Duration(delay))
			}
		}()
	}
	wg.Wait()

	if time.Since(start) > time.Second {
		iterations *= 0.9
	} else {
		iterations *= 1.1
	}

	punches++

	if punches > 5 {
		delay *= 0.9
	}
}
