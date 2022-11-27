package go_module_phytagoras

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Test")

	m.Run()

	fmt.Println("Setelah tes")
}

func TestFindC(t *testing.T) {
	res := FindC(4.0, 3.0)
	assert.Equal(t, 5.0, res)
}

func TestZeroOperand(t *testing.T) {
	t.Run("Left", func(t *testing.T) {
		res := FindC(0, 5.0)
		require.Equal(t, res, 0.0)
	})
	t.Run("Right", func(t *testing.T) {
		res := FindC(3.0, 0.0)
		require.Equal(t, res, 0.0)
	})
}

func TestTable(t *testing.T) {

	tests := []struct {
		name         string
		leftOperand  float64
		rightOperand float64
		expected     float64
	}{
		{
			name:         "All",
			leftOperand:  3.0,
			rightOperand: 4.0,
			expected:     5.0,
		},
		{
			name:         "Left Zero",
			leftOperand:  0.0,
			rightOperand: 4.0,
			expected:     0.0,
		},
		{
			name:         "Right Zero",
			leftOperand:  3.0,
			rightOperand: 0.0,
			expected:     0.0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := FindC(test.leftOperand, test.rightOperand)
			require.Equal(t, res, test.expected)
		})
	}
}
