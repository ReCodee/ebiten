// Copyright 2022 The Ebitengine Authors
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
// limitations under the License

package ebitenutil_test

import (
	"embed"
	"testing"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed *.png
var images embed.FS

func TestNewImageFromFileSystem(t *testing.T) {
	img, _, err := ebitenutil.NewImageFromFileSystem(images, "text.png")
	if err != nil {
		t.Fatal(err)
	}
	w, h := img.Size()
	if got, want := w, 192; got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
	if got, want := h, 128; got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
