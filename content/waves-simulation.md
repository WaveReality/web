+++
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_waves1d" title="Basic 1D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
	func(sim *wavesim.Sim) { // config
        sim.Config.GPU = true
		sim.Config.Equation = wavesim.Wave1D
		sim.Config.Size.Set(80, 1, 1)
		sim.ViewInit(wavesim.Wave1DViewAll)
	},
	func(sim *wavesim.Sim) { // init
		sim.MovingWavePacket(wavesim.WavePos, math32.X, math32.Vec3i(40, 0, 0), -1, 8, 8, 0, 1.5)
	})
```

<div>

This simulation runs the 1D wave equation starting with a moving wave packet initial state.

</div>
