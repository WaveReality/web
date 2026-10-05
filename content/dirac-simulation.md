+++
Name = "Dirac Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_dirac" title="Dirac in 3D" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Dirac
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

This simulation runs the [[Dirac]] equation in 3D.

This builds on the [[complex KG simulation]], with more complete spin-based coupling with the [[Maxwell]] EM field. Use the [[#sim_dirac:Root Stats Plot]] to view the `Charge` value and other key stats.

The single most important qualitative phenomenon to observe is how spin in this set of four second-order wave equations emerges only through the coupling with the EM field, and is otherwise not present in the free electron. You can observe this by comparing `Spin at Rest` with `Spin Precession` as described below: in the first case, there is no EM field, and the `A2s` wave remains inactive (it is initialized to 0). In the latter case, there is an EM field, and the `A2s` spins relative to the `A1s` etc.

The [[#sim_dirac:Config]] initial configurations are as follows:

* `Spin at Rest:` A lump of charge at rest with spin along Z; nothing happens to the spin, because a free particle has no effective spin, which only emerges through coupling with EM field.

* `Spin Packet:` A travelling spin-1/2 packet along X (as in other wave packet cases), showing the effects of spin and mass relative to basic wave.

* `Spin in Potential:` An electron lump offset from a fixed 1/r well, pulled in by it: the external-field case, with SelfField off so A0 stays as set.

* `Spin Precession:` Spin along X in a uniform B magnetic field along Z: it precesses at the Larmor rate with g = 2; look at the `SigX` and `SigY` values in the stats plot.

* `Dirac Hydrogen...` simulates an electron in an `S` or `P` orbital configuration within the hydrogen atom, where the positive charge provides the attractive force.

* `Dirac Oscillator:` simulates the quantum harmonic oscillator using Dirac waves.

* `Chiral Oscillation:` A lump started purely right-chiral: the mass turns it into the left one and back at m c^2 / hbar, which relates to the Weyl version where mass mediates conversion back and forth.

</div>
