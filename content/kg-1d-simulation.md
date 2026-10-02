+++
Name = "KG 1D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_kg1d" title="Klein-Gordon in 1D" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.KleinGordon
		sim.Params.ThreeD.SetBool(false)
		sim.ViewInitFunc = func(view *wavesim.View) {
            wavesim.ViewInitFour(view)
            wavesim.ViewInitBars1D(view)
        }
	    sim.Config.Size.Set(80, 1, 1)
    },
    func(sim *wavesim.Sim) { // init
	})
```


<div>

This simulation runs the [[Klein-Gordon]] equation in 1D. See also [[KG 3D Simulation]] for the 3D version.

</div>
