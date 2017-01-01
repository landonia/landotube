// Copyright 2013 Landon Wainwright. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Starts up the blog system using the default values
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/landonia/golog"
	"github.com/landonia/simplegoblog/blog"
)

// Starts a new simple go blog server
func main() {

	// Define flags
	var postsdir, templatesdir, assetsdir, address, loglevel string

	flag.StringVar(&postsdir, "pdir", "../posts", "the directory for storing the posts")
	flag.StringVar(&templatesdir, "tdir", "../templates", "the directory containing the templates")
	flag.StringVar(&assetsdir, "adir", "../assets", "the directory containing the assets")
	flag.StringVar(&address, "address", ":8080", "the host:port to run the blog on")
	flag.StringVar(&loglevel, "loglevel", "debug", "The log level to use")
	flag.Parse()
	golog.LogLevel(loglevel)

	// Create a new configuration containing the info
	config := &blog.Configuration{
		Title:           "Life thru a Lando",
		Postsdir:        postsdir,
		Templatesdir:    templatesdir,
		Assetsdir:       assetsdir,
		NoOfRecentPosts: 4,

		// We want to allow a short burst of 10 requests and otherwise it will allow 1 request/sec
		RequestHandlerLimit: blog.ThrottleLimit{
			Max: 10,          // Allow a max bucket amount of 10 requests
			TTL: time.Second, // Only fill a token in the bucket every second
		},
	}

	// Create a new data structure for storing the data
	b := blog.New(config)

	// Start the blog server
	err := b.Start(address)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
