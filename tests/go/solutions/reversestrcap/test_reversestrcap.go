package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arg1 := []string{"First SMALL TesT"}
	arg2 := []string{"SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e"}
	args := [][]string{arg1, arg2}

	for i := 0; i < 15; i++ {
		args = append(args, z01.MultRandAlnum())
	}

	args = append(args, []string{""})

	for _, v := range args {
		z01.ChallengeMain(v...)
	}
	z01.ChallengeMain()
}