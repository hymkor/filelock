filelock
========

`filelock` is a program that keeps the file specified by its first argument in use until EOF is encountered on standard input.
it was created for testing file locking.

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/filelock/releases) and extract the executable.

### for scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/filelock/master/filelock.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install filelock
```
