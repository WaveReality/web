+++
Name = "KG 1D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_kg1d" title="Klein-Gordon in 1D" collapsed="true"}
```Goal
wavesim.Embed(b,
	func(sim *wavesim.Sim) { // config
        sim.Config.GPU = true
		sim.Params.C = 0.5
		sim.Params.ThreeD.SetBool(false)
		sim.Config.Equation = wavesim.KleinGordon
		sim.Config.Size.Set(500, 1, 1)
		sim.ViewInit(wavesim.Wave1DViewAll)
		sim.WaveStats()
	},
	func(sim *wavesim.Sim) { // init
		sim.MovingWavePacketConfig(wavesim.WavePos, wavesim.WaveVel, math32.X, math32.Vec3(-1, -1, -1), -1, 0, 1)
	})
```

<div>

This simulation runs the [[Klein-Gordon]] equation in 1D. See also [[KG 3D Simulation]] for the 3D version.

</div>
