package kit

import (
	"testing"
)

func BenchmarkDebugLog(b *testing.B) {
	// Ensure debug is disabled for benchmark to test overhead when off, 
    // or enabled if that's the target. 
    // Here we just test the function call.
	for i := 0; i < b.N; i++ {
		DebugLog("test")
	}
}
