//go:build !debug

package kit

func DebugLog(v ...any) {
	// no-op for non-debug builds to improve coverage metrics of main logic
}
