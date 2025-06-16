package mkpsxgo

import (
	"fmt"
	"os"
	"os/exec"
)

func RunMkpsxiso(args ...string) error {
	// Ensure the necessary binaries are present
	panicIfBinariesMissing()

	// Create the command with the provided arguments
	cmd := exec.Command(MkpsxisoPath, args...)

	// Set the command's output to standard output and standard error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running command: %s %s\n", MkpsxisoPath, args)

	// Run the command and return any error that occurs
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run mkpsxiso with arguments %v: %w", args, err)
	}

	return nil
}

type MakeISOOptions struct {
	Overwrite bool   // Always overwrite ISO image files
	Quiet     bool   // Quiet mode (suppress all but warnings and errors)
	Output    string // Specify output file (overrides image_name attribute)
	CueFile   string // Specify cue sheet file (overrides cue_sheet attribute)
	Label     string // Specify volume ID (overrides volume element)
	NoXA      bool   // Do not generate CD-XA extended file attributes (plain ISO9660) (XA data can still be included but not recommended)
}

func MakeISO(xmlPath string, options *MakeISOOptions) error {
	// Ensure the necessary binaries are present
	panicIfBinariesMissing()

	// Validate the xmlPath
	if xmlPath == "" {
		return fmt.Errorf("xmlPath cannot be empty")
	}

	// Prepare the arguments for mkpsxiso
	args := []string{}

	if options != nil {
		if options.Overwrite {
			args = append(args, "-y")
		}
		if options.Quiet {
			args = append(args, "-q")
		}
		if options.Output != "" {
			args = append(args, "-o", options.Output)
		}
		if options.CueFile != "" {
			args = append(args, "-c", options.CueFile)
		}
		if options.Label != "" {
			args = append(args, "-l", options.Label)
		}
		if options.NoXA {
			args = append(args, "-noxa")
		}
	}

	// Add the xmlPath as the last argument
	args = append(args, xmlPath)

	return RunMkpsxiso(args...)
}

// RebuildXML rebuilds the specified XML project file using mkpsxiso's newest schema.
func RebuildXML(xmlPath string) error {
	if xmlPath == "" {
		return fmt.Errorf("xmlPath cannot be empty")
	}
	return RunMkpsxiso("-rebuildxml", xmlPath)
}

// GenerateLBALog generates a log of file LBA locations.
// Optionally, it can also suppress ISO generation.
type LBAOptions struct {
	SuppressISOGen bool // -noisogen
}

// GenerateLBALog creates a log of file LBA locations.
// The output file for the LBA log is specified by lbaLogFilePath.
// If suppressISOGen is true, the ISO image will not be generated (-noisogen).
func GenerateLBALog(xmlPath, lbaLogFilePath string, opts *LBAOptions) error {
	if xmlPath == "" || lbaLogFilePath == "" {
		return fmt.Errorf("XML project file path and LBA log file path cannot be empty")
	}

	args := []string{"-lba", lbaLogFilePath}
	if opts != nil && opts.SuppressISOGen {
		args = append(args, "-noisogen")
	}
	args = append(args, xmlPath)

	return RunMkpsxiso(args...)
}

// GenerateLBAHeader generates a C header file of file LBA locations.
// Optionally, it can also suppress ISO generation.
func GenerateLBAHeader(xmlPath, lbaHeaderFilePath string, opts *LBAOptions) error {
	if xmlPath == "" || lbaHeaderFilePath == "" {
		return fmt.Errorf("XML project file path and LBA header file path cannot be empty")
	}

	args := []string{"-lbahead", lbaHeaderFilePath}
	if opts != nil && opts.SuppressISOGen {
		args = append(args, "-noisogen")
	}
	args = append(args, xmlPath)

	return RunMkpsxiso(args...)
}
