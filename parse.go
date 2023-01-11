package goupx

import (
	"fmt"
	"regexp"
	"strconv"
)

// parsedError contain information from command error message.
type parsedError struct {
	errorMessage error
}

// parseError extract information from command error message.
func parseError(errorMessage []byte) parsedError {
	return parsedError{
		errorMessage: fmt.Errorf("%v", string(errorMessage)),
	}
}

// GetErrorMessage return the error value from the command execution.
func (pe *parsedError) GetErrorMessage() error {
	return pe.errorMessage
}

// parsedOutput contain information from command output message.
type parsedOutput struct {
	fileSizeOrig     uint64
	fileSizeCompress uint64
	ratio            float64
	format           string
	name             string
}

// parseOutput extract information from command output message.
func parseOutput(outputMessage []byte) parsedOutput {
	upxParsed := parsedOutput{}
	re := regexp.MustCompile(`(?P<originalSize>\d+)\s[<\->]+\s*(?P<compressedSize>\d+)\s*(?P<ratio>\d+.\d+)\%\s*(?P<format>\w+\/\w+)\s*(?P<name>\w+)`)
	match := re.FindStringSubmatch(string(outputMessage))
	for i, name := range re.SubexpNames() {
		switch name {
		case "originalSize":
			fileSize, _ := strconv.ParseUint(match[i], 10, 64)
			upxParsed.fileSizeOrig = fileSize
		case "compressedSize":
			fileSize, _ := strconv.ParseUint(match[i], 10, 64)
			upxParsed.fileSizeCompress = fileSize
		case "ratio":
			ratio, _ := strconv.ParseFloat(match[i], 64)
			upxParsed.ratio = ratio
		case "format":
			upxParsed.format = match[i]
		case "name":
			upxParsed.name = match[i]
		default:
		}
	}
	return upxParsed
}

// GetOriginalFileSize return the original file size value from the command execution.
func (po *parsedOutput) GetOriginalFileSize() uint64 {
	return po.fileSizeOrig
}

// GetCompressedFileSize return the compressed file size value from the command execution.
func (po *parsedOutput) GetCompressedFileSize() uint64 {
	return po.fileSizeCompress
}

// GetRatio return the ratio value from the command execution.
func (po *parsedOutput) GetRatio() float64 {
	return po.ratio
}

// GetFormat return the format value from the command execution.
func (po *parsedOutput) GetFormat() string {
	return po.format
}

// GetName return the name value from the command execution.
func (po *parsedOutput) GetName() string {
	return po.name
}
