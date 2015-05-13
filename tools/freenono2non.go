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

// Takes an input files and converts it from FreeNono's xml format to ours.
// While Go may not be one's first thought for such a task, I wanted to get
// practice with it.  And actually, their xml parser is rather easy to use.

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Nonogram struct {
	Author string `xml:"author,attr"`
	Name string `xml:"name,attr"`
	Height int `xml:"height,attr"`
	Width int `xml:"width,attr"`
	Lines []string `xml:"line"`
}

type Nonograms struct {
	Nono Nonogram `xml:"Nonograms>Nonogram"`
}

func main() {
	var reader io.Reader
	var err error
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
		reader, err = os.Open(file)
		if err != nil {
			panic(err)
		}
	} else {
		reader = os.Stdin
	}

	v := Nonograms{}
	err = xml.NewDecoder(reader).Decode(&v)
	if err != nil {
		panic(err)
	}

	// Get rid of any whitespace in lines
	for i := range v.Nono.Lines {
		v.Nono.Lines[i] = strings.Replace(v.Nono.Lines[i], " ", "", -1)
	}

	rowHints := make([]string, 0, v.Nono.Height)
	for _, line := range v.Nono.Lines {
		rowChunks := strings.Split(line, "_")
		currentRowHints := make([]string, 0)
		for _, chunk := range rowChunks {
			if l := len(chunk); l > 0 {
				currentRowHints = append(currentRowHints, strconv.Itoa(l))
			}
		}
		rowHints = append(rowHints, strings.Join(currentRowHints, ","))
	}

	colHints := make([]string, 0, v.Nono.Width)
	for x := 0; x < v.Nono.Width; x++ {
		col := ""
		for y := 0; y < v.Nono.Height; y++ {
			col += string(v.Nono.Lines[y][x])
		}
		colChunks := strings.Split(col, "_")
		currentColHints := make([]string, 0)
		for _, chunk := range colChunks {
			if l := len(chunk); l > 0 {
				currentColHints = append(currentColHints, strconv.Itoa(l))
			}
		}
		colHints = append(colHints, strings.Join(currentColHints, ","))
	}

	goal := strings.Join(v.Nono.Lines, "")
	goal = strings.Replace(goal, "_", "0", -1)
	goal = strings.Replace(goal, "x", "1", -1)

	if file != "" {
		fmt.Println(`catalogue "FreeNono ` + filepath.Base(file) + `"`)
	}
	fmt.Println(`title "` + v.Nono.Name + `"`)
	fmt.Println(`by "` + v.Nono.Author + `"`)
	fmt.Println(`width ` + strconv.Itoa(v.Nono.Width))
	fmt.Println(`height ` + strconv.Itoa(v.Nono.Height))
	fmt.Println()
	fmt.Println("rows")
	fmt.Println(strings.Join(rowHints, "\n"))
	fmt.Println()
	fmt.Println("columns")
	fmt.Println(strings.Join(colHints, "\n"))
	fmt.Println()
	fmt.Println(`goal "` + goal + `"`)
}
