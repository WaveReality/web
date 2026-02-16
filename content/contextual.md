+++
bibfile = "mechphys.json"
+++

The widely discussed  distinction between **contextual** vs. "real" variables is a source of considerable confusion within the QM world ([[@Shimony84]]; [[@Gudder70]]; [[@Khrennikov01]]; [[@Rovelli96]]. A contextual variable is effectively something that cannot be discretely localized and quantified --- its value depends in some necessary way on the surrounding _context_, which is usually taken to mean the state of the measurement apparatus. In the standard Copenhagen interpretation, _everything_ could be described as being contextual, given that _nothing_ is thought to exist in any localized, definite way prior to the measurement process.

However, in the pilot-wave framework, the positions of the particles (and _only_ these variables) are given a privileged status as being _real_, non-contextual variables, and _everything else_ about the quantum state is relegated to the usual _contextual_ status. Another way of stating this is that everything that must be computed from the wave function itself is contextual, and only the position values are excluded from this status.

For example, [[@^Norsen14]] analyzed a pilot-wave model of _spin_, and showed very clearly that it is contextual in nature. There are no predefined, definite spin values for any particles (indeed this is mathematically impossible as discussed in a moment), and the interaction with a spin-detecting Stern-Gerlach magnetic field apparatus is responsible for _creating_ a definite spin value along a given axis, where none existed previously. The same logic applies to the momentum and energy of the particles, which also depend on the wave function, as elaborated in [[@^Norsen17]].

Is there something fundamental within quantum theory that promotes position variables to this privileged status that they have in the pilot-wave framework?  In general, the answer appears to be "no", although there is ongoing philosophical debate on this topic ([[@Schoeren22]]; [[@Wallace20]]; [[@North12]]; [[@NeyAlbert13]]). In the standard matrix mechanics approaches, the quantum state is chosen to be whatever is most convenient, and this could be a momentum basis instead of position, for example. However, if it turns out that the position variables somehow do provide a uniquely useful basis for "real" variables, as in the pilot-wave model, that would be an intriguing result.

There is a further quantum property that is critical to appreciate in the context of contextuality, which is whether different quantum variables _commute_ with each other or not. Mathematically, two variables $A$ and $B$ commute if their order of application (multiplication) doesn't affect the result:

{id="eq_commuting" title="Commuting variables"}
$$
A B = B A
$$

or:

$$
AB + BA = 0
$$

In the mathematics of matrix mechanics, these variables are actually "operators", that are like a measurement operation that extracts the value of a given variable. In any case, the key physical meaning of commuting variables is that _the order in which you apply the measurements_ doesn't matter. For example, if you measure the position of a particle in the X axis, that does not affect your measurement of the position in the Y axis, so these position operators commute. However, measuring the position of a particle specifically does _not_ commute with measuring its momentum, which is the key point of the Heisenberg uncertainty principle: the more accurately you measure position, the more that destroys your ability to measure momentum, and vice-versa.

Critically, the different components of _spin_ do not commute with each other, which is why it is impossible for any quantum state to have a definite value for each of these different spin components ([[@KochenSpecker90]]; [[@Spekkens05]]). Thus, spin _must_ be a contextual value (as shown in [[@Norsen14]]), where the measurement process effectively rotates the spin axis onto one dimension, which at the same time makes the other axes indeterminate.

{id="figure_polarization" style="height:20em"}
![Demonstration that polarization actually rotates the "photons" in light --- the polarized lens closest to the camera is oriented perpendicular to the polarization of the LCD screen, and thus blocks nearly all of that light. However, the other lens interposed between it and the screen rotates the light by roughly 45 degrees, so that it can then make it through the lens, as seen in their overlapping region.](media/fig_double_polarization_expt.jpg)

This contextual property of spin can also be easily demonstrated for polarization of light, as shown in [[#figure_polarization]]. The first pass through a polarizing sunglass lens effectively rotates polarized light onto that axis, such that it can then pass through a second lens, whereas it isn't able to pass through that second lens directly.

