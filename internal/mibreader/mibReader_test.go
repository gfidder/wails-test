package mibreader

import (
	"fmt"
	"testing"

	"github.com/sleepinggenius2/gosmi/types"
	"github.com/stretchr/testify/assert"
)

func TestAppendOidNumbers(t *testing.T) {
	parameters := []struct {
		oidNum    string
		newOidNum types.SmiSubId
		expected  string
	}{
		{".1", 3, ".1.3"},
		{".1.3.6.1.3.1.585.4", 6175, ".1.3.6.1.3.1.585.4.6175"},
	}

	for i, parameter := range parameters {
		t.Run(fmt.Sprintf("Testing [%v]", i), func(t *testing.T) {
			actual := appendOidNumber(parameter.oidNum, parameter.newOidNum)
			assert.Equal(t, parameter.expected, actual)
		})
	}
}
