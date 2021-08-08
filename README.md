Banwoldang
==========

A helper tool for [gThumb](https://wiki.gnome.org/Apps/Gthumb)

Build
-----

```
go build
```

Run
---

```bash
go run ./main.go  --source-dir=/your/dir  --link-dir=/dir/for/link-tree
```

or

```bash
./banwoldang  --source-dir=/your/dir  --link-dir=/dir/for/link-tree

```

### Flags

* ***--checksum=true*** Creates checksums from the media fils CONTENT and us it as file name for the destination links. This ist god to find duplicates. 
* ***--checksum=false*** Creates checksums from the media fils NAMES and us it as file name for the destination links.



- [rename files](https://www.geeksforgeeks.org/how-to-rename-and-move-a-file-in-golang/)
- [Parsing XML Files](https://tutorialedge.net/golang/parsing-xml-with-golang/)
- [List all files (recursively) in a directory](https://yourbasic.org/golang/list-files-in-directory/)