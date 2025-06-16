# mkpsxgo
`mkpsxgo` is a library that provides a convenient Go API for interacting with the `mkpsxiso` and `dumpsxiso` cli tools.

### Credits
This library is merely a wrapper around the excellent `mkpsxiso` and `dumpsxiso` tools by Lameguy64. You can find their work here: [https://github.com/Lameguy64/mkpsxiso](https://github.com/Lameguy64/mkpsxiso)

### Installation
``go get github.com/murphy28/mkpsxgo``

### Features
This library features easy to use functions for interacting with `mkpsxiso` and `dumpsxiso`, and will automatically download the binaries for your OS.

### Dumping an ISO

```go
func main() {
	// Ensure mkpsxiso and dumpsxiso binaries are ready.
	if err := mkpsxgo.EnsureBinaries(); err != nil {
		log.Fatalf("Failed to ensure mkpsxiso binaries: %v", err)
	}

	// Example paths for dumping an ISO.
	sourceIsoPath := "/path/to/your/Kula Quest (Japan).bin" // Replace with your actual ISO path
	extractedFilesDir := "/path/to/extracted_data"
	rebuiltXMLPath := "/path/to/rebuilt_game.xml"

	dumpOpts := &mkpsxgo.DumpISOOptions{
		ExtractPath: extractedFilesDir,
		XMLPath:     rebuiltXMLPath,
	}

	// Dump the ISO to extract contents and XML.
	if err := mkpsxgo.DumpISO(sourceIsoPath, dumpOpts); err != nil {
		log.Printf("Error dumping ISO: %v", err)
	}
}
```

### Building an ISO

```go
func main() {
	// Ensure mkpsxiso and dumpsxiso binaries are ready.
	if err := mkpsxgo.EnsureBinaries(); err != nil {
		log.Fatalf("Failed to ensure mkpsxiso binaries: %v", err)
	}

	// Example paths for building an ISO.
	projectXMLPath := "/path/to/rebuilt_game.xml"
	outputIsoPath := "/path/to/Kula Quest (Japan).bin"
	outputCuePath := "/path/to/Kula Quest (Japan).cue"

	makeIsoOpts := &mkpsxgo.MakeISOOptions{
		Output:    outputIsoPath,
		CueFile:   outputCuePath,
		Overwrite: true,
	}

	// Build the ISO from the XML project file.
	if err := mkpsxgo.MakeISO(projectXMLPath, makeIsoOpts); err != nil {
		log.Printf("Error building ISO from XML: %v", err)
	}
}
```