// Copyright 2016 The ahmdrz.
// license that can be found in the LICENSE file

/*
Package vision is a library for micrsoft vision api

Let's start to make a request to bing vision
Full documention is on https://dev.projectoxford.ai/docs/services/56f91f2d778daf23d8ec6739/

Installation : go get github.com/ahmdrz/microsoft-vision-golang

First of all , let me to explain how to use image from url.


Tag method :

	func main() {
		vision, err := vision.New("<KEY>")
	    if err != nil {
	        panic(err)
	    }

	    result, err := vision.Tag("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println(result)
	}

Analyze method :

	func main() {
		vis, err := vision.New("<KEY>")
	    if err != nil {
	        panic(err)
	    }

	    result, err := vis.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg",
	        vision.VisualFeatures{
	            Adult: true,
	            Tags:  true,
	        })
	    if err != nil {
	        panic(err)
	    }
	}


------

In analyze method we have VisualFeatures , Source of VisualFeatures struct is
	type VisualFeatures struct {
		Categories  bool
		Tags        bool
		Description bool
		Faces       bool
		ImageType   bool
		Color       bool
		Adult       bool
	}
... and func (order VisualFeatures) String() (string, error) can make a string from this struct

*/
package vision
