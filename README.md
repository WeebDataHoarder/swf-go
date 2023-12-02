# swf-go

Modular Shockwave Flash reader, parser and decoder in Go.

Compared to [kelvyne/swf](https://github.com/kelvyne/swf), it also decodes types and Tag contents, including sprites.

### Example

```go
package example

import (
    "errors"
    "git.gammaspectra.live/WeebDataHoarder/swf-go"
    "git.gammaspectra.live/WeebDataHoarder/swf-go/tag"
    "io"
    "os"
)

func main() {
    f, err := os.Open("flash.swf")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    reader, err := swf.NewReader(f)
    if err != nil {
        panic(err)
    }
    defer reader.Close()

    for {
        t, err := reader.Tag()
        if err != nil {
            if errors.Is(err, tag.ErrUnknownTag) {
                //unknown tag, cannot decode
                continue
            }
            if errors.Is(err, io.EOF) {
                //file is completely read
                panic("EOF reached without End tag")
            }
            panic(err)
        }

        //Handle tags
        switch t.(type) {
            case *tag.End:
            //end of file
            break
        }
    }
}
```