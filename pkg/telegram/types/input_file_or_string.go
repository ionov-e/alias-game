package types

import "encoding/json"

type InputFileOrString struct {
	File        *InputFile // Represents a file to be uploaded
	FileIDOrURL string     // Represents a file ID or URL as a string
}

// InputFile represents a file to be uploaded via multipart/form-data.
type InputFile struct {
	FilePath string // Path to the file on the local disk
	Name     string // Name of the file for multipart
}

// MarshalJSON ensures the correct value is serialized based on the set field.
func (i InputFileOrString) MarshalJSON() ([]byte, error) {
	if i.File != nil {
		// For files, return the attach:// format
		return json.Marshal("attach://" + i.File.Name)
	}
	// For strings, return the string directly
	return json.Marshal(i.FileIDOrURL)
}

// NewInputFile creates a new InputFileOrString for uploading files.
func NewInputFile(filePath, name string) InputFileOrString {
	return InputFileOrString{
		File: &InputFile{
			FilePath: filePath,
			Name:     name,
		},
	}
}

// NewFileIDOrURL creates a new InputFileOrString for file ID or URL.
func NewFileIDOrURL(fileIDOrURL string) InputFileOrString {
	return InputFileOrString{
		File:        nil,
		FileIDOrURL: fileIDOrURL,
	}
}
