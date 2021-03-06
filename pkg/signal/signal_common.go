package signal

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

// ParseSignal translates a string to a valid syscall signal.
// It returns an error if the signal map doesn't include the given signal.
func ParseSignal(rawSignal string) (syscall.Signal, error) {
	s, err := strconv.Atoi(rawSignal)
	if err == nil {
		if s == 0 {
			return -1, fmt.Errorf("invalid signal: %s", rawSignal)
		}
		return syscall.Signal(s), nil
	}
	sig, ok := SignalMap[strings.TrimPrefix(strings.ToUpper(rawSignal), "SIG")]
	if !ok {
		return -1, fmt.Errorf("invalid signal: %s", rawSignal)
	}
	return sig, nil
}

// ParseSignalNameOrNumber translates a string to a valid syscall signal.  Input
// can be a name or number representation i.e. "KILL" "9".
func ParseSignalNameOrNumber(rawSignal string) (syscall.Signal, error) {
	basename := strings.TrimPrefix(rawSignal, "-")
	s, err := ParseSignal(basename)
	if err == nil {
		return s, nil
	}
	for k, v := range SignalMap {
		if k == strings.ToUpper(basename) {
			return v, nil
		}
	}
	return -1, fmt.Errorf("invalid signal: %s", basename)
}
