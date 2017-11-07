/*Original code from Richard Lehane github.com/richardlehane*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	version = "v0.0.1-beta"
	vers    bool
	fname   string
	badfood bool

	size float64 // megabytes
)

func init() {
	flag.Float64Var(&size, "size", 0.0, "size of random file to create (mb).")
	flag.StringVar(&fname, "fname", "samplefile", "filename of object to create.")
	flag.BoolVar(&badfood, "badfood", false, "output Artefactual test virus in middle of stream")
	flag.BoolVar(&vers, "version", false, "Return version.")
}

const artefactualBadfood = "\x00\xAF\xBA\xDF\x00\xD0"

func createFile() {
	if size > 0 && fname != "" {
		sz := int(size * 1024 * 1024)
		bigBuff1 := make([]byte, sz/2)
		bigBuff2 := []byte("")
		if badfood {
			bigBuff2 = []byte(artefactualBadfood)
		}
		bin := append(bigBuff1[:], bigBuff2[:]...)
		bin = append(bin[:], bigBuff1[:]...)
		ioutil.WriteFile(fname, bin, 0666)
		fmt.Println("Outputting", fname)
	}
}

func main() {

	flag.Parse()
	var verstring = "badfood "
	if vers {
		fmt.Fprintf(os.Stderr, "%s %s \n", verstring, version)
		os.Exit(0)
	} else if flag.NFlag() < 2 { // can access args w/ len(os.Args[1:]) too
		fmt.Fprintln(os.Stderr, "Usage:  badfood [-size ...] [-fname ...]")
		fmt.Fprintln(os.Stderr, "                  OPTIONAL: [-badfood]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Output: [BIN] {fname}")
		fmt.Fprintf(os.Stderr, "Output: [STRING] '%s ...'\n\n", verstring)
		flag.Usage()
		os.Exit(0)
	} else {
		createFile()
	}
}
