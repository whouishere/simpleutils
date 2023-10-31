# simpleutils

[![build status](https://github.com/whouishere/simpleutils/actions/workflows/build.yml/badge.svg)](https://github.com/whouishere/simpleutils/actions)

simpleutils is as an alternative (though not a replacement) coreutils package.

Note that simpleutils is not supposed to be a full and 100% faithful replacement for GNU coreutils nor it is meant to be POSIX complient.
I just wanted to create this for sheer fun and to challenge myself.

Though I do want to make the code as simple as possible, always __trying__ to follow the Unix philosophy and the KISS principle, note that I am not perfect nor an experienced programmer.
If you notice something can be better, especially more simple, please contribute to the code!

## Building
The only requirement to build the coreutils is GNU Make.

Then you can use `make` to build every utility, and then `make install` to install them into `~/.local/bin`.

By default, each binary will have the `su-` prefix, in order to not meddle with your system's coreutils. You may change the prefix by passing the `BIN_PREFIX` variable to `make` and `make install`:
```
make BIN_PREFIX=<prefix>
make install BIN_PREFIX=<prefix>
```
(replacing `<prefix>` with you desired prefix)

In the same fashion, you can also pass in a different install path to `make install` with the `PREFIX` variable (together with `BIN_PREFIX` or not):
```
make BIN_PREFIX=<prefix>
make install BIN_PREFIX=<prefix> PREFIX=<install path>
```
(replacing `<install path>` with your desired install directory)

If you just want to test out some utility, you can also use the `make run` command:
```
make run UTIL=<utility> ARGS="<arguments>"
```
Replace `<utility>` with your desired utility (`cat`, for example) and `<arguments>` with any command line arguments you wish to pass to that utility (`--help`, for example).

## Progress

| Utility  | Completed | Notes |
| -------- | --------- | ----- |
| cat      | 🟨 | Functional |
| cp       | 🟨 | Simple file copies only |
| dirname  | ✅ |            |
| false    | ✅ |            |
| ln       | 🟨 | Basic functionality only |
| mkdir    | 🟨 | Basic functionality only |
| mv       | 🟨 | Basic functionality only |
| printenv | ✅ |            |
| pwd      | ✅ |            |
| rm       | ✅ | Files only |
| rmdir    | ✅ |            |
| touch    | 🟨 | Basic functionality only |
| true     | ✅ |            |
| whoami   | ✅ |            |

## Acknowledgements and references
- [busybox](https://busybox.net/)
- [GNU coreutils](https://www.gnu.org/software/coreutils/)
- [uutils](https://github.com/uutils/coreutils)
- [linux.die.net](https://linux.die.net/)
- [pubs.opengroup.org](https://pubs.opengroup.org/onlinepubs/9699919799.2018edition/) (POSIX specifications)

## Main goals
- [ ] Finish simpler programs
- [ ] Make simpler programs feature complete and/or equivalent to GNU
- [ ] Write some kind of documentation for every program
- [ ] Finish all coreutils
