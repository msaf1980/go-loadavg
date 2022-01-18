// +build linux

package loadavg

import (
	"errors"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"

	"github.com/msaf1980/go-stringutils"
)

func Parse() (la [3]float32, err error) {
	switch runtime.GOOS {
	case "linux":
		return parse_linux()
	default:
		return [3]float32{0.0}, errors.New("loadavg unimplemented on " + runtime.GOOS)
	}
}

func parse_linux() (la [3]float32, err error) {
	raw, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return la, err
	}

	s := strings.TrimRight(stringutils.UnsafeString(raw), "\n")
	buf := make([]string, 4)
	values := stringutils.SplitN(s, " ", buf)
	if len(values) != 4 {
		return la, errors.New("/proc/loadavg field count mismatch")
	}

	n, err := strconv.ParseFloat(values[0], 32)
	if err != nil {
		return la, errors.New("LoadAverage1 parse error")
	}
	la[0] = float32(n)

	n, err = strconv.ParseFloat(values[1], 64)
	if err != nil {
		return la, errors.New("LoadAverage5 parse error")
	}
	la[1] = float32(n)

	n, err = strconv.ParseFloat(values[2], 64)
	if err != nil {
		return la, errors.New("LoadAverage10 parse error")
	}
	la[2] = float32(n)

	return la, nil
}
