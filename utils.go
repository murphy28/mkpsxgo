package mkpsxgo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
)

// getDownloadInfo retrieves the download information (URL and checksum)
// based on the operating system.
func getDownloadInfo() (string, string, error) {
	var downloadURL, checksum string

	// Determine if the system is Windows or Linux using the GOOS environment variable.
	if runtime.GOOS == "windows" {
		// These constants are expected to be defined elsewhere, e.g., in constants.go
		// For demonstration, let's assume they are globally available or imported.
		downloadURL = DownloadWindowsURL
		checksum = ChecksumWindows
	} else if runtime.GOOS == "linux" {
		downloadURL = DownloadLinuxURL
		checksum = ChecksumLinux
	}

	// If the environment variable is not set or not recognized, return an error.
	if downloadURL == "" || checksum == "" {
		return "", "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return downloadURL, checksum, nil
}

// calculateSHA256 computes the SHA256 checksum of a file.
func calculateSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s for checksum calculation: %w", filePath, err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to read file %s for checksum calculation: %w", filePath, err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func areBinariesPresent() (bool, bool) {
	// Check if mkpsxiso exists
	mkpsxisoFound := false
	if _, err := os.Stat(MkpsxisoPath); err == nil {
		mkpsxisoFound = true
	}

	// Check if dumpsxiso exists
	dumpsxisoFound := false
	if _, err := os.Stat(DumpsxisoPath); err == nil {
		dumpsxisoFound = true
	}

	return mkpsxisoFound, dumpsxisoFound
}

func panicIfBinariesMissing() {
	mkpsxiso, dumpsxiso := areBinariesPresent()
	if !mkpsxiso || !dumpsxiso {
		panic(fmt.Sprintf("Required binaries not found: mkpsxiso (%s), dumpsxiso (%s). Please run mkpsxgo.EnsureBinaries() to download them.", MkpsxisoPath, DumpsxisoPath))
	}
}
