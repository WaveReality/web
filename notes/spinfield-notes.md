+++
Categories = ["Spinfield Model"]
bibfile = "mechphys.json"
+++

## spin-motion coupling etc

Would like to be able to say something like this:

Phenomenologically, the charge of the electron emerges as a two coupled CVHOs oscillating 90 degrees out of phase with each other, while the neutrino has only one or the other of these two CVHOs. In other words, the electron is essentially two coupled neutrinos, trapped together so that they produce charge in their phase relationship. Furthermore, the muon is a heavy version of an electron, which quickly decays into an electron and two neutrinos. This can be understood as an excited state of the electron, with _extra spin energy_ reflected in the extra neutrinos that are emitted upon decay.

This structural, compositional model of the leptons illustrates the conceptual power of the Spinfield Model.

HOWEVER: how exactly does the motion stuff work? what is the chirality of the neutrinos? The ORD model has things moving within internal states but they are actual position states -- spin seems positional? 2x additional states within. Maybe distributing over 2 coupled states is OK? Always moving across these states but probabilistically jumping to other states to produce motion? how does it work in 3D?

And everything is always oscillating, so what is really moving? everything is always there and always has a phase!?

Current model is to have one central T0 heartbeat, and then motion and charge as other states relative to that.. 

## Photons

Unlike fermion particles such as electrons, the EM field is not amenable to a discrete particle-like framework: [[photon]]s have many problematic issues as discrete particles of the EM field. Therefore, it makes more sense to retain a "classical" [[Maxwell]] EM field, interacting with the discretized fermion cells, as in the [[semiclassical]] approach developed by a number of researchers (see [[@Struyve20]]; [[@Santos15]]).

## Virtual particles, the discrete lattice, and probability waves

In particular, a simple schema is that the probability waves associated with the standard interpretation of QM reflect a rippling propagation of probability factors across virtual particle slots in the matrix, with a real particle having a special status as being the current "true" location. Each possible jump to a neighboring cell involves a full transition matrix dependent upon the total energy (mass + kinetic) of the source: if the source is sufficiently energetic, it has some probability of activating a different combination of real particles as it makes the leap, accounting for the splitting tracks observed in particle accelerator experiments. Perhaps some of the "trace" in the matrix represents residual bits of this probability field propagating out and being left behind as real particles move around.

* if we have SHO field, how is it different from [[QED]]?  could we just have some kind of wave dynamic operating over the SHOs?  what distinguishes the SHOs from the other wave variables? need to think outside the box a bit about this parcellation into particles vs. waves -- is there some kind of more synthetic model there that I'm not seeing?

## Stochastic KG

This is a first-pass attempt at an integrated wave-particle system based on the complex KG system.

* The mass-drag term only exists at one discrete point, otherwise it is just standard wave equation.

* The probability of a motion jump to the next cell is computed using the standard current density equation:

{id="eq_kg-current" title="Klein-Gordon current density"}
$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \chi^* \vec{\nabla} \chi - \chi \vec{\nabla} \chi^* \right)
$$

In terms of the underlying scalar state variables (and again for natural units), this is:

$$
\vec{J} = \frac{e}{m_0} (\phi_a \vec{\nabla} \phi_b - \phi_b \vec{\nabla} \phi_a)
$$

* The probability of staying in the same location is from the charge density:

{id="eq_kg-charge" title="Klein-Gordon charge density"}
$$
\rho \equiv \frac{i \hbar e}{2m_0c^2} \left( \chi^* \frac{\partial \chi}{\partial t} - \chi \frac{\partial \chi^*}{\partial t} \right)
$$

$$
\rho_i = \frac{e}{m_0} ({\phi_b}_i \dot {\phi_a}_i - {\phi_a}_i \dot {\phi_b}_i)
$$

* everything is normalized by the sum, and then the stochastic choice is made on these normalized probabilities.

* if staying put, then the energy equivalent of the mass-dependent factor:

$$
\frac{c^2 m_0^2}{\hbar^2} \phi
$$

is converted into an acceleration of the complex state variables. How exactly??

* if moving, this same quantity is used to drive the motion in the new cell that is chosen as a weighted function of the projection of the charge density vector onto the laplacian neighbor vectors.

again, the force drives the rotation of the complex state values but how exactly?

