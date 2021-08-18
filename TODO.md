
- Renaming source file (checksum) 
- flag for creating empty manifests
- Flag for creating xml config file
- Tag statistics


```golang
// rename
e := os.Rename(Original_Path, New_Path)
if e != nil {
    cl.ErrorLogger.Fatal(e)
}
```


