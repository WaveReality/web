+++
Name = "Weyl"
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

The Weyl equations, developed by Hermann Weyl as a simplification of the [[Dirac]] wave function for a spin 1/2 particle ([[fermion]]) without any rest mass, provide a possible model of the [[neutrino]] if it were massless (which it does not appear to be). Because the neutrino mass is very small, it turns out that the Weyl equations nevertheless provide a reasonable approximation.

Furthermore, the Weyl equations are having a bit of a recent comeback in the [[field theory]] framework for high-energy situations, where the rest mass is a relatively small portion of the total energy. The simpler two-component nature of the Weyl equations (versus the 4 required for the first-order Dirac equation) turns out to be advantageous in simplifying many important computations, forming the basis for the _spinor helicity formalism_ ([[@ElvangHuang14]]).

For our purposes, the key feature of these wave functions (there are two of them, one for each _helicity_), is that they provide a simpler introduction to the central properties of [[spin]]. A critical conclusion is that:

> **spin 1/2 is fundamentally a first-order phenomenon.**

At a basic mathematical level, the standard second-order [[wave]] equation involves _squared_ quantities, in that a second-order derivative is the repeated application of the derivative operation (derivative squared). This squaring effectively causes the spin 1/2 case to effectively rotate all the way around!

This situation is essentially the same as the relationship between the second-order [[Klein-Gordon]] (KG) equation versus the first-order [[Schrodinger]] equation. KG operates on a single real value, but requires maintaining the velocity in addition to the wave value, while Schrödinger's equation operates on a complex value. The two degrees of freedom in the complex number essentially map onto the position and velocity variables in the second-order equation.

Interestingly, Dirac derived his first-order wave equation directly from the second-order KG equation, by effectively taking the square root. Any time you take a square root, there are 2 solutions of different signs, and this is where the doubling of wave state variables comes from.

In effect, the spinning that happens in the Weyl equations represents a rotation through the two coupled complex wave components:

{id="eq_two-component" title="two-component complex wave state"}
$$
\psi = \begin{bmatrix} \phi_{1a} + i \phi_{1b} \\ \phi_{2a} + i \phi_{2b} \end{bmatrix}
$$

This rotation happens at the speed of light, as the wave itself also propagates at light-speed. The two different _helicity_ variants of the Weyl equations reflects whether this rotation is effectively spinning in a right-hand or left-hand manner relative to the wave propagation direction. You can see this on your own hands, by having your thumb represent the motion direction, and the curled-around fingers represent the rotation direction.

The fist order, right-hand helicity of the Weyl equation in [[four-vector]] notation looks like this:

{id="eq_weyl-first" title="first-order right-handed Weyl equation"}
$$
\sigma^{\mu} \partial_{\mu} \psi = 0
$$

where $\sigma^{\mu}$ are the Pauli spin matricies, that provide the essential rotation dynamics:

{id="eq_pauli" title="Pauli matricies"}
$$
\vec{\sigma} = \left( \begin{bmatrix} 0 & 1\\
1 & 0 \end{bmatrix}, \begin{bmatrix} 0 & -i \\
i & 0 \end{bmatrix},  \begin{bmatrix} 1 & 0 \\
0 & -1 \end{bmatrix} \right)
$$

You can see how these basically map one variable onto another in the first two cases, where everything is off-diagonal. The third, _z_ case does a flip due to the minus sign -- this is the direction of motion axis.

These spin matricies are in fact defined by the fact that applying them twice gets you back to where you started, which is what creates the alignment between the spinless second-order KG equation and the spinning first-order Dirac or Weyl equations ([[@Brown58]]; [[@Tonin59]]; [[@Marx67]]; [[@Marx70]]; [[@Case57]]; [[@Diaz-CruzLopezMeza-AldamaEtAl15]]; [[@KibblePolkinghorne58]]; [[@BarutMullen62]]; [[@BabinFigotin14]]; [[@Cardoso93]]; [[@Veblen33]]; [[@DreinerHaberMartin10]])

We are left with a fundamental challenge here: if spin is fundamentally a first-order phenomenon, we cannot capture first-order waves in an isomorphic manner within the [[cellular automaton]] framework. If it is just standard second-order wave propagation as in the KG equation, then what's the big deal?

Interestingly, the second-order version of the Dirac equation only involves spin in the context of the coupling with the [[Maxwell]] EM field. The components themselves are not spinning into each other.

Thus, the fundamental question is: is spin only revealed in the interaction with the EM field, or is there some more fundamental phenomenon that we're missing out on in the second-order versions?

The neutrino provides strong indications that a second-order wave function is insufficient to account for observed spin 1/2 phenomena. For example, spin is a conserved property, and the emission of an electron and a neutrino during beta decay (involving the [[weak]] force) must preserve the total spin. Furthermore, the difference between a neutrino and an antineutrino is entirely in terms of their spin.

Thus, the conclusion from all these considerations is that:

> **spin 1/2 is fundamentally a particle-level phenomenon** 

which must be captured by the [[stochastic motion]] model.

