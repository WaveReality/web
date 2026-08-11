+++
Categories = ["Standard Model"]
Name = "Complex KG"
bibfile = "mechphys.json"
+++

<!--- todo: [[@MallickChandrashekar16]] -->

In [[Klein-Gordon]] we saw that a major problem with the scalar KG equation is that it doesn't represent any kind of conserved value: you cannot compute some constant, unchanging value from the $\varphi$ state variables under this equation. This is a problem if you want to develop a probabilistic interpretation of the wave, as in [[Schrodinger]]'s equation for the standard [[Copenhagen]] interpretation. But it is also a problem for other interpretations as well.

In particular, the KG waves seem to actually represent **waves of charge** (i.e., [[matter wave]]s), because charge is also strictly conserved, and it comes in both positive and negative varieties, which the KG waves produce. By contrast, the Schrödinger wave only represents a positive-valued quantity, which fits better with the probabilistic interpretation. Indeed, the authors that do write extensively about the KG equation adopt this interpretation ([[@Greiner00]]; [[@Gingrich04]]; [[@MandlShaw13]]), and this idea was pursued initially by Schrödinger in 1926.

Furthermore, we will see that this charge interpretation fits naturally with the coupling of this KG equation to the electromagnetic (EM) field ([[Maxwell]]'s equations), where the conserved charge value acts just like the electric charge in driving the field. However, to fully accomplish this coupling in a physically accurate way, the wave function needs [[spin]], which is what the [[Dirac]] wave function adds on top of the KG waves. This coupling of charge waves and EM waves has been pursued more recently in neoclassical self-coupled field theory ([[@JaynesCummings63]]; [[@CrispJaynes69]]; [[@BarutVanHuele85]]; [[@BarutDowling90]]; [[@Crisp96]]; [[@FinsterSmollerYau99a]]; [[@Radford03]]; [[@MasielloDeumensOhrn05]]).

The first step toward a more complete wave function for something like an [[electron]] is to introduce a wave state with [[complex number]]s ($\phi$) instead of scalars ($\varphi$), which then supports the computation of a conserved quantity across the two complex state values. When we translate this complex wave function into two separate second-order wave equations without an $i$ imaginary number factor, the KG equation is identical to computing two separate KG equations on each of the two scalar values represented by the complex state variable (i.e., $\varphi_a$ and $\varphi_b$). Note that this was not true of Schrödinger's wave equation, which is first order and has an $i$ term that causes the $a$ and $b$ terms to intermix as the wave unfolds.

{id="eq_kg-complex" title="Klein-Gordon on complex state"}
$$
\frac{\partial^2 \phi}{\partial t^2} = (\nabla^2 - m_0^2) \phi
$$

Because differentiation operates independently on the two separate scalar variables in a complex number, this is equivalent to two parallel scalar KG equations, which we can write as:

$$d
\frac{\partial^2 \varphi_a}{\partial t^2} = \left( \nabla^2 - m_0^2 \right) \varphi_a
$$

$$
\frac{\partial^2 \varphi_b}{\partial t^2} = \left( \nabla^2 - m_0^2 \right) \varphi_b
$$

where again the $\varphi_a$ indicates a scalar state variable representing the real $a$ component of $\phi$, and $\varphi_b$ represents the imaginary $b$ value.

Ok, so how do you get a charge out of that, so to speak? As with Schrödinger's equation, the procedure involves multiplying by the complex conjugate ($\phi^* = \varphi_a - i \varphi_b$), which generally produces the overall magnitude or length of the vector represented by the two components of the complex number: $\varphi_a^2 + \varphi_b^2$.

See this page for all of the details regarding the [[conservation]] properties of wave functions. In summary, if you compute the sum of the complex conjugate across all of space (actually an integral, using continuous equations), and set it equal to zero (so that it never changes), you end up with an expression for the density and motion (current) of a quantity that is conserved (i.e., the charge).

The resulting expression for computing the density of charge (typically written as $\rho$, which is the Greek letter "rho"), which is to say, the amount of charge per cubic state unit, is:

{id="eq_kg-charge" title="Klein-Gordon charge density"}
$$
\rho \equiv \frac{i \hbar e}{2m_0c^2} \left( \phi^* \frac{\partial \phi}{\partial t} - \phi \frac{\partial \phi^*}{\partial t} \right)
$$

where the $e$ value is a constant (1.6e-19 in SI units) that converts natural units into proper units of charge.

This can be expressed in terms of the underlying scalar state variables that make up the complex state ($\phi = \varphi_a + i \varphi_b$), and their first temporal derivatives ($\dot \varphi_a$ and $\dot \varphi_b$), and using natural units where $c=\hbar=1$, as:

$$
\rho_i = \frac{e}{m_0} ({\varphi_b}_i \dot {\varphi_a}_i - {\varphi_a}_i \dot {\varphi_b}_i)
$$

This is directly computable for each cubic cell $i$ in the system.

It becomes clear when explicitly written out in this manner that charge represents a coupling of the two otherwise independently-updated variables in the complex number, and this suggests why a single scalar number cannot represent a conserved charge value.

{id="sim_cc" title="Complex charge" collapsed="true"}
```Goal
ai := 1.0
bphase := 90.0
c := 1.0
mass := 0.1
hbar := 0.5
e := 1.0
csq := c * c
mcOverHSq := (mass * mass * csq) / (hbar * hbar)
heOver2mCSq := (hbar * e) / (2.0 * mass * csq)

var bphaseStr, massStr, hbarStr, msgStr string

##
totalTime := 100
sa := zeros(totalTime)
sb := zeros(totalTime)
da := zeros(totalTime)
db := zeros(totalTime)
chg := zeros(totalTime)
##

func valUpdate() {
    mcOverHSq = (mass * mass * csq) / (hbar * hbar)
    heOver2mCSq = (hbar * e) / (2.0 * mass * csq)
    bphaseStr = fmt.Sprintf("b phase: %4.0f", bphase)
    massStr = fmt.Sprintf("mass: %4.1f", mass)
    hbarStr = fmt.Sprintf("hbar: %4.1f", hbar)
    bi := ai * float64(math32.Cos(math32.DegToRad(float32(-bphase))))
    bvi := math.Sqrt(mcOverHSq) * ai * float64(math32.Sin(math32.DegToRad(float32(-bphase))))
    ##
    mf := array(mcOverHSq)
    cf := array(heOver2mCSq)
    ap := array(ai)
    bp := array(bi)
    av := array(0.0)
    bv := array(bvi)
    ##
    for t := range 100 {
        ##
        sa[t] = ap
        sb[t] = bp
        da[t] = av
        db[t] = bv
        chg[t] = cf * (bp * av - ap * bv)
        
        av -= mf * ap
        ap += av
        
        bv -= mf * bp
        bp += bv
        ##
    }
    msgStr = fmt.Sprintf("<b>Charge: %7.3g </b>", chg.Float(99))
}

valUpdate()

plotStyler := func(s *plot.Style) {
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
}
plot.SetStyler(sa, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
al := plots.NewLine(fig1, sa)
dal := plots.NewLine(fig1, da)
bl := plots.NewLine(fig1, sb)
dbl := plots.NewLine(fig1, db)
chgl := plots.NewLine(fig1, chg)
fig1.Legend.Add("a", al)
fig1.Legend.Add("av", dal)
fig1.Legend.Add("b", bl)
fig1.Legend.Add("bv", dbl)
fig1.Legend.Add("chg", chgl)

msgTx := core.NewText(b)
msgTx.Styler(func(s *styles.Style) {
    s.Min.X.Ch(80) // clean rendering with variable width content
})
core.Bind(&msgStr, msgTx)

func updt() {
    valUpdate()
    al.SetData(sa)
    dal.SetData(da)
    bl.SetData(sb)
    dbl.SetData(db)
    chgl.SetData(chg)
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

addSlider(&bphaseStr, &bphase, -180, 180)
addSlider(&massStr, &mass, 0.1, 1.0)
addSlider(&hbarStr, &hbar, 0.1, 1.0)
```

[[#sim_cc]] demonstrates how this works, in terms of two simple [[harmonic oscillator]] variables _a_ and _b_, which are set to be a specific phase apart from each other (+90 degrees shifts _b_ to the _left_ (earlier) relative to _a_, while -90 shifts to the right, due to the trigonometric convention of 0 degrees being at 1,0 and proceeding counter-clockwise from there). Regardless of the phase relationship, the computed charge value remains constant across the cycles of oscillation. However, critically, the value of the charge is directly a function of this phase relationship, with a maximum of 0.5 when the _b_ value is +90 degrees in relation to the _a_ value, and a minimum of -0.5 for -90 degrees, and zero for 0 or 180 degrees. These relationships are fairly obvious once you appreciate the relationship between velocity and position for each of the variables (which are 90 degrees out of phase with each other, always), and how they enter into the charge equation.

Critically, complex numbers are _always_ 90 degrees out of phase with each other by the very nature of the complex plane. Thus, even though the separate real-valued wave functions are independently updated, it is critical that these two wave states are _initialized_ with the 90 degree phase relationship appropriate for complex numbers, which will then determine the sign of the charge value represented.

Perhaps the most important feature of this charge equation is that it can be either positive or negative, as a function of the phase relationship. This is not true of the corresponding expression for Schrödinger's equation, which is "definitely positive", or, in mathematical terminology, "positive definite". This is one of the major reasons why standard quantum physics has strongly embraced Schrödinger's equation, and not KG: KG does not fit with the standard probabilistic framework, where the wave describes a probability, and a probability is always positive.

Interestingly, the Dirac equation, which standard physics has adopted as a model of the electron (and we'll cover later), also produces negative "probabilities", but these have been (correctly, in our framework) reinterpreted as representing antiparticles (i.e., particles with an opposite charge). The antiparticle of the electron is the **positron**, and it is just like an electron, except it has the opposite charge. Historically, this antiparticle nature of the Dirac equation was regarded as a major problem, until positrons were subsequently discovered, and then Dirac looked like a genius for having made such a bold prediction. Nevertheless, there seems to be some residual discomfort in all this, and many treatments of quantum electrodynamics marginalize the Dirac equation in favor of a largely particle-based treatment. We return to these issues later.

We can also derive an expression for the motion of charge over space, which is the _charge current density_ $\vec{J}$:

{id="eq_kg-current" title="Klein-Gordon current density"}
$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right)
$$

In terms of the underlying scalar state variables (and again for natural units), this is:

$$
\vec{J} = \frac{e}{m_0} (\varphi_a \vec{\nabla} \varphi_b - \varphi_b \vec{\nabla} \varphi_a)
$$

This value indicates how much charge is moving in each of the three different coordinate directions; the $\vec{}$ symbol on top of $\vec{J}$ indicates that this is a vector, containing a separate real scalar value for each direction: $\vec{J} = (J_x, J_y, J_z)$. As noted earlier the $\vec{\nabla}$ symbol is the vector gradient operator, which computes the rate of change of the values along each dimension:

$$
\vec{\nabla} \equiv \left(\frac{\partial}{\partial x}, \frac{\partial}{\partial y}, \frac{\partial }{\partial z}\right)
$$

{id="figure_vgrad" style="height:20em"}
![The gradient of a scalar field ($\vec{\nabla}$), which produces a vector field describing the slope at each point in space. These gradient vectors point in the direction of maximum "downhill" slope. In this example the scalar field is a circularly-symmetric bump.](media/fig_vgrad.jpg)

This just means that this $\vec{\nabla}$ takes a three-dimensional field, in this case the field of our wave value $\varphi_a$ or $\varphi_b$ distributed over space, and computes how steeply this field is changing in each of the three different directions [[#figure_vgrad]]. If we assume that this value is actually computed in our model, then we'll need a way of computing the gradient $\vec{\nabla}$ in discrete space-time. This is covered in the next section.

Before proceeding, we look ahead to the next major development. We have ways of computing the density and current of charge ($\rho$, $\vec{J}$), which drive the electromagnetic field. Thus, we need to think of these variables as physically real values, which can be computed directly from the underlying wave state variables, that give rise to long-range electromagnetic forces, through which our particle-waves interact. The next step is to see how the electromagnetic fields can push our particle waves around.

See [[discrete gradient]] for how to compute the gradient in the discrete CA framework.

So, in the end, the computation of the current, which is a vector having three separate components ($J_x, J_y, J_z$), looks like this:

$$
J_x = \frac{e}{m_0} \left[ \varphi_a \left( \sum_{j \in N_{X}} k_j {(\varphi_b}_{j+} - {\varphi_b}_{j-}) \right) - \varphi_b \left( \sum_{j \in N_{X}} k_j ({\varphi_a}_{j+} - {\varphi_a}_{j-}) \right) \right]
$$

$$
J_y = \frac{e}{m_0} \left[ \varphi_a \left( \sum_{j \in N_{Y}} k_j ({\varphi_b}_{j+} - {\varphi_b}_{j-}) \right) - \varphi_b \left( \sum_{j \in N_{Y}} k_j ({\varphi_a}_{j+} - {\varphi_a}_{j-}) \right) \right]
$$

$$
J_z = \frac{e}{m_0} \left[ \varphi_a \left( \sum_{j \in N_{Z}} k_j ({\varphi_b}_{j+} - {\varphi_b}_{j-}) \right) - \varphi_b \left( \sum_{j \in N_{Z}} k_j ({\varphi_a}_{j+} - {\varphi_a}_{j-}) \right) \right]
$$

Again, it does not look as simple as before, but nevertheless it is necessary to have a current to be able to drive the magnetic field in an manner consistent with known physics. Specifically, the electromagnetic field equations require both $\rho$ and $\vec{J}$ values as their sources (see [[Maxwell]]). In addition, this gradient operation is necessary for several other computations in our model, so, like the laplacian, it can be thought of as one of just a few basic operations that take place over the neighborhood of cells.

## Minimal Coupling of Charge Waves with Electromagnetic Fields

At this point, we have a charged wave that can generate an electromagnetic field according to the charge density $\rho$ and current density $\vec{J}$, and we know how this electromagnetic field propagates according to wave equations. Now, we need to have that electromagnetic field interact with the charge wave to produce actual forces on our model. This occurs by introducing new terms into the complex KG wave equation, which, intuitively speaking, act as external driving forces on this charge wave, in much the same way that the charge and current act as driving forces on the electromagnetic wave equations.

In the electromagnetic field equations, the driving force from charge $\rho$ adds into the second-order temporal derivative $\frac{\partial^2 {}}{\partial t^2}$ ([[maxwell#eq_scalar-pot-chg]] in [[Maxwell]]):

$$
\frac{\partial^2 {A_0}}{\partial t^2} = \nabla^2 A_0 + \frac{1}{\epsilon_0} \rho
$$

However, in Schrödinger's equation, external forces enter as a potential ($V$), in the first-order derivative  $\frac{\partial {}}{\partial t}$:

$$
i \hbar \frac{\partial {\phi}}{\partial t} = \frac{\hbar^2}{2m} \nabla^2 \phi + V \phi
$$

This makes sense, because force is the derivative of a potential, so potential is a first-order factor, and force is a second-order factor.

Our KG (Klein-Gordon) charge wave equation is a second-order equation, expressed in terms of $\frac{\partial^2 {}}{\partial t^2}$, and therefore we need to include external driving forces, not potentials. However, for various reasons, it is necessary to derive such an equation starting from the potential. To do this, we can re-derive a second-order wave equation by replacing the first-order derivative with the first-order derivative with the external driving potential:

$$
\partial^\mu \rightarrow \partial^\mu - \frac{e}{c} {A}^\mu
$$

We can then substitute this first-order four-vector derivative into the four-vector version of the KG wave equation (in natural units), which is ([[klein-gordon#eq_kg-4vec]]):

$$
\partial_\mu \partial^\mu \phi = - m_0^2 \phi
$$

because, as we've noted before:

$$
\partial_\mu \partial^\mu = \frac{\partial^2 {}}{\partial t^2} - \nabla^2
$$

So the compact form of the KG wave equation with minimal coupling is therefore:

$$
\left( \partial_\mu - e {A}_\mu \right) \left(\partial^\mu - e {A}^\mu \right) \phi = -m_0^2 \phi
$$

To get all the units right, and perhaps add some conceptual clarity, we can do the same thing with the four-momentum version of the wave equation, which is:

$$
\hat{p}^\mu \hat{p}_\mu \phi = m_0^2 c^2 \phi
$$

where the momentum operator is essentially just the four-derivative, plus the pesky $i$ and $\hbar$ factors:

$$
\hat{p}^\mu = i \hbar \partial^\mu
$$

So now we can say that the electromagnetic potential pushes directly on the momentum of the wave-particle:

$$
\hat{p}^\mu \rightarrow i \hbar \partial^\mu - \frac{e}{c} {A}^\mu
$$

and the same goes for the covariant forms:

$$
\hat{p}_\mu \rightarrow i \hbar \partial_\mu - \frac{e}{c} {A}_\mu
$$

This notion of the potential pushing directly on the momentum of the particle hopefully makes good intuitive sense, even if all the associated mathematics does not. In any case, the resulting KG wave equation becomes:

$$
\left({\hat{p}^\mu} - \frac{e}{c}{A}^\mu \right) \left({\hat{p}_\mu} - \frac{e}{c}{A}_\mu \right)\phi = m_0^2 c^2 \phi
$$

which can also just be written more compactly as a squared expression:

$$
\left(i \hbar \partial_\mu - \frac{e}{c}{A}_\mu \right)^2\phi = m_0^2 c^2 \phi
$$

When this equation is crunched through to produce separate time and space derivatives (as detailed in a subsequent section), we get a standard second-order wave update equation plus a few extra terms (in natural units):

$$
\frac{\partial^2 {\phi}}{\partial t^2} = \nabla^2 \phi - m_0^2\phi - 2 i e \left(A_0 \frac{\partial
 {\phi}}{\partial t} + \vec{A} \cdot \vec{\nabla} \phi \right) + e^2 \phi \left(A_0^2 - \vec{A}^2\right)
$$

Including all of the various constants, it is:

$$
\frac{\partial^2 {\phi}}{\partial t^2} = c^2 \left( \nabla^2 - \frac{m_0^2 c^2}{\hbar^2} \right) \phi -
 \frac{2 i e}{\hbar} \left(A_0 \frac{\partial {\phi}}{\partial t} + c \vec{A} \cdot \vec{\nabla} \phi \right) + \frac{e^2 \phi}{\hbar^2} \left(A_0^2 - \vec{A}^2\right)
$$

This equation amounts to the basic KG wave equation, plus terms that involve the interaction between the charge wave and the electromagnetic field potentials $A_0$ and $\vec{A}$. For example, in the second term of this equation, the vector potential $\vec{A}$ "pushes" on the gradient of the wave function $\vec{\nabla} \phi$, and the scalar potential $A_0$ pushes on the temporal derivative $\frac{\partial {\phi}}{\partial t}$. Notice that these interactions are all first-order, in terms of the potentials and first-order derivatives of the wave equations. The second-order electromagnetic force fields $\vec{E}$ and $\vec{B}$ do not appear at all! This is despite the fact that these are widely regarded as the primary observables of electromagnetic force. Also, the electromagnetic terms introduce a coupling between the two components of the complex variable $\phi$, because of the presence of the imaginary number $i$ in this term. Therefore, it is only the "free" particle (without electromagnetic forces) that has the completely uncoupled scalar components.

Before we pause to reflect more, we need to take two more steps. First, because $\phi$ is a complex variable, we need to further compute the separate real-valued update equations for $\varphi_a$ and $\varphi_b$ to simulate this in our model. Second, we need to update the charge and current equations for this new version of the KG wave equations.

The resulting CA model update equations (including all the relevant $c$ and $\hbar$ factors) are:

$$
\ddot \varphi_a^{t+1} = c^2 \left(\nabla^2 \varphi_a - \frac{m_0^2c^2}{\hbar^2} \varphi_a \right) + \frac{2 e}{\hbar} \left(A_0 \dot \varphi_b + c \vec{A} \cdot \vec{\nabla} \varphi_b \right) + \frac{e^2}{\hbar^2} \varphi_a \left(A_0^2 - \vec{A}^2 \right)
$$

and:

$$
\ddot \varphi_b^{t+1} = c^2 \left(\nabla^2 \varphi_b - \frac{m_0^2c^2}{\hbar^2} \varphi_b \right) - \frac{2 e}{\hbar} \left(A_0 \dot \varphi_a + c \vec{A} \cdot \vec{\nabla} \varphi_a \right) + \frac{e^2}{\hbar^2} \varphi_b \left(A_0^2 - \vec{A}^2 \right)
$$

(where for simplicity the right hand side variables are implicitly taken at time $t$ for cell $i$, and the discrete versions of $\nabla^2$ and $\vec{\nabla}$ presented earlier are assumed). Note that we again need the first-order spatial gradient operator $\vec{\nabla}$ as a basic computation in our model, but otherwise all the variables are local to the system.

The [[conservation|conserved]] charge and current values computed by this equation must also be updated. The coupling with the electromagnetic field has introduced additional factors here, which depend on the electromagnetic potentials $A_0$ and $\vec{A}$. For the charge density $\rho$, we have:

$$
\rho = \frac{i \hbar e}{2m_0c^2} \left( \phi^* \frac{\partial {\phi}}{\partial t} - \phi \frac{\partial {\phi^*}}{\partial t} \right) - \frac{e^2}{m_0 c^2} A_0 \phi \phi^*
$$

If you compare with the original charge equation for the complex KG equation ([[#eq_kg-charge]]), it is the same except for the last term. Similarly, the current density $\vec{J}$ is:

$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right) - \frac{e^2}{m_0 c} \vec{A} \phi \phi^*
$$

again with an extra term at the end relative to [[#eq_kg-current]].

These equations need to be converted into computational expressions in terms of the separate complex components, as before:

$$
\rho = \frac{\hbar e}{m_0c^2} (\varphi_b \dot \varphi_a - \varphi_a \dot \varphi_b) - \frac{e^2}{m_0 c^2} A_0 (\varphi_a^2 + \varphi_b^2)
$$

$$
\vec{J} = \frac{\hbar e}{m_0} (\varphi_a \vec{\nabla} \varphi_b - \varphi_b \vec{\nabla} \varphi_a) - \frac{e^2}{m_0c} \vec{A} (\varphi_a^2 + \varphi_b^2)
$$

At this point, we have reached an important milestone --- if you take the equations just presented above, this describes a particle as a distributed wave of charge that gets pushed around by the electromagnetic field potentials ($A_0$ and $\vec{A}$). Furthermore, this wave of charge produces electromagnetic fields, in terms of charge and current densities $\rho$ and $\vec{J}$. Thus, we finally have a complete system of equations that can potentially simulate charged particles whizzing around and interacting with each other. In other words, we finally have the potential to make direct contact with observable physics! Indeed, you can explore the behavior of this system in the model, by using the Complex Coupled KG wave equations setting.

## Numerical Issues with Coupling: Symmetry Breaking

When you actually simulate these equations on the computer, something very interesting (and initially distressing) happens --- they blow up! As you run the equations over time in the presence of a fixed electromagnetic field, the total charge value, far from being a constant, increases steadily, and eventually you end up with numbers approaching infinity. This is not because the math is wrong (after very thorough checking!), but because of the coupling between the two elements of the complex variable that occurs in the update equation. Specifically, the update of $\varphi_a$ depends on $\dot \varphi_b$ and $\vec{\nabla} \varphi_b$, and vice-versa. This interdependency creates numerical instabilities when we simply substitute in the discrete computed values at each time step. In particular, because the change in $\varphi_a$ depends on the change in $\varphi_b$, this cycle of dependency can get out of whack.

Intuitively, the electromagnetic potentials drive a rotation through the $\varphi_a$ and $\varphi_b$ variables, which is evident in the fact that they subtract from $\varphi_b$ but add to $\varphi_a$ --- these opposite signs are the signature of a rotation (and incidentally are caused by the presence of the imaginary $i$ numbers in the equations). To the extent that the potentials are pushing the $\varphi_a$ variable up, there should be an equal and opposite pushing of the $\varphi_b$ variable down, causing the rotation. However, if the $\varphi_b$ variable only has the "old" data from the previous time step about how much $\varphi_a$ got pushed up, then it doesn't compensate correctly in how much it gets pushed down. Thus, you end up with a "leak" in the system, where instead of rotating nicely in place, the system starts to fly out of control, spinning wider and wider circles each time.

The solution to this problem is to *break the symmetry* between $\varphi_a$ and $\varphi_b$ in these update equations. Instead of updating each of them at the same time, based on the prior values of the other, we choose one variable ($\varphi_a$, arbitrarily) and update its values first. Then, when we compute $\varphi_b$, we use the _current value_ of $\dot \varphi_a$ in the update equation for $\varphi_b$. This prevents the rotation between these variables from getting out of whack, and restores numerical stability to the system.

<!--- todo: run KG invr5 case, plot figure! -->

`\begin{figure}`
` \centering\includegraphics[height=2in]{fig.dirac_invr5_50.eps}   \caption{\small Total charge for updating of the charge wave     equation over 100,000 time steps in the presence of a fixed $1/r$ potential with magnitude .5, in a universe of $50^3$ cubes.  The`  
`   instantaneous total charge varies considerably over time, but the average     across time is constant, demonstrating that total charge is conserved on     average, but not at each moment.  If no potential is present, then charge     is identically conserved from one time step to the next.}`  
` \label{fig.kg_invr5_50} \end{figure}`

However, if you run this equation in a static electromagnetic field, it is clear that the total amount of charge in the model at any one time changes over time (Figure~\ref{fig.kg_invr5_50}). This is flies in the face of the fancy math that says that these equations conserve charge! Somewhat amazingly, however, if you run the system long enough, it becomes clear that the _average_ amount of charge never changes.

Overall, we have now happened upon a very interesting situation. The breaking of symmetry between the two variables in the complex wave state, forced upon us by implementational considerations, actually fits at least qualitatively with a known and otherwise very puzzling property of physics. The weak force also breaks symmetry in a very similar way: there is a preferred direction of rotation in the weak force. Although it is not yet clear (to me at least) that this preferred rotational direction in the weak force maps identically onto this preferred rotation direction, it is nevertheless a tantalizing possibility. The weak force has been integrated with the electromagnetic force, as the _electroweak_ force --- it is possible that the effects described by the electroweak force correspond in some way to the oscillations in charge value that are observed in our model. We will return to these issues later.

Meanwhile, we nee need to introduce just a bit more complexity into our KG wave equation before we have a fully satisfactory model of a fundamental particle of nature: the electron (and its antiparticle, the positron). This extra bit of complexity extends the phenomenon of rotation that we've just been discussing, to account for the strange quantum mechanical property of **spin**. The resulting equation goes by the name of the second-order Dirac equation. Once we have that, we will have a complete system that, if all the math is correct, should make direct and numerically accurate contact with observable phenomena!

Before moving on to spin, there are two remaining loose ends for the present set of equations. First, there is an interesting interpretation of the way that the electromagnetic potential interacts with our charge wave, called **local gauge invariance**, which provides a template for exploring the other two forces of nature: the weak and strong forces (which we will not cover further in this paper). Second, we have the actual mechanics of deriving the above equations.

## Local Gauge Invariance

If you look at it in the right way, the electromagnetic field can be seen as a way of canceling out an extra degree of freedom present in the complex KG wave field equations. As discussed in [[Maxwell]], the Lorenz gauge is an example of a situation where we introduced some extra constraints on the field variables, and this eliminated a degree of freedom in the electromagnetic equations, and also made them appear a lot simpler than they otherwise would.

This notion of a _gauge_ is very general: it just means that whenever you have some unconstrained values in your equations (i.e., values that can change without changing the observable results that you can measure in physics experiments), then you need to apply some kind of gauge to fix these variables. In the Lorenz gauge example, the extra degree of freedom comes from the fact that the observable variables are the force vectors $\vec{E}$ and $\vec{B}$, which are essentially derivatives of the underlying potentials $A_0$, $\vec{A}$. Thus, the raw values of the potentials can be moved up or down, and it won't change the slope of the fields, and therefore it won't change the observable force vectors.

However, it is not clear exactly how to reconcile such an argument with the fact that the potentials appear directly in our coupling equations, and also are observable in terms of the Arahnov-Bohm effect, as discussed elsewhere.

Nevertheless, here is the argument for the electromagnetic coupling being a form of local gauge invariance. If you just take our basic complex KG wave equation, you can get exactly the same overall behavior if you multiply the thing by a "phase transformation" (it is often said that gauge invariance should really be phase invariance) which is basically just a rotation along the complex axes. This is exactly the kind of rotation discussed above. The generic form of a rotation in complex numbers is to multiply by an exponential term:

$$
\phi(x) \rightarrow \exp \left(\frac{ie}{\hbar c} \chi \right) \phi(x)
$$

Where the $\chi$ term is the amount that you're rotating (the rest are just convenient constants for getting $\chi$ into the right units) --- think of it as some number of degrees of rotating the underlying $\varphi_a$ value into $\varphi_b$ (and vice-versa) for the complex number $\phi$.

If $\chi$ is independent of location ($x$), then it is just a constant and nothing happens. This is a *global* gauge/phase transformation, and it is not very interesting. However, if $\chi$ is now itself a function of location (i.e., $\chi(x)$ ), this is a *local* gauge transformation, and this is where the electromagnetic coupling comes in. If you have such a local variable, and you take the derivative of the resulting overall system that includes this locally-varying thing, you get this extra term for the derivative of $\chi(x)$ with respect to x:

$$
\partial_\mu \phi(x) \rightarrow \exp \left( \frac{ie}{\hbar c} \chi(x) \right) \left(\partial_\mu + \frac{ie}{\hbar c} \partial_\mu \chi(x) \right) \phi(x)
$$

So now your nice wave equation is a mess, and it varies from one place to another as a function of this $\partial_\mu \chi(x)$ term. So here is the trick: you basically just add this annoying term into the overall EM potential field (which is OK because such additions do not change the gradients and therefore do not affect observable EM field vectors):

$$
{A}_\mu(x) \rightarrow {A}_\mu(x) + \partial_\mu \chi(x)
$$

And then you just subtract this whole thing back out from your messy equation, and this gives you something just slightly less messy:

$$
\left(i \hbar \partial_\mu - \frac{e}{c} {A}_\mu\right)\phi
$$

So, this ends up being the same thing as the minimal coupling described earlier. Somehow, this whole process seems like a rather contrived way to end up with something that already made quite a bit of sense before hand. However, as noted earlier, the true payoff of such a procedure appears to come in addressing the weak and strong forces, which we leave for a future refinement of the model.

## Summary

We now have a full electrodynamic system with bidirectional interactions between a wave of charge and the electromagnetic force field. The wave equation remains at the core of both the charge wave and the electromagnetic field equations, but we did have to add in a few extra first-order spatial gradient computations here and there. All of this was necessary to get our electrodynamics working according to known physical laws.

It may not be quite as elegant as our simplest system of a pure wave field for forces, and a simple scalar KG wave equation for our particle, but elegance cannot override physical facts. Still, all the broad implications of those core wave equations (special relativity, Newtonian-like dynamics, quantum physics, etc) all hold for our more elaborated equations.

Perhaps the most surprising development here is that we had to break the symmetry between the two components of the complex state variable, and introduce a preferred direction of rotation, in order to avoid numerical instability. This numerical instability results when variables are coupled, and the rotation through the $\varphi_a$ and $\varphi_b$ variables produced by the interaction with the electromagnetic field is manifest as such a coupling. This break in symmetry and preferred direction of rotation bears a tantalizing resemblance to features of the weak force, suggesting a possible explanation for an otherwise very strange aspect of nature.

Our overall model at this point consists of a small handful of locally computable equations, which can be readily simulated on a computer. As these equations play out, they deterministically, locally, and automatically generate physics that should be largely consistent with what we know about the world. However, as we noted before, our equations are missing one critical piece, which is _spin_.

Interestingly, the introduction of spin in the [[Dirac]] function creates a parity in the number of variables participating in the electromagnetic field equations (four) with those in the particle charge wave, which is currently two, but will now double to four. Overall, the elegance of having four variables each in two systems of interacting wave equations, which unfold in a four-dimensional space-time, seems suspiciously neat.

