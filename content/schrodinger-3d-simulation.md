+++
Name = "Schrodinger 3D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_sch3d" title="Schrodinger in 1D" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Schrodinger
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

This simulation runs the [[Schrodinger]] equation in 3D. See also [[Schrodinger 1D Simulation]] for the 1D version.

</div>
