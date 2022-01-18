package loadavg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	gotLa, err := Parse()
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}

	assert.Greater(t, gotLa[0], float32(0.0), "loadavg[0]")
	assert.Greater(t, gotLa[1], float32(0.0), "loadavg[1]")
	assert.Greater(t, gotLa[2], float32(0.0), "loadavg[2]")
}
