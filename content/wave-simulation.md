+++
Categories = ["Simulations"]
bibfile = "mechphys.json"
+++

{id="sim_wave1d" title="Basic 1D Waves" collapsed="true"}
```Goal
wavesim.Embed(b,
    func(sim *wavesim.Sim) {
	    sim.Config.Equation = wavesim.Wave
		sim.Params.ThreeD.SetBool(false)
        sim.Params.C = 1.0
		sim.ViewInitFunc = func(view *wavesim.View) {
            wavesim.ViewInitFour(view)
            wavesim.ViewInitBars1D(view)
        }
	    sim.Config.Size.Set(128, 1, 1)
    },
    func(sim *wavesim.Sim) { // init
	})
```

<div>

This simulation runs the 1D wave equation. You should be able to see how the standard second-order [[wave]] equation arises as a function of the difference between each cell and its neighbors (i.e., the **spatial gradient**). This spatial gradient drives the second-order **acceleration** of the velocity at each cell.

The default configuration starts with a moving wave packet initial state. To dive right in and see the wave in action, you can click the [[#sim_wave1d:Step 1]] button to update the velocity from the spatial gradient force in each cell, and then update the position from that velocity. Use the other Step buttons or `Run` and `Stop` to watch things unfold. 

The [[#sim_wave1d:Init]] button restarts everything. In general it is a good idea to do this after changing parameters, although parameter updates do take effect on the next Run or Step action. Often, the initial configuration depends on the parameters, so it might not work properly unless you Init.

The four different panels are showing the current and previous state **Pos** = position values (on the lower row), and the current and previous **Vel** velocity values on the upper row.

You can use the Zoom button (with a magnifying-glass icon, sort of like this: 🔍) on the lower toolbar to zoom the range of displayed values, so you can see more clearly how the current and previous values are offset by one state cell on the horizontal axis, and likewise the velocity values are 90 degrees out of phase with the position values, such that the most rapidly-moving values are those that have the smallest position values.

The speed-of-light factor `C` (in the `Params` section on the left panel) is set to 1 -- you can also set that to lower values and see what effect that has, on slowing the wave propagation rate.

The [[#sim_wave1d:Config]] menu allows you to select different initial configurations -- try the `Wave Pulse` option instead of the wave packet. It generates a wave blob that sloshes around in the space.

You can view the [[#sim_wave1d:Root Stats Plot]] to see various summary statistics as the simulation runs. This includes the different components of energy (`Kinetic`, `Potential`, and the sum as `Energy`), along with stats that track the velocity and width of the wave. Mouse over each label to see what it is recording.

The [[#sim_wave1d:Panel]] selector in the View toolbar allows you to move through the 4 different panels, going in a left-right, bottom-top order. You can see the variable update as you move through, and this allows you to change which variable (or any other view setting) for that panel.

## Edges

The [[#sim_wave1d:Edges]] setting in `Params` controls what happens at the edges, as explained in [[wave#Dealing with edges]]. Explore these settings so you can see what they do.

## Superposition

Select the `Symmetric Packet` configuration in [[#sim_wave1d:Config]], which is like `Wave Packet` except it doesn't set the velocity values, so the wave packet ends up going in both directions. This provides a good way to see the superposition phenomenon, where the two waves will wrap around (if `Edges` = `Wrap`) or bounce back (for `Fixed`) and when they meet again, at specific steps, the `Pos` values will shrink significantly, and overall it is hard to tell just by looking at the wave state that there would be two waves moving in opposite directions embedded in that one state.

This happens maximally at step 126, so you can get there quickly by doing [[#sim_wave1d:Step 100]] and then two [[#sim_wave1d:Step 10]] and 6 [[#sim_wave1d:Step 1]].

If you just Run the sim and watch it run, you can see that the wave equation is linear, so that the waves just pass right through each other instead of somehow bouncing off of each other.

## Potential interactions

The `Oscillator` configuration creates a potential "well" that traps the wave state. For this to work in this 1D case, you need to set the `C` speed-of-light to 0.25. Press `Init` after updating the parameters. 

Click on the `V` variable to see the resulting potential well: this is strongly negative in the surrounding region, and goes up to 0 in the middle. Go ahead and `Step` / `Run` the model and see that the wave is now trapped in this little well, and it just oscillates back and forth. This is the basis of the _quantum harmonic oscillator_ that we'll see in other configurations where it behaves a bit more cleanly.

</div>
