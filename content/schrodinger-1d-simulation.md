+++
Name = "Schrodinger 1D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_kg1d" title="Schrodinger in 1D" collapsed="true"}
```Goal
wavesim.Embed(b,
	func(sim *wavesim.Sim) { // config
        sim.Config.GPU = true
		sim.Params.C = 0.5
		sim.Params.Hbar = 0.5
		sim.Params.ThreeD.SetBool(false)
		sim.Config.Equation = wavesim.Schrodinger
		sim.Config.Size.Set(500, 1, 1)
		sim.ViewInit(wavesim.Cab1DViewAll)
		sim.SchrodingerStats()
	},
	func(sim *wavesim.Sim) { // init
		sim.MovingWavePacketConfig(wavesim.CabPosA, wavesim.CabPosB, math32.X, math32.Vec3(-1, -1, -1), -1, 0, 1)
	})
```

<div>

This simulation runs the [[Schrodinger]] equation in 1D. See also [[Schrodinger 3D Simulation]] for the 3D version.

</div>
