// Copyright (c) 2026, The WaveReality Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/text/csl"
	_ "cogentcore.org/core/text/tex" // include this to get math
	"cogentcore.org/core/tree"
	_ "cogentcore.org/lab/yaegilab"
)

// NOTE: you must make a symbolic link to the zotero MechPhys CSL file as mechphys.json
// in this directory, to generate references and have the generated reference links
// use the official APA style. https://www.zotero.org/groups/2525742/mechphys
// Must configure using BetterBibTeX for zotero: https://retorque.re/zotero-better-bibtex/
// todo: include link for configuring here

//go:generate mdcite -vv -refs ./mechphys.json -d ./content

//go:embed content citedrefs.json
var econtent embed.FS

//go:embed icon.svg
var icon string

func main() {
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
				ctx.LinkButton(w, "https://github.com/WaveReality/web")
				w.SetText("GitHub").SetIcon(icons.GitHub)
			})
		})
	})

	b.RunMainWindow()
}
