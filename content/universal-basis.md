+++
Categories = ["Interpretations"]
bibfile = "mechphys.json"
+++

A strong constraint for understanding quantum physics can be derived from the attempt to define a **universal basis** space that can be used to represent and compute _any arbitrary physics problem_ that can occur in nature. Logically, Nature must somehow operate within such a basis space, assuming that it is not dynamically reconfiguring itself on the fly for different situations at different points in space / time. The value of this universal basis constraint is that it is entirely abstract and functionally-defined, and has no dependencies on intuitions based on macroscopic physics, and yet it appears to be strongly constraining.

Indeed, there are no existing frameworks that satisfy this constraint for the full scope of the [[Standard Model]], and the goal of the [[Spinfield Model]] can be see as providing such a universal basis space, based on the [[cellular automaton]] (CA) framework, which serves as a recipe for constructing such universal basis representations.

The universal basis constraint can be used to determine if a given framework is a [[calculational tool]], most of which involve the use of _bespoke_ representational spaces that extract the most relevant and convenient basis for calculation of a given problem. For example, the use of $r$ in the $1/r^2$ inverse-square-law form of the Coulomb electric force law (or Newton's law of gravity) is an example of a non-universal representational variable. Do we imagine that Nature is somehow computing and storing these _r_ values for each possible pair of charged particles in the universe? And updating them in real-time as they move? And what happens when such particles are destroyed or created?

In short, is clear that this is not a universal basis representation, _because the representation must be updated in highly non-local ways, with a required capacity that is both highly variable over time and exponential in size_. The need to associate particles with their properties in a non-local way creates a significant logical problem: do we imagine there is a giant lookup table somewhere, cross-referenced against every other particle? How does all of this drive local forces experienced by each such particle, in an efficient, parallel manner?

These are fundamentally **information processing** issues, which are independent of any given algorithm or detailed programming language choices: the issue is about the availability of the _information_ needed to drive the forces that then update the state at the next moment in time. This kind of abstract information processing framework is what the cellular automaton provides: a way of organizing the data and the computation so that everything can happen locally, in parallel, with fully determined finite computation at each location and time step -- no open-ended search or long-range variable lookup routines required.

[[Wave]] equations provide a suitable universal basis representation for electromagnetism (EM) in the Lorenz gauge ([[Maxwell]]), and likewise for general relativity. All of the relevant state is available locally to drive the updating of the wave state over time, and a fixed set of state variables organized in 3D space can be used to represent any arbitrary EM configuration that can occur in Nature.

According to these criteria, [[configuration space]] is another example of a bespoke calculational space, where the dimensionality and basis are chosen specifically for the problem at hand (often a simple 2D spin or qubit space), and the tensor product nature results in complete expressive power to capture any kind of interaction or correlation. However, as discussed on that page, there is no plausible configuration space that provides a universal basis ([[@Wallace21]]), due to the same kinds of issues as the inverse-square-law.

The **complementarity** of variables in quantum physics (e.g., position vs. momentum) also poses a specific challenge for finding a universal basis. For example, the Fock space used in [[field theory]] adopts the momentum basis by using Fourier space, where the momentum of a given particle can be represented precisely. But, to keep things tractable, such particles are represented in a "pure" manner, as an infinite sine wave, so any information about a localized position is lost. Adding such information to the Fourier space requires adding many additional sine waves at different phases and frequencies, which then creates a significant burden for tracking all of these associated waves and calculating over them. Likewise, spin and other 2D basis spaces do not represent either momentum or position information.

## Position as a universal basis

J. S. [[Bell]] persuasively argued that all measurement in quantum physics ultimately boils down to measuring positions of massive particles ([[fermion]]s), thereby establishing the privileged status of the position basis ([[@Bell82]], p. 996):

> The second moral is that in physics the only observations we must consider are position observations, if only the positions of instrument pointers. It is a great merit of the de Broglie-Bohm picture to force us to consider this fact. If you make axioms, rather than definitions and theorems, about the "measurement" of anything else, then you commit redundancy and risk inconsistency.

This idea was further developed and formalized ([[@DaumerDurrGoldsteinEtAl96]]; [[@DurrGoldsteinZanghi04]]), all of which is based on the [[pilot-wave]] framework, which provides an unambiguous framework for understanding what a measurement is actually doing, as compared to the standard [[Copenhagen]] framework which leaves this entirely underspecified.

## Other proposals

There does not appear to be a significant literature on the goal of establishing universal basis space representations of the form defined above. [[@^Stoica25]] appeals to the use of a basis space for the universe in deriving the Born probability rule. The requirement of a representation to make sense of a situated observer within the formalism is also related ([[@Barrett21]]), and is also used to support the [[pilot wave]] approach.

