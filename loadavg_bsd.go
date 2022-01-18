// +build freebsd dragonfly netbsd openbsd darwin

package loadavg

// #include <stdlib.h>
import "C"

import (
	"errors"
	"syscall"
	"time"
	"unsafe"
)

func Parse() (la [3]float64, err error) {
	avg := []C.double{0, 0, 0}

	n := C.getloadavg(&avg[0], C.int(len(avg)))

	if n == -1 {
		return [3]float32{0, 0, 0}, errors.New("load average unavailable")
	}
	return [3]float32{float32(avg[0]), float32(avg[1]), float32(avg[2])}, nil
}
