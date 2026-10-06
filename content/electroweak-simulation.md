+++
Name = "Electroweak Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_ew" title="Electroweak" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Electroweak
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

This simulation runs the full [[weak#electroweak]] system, including the [[Higgs]] field.

* `HS0a` is the real value (`a`) of the un-charged component of the Higgs doublet, which is where the vacuum expectation value ends up going after the symmetry breaking. This is the component of the field that is termed the "Higgs boson" field. The others end up becoming effectively part of the $W^\pm$ and $Z^0$ boson fields, due to the way that the coupling with the other gauge fields works.

Use the [[#sim_ew:Root Stats Plot]] to view key stats.

Use the [[#sim_ew:Rescale]] button to rescale the wave display to the current range of the wave state values being displayed.

The [[#sim_ew:Config]] initial configurations are as follows:

* `Higgs Broken:` Broken symmetry with Higgs field (`Hs0a`) at vacuum expectation value. Not much happens here, except that the vacuum expectation value remains steady in its local attractor state.

* `Higgs Symmetric:` Higgs starts at zero plus noise and falls off the top of the Mexican hat. The magnitude of the Higgs fields (`Hmag`) increases due to the coupling with the gauge field activity. Because the system is tiny and closed, it all just ends up bouncing around and not settling out to a steady expectation value, but you can see the basic principle at work.

* `EM Pulse:` Transverse EM pulse travelling along X at exactly C; Note that `BYs` and `W3Ys` are _aligned_ and thus cancel out the mass factor that arises due to the self-interaction term -- you can see the net effect in the `AYs` field which integrates the BYs and W3Ys components according to the weak mixing angle factors. But really, the work has already been done by virtue of the fact that, independently, the `B` and `W3` fields are experiencing no mass themselves, and thus moving at C. 

* `Z Pulse:` Z boson pulse along X at about 0.75 C, slower than light because the `BYs` and `W3Ys` are at 180 degree opposite phases, so they do not cancel out, and thus do experience a mass factor, by contrast with the EM pulse. You can switch the Panel 1 to view `ZY` to see the resulting Z mixture of these two source fields. Note how the mixing angle ends up rectifying the phase misalignment in the sources. In addition, the `Hs0a` Higgs boson field gets a big "dent" in it due to the coupling of the Z boson with this field, whereas the EM pulse had absolutely no effect. This dent is much larger than it would be in reality, due to the way the parameters are set to make an easily-observable slowdown due to the mass, given the small scale of this simulation.

* `Lepton Packets:` A neutrino and an electron as the same wave in the two halves of one doublet: the Higgs gives one of them a mass and not the other, due to the coupling dynamics. Set the Panel 0 to view `Nu1a` and Panel 1 to `El1a` to watch them propagate, and observe that the neutrino moves faster.

* `W Collision:` Two W packets cross and generate longitudinal W3 where they overlap, demonstrating the non-Abelian W coupling in action; watch `W3Xs` and `ZX`.

</div>
