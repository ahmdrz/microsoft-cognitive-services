# microsoft-cognitive-services
Put intelligence APIs to work, Full library for microsoft cognitive services

[![Build Status](https://travis-ci.org/ahmdrz/microsoft-cognitive-services.svg?branch=master)](https://travis-ci.org/ahmdrz/microsoft-cognitive-services)

***

Tap into the power of machine learning with easy-to-use REST APIs. [Get started](https://www.microsoft.com/cognitive-services)

### List of APIs :
|Name|Link|Library info|Contains|
|----|----|----|---|
|Computer vision|[link](https://www.microsoft.com/cognitive-services/en-us/computer-vision-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#computer-vision)|True|

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
