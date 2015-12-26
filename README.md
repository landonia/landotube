# landotube

My blog landotube created using the [simplegoblog](https://github.com/landonia/simplegoblog).

## Overview

I have uploaded my blog to github as an example on how to create a blog using [simplegoblog](https://github.com/landonia/simplegoblog).
All changes to the blog will be committed here and available for all to see.

## Maturity

This is my first actual application written using the Go Language and therefore you should expect
some bugs.

## Installation

With a healthy Go Language installed, simply run `go get github.com/landonia/landoblog`

## Out of Box Example

With a healthy Go Language installed, simply run `go build github.com/landonia/landotube/main`

## Custom Example

	package main

	import (
		"fmt"
		"github.com/landonia/simplegoblog/blog"
	)

	func main() {
		blog := SimpleBlog.Create(directory)

		// If you want ot overwrite the NotFound page
		blog.NotFoundPage("page.html")

		// Or
		blog.NotFoundFunction(func(){})

		// Any custom handlers
		err := blog.AddCustomHandler("/custom", func() {})
		if err != nil {
			panic(err.Error())
		}

		// Start the server
		err != blog.Start()
		if err != nil {
			panic(err.Error())
		}
	}

## Future

As the blog posts are marshalled to/from json and written to disk it would make sense
to add a feature that would allow you to use a json backed data store such as mongodb.

## About

simplegoblog was written by [Landon Wainwright](http://www.landotube.com) | [GitHub](https://github.com/landonia).

Follow me on [Twitter @landoman](http://www.twitter.com/landoman)! Although I don't really tweet much tbh.
