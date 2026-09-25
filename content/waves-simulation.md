+++
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_waves1d" title="Basic 1D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Wave
	    sim.Config.Size.Set(80, 1, 1)
    },
    func(sim *wavesim.Sim) { // init
	    wavesim.WavePacket(sim)
	})
```

<div>

This simulation runs the 1D wave equation starting with a moving wave packet initial state.

</div>
