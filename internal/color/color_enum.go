// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package color

import (
	"fmt"
	"strings"
)

const (
	// ColorGray is a Color of type gray.
	ColorGray Color = "#bfbfbf"
	// ColorGreen is a Color of type green.
	ColorGreen Color = "#008000"
	// ColorOrange is a Color of type orange.
	ColorOrange Color = "#ff8c00"
	// ColorRed is a Color of type red.
	ColorRed Color = "#ff0000"
	// ColorWhite is a Color of type white.
	ColorWhite Color = "#ffffff"
	// ColorBlack is a Color of type black.
	ColorBlack Color = "#000000"
)

var ErrInvalidColor = fmt.Errorf("not a valid Color, try [%s]", strings.Join(_ColorNames, ", "))

var _ColorNames = []string{
	string(ColorGray),
	string(ColorGreen),
	string(ColorOrange),
	string(ColorRed),
	string(ColorWhite),
	string(ColorBlack),
}

// ColorNames returns a list of possible string values of Color.
func ColorNames() []string {
	tmp := make([]string, len(_ColorNames))
	copy(tmp, _ColorNames)
	return tmp
}

// ColorValues returns a list of the values for Color
func ColorValues() []Color {
	return []Color{
		ColorGray,
		ColorGreen,
		ColorOrange,
		ColorRed,
		ColorWhite,
		ColorBlack,
	}
}

// String implements the Stringer interface.
func (x Color) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Color) IsValid() bool {
	_, err := ParseColor(string(x))
	return err == nil
}

var _ColorValue = map[string]Color{
	"#bfbfbf": ColorGray,
	"#008000": ColorGreen,
	"#ff8c00": ColorOrange,
	"#ff0000": ColorRed,
	"#ffffff": ColorWhite,
	"#000000": ColorBlack,
}

// ParseColor attempts to convert a string to a Color.
func ParseColor(name string) (Color, error) {
	if x, ok := _ColorValue[name]; ok {
		return x, nil
	}
	return Color(""), fmt.Errorf("%s is %w", name, ErrInvalidColor)
}

// MarshalText implements the text marshaller method.
func (x Color) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Color) UnmarshalText(text []byte) error {
	tmp, err := ParseColor(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Set implements the Golang flag.Value interface func.
func (x *Color) Set(val string) error {
	v, err := ParseColor(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *Color) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *Color) Type() string {
	return "Color"
}