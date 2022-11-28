# KEG text search

Search text for matches inside a KEG repository.

## Install

```sh
go install github.com/BuddhiLW/keg-search@latest
```

## How to use it

Let's say your KEG is at your home directory, named _keg_. So, use `-p ~/keg` to
search and parse the KEG repository, in this directory.

Search for the regex: "Idiot", and use case sensitivity `-c true` [^1]. If
omitted, the search will be case insensitive as default.

One-liner:

```sh
keg-search -p ~/keg -reg "Idiot" -c true
```

Output:

```sh
/home/buddhilw/keg/3/README.md
# Dostoievsky's The Idiot (20221125132220)
```

### NOTES

[^1]: pass any string to `-c` and the search will become case sensitive; default: insensitive.
