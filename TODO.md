TODOs
-----


- Refactoring Configuration Object and his using
- Logging cleanup
- Log level
- Create a file with statistic information
- support hard links
- Join categories doublet files
- Add doublet category
- Remove special characters in file names
- Make in ***mediainformation/ExtractFileNameParts.go**** the ***"if len(listOfParts[i]) < 4 {"*** mor flexible over configuration.
- Add more statistic
- Godoc

```golang
// rename
e := os.Rename(Original_Path, New_Path)
if e != nil {
    cl.ErrorLogger.Fatal(e)
}
```


