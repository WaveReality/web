+++
bibfile = "mechphys.json"
+++

Historically, the early development of this wave-like property of particles was focused on understanding the nature of simple atoms like hydrogen, which has a single electron orbiting a nucleus. The dominant classical physical model of the atomic system in the early 1900's was the Rutherford model of 1911, with electrons as tiny points of charge and mass, orbiting a nucleus, much like planets orbiting the sun. This model had important failings, which the full development of quantum mechanics resolved, thus cementing the demise of the classical worldview, and solidly establishing quantum mechanics.

The major problem with the classical atom was that it is fundamentally unstable: the electron should emit electromagnetic radiation as it orbits around the nucleus, and thus lose energy. As it loses energy, the orbit must get tighter, and eventually the electron should just collapse into the nucleus, just like one of those quarters you roll around in a gravity well at a science museum. Furthermore, as its orbit gets tighter, it should emit higher frequency radiation, predicting a continuous and increasingly high frequency emission spectrum. Instead, it was known that atoms emit consistent, discrete frequencies of radiation.

In 1913, Niels Bohr provided an apparent solution to the problem, leveraging the emerging ideas of Planck and Einstein that were derived from the _photoelectric effect_ and the properties of _blackbody radiation_: the energy of "photons" (particle-like elements of EM radiation) is proportional to their frequency (and thus inversely proportional to wavelength), times a mysterious new physical constant named after Planck, $h$:

{id="eq_energy" title="Energy = h times frequency"}
$$
E = h \nu
$$

Bohr postulated that electrons can only have orbits where the angular momentum (i.e., the effective period of the orbit) is restricted to an integer multiple $n$ of Planck's constant:

{id="eq_borh" title="Bohr wavelength = integer multiple of h"}
$$
L = n \hbar = n \frac{h}{2\pi}
$$

Although the reason for this restriction was not clear, it immediately made sense of a great deal of data, including the Rydberg formula for hydrogen emission spectra. The justification for Bohr's restriction on atomic orbits came in 1924, when Louis de Broglie proposed that electrons have a wave-like nature, and thus the only frequencies of electron wave vibration that are stable are standing waves. Standing waves must have an integer number of wavelengths, such that within the orbiting electron model, the electron orbits are constrained to have an integer number of such waves per orbit.

Shortly thereafter, in 1926, Erwin Schrodinger developed his famous wave equation, which then gave a complete mathematical description of the behavior of bound electrons in atomic systems, which made sense of even more data than Bohr's original model.

The experimental confirmation of de Broglie's matter wave hypothesis came in 1927 in an experiment by Davisson and Germer, who found that electrons moving through a crystal exhibit a diffraction pattern, consistent with a wave-like property. Calculations showed that the de Broglie wavelength predicted for the electrons fit the observed diffraction pattern quite well:

{id="eq_momentum" title="Momentum = h / de Broglie wavelength"}
$$
p = \frac{h}{\lambda}
$$

{id="eq_debh" title="de Broglie wavelength = h / momentum"}
$$
\lambda = \frac{h}{p}
$$

This wavelength is about .165 nanometers for the electrons in the Davisson-Germer experiment (very tiny, but enough to produce a measurable diffraction pattern through the crystal).

Both de Broglie and Schrodinger thought that these [[matter waves]] were real physical things, like light waves. Furthermore, de Broglie suggested that the wave acted to _guide_ the point particle electron around, in his [[pilot-wave]] theory. Schrodinger initially had an even more radical view, which abandoned the point electron entirely: he thought his wave equation described a wave of _charge density_ that _is_ the actual electron, without any need for a dual particle-like entity.

However, both of these attempts to provide a "physically realistic" perspective on the phenomena were quickly abandoned in the face of further evidence suggesting that the wave function fundamentally describes the _probability_ that a particle might appear at a given point in space when measured. As such, the wave is somehow "non physical",  and yet exerts physically-measurable effects.

This movement away from a physical model was furthered by Heisenberg and his [[uncertainty principle]], and the broader principle of _complementarity_: there is a minimum, irreducible amount of ambiguity or uncertainty in the quantum world, and you can _either_ push things more toward one perspective (e.g., the wave-like aspect) _or_ the other (the particle-like aspect), but not both at the same time. This philosophical stance was based initially on the simple fact that the uncertainty in the position ($x$) or momentum ($p$) of a particle has a minimum bound of a factor of Planck's constant:

{id="eq_uncertainty" title="Uncertainty in position * momentum"}
$$
\sigma_x \sigma_p \geq \frac{\hbar}{2}
$$

Thus, any attempt to decrease the uncertainty in location $\sigma_x$ necessarily increases the uncertainty in momentum $\sigma_p$. This can be seen as a natural consequence of [[matter waves]], and more generally, because a wave is a spatially distributed thing, it is very hard to pin down precisely.

From this complementarity principle, Bohr and Heisenberg developed the [[Copenhagen]] interpretation of QM in the late 1920's, and this is still dominant today. Central to this interpretation is the notion that the physical world operates in two complementary modes: you are _either_ making a _measurement_, which causes the wave function to _collapse_ down to a single discrete particle-like point, _or_ physics is otherwise evolving according to the wave function, which critically preserves all the quantum uncertainty, and just rotates it around in a _unitary_ manner over time.

During this wave-function mode, the mathematical picture suggests that there is no definitive underlying state of the world: everything is in some kind of probabilistic superposition of possible states. Only once you measure something does it actually exist in any kind of definite way, leading to the mantra that "the world only exists when you measure it".

This strong discretization of the laws of physics is at the root of many seeming paradoxes and puzzles in understanding the quantum world: what exactly defines a "measurement" at a fundamental level?  How can the wave function, which could conceivably spread out over large macroscopic spaces over time, instantaneously collapse down to a single point within that entire space?  Despite these conceptual difficulties, the mathematics of the framework allow straightforward calculations that match the outcomes of actual experiments, leading to a general attitude of "shut up and calculate": don't bother with unnecessary considerations of the actual underlying physical ontology, just do the math!

However, in the 1950's, David Bohm reinvented the original _pilot-wave_ model of de Broglie, and showed that in fact it can fully explain the same phenomena as the standard Copenhagen QM framework. Critically, in this _de Broglie-Bohm pilot-wave_ framework, _particles always have a definite well-defined location_. There is no longer a complementary discretization of the world into measurement vs. wave-function evolution phases: the two are _always_ operating hand-in-glove, all the time.

{id="figure_double-slit-deb" style="height:20em"}
![Trajectories for particles in the double-slit experiment computed according to the de Broglie-Bohm pilot-wave model. The interference effects can be seen as relatively localized bumps in the trajectories, corresponding to steep gradients in the Schrodinger wave equation. Critically, the underlying trajectories are considered to exist at all points even if you don't happen to observe them.](media/fig_double_slit_debroglie_bohm.png)

[[#figure_double-slit-deb]] shows what the underlying trajectories of particles under the pilot-wave framework look like in a double-slit experiment, and [[#figure_double-slit-kocsis]] shows some recent data from an experiment where _weak measurements_ that minimally disturb the system allow one to infer particle trajectories, which look remarkably similar to those predicted by the pilot-wave model ([[@KocsisBravermanRavetsEtAl11]]).

{id="figure_double-slit-kocsis" style="height:20em"}
![Reconstructed trajectories of photons in a double-slit experiment using a weak measurement technique that allows aggregate trajectory information to be reconstructed over many repeated samples that are post-sorted according to a weak additional modulation of the system --- these are not individual particle trajectories. There is a striking correspondence to the predictions of the de Broglie-Bohm model. Figure from Kocsis et al, 2011.](media/fig_double_slit_kocsis_et_al_11.png)


