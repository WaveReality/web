+++
Name = "Spinfield Model"
bibfile = "mechphys.json"
+++

Within this broad framework of discrete particles interacting via continuous waves, there are many degrees of freedom in specifying the details, in a way that ends up being physically accurate. Here are some of the primary questions:

* An original inspiration for the stochastic particle approach was Feynman's path integral approach to [[QED]], which is based on the idea that every possible pathway is somehow being sampled with an associated probability, via _virtual particles_ that are only transiently extant. This "quantum foam" provides a different kind of picture relative to a lone "real" discrete particle bumping around in space. Instead, at every moment, particles are constantly popping into and out of existence, most with tiny probabilities, and this is what shapes the behavior of the "real" particles that are observed.

    The key question here is whether the continuous wave function plus stochastic behavior in a real particle ends up capturing everything that is otherwise captured by the virtual particles? From a complexity management perspective, it makes sense in any case to start with the much simpler frameork of the known [[Dirac]] wave function operating on a real [[electron]], and see how far that goes.
    
* How does the particle influence the wave fields? This is the _back reaction_ question that has long bedeviled attempts to understand the detailed "mechanistic" physics taking place up close around a charged particle ([[@FordOConnell91]]; [[@FordOConnell93]]). Specifically, very near the particle itself, the electromagnetic potential would be very strong, and thus correspond to a very high energy density, which should then end up producing virtual particles etc. This is known as the _ultraviolet divergence_ problem. It can be "solved" by having some kind of cutoff at short length scales, but that seems fairly arbitrary. Also, as the particle moves discretely to neighboring grid points, that creates ripples in the force field. 

    One solution, consistent with the [[pilot-wave]] approach, is that the discrete particle itself doesn't actually generate any back reaction, and is instead purely a "surfer" on the waves. This would imply that the [[Dirac]] field itself would generate the corresponding [[Maxwell]] EM field, which also solves many conceptual difficulties because it is smoothly distributed over space. In short, the same arguments that favor the Dirac field as a kind of antenna for the force field also apply in the reverse direction, for the generation of force fields from charged matter.

    However, the Dirac field itself is likely representing both [[epistemological vs ontic]] contributions to overall uncertainty. So in principle, the _actual_ current location of the particle should be driving updating of the Dirac field. But there is little in the way of existing guidance for how this might happen. Mathematically, this should be just like the state right after a particle has been localized to one point,  but then the momentum is gone. Considerable work needs to be done investigating useful continuous momentum wave functions that can be driven by the localization of a discrete particle.

    A further wrinkle is the possibility that particles leave some kind of _trace_ on the wave field that then influences subsequent particles, as a way of explaining the apparent [[non-locality]] phenomena ([[@Sciarretta18]]; [[@Sciarretta21]]).
    
* Hidden particle states. The [[@^Ord96]] model involves 2 state variables for each particle, that correspond qualitatively to the [[spin]] degrees of freedom in a fermion particle. The [[Dirac]] equation in its second order formulation likewise has 4 wave state variables that mutually interact to produce spin, via the _spinor_ dynamics. These also produce the phase dynamics of quantum wave functions, and are important for the conservation properties of these functions. Thus, the possibility of these internal state degrees of freedom, beyond just a single binary state value, must be properly explored, especially as it might interact with the Dirac spinor states.

    One intuition is that the resting energy / mass of a particle is associated with this constant cycling through the [[spin]] states, and that somehow this cycling dynamic within a single cell is capable of maintaining some kind of momentum value, as it couples with the Dirac-like wave state that it is generating. This may be related to the phenomenon of _zitterbewegung_ ([[@Hestenes08]]; [[@Hestenes90]]]; [[@Sidharth09]]; [[@RomanRosoPlaja03]]; [[@BarutBracken81]]; [[@WangZhang01]]). Thus, the actual momentum represents a _spatial imbalance_ in this constant internal spin motion within the particle state.
    
    One tempting idea in this space is that you could somehow imagine a spatially distributed collection of "sub states" for each particle, kind of like a very compact localized wave function. However, once you stray outside of a single discrete cell, it becomes logistically very difficult to keep track of these distributed states, and preventing them from leaking out all over the place.
    
    The initial specific model here would involve at least two values for each cubic 3D axis (6 total), plus two for the "staying still" component, and the relative balance between the values along each axis determines the net momentum along that axis. These values are then updated in response to the gradient along each such axis, computed across the 26 point cubic neighborhood as usual ([[discrete gradient]]). The same technique as used in the [[Klein-Gordon]] and [[Dirac]] wave functions, where the mass term drags against the wave-based oscillation frequency, could be used to obtain the fundamental quantum frequency relationship.

* Multiparticle states are also an essential constraint on the system. The wave state variables must be shared across all particles of a given type (otherwise you're back in [[configuration space]]), but if they are holding important state for an individual particle, then how does that work when another particle's wave function gets close by? For example, if it is spinning the other way? Does that interfere with the spin of the other guy? This is also a good reason to keep the wave function less widely distributed, as should happen when the particle is the generator.

* The particle zoo: muon and tau vs [[electron]]s, antiparticles, [[neutrino]]s are "pure spin" without any charge. All of this suggests that charge and spin are two separate factors that could potentially dissociate, but yet stick together. Neutrinos are always left-handed, and anti-neutrinos are right-handed: this is a key constraint on the nature of spin and mass. Charge gives rise to (lots of) mass, but spin does not (much)? Maybe just kinetic energy in spin? The spin model should definitely accommodate neutrinos. Also, they have muon and tau flavors. So overall, this all sounds like a factorizable dynamic system.

    The muon decays after an average lifetime of 2.2 millionths of a second into an electron, a neutrino, and an antineutrino. Ok, so the muon is somehow a regular electron plus these two additional neutrinos worth of spin? And that adds up to 200x the rest mass of the electron, and yet the neutrinos themselves are nearly massless! And all of this is tied up with the [[weak]] force, which must then be considered. Perhaps the short-range nature of this force plays a critical role in the localized stochastic dynamics. Also, the random nature of weak-mediated decay processes is broadly consistent with the true randomness required here.

Overall, this is a relatively large space to explore, and there are many potential tradeoffs in terms of the risks of taking on too much versus, paradoxically, too little, where it may be the case that adding additional complexity (e.g., in the form of the quantum foam of virtual particles and multiparticle states more generally) solves problems that would otherwise plague simpler, reduced, single-particle frameworks.

## Photons    

Unlike fermion particles such as electrons, the EM field is not amenable to a discrete particle-like framework: [[photons]] have many problematic issues as discrete particles of the EM field. Therefore, it makes more sense to retain a "classical" [[Maxwell]] EM field, interacting with the discretized fermion cells, as in the [[semiclassical]] approach developed by a number of researchers (see [[@Struyve20]]; [[@Santos15]]).

## Virtual particles, the discrete lattice, and probability waves

Virtual particles are an essential feature of QED / quantum field theory, and yet their "ontological" status is clearly somewhat confusing: they aren't the "real" particles that we observe, and yet their fleeting existence is necessary for the theory to work, so in some sense they must be just as real as the "real" particles.

The discrete particle lattice framework provides a potential resolution to this conundrum. If any given "real" particle can potentially occupy any given cell in a discrete lattice, then there must effectively be a "slot" reserved for such a particle in each cell. These empty slots could provide an appealing basis for virtual particles, and the propagation and interactions of particles in the matrix.

In particular, a simple schema is that the probability waves associated with the standard interpretation of QM reflect a rippling propagation of probability factors across virtual particle slots in the matrix, with a real particle having a special status as being the current "true" location. Each possible jump to a neighboring cell involves a full transition matrix dependent upon the total energy (mass + kinetic) of the source: if the source is sufficiently energetic, it has some probability of activating a different combination of real particles as it makes the leap, accounting for the splitting tracks observed in particle accelerator experiments. Perhaps some of the "trace" in the matrix represents residual bits of this probability field propagating out and being left behind as real particles move around.

It is essential that these probability computations are all propagated in terms of _amplitudes_, not the probability values themselves, which are obtained by the product with the complex conjugate ("squaring").

## TODO

* zitterbewegung and helical spin in electron: [[@Hestenes08]]; [[@Hestenes90]]]; [[@Sidharth09]]; [[@RomanRosoPlaja03]]; [[@BarutBracken81]]; [[@WangZhang01]]

Key points: 

* key point about simple [[harmonic oscillator]] (SHO): the height position _itself_ provides the acceleration force pulling back -- this is even simpler than a wave! So there is a 4 vector of SHO that represent the momentum of the particle, and are coupled to the wave variables! The time-like one is the "heartbeat" of the particle, representing the rest mass, while the 3 spatial ones represent the velocity direction.

* each velocity axis (X,Y,Z) has a phase relative to the central time-like beat, and this phase represents the -1..+1 velocity value. This phase relationship is now demonstrated in [[complex KG]]. The extreme nutrino-level particle would have the extreme case. There would have to be some kind of coordination across the 3 pairs, so the total norm could not exceed 1.

* the resting state is all four oscillating in sync with 0 phase, and then forces act by boosting or lagging an axis wave relative to the central one. this could be pretty natural. the amount of boost needs to be dependent on existing phase to capture relativistic effects.

* the wave oscillation frequency is determined by the relativistic E^2 energy which goes up as momentum increases. This captures the key momentum / frequency relationship with QM. Never quite gets to CSq with massive particles.

* probability of jumping is then proportional to these values. can just implement that. it might emerge more naturally from some kind of phase offset dynamic, but that can be a later stage.

* coupling to surrounding wave is directly via drive from the central time-like oscillator, which provides the driving input to the field in its neighborhood.

* An acceleration kick just bumps the phase? doesn't capture the conservation dynamics among the SHOs.

Fundamental problem with localist particle driving the surrounding field:

* driven wave amplitude falloff is *always* 1/r, regardless of time constant!
* this sharply peaked wave field then gets very distorted when the particle makes its quantum jumps.
* if the surrounding field is supposed to be an antenna, this self-turbulence is going to create sharp high-frequency discontinuities.

Basically, the quantum jumps are the problem!  But they are inevitable. any graded, deterministic movement is going to be subject to the aliasing problem, or is going to include some amount of randomness. There is no advantage to some more complex kind of randomness.  See particle_graded_move.goal for experiment that demonstrates this (to my satisfaction).

Another alternative is to somehow propagate the driver in a more distributed fashion, so the local landscape is flatter.

The real problem is that there is a major disconnection between the initial "wave packet" that corresponds to what the matter wave *should* look like, and what a point-particle locally driving the wave field actually produces. The spatial wavelength doesn't look anything like the wave packet. It is just a pointwise oscillating blob with sharp 1/r falloff.

This is the fundamental back-reaction problem from the particle onto the field. it is completely unsolved ([[@Rohrlich00]]; [[@Rohrlich99]]; [[@Rohrlich97]]; [[@Baez20]]). whereas the SHO provides a very nice solution for the particle dynamics itself, I don't have any real plausible guidance here.

Key point: if self-field from particle is perfectly symmetric, and it is the _gradient_ of the field that drives forces, then the self-field problem can disappear! Just make it sufficiently smooth, and there is no problem!  Only another particle, which is some spatially asymmetric location away (the only spatially symmetric location is the self-location), can then cause force interactions.

Of course, the two slit example shows self-interactions, but here the self field "comes back" to interfere.. Normally it always just radiates _away_ (why exactly? key point!)

* if we have SHO field, how is it different from [[QED]]?  could we just have some kind of wave dynamic operating over the SHOs?  what distinguishes the SHOs from the other wave variables? need to think outside the box a bit about this parcellation into particles vs. waves -- is there some kind of more synthetic model there that I'm not seeing?

* first target should be the neutrino! not the electron! make a neutrino as a stochastic propagating SHO guy! why is it going at the speed of light?  Why is its phase dynamic set like that?  How does it 

Unscrambling the omlette of epistemological vs ontological is very tricky!  also the point guy will be scrambling around spreading its point energy all over the place -- don't want to include any memory process here (markovian) beyond the basic memory of the velocity via the SHO's.

How about a diffusion process that communicates the particle driver signal with only weak falloff proximally? basically a gaussian with a nice flat top?

so far, using an exponential factor with a max tends to produce anisotrophic behavior.

using laplacian form produces 1/r.  maybe just need a stronger recovery term, but it could blow up?


* so there isn't any randomness!?

## Stochastic KG

This is a first-pass attempt at an integrated wave-particle system based on the complex KG system.

* The mass-drag term only exists at one discrete point, otherwise it is just standard wave equation.

* The probability of a motion jump to the next cell is computed using the standard current density equation:

{id="eq_kg-current" title="Klein-Gordon current density"}
$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right)
$$

In terms of the underlying scalar state variables (and again for natural units), this is:

$$
\vec{J} = \frac{e}{m_0} (\varphi_a \vec{\nabla} \varphi_b - \varphi_b \vec{\nabla} \varphi_a)
$$

* The probability of staying in the same location is from the charge density:

{id="eq_kg-charge" title="Klein-Gordon charge density"}
$$
\rho \equiv \frac{i \hbar e}{2m_0c^2} \left( \phi^* \frac{\partial \phi}{\partial t} - \phi \frac{\partial \phi^*}{\partial t} \right)
$$

$$
\rho_i = \frac{e}{m_0} ({\varphi_b}_i \dot {\varphi_a}_i - {\varphi_a}_i \dot {\varphi_b}_i)
$$

* everything is normalized by the sum, and then the stochastic choice is made on these normalized probabilities.

* if staying put, then the energy equivalent of the mass-dependent factor:

$$
\frac{c^2 m_0^2}{\hbar^2} \varphi
$$

is converted into an acceleration of the complex state variables. How exactly??

* if moving, this same quantity is used to drive the motion in the new cell that is chosen as a weighted function of the projection of the charge density vector onto the laplacian neighbor vectors.

again, the force drives the rotation of the complex state values but how exactly?

