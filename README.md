# golog

A simple logger for go, similar to the builtin "log" package.

Each log line is prefixed with its RFC3339 timestamp in UTC with millisecond resolution, folowed by a vertical pipe, e.g.

```
2022-12-11T20:06:02.461Z | Hello, World!
```

The `Debug`, `Error`, and `Fatal` functions print to stderr their respective levels in all caps before the given message.

```
2022-12-11T20:06:33.886Z | DEBUG: only prints if the "DEBUG" env var is set.
2022-12-11T20:06:33.886Z | ERROR: see also "Err{,f}" for error types.
2022-12-11T20:06:33.886Z | FATAL: exits the process with code 1.
```

The `...f` variants of the above functions treat the message as a format string and take variable arguments.

The `Debug{,f}` functions do nothing unless the `DEBUG` environment variable is non-empty.

The `Fatal{,f}` functions exit the process with exit code 1 after printing their messages.

The `Must{,1,2,3}` functions log a fatal error if the given error is not nil, returning the corresponding number of additional arguments.

# Error reporting

TODO

The `Err{,f}` functions log the given error, similar to `Error{,f}`. All `Err{,or}{,f}` variants handle the given error/message using the given handler

The `func HandleError(error)` callback function can handle errors, for example, by reporting to your error tracking software.
