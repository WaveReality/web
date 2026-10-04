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

This builds on the [[complex KG simulation]], with more complete spin-based coupling with the [[Maxwell]] EM field. Use the [[#sim_dirac:Root Stats Plot]] to view the `Charge` value.

The [[#sim_dirac:Config]] initial configurations are as follows:

TODO: set Mu0 lower! also in KGC

* `Spin at Rest:` shows the conservation of charge (view the plot), as Gaussian wave blob oscillates and disperses over time. 

* `Spin Packet:` shows charge generated from a moving wave packet. Run after Charge Self Field to get the self-field effects as well.

* `Charge at Rest Anti:` has the opposite charge value.

* `Charge Uniform:` has a uniform charge distribution.

* `Charge Self Field:` turns on the `EM` and `SelfField` flags, and sets `Mu0` to a small value, so you can see the Charge value computed from the KG wave actually drive the EM `A` four-potential, which in turn then feeds back and couples with the KG field. You can select these variables to view (use the `Vectors` view for `AXs` for example, and use the [[#sim_dirac:Rescale]] button to auto-scale the range).

* `Scalar Hydrogen...` simulates an electron in an `S` or `P` orbital configuration within the hydrogen atom.

* `Scalar Oscillator:` simulates the quantum harmonic oscillator using KG complex waves.

</div>
