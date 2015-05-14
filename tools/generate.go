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

// Creates a randomly generated nongram.  May not be solvable or unique or fun.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: generate.go HEIGHT WIDTH SEED")
		os.Exit(1)
	}
	height, err := strconv.ParseUint(os.Args[1], 10, 0)
	if err != nil { fmt.Println("Could not understand height"); os.Exit(1) }
	width, err := strconv.ParseUint(os.Args[2], 10, 0)
	if err != nil { fmt.Println("Could not understand width"); os.Exit(1) }
	seed, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil { fmt.Println("Could not understand the seed"); os.Exit(1) }

	rand.Seed(seed)

	goal := ""
	for i := uint64(0); i < width * height; i++ {
		goal += strconv.Itoa(rand.Intn(2))
	}

	rowHints := make([]string, height)
	for y := uint64(0); y < height; y++ {
		row := goal[y * width:y * width + width]
		rowChunks := strings.Split(row, "0")
		var currentRowHints []string
		for _, chunk := range rowChunks {
			if l := len(chunk); l > 0 {
				currentRowHints = append(currentRowHints, strconv.Itoa(l))
			}
		}
		rowHints[y] = strings.Join(currentRowHints, ",")
	}

	colHints := make([]string, width)
	for x := uint64(0); x < width; x++ {
		col := ""
		for y := uint64(0); y < height; y++ {
			col += string(goal[y * width + x])
		}
		colChunks := strings.Split(col, "0")
		var currentColHints []string
		for _, chunk := range colChunks {
			if l := len(chunk); l > 0 {
				currentColHints = append(currentColHints, strconv.Itoa(l))
			}
		}
		colHints[x] = strings.Join(currentColHints, ",")
	}

	fmt.Println(fmt.Sprintf(`catalogue "generate.go with seed %v"`, seed))
	fmt.Println(fmt.Sprintf(`title "Random #%v"`, seed))
	fmt.Println(`by "Michael Terry"`)
	fmt.Println(`copyright "© 2015 Michael Terry"`)
	fmt.Println(`license CC-BY-SA-4.0`)
	fmt.Println(fmt.Sprintf(`height %v`, height))
	fmt.Println(fmt.Sprintf(`width %v`, width))
	fmt.Println()
	fmt.Println("rows")
	fmt.Println(strings.Join(rowHints, "\n"))
	fmt.Println()
	fmt.Println("columns")
	fmt.Println(strings.Join(colHints, "\n"))
	fmt.Println()
	fmt.Println(fmt.Sprintf(`goal "%v"`, goal))
}
