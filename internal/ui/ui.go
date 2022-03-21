// Copyright 2022 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2/internal/atlas"
	"github.com/hajimehoshi/ebiten/v2/internal/mipmap"
)

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)

type TouchID int

// RegularTermination represents a regular termination.
// Run can return this error, and if this error is received,
// the game loop should be terminated as soon as possible.
var RegularTermination = errors.New("regular termination")

type FPSModeType int

const (
	FPSModeVsyncOn FPSModeType = iota
	FPSModeVsyncOffMaximum
	FPSModeVsyncOffMinimum
)

type CursorMode int

const (
	CursorModeVisible CursorMode = iota
	CursorModeHidden
	CursorModeCaptured
)

type CursorShape int

const (
	CursorShapeDefault CursorShape = iota
	CursorShapeText
	CursorShapeCrosshair
	CursorShapePointer
	CursorShapeEWResize
	CursorShapeNSResize
)

type WindowResizingMode int

const (
	WindowResizingModeDisabled WindowResizingMode = iota
	WindowResizingModeOnlyFullscreenEnabled
	WindowResizingModeEnabled
)

func (u *UserInterface) imageAt(mipmap *mipmap.Mipmap, x, y int) (r, g, b, a byte, err error) {
	return mipmap.At(graphicsDriver(), x, y)
}

func (u *UserInterface) dumpScreenshot(mipmap *mipmap.Mipmap, name string, blackbg bool) error {
	return mipmap.DumpScreenshot(graphicsDriver(), name, blackbg)
}

func (u *UserInterface) dumpImages(dir string) error {
	return atlas.DumpImages(graphicsDriver(), dir)
}
