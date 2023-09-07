package fileutils

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	ptr "github.com/ShindouMihou/go-little-utils"
	"github.com/ShindouMihou/go-little-utils/buffer"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Close is a little handy function that is used to print when a file failed to close.
func Close(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println("err: failed to close ", f.Name(), " body")
	}
}

// MkdirParent creates the parent folders of the path.
func MkdirParent(file string) error {
	if strings.Contains(file, "\\") || strings.Contains(file, "/") {
		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// Create creates a new file by first creating the parent folders before actually creating
// the file, ensuring that there is a directory for the file.
func Create(file string) (*os.File, error) {
	if err := MkdirParent(file); err != nil {
		return nil, err
	}
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Save creates the file in question, and pushes the data, this does not override the existing content.
func Save(file string, data []byte) error {
	f, err := Create(file)
	if err != nil {
		return err
	}
	defer Close(f)
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// SaveOrOverwrite creates the file in question, and pushes the data. If the file already exists,
// then it truncates the file before pushing the data.
func SaveOrOverwrite(file string, data []byte) error {
	f, err := Create(file)
	if err != nil {
		return err
	}
	defer Close(f)
	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// SaveBuffer creates the file in question, and pushes the data from the buffer
// by a size of 4096 bytes, this does not override the existing content.
func SaveBuffer(file string, data bufio.Reader) error {
	f, err := Create(file)
	if err != nil {
		return err
	}
	defer Close(f)
	if err := buffer.Read(data, 4_096, func(bytes []byte) error {
		if _, err := f.Write(bytes); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// SaveOrOverwriteBuffer creates the file in question, and pushes the data by 4096 bytes each. If the file already exists,
// then it truncates the file before pushing the data.
func SaveOrOverwriteBuffer(file string, data bufio.Reader) error {
	f, err := Create(file)
	if err != nil {
		return err
	}
	defer Close(f)
	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}
	if err := buffer.Read(data, 4_096, func(bytes []byte) error {
		if _, err := f.Write(bytes); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// CopyWithHash copies the source to the destination while also creating a hash of the original,
// then returning the hash.
func CopyWithHash(source, dest string) (*string, error) {
	f, err := Create(dest)
	if err != nil {
		return nil, err
	}
	defer Close(f)
	r, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer Close(r)
	hash := sha256.New()
	r2 := io.TeeReader(r, hash)
	_, err = io.Copy(f, r2)
	if err != nil {
		return nil, err
	}
	return ptr.Ptr(hex.EncodeToString(hash.Sum(nil))), nil
}

// SanitizeFilePath simply sanitizes the file path by removing all relative paths and replacing spaces with
// underscores.
func SanitizeFilePath(key string) string {
	key = filepath.Clean(filepath.Base(key))
	key = strings.ToLower(key)
	key = strings.ReplaceAll(key, " ", "_")
	return key
}

var homeDirectory = ""

// GetHomeDir gets the home directory of the system, this panics if it cannot get the home directory.
// The result of the initial load is cached.
func GetHomeDir() string {
	if homeDirectory == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic("failed to get home directory")
		}
		homeDirectory = home
	}
	return homeDirectory
}

// JoinHomePath joins the other path with the home directory, this panics if it cannot get the home directory.
func JoinHomePath(paths ...string) string {
	return filepath.Join(GetHomeDir(), filepath.Join(paths...))
}

var workingDirectory = ""

// GetWorkingDirectory gets the working directory of the executable, this panics if it cannot get the working directory.
// The result of the initial load is cached.
func GetWorkingDirectory() string {
	if workingDirectory == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic("failed to get working directory")
		}
		workingDirectory = wd
	}
	return workingDirectory
}

// JoinWorkingDirectory joins the other path with the working directory, this panics if it cannot get the working directory.
func JoinWorkingDirectory(paths ...string) string {
	return filepath.Join(GetWorkingDirectory(), filepath.Join(paths...))
}
