package loadavg

import (
	"fmt"
	"runtime"
)

func Parse() (la [3]float64, err error) {
	return la, fmt.Errorf("%s not supported", runtime.GOOS)
}
