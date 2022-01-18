// +build freebsd dragonfly netbsd openbsd darwin

package loadavg

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/msaf1980/go-stringutils"
)

func Parse() (la [3]float64, err error) {
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

	la[0], err = strconv.ParseFloat(values[0], 64)
	if err != nil {
		return la, errors.New("LoadAverage1 parse error")
	}

	la[1], err = strconv.ParseFloat(values[1], 64)
	if err != nil {
		return la, errors.New("LoadAverage5 parse error")
	}

	la[2], err = strconv.ParseFloat(values[2], 64)
	if err != nil {
		return la, errors.New("LoadAverage10 parse error")
	}

	return la, nil
}
