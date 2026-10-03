+++
Name = "KG 3D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_kg3d" title="Klein-Gordon in 3D" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.KleinGordon
		sim.Params.ThreeD.SetBool(true)
		sim.ViewInitFunc = func(view *wavesim.View) {
            wavesim.ViewInitFour(view)
        }
	    sim.Config.Size.Set(64, 64, 64)
    },
    func(sim *wavesim.Sim) { // init
	})
```

<div>

This simulation runs the [[Klein-Gordon]] equation in 3D. See also [[KG 1D Simulation]] for the 1D version.

As with the 1D version, the goal here is to explore the effects of the `Mass` parameter on wave propagation, this time in 3D. The `Oscillator` configuration is interesting to watch in this case.

</div>
