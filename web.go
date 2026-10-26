// Copyright (c) 2026, The WaveReality Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/text/csl"
	"cogentcore.org/core/text/tex/texcache"
	"cogentcore.org/core/tree"
	"github.com/WaveReality/waves/mvmwaves"
)

// NOTE: you must make a symbolic link to the zotero MechPhys CSL file as mechphys.json
// in this directory, to generate references and have the generated reference links
// use the official APA style. https://www.zotero.org/groups/2525742/mechphys
// Must configure using BetterBibTeX for zotero: https://retorque.re/zotero-better-bibtex/
// todo: include link for configuring here

//go:generate mdcite -vv -refs ./mechphys.json -d ./content
//go:generate go run ./genmath.go

//go:embed content citedrefs.json
var econtent embed.FS

//go:embed mathcache.json.gz
var mathcache embed.FS

//go:embed icon.svg
var icon string

// simNames are the names of the simulations of the sims package, which
// the pages embed as <sim-name> elements.
var simNames = []string{"wave1d", "wave3d", "maxwell", "kg3d", "kg1d", "kgc",
	"schrodinger3d", "schrodinger1d", "weyl", "dirac", "electroweak"}

// embedSim is [sims.Embed], which js builds load only when the first
// simulation is shown (see [Web] DeferredFunctions in core.toml). It is
// only set on non-generatehtml platforms so that xyz does not interfere
// with html generation.
var embedSim func(parent tree.Node, name string) bool

func main() {
	// The code blocks of the content are compiled into the app by the core
	// tool (contentcode_gen.go, see [ContentCode] in core.toml). Edited code
	// runs in the mvm interpreter, which js builds load when code is first
	// edited (see [Web] DeferredFunctions in core.toml).
	content.BindTextEditor = mvmwaves.BindTextEditor
	texcache.OpenFS(mathcache, "mathcache.json.gz")
	texcache.SetShapeMath()
	core.AppIcon = icon
	content.Settings.SiteTitle = "Wave Reality"
	content.OfflineURL = "https://wavereality.org"
	b := core.NewBody(content.Settings.SiteTitle)
	ct := content.NewContent(b).SetContent(econtent)
	ctx := ct.Context
	refs, err := csl.OpenFS(econtent, "citedrefs.json")
	if err == nil {
		ct.References = csl.NewKeyList(refs)
	}
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(func(p *tree.Plan) {
			ct.MakeToolbar(p)
			ct.MakeToolbarPDF(p)
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/WaveReality/")
				w.SetText("GitHub").SetIcon(icons.GitHub)
			})
		})
	})

	AddSims(ctx)

	b.RunMainWindow()
}

func AddSims(ctx *htmlcore.Context) {
	if embedSim == nil {
		return
	}
	for _, nm := range simNames {
		snm := "sim-" + nm
		ctx.ElementHandlers[snm] = func(ctx *htmlcore.Context) bool {
			fr := core.NewFrame(ctx.BlockParent)
			fr.SetName(snm)
			fr.Styler(func(s *styles.Style) {
				s.Direction = styles.Column
				s.Grow.Set(1, 1)
			})
			embedSim(fr, nm)
			return true
		}
	}
}
