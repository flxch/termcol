package termcol

import (
    "testing"
    "fmt"
)


// The following tests do not fail.  Run them as `go test -v` and check
// the (colorful and blinking) output.

func TestForeground(t *testing.T) {
    t.Logf("%sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(FgBlack), Effect(FgRed), Effect(FgGreen), Effect(FgYellow), Effect(FgBlue), Effect(FgMagenta), Effect(FgCyan), Effect(FgWhite),
        Effect())
}

func TestBackground(t *testing.T) {
    t.Logf("%sBlack%sRed%sGreen%sYellow%sBlue%sMagenta%sCyan%sWhite%s\n",
        Effect(BgBlack), Effect(BgRed), Effect(BgGreen), Effect(BgYellow), Effect(BgBlue), Effect(BgMagenta), Effect(BgCyan), Effect(BgWhite),
        Effect())
}

func TestStyle(t *testing.T) {
    t.Logf("%sBold%sDim%sItalic%sUnderline%sBlink%sReverse%sHide%s\n",
        Effect(Bold, FgWhite), Effect(BoldOff, Dim), Effect(DimOff, Italic), Effect(ItalicOff, Underline), Effect(UnderlineOff, Blink), Effect(BlinkOff, Reverse), Effect(ReverseOff, Hide),
        Effect())
}

func TestCombinedStyle(t *testing.T) {
    t.Logf("%sBold%sDim%sItalic%sUnderline%sBlink%sReverse%sHide%s\n",
        Effect(Bold, FgWhite), Effect(Dim), Effect(Italic), Effect(Underline), Effect(Blink), Effect(Reverse), Effect(Hide),
        Effect())
}


func ExampleEffect() {
    // Delete or set to false to see colorized output.
    NoColor = true

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

    NoColor = false

    // Output:
    // Foreground: BlackRedGreenYellowBlueMagentaCyanWhite
    // Background: BlackRedGreenYellowBlueMagentaCyanWhite
    // Text Text also underlined Text
}

func ExampleSetEffect() {
    // Delete or set to false to see colorized output.
    NoColor = true

    SetEffect()
    fmt.Printf("default")

    // Output some text in red (foreground) and blue (background)
    SetEffect(FgRed, BgBlue)
    fmt.Printf("colored")

    // Output some text witn reversed colors, i.e., blue (foreground) and red (background)
    SetEffect(Reverse)
    fmt.Printf("reversed")

    // Back to red (foreground) and blue (background) plus underlined
    SetEffect(ReverseOff, Underline)
    fmt.Printf("underlined and colored")

    // Turn off all effects
    SetEffect()
    fmt.Printf("default again\n")

    NoColor = false

    // Output:
    // defaultcoloredreversedunderlined and coloreddefault again
}


func ExampleColorCode() {
    const (
        Info = 1
        Warn = 2
        Err  = 3
    )
    coloring := ColorCode{
        NoColor: true, // Delete or set to false to see colorized output.
        Translate: map[int][]int{
            Info: []int{Bold, BgYellow},
            Warn: []int{Bold, BgMagenta},
            Err:  []int{Bold, BgRed},
        },
    }

    fmt.Printf("%s[I]%s Some information.\n", coloring.Effect(Info), coloring.Effect())
    fmt.Printf("%s[W]%s A warning.\n", coloring.Effect(Warn), coloring.Effect())
    fmt.Printf("%s[E]%s An error message.\n", coloring.Effect(Err), coloring.Effect())
    fmt.Printf("%sno color effect, since filtered out\n", coloring.Effect(Underline, Blink, FgGreen))

    // Output:
    // [I] Some information.
    // [W] A warning.
    // [E] An error message.
    // no color effect, since filtered out
}
