/*
 Copyright 2021 Scott Nichols
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"tableflip.dev/ggsv/pkg/git"
)

func main() {
	if len(os.Args) != 4 || os.Args[1] != "next-patch" {
		fmt.Fprintf(os.Stderr, "usage:\n ggsv next-patch <git url> <git-ref|version|branch>\n")
		os.Exit(1)
	}

	// TODO: update this to a better cli experience.
	remote := os.Args[2]
	ref := os.Args[3]

	repo, err := git.GetRepo("HEAD", remote)
	if err != nil {
		panic(err)
	}

	next, err := repo.Next(ref)
	fmt.Printf("v%s", next)
}
