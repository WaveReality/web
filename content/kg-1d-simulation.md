+++
Name = "KG 1D Simulation"
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

<sim-kg1d>

<div>

This simulation runs the [[Klein-Gordon]] equation in 1D. See also [[KG 3D Simulation]] for the 3D version.

Relative to the [[wave simulation]], the primary difference here is the effect of `Mass` (adjustable in the `Params` fields on the left) on the rate of propagation of the wave. You can try changing this Mass parameter, hitting `Init`, and then `Step 100` to go a fixed number of steps with different Mass values. You should see that the wave packet travels slower as the mass increases.

You can try all the other [[#sim_kg1d:Config]] initial configurations, which are the same as those available in the [[wave simulation]], to see how the additional mass factor alters the behavior of these cases.

</div>
