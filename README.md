# KEG text search

Search text for matches inside a KEG repository.

## Install

```sh
go install github.com/BuddhiLW/keg-search@latest
```

## How to use it

In a straight-forward manner,

```sh
keg-search -reg "Idiot"
```

will return any matches under `~/keg` to the `"Idiot"` regex, in a
case-insensitive search.

### All optional arguments

**-p**: `path`
Let's say your KEG is at a directory named _/home/user/path-to/keg_. So, use `-p ~/path-to/keg` to search and parse the KEG repository, in that directory [^1].

**-c**: `case-sensitivity`
Further, let's look for the regex `"Idiot"`, in a case-sensitive manner. use
case sensitivity `-c true` option [^2]. Defaults to: case-insensitive.

**-s**: `surrounding-context`
Also, in our search, we can specify how much characters we want as context. This
is, how many characters right, and left, from the pattern-match we are searching
for.

One-liner:

```sh
keg-search -p ~/keg -reg "Idiot" -c true -s 46
```

Output:

```sh
/home/buddhilw/keg/3/README.md
[20 25]
# Dostoievsky's The Idiot (20221125132220)

A novel extremely intricate
```

### NOTES

[^1]: **Default:** looks for `~/keg` if _path_ argument is omitted.
[^2]: pass any string to `-c` and the search will become case sensitive; default: insensitive.
