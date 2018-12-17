package emoji

import (
	"fmt"
	"os/exec"
	"path"
)

//ConvertToPng converts the give svgPath into a new png file at the path
//pngOutput using the specified height and width in pixels for the resulting
//image
func ConvertToPng(svgPath string, pngOutPath string, width int, height int) {

	inkscapePath := path.Join("C:", "Program Files", "Inkscape", "inkscape.exe")
	command := exec.Command(
		inkscapePath,
		fmt.Sprintf("--file=%s", svgPath),
		fmt.Sprintf("--export-png=%s", pngOutPath),
		fmt.Sprintf("--export-width=%d", width),
		fmt.Sprintf("--export-height=%d", width))
	err := command.Run()
	if err != nil {
		check(err)
	}
}
