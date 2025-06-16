package mkpsxgo

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
	CurrentMkpsxisoVersion = "2.10"

	DownloadLinuxURL   = "https://github.com/Lameguy64/mkpsxiso/releases/download/v" + CurrentMkpsxisoVersion + "/mkpsxiso-" + CurrentMkpsxisoVersion + "-Linux.zip"
	DownloadWindowsURL = "https://github.com/Lameguy64/mkpsxiso/releases/download/v" + CurrentMkpsxisoVersion + "/mkpsxiso-" + CurrentMkpsxisoVersion + "-win64.zip"

	ChecksumLinux   = "a26181df7c7409d7233c39ddb018d53c599f6925ad49bd65f80f5a5888b734af"
	ChecksumWindows = "158c027f28fff4ad4ac227fe686b65692a14f6a4bbab39e99d8f9455e4b70e6b"
)

// Declare binary paths
var UserDirectory string
var MkpsxisoPath string
var DumpsxisoPath string

// Initialize the paths for mkpsxiso and dumpsxiso binaries
func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home directory: " + err.Error())
	}

	UserDirectory = filepath.Join(userDir, ".mkpsxgo")

	MkpsxisoPath = filepath.Join(userDir, ".mkpsxgo", "bin", "mkpsxiso")
	DumpsxisoPath = filepath.Join(userDir, ".mkpsxgo", "bin", "dumpsxiso")

	// Ensure the .mkpsxgo/bin directory exists
	if err := os.MkdirAll(filepath.Dir(MkpsxisoPath), 0755); err != nil {
		panic("Failed to create .mkpsxgo/bin directory: " + err.Error())
	}

	// Append .exe for Windows binaries
	if runtime.GOOS == "windows" {
		MkpsxisoPath += ".exe"
		DumpsxisoPath += ".exe"
	}
}
