package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		fmt.Println("open error")
		return
	}
	defer in.Close()

	_, err = os.Stat(filepath.Dir(dst))
	if err != nil {
		fmt.Printf("attempting to copy %s to %s. directory %s does not exist", src, dst, filepath.Dir(dst))
		return
	}

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("create error")
		return
	}

	defer func() {
		if e := out.Close(); e != nil {
			fmt.Println("close error")
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		fmt.Println("copy error")
		return
	}

	err = out.Sync()
	if err != nil {
		fmt.Println("sync error")
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		fmt.Println("stat error")
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		fmt.Println("chmod error")
		return
	}

	return
}

func copyDir(src string, dst string) (err error) {
	// fmt.Println("COPY DIR", src, dst)
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("copy: destination already exists %s", dst)
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

// CopyChildren ...
func copyChildren(fromDir, toDir string) error {
	// copy the outputDir to the sites dir
	files, err := ioutil.ReadDir(fromDir)
	if err != nil {
		return err
	}

	// automatically recycle dist
	err = os.RemoveAll(toDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(toDir, 0770)
	if err != nil {
		return err
	}

	for _, f := range files {

		fullPath := filepath.Join(fromDir, f.Name())

		if f.IsDir() {
			err = copyDir(fullPath, filepath.Join(toDir, f.Name()))
		} else {
			err = copyFile(fullPath, filepath.Join(toDir, f.Name()))
		}

		if err != nil {
			return err
		}
	}

	return nil
}
