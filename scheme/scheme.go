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

package scheme

import (
	"github.com/gio-eui/md3-colors/palettes"
	"image/color"
)

// Scheme is a collection of colors that are used to represent the UI of an app.
type Scheme struct {
	// The primary key color is used to derive roles for key components across the UI,
	// such as the FAB, prominent buttons, active states, as well as the tint of elevated surfaces.
	Primary            color.NRGBA // primary40 / primary80
	OnPrimary          color.NRGBA // primary90 / primary30
	PrimaryContainer   color.NRGBA // primary100 / primary20
	OnPrimaryContainer color.NRGBA // primary10 / primary90
	InversePrimary     color.NRGBA // primary80 /primary40

	// The secondary key color is used for less prominent components in the UI such as filter chips,
	// while expanding the opportunity for color expression.
	Secondary            color.NRGBA // secondary40 / secondary80
	OnSecondary          color.NRGBA // secondary90 / secondary30
	SecondaryContainer   color.NRGBA // secondary100 / secondary20
	OnSecondaryContainer color.NRGBA // secondary10 / secondary90

	// The tertiary key color is used to derive the roles of contrasting accents that can be used
	// to balance primary and secondary colors or bring heightened attention to an element.
	// The tertiary color role is left for teams to use at their discretion and is intended
	// to support broader color expression in products.
	Tertiary            color.NRGBA // tertiary40 / tertiary80
	OnTertiary          color.NRGBA // tertiary90 / tertiary30
	TertiaryContainer   color.NRGBA // tertiary100 / tertiary20
	OnTertiaryContainer color.NRGBA // tertiary10 / tertiary90

	// The custom key color is used to derive the roles of contrasting accents that can be used
	// to balance primary and secondary colors or bring heightened attention to an element.
	// The custom color role is left for teams to use at their discretion and is intended
	// to support broader color expression in products.
	Custom            color.NRGBA // custom40 / custom80
	OnCustom          color.NRGBA // custom90 / custom30
	CustomContainer   color.NRGBA // custom100 / custom20
	OnCustomContainer color.NRGBA // custom10 / custom90

	//
	Error            color.NRGBA // error40 / error80
	OnError          color.NRGBA // error90 / error30
	ErrorContainer   color.NRGBA // error100 / error20
	OnErrorContainer color.NRGBA // error10 / error90

	// The neutral key color is used to derive surface color roles for backgrounds, as well
	// as colors used for high emphasis text and icons.
	// The neutral variant key color is used to derive color roles for medium emphasis elements
	// like text, icons, and component outlines
	Surface                 color.NRGBA // neutral98 / neutral6
	SurfaceDim              color.NRGBA // neutral87 / neutral6
	SurfaceBright           color.NRGBA // neutral98 / neutral24
	SurfaceContainerLowest  color.NRGBA // neutral100 / neutral4
	SurfaceContainerLow     color.NRGBA // neutral96 / neutral10
	SurfaceContainer        color.NRGBA // neutral94 / neutral12
	SurfaceContainerHigh    color.NRGBA // neutral92 / neutral17
	SurfaceContainerHighest color.NRGBA // neutral90 / neutral22
	SurfaceVariant          color.NRGBA // neutral-variant90 / neutral-variant30
	OnSurface               color.NRGBA // neutral0 / neutral90
	OnSurfaceVariant        color.NRGBA // neutral-variant30 / neutral-variant80
	InverseSurface          color.NRGBA // neutral20 / neutral90
	InverseOnSurface        color.NRGBA // neutral95 / neutral20

	Background   color.NRGBA // neutral98 / neutral6
	OnBackground color.NRGBA // neutral0 / neutral90

	Outline        color.NRGBA // neutral-variant50 / neutral-variant60
	OutlineVariant color.NRGBA // neutral-variant80 / neutral-variant30
	Shadow         color.NRGBA // neutral0 / neutral0
	ShadowTint     color.NRGBA // primary / primary
	Scrim          color.NRGBA // neutral0 / neutral0

	PrimaryTone        color.NRGBA // primary50
	SecondaryTone      color.NRGBA // secondary50
	TertiaryTone       color.NRGBA // tertiary50
	CustomTone         color.NRGBA // custom50
	NeutralTone        color.NRGBA // neutral50
	NeutralVariantTone color.NRGBA // neutral-variant50
	ErrorTone          color.NRGBA // error50

	primaryTone        *palettes.TonalPalette
	secondaryTone      *palettes.TonalPalette
	tertiaryTone       *palettes.TonalPalette
	customTone         *palettes.TonalPalette
	neutralTone        *palettes.TonalPalette
	neutralVariantTone *palettes.TonalPalette
	errorTone          *palettes.TonalPalette
}

// Light creates a light scheme based on the given color.
func Light(primary int, secondary int, tertiary int, neutral int, neutralVariant int) *Scheme {
	return lightFromTonalPalette(palettes.NewTonalPaletteFromInt(primary), palettes.NewTonalPaletteFromInt(secondary), palettes.NewTonalPaletteFromInt(tertiary), palettes.NewTonalPaletteFromInt(neutral), palettes.NewTonalPaletteFromInt(neutralVariant))
}

// Dark creates a dark scheme based on the given color.
func Dark(primary int, secondary int, tertiary int, neutral int, neutralVariant int) *Scheme {
	return darkFromTonalPalette(palettes.NewTonalPaletteFromInt(primary), palettes.NewTonalPaletteFromInt(secondary), palettes.NewTonalPaletteFromInt(tertiary), palettes.NewTonalPaletteFromInt(neutral), palettes.NewTonalPaletteFromInt(neutralVariant))
}

// lightFromTonalPalette creates a light scheme based on the given core palette.
func lightFromTonalPalette(primary *palettes.TonalPalette, secondary *palettes.TonalPalette, tertiary *palettes.TonalPalette, neutral *palettes.TonalPalette, neutralVariant *palettes.TonalPalette) *Scheme {
	s := &Scheme{}
	s = s.WithPrimaryTonalPalette(primary, false)
	s = s.WithSecondaryTonalPalette(secondary, false)
	s = s.WithTertiaryTonalPalette(tertiary, false)
	s = s.WithNeutralTonalPalette(neutral, false)
	s = s.WithNeutralVariantTonalPalette(neutralVariant, false)
	s = s.WithErrorTonalPalette(ErrorTonalPalette, false)
	return s
}

// darkFromTonalPalette creates a dark scheme based on the given core palette.
func darkFromTonalPalette(primary *palettes.TonalPalette, secondary *palettes.TonalPalette, tertiary *palettes.TonalPalette, neutral *palettes.TonalPalette, neutralVariant *palettes.TonalPalette) *Scheme {
	s := &Scheme{}
	s = s.WithPrimaryTonalPalette(primary, true)
	s = s.WithSecondaryTonalPalette(secondary, true)
	s = s.WithTertiaryTonalPalette(tertiary, true)
	s = s.WithNeutralTonalPalette(neutral, true)
	s = s.WithNeutralVariantTonalPalette(neutralVariant, true)
	s = s.WithErrorTonalPalette(ErrorTonalPalette, true)
	return s
}

// WithPrimaryTonalPalette sets the primary tonal palette of the scheme.
func (s *Scheme) WithPrimaryTonalPalette(primaryTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.Primary = s.nrgba(primaryTone.Tone(80))
		s.OnPrimary = s.nrgba(primaryTone.Tone(20))
		s.PrimaryContainer = s.nrgba(primaryTone.Tone(30))
		s.OnPrimaryContainer = s.nrgba(primaryTone.Tone(90))
		s.InversePrimary = s.nrgba(primaryTone.Tone(40))
	} else {
		s.Primary = s.nrgba(primaryTone.Tone(40))
		s.OnPrimary = s.nrgba(primaryTone.Tone(100))
		s.PrimaryContainer = s.nrgba(primaryTone.Tone(90))
		s.OnPrimaryContainer = s.nrgba(primaryTone.Tone(10))
		s.InversePrimary = s.nrgba(primaryTone.Tone(80))
	}
	s.ShadowTint = s.nrgba(primaryTone.GetKeyColor().ToInt())
	s.PrimaryTone = s.nrgba(primaryTone.Tone(50))
	s.primaryTone = primaryTone
	return s
}

// WithSecondaryTonalPalette sets the secondary tonal palette of the scheme.
func (s *Scheme) WithSecondaryTonalPalette(secondaryTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.Secondary = s.nrgba(secondaryTone.Tone(80))
		s.OnSecondary = s.nrgba(secondaryTone.Tone(20))
		s.SecondaryContainer = s.nrgba(secondaryTone.Tone(30))
		s.OnSecondaryContainer = s.nrgba(secondaryTone.Tone(90))
	} else {
		s.Secondary = s.nrgba(secondaryTone.Tone(40))
		s.OnSecondary = s.nrgba(secondaryTone.Tone(100))
		s.SecondaryContainer = s.nrgba(secondaryTone.Tone(90))
		s.OnSecondaryContainer = s.nrgba(secondaryTone.Tone(10))
	}
	s.SecondaryTone = s.nrgba(secondaryTone.Tone(50))
	s.secondaryTone = secondaryTone
	return s
}

// WithTertiaryTonalPalette sets the tertiary tonal palette of the scheme.
func (s *Scheme) WithTertiaryTonalPalette(tertiaryTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.Tertiary = s.nrgba(tertiaryTone.Tone(80))
		s.OnTertiary = s.nrgba(tertiaryTone.Tone(20))
		s.TertiaryContainer = s.nrgba(tertiaryTone.Tone(30))
		s.OnTertiaryContainer = s.nrgba(tertiaryTone.Tone(90))
	} else {
		s.Tertiary = s.nrgba(tertiaryTone.Tone(40))
		s.OnTertiary = s.nrgba(tertiaryTone.Tone(100))
		s.TertiaryContainer = s.nrgba(tertiaryTone.Tone(90))
		s.OnTertiaryContainer = s.nrgba(tertiaryTone.Tone(10))
	}
	s.TertiaryTone = s.nrgba(tertiaryTone.Tone(50))
	s.tertiaryTone = tertiaryTone
	return s
}

// WithCustomTonalPalette sets the custom tonal palette of the scheme.
func (s *Scheme) WithCustomTonalPalette(customTone *palettes.TonalPalette, isDark bool) *Scheme {
	if customTone == nil {
		return s
	}
	if isDark {
		s.Custom = s.nrgba(customTone.Tone(80))
		s.OnCustom = s.nrgba(customTone.Tone(20))
		s.CustomContainer = s.nrgba(customTone.Tone(30))
		s.OnCustomContainer = s.nrgba(customTone.Tone(90))
	} else {
		s.Custom = s.nrgba(customTone.Tone(40))
		s.OnCustom = s.nrgba(customTone.Tone(100))
		s.CustomContainer = s.nrgba(customTone.Tone(90))
		s.OnCustomContainer = s.nrgba(customTone.Tone(10))
	}
	s.CustomTone = s.nrgba(customTone.Tone(50))
	s.customTone = customTone
	return s
}

// WithNeutralTonalPalette sets the neutral tonal palette of the scheme.
func (s *Scheme) WithNeutralTonalPalette(neutralTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.Surface = s.nrgba(neutralTone.Tone(6))
		s.SurfaceDim = s.nrgba(neutralTone.Tone(6))
		s.SurfaceBright = s.nrgba(neutralTone.Tone(24))
		s.SurfaceContainerLowest = s.nrgba(neutralTone.Tone(4))
		s.SurfaceContainerLow = s.nrgba(neutralTone.Tone(10))
		s.SurfaceContainer = s.nrgba(neutralTone.Tone(12))
		s.SurfaceContainerHigh = s.nrgba(neutralTone.Tone(17))
		s.SurfaceContainerHighest = s.nrgba(neutralTone.Tone(22))
		s.OnSurface = s.nrgba(neutralTone.Tone(90))
		s.InverseSurface = s.nrgba(neutralTone.Tone(90))
		s.InverseOnSurface = s.nrgba(neutralTone.Tone(20))
		s.Background = s.nrgba(neutralTone.Tone(6))
		s.OnBackground = s.nrgba(neutralTone.Tone(90))
		s.Shadow = s.nrgba(neutralTone.Tone(0))
		s.Scrim = s.nrgba(neutralTone.Tone(0))
	} else {
		s.Surface = s.nrgba(neutralTone.Tone(98))
		s.SurfaceDim = s.nrgba(neutralTone.Tone(87))
		s.SurfaceBright = s.nrgba(neutralTone.Tone(98))
		s.SurfaceContainerLowest = s.nrgba(neutralTone.Tone(100))
		s.SurfaceContainerLow = s.nrgba(neutralTone.Tone(96))
		s.SurfaceContainer = s.nrgba(neutralTone.Tone(94))
		s.SurfaceContainerHigh = s.nrgba(neutralTone.Tone(92))
		s.SurfaceContainerHighest = s.nrgba(neutralTone.Tone(90))
		s.OnSurface = s.nrgba(neutralTone.Tone(0))
		s.InverseSurface = s.nrgba(neutralTone.Tone(20))
		s.InverseOnSurface = s.nrgba(neutralTone.Tone(95))
		s.Background = s.nrgba(neutralTone.Tone(98))
		s.OnBackground = s.nrgba(neutralTone.Tone(0))
		s.Shadow = s.nrgba(neutralTone.Tone(0))
		s.Scrim = s.nrgba(neutralTone.Tone(0))
	}
	s.NeutralTone = s.nrgba(neutralTone.Tone(50))
	s.neutralTone = neutralTone
	return s
}

// WithNeutralVariantTonalPalette sets the neutral variant tonal palette of the scheme.
func (s *Scheme) WithNeutralVariantTonalPalette(neutralVariantTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.SurfaceVariant = s.nrgba(neutralVariantTone.Tone(30))
		s.OnSurfaceVariant = s.nrgba(neutralVariantTone.Tone(80))
		s.Outline = s.nrgba(neutralVariantTone.Tone(60))
		s.OutlineVariant = s.nrgba(neutralVariantTone.Tone(30))
	} else {
		s.SurfaceVariant = s.nrgba(neutralVariantTone.Tone(90))
		s.OnSurfaceVariant = s.nrgba(neutralVariantTone.Tone(30))
		s.Outline = s.nrgba(neutralVariantTone.Tone(50))
		s.OutlineVariant = s.nrgba(neutralVariantTone.Tone(80))
	}
	s.NeutralVariantTone = s.nrgba(neutralVariantTone.Tone(50))
	s.neutralVariantTone = neutralVariantTone
	return s
}

// WithErrorTonalPalette sets the error tonal palette of the scheme.
func (s *Scheme) WithErrorTonalPalette(errorTone *palettes.TonalPalette, isDark bool) *Scheme {
	if isDark {
		s.Error = s.nrgba(errorTone.Tone(80))
		s.OnError = s.nrgba(errorTone.Tone(20))
		s.ErrorContainer = s.nrgba(errorTone.Tone(30))
		s.OnErrorContainer = s.nrgba(errorTone.Tone(90))
	} else {
		s.Error = s.nrgba(errorTone.Tone(40))
		s.OnError = s.nrgba(errorTone.Tone(100))
		s.ErrorContainer = s.nrgba(errorTone.Tone(90))
		s.OnErrorContainer = s.nrgba(errorTone.Tone(10))
	}
	s.ErrorTone = s.nrgba(errorTone.Tone(50))
	s.errorTone = errorTone
	return s
}

// WithPrimary sets the primary color of the scheme.
func (s *Scheme) WithPrimary(primary int) *Scheme {
	s.Primary = s.nrgba(primary)
	return s
}

// WithOnPrimary sets the on-primary color of the scheme.
func (s *Scheme) WithOnPrimary(onPrimary int) *Scheme {
	s.OnPrimary = s.nrgba(onPrimary)
	return s
}

// WithPrimaryContainer sets the primary container color of the scheme.
func (s *Scheme) WithPrimaryContainer(primaryContainer int) *Scheme {
	s.PrimaryContainer = s.nrgba(primaryContainer)
	return s
}

// WithOnPrimaryContainer sets the on-primary container color of the scheme.
func (s *Scheme) WithOnPrimaryContainer(onPrimaryContainer int) *Scheme {
	s.OnPrimaryContainer = s.nrgba(onPrimaryContainer)
	return s
}

// WithInversePrimary sets the inverse primary color of the scheme.
func (s *Scheme) WithInversePrimary(inversePrimary int) *Scheme {
	s.InversePrimary = s.nrgba(inversePrimary)
	return s
}

// WithSecondary sets the secondary color of the scheme.
func (s *Scheme) WithSecondary(secondary int) *Scheme {
	s.Secondary = s.nrgba(secondary)
	return s
}

// WithOnSecondary sets the on-secondary color of the scheme.
func (s *Scheme) WithOnSecondary(onSecondary int) *Scheme {
	s.OnSecondary = s.nrgba(onSecondary)
	return s
}

// WithSecondaryContainer sets the secondary container color of the scheme.
func (s *Scheme) WithSecondaryContainer(secondaryContainer int) *Scheme {
	s.SecondaryContainer = s.nrgba(secondaryContainer)
	return s
}

// WithOnSecondaryContainer sets the on-secondary container color of the scheme.
func (s *Scheme) WithOnSecondaryContainer(onSecondaryContainer int) *Scheme {
	s.OnSecondaryContainer = s.nrgba(onSecondaryContainer)
	return s
}

// WithTertiary sets the tertiary color of the scheme.
func (s *Scheme) WithTertiary(tertiary int) *Scheme {
	s.Tertiary = s.nrgba(tertiary)
	return s
}

// WithOnTertiary sets the on-tertiary color of the scheme.
func (s *Scheme) WithOnTertiary(onTertiary int) *Scheme {
	s.OnTertiary = s.nrgba(onTertiary)
	return s
}

// WithTertiaryContainer sets the tertiary container color of the scheme.
func (s *Scheme) WithTertiaryContainer(tertiaryContainer int) *Scheme {
	s.TertiaryContainer = s.nrgba(tertiaryContainer)
	return s
}

// WithOnTertiaryContainer sets the on-tertiary container color of the scheme.
func (s *Scheme) WithOnTertiaryContainer(onTertiaryContainer int) *Scheme {
	s.OnTertiaryContainer = s.nrgba(onTertiaryContainer)
	return s
}

// WithError sets the error color of the scheme.
func (s *Scheme) WithError(err int) *Scheme {
	s.Error = s.nrgba(err)
	return s
}

// WithOnError sets the on-error color of the scheme.
func (s *Scheme) WithOnError(onError int) *Scheme {
	s.OnError = s.nrgba(onError)
	return s
}

// WithErrorContainer sets the error container color of the scheme.
func (s *Scheme) WithErrorContainer(errorContainer int) *Scheme {
	s.ErrorContainer = s.nrgba(errorContainer)
	return s
}

// WithOnErrorContainer sets the on-error container color of the scheme.
func (s *Scheme) WithOnErrorContainer(onErrorContainer int) *Scheme {
	s.OnErrorContainer = s.nrgba(onErrorContainer)
	return s
}

// WithSurface sets the surface color of the scheme.
func (s *Scheme) WithSurface(surface int) *Scheme {
	s.Surface = s.nrgba(surface)
	return s
}

// WithSurfaceDim sets the dim surface color of the scheme.
func (s *Scheme) WithSurfaceDim(surfaceDim int) *Scheme {
	s.SurfaceDim = s.nrgba(surfaceDim)
	return s
}

// WithSurfaceBright sets the bright surface color of the scheme.
func (s *Scheme) WithSurfaceBright(surfaceBright int) *Scheme {
	s.SurfaceBright = s.nrgba(surfaceBright)
	return s
}

// WithSurfaceContainerLowest sets the lowest surface container color of the scheme.
func (s *Scheme) WithSurfaceContainerLowest(surfaceContainerLowest int) *Scheme {
	s.SurfaceContainerLowest = s.nrgba(surfaceContainerLowest)
	return s
}

// WithSurfaceContainerLow sets the low surface container color of the scheme.
func (s *Scheme) WithSurfaceContainerLow(surfaceContainerLow int) *Scheme {
	s.SurfaceContainerLow = s.nrgba(surfaceContainerLow)
	return s
}

// WithSurfaceContainer sets the surface container color of the scheme.
func (s *Scheme) WithSurfaceContainer(surfaceContainer int) *Scheme {
	s.SurfaceContainer = s.nrgba(surfaceContainer)
	return s
}

// WithSurfaceContainerHigh sets the high surface container color of the scheme.
func (s *Scheme) WithSurfaceContainerHigh(surfaceContainerHigh int) *Scheme {
	s.SurfaceContainerHigh = s.nrgba(surfaceContainerHigh)
	return s
}

// WithSurfaceContainerHighest sets the highest surface container color of the scheme.
func (s *Scheme) WithSurfaceContainerHighest(surfaceContainerHighest int) *Scheme {
	s.SurfaceContainerHighest = s.nrgba(surfaceContainerHighest)
	return s
}

// WithSurfaceVariant sets the surface variant color of the scheme.
func (s *Scheme) WithSurfaceVariant(surfaceVariant int) *Scheme {
	s.SurfaceVariant = s.nrgba(surfaceVariant)
	return s
}

// WithOnSurface sets the on-surface color of the scheme.
func (s *Scheme) WithOnSurface(onSurface int) *Scheme {
	s.OnSurface = s.nrgba(onSurface)
	return s
}

// WithOnSurfaceVariant sets the on-surface variant color of the scheme.
func (s *Scheme) WithOnSurfaceVariant(onSurfaceVariant int) *Scheme {
	s.OnSurfaceVariant = s.nrgba(onSurfaceVariant)
	return s
}

// WithInverseSurface sets the inverse surface color of the scheme.
func (s *Scheme) WithInverseSurface(inverseSurface int) *Scheme {
	s.InverseSurface = s.nrgba(inverseSurface)
	return s
}

// WithInverseOnSurface sets the inverse on-surface color of the scheme.
func (s *Scheme) WithInverseOnSurface(inverseOnSurface int) *Scheme {
	s.InverseOnSurface = s.nrgba(inverseOnSurface)
	return s
}

// WithBackground sets the background color of the scheme.
func (s *Scheme) WithBackground(background int) *Scheme {
	s.Background = s.nrgba(background)
	return s
}

// WithOnBackground sets the on-background color of the scheme.
func (s *Scheme) WithOnBackground(onBackground int) *Scheme {
	s.OnBackground = s.nrgba(onBackground)
	return s
}

// WithOutline sets the outline color of the scheme.
func (s *Scheme) WithOutline(outline int) *Scheme {
	s.Outline = s.nrgba(outline)
	return s
}

// WithOutlineVariant sets the outline variant color of the scheme.
func (s *Scheme) WithOutlineVariant(outlineVariant int) *Scheme {
	s.OutlineVariant = s.nrgba(outlineVariant)
	return s
}

// WithShadow sets the shadow color of the scheme.
func (s *Scheme) WithShadow(shadow int) *Scheme {
	s.Shadow = s.nrgba(shadow)
	return s
}

// WithShadowTint sets the shadow tint color of the scheme.
func (s *Scheme) WithShadowTint(shadowTint int) *Scheme {
	s.ShadowTint = s.nrgba(shadowTint)
	return s
}

// WithScrim sets the scrim color of the scheme.
func (s *Scheme) WithScrim(scrim int) *Scheme {
	s.Scrim = s.nrgba(scrim)
	return s
}

// nrgba converts an ARGB color to a color.NRGBA.
func (s *Scheme) nrgba(argb int) color.NRGBA {
	return color.NRGBA{
		R: uint8(argb >> 16),
		G: uint8(argb >> 8),
		B: uint8(argb),
		A: uint8(argb >> 24),
	}
}
