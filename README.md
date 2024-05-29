# ascii-art-gen

A CLI tool that can generate ASCII Art from any image.  
![demo](https://github.com/ayanami77/ascii-art-gen/assets/107479598/1246092b-11d6-4a2f-a16c-e9afbcf0ec92)

## Usage

Here is basic command.

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

## Contributions

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Submit a pull request to the original repository.

## License

MIT
