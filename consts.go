package termcol

// Default setting (turn off all effects)
const (
    reset  = "\x1b[0m"
)


// Internals
const (
    black           = 0
    red             = 1
    green           = 2
    yellow          = 3
    blue            = 4
    magenta         = 5
    cyan            = 6
    white           = 7

    off             = 20
    fg              = 30
    bg              = 40
    bright          = 60
)

// Foreground colors
const (
    FgBlack         = fg + black
    FgRed           = fg + red
    FgGreen         = fg + green
    FgYellow        = fg + yellow
    FgBlue          = fg + blue
    FgMagenta       = fg + magenta
    FgCyan          = fg + cyan
    FgWhite         = fg + white

    FgBrightBlack   = fg + bright + black
    FgBrightRed     = fg + bright + red
    FgBrightGreen   = fg + bright + green
    FgBrightYellow  = fg + bright + yellow
    FgBrightBlue    = fg + bright + blue
    FgBrightMagenta = fg + bright + magenta
    FgBrightCyan    = fg + bright + cyan
    FgBrightWhite   = fg + bright + white
)

// Background colors
const (
    BgBlack         = bg + black
    BgRed           = bg + red
    BgGreen         = bg + green
    BgYellow        = bg + yellow
    BgBlue          = bg + blue
    BgMagenta       = bg + magenta
    BgCyan          = bg + cyan
    BgWhite         = bg + white

    BgBrightBlack   = bg + bright + black
    BgBrightRed     = bg + bright + red
    BgBrightGreen   = bg + bright + green
    BgBrightYellow  = bg + bright + yellow
    BgBrightBlue    = bg + bright + blue
    BgBrightMagenta = bg + bright + magenta
    BgBrightCyan    = bg + bright + cyan
    BgBrightWhite   = bg + bright + white
)

// Styles
const (
    Bold         = 1
    Dim          = 2
    Italic       = 3
    Underline    = 4
    Blink        = 5
    Reverse      = 7
    Hide         = 8

    BoldOff      = off + Bold
    DimOff       = off + Dim
    ItalicOff    = off + Italic
    UnderlineOff = off + Underline
    BlinkOff     = off + Blink
    ReverseOff   = off + Reverse
    HideOff      = off + Hide
)
