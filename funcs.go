package termcol

import (
    "strings"
    "strconv"
    "fmt"
    "io"
)


// `effect` returns the escape sequence as a string for setting the specified
// effects `effs`. 
func effect(effs ...int) string {
    if len(effs) == 0 {
        return reset
    }

    var b strings.Builder
    b.WriteString("\x1b[")

    for i, eff := range effs {
        b.WriteString(strconv.Itoa(eff))
        // If not the last effect, separate the effect from the previous one by
        // a semicolon.
        if i < len(effs) - 1 {
            b.WriteByte(';')
        }

    }

    b.WriteByte('m')
    return b.String()
}


// If `NoColor` is set to true, no terminal color escape sequences are used.  For
// instance, the function `Effect` always returns the empty string independently
// from its arguments.  The default value of `NoColor` is false.
var NoColor bool


// `Effect` returns a string for the effects `effs`.  The string consists of the 
// respective terminal color escape sequences.  If `effs` is empty, all effects
// are turned off.  The returned string can, e.g., be used in fmt.Printf calls.
func Effect(effs ...int) string {
    if NoColor {
        return ""
    }
    return effect(effs...)
}


// `SetEffect` is similar to `Effect` but directly prints (stdout) the respective
// escape sequence.
func SetEffect(effs ...int) error {
    if NoColor {
        return nil
    }
    _, err := fmt.Printf("%s", effect(effs...))
    return err
}

// `FSetEffect` is similar to `SetEffect` but uses the writer `w`.
func FSetEffect(w io.Writer, effs ...int) error {
    if NoColor {
        return nil
    }
    _, err := fmt.Fprintf(w, "%s", effect(effs...))
    return err
}


type ColorCode struct {
    // If true, ignore color codes.
    NoColor bool
}

func (c ColorCode) Effect(effs ...int) string {
    if c.NoColor {
        return ""
    }
    return effect(effs...)
}

func (c ColorCode) SetEffect(effs ...int) error {
    if c.NoColor {
        return nil
    }
    _, err := fmt.Printf("%s", effect(effs...))
    return err
}

func (c ColorCode) FSetEffect(w io.Writer, effs ...int) error {
    if c.NoColor {
        return nil
    }  
    _, err := fmt.Fprintf(w, "%s", effect(effs...))
    return err
}
