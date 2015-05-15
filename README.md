## Nonograms

Nonograms (or griddlers, paint by numbers, picross, hanjie, or a million other names) are picture logic puzzles in which cells in a grid must be colored or left blank according to numbers at the side of the grid to reveal a hidden picture.

It's a popular type of puzzle, and most platforms have several apps that let users solve and/or create nonograms.

## Licenses

However, nonogram puzzles that are freely distributable are actually hard to find.  Many web sites let users create nonograms, but few (none?) allow those nonograms to be placed under any sort of license that allows redistribution.

The few open source nonogram programs that do ship with puzzles either have a paltry amount by default or ship clearly dubious-in-origin puzzles (i.e. ripped straight from a commercial game).

This database is an attempt to improve on the status quo.

## Database

All puzzle files can be found in the `db` directory.  Related puzzles will be collected in subdirectories, along with relevant information in README.md files.

## Quality

This database will only ship puzzles that have a unique solution that can be solved via logic.

Additionally, it is desirable that the puzzle reveal a picture and have a name that does not obviously give away what the picture is.

## Random Puzzles

There is also a generator included in the `tools` folder.  To find the first three solvable 5x5 nonograms that are rated difficulty 4 or higher, use the following command (with an install of Jan Wolter's pbnsolve at `~/pbnsolve`):

    cd tools
    PATH=~/pbnsolve:$PATH ./findpuzzle.go -d 4+ -n 3 5 5 0

This won't generate a pretty picture.  But they are still good logic puzzles.

## Format

An additional complication for writing a nonogram app is that there are many many formats ([at least 26](http://webpbn.com/export.cgi)) for puzzle files.  None are particularly bizarre or innovative.  They are all just different.

This database uses one format.  It's an existing format (Steve Simpson's `non` format) extended slightly (to support license information and multi-color puzzles).  Details are in FORMAT.md.
