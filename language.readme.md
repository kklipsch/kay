# Golang 

The golang implementation.  See the [Makefile](Makefile) for dependencies and special build parameters.

This relies on a fork of a commonly used cli library available in my github.

## Concepts

There are several concepts in this solution:
- **Kaydir** the file system folder with the content to keep track of.  Inside of that folder is a hidden meta folder named .kay which is what indicates that this is a Kaydir.  Inside of the .kay folder is an index folder which contains the serialized index.
- **Chapter** the content files.  There can be multiple revisions of each year and each revision is known as a chapter.
- **Record** the meta data for an individual chapters.
- **Index** the collection of records for the kaydir.

In general there is a go file per command in addition to utility files.  So there is an add.go file that has the code for the kay add subcommand.
