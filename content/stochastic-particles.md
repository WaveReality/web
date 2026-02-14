+++
bibfile = "mechphys.json"
+++

There are two key principles that allow discrete particles in a cubic lattice (as in the [[cellular automaton]] framework) to exhibit fully continuous, isotrophic motion:

* Stochastic Brownian movement enables a discrete particle to exhibit macroscopically continuous overall behavior, based on the original work of [[@^Nelson66]], [[@^Ord96]], and subsequent work by [[@^Sciarretta18]]; [[@^Sciarretta21]] and others reviewed therein. Furthermore, this motion naturally exhibits the main properties of the quantum wave function (e.g., the [[Schrodinger]] equation as shown by [[@Nelson66]]), where a slow drift rate produces a wide cloud of space where particle could be, corresponding to a long wavelength in the probability cloud that the Schrodinger wave function describes. However, when the particle has high momentum, it moves more deterministically in a given direction, resulting in a narrower range of variance around the particle's mean trajectory, resulting in a narrower effective wavelength.

* These discrete particles rely on other continuous state values embedded in the same discrete lattice to drive the _probabilities_ underlying their stochastic behavior. In this way, the framework is naturally consistent with the essence of the [[pilot wave]] framework of de Broglie and Bohm, where a continuous-valued wave function guides the motion of a discrete particle in some way.

{id="figure_pf-origin" style="height:20em"}
![Stochastic origin of quantum momentum / frequency relationship. The momentum on the left is 0.5c while on the right is 0. The distribution of position is on the vertical axis, while time is on the horizontal axis, with each point centered at the origin in the center (i.e., the temporal autocorrelation function). The variance on the left is half of that on the right.](media/fig_asmom5_0_autoc.png)

Overall, the discrete particle framework does a great job of keeping the accounting tight, in comparison to the fundamentally sloppy [[matter waves]] that end up diffusing over time and space. The discrete particles can be strictly conserved, and always tightly localized, while the surrounding wave functions diffuse outward from this central island of stability.

However, a system based purely on discrete particles has a very hard time managing interactions among the particles via forces: the odds of any other discrete particle "hitting" another are very small, whereas these big sloppy waves provide a nice continuous, broadly diffusing and saturating medium for force transmission. In effect, the surrounding wave field is essential for the particle to be able to properly "sense" the force field effects (from electromagnetic, weak and strong forces). Thus, the wave field functions like an _antenna_ or "whiskers" in sensing forces over a broader space, beyond its own singular cell.

From a "design" perspective, this _hybrid_ framework of discrete particles and continuous wave fields represents a "best of both worlds" solution, relative to something based exclusively on discrete particles, or exclusively on continuous wave functions. This then provides a compelling reason for the otherwise central puzzle of the quantum world: _why_ is there this wave-particle duality? Why is nature so strangely complicated in this way?

Relative to the standard QM interpretations, the pilot-wave nature of this approach avoids all the impossible conundrums of instantaneous wave collapse at the point of measurement, as the discrete particles are there at every point in time and space, even when you're not looking. The primary challenge for such a framework is dealing with the apparent [[non locality]] of quantum entanglement.

## Remaining degrees of freedom

Within this broad framework of discrete particles interacting via continuous waves, there are many degrees of freedom in specifying the details, in a way that ends up being physically accurate. Here are some of the primary questions:

* Is the quantum wave function real? Or can you somehow avoid it entirely, with the stochastic behavior of discrete particles producing the observed quantum-like effects?

    Although it has been shown that the [[Schrodinger]] wave function emerges naturally from stochastic Brownian motion ([[@Nelson66]]; [[@Ord96]]), these approaches don't address how the underlying continuous drift rate (momentum) parameter might be represented, and available to influence the particle's motion. Furthermore, the purely local effects of the EM force field available to a discrete particle are likely not sufficient to capture all the known physical effects, particularly with respect to spin and magnetic field interactions. Thus, as in the full de Broglie-Bohm [[pilot wave]] framework, it likely makes more sense to think in terms of a real quantum wave function that interacts in a continuous, distributed manner with the distributed EM force field, and then this quantum wave locally influences the behavior of the stochastic particle (as a function of the local gradient).

* An original inspiration for the stochastic particle approach was Feynman's path integral approach to [[QED]], which is based on the idea that every possible pathway is somehow being sampled with an associated probability, via _virtual particles_ that are only transiently extant. This "quantum foam" provides a different kind of picture relative to a lone "real" discrete particle bumping around in space. Instead, at every moment, particles are constantly popping into and out of existence, most with tiny probabilities, and this is what shapes the behavior of the "real" particles that are observed.

    The key question here is whether the continuous wave function plus stochastic behavior in a real particle ends up capturing everything that is otherwise captured by the virtual particles? From a complexity management perspective, it makes sense in any case to start with the much simpler perspective of the known [[Dirac]] wave function operating on a real [[electron]], and see how far that goes.
    
* How does the particle influence the wave fields? This is the _back reaction_ question that has long bedeviled attempts to understand the detailed "mechanistic" physics taking place up close around a charged particle. Specifically, very near the particle itself, the electromagnetic potential would be very strong, and thus correspond to a very high energy density, which should then end up producing virtual particles etc. Also, as the particle moves discretely to neighboring grid points, that creates ripples in the force field. 

    One solution, consistent with the [[pilot wave]] approach, is that the discrete particle itself doesn't actually generate any back reaction, and is instead purely a surfer on the waves. This would imply that the [[Dirac]] field itself would generate the corresponding [[Maxwell]] EM field, which solves many of the conceptual difficulties because it is smoothly distributed over space. In short, the same arguments that favor the Dirac field as a kind of antenna for the force field also apply in the reverse direction, for the generation of force fields from charged matter.

    However, the Dirac field itself is likely representing both [[epistemological vs ontic]] contributions to overall uncertainty. So in principle, the _actual_ current location of the particle should be driving updating of the Dirac field. But there is little in the way of existing guidance for how this might happen.
    
    A further wrinkle is the possibility that particles leave some kind of _trace_ on the wave field that then influences subsequent particles, as a way of explaining the apparent [[non-locality]] phenomena ([[@Sciarretta18]]; [[@Sciarretta21]]). More generally, the way in which the system behaves in a multi-particle situation is of utmost importance. The 

* Hidden particle states. The [[@Ord96]] model involves 2 additional state variables for each particle, that correspond qualitatively to the spin degrees of freedom in a fermion particle. The Dirac equation in its second order formulation likewise has 4 wave state variables that mutually interact to produce spin, via the _spinor_ dynamics. These also produce the phase dynamics of quantum wave functions, and are important for the conservation properties of these functions. Thus, the possibility of these internal state degrees of freedom, beyond just a single binary state value, must be properly explored, especially as it might interact with the Dirac spinor states.

Overall, this is a relatively large space to explore, and there are many potential tradeoffs in terms of the risks of taking on too much versus, paradoxically, too little, where it may be the case that adding additional complexity (e.g., in the form of the quantum foam of virtual particles and multiparticle states more generally) solves problems that would otherwise plague simpler, reduced, single-particle frameworks.

## Photons    

Unlike fermion particles such as electrons, the EM field is not amenable to a discrete particle-like framework: "photons" have many problematic issues as discrete particles of the EM field, as elaborated below. Therefore, it makes more sense to retain a "classical" [[Maxwell]] EM field, interacting with the discretized fermion cells, as in the _semiclassical_ approach developed by a number of researchers (see [[@Struyve20]]; [[@Santos15]]).

## Virtual particles, the discrete lattice, and probability waves

Virtual particles are an essential feature of QED / quantum field theory, and yet their "ontological" status is clearly somewhat confusing: they aren't the "real" particles that we observe, and yet their fleeting existence is necessary for the theory to work, so in some sense they must be just as real as the "real" particles.

The discrete particle lattice framework provides a potential resolution to this conundrum. If any given "real" particle can potentially occupy any given cell in a discrete lattice, then there must effectively be a "slot" reserved for such a particle in each cell. These empty slots could provide an appealing basis for virtual particles, and the propagation and interactions of particles in the matrix.

In particular, a simple schema is that the probability waves associated with the standard interpretation of QM reflect a rippling propagation of probability factors across virtual particle slots in the matrix, with a real particle having a special status as being the current "true" location. Each possible jump to a neighboring cell involves a full transition matrix dependent upon the total energy (mass + kinetic) of the source: if the source is sufficiently energetic, it has some probability of activating a different combination of real particles as it makes the leap, accounting for the splitting tracks observed in particle accelerator experiments. Perhaps some of the "trace" in the matrix represents residual bits of this probability field propagating out and being left behind as real particles move around.

It is essential that these probability computations are all propagated in terms of _amplitudes_, not the probability values themselves, which are obtained by the product with the complex conjugate ("squaring").

