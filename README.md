Banwoldang
==========

A helper tool for [gThumb](https://wiki.gnome.org/Apps/Gthumb). Banwoldang is 
only a random name.

concet / idea
-------------

The media browser gThumb stores tags in xml files without database. A media 
file has an XML file with meta information (e.g. tags). The advantage is that 
the solution is very robust and simple. If meta data is stored in a database 
and the directories are moved, the links in the database are broken and the 
meta data (like tags) are lost. This does not happen with gThumb. You can 
also read and write the meta data in the XML file quite well without causing 
major damage. 

But a disadvantage is the slow search function. Because gThumb has to open 
and read every meta file during a search query. I had the idea to write a 
program that reads all meta data and creates a shadow directory with links. 
Each tag is a link in this directory tree. You can then move very fast with 
gThumb through this directory tree, because gThumb makes no difference 
between Sof Links umd Hard Links. The XML files with the meta data are also 
created in this link directory tree. So you can also edit the tags persistent 
in this tree with gThumb.

Build
-----

```
go build
```

Run
---

```bash
go run ./main.go  -f example/config.yaml
```

or

```bash
./banwoldang  -f  example/config.yaml

```

Flags
-----

* ***-f <filename>*** The path of the configuration file

Configuration
-------------

You find an example in [example/config.yaml](example/config.yaml)

ffmpeg support
--------------

For ffmpeg support, install ffmpeg and add in your configuration:

```yaml
ffmpeg_support: true

```

