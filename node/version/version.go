package version

import (
	"fmt"
	"runtime"
)

var (
	Version = "1.0.0"
)

func Info(name, commit, buildDate string) {
	fmt.Println("Version:", Version)
	if commit != "" {
		fmt.Println("Git Commit:", commit)
	}
	if buildDate != "" {
		fmt.Println("Build Date:", buildDate)
	}
	if name != "" {
		fmt.Println("Binary Name:", name)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
}
