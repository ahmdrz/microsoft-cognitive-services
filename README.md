# microsoft-cognitive-services
Put intelligence APIs to work, Full library for microsoft cognitive services

[![Build Status](https://travis-ci.org/ahmdrz/microsoft-cognitive-services.svg?branch=master)](https://travis-ci.org/ahmdrz/microsoft-cognitive-services)

***

Tap into the power of machine learning with easy-to-use REST APIs. [Get started](https://www.microsoft.com/cognitive-services)

### List of APIs :
|Name|Link|Library info|Contains|Coverage percent|
|----|----|----|---|---|
|Computer vision|[link](https://www.microsoft.com/cognitive-services/en-us/computer-vision-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#computer-vision)|True|90%|
|Emotion|[link](https://www.microsoft.com/cognitive-services/en-us/emotion-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#emotion)|True|60%|
|Face|[link](https://www.microsoft.com/cognitive-services/en-us/face-api)|[link](https://www.microsoft.com/cognitive-services/en-us/emotion-api)|[see](https://github.com/ahmdrz/microsoft-cognitive-services#face)|True|10%|

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
  \\ please wait ... :smile:
```

### License

Visit license file

### Contrib

Help me to improve it by pull requests.
