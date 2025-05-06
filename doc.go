package termcol

// The package contains mainly constants for setting terminal colors.  The
// function `Effect` can, e.g., be used in `fmt.Printf` calls to set foreground
// and background colors.  The arguments to this function are any combination of
// the package constants.  It returns a string for setting the specified color
// combination in a terminal.

// Note that `Effect()` resets to the default setting.

// When setting the package variable `NoColor` to true, the function Effect
// always returns the empty string independently from its arguments.

// The type `ColorCode` allows one to turn off coloring schemes "locally".
// Concretely, `NoColor` is local to each `ColorCode` instance.
