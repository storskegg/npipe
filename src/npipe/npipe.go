package npipe

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"runtime"
)

// Nothing really to see here, yet.
func mkServerPipe(path string) (string, error) {
	switch os := runtime.GOOS; os {
	case "darwin":
	case "freebsd":
	case "linux":
		// Nothing to see here.
		break
	case "windows":
		pipe, err := winio.DialPipe(path, nil)
		if err != nil {
			return "", err
		}
		fmt.Println(pipe)
	default:
		return "", error(os + " is not yet supported.")
	}
	return path, nil
}
