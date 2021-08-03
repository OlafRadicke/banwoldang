- Checksum-tree
  - By file name
  - By category count
- Renaming source file (checksum) 
- Generate category template in /home/<user>/.config/gthumb/tags.xml


```golang
e := os.Rename(Original_Path, New_Path)
if e != nil {
    log.Fatal(e)
}
```



```xml
<?xml version="1.0" encoding="UTF-8"?>
<tags version="1.0">
  <tag value="example-01"/>
  <tag value="example-02"/>
  <tag value="example-03"/>  
  <tag value="example-04"/>
</tags>

```