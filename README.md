Banwoldang
==========

A helper tool for [gThumb](https://wiki.gnome.org/Apps/Gthumb). Banwoldang is
only a random name.

[[_TOC_]]

concept / idea
--------------

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

Content of the link tree
------------------------

### 00-cat-count

Counts the number of tags in the meta files. This is useful when you want to
find files that have no or few tags.

### 00-checksum

Here a checksum is formed for each found file and used as link address.
If there are several files with the same content, they get the same name.
In case of duplicates this will fail. This is helpful to find duplicates.
After the run you can search the log (./logs/duplicates.log) for
corresponding entries to identify the duplicates.

If the data is very large, it takes a long time to create the checksums.
The creation of (real) checksums can be disabled in the configuration with:

```yaml
real_checksum: false
```

### 00-old-name-parts

The filenames of the media files are split into their parts and this is
grouped into directories. For example two files with the following names:

- Tom-2002-in_berlin_shoppen.jpg
- 1985/07-Cathy-and-celine@holiday.jpg

Result in the following directories being created:

- 2002
- BERLIN
- SHOPPING
- 1985
- CATHY
- CELINE
- HOLIDAY

"Tom", "07" and "in" are discarded because they are too short.

This is useful to suggest useful tags and to tag several files at once.

### categories

For each tag found a category was created.


### duration

Video movies are looked how long they are. For each duration there is a
separate directory. It is always rounded up to full minutes.

### resolution

If the media is a movie, it will check what resolution it has and set each for

- height
- width

A separate directory is generated

### Statistics

Statistics are also created during the run. The following files are located in
the base directory of the link directory:

- statistics_duration.csv
- statistics.md
- statistics_name-parts.csv
- statistics_resolution_height.csv
- statistics_resolution_width.csv
- statistics_tag.csv

Dos and don'ts
--------------

The tree with the links is not meant to be edited. The metadata of the media
files can be viewed and also edited with gthumb. If you run banwoldang again,
the link tree will not be synchronized and obsolete links will not be deleted.
Instead, the entire link tree should be deleted and recreated. The only
***"place of truth"*** is and remains the directory with the media files and
their original comment files.

Install requirements
--------------------

enter:

```bash
go get
```
Build
-----

```
go build
```

Run
---

```bash
go run ./banwoldang.go  -f example/config.yaml
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

Logging
-------

This log files are written:

- ***./logs/error.log***: error outputs
- ***./logs/info.log***: (Dibugging) information
- ***./logs/orphan-comments.log***: comment files for which no media file was found

Helper scrpt
------------

Enter

```bash
./scripts/empty_logs.sh
```

...to clean up the logs