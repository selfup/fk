# Fahrenheit and Kelvin (fk)

Windows (F), Linux (K), and macOS (C)

Tiny shell commands that work the same everywhere!

```shell
(mkdir)
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk mkdir foo/bar

(write file)
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk wr foo/bar/example.txt "hello world" 0775

(file exists)
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk fex "foo/bar/example.txt"

(file read)
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk rd "foo/bar/example.txt"
hello world

(list dir contents)
selfup@win42 MINGW64 ~/go/src/github.com/selfup/fk (master)
$ fk ls .
d: .git
f: .gitignore
f: LICENSE
f: README.md
f: main.go
```
