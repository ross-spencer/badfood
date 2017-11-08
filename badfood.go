/*Original code from Richard Lehane github.com/richardlehane*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	version = "v0.0.2-nukem"
	vers    bool
	fname   string
	badfood bool
	nukem   bool
	eicar   bool
	size    float64 // megabytes
)

func init() {
	flag.Float64Var(&size, "size", 0.0, "size of random file to create (mb).")
	flag.StringVar(&fname, "fname", "samplefile", "filename of object to create.")
	flag.BoolVar(&badfood, "badfood", false, "output Artefactual test virus in middle of stream")
	flag.BoolVar(&nukem, "nukem", false, "output Nukem virus header in middle of stream")
	flag.BoolVar(&eicar, "eicar", false, "Output Eicar virus at beginning of file")
	flag.BoolVar(&vers, "version", false, "Return version.")
}

const artefactualBadfood = "\x00\xAF\xBA\xDF\x00\xD0"

//Win.Trojan.Nukem-1:0:*:
const artefactualNukem = "\x4e\x75\x6b\x65\x5f\x65\x4d\x00\x00\x90\x90\x90\xe8\x78\x26\x00\x00\x84\xd2\x84\xd2\x7e\x05\xe8\x60\x26\x00\x00\xc3\x90\x90\x90\x55\x8b\xec\x83\xc4\xd8\x53\xe8\x5d\x26\x00\x00\x8b\xda\x89\x45"

//Eicar-Test-Signature
const artefactualEicar = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"

func createFile() {
	if size > 0 && fname != "" {
		sz := int(size * 1024 * 1024)
		bigBuff1 := make([]byte, sz/2)
		bigBuff2 := []byte("")
		if badfood {
			bigBuff2 = []byte(artefactualBadfood)
		} else if nukem {
			bigBuff2 = []byte(artefactualNukem)
		} else if eicar {
			bigBuff2 = []byte(artefactualEicar)
		}

		var bin []byte
		if eicar {
			bin = append(bigBuff2[:], bigBuff1[:]...)
		} else {
			bin = append(bigBuff1[:], bigBuff2[:]...)
		}
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
