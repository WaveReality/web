// Copyright (c) 2026, The WaveReality Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sims has the simulations that the pages of the site embed, which
// js builds load only when the first one is shown, as Embed is in the
// Web.DeferredFunctions of core.toml.
package sims

import (
	"maps"
	"slices"

	"cogentcore.org/core/tree"
	"github.com/WaveReality/waves/wavesim"
)

// Embed embeds the simulation of the given name in the given parent,
// returning false if there is no simulation of that name.
func Embed(parent tree.Node, name string) bool {
	em, ok := sims[name]
	if ok {
		em(parent)
	}
	return ok
}

func std1D(sim *wavesim.Sim) {
	sim.Params.ThreeD.SetBool(false)
	sim.ViewInitFunc = func(view *wavesim.View) {
		wavesim.ViewInitFour(view)
		wavesim.ViewInitBars1D(view)
	}
	sim.Config.Size.Set(128, 1, 1)
}

func std3D(sim *wavesim.Sim) {
	sim.Params.ThreeD.SetBool(true)
	sim.ViewInitFunc = func(view *wavesim.View) {
		wavesim.ViewInitFour(view)
	}
	sim.Config.Size.Set(64, 64, 64)
}

// sims are the functions that embed each simulation, by name.
var sims = map[string]func(tree.Node){
	"wave1d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Wave
			sim.Params.C = 1.0
			std1D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"wave3d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Wave
			sim.Config.PacketSlab = false
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"maxwell": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Maxwell
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"kg3d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.KleinGordon
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"kg1d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.KleinGordon
			sim.Params.C = 1.0
			std1D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"kgc": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.KleinGordonC
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"schrodinger3d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Schrodinger
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"schrodinger1d": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Schrodinger
			std1D(sim)
			// sim.Config.Size.Set(80, 1, 1)
		},
			func(sim *wavesim.Sim) {})
	},
	"weyl": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Weyl
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"dirac": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Dirac
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
	"electroweak": func(b tree.Node) {
		wavesim.Embed(b, func(sim *wavesim.Sim) {
			sim.Config.Equation = wavesim.Electroweak
			std3D(sim)
		},
			func(sim *wavesim.Sim) {})
	},
}

// Names returns the names of the simulations.
func Names() []string {
	return slices.Sorted(maps.Keys(sims))
}
