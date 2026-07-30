+++
bibfile = "mechphys.json"
+++

Wave equations provide such a natural explanation of so many physical phenomena, that it is very tempting to imagine that the entirety of physics can be produced exclusively through various coupled wave functions, as in the coupled [[Dirac]] -- [[Maxwell]] system. However, having implemented such a system, it is evident that the Achilles heel of waves is a fundamental problem: they inevitably just spread out and diffuse over time.

The obvious solution to this problem is to have some kind of discrete particle that cannot spread out, and remains localized in a specific point in space at any given point in time. In this context, the fundamental wave-particle duality in quantum physics suddenly starts to look like an essential _feature_ of the system, instead of some kind of paradoxical bug that nature threw at us for no good reason.

The basic [[wave]] equation represents a fundamental interaction between **space** (the spatial gradient, i.e., the _momentum_ operator in the [[Hamiltonian]]) and **time** (acceleration within a single grid element, i.e., the Hamiltonian _energy_ operator). In this context, consistent with [[special relativity]], a massive particle lives strictly in the time component. This can be seen in the [[Klein-Gordon]] equation:

{id="eq_kg" title="Klein-Gordon equation"}
$$
\frac{\partial^2 {\varphi}}{\partial t^2} = c^2 \nabla^2 \varphi - \frac{c^2 m_0^2}{\hbar^2} \varphi
$$

where the mass-dependent factor enters without any spatial gradient operator, and directly subtracts away from the energy factor (acceleration) on the left-hand side. This lack of spatial extent is consistent with a discrete particle entity, that somehow consumes energy from the wave field in proportion to its rest mass. This is the qualitative picture for how our discrete particles emerge, with the energy consumption propelling an internal oscillation within the discrete cell in the cubic lattice of the [[cellular automaton]] (CA) framework.

An obvious problem with this notion of something being contained entirely within a discrete cell is that it becomes challenging to imagine how it might ever move to another such cell. Such a move would have to happen in a discrete jump, creating a major discontinuity in the overall wave state, and potentially making the particle trajectory seemingly discontinuous and anisotrophic. 

The only way to overcome those difficulties is to use **stochastic** discrete jumps, such that, on longer time averages, the timing and spatial distribution of such jumps smooths out into a continuous, isotrophic distribution. Thus, contrary to Einstein's oft-cited objection that "God does not play dice with the universe", in fact it seems that an essential form of randomness is _necessary_ for discrete particles to move in a physically plausible manner.

Furthermore, although the continuum limit is mathematically approachable through the tools of calculus, it is problematic from a physics perspective due to the nearly-infinite field strengths (and thus energies) that would be present in the immediate vicinity of a charged particle. This _ultraviolet catastrophe_ is a recurring theme throughout the [[history]] of quantum physics, and it is nicely resolved through the use of the discrete cubic lattice of the CA framework.

## Quantum cellular automata and random walks

There is a growing literature on quantum cellular automata (QCA) and the quantum random walk (QW) model, going back to the influential papers by [[@^Feynman82]], [[@^AharonovDavidovichZagury93]], and [[@^Bialynicki-Birula94]]. This approach has become somewhat more active recently due to the practical application of these models for designing quantum algorithms ([[@ChildsCleveDeottoEtAl03]]; [[@ChiribellaDArianoPerinotti11]]; [[@DAriano17]]; [[@Kempe09]]). Although this work shares many essential properties in common with the present approach, it is not directly applicable, due to some differences in basic assumptions.

The central idea behind the random walk model proposed originally by [[@^AharonovDavidovichZagury93]] is that a shift operator that translates the quantum state in either the left or right direction along a discrete 1D lattice can be driven by a quantum measurement process on a _different_ element of the quantum state (e.g., the spin), which thus would produce a random sample according to the usual probabilistic quantum theory. This "internal" source of randomness will thus produce an emergent random walk dynamic as the state evolves (translates) over time, in the same way that thermal noise produces Brownian random walk motion as originally analyzed by Einstein.

The results from these analyses show that standard wave functions such as the [[Dirac]], [[Maxwell]] and [[Weyl]] functions emerge from these random walk processes. A closely related analysis in 3D by [[@^Bialynicki-Birula94]] showed the emergence of the Weyl and Maxwell wave functions in the context of a cellular automaton update rule with _unitary_ update rules.

The basic intuition is that these unitary update rules cause the propagation of a wave-like pattern at the speed of light (one unit cell per unit time), and if multiple internal cell states are present, this propagation can also include the [[spin]] property where the state rotates through these internal states as it also propagates. This is the essential feature of the Weyl wave functions, which describe a "pure spin" particle such as a massless [[neutrino]] that travels at the speed of light while spinning in one fixed helical rotation.

Critically, these approaches do _not_ provide a model of a discrete particle moving with graded momentum (velocities) in an isotrophic manner along a cubic grid, which is what we develop in [[stochastic particles]]. That work builds on the approach originated by [[@Nelson66]] in analyzing single-particle Brownian motion, using equations of motion initially developed by [[@^Sciarretta18]].

## todo

The seminal work in analyzing stochastic discrete particle motion within the context of quantum physics was done by [[@^Nelson66]], who showed that a form of stochastic discrete particle motion actually results in the [[Schrodinger]] wave function in the continuous time-average limit. Interestingly, this work builds on the original work by Einstein on Brownian random-walk motion, back in 1905. Subsequent work has developed these ideas in multiple ways ([[@Cufaro-PetroniVigier83]]; [[@CufaroPetroniVigier79]]; [[@Ord96]]; [[@^Sciarretta18]]; [[@^Sciarretta21]] and others reviewed therein).

Critically none of this existing work involves an integrated wave-particle [[duality]]; it focuses exclusively on the time-average distributions of discrete particle motion. Thus, unlike the [[pilot-wave]] framework, the quantum wave function in a purely stochastic particle model is entirely [[epistemic vs ontic|epistemic]]: it just describes the expected value of a discrete particle's random walk trajectories over time. There is no physical reality to such a wave. By contrast, our goal here is to derive an integration of discrete particle motion with quantum wave functions.

{id="figure_pf-origin" style="height:20em"}
![Stochastic origin of quantum momentum / frequency relationship. The momentum on the left is 0.5c while on the right is 0. The distribution of position is on the vertical axis, while time is on the horizontal axis, with each point centered at the origin in the center (i.e., the temporal autocorrelation function). The variance on the left is half of that on the right.](media/fig_asmom5_0_autoc.png)

One key intuition for why discrete particle motion naturally exhibits quantum wave-like behavior is that a slow drift rate produces a wide cloud of space where particle could be, corresponding to a long wavelength in the probability cloud that the Schrödinger wave function describes. However, when the particle has high momentum (velocity), it moves more deterministically in a given direction, resulting in a narrower range of variance around the particle's mean trajectory, resulting in a narrower effective wavelength ([[#figure_pf-origin]]).

While this pure particle-based approach is appealing in its simplicity, it does not appear to provide an explanation for phenomena such as the [[double-slit]] experiment, where somehow a particle can interfere _with itself_, but only if the other slit is open. Furthermore, it cannot be the case that these interference effects only arise in the rare cases when a discrete particle happens to wander so aimlessly as to go through both slits somehow.

Furthermore, discrete particles require some other continuous state values embedded in the same discrete lattice to drive the _probabilities_ underlying their stochastic behavior. The nature of such values is often unaddressed in existing frameworks. The [[pilot-wave]] framework of de Broglie and Bohm provides a natural mechanism for these continuous momentum values, where the continuous-valued wave function guides the motion of a discrete particle in some way.

Overall, the discrete particle framework does a great job of keeping the accounting tight, in comparison to the fundamentally sloppy [[matter waves]] that end up diffusing over time and space. The discrete particles can be strictly conserved, and always tightly localized, while the surrounding wave functions diffuse outward from this central island of stability.

However, a system based purely on discrete particles has a very hard time managing interactions among the particles via forces: the odds of any other discrete particle "hitting" another are very small, whereas these big sloppy waves provide a nice continuous, broadly diffusing and saturating medium for force transmission. In effect, the surrounding wave field is essential for the particle to be able to properly "sense" the force field effects (from electromagnetic, weak and strong forces). Thus, the wave field functions like an _antenna_ or "whiskers" in sensing forces over a broader space, beyond its own singular cell.

From a "design" perspective, this _hybrid_ framework of discrete particles and continuous wave fields represents a "best of both worlds" solution, relative to something based exclusively on discrete particles, or exclusively on continuous wave functions. This then provides a compelling reason for the otherwise central puzzle of the quantum world: _why_ is there this wave-particle duality? Why is nature so strangely complicated in this way?

Relative to the standard QM interpretations, the pilot-wave nature of this approach avoids all the impossible conundrums of instantaneous wave collapse at the point of measurement, as the discrete particles are there at every point in time and space, even when you're not looking. The primary challenge for such a framework is dealing with the apparent [[non locality]] of quantum entanglement.

## Remaining degrees of freedom

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

* The particle zoo: muon and tau vs [[electron]]s, antiparticles, [[neutrinos]] are "pure spin" without any charge. All of this suggests that charge and spin are two separate factors that could potentially dissociate, but yet stick together. Neutrinos are always left-handed, and anti-neutrinos are right-handed: this is a key constraint on the nature of spin and mass. Charge gives rise to (lots of) mass, but spin does not (much)? Maybe just kinetic energy in spin? The spin model should definitely accommodate neutrinos. Also, they have muon and tau flavors. So overall, this all sounds like a factorizable dynamic system.

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

## Stochastic particle equations

The basic behavior of a stochastic particle independent of any waves is described by [[@^Sciarretta18]] (in the 1D case; [[@Sciarretta21]] extends to the 3D case with spin). In this non-relativistic model, the particle has an associated real-valued 3-component normalized (range -1..1) momentum vector $\vec{\nu}$ that drives a _stable_ trajectory over time, despite the stochastic nature of each movement step. Any forces accumulate in this momentum vector, and it propagates with the particle. The motion of the particle is defined per dimension $\mu$ (per the [[four-vector]] notation), in reference to an energy-like factor:

todo: figure out if e should in fact be across all dims -- seems like it should. just run an empirical sim.

$$
e_{\mu} = \frac{1 + \nu_{\mu}^2}{2} 
$$

The probability of moving in a positive or negative direction, or staying put, along a given dimension is given by:

$$
P(+1) = \frac{e_{\mu} + \nu_{\mu}}{2}
$$

$$
P(-1) = \frac{e_{\mu} - \nu_{\mu}}{2}
$$

$$
P(0) = 1 - e_{\mu}
$$

The energy-like factor $e$ is .5 when the momentum $\nu=0$, and thus there is a .25 probability of going in either direction, and a .5 probability of staying in the same location. As momentum increases in say the positive direction, the probability of moving in the positive direction increases in a proportional manner, while the opposite direction decreases. The squaring of the energy factor ensures a division among the staying and move directions works out.

It would be useful to derive a relativistic version of these individual motion equations, where the effective mass, and thus momentum, increases with velocity. This is what the wave function naturally does.

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



Weyl wave couples spin with direction as a helical thing.

There is a literature on coupling of a stochastic particle with a "heat bath", somewhat like the [[zero-point]] field, and trying to understand the aggregate behavior of such a system. [[@^DunkelHanggi05a]], [[@DunkelHanggi05]] provide a relatively accessible treatment, building on foundational work ([[@Dudley65]], [[@Dudley73]], [[@GuerraRuggiero78]], [[@Nakagomi88]]). This all builds on Langevin equations, which are stochastic equations of motion, with connections to Ornstein-Ullenbach and Fokker-Planck etc. The specific restriction to heat bath dynamics vs. some kind of other intrinsic stochastic process is perhaps overly restrictive, but they nevertheless have a four-vector representation that seems to involve a conservation of energy between the time and momentum factors, which is really the essential calculus for the SHO model.

One key "no-go" finding from [[@^Dudley65]] is that a purely Markovian position-based system doesn't capture particle motion -- you _need_ an additional momentum / velocity vector as part of the state. This is definitely key.

* compton wavelength as function of rest mass -- does this fall out?
* $\nu$ is already 0-c normalized -- v/c
* wave function trades energy against momentum -- momentum is $\gamma m0 \nu$, $E^2 = p^2 + m0^2$ so the 1 in above eq is like m0^2 -- not clear where the 1/2 comes from but whatever.
* key idea that m0 is the internal motion of the particle rotating through spin, so need to just have that always going on as an "anchor", and then there are these extra $\nu_{\mu}$ factors where the x^2 + y^2 + z^2 hypotenuse of the momentum-velocity is v^2 relative to c^2 -- i.e., need to constrain total length to c^2.

* probabilistically when it stays still it rotates the internal spin. if it never stays still it never rotates the spin. the nutrino very rarely rotates the spin, but does sometimes. spin coupling couples the two helicies of the Weyl. need to go back to that. are there 2 neutrinos trapped inside one electron??

* todo:

15. Ord GN, Schrödinger's Equation and Classical Brownian Motion, Fortschr. Phys. 46, 6–8, 889–
896 (1998).
16. Janaswamy R, Transitional probabilities for the 4-state random walk on a lattice, J. Phys. A: Math.
Theor. 41, 1–11 (2008).
17. Badiali JP, Entropy, time-irreversibility and the Schrödinger equation in a primarily discrete
spacetime, J. Phys. A: Math. Gen. 38(13), 2835–2848 (2005).
32. Snyder HS, Quantized space-time, Phys. Rev. 71(1), 38-41 (1947).
33. Finkelstein D, Saller H, and Tang Z, Quantum spacetime, in: P. Pronin, et al. (Eds.), Gravity,
Particles and Space Time, World Scientific, Singapore (1996).
34. Sidharth BG, The Thermodynamic Universe: Exploring the Limits of Physics, World Scientific,
Singapore (2008).

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

