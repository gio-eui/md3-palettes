// Copyright (c) 2023 https://github.com/gio-eui
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR a PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// SPDX-License-Identifier: MIT

package palette

import "github.com/gio-eui/md3-palettes/scheme"

type Palette struct {
	Light  *scheme.Scheme
	Dark   *scheme.Scheme
	Active *scheme.Scheme

	IsDark bool
}

// NewPaletteFromInt creates a new palette from a primary color.
// Each color is formatted as an int representing an argb color.
// For example, 0xff000000 is black, 0xffffffff is white, 0xffff0000 is red, etc.
func NewPaletteFromInt(primary int, secondary int, tertiary int, neutral int, neutralVariant int) *Palette {
	// Light scheme
	light := scheme.Light(primary, secondary, tertiary, neutral, neutralVariant).WithErrorTonalPalette(scheme.ErrorTonalPalette, false)
	// Dark scheme
	dark := scheme.Dark(primary, secondary, tertiary, neutral, neutralVariant).WithErrorTonalPalette(scheme.ErrorTonalPalette, true)
	// Active scheme is by default the light scheme
	active := light
	isDark := false

	// Create the palette
	return &Palette{
		Light:  light,
		Dark:   dark,
		Active: active,
		IsDark: isDark,
	}
}

// NewDefaultPalette creates a new palette with the default colors.
func NewDefaultPalette() *Palette {
	// Light scheme
	light := scheme.Light(
		scheme.PrimaryTonalPalette.Tone(50),
		scheme.SecondaryTonalPalette.Tone(50),
		scheme.TertiaryTonalPalette.Tone(50),
		scheme.NeutralTonalPalette.Tone(50),
		scheme.NeutralVariantTonalPalette.Tone(50),
	).WithErrorTonalPalette(scheme.ErrorTonalPalette, false)
	// Dark scheme
	dark := scheme.Dark(
		scheme.PrimaryTonalPalette.Tone(50),
		scheme.SecondaryTonalPalette.Tone(50),
		scheme.TertiaryTonalPalette.Tone(50),
		scheme.NeutralTonalPalette.Tone(50),
		scheme.NeutralVariantTonalPalette.Tone(50),
	).WithErrorTonalPalette(scheme.ErrorTonalPalette, true)
	// Active scheme is by default the light scheme
	active := light
	isDark := false

	// Create the palette
	return &Palette{
		Light:  light,
		Dark:   dark,
		Active: active,
		IsDark: isDark,
	}
}

// SwitchMode changes the active scheme between light and dark.
func (p *Palette) SwitchMode(isDark bool) {
	if isDark {
		p.Active = p.Dark
	} else {
		p.Active = p.Light
	}
	p.IsDark = isDark
}
