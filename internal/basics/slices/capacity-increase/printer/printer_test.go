package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSliceCapacityIncrease(t *testing.T) {
	capacities := []int{4, 8, 16, 32, 64, 128, 256, 512, 848, 1280}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintSliceCapacityIncrease(capacities)

	w.Close()
	os.Stdout = oldStdout

	var capturedOutput bytes.Buffer
	io.Copy(&capturedOutput, r)

	expectedOutput := `----- Slice capacity increase -----
================ =============== ============
Current Capacity Increase Factor New Capacity
---------------- --------------- ------------
4                2.00            8
8                2.00            16
16               2.00            32
32               2.00            64
64               2.00            128
128              2.00            256
256              2.00            512
512              1.66            848
848              1.51            1280
================ =============== ============
`
	assert.Equal(t, expectedOutput, capturedOutput.String())
}
