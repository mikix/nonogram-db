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

// Searches our database tree and makes sure all files meet our quality
// standards, both in term of puzzle quality and licensing.
// Make sure pbnsolve is in $PATH.

// Go is probably not the best choice for this, but I wanted to get more
// practice with the language.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var toolspath string
var done chan bool
var count int

func checkIt(path string) (status bool) {
	defer func() {done <- status}()
	if filepath.Ext(path) != ".non" {
		return true
	}

	// There are a few checks we want to make.
	// First, is this even a valid puzzle?  We use addgoal for that.
	// Second, are all the metadata we insist upon present?

	cmd := exec.Command(filepath.Join(toolspath, "addgoal"), path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(string(output))
		return false
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	content := string(data)

	if ok, _ := regexp.MatchString(`(?m:^)title ".+?"(?m:$)`, content); !ok {
		fmt.Println(path, "is missing a valid title")
		return false
	}
	if ok, _ := regexp.MatchString(`(?m:^)by ".+?"(?m:$)`, content); !ok {
		fmt.Println(path, "is missing a valid author")
		return false
	}
	if ok, _ := regexp.MatchString(`(?m:^)copyright ".+?"(?m:$)`, content); !ok {
		fmt.Println(path, "is missing a valid copyright")
		return false
	}
	// only ship standard SPDX license (i.e. make sure we don't have quotes)
	if ok, _ := regexp.MatchString(`(?m:^)license [^"]+?(?m:$)`, content); !ok {
		fmt.Println(path, "is missing a valid license")
		return false
	}

	// Everything else was taken care of by addgoal

	return true
}

func walkIt(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Parallelism is not concurrency, yeah.  But still.
	count++
	go checkIt(path)
	return nil
}

func main() {
	execpath, err := filepath.Abs(os.Args[0])
	if err != nil { panic(err) }

	if _, err = exec.LookPath("pbnsolve"); err != nil {
		fmt.Println("You forgot to set the PATH for pbnsolve")
		os.Exit(1)
	}

	toolspath = filepath.Join(execpath, "..")
	dbpath := filepath.Join(toolspath, "../db")
	done = make(chan bool, 100)

	err = filepath.Walk(dbpath, walkIt)
	if err != nil { panic(err) }

	// Wait for goroutines
	success := true
	for i := 0; i < count; i++ {
		if !<-done {
			success = false
		}
	}
	if !success {
		os.Exit(1)
	}
}
