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
	"time"
)

// Starts a new simple go blog server
func main() {

	// Define flags
	var postsdir, templatesdir, assetsdir, address string
	flag.StringVar(&postsdir, "pdir", "../posts", "the directory for storing the posts")
	flag.StringVar(&templatesdir, "tdir", "../templates", "the directory containing the templates")
	flag.StringVar(&assetsdir, "adir", "../assets", "the directory containing the assets")
	flag.StringVar(&address, "address", ":8080", "the host:port to run the blog on")
	flag.Parse()

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
			Ttl: time.Second, // Only fill a token in the bucket every second
		},

		// We want to allow a larger burst as the assets will likely be fetched in a larger chunk but once fetched it is unlikely
		// they will be requested again for a while
		AssetHandlerLimit: blog.ThrottleLimit{
			Max: 20,               // Allow an initial bucket of 20 to allow all the required assets to be delivered
			Ttl: time.Second * 10, // Only fill a token in the bucket every 10 seconds
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
