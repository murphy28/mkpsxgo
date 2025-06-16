package mkpsxgo

import (
	"fmt"
	"os"
	"os/exec"
)

func RunDumpsxiso(args ...string) error {
	// Ensure the necessary binaries are present
	panicIfBinariesMissing()

	// Create the command with the provided arguments
	cmd := exec.Command(DumpsxisoPath, args...)

	// Set the command's output to standard output and standard error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running command: %s %s\n", DumpsxisoPath, args)

	// Run the command and return any error that occurs
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run dumpsxiso with arguments %v: %w", args, err)
	}

	return nil
}

type DumpISOOptions struct {
	ExtractPath     string // Optional destination directory for extracted files. (Defaults to dumpsxiso dir)
	XMLPath         string // Optional XML name/destination of MKPSXISO compatible script for later rebuilding. (Defaults to dumpsxiso dir)
	SortByDirectory bool   // Outputs a "pretty" XML script where entries are grouped in directories, instead of strictly following their original order on the disc.
	Codec           string // Codec to encode CDDA/DA audio. wave is default. Supported codecs: "wave", "pcm", "flac"
	PathTable       bool   // Instead of going through the file system, go to every known directory in order; helps with deobfuscating
}

func DumpISO(isoPath string, options *DumpISOOptions) error {
	// Ensure the necessary binaries are present
	panicIfBinariesMissing()

	// Validate the isoPath
	if isoPath == "" {
		return fmt.Errorf("isoPath cannot be empty")
	}

	// Prepare the arguments for dumpsxiso
	args := []string{isoPath}

	if options != nil {
		if options.ExtractPath != "" {
			args = append(args, "-x", options.ExtractPath)
		}
		if options.XMLPath != "" {
			args = append(args, "-s", options.XMLPath)
		}
		if options.SortByDirectory {
			args = append(args, "-S")
		}
		if options.Codec != "" {
			args = append(args, "-e", options.Codec)
		}
		if options.PathTable {
			args = append(args, "-pt")
		}
	}

	return RunDumpsxiso(args...)
}
