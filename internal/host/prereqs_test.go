package host

import "testing"

func TestWSLVersionFromRelease(t *testing.T) {
	for release, want := range map[string]int{
		// WSL2 kernels: the suffix was added later; older and custom
		// builds carry only "microsoft-standard".
		"5.15.153.1-microsoft-standard-WSL2":  2,
		"6.6.87.2-microsoft-standard-WSL2":    2,
		"4.19.43-microsoft-standard":          2,
		"4.19.104-microsoft-standard":         2,
		"5.10.16.3-microsoft-standard-wsl2\n": 2,
		// WSL1's Microsoft-suffixed Linux release.
		"4.4.0-19041-Microsoft":   1,
		"4.4.0-17763-microsoft":   1,
		"4.4.0-19041-Microsoft\n": 1,
		// Native Linux kernels.
		"6.8.0-45-generic": 0,
		"6.6.30-azure":     0,
		"":                 0,
		"   \n":            0,
	} {
		if got := wslVersionFromRelease(release); got != want {
			t.Errorf("wslVersionFromRelease(%q) = %d, want %d", release, got, want)
		}
	}
}
