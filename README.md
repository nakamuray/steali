steali -- steal first N lines from stdin and display it
=======================================================

`steali` steal first N (default:1) lines from stdin and display it.
After that, it execute a command specified by command line arguments and pass remaining input to the command.

```
Usage: steali [-e] [-n int] command [args]
  -e    output lines to stderr instead of stdout
  -n int
        how many lines to steal (default 1)
```

example:

```
$ ps aux | steali grep foo
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
u        25053  0.0  0.0 103224  4832 pts/13   Sl+  20:35   0:00 steali grep foo
```
