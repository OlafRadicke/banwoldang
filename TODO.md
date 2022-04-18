- support hard links
- Join categories doublet files
- Add doublet category


```golang
// rename
e := os.Rename(Original_Path, New_Path)
if e != nil {
    cl.ErrorLogger.Fatal(e)
}
```


