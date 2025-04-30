package termcol

import (
    "testing"
    "fmt"
)


// Tests do not fail.  Run them as `go test -v` and check the (colorful and
// blinky) output.
func TestForeground(t *testing.T) {
    fmt.Printf("%sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(FgBlack), Effect(FgRed), Effect(FgGreen), Effect(FgYellow), Effect(FgBlue), Effect(FgMagenta), Effect(FgCyan), Effect(FgWhite), Effect())
}

func TestBackground(t *testing.T) {
    fmt.Printf("%sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(BgBlack), Effect(BgRed), Effect(BgGreen), Effect(BgYellow), Effect(BgBlue), Effect(BgMagenta), Effect(BgCyan), Effect(BgWhite), Effect())
}


func TestStyle(t *testing.T) {
    fmt.Printf("%sBold%sDim%sItalic%sUnderline%sBlink%sReverse%sHide%s\n",
        Effect(Bold, FgWhite), Effect(BoldOff, Dim), Effect(DimOff, Italic), Effect(ItalicOff, Underline), Effect(UnderlineOff, Blink), Effect(BlinkOff, Reverse), Effect(ReverseOff, Hide), Effect())
}

func TestCombinedStyle(t *testing.T) {
    fmt.Printf("%sBold%sDim%sItalic%sUnderline%sBlink%sReverse%sHide%s\n",
        Effect(Bold, FgWhite), Effect(Dim), Effect(Italic), Effect(Underline), Effect(Blink), Effect(Reverse), Effect(Hide), Effect())
}


func TestCombinations(t *testing.T) {
    SetEffect()
    fmt.Printf("default")
    SetEffect(FgRed, BgBlue)
    fmt.Printf("colored")
    SetEffect(Reverse)
    fmt.Printf("reversed")
    SetEffect(ReverseOff, Underline)
    fmt.Printf("underlined and colored")
    SetEffect()
    fmt.Printf("default again\n")
}


func ExampleEffect() {
    // Output text in various foreground colors
    fmt.Printf("Foreground: %sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(FgBlack), Effect(FgRed), Effect(FgGreen), Effect(FgYellow), Effect(FgBlue), Effect(FgMagenta), Effect(FgCyan), Effect(FgWhite),
        Effect())

    // Output text in various background colors
    fmt.Printf("Background: %sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(BgBlack), Effect(BgRed), Effect(BgGreen), Effect(BgYellow), Effect(BgBlue), Effect(BgMagenta), Effect(BgCyan), Effect(BgWhite),
        Effect())

    // Output text in red, bold, and underlined
    fmt.Printf("%sText %sText also underlined%s Text%s\n",
        Effect(FgRed, Bold), Effect(Underline), Effect(UnderlineOff), 
        Effect())
}
