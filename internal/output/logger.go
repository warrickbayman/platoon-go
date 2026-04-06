package output

import (
	"fmt"
	"os"
	"time"
)

func WriteToFile(path string, content string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(fmt.Errorf("failed to create file: %w", err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to close file: %w", err))
		}
	}(f)

	t := time.Now().Format("2006-01-02 15:04:05")

	_, er := fmt.Fprintf(f, "[%s] %s\n", t, content)

	if er != nil {
		fmt.Println(fmt.Errorf("failed to write to log file: %w", er))
	}
}

func ClearFile(path string) {
	err := os.Remove(path)

	if err != nil {
		fmt.Println(fmt.Errorf("failed to clear log file: %w", err))
	}
}
