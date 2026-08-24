//go:build !windows

package app

func enableTerminalANSI() bool {
	return true
}
