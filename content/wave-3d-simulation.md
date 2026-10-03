+++
Name = "Wave 3D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_wave3d" title="3D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Wave
		sim.Config.PacketSlab = false
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

This simulation runs the 3D version of the wave equation. See [[wave simulation]] for the 1D version which has all the introductory overview for the simulator.

The key point here is to see that the 3D discrete Laplacian operator produces perfectly isotropic and smooth waves in three dimensions. Due to technical issues with this Laplacian operator (i.e., its _spectral radius_), the speed-of-light must be set slower, so it is at the default of 0.25 used for most of the other simulations.

The 3D version of the display shows a single slice through the 3D state, in the X (horizontal) by Z (depth) plane, at a given Y axis slice level. You can use the arrow buttons in the bottom view toolbar to move around in the space (up and down the Y axis), and, if you have zoomed in, you can also move the X axis position back and forth.

You can just explore all of the same [[#sim_wave3d:Config]] options that you looked at in the 1D version. The `Oscillator` case in particular is a bit more interesting looking here (and doesn't require any different parameters).

The `PacketSlab` option in `Config` determines whether the wave packets are generated as an entire consistent "slab" across the Z (depth) axis, or whether that axis also has a Gaussian envelope applied, as the X axis does. This Gausssian envelope will allow you to observe the spreading of the wave over time, while the slab case only exhibits dispersion in the X axis.

For the `SymmetricPacket` case showing superposition, this is best with `PacketSlab = true`, and occurs at Step 271 with the default C = 0.25 and 64 cube size.

</div>
