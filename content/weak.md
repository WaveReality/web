+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

The **weak** force on its own is responsible for various forms of _decay_ or _transmutation_ from one type of particle to another, most notably in the case of _beta decay_ where for example a proton decays into a neutron within the nucleus. This is a relatively rare event, and the weak force is, after all, "weak", so in general it perhaps doesn't get as much attention as it otherwise might.

However, the **electroweak** unification of the electromagnetic force (i.e., [[Maxwell]]'s equations) together with the weak force, along with the central role of the [[Higgs]] field in this framework, provides a profoundly different picture of the fundamental nature of an [[electron]] and its interaction with the electromagnetic field. This electroweak picture also integrates the [[neutrino]]s with their charged [[lepton]] partners like the electron, and is intimately connected with the _flavors_ or [[generation]]s of particles, e.g., muon and tau.

Thus, it is perhaps not an exaggeration to say that the electroweak framework is as revolutionary for understanding the nature of leptons as the [[quark]] framework is for understanding the hadrons: it provides an entirely different fundamental basis space, that parsimoniously integrates a wide range of disparate phenomena into a more coherent framework. To understand the most basic nature of the electron and electromagnetic interactions, understanding the electroweak model seems essential.

However, the counter-argument is that the actual physical implications of this revolutionary rearrangement of the quantum furniture are somewhat difficult to detect, which presumably why this part of the story gets less attention than it otherwise might. The famously successful [[QED]] model, which preceded the electroweak model by more than a decade, persists unchanged!

This fact says a lot about how different [[calculational tool]]s can be used to describe the same phenomena, but if the goal is to understand something about the underlying physical mechanisms, then it seems that there is much to learn from the electroweak framework. In that respect, there are many features of the electroweak framework that align with aspects of the [[Spinfield Model]], so it can thus can provide essential guidance on how that model should be configured.

It is strongly recommended to read [[gauge theory]] and [[Higgs]] before proceeding, as these contain essential background information necessary for understanding what follows. You should understand how gauge theory describes the interaction between particle and force fields, and how this interaction can give rise to a dynamic mass-like factor, which shows up in the conserved current density expression. This is precisely where the mass of the weak bosons will arise in the electroweak framework, for precisely the same reason. 

The role of the Higgs field is, first and foremost, to provide the source of a non-zero "fuel" for the current density, through the mechanism of spontaneous symmetry breaking as a result of the specific shape of the Higgs potential. The specific structure of this field, as a scalar complex field (spin 0, i.e., [[complex KG]]) with two elements, is also critical for how it interacts with the four different force fields present in the electroweak system, to deal with the otherwise problematic Goldstone boson problem and to allow one effective force field to remain massless (which then represents the Maxwell EM field), while three others take on mass terms, which turn out to be consistent with the measured masses of the weak bosons.

## Electroweak

The electroweak framework is based on three fields, which together comprise a total of 4 four-vector potential-like fields, and the Higgs doublet, which has 4 real-valued numbers, organized into two complex values.

* The **Higgs doublet** $\Psi$, which has 2 complex field elements (4 real numbers), one of which interacts with electric charge ($\chi^+$) and another that is electrically neutral ($\chi^0$). As a result of the spontaneous symmetry breaking property of the Higgs potential, the neutral component acquires a stable expected value of $v_h = 246 GeV$, which is what then provides the mass term for all other wave fields.

* The **weak hypercharge** four-potential $B_\mu$, which is essentially just like the EM four-potential $A_\mu$, with four real-valued numbers, $B_0$ is the scalar electric potential, and the remaining three terms ($B_{(x,y,z)}$) are the vector potential. The source of this potential is the weak hypercharge value $Y_W$ which acts like the $Q$ charge for the EM potential. It interacts with the Higgs doublet through the same local gauge invariance mechanism (see [[gauge theory]]) that describes minimal coupling with the EM field, exactly as done in the case of the [[complex KG]] equation.

* The **weak isospin** triplet of four-vectors, $W^a_\mu$ where $a=1,2,3$, and each such four-vector is mostly like the EM four-potential $A_\mu$. The first two of these four-vectors, $W^1_\mu$ and $W^2_\mu$ together end up as the components of the $W^\pm$ bosons, while the third one $W^3_\mu$ contributes to both the final electromagnetic field potential ($A^\mu$) and the neutral $Z^0$ boson.

These fields are mixed together in specific combinations, using a standard unitary rotation matrix defined by an angle $\theta_w$ known as the Weinberg or _weak mixing_ angle, to obtain the actual _observable_ fields, as follows:

The EM four-potential, which propagates without any mass:

$$
A_\mu = cos \theta_w B_\mu + sin \theta_w W^3_\mu
$$

The four-potential associated with the weak neutral boson $Z^0$, which has a (large) mass:

$$
Z_\mu = -sin \theta_w B_\mu + cos \theta_w W^3_\mu
$$

And the four-potentials associated with the two charged weak bosons $W^\pm$:

$$
W^+ = \frac{1}{\sqrt2} \left( W^1_\mu - W^2_\mu \right)
$$

$$
W^- = \frac{1}{\sqrt2} \left( W^1_\mu + W^2_\mu \right)
$$

Furthermore, the electric charge $Q$ is also a mixture of two quantum numbers that are associated with different types of particles, _weak hypercharge_ ($Y$) and _weak isospin_ (third component: $T^3$):

$$
Q = Y + T^3
$$

{id="table_weak-qs" title="weak hypercharge and isospin quantum numbers"}
| Fermion                | Y    | T^{1,2} | T^3  | Q  |
|------------------------|------|---------|------|----|
| neutrino (left chiral) | -1/2 | 1/2     | 1/2  | 0  |
| electron, left chiral  | -1/2 | 1/2     | -1/2 | -1 |
| electron, right chiral | -1   | 0       | 0    | -1 |
| Higgs $\chi^+$         | 1/2  | 1/2     | 1/2  | 1  |
| Higgs $\chi^0$         | 1/2  | 1/2     | -1/2 | 0  |

[[#table_weak-qs]] shows these quantum numbers for neutrinos and electrons (the leptons). This table, plus the mixing factors shown above, clearly show that the electroweak framework provides an entirely different factorization of the theoretically fundamental lepton particles. There are two fundamentally different versions of the electron, and the left-chiral electron and the neutrino share most of their quantum numbers, so they would seem to be much more similar at the electroweak level than we would otherwise expect given their overall properties.

Thus, it seems that one's understanding of what an electron really is changes dramatically under this electroweak framework. Furthermore, every one of these electroweak fields, including in addition the particle fields associated with the leptons, is coupling to the very same Higgs doublet field: what are the implications of this mutual interaction through this shared field?

A fundamental question raised by all of this, is _why_ are the observable fields a mixture and not just equivalent to the corresponding weak hypercharge or weak isospin fields that we started with? This is described as being a consequence of the spontaneous symmetry breaking dynamic in the Higgs field, but the story is quite a bit more complicated than that. The symmetry breaking is _necessary_ for these mixture factors to actually matter, because prior to this symmetry breaking, all of the four-potentials are essentially equivalent, acting like the EM massless four-potential. Once the symmetry is broken and the Higgs field has a vacuum expectation value, then all of the differences that were actually built into the system start to matter.

In particular, the gauge theory derivation of how the Higgs field couples to the weak hypercharge and weak isospin fields results in a **covariant derivative** that cancels out the local gauge factors arising from the gauge fields, of the form:

$$
D_\mu = \partial_\mu - i g' Y B_\mu - i g W^a_\mu T^a
$$

Where $g'$ and $g$ are arbitrary real-valued coupling parameters for each of the respective fields ($B_\mu$ and $W^a_\mu$), and $Y$ is the weak hypercharge quantum number value for the Higgs field (a real number), which is set to 1/2 by convention. The $T^a$ is a set of 3 different coupling matricies, one for each of the three $W^a_\mu$ four-vector components, that are defined by the SU(2) unitary rotation group -- i.e., a set of basis vectors that perform unitary rotations of a 2x2 matrix, which is what the Higgs field is (2 complex values for each of 2 doublets).

These are none other than the [[Pauli matricies]], which are also used in the [[Dirac]] equation, and define the property of [[spin]] in the quantum world. This is why the $W$ field is called _isospin_, because it causes spinning. Each of these are also multiplied by the conventional 1/2 factor, but it turns out that the third Pauli matrix has a -1 on the bottom-right diagonal, which therefore gives the neutral Higgs doublet component $\chi^0$ a -1/2 effective $T^3$ quantum number.

This negative quantum number means that the $W^3$ component force field has a negative contribution, while the $B_\mu$ field has a positive contribution, and any field configuration that has positive values in each of these fields, in the correct ratio as defined below, will end up with no mass, due to the cancellation of these positive and negative contributions.

The covariant derivative is used in the [[Lagrangian]] for the Higgs field (see [[gauge theory]] for details on how this is all computed, using a simpler single-valued complex KG particle field):

$$
\mathcal{L}_h = |D_\mu \Psi|^2 + V(\Psi)
$$

where $V(\Psi)$ is the Higgs potential that leads to symmetry breaking, as described in detail in [[Higgs]], and the Higgs field $\Psi$ is the (1/2 normalized) complex doublet:

{id="eq_state" title="Higgs complex doublet"}
$$
\Psi \equiv \frac{1}{\sqrt2} \binom{\chi^+}{\chi^0}
$$

From all of this, we can now see that the _reason_ that electric charge and the electromagnetic field are a mixture of the $B_\mu$ and $W^3$ fields is because of the way that these gauge coupling factors work. In particular, the only way for the EM field to remain fully massless (and thus have long-range effects) is for it to have opposite effects on the current density computed by the gauge theory analysis of the electroweak system, so that these effects _perfectly_ cancel out, leaving zero current, and thus, zero mass. We will see this in detail below.

This cancellation is captured qualitatively by the definition of electric charge in terms of the sum of the factors for each field, and quantitatively it is given by the ratio of the coupling parameters $g'$ and $g$, which determines the weak mixing angle:

$$
\sin \theta_w = \frac{g'}{\sqrt{g^2 + g'^2}}
$$

These coupling parameters are determined experimentally, and are not predicted by the theory itself. Nevertheless, their impacts have far-reaching implications on a range of other parameters and experimental outcomes, all of which have been strongly confirmed. As with everything else in the Standard Model at this point, all of this seemingly crazy complexity has been experimentally validated to high levels of precision, with this electroweak theory making strong predictions for things that had not yet been seen experimentally before.

Furthermore, the gauge theory framework that drives so much of the phenomenology here is fundamentally based on describing the [[conservation]] of energy and charge in the equations of motion, derived from the Lagrangian framework. Thus, ultimately, these conservation dynamics are what is driving everything. And the only way to get the math to all work out was to do it with massless gauge fields that work like the EM four-potential, and then have the mass emerge dynamically from within that fundamentally massless context, which the specific form of the Higgs potential accomplishes.

Thus, even though this seems like a very complicated way of going about everything, in the end each step is strongly constrained, and it is difficult to see how a different kind of system could be derived.

## Phenomenology of the electroweak system

Go to the [[electroweak simulation]] to explore the full electroweak system of coupled fields, with several important demonstrations of its phenomenology.

### The EM field is _defined_ by the ratio of local field strengths

In the following analysis, we see precisely how the current density that is computed for the coupled electroweak system will turn out to be zero whenever the local field values for $B_\mu$ and $W^3$ are in a specific positive-valued ratio. The result is that any waves that have this particular configuration will propagate over a long range, because they are effectively massless, while all the massive configurations will end up decaying over a relatively short range. Thus, the Higgs field acts much like a filter: the only thing that gets through is light itself (EM radiation over the effectively massless field configuration), while everything else is effectively the massive weak bosonic fields $W^\pm$ and $Z^0$.

Everything below concerns only the two neutral gauge fields $W^3_\mu$ and $B_\mu$. The charged $W^{1,2}$ and the Higgs fluctuations play no part. Each gauge field obeys a wave equation driven by a current, and the Higgs doublet is what supplies that current:

{id="eq_eom_gen" title="equations of motion for the neutral gauge fields"}
$$
\square\, W^3_\mu = j^3_\mu, \qquad \square\, B_\mu = j^Y_\mu
$$

{id="eq_currents" title="the currents the Higgs doublet supplies"}
$$
j^3_\mu = 2 g\, \mathrm{Im}\!\left[\Phi^\dagger T^3 D_\mu \Phi\right], \qquad
j^Y_\mu = 2 g' Y\, \mathrm{Im}\!\left[\Phi^\dagger D_\mu \Phi\right]
$$

The key factor in these current equations is the nature of the covariant derivative, which will end up bringing the force fields into the overall current expression, via the _seagull_ term that we saw in the gauge theory derivation.

{id="eq_covd" title="covariant derivative, neutral sector"}
$$
D_\mu \Phi = \left(\partial_\mu - i G_\mu\right)\Phi, \qquad
G_\mu = g W^3_\mu T^3 + g' Y B_\mu
$$

With $T^3 = \mathrm{diag}(+\tfrac12, -\tfrac12)$ and $Y = \tfrac12$, which is a multiple of the identity, $G_\mu$ is diagonal:

{id="eq_gmat" title="the neutral gauge matrix"}
$$
G_\mu = \frac{1}{2}
\begin{pmatrix} g W^3_\mu + g' B_\mu & 0 \\[2pt] 0 & -g W^3_\mu + g' B_\mu \end{pmatrix}
$$

$W^3$ enters with **opposite signs** in the two entries, because $T^3$ distinguishes the components. $B$ enters with the **same** sign in both, because hypercharge does not. That asymmetry is the origin of everything that follows: positive values of $W^3$ will end up canceling out positive values of $B$.

Now use the actual state of the Higgs doublet. It is uniform, so $\partial_\mu \Phi = 0$ and the covariant derivative is *entirely* the gauge term. And it is $\Phi = (0,\ v/\sqrt2)$ -- the upper charged component is exactly zero, while the vacuum expectation value is concentrated in the lower neutral component. Note that this specific configuration is _essential_ for all of the mixing logic described here to actually work: the 0 cancels out possible contributions from the other W components. Thus, the upper entry of $G_\mu$ acts on nothing, due to the 0:

{id="eq_dphi" title="the covariant derivative at the broken Higgs doublet"}
$$
D_\mu \Phi = -i\, G_\mu \Phi
= \left(0,\; \tfrac{i}{2}\left(g W^3_\mu - g' B_\mu\right) \tfrac{v}{\sqrt2}\right)
$$

Only one scalar survives. Give it a name:

{id="eq_sdef" title="the one combination the Higgs doublet responds to"}
$$
S_\mu \;\equiv\; g W^3_\mu - g' B_\mu
$$

Two neutral fields, but the Higgs doublet presents only _one non-zero component_ for them to couple through.

Put that $D_\mu\Phi$ into the currents. Both reduce to the same scalar $S_\mu$, with different prefactors:

{id="eq_eom_mass" title="the currents become mass terms"}
$$
\square\, W^3_\mu = -\frac{v^2}{4}\, g\, S_\mu, \qquad
\square\, B_\mu = +\frac{v^2}{4}\, g'\, S_\mu
$$

The resulting current factors act like a mass, and note that the $S_\mu$ term contains both the $W^3$ and $B$ field values, with opposite signs, so this is why the specific field levels in each field determine the resulting effective mass value at each point. There is no sum over $\mu$ and no magnitude anywhere, which is why it is local to each field component at each point. Each spacetime index $\mu$ -- the scalar potential and the three vector components -- has its own $S_\mu$ and thus its own mass term, computed from the field values at that one lattice site.

Because both equations are driven by the same $S_\mu$, two particular combinations decouple. Take $g$ times the first plus $g'$ times... more simply, form:

{id="eq_pz" title="the massless and massive combinations"}
$$
\begin{aligned}
P_\mu &\equiv g' W^3_\mu + g B_\mu
& \square P_\mu &= -\tfrac{v^2}{4} g g' S_\mu + \tfrac{v^2}{4} g g' S_\mu = 0 \\[4pt]
Z_\mu &\equiv g W^3_\mu - g' B_\mu
& \square Z_\mu &= -\tfrac{v^2}{4}\left(g^2 + g'^2\right) S_\mu
\end{aligned}
$$

$P_\mu$ is the photon (EM field), up to normalization: dividing by $\sqrt{g^2+g'^2}$ turns $(g', g)$ into $(\sin\theta_W, \cos\theta_W). $Z_\mu$ is the $Z$, likewise up to normalization, and it obeys a massive wave equation with $M_Z^2 = \tfrac{v^2}{4}(g^2+g'^2)$.

In summary, when you work through the final bit of math, there are two specific combinations of wave state magnitudes, one that results in zero effective mass (the photon) and one that has mass (the Z boson).

{id="eq_configs" title="the two pulse configurations"}
$$
\begin{aligned}
\text{photon:}\quad (W^3, B) &\propto (g',\, g)
& S &= g g' - g' g = 0 \\[2pt]
\text{Z:}\quad (W^3, B) &\propto (g,\, -g')
& S &= g^2 + g'^2 \neq 0
\end{aligned}
$$

By virtue of the filtering argument given above, nothing specifically requires that precisely the right ratio of field values must be present, because anything that is not well-aligned will be filtered out! For example, a pure $W^3$ pulse produces _both_ photon and Z fields, moving at different speeds.

### Does the electron make only EM field waves?

It does not — and it is worth being clear that this is the gauge coupling's business, not the Yukawa's. The Yukawa term gives the electron a mass; what it couples to is fixed entirely by $T^3$ and $Y$.

Rewrite the fermion's gauge matrix in the photon and $Z$ directions, using $W^3 = \sin\theta_W A + \cos\theta_W Z$ and $B = \cos\theta_W A - \sin\theta_W Z$:

{id="eq_ncoup" title="the neutral couplings of a fermion"}
$$
g W^3 T^3 + g' Y B
\;=\; \underbrace{e\,Q}_{\text{to } A}\, A
\;+\; \underbrace{\sqrt{g^2+g'^2}\left(T^3 - \sin^2\!\theta_W\, Q\right)}_{\text{to } Z}\, Z
$$

with $Q = T^3 + Y$ and $e = g g'/\sqrt{g^2+g'^2}$.

The photon coupling collapses to $e Q$ — that is what $Q = T^3 + Y$ buys, and it is zero for the neutrino. The $Z$ coupling is a *different* combination, and it vanishes for nothing here:

| | $T^3$ | $Q$ | photon, $Q$ | $Z$, $T^3 - \sin^2\!\theta_W Q$ |
|---|---|---|---|---|
| $\nu_L$ | $+\tfrac12$ | $0$ | $0$ | $+0.500$ |
| $e_L$ | $-\tfrac12$ | $-1$ | $-1$ | $-0.277$ |
| $e_R$ | $0$ | $-1$ | $-1$ | $+0.223$ |

So an electron sources both fields at once. Putting a lump of each in an empty box and reading off the fields it creates:

```
electron   max |A0| 1.260e-01    max |Z0| 8.386e-02
neutrino   max |A0| 1.124e-08    max |Z0| 1.514e-01
```

The neutrino makes no photon field and a large $Z$ field; the electron makes both. Their $Z$ fields stand in the ratio $1.805$, against the $0.500 / 0.277 = 1.805$ the couplings predict.

Note also that $e_L$ and $e_R$ have $Z$ couplings of **opposite sign** while their photon couplings are identical. That difference is parity violation, and it is why the weak interaction distinguishes handedness while electromagnetism does not.

### Then why does only electromagnetism reach us?

Because of the mass, not the coupling. A massless field around a static source falls off as $1/r$; a massive one is screened,

{id="eq_yukawa_range" title="range of a massive field"}
$$
\frac{1}{r} \;\longrightarrow\; \frac{e^{-M_Z r}}{r}
$$

so the $Z$ part of an electron's field is gone beyond $r \sim 1/M_Z$, about $2\times10^{-3}$ fm. Every electron is radiating $Z$ field as vigorously as photon field; it simply does not get anywhere.

This closes the loop with the earlier section. A random excitation of $W^3$ and $B$ splits into both modes, and so does the field around an electron. What makes electromagnetism look like the only long-range neutral force is not that anything is emitted selectively, but that one of the two modes has a mass and the other does not.


### The weak force bosons W are very hard to activate

* heavy
* requires precise phase coupling between electron and neutrino -- forbidden!

### What happens to the other 3 components of the Higgs field?

(eaten by W, Z)


## Derivation of gauge coupling

$$
\Phi = \begin{pmatrix}\chi^+\\ \chi^0\end{pmatrix}
$$

$$
T^a = \frac{\sigma^a}{2}
$$

$$
Y_\Phi = +\tfrac12
$$

$$
\mathcal{L} = (\partial_\mu\Phi)^\dagger(\partial^\mu\Phi) - V(\Phi^\dagger\Phi)
$$

Take a **constant** transformation $\Phi \to U\Phi$ with

$$
U = e^{i\alpha^a T^a}\,e^{i\beta Y} \in SU(2)\times U(1)
$$

The potential is fine because $U$ is unitary: $\Phi^\dagger U^\dagger U \Phi = \Phi^\dagger\Phi$. The kinetic term is fine because $U$ passes through the derivative: $\partial_\mu(U\Phi) = U\partial_\mu\Phi$.

Now let $\alpha^a(x)$, $\beta(x)$ vary. The potential still survives — unitarity holds pointwise. But

$$
\partial_\mu\left(U(x)\Phi\right) = U\,\partial_\mu\Phi + \underbrace{(\partial_\mu U)\,\Phi}_{\text{the problem}}
$$

That second term wrecks the kinetic term. Everything that follows exists to cancel it.

Introduce a derivative $D_\mu$ and **demand** that $D_\mu\Phi$ transform exactly like $\Phi$ itself:

$$
\boxed{\;(D_\mu\Phi)' = U\,(D_\mu\Phi)\;}
$$

If that holds, the kinetic term is automatically invariant, since $(D_\mu\Phi)'^\dagger(D^\mu\Phi)' = (D_\mu\Phi)^\dagger U^\dagger U (D^\mu\Phi)$. Write the ansatz

$$
D_\mu = \partial_\mu - i\,\mathcal{G}_\mu
$$

$$
\mathcal{G}_\mu \equiv g\,W^a_\mu T^a + g'\,Y B_\mu
$$

and impose the box. Left side:

$$
\partial_\mu(U\Phi) - i\mathcal{G}'_\mu U\Phi = U\partial_\mu\Phi + (\partial_\mu U)\Phi - i\mathcal{G}'_\mu U\Phi
$$

Right side:

$$
U\partial_\mu\Phi - iU\mathcal{G}_\mu\Phi
$$

Equate, cancel $U\partial_\mu\Phi$, and strip $\Phi$ (it's arbitrary):

$$
(\partial_\mu U) - i\,\mathcal{G}'_\mu U = -i\,U\mathcal{G}_\mu
$$

Right-multiply by $U^\dagger$:

$$
\mathcal{G}'_\mu = U\,\mathcal{G}_\mu\,U^\dagger - i\,(\partial_\mu U)U^\dagger
$$

**That single equation is the entire gauge-field transformation law**, and it was forced on us — we chose nothing.

### Split into the SU(2) and U(1) pieces

Write $U = U_2 U_1$ with $U_2 = e^{i\alpha^aT^a}$ and $U_1 = e^{i\beta Y}$. Since $U_1$ is a phase times the identity it commutes with everything, and

$$
(\partial_\mu U)U^\dagger = (\partial_\mu U_2)U_2^\dagger + i(\partial_\mu\beta)\,Y
$$

Matching the traceless-matrix part against the identity part:

$$
g\,W'^a_\mu T^a = U_2\left(g W^a_\mu T^a\right)U_2^\dagger - i(\partial_\mu U_2)U_2^\dagger
$$

$$
B'_\mu = B_\mu + \frac{1}{g'}\partial_\mu\beta
$$

Put $U_2 \approx 1 + i\alpha^aT^a$. For the homogeneous piece, using $[T^a,T^b] = i\varepsilon^{abc}T^c$:

$$
U_2\left(W^bT^b\right)U_2^\dagger \approx W^bT^b + i\alpha^a W^b\left(i\varepsilon^{abc}T^c\right) = W^bT^b - \varepsilon^{abc}\alpha^a W^b T^c
$$

For the inhomogeneous piece, $-i(\partial_\mu U_2)U_2^\dagger \approx (\partial_\mu\alpha^a)T^a$. Collecting:

$$
\delta W^a_\mu = \frac{1}{g}\partial_\mu\alpha^a + \varepsilon^{abc}W^b_\mu\,\alpha^c, \qquad \delta B_\mu = \frac{1}{g'}\partial_\mu\beta, \qquad \delta\Phi = i\left(\alpha^aT^a + Y\beta\right)\Phi
$$

**Note the structural difference.** $B$ gets only the inhomogeneous shift. $W$ gets that *plus* a homogeneous rotation $\varepsilon^{abc}W^b\alpha^c$ — because $W$ lives in the adjoint representation, i.e. it is charged under its own group. That extra term is the seed of every W self-interaction.

Also note the gauge field is **not a tensor**. The $\partial_\mu\alpha/g$ term is inhomogeneous, which is precisely what makes it a *connection* and what lets it absorb the offending $(\partial_\mu U)\Phi$.

Apply $\delta$ to $D_\mu\Phi = \partial_\mu\Phi - ig W^aT^a\Phi - ig'YB\Phi$, writing $\Lambda \equiv \alpha^aT^a + Y\beta$:

$$
\delta(D_\mu\Phi) = \underbrace{i(\partial_\mu\Lambda)\Phi + i\Lambda\partial_\mu\Phi}_{\text{from }\delta\Phi} \underbrace{- i(\partial_\mu\alpha^a)T^a\Phi - ig\,\varepsilon^{abc}W^b\alpha^c T^a\Phi}_{\text{from }\delta W} \underbrace{+\, gW^aT^a\Lambda\Phi}_{\text{from }\delta\Phi\text{ in the }W\text{ term}} \underbrace{- iY(\partial_\mu\beta)\Phi}_{\text{from }\delta B} + \underbrace{g'YB\Lambda\Phi}_{\text{from }\delta\Phi\text{ in the }B\text{ term}}
$$

**The inhomogeneous terms cancel.** Since $\partial_\mu\Lambda = (\partial_\mu\alpha^a)T^a + Y\partial_\mu\beta$, the $i(\partial_\mu\Lambda)\Phi$ piece exactly kills the $-i(\partial_\mu\alpha^a)T^a\Phi$ and $-iY(\partial_\mu\beta)\Phi$ terms. That's the whole job of the $\partial\alpha/g$ in $\delta W$.

**The commutator terms cancel too.** What's left must equal $i\Lambda(D_\mu\Phi)$, which requires

$$
-ig\,\varepsilon^{abc}W^b\alpha^c T^a \;=\; gW^a\left[\Lambda, T^a\right]
$$

Evaluate the right side: only the $SU(2)$ part of $\Lambda$ contributes, giving $gW^a\alpha^b\left(i\varepsilon^{bac}T^c\right)$. Relabelling both sides so the generator index is $c$, the left gives $-\varepsilon^{cab}W^a\alpha^b$ and the right gives $+\varepsilon^{bac}W^a\alpha^b$. Since $\varepsilon^{cab} = \varepsilon^{abc}$ and $\varepsilon^{bac} = -\varepsilon^{abc}$, both equal $-\varepsilon^{abc}$. They match.

$$
\Longrightarrow \quad \delta(D_\mu\Phi) = i\Lambda\,(D_\mu\Phi) \quad\checkmark
$$

$D_\mu\Phi$ transforms exactly like $\Phi$. The kinetic term $(D_\mu\Phi)^\dagger(D^\mu\Phi)$ is now locally invariant.

$$
\mathcal{G}_\mu = gW^a_\mu T^a + g'Y B_\mu = \frac12\begin{pmatrix} gW^3_\mu + g'B_\mu & g\left(W^1_\mu - iW^2_\mu\right)\\[4pt] g\left(W^1_\mu + iW^2_\mu\right) & -gW^3_\mu + g'B_\mu\end{pmatrix}
$$

Everything from here is reading off entries. Note the $\mp g W^3$ on the diagonal — that sign flip is $T^3 = \pm\tfrac12$, and it is why the two components feel opposite $W^3$ and identical $B$.

### Higgs with the vacuum expectation

Put $\Phi \to \langle\Phi\rangle = (0,\, v/\sqrt2)^T$. Then $\partial_\mu\langle\Phi\rangle = 0$, so only the last term of

$$
|D_\mu\Phi|^2 = (\partial_\mu\Phi)^\dagger(\partial^\mu\Phi) - i(\partial_\mu\Phi)^\dagger\mathcal{G}^\mu\Phi + i\Phi^\dagger\mathcal{G}_\mu(\partial^\mu\Phi) + \Phi^\dagger\mathcal{G}_\mu\mathcal{G}^\mu\Phi
$$

survives. Acting on the VEV picks out the **second column**:

$$
\mathcal{G}_\mu\langle\Phi\rangle = \frac{v}{2\sqrt2}\begin{pmatrix} g\left(W^1_\mu - iW^2_\mu\right) \\[3pt] -gW^3_\mu + g'B_\mu \end{pmatrix}
$$

$$
\left|\mathcal{G}_\mu\langle\Phi\rangle\right|^2 = \frac{v^2}{8}\left[g^2\left(W^1_\mu W^{1\mu} + W^2_\mu W^{2\mu}\right) + \left(gW^3_\mu - g'B_\mu\right)^2\right]
$$

With $W^\pm_\mu = (W^1_\mu \mp iW^2_\mu)/\sqrt2$, so that $W^1W^1 + W^2W^2 = 2W^+_\mu W^{-\mu}$:

$$
\Longrightarrow\quad \frac{g^2v^2}{4}\,W^+_\mu W^{-\mu} \;+\; \frac{v^2}{8}\left(gW^3_\mu - g'B_\mu\right)^2
$$

$$
m_W = \frac{gv}{2}, \qquad m_Z = \frac{v\sqrt{g^2+g'^2}}{2}, \qquad m_\gamma = 0
$$

The photon combination is simply absent from that expression — there is no term to give it a mass.

### What the full expansion contains

Restoring $\Phi = \left(G^+,\ (v+h+iG^0)/\sqrt2\right)^T$, the same $|D_\mu\Phi|^2$ generates, in one stroke:

- kinetic terms for $h$ and the Goldstones
- the mass terms above, from $v^2$
- $hVV$ and $hhVV$ couplings, from the $2vh$ and $h^2$ pieces of $(v+h)^2$
- the bilinear mixing $m_W W^\mu\partial_\mu G^\mp + m_Z Z^\mu\partial_\mu G^0$ — the term $R_\xi$ gauge exists to cancel
- assorted $VGG$ vertices

All of it from one term, with no free parameters beyond $g$, $g'$, $v$. That rigidity is the point.

