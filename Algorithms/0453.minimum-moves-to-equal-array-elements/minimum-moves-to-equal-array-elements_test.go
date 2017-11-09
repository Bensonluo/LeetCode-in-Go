package Problem0453

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{1, 2, 3},
		3,
	},

	// 可以有多个 testcase
}

func Test_minMoves(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minMoves(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_minMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minMoves(tc.nums)
		}
	}
}
