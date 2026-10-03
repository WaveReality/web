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

Relative to the [[wave simulation]], the primary difference here is the effect of `Mass` (adjustable in the `Params` fields on the left) on the rate of propagation of the wave. You can try changing this Mass parameter, hitting `Init`, and then `Step 100` to go a fixed number of steps with different Mass values. You should see that the wave packet travels slower as the mass increases.

</div>
