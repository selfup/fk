# Fahrenheit and Kelvin (fk)

Windows (F), Linux (K), and macOS (C)

Tiny shell commands that work the same everywhere!

```shell
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk mkdir foo/bar

selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk wr foo/bar/example.txt "hello world" 0775

selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk fex "foo/bar/example.txt"

selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk rd "foo/bar/example.txt"
hello world
```
