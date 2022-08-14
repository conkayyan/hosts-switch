package hosts

import "os"

func getHostPath() string {
	return os.Getenv("SystemRoot") + "\\System32\\drivers\\etc\\hosts"
}
