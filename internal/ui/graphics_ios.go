// Copyright 2019 The Ebiten Authors
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

//go:build ios && !ebitengl && !ebitencbackend
// +build ios,!ebitengl,!ebitencbackend

package ui

import (
	"fmt"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver/metal"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver/metal/mtl"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver/opengl"
)

func Graphics() graphicsdriver.Graphics {
	if runtime.GOARCH == "386" || runtime.GOARCH == "amd64" {
		return opengl.Get()
	}

	if _, err := mtl.CreateSystemDefaultDevice(); err != nil {
		panic(fmt.Sprintf("mobile: mtl.CreateSystemDefaultDevice failed on iOS: %v", err))
	}
	return metal.Get()
}
