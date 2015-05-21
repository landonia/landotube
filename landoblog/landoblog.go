// Copyright 2013 Landon Wainwright. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Starts up the blog system using the default values
package main

import (
	"flag"
	"github.com/landonia/simplegoblog/blog"
	"log"
	"os"
)

// Starts a new simple go blog server
func main() {

	// Define flags
	var postsdir, templatesdir, assetsdir, port string
	flag.StringVar(&postsdir, "pdir", "../posts", "the directory for storing the posts")
	flag.StringVar(&templatesdir, "tdir", "../templates", "the directory containing the templates")
	flag.StringVar(&assetsdir, "adir", "../assets", "the directory containing the assets")
	flag.StringVar(&port, "port", "8080", "the port to run the blog on")
	flag.Parse()

	// Create a new configuration containing the info
	config := &blog.Configuration{Title: "Life thru a Lando", DevelopmentMode: true, Postsdir: postsdir, Templatesdir: templatesdir, Assetsdir: assetsdir}

	// Create a new data structure for storing the data
	b := blog.New(config)

	// Start the blog server
	err := b.Start(port)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
