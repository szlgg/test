package gerroraa

import "runtime"

// stack represents a stack of program counters.
type stack []uintptr

const (
	// maxStackDepth marks the max stack depth for error back traces.
	maxStackDepth = 32
)

func callers(skip ...int) stack {
	var (
		pcs [maxStackDepth]uintptr
		n   = 3
	)
	if len(skip) > 0 {
		n += skip[0]
	}
	return pcs[:runtime.Callers(n, pcs[:])]
}