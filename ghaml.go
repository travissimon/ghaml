package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"unicode"
	"unicode/utf8"
)

type ghamlConfig struct {
	goBuildAfter bool
	forceCompile bool
	verbose      bool
	clean        bool
	workingDir   string
}

func main() {
	var nogo = flag.Bool("nogo", false, "prevent running the 'go build' command")
	var f = flag.Bool("f", false, "force compilation of all haml files")
	var v = flag.Bool("v", false, "prints the name of the files as they are compiled")
	var clean = flag.Bool("clean", false, "cleans generated *.go files")
	flag.Parse()

	cfg := &ghamlConfig{
		goBuildAfter: *nogo == false,
		forceCompile: *f,
		verbose:      *v,
		clean:        *clean,
	}

	wdStr, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	if cfg.clean {
		filepath.Walk(wdStr, makeWalkFunc(checkFileForDeletion, cfg))
		return
	}

	// create closure to pass our config into a WalkFunc
	filepath.Walk(wdStr, makeWalkFunc(checkFileForCompilation, cfg))

	if cfg.goBuildAfter {
		runGoBuild(cfg)
	}
}

// Runs 'go build' on the application.
// This is a convenience so that the user
// doesn't have to repeatedly execute two commands
func runGoBuild(cfg *ghamlConfig) {
	if cfg.verbose {
		fmt.Println("Executing 'go build' command")
	}
	goCmd := exec.Command("go")
	goCmd.Args = append(goCmd.Args, "build")
	goCmd.Dir = cfg.workingDir

	output, err := goCmd.Output()
	if len(output) > 0 {
		fmt.Printf("%s\n", output)
	}
	if err != nil {
		fmt.Println(err)
	}
}

// create a closure to pass our config into a WalkFunc
func makeWalkFunc(fn func(string, os.FileInfo, error, *ghamlConfig) error, cfg *ghamlConfig) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		return fn(path, f, err, cfg)
	}
}

// Removes all generated files from the working directory.
func checkFileForDeletion(path string, f os.FileInfo, err error, cfg *ghamlConfig) error {
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	if isHaml(f.Name()) {
		goFilename := getGoFilename(path, f)
		dir := filepath.Dir(path)
		goFilePath := filepath.Join(dir, goFilename)
		exists, err := exists(goFilePath)
		if err != nil {
			panic(err)
		}
		if cfg.verbose {
			if exists {
				fmt.Printf("Removing file: %s\n", goFilename)
			} else {
				fmt.Printf("No corresponding go file found for: %s\n", f.Name())
			}
		}
		if exists {
			os.Remove(goFilePath)
		}
	}
	return nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Checks a file to see if it is a haml file, and if it needs to be compiled
func checkFileForCompilation(path string, f os.FileInfo, err error, cfg *ghamlConfig) error {
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	if isHaml(f.Name()) {
		needsCompiling := doesFileNeedsCompiling(path, f)
		compile := cfg.forceCompile || needsCompiling // are we forcing compiling of all files?
		if cfg.verbose {
			if compile {
				fmt.Printf("Compiling: %s\n", f.Name())
			} else {
				fmt.Printf("%s is up to date - skipping\n", f.Name())
			}
		}
		if compile {
			return compileFile(path, f)
		}
	}
	return nil
}

// Utility method to check if this is a haml file (*.haml)
func isHaml(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".haml"
}

// lets us know if we need to compile a haml file
// Files need compiling if there is no *.go file, or if it is
// older than the corresponding .haml file
func doesFileNeedsCompiling(filepath string, f os.FileInfo) bool {
	goFilename := getGoFilename(filepath, f)
	goFileExists, err := exists(goFilename)

	if err != nil {
		panic(err)
	}

	// simple condition - 
	if !goFileExists {
		return true
	}

	goFileInfo, _ := os.Stat(filepath)
	// check to see if haml file is newer
	return f.ModTime().Before(goFileInfo.ModTime())
}

// Compiles a haml file into a go file
func compileFile(path string, f os.FileInfo) error {
	hamlFile, err := os.Open(path)
	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(hamlFile)
	if err != nil {
		return err
	}

	if err := hamlFile.Close(); err != nil {
		return err
	}

	parser := NewParser(f.Name(), string(contents))
	parser.Parse()

	goFileStr := getGoFilename(path, f)

	dir := filepath.Dir(path)
	goFile, err := os.Create(filepath.Join(dir, goFileStr))
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(goFile)
	defer func() {
		writer.Flush()
		if err := goFile.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	ext := filepath.Ext(f.Name())
	rootNameStr := getProperCase(TrimSuffix(f.Name(), ext))
	viewWriter := NewViewWriter(writer, parser.context, parser.root, rootNameStr)
	viewWriter.WriteView()

	return nil
}

// Utility function to get the go version of a haml filename
func getGoFilename(path string, f os.FileInfo) string {
	ext := filepath.Ext(f.Name())
	rootNameStr := TrimSuffix(f.Name(), ext)
	return rootNameStr + ".go"
}

// from Go 1.1 sources
// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string {
	if HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Gets a proper case for a string (i.e. capitalised like a name)
func getProperCase(s string) string {
	// Generously 'borrowing' from strings.Map

	// In the worst case, the string can grow when mapped, making
	// things unpleasant.  But it's so rare we barge in assuming it's
	// fine.  It could also shrink but that falls out naturally.
	maxbytes := len(s) // length of b
	nbytes := 0        // number of bytes encoded in b
	var b []byte = make([]byte, maxbytes)

	needsCapitalising := true

	for _, c := range s {
		r := c

		// chars to skip, causing a capitalisation
		if r == '.' || r == '-' || r == '_' || r == ' ' {
			needsCapitalising = true
			continue
		}

		if needsCapitalising {
			r = unicode.ToTitle(r)
			needsCapitalising = false
		}

		if r >= 0 {
			wid := 1
			if r >= utf8.RuneSelf {
				wid = utf8.RuneLen(r)
			}
			if nbytes+wid > maxbytes {
				// Grow the buffer.
				maxbytes = maxbytes*2 + utf8.UTFMax
				nb := make([]byte, maxbytes)
				copy(nb, b[0:nbytes])
				b = nb
			}
			nbytes += utf8.EncodeRune(b[nbytes:maxbytes], r)
		}
	}
	return string(b[0:nbytes])
}
