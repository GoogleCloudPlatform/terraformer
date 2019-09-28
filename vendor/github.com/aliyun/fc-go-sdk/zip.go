package fc

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TempZipDir zips everything from source dir into a temporary zip file which doesn't include the source dir but its content.
// Return the location of the temporary zip file.
func TempZipDir(dir string) (string, error) {
	// Collect files to zip.
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	files := []string{}
	for _, f := range fs {
		files = append(files, filepath.Join(dir, f.Name()))
	}
	return TmpZip(files)
}

// TmpZip everything from source file into a temporary zip file.
// Return the location of the temporary zip file.
func TmpZip(files []string) (string, error) {
	zipfile, err := ioutil.TempFile("", "fc_temp_file_")
	if err != nil {
		return "", err
	}

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	for _, source := range files {
		if err := compress(source, archive); err != nil {
			return "", err
		}
	}

	return zipfile.Name(), nil
}

// Zip everything from the source (either file/directory) recursively into target zip file.
func Zip(files []string, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	for _, f := range files {
		if err := compress(f, archive); err != nil {
			return err
		}
	}

	return nil
}

// Compress zips source file into archive file.
func compress(source string, archive *zip.Writer) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if srcInfo.IsDir() {
			header.Name = filepath.Join(filepath.Base(source), strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if info.Mode()&os.ModeSymlink != 0 {
			// handle symbol link file
			dest, err := os.Readlink(path)
			if err != nil {
				return err
			}
			_, err = writer.Write([]byte(dest))
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
}

// ZipDir zip up a directory and preserve symlinks and empty directories.
func ZipDir(srcDir string, output io.Writer) error {
	zipWriter := zip.NewWriter(output)
	defer zipWriter.Close()

	// Convert the input dir path to absolute path if necessary.
	srcDir, err := filepath.Abs(srcDir)
	if err != nil {
		return err
	}

	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// handle the decode issue for windows : set default utf8.
		header.Flags |= 1 << 11
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		// handle windows path issue: replace "\\" to "/"
		header.Name = filepath.ToSlash(relPath)
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			dest, err := os.Readlink(path)
			if err != nil {
				return err
			}
			_, err = writer.Write([]byte(dest))
			return err
		}
		if info.IsDir() {
			// dir need create one entry to avoid of setting default permission
			_, err = writer.Write([]byte(""))
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
}
