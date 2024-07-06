# ascii-art-gen

A CLI tool for generating ASCII Art from any image you like.

![ascii-art-gen](https://github.com/ayanami77/ascii-art-gen/assets/107479598/31ac8872-90d0-4a59-a99a-ae4e9d6587df)

## Usage

Here is basic command. We support `.jpg (jpeg)` and `.png` image.

```
$ ascii-art-gen ./cat.png
```

You can use `--threshold` or `-t` option to use any threshold. The default threshold is 128.

```
$ ascii-art-gen -t 110 ./cat.png
```

You can use `--magnification` or `-m` option to control the size of the ascii art. The default magnification is 1.0.

```
$ ascii-art-gen -m 2.0 ./cat.png
```

## Project setup

- run app
```
$ go run main.go <image>
```

- test
```
$ make test
```

- build
```
$ make build
```

## License

MIT
