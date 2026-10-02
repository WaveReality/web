+++
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_waves1d" title="Basic 1D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Wave
		sim.Params.ThreeD.SetBool(false)
        sim.Params.C = 1.0
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

This simulation runs the 1D wave equation. The goal is to understand how the standard second-order [[wave]] equation arises as a function of the difference between each cell and its neighbors (i.e., the **spatial gradient**). This spatial gradient drives the second-order **acceleration** of the velocity at each cell.

The default configuration starts with a moving wave packet initial state. 

</div>
