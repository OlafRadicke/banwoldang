
- Create a file with statistic information
- support hard links
- Join categories doublet files
- Add doublet category
- Remove special characters in file names
- read out information with https://github.com/vansante/go-ffprobe/blob/v2/probedata.go and
  - Sort by Width and Height
  - Sort by Duration
  - Config switch

```golang
// rename
e := os.Rename(Original_Path, New_Path)
if e != nil {
    cl.ErrorLogger.Fatal(e)
}
```


