package mkpsxgo

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// downloadFile downloads a file from the given URL to a temporary location.
// It returns the path to the downloaded file.
func downloadFile(url string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "mkpsxgo-download")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %w", err)
	}

	filePath := filepath.Join(tmpDir, "makepsxiso.zip")
	file, err := os.Create(filePath)
	if err != nil {
		os.RemoveAll(tmpDir) // Clean up if file creation fails
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", fmt.Errorf("failed to download from %s: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		os.RemoveAll(tmpDir)
		return "", fmt.Errorf("failed to download binaries: received status %s", response.Status)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", fmt.Errorf("failed to copy download content: %w", err)
	}

	return filePath, nil
}

// verifyChecksum calculates the SHA256 checksum of the file at filePath
// and compares it against the expectedChecksumHex.
func verifyChecksum(filePath, expectedChecksumHex string) error {
	calculatedChecksumHex, err := calculateSHA256(filePath)
	if err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}

	if expectedChecksumHex != calculatedChecksumHex {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksumHex, calculatedChecksumHex)
	}
	return nil
}

// unzipBinaries unzips the given zipFile and moves the
// mkpsxiso and dumpsxiso binaries to the user's .mkpsxgo/bin directory.
func unzipBinaries(zipFile string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("failed to open zip archive: %w", err)
	}
	defer reader.Close()

	userDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	binariesDir := filepath.Join(userDir, ".mkpsxgo", "bin")
	if err := os.MkdirAll(binariesDir, 0755); err != nil {
		return fmt.Errorf("failed to create binaries directory %s: %w", binariesDir, err)
	}

	for _, file := range reader.File {
		fileName := file.FileInfo().Name()

		if strings.Contains(fileName, "mkpsxiso") || strings.Contains(fileName, "dumpsxiso") {
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("failed to open file %s in zip: %w", fileName, err)
			}
			defer rc.Close() // Ensure all opened files in zip are closed

			destFilePath := filepath.Join(binariesDir, fileName)
			destFile, err := os.Create(destFilePath)
			if err != nil {
				return fmt.Errorf("failed to create destination file %s: %w", destFilePath, err)
			}

			if _, err := io.Copy(destFile, rc); err != nil {
				destFile.Close() // Close on error
				return fmt.Errorf("failed to copy %s to %s: %w", fileName, destFilePath, err)
			}
			destFile.Close() // Close after successful copy

			if err := os.Chmod(destFilePath, 0755); err != nil {
				return fmt.Errorf("failed to make %s executable: %w", destFilePath, err)
			}
		}
	}
	return nil
}
