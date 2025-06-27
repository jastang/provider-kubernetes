//go:build !uxp
// +build !uxp

package bootcheck

func CheckEnv() error {
	// No-op by default. Use build tags for build-time isolation of custom preflight checks.
	return nil
}
