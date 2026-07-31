+++
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_waves1d" title="Basic 1D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
	func(sim *wavesim.Sim) { // config
        sim.Config.GPU = true
		sim.Params.C = 1
		sim.Params.ThreeD.SetBool(false)
		sim.Config.Equation = wavesim.Wave
		sim.Config.Size.Set(80, 1, 1)
		sim.ViewInit(wavesim.Wave1DViewAll)
		sim.WaveStats()
	},
	func(sim *wavesim.Sim) { // init
		sim.PosWavePacket(wavesim.WavePos, math32.X, math32.Vec3(-1, -1, -1), -1, 8, 8, 0, 1.5)
	})
```

<div>

This simulation runs the 1D wave equation starting with a moving wave packet initial state.

</div>
