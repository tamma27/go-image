#

## 1. Description

API:

- Png2jpg

- Webp2jpg

- Watermar

## 2. Example

```go
func main() {
  Png2jpg("./image/image.png", "./image.png.jpg")
  Watermar("./image.png.jpg")

  Webp2jpg("./image/image.webp", "./image.webp.jpg")
  Watermar("./image.webp.jpg")

  Watermar("./image/image.jpg")
}
```
