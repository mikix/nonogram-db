## Rules

* The `non` file extension should be used to identify this format.

* Parsers must ignore any line that they don't recognize.

* There is no guaranteed order of keys, except that the `height` and `width` come before the `columns`, `rows`, and `goal` keys.

* The only required keys are `width`, `height`, `rows`, and `columns`.

* Encoding is always utf8.

* String values must be quoted and may have HTML escape codes in them.

* Blank lines can appear anywhere in between keys and should be ignored.

## Keys

* `catalogue`: string; freeform description of puzzle origin

* `title`: string; user-presentable name of puzzle

* `by`: string; author

* `copyright`: string; freeform description of copyright

* `license`: either license code from [SPDX][spdx] (non-quoted-string) or a quoted string freeform description
[spdx]: http://spdx.org/licenses/

* `color`: followed by a character designation (usually a letter), then a space, then a six-digit hex color like #808080 (an example would be `color a #ff0000`); can be specified multiple times for multiple characters; these characters are used in the `rows`, `columns`, and `goal` keys; color designation characters may be present in such even if no `color` keys were specified, in which case the puzzle interface can pick its own colors

* `width`: int; number of columns

* `height`: int; number of rows

* `rows`: starts a sequence of lines, in number equal to `height`; represents the hints for each row; blank lines may be present if a row has no hints; each hint is a number, followed by an optional color character, separated by a comma from the next hint (an example multi-colored line would be `3b,1d,6b,4c,3a,1b,2b`)

* `columns`: same deal as `rows` but for columns

* `goal`: string; sequence of answer characters from top left of the puzzle along the first row, then the second row, etc.; 0 is a blank; anything else is a filled spot, likely either a 1 or a color character designation

## Bundling

You can bundle more than one puzzle in the same file.  Just separate them with a divider of four `=` characters like `====` on a line by itself.  Use the file extension `nonpack` to denote such bundles.  This can be useful to logically group a sequence of puzzles and avoid many tiny files.

Obviously, you should probably only bundle files with compatible licenses.

You may also consider compressing your nonpacks with zlib and shipping them as `nonopack.gz` files.

## Identification

In order to identify a given puzzle (e.g. to store state data about it), you may be tempted to use the filename or puzzle name.  But I'd recommend a hash of the `rows` and `columns` hint data.  That will survive name changes and should still uniquely identify a given puzzle.  If the hints change, it's a different puzzle after all!

## Example

    catalogue "webpbn.com #1"
    title "Demo Puzzle from Front Page"
    by "Jan Wolter"
    copyright "&copy; Copyright 2004 by Jan Wolter"
    license CC-BY-3.0
    width 5
    height 10
    
    rows
    2
    2,1
    1,1
    3
    1,1
    1,1
    2
    1,1
    1,2
    2
    
    columns
    2,1
    2,1,3
    7
    1,3
    2,1
    
    goal "01100011010010101110101001010000110010100101111000"

## Changes

* 2015-05-11: Add the `color` key (and associated use of the color designation character in the `rows`, `columns`, and `goal` keys).  Also add the `license` key and bundling support.

* 2011?: Steve Simpson's original format.
