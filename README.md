# go-loadavg
A cross-platform golang library for retrieving load average.

## How to use

```go
package main

import (
	"github.com/msaf1980/go-loadavg"
	"log"
)

func main() {
	la, err := loadavg.Parse()
	if err != nil {
		log.Fatal(err)
	}

    switch runtime.GOOS {
	case "linux":
		log.Printf("load averages: loadavg1 %.2f loadavg5 %.2f loadavg10 %.2f", la[0], la[1], la[2])
	default:
		log.Printf("load averages: loadavg1 %.2f loadavg5 %.2f loadavg15 %.2f", la[0], la[1], la[2])
	}
}

```

## Supported platform

|    GOOS   |     Source     |   Support   |            Result              |
|:---------:|:--------------:|:-----------:|:------------------------------:|
|   darwin  |  getloadavg()  | coming soon |                                |
| dragonfly |  getloadavg()  | coming soon |                                |
|  freebsd  |  /proc/loadavg | coming soon |                                |
|   linux   |  /proc/loadavg |      O      | load_avg1 load_avg5 load_avg10 |
|   netbsd  |  getloadavg()  | coming soon |                                |
|  openbsd  |  getloadavg()  | coming soon |                                |
|   plan9   | /dev/sysstat ? |      ?      |                                |
|  solaris  |        ?       |      ?      |                                |
|  windows  |        ?       |      ?      |                                |

## License
This code is licensed under the MIT license.
