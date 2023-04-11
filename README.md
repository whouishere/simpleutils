# simpleutils
simpleutils is as an alternative (though not a replacement) coreutils package.

Note that simpleutils is not supposed to be a full and 100% faithful replacement for GNU coreutils nor it is meant to be POSIX complient.
I just wanted to create this for sheer fun and to challenge myself.

Though I do want to make the code as simple as possible, always __trying__ to follow the Unix philosophy and the KISS principle, note that I am not perfect nor an experienced programmer.
If you notice something can be better, especially more simple, please contribute to the code!

## Running
If you really want to test some of the utilities out, you may install [Task](https://taskfile.dev/) first in order to use the build system.

With that out of the way, it is possible to build every utility using the `task build`.

If you just want to test out some utility though, you can use the `task run UTIL=<utility> -- <args>` command. Substitute `<utility>` with an utility, `cat` as an example, and `<args>` with the desired command line arguments, like `--help`.

## Progress
<div style="text-align: center;">

| Utility  | Completed | Notes |
| -------- | --------- | ----- |
| cat      | 🟨 | Functional |
| dirname  | ✅ |            |
| false    | ✅ |            |
| printenv | ✅ |            |
| pwd      | ✅ |            |
| rmdir    | ✅ |            |
| true     | ✅ |            |
| whoami   | ✅ |            |

</div>

## Main goals
- [ ] Finish the simpler programs
- [ ] Make the simpler programs be feature complete and/or equivalent to GNU
- [ ] Write some kind of documentation for every program
- [ ] Finish all coreutils
