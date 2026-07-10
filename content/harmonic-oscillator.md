+++
bibfile = "mechphys.json"
+++

The **simple harmonic oscillator** (SHO) captures the core oscillatory behavior of [[waves]], without any spatial dimensions to bother with. It can be seen as the 0-dimensional version of a wave, where the force that drives the oscillation comes not from neighbors, but from the position (height) of the wave itself. As such, it provides a potentially interesting role in the mechanics of [[stochastic particles]] because it can be entirely localized to one discrete grid cell within the [[cellular automaton]] framework. Thus, a particle in this view can be considered to be a simple harmonic oscillator that periodically jumps between cells.

The basic equations from [[waves]] for the SHO are:

{id="eq_force" title="restoring force"}
$$
f = -c^2 y^t
$$

where $c$ is the basic rate update constant, analogous to the speed of light in waves, which determines the effective strength of the restoring force, and thus the oscillation rate. The _t_ suffix indicates the time step (only for variables that require integration over time). Everything else from this point onward is the same, in the basic Newtonian physics framework of acceleration, velocity, and position:

{id="eq_a" title="acceleration"}
$$
a = \frac{f}{m}
$$

{id="eq_" title="new velocity"}
$$
v^{t+1} = v^t + a
$$

{id="eq_" title="new state"}
$$
y^{t+1} = y^t + v^{t+1}
$$

{id="sim_sho" title="Simple harmonic oscillator" collapsed="true"}
```Goal
ip := 1.0
c := 0.2
mass := 0.5
csq := c * c

var massStr, cStr, msgStr string

##
totalTime := 100
sp := zeros(totalTime)
sv := zeros(totalTime)
pE := zeros(totalTime)
kE := zeros(totalTime)
tE := zeros(totalTime)
##

func valUpdate() {
    massStr = fmt.Sprintf("mass: %4.1f", mass)
    cStr = fmt.Sprintf("c: %4.1f", c)
    csq = c * c
    ##
    mf := array(mass)
    cf := array(csq)
    p := array(ip)
    pp := array(ip)
    pv := array(0.0)
    mv := array(0.0)
    v := array(0.0)
    pot := array(0.0)
    kin := array(0,0)
    ##
    for t := range 100 {
        ##
        pv = 1.0 * v 
        pp = 1.0 * p
        v = pv - (cf * pp) / mf
        p = pp + v
        
        mv = 0.5 * (v + pv) // midway
        pot = 0.5 * pp * pp
        kin = (mv * mv * mf) / (2.0 * cf)
        
        sp[t] = p
        sv[t] = v
        pE[t] = pot
        kE[t] = kin
        tE[t] = pot + kin
        ##
    }
    msgStr = fmt.Sprintf("<b>Total E: %7.3g </b>", tE.Float(99))
}

valUpdate()

plotStyler := func(s *plot.Style) {
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
}
plot.SetStyler(sp, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
pl := plots.NewLine(fig1, sp)
vl := plots.NewLine(fig1, sv)
pEl := plots.NewLine(fig1, pE)
kEl := plots.NewLine(fig1, kE)
tEl := plots.NewLine(fig1, tE)
fig1.Legend.Add("p", pl)
fig1.Legend.Add("v", vl)
fig1.Legend.Add("pE", pEl)
fig1.Legend.Add("kE", kEl)
fig1.Legend.Add("tE", tEl)

msgTx := core.NewText(b)
msgTx.Styler(func(s *styles.Style) {
    s.Min.X.Ch(80) // clean rendering with variable width content
})
core.Bind(&msgStr, msgTx)

func updt() {
    valUpdate()
    pl.SetData(sp)
    vl.SetData(sv)
    pEl.SetData(pE)
    kEl.SetData(kE)
    tEl.SetData(tE)
    msgTx.UpdateRender()
    pw.NeedsRender()
}

func addSlider(label *string, val *float64, mnVal, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
    core.Bind(label, tx)
	sld := core.NewSlider(b).SetMin(mnVal).SetMax(mxVal).SetEnforceStep(true)
    if mxVal > 10 {
        sld.SetStep(10)
    } else {
        sld.SetStep(0.1)
    }
	sld.SendChangeOnInput()
	sld.OnChange(func(e events.Event) {
		updt()
		tx.UpdateRender()
	})
	core.Bind(val, sld)
}

addSlider(&massStr, &mass, 0.1, 1.0)
addSlider(&cStr, &c, 0.1, 1.0)
```

The equivalent of the wavelength for the SHO is the _period_, in time, for the cycle to repeat.

