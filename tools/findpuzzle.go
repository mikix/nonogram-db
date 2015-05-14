///bin/true; exec /usr/bin/env go run "$0" "$@"
// -*- Mode: Go; indent-tabs-mode: t; tab-width: 8 -*-
/*
 * Copyright 2015 Michael Terry
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	if _, err := exec.LookPath("pbnsolve"); err != nil {
		fmt.Println("You forgot to set the PATH for pbnsolve")
		os.Exit(1)
	}

	n := flag.Uint("n", 1, "Number of puzzles to generate")
	d := flag.String("d", "1+", "Difficulty of puzzles to generate")
	flag.Parse()

	if len(flag.Args()) != 3 {
		fmt.Println(flag.Args())
		fmt.Println("Usage: findpuzzle.go HEIGHT WIDTH SEED")
		os.Exit(1)
	}
	height, err := strconv.ParseUint(flag.Arg(0), 10, 0)
	if err != nil { fmt.Println("Could not understand height"); os.Exit(1) }
	width, err := strconv.ParseUint(flag.Arg(1), 10, 0)
	if err != nil { fmt.Println("Could not understand width"); os.Exit(1) }
	seed, err := strconv.ParseInt(flag.Arg(2), 10, 64)
	if err != nil { fmt.Println("Could not understand the seed"); os.Exit(1) }

	minDifficulty := 1
	maxDifficulty := 0

	if (*d)[len(*d) - 1] == '+' {
		minDifficulty, _ = strconv.Atoi((*d)[:len(*d) - 1])
	} else {
		minDifficulty, _ = strconv.Atoi(*d)
		maxDifficulty = minDifficulty
	}

	for count := uint(0); count < *n; {
		cmd := exec.Command("go", "run", "./generate.go", fmt.Sprintf("%v", height), fmt.Sprintf("%v", width), fmt.Sprintf("%v", seed))
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		file, _ := ioutil.TempFile(".", "")
		file.Write(output)
		file.Close()

		cmd = exec.Command("pbnsolve", "-u", "-aLHEC", "-t", "-f", "non", file.Name())
		output, err = cmd.Output()
		if err != nil {
			panic(err)
		}

		if ok, _ := regexp.Match(".*UNIQUE.*", output); ok {
			diffexp, _ := regexp.Compile("\nLines Processed:.* .((?m:.*))00%")
			matches := diffexp.FindSubmatch(output)
			difficulty, _ := strconv.Atoi(string(matches[1]))

			if difficulty >= minDifficulty && (maxDifficulty == 0 || difficulty <= maxDifficulty) {
				newfile := fmt.Sprintf("%vx%v.%v.%v.non", width, height, difficulty, seed)
				os.Rename(file.Name(), newfile)
				fmt.Println("Wrote", newfile)
				count++
			}
		}
		seed++

		os.Remove(file.Name())
	}
}
