package telebot

import (
	"fmt"
	"os"
)

// File object represents any sort of file.
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
	FileBytes []byte `json:"file_bytes,omitempty"`

	// Local absolute path to file on local file system.
	filename string
}

// NewFile attempts to create a File object, leading to a real
// file on the file system, that could be uploaded later.
//
// Notice that NewFile doesn't upload file, but only creates
// a descriptor for it.
func NewFile(path string) (File, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return File{}, fmt.Errorf("telebot: '%s' does not exist", path)
	}

	return File{filename: path}, nil
}

// NewFileBytes creates a File object using byte array
// In this case file is not presented in file system.
//
// fileName in this method used only for headers, so could be any
// with correct file extension
func NewFileBytes(data []byte, fileName string) (File) {
	return File{FileBytes: data, filename: fileName}
}

// Exists says whether the file presents on Telegram servers or not.
func (f File) Exists() bool {
	return f.FileID != ""
}

// Local returns location of file on local file system, if it's
// actually there, otherwise returns empty string.
func (f File) Local() string {
	return f.filename
}
