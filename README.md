# golog

A simple logger for go.

The `Debug`, `Print`, `Error`, and `Fatal` functions print their level before the given message.

The `...f` variants of the above functions treat the message as a format string and take variable arguments.

The `Debug{,f}` functions do nothing unless the `DEBUG` environment variable is non-empty.

The `Fatal{,f}` functions exit the process with exit code 1 after printing their messages.

# Error reporting

The `Err{,f}` functions log the given error message, similar to `Error{,f}`.

TODO: the `func HandleError(error)` callback function can handle errors, for example, by reporting to your error tracking software.
