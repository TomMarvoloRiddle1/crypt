package zipping

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func ZipTargetFolder(originFolder string, exportedZip string) {

	//creates zip
	file, err := os.Create(exportedZip)
	if err != nil {
		fmt.Println("zip creation error 1")
		panic(err)
	}
	//idiomatic data leak prevention, defer moves schedules before main() returns
	defer file.Close()

	//begin zip recursion
	writer := zip.NewWriter(file)
	// func() = "anonymous function" used as a sub codeblock inside main(), CREATED ON THE SPOT for internal use (as a "sub func")
	defer func() {
		//similar to file.Close() to do IO operations, handles the errors within codeblock
		if err := writer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing zip writer: %v\n", err)
		}
		// () = calling anon function, scheduled due to defer
	}()

	// THIS is the important part for WalkDir, as its the custom logic that handles what is done with recursively returned Dir data
	walker := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accessing path %q: %v\n", path, err)
			return err
		}

		//true = directory, no need to write the data
		if d.IsDir() {
			return nil
		}

		//if != Dir, its a file, so the data is returned as srcFile
		srcFile, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %q: %v\n", path, err)
			return err
		}
		defer srcFile.Close()

		// determines relativepath from originFolder for subDirectories and returns string
		relativePath, err := filepath.Rel(originFolder, path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting relative path for %q: %v\n", path, err)
			return err
		}
		//sets value of relativePath in correspondance to root ./
		relativePath = filepath.ToSlash(relativePath)

		// checks relative path for files, and generates *io.Copy "entry"; prepping for writing
		// writer is for new .zip, this is like 1/2 of os.Create
		zipEntry, err := writer.Create(relativePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating entry in zip for %q (as %q): %v\n", path, relativePath, err)
			return err
		}

		//copy file into zip entry
		// actually writes the "entry" prepped data, like 2/2 of os.Create
		_, err = io.Copy(zipEntry, srcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error copying content from %q to zip entry %q: %v\n", path, relativePath, err)
			return err
		}

		return nil
	}

	//the recursive file walking itself
	err = filepath.WalkDir(originFolder, walker)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking the path %q: %v\n", originFolder, err)
		// panic(err) // Or handle more gracefully
	} else {
		fmt.Printf("Successfully created %s from %s\n", exportedZip, originFolder)
	}
}
