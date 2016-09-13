# microsoft-cognitive-services
> Put intelligence APIs to work, Full library for microsoft cognitive services

[![Build Status](https://travis-ci.org/ahmdrz/microsoft-cognitive-services.svg?branch=master)](https://travis-ci.org/ahmdrz/microsoft-cognitive-services)
![Progress](http://progressed.io/bar/15)

***

*I will improve this library just for improve my dominance on Golang.*

<img src="https://i.imgflip.com/1ag7x6.jpg"/>

Tap into the power of machine learning with easy-to-use REST APIs. [Get started](https://www.microsoft.com/cognitive-services)

**Subscription key** : Subscription key which provides access to this API. Found in your subscriptions.

### List of APIs :
|Name|Link|Library info|Reference|Coverage percent|
|----|----|----|---|---|
|Computer vision|[link](https://www.microsoft.com/cognitive-services/en-us/computer-vision-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#computer-vision)|[![GoDoc](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/computer-vision?status.svg)](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/computer-vision)|![Progress](http://progressed.io/bar/90)|
|Emotion|[link](https://www.microsoft.com/cognitive-services/en-us/emotion-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#emotion)|[![GoDoc](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/emotion?status.svg)](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/emotion) |![Progress](http://progressed.io/bar/60)|
|Face|[link](https://www.microsoft.com/cognitive-services/en-us/face-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#face)|[![GoDoc](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/face?status.svg)](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/face) |![Progress](http://progressed.io/bar/18)|
|WebSearch|[link](https://www.microsoft.com/cognitive-services/en-us/bing-web-search-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#websearch)|[![GoDoc](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/web-search?status.svg)](https://godoc.org/github.com/ahmdrz/microsoft-cognitive-services/web-search) |![Progress](http://progressed.io/bar/60)|

***

### Computer Vision

```bash
  go get github.com/ahmdrz/microsoft-cognitive-services/computer-vision
```

Sample 

```go
// test project main.go
package main

import (
    "fmt"

    "github.com/ahmdrz/microsoft-cognitive-services/computer-vision"
)

func main() {
    fmt.Println("Hello World!")
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

    fmt.Println(result)
}
```

### Emotion

```bash
  go get github.com/ahmdrz/microsoft-cognitive-services/emotion
```

Sample 

```go
// test project main.go
package main

import (
    "fmt"

    "github.com/ahmdrz/microsoft-cognitive-services/emotion"
)

func main() {
    fmt.Println("Hello World!")
    emo, err := emotion.New("<KEY>")
    if err != nil {
        panic(err)
    }

    result, err := emo.Recognize("https://portalstoragewuprod.azureedge.net/emotion/recognition1.jpg")
    if err != nil {
        panic(err)
    }

    fmt.Println(result)
}
```

### Face

```bash
  go get github.com/ahmdrz/microsoft-cognitive-services/face
```

Sample 

```go
// test project main.go
package main

import (
    "fmt"

    "github.com/ahmdrz/microsoft-cognitive-services/face"
)

func main() {
  	fmt.Println("Hello World!")
  	facevar, err := face.New("<KEY>")
	if err != nil {
		panic(err)
	}

	results, err := facevar.Detect("https://portalstoragewuprod.azureedge.net/media/Default/Documentation/Face/Images/FaceFindSimilar.QueryFace.jpg",
		face.DetectOrder{
			FaceLandmarks: true,
			FaceAttributes: face.FaceAttributesOrder{
				Age: true,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result.FaceAttributes.Age)
	}
}
```

### License

Visit license file

### Contrib

Help me to improve it by pull requests.
