# kay - A command line tool to keep track of autobiography chapters organized by year.

## Preamble

Every year for most of my adult life, my grandmother has given me (and all her other grandchildren) a chapter of her autobiography.  The chapter is organized by age, so for example, on your 25th birthday you receive the chapter of her story from her own 25th year.  She does not fully rewrite the chapter if it has already been written but she does occasionally edit details or add small pieces.  Further, she did not start this with her first year and only writes the chapters as they are needed.  For example, if she has grandchildren hitting the ages 34, 28, 19 and 7 in a year, but she has years 28 and 19 already written, that year she will add years 34 and 7 along with her edits of years 28 and 19.  This has led to an autobiography that is delivered out of order and has gaps in it, as well as there being many different versions of each story.  In addition to the stories she sends, she frequently includes pictures or other mementos related to her autobiography that add valuable context to the stories she tells.

I started collecting these stories in order to organize them and keep a central repository for myself and anyone else who wanted to access the stories.  Typically, they are delivered in the mail on or near your birthday.  They are standard printouts with hand written edits or personal notes.  Over the years, I would also occasionally receive printouts or xerox copies of the versions of the chapters she had prepared for her other grandchildren through the year.  I scan all of the chapters I receive and save them.  Also, every once in awhile someone will take a dump of the source files from my grandmother's computer.  This usually is then delivered to me on a thumb-drive or Dropbox share with no organization.  I therefore have collected a hodge-podge of .doc, .docx, .txt, .rtf, .txt, .jpg, .jpeg, .png and .pdf chapters and appendixes.  Answering questions about who received what story, when it was written and what years are missing has become an unmanageable task.

Further, I publish a website for others to be able to access.  Organizing this website has been a manual, tedious and error prone process. Due to this the website is usually out of date, is organized poorly, and doesn't receive much attention from a visual perspective.

Finally, at some point I would like to investigate creating a 'vanity' book out of these chapters so being able to organize and make sense of them quickly will be a boon.

## What I Want to Achieve Part I: Make the task of maintaining this autobiography easy.

In general I want to be able to do the following (and more):

- Take a large directory of files with a few different naming conventions and catalog the files by year, and if possible by "tag".  Tags should be anything, but are envisioned to be used to manage grandchildren names.  It should be assumed that some of these files have already been cataloged and duplicate catalog details should not be a problem.
- Determine what files have not been cataloged.
- Determine what years are represented in a catalog and what years are missing.
- Determine what tags are represented in a catalog.
- Find files by year and tag.
- Allow for annotations of files.
- Allow for the existence of appendix style entries, with tags & notes.
- Enable easy comparison of different versions of a document from the same year.
- Generate a **static** website.
- **Bonus:** Make the website with templates. 
- **Bonus:** Preparations for the 'vanity' book.
- **Bonus:** Allow for multiple platform execution (Mac, Linux & Windows).

What the kay program specifically should not do:

- VCS style work.  Version changes, modification detection etc. is best done by a VCS and therefore will be done by one.  Nothing in the design of the program should interfere with this though, and in fact, preference should be given to solutions that are more easily managed via a VCS.
- Publishing of the website.  Hosting parameters might change over the lifetime of this project, so we should not build host specific workflows in the tool.  The design of the program should make it easy to automate with external tools the publication of the project.

The kay program is envisioned to be a command line utility with git-like subcommands.  While not a requirement, it is assumed that the catalog will exist as human readable plain-text files storeable in a VCS.  Further, any command line utility is made easier to use with command completion so bash support of command completion is a **strong** positive indicator.

## What I Want to Achieve Part II: Use this as a base problem set when learning new programming languages.

This program is typically the first thing I try to write in a new language after 'Hello World'.  There have been aborted attempts written in Haskell and an early version of Rust.  The time has come to actually finish the program though, so the priority now is on completion.  That said, this remains an interesting first problem in a language as it is well defined, constrained but has several interesting properties (argument parsing, file io, opportunities for types etc.).  Therefore this repo will be organized in such a way as to make multiple versions in multiple different languages easy.

### Branching Model.

The master branch will be any language neutral documentation, utilities and the starting point for each new language.  To start a new language ought to be as simple as making a new branch with that languages name from the HEAD of the master branch.  Further, each branch should be organized in such a way that merging changes from master is easy, even if the actual language specific implementations are widely divergent. The language that is currently being actively developed should be set as the 'default' branch.  Language specific documentation should be found in the [language.readme.md](language.readme.md) file.

### The integration tests.

In the previous uncompleted implementations all of the testing was done in the idiomatic unit test tools of the language.  When I abandoned these implementations I also lost all of the built in tests.  In order to avoid this in the future a suite of integration tests in the form of an ansible playbook have been provided in the [integration](integration) folder.

### Caveats about the languages represented.

The languages represented in this repo represent the *first* attempt at writing a project in the language for me.  Therefore they are frequently the **worst** possible representation of that language. The code is not expected to be idiomatic, clever, or particularly good.  It is expected to be clear, working and tested.  I choose the languages based on my interest, not necessarily on it being the best language for the job (this project is probably trivial in Python for instance).  Also, the problem itself will not necessarily show the best characteristics for all languages.  For instance it is largely IO driven so pure languages might not make sense.  It needn't have much concurrency so languages optimized for that won't shine and it isn't big enough to truly benefit from a robust type system.  Also the nature of the solution, a small quickly run command line app will severely penalize some platforms (the JVM jumps to mind).
