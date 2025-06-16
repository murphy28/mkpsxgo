package mkpsxgo

import (
	"os"
	"path/filepath"
)

// EnsureBinaries orchestrates the entire process of downloading,
// checksumming, unzipping, and moving the binaries to
// the appropriate location.
func EnsureBinaries() error {
	// Check if the binaries already exist
	mkpsxisoFound, dumpsxisoFound := areBinariesPresent()

	// If both binaries are found, no need to download
	if mkpsxisoFound && dumpsxisoFound {
		return nil
	}

	// Get the download URL and expected checksum
	downloadURL, expectedChecksumHex, err := getDownloadInfo()
	if err != nil {
		return err
	}

	// Download the file to a temporary location
	tmpFile, err := downloadFile(downloadURL)
	if err != nil {
		return err
	}
	defer os.RemoveAll(filepath.Dir(tmpFile)) // Clean up temporary directory

	// Verify the checksum of the downloaded file
	if err := verifyChecksum(tmpFile, expectedChecksumHex); err != nil {
		return err
	}

	// Unzip the downloaded file to the user directory
	if err := unzipBinaries(tmpFile); err != nil {
		return err
	}

	return nil
}

func Test() {
	panicIfBinariesMissing()
}
