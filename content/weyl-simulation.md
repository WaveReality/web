+++
Name = "Weyl Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_weyl" title="Weyl in 3D" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Weyl
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

This simulation runs the [[Weyl]] equation in 3D.

Overall, you can observe that it is rather "jittery" compared to the smooth nature of the second-order [[Dirac]] equation, and prone to generating high-frequency noisy modes. This is because of its first-order nature, despite using appropriate numerical integration techniques to deal with the potential issues.

The [[#sim_weyl:Config]] initial configurations are as follows:

* `Neutrino Packet:` A massless left-handed helicity wave packet along X, with no paired right-handed half to couple with, so it travels at light-speed c.

* `Neutrino Doubler:` A buggy case showing lattice effects by placing alternating signs from one cell to the next, which causes it to move backward instead of forward.

* `Electron Packet:` A travelling spin-1/2 packet along X with both left and right-handed helicity coupled through the mass term, which causes wave energy to move _between_ the two sides, thereby slowing both down. This is how the same wave equation can represent a massless neutrino (uncoupled, one helicity only) and a massive electron (coupled, both helicities).

* `Electron at Rest:` ?

* `Electron Chiral Flip:` An electron starting with everything in the left-hand side, but mass causes it to move entirely to the right-hand side, and then back again. Mega [[zitterbewegung]].

* `Electron in Field:` Adds a uniform electric field, which drives momentum, except if WeylQ is 0, in which case it is a neutrino that ignores the field.

* `Weyl Oscillator:` simulates the quantum harmonic oscillator using Weyl waves.

* `Weyl Hydrogen...` simulates an electron in an `S` orbital configuration within the hydrogen atom, where the positive charge provides the attractive force.

</div>
