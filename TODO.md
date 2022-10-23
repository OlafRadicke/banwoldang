
- Logging cleanup
- Log level
- Create a file with statistic information
- support hard links
- Join categories doublet files
- Add doublet category
- Remove special characters in file names


```golang
// rename
e := os.Rename(Original_Path, New_Path)
if e != nil {
    cl.ErrorLogger.Fatal(e)
}
```


