# prp

prp (Print Relative Path) command to generate relative path between two absolute paths.
Sometimes import paths can get very long and annoying to figure out manually.
Use prp when you don't want to figure out and manually write the relative path to some file that 
ends up being 4 directories back then 3 forward.

## Installation
```
$ brew tap utk9/homebrew-prp
$ brew install prp
```

### Usage
You can give `prp` either one or two arguments. If two are provided, it uses the
first as `from` and the second as `to`. If one if provided, it uses the
given path as `to` and the current working directory (from `pwd`) as `from`.

### Example

```
/dir1/
	dir2/
		dir3/
			dir4/
				file1.ext
		dir5/
			file2.ext
	dir6/
		dir7/
			file3.ext

/dir1/dir2/dir3/dir4 $ prp /dir1/dir6/dir7/file.ext
./../../../dir6/dir7/file3.ext

~ $ prp /dir1/dir2/fir5/ /dir1/dir6/dir7/file.ext
./../../dir6/dir7/file.ext
```