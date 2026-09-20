+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

The **weak** force on its own is responsible for various forms of _decay_ or _transmutation_ from one type of particle to another, most notably in the case of _beta decay_ where for example a proton decays into a neutron within the nucleus. This is a relatively rare event, and the weak force is, after all, "weak", so in general it perhaps doesn't get as much attention as it otherwise might.

However, the **electroweak** unification of the electromagnetic force (i.e., [[Maxwell]]s equations) together with the weak force, along with the central role of the [[Higgs]] field in this framework, provides a profoundly different picture of the fundamental nature of an [[electron]] and its interaction with the electromagnetic field. This electroweak picture also integrates the [[neutrino]]s with their charged [[lepton]] partners like the electron, and is intimately connected with the _flavors_ or [[generation]]s of particles, e.g., muon and tau.

Thus, it is not an exaggeration to say that the electroweak framework is as revolutionary for understanding the nature of leptons as the [[quark]] framework is for understanding the hadrons: it provides an entirely different fundamental basis space, that parsimoniously integrates a wide range of disparate phenomena into a more coherent framework. It is simply not sensible to try to understand the most basic nature of the electron and electromagnetic interactions without understanding the electroweak model.

However, the counter-argument is that the actual physical implications of this revolutionary rearrangement of the quantum furniture are somewhat difficult to detect, which presumably why this part of the story gets less attention than it otherwise might.

Interestingly, there are many features of this electroweak framework that align with aspects of the [[Spinfield Model]], which thus can provide essential guidance on how that model should be configured. Therefore, it is of central importance to the current effort.

## Electroweak

The electroweak framework is based on three fields, which together comprise a total of 4 four-vector potential-like fields, and the Higgs doublet, which has 4 real-valued numbers, organized into two complex values.

* The **Higgs doublet** $\Psi$, which has 2 complex field elements (4 real numbers), one of which interacts with electric charge ($\phi^+$) and another that is electrically neutral ($\phi^0$). As a result of the spontaneous symmetry breaking property of the Higgs potential, the neutral component acquires a stable expected value of $v_h = 246 GeV$, which is what then provides the mass term for all other wave fields.

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
W^+ = \frac{1}{\sqrt2} \left( W^1_\mu - W^2_\mu)
$$

$$
W^- = \frac{1}{\sqrt2} \left( W^1_\mu + W^2_\mu)
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
| Higgs $\phi^+$         | 1/2  | 1/2     | 1/2  | 1  |
| Higgs $\phi^0$         | 1/2  | 1/2     | -1/2 | 0  |

[[#table_weak-qs]] shows these quantum numbers for neutrinos and electrons (the leptons). This table, plus the mixing factors shown above, clearly show that the electroweak framework provides an entirely different factorization of the theoretically fundamental lepton particles. There are two fundamentally different versions of the electron, and the left-chiral electron and the neutrino share most of their quantum numbers, so they would seem to be much more similar at the electroweak level than we would otherwise expect given their overall properties.

Thus, it seems that one's understanding of what an electron really is changes dramatically under this electroweak framework. Furthermore, every one of these electroweak fields, including is coupling to the very same Higgs doublet field: what are the implications of this mutual interaction through this shared field?

A fundamental question raised by all of this, is are the observable fields a mixture and not just equivalent to the corresponding weak hypercharge and weak isospin fields that we started with? This is described as being a consequence of the spontaneous symmetry breaking dynamic in the Higgs field, but the story is quite a bit more complicated than that. The symmetry breaking is _necessary_ for these mixture factors to actually matter, because prior to this symmetry breaking, all of the four-potentials are essentially equivalent, acting like the EM massless four-potential. Once the symmetry is broken and the Higgs field has a vacuum expectation value, then all of the differences that were actually built into the system start to matter.

In particular, the gauge theory derivation of how the Higgs field couples to the weak hypercharge and weak isospin fields results in a **covariant derivative** that cancels out the local gauge factors arising from the gauge fields, of the form:

$$
D_\mu = \partial_\mu - i g' Y B_\mu - i g W^a_\mu T^a
$$

Where $g'$ and $g$ are arbitrary real-valued coupling parameters for each of the respective fields ($B_\mu$ and $W^a_\mu$), and $Y$ is the weak hypercharge quantum number value for the Higgs field (a real number), which is set to 1/2 by convention. The $T^a$ is a set of 3 different coupling matricies, one for each of the three $W^a_\mu$ four-vector components, that are defined by the SU(2) unitary rotation group -- i.e., a set of basis vectors that perform unitary rotations of a 2x2 matrix, which is what the Higgs field is (2 complex values for each of 2 doublets).

These are none other than the [[Pauli matricies]], which are also used in the [[Dirac]] equation, and define the property of [[spin]] in the quantum world. This is why the $W$ field is called _isospin_, because it causes spinning. Each of these are also multiplied by the conventional 1/2 factor, but it turns out that the third Pauli matrix has a -1 on the bottom-right diagonal, which therefore gives the neutral Higgs doublet component $\phi^0$ a -1/2 effective $T^3$ quantum number.

This covariant derivative is used in the [[Lagrangian]] for the Higgs field:

$$
\mathcal{L}_h = |D_\mu \Psi|^2 + V(\Psi)
$$

where $V(\Psi)$ is the Higgs potential that leads to symmetry breaking, as described in detail in [[Higgs]], and the Higgs field $\Psi$ is the (1/2 normalized) complex doublet:

{id="eq_state" title="Higgs complex doublet"}
$$
\Psi \def \frac{1}{\sqrt2} \binom{\phi^+}{\phi^0}
$$

From all of this, we can now see that the _reason_ that electric charge and the electromagnetic field are a mixture of the $B_\mu$ and $W^3$ fields is because of the way that these gauge coupling factors work. In particular, the only way for the EM field to remain fully massless (and thus have long-range effects) is for it to have opposite effects on the underlying Higgs field, so that these effects _perfectly_ cancel out. This cancellation is captured qualitatively by the definition of electric charge in terms of the sum of the factors for each field, and quantitatively it is given by the ratio of the coupling parameters $g'$ and $g$, which determines the weak mixing angle:

$$
\sin \theta_w = \frac{g'}{\sqrt{g^2 + g'^2}}
$$

These coupling parameters are determined experimentally, and are not predicted by the theory itself. Nevertheless, their impacts have far-reaching implications on a range of other parameters and experimental outcomes, all of which have been strongly confirmed. As with everything else in the Standard Model at this point, all of this seemingly crazy complexity has been experimentally validated to high levels of precision, with this electroweak theory making strong predictions for things that had not yet been seen experimentally before.

Furthermore, the gauge theory framework that drives so much of the phenomenology here is fundamentally based on describing the [[conservation]] of energy and charge in the equations of motion, derived from the Lagrangian framework. Thus, ultimately, these conservation dynamics are what is driving everything. And the only way to get the math to all work out was to do it with massless gauge fields that work like the EM four-potential, and then have the mass emerge dynamically from within that fundamentally massless context, which the specific form of the Higgs potential accomplishes.

Thus, even though this seems like a very complicated way of going about everything, in the end each step is strongly constrained, and it is difficult to see how a different kind of system could be derived.

## Phenomenology of the weak force

### EM does _not_ couple to the Higgs field in the end

The weak mixing angle ensures that the EM field $A_\mu$ has _precisely_ canceling effects on the Higgs field, meaning that it has no net effect. This seems like a bit of a strong coincidence. Or is it definitional: given that a zero exists, that specific mixture is _selected_ empirically as the one that has long-range propagation. So it doesn't have to have been done "a priori" or magically.

But if these other wave fields are really what is there, then how exactly does this specific mixture manage to somehow emerge as a linear mixture _after the fact_?  todo: ask claude about this

### The weak force bosons W are very hard to activate

* heavy
* requires precise phase coupling between electron and neutrino -- forbidden!

### What happens to the other 3 components of the Higgs field?

(eaten by W, Z)


## Derivation of gauge coupling

$$
\Phi = \begin{pmatrix}\phi^+\\ \phi^0\end{pmatrix}
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

