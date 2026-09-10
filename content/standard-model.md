+++
Name = "Standard Model"
bibfile = "mechphys.json"
+++

The **Standard Model** of quantum physics defines a set of fundamental [[particle]]s and their properties, which interact according to force fields carried by [[boson]]s. It has been tested empirically in high-energy collider and cosmic-ray experiments, and through its astrophysical implications, and has been found to be consistent with all of the data accumulated to date.

This serves as the target of explanation for this project. All of the pages describing elements of this Standard Model, including more basic physics principles, are collected under this heading.

The mathematical structure of the Standard Model is based on a set of 7 [[conservation]] laws, 4 of which derive from [[special relativity]] and the remaining 3 use [[gauge theory]] to develop force interactions that conserve the relevant source charge and the energy it generates. These force interactions describe [[wave]] fields associated with each of the fundamental particles, each characterized by [[Lagrangian]] terms. They are derived by introducing a local gauge invariance constraint into a source field, which then defines associated parameterized, independent source currents into a field at each point in space.

When viewed at this abstract level in terms of a system defined by a small set of conservation laws and the basic constraints of special relativity, it is more plausible to imagine that there could be another way of implementing these same conservation laws, and the relevant interactions between waves and particles, that would then produce the same overall behavior. Indeed, much of the complexity of the Standard Model comes from the quantum [[field theory]] mechanisms used to generate quantitative predictions, whereas its overarching principled foundations are remarkably simple and general.

As one salient example of how different levels of analysis can give the same results, a large amount of the work done in quantum physics uses the much simpler, older, abstract framework of the [[Schrodinger]] wave equation and the [[Hilbert space]] formalism, which capture the most basic level of conservation laws. This abstract framework is appropriate for _low energy_ situations in the "ordinary world" where electrons and protons remain stable over time, they are moving much slower than the speed of light, and the distance scales are relatively large, such that the relevant force levels are relatively low.

The central features of this abstract framework, which are carried through in the more complex and detailed Standard Model, are:

* The **unitary** time evolution of the Schrödinger equation captures the fundamental conservation laws: nothing is fundamentally ever gained or lost, it just _rotates_ around the [[complex number]] plane.

* The **wave** nature of physics is manifest by the wave properties of the Schrödinger equation, capturing the key property of linear **superposition**, and the **phase-based** properties of **interference** and **amplification**. It also captures the basic relationships between wave **frequency** and **energy, momentum, and rest mass**, as in the [[Planck relation]] and [[special relativity]]. In this respect, the [[Klein-Gordon]] wave equation is recognized as the more general, relativistically-correct version of the Schrödinger equation, which must be therefore understood as a simplified [[tools vs models|calculational tool]].

* To enable the linear Schrödinger equation to represent these wave-based, phase-driven correlations across different particles, the exponentially-large [[configuration space]] is used to encode the quantum state. There are multiple reasons to believe that the exponential size of this space is indeed exponentially excessive, as discussed in that page. The Standard Model typically uses a different kind of state space, known as a Fock space, which represents the momentum / energy of the relevant particles in terms of a Fourier space. Creation and annihilation operators add or remove sine waves to this space, with frequencies corresponding to the relevant particle energies. This infinite-dimensional state space must also be recognized as a calculational tool, which makes the math much simpler.

The full complexity of the Standard Model relative to this more abstract framework is needed for explaining what happens in high-energy particle accelerators and cosmic events, when particles start changing into other particles, or disappear completely into a flash of energy.

The ultra-microscopic realm of the Standard Model is thus the appropriate level for understanding what is really going on "under the hood", and is therefore the primary focus here. So far, there are many indications that some of the "messy" details present in this model provide useful constraints and inspiration for the [[Spinfield Model]], which are otherwise absent in the abstract framework.

## Pedagogical presentation

The following sequence of pages develops an understanding of the wave equations up through the [[Dirac]] equation for the [[electron]] and all other [[fermion]] particles, and is thus the foundational equation of the Standard Model:

* [[Conservation]] describes how essentially the entire Standard Model can be seen as a consequence of basic conservation laws: conserving energy, charge, spin, etc. This provides a broad principled framework for understanding how a different mathematical model that exhibits the same conservation properties could thus account for the same data.

* [[Wave]]s are the natural consequence of conservation laws, and this page provides the essential foundation for understanding the phenomenology and mathematical formulation of waves, leveraging the [[cellular automaton]] (CA) framework to understand concretely how waves emerge from local neighborhood interactions. This CA framework is used for the graphical, interactive [[Simulations]] that enable a concrete understanding of wave dynamics.

* [[Maxwell]] presents the wave equations for the electromagnetic (EM) force field, as a classical field, implemented through CA-based wave functions. [[Special relativity]] explains this foundational pillar of physics, which can be seen as a natural consequence of the wave equations.

* The [[Klein-Gordon]] (KG) equation provides the simplest version of a relativistically-accurate quantum matter wave function, which can be derived directly from the [[Hamiltonian]] based on the total energy expression in special relativity. The KG equation is arguably the single most important equation in all of quantum physics, as it provides the starting foundation for all the other wave functions, from the simplified  [[Schrodinger]] equation to the more complex [[Dirac]] equation.

* In [[complex KG]], we apply the KG wave functions to a complex-valued wave state, which enables a conserved quantity that acts like electrical charge to be computed. Although the [[complex number]]s involved seem mysterious, we see that they are just a mathematical convenience, and we can eliminate them entirely in our second-order wave functions, which are computed using only real-valued numbers. This complex KG system can be directly coupled to [[Maxwell]]'s equations using the [[gauge theory]] principle of establishing a _local_ form of gauge invariance. This gets us very close to the goal of capturing all of the properties of an [[electron]].

* Interestingly, this complex KG system provides a good model of the [[Higgs]] field, and the somewhat mysterious phenomenon of _spontaneous symmetry breaking_, which ends up giving mass to all massive particles. Understanding this dynamic is critical for developing the Spinfield Model.

* The final step in the wave function development sequence is the second-order [[Dirac]] wave function, which builds on the complex KG system, and captures the phenomenon of [[spin]]. This then provides a physically complete description of the quantum dynamics of a [[fermion]] particle like the electron, and is largely responsible for the amazing predictive accuracy of the _quantum electrodynamics_ [[QED]] framework within the Standard Model.

* Using wave equations as building blocks, the full Standard Model can be specified in terms of a [[#Lagrangian]] equation, representing an interconnected system of partial differential equations. These are all expressed in [[four-vector]] notation with differential operators specified by the [[gauge theory]] transformations that implement the 7 conservation laws. In this form, it is entirely _classical_ and continuous -- no different in any essential respect from Maxwell's wave equations. 

* The classical system of interconnected wave equations is _nonlinear_ in the weak, strong, and Higgs interaction terms ([[@GreinerMuller96]]; [[@Pich12]]; [[@Hansson00]]), and does not provide a suitable mechanism for representing the creation and annihilation of discrete particles, or for analyzing the probabilities of different kinds of _localized_ interactions among discrete particles. All of these problems are managed by using a _perturbation_ framework that iteratively peals away different "layers" (polynomial orders) from the otherwise intractable nonlinear system, using essentially the same power-series logic as the Taylor series.

* The _Fock space_ representation of quantum [[field theory]] provides a purely momentum-space, Fourier representation of all the wave functions, in combination with the [[Feynman diagram]]s that enumerate all the possible discrete, localized, particle-based interactions. By working with this combination of the discrete, localized Feynman diagrams to construct the momentum-space Fock representation of the waves used in the field theory, both the particle and wave [[duality]] aspects of the quantum world are integrated together.

* The [[path integral]] formulation in which the Feynman diagrams are used considers all possible effects that might take place when two or more particles interact, which is used to construct a corresponding _scattering matrix_ (_S-matrix_), computing the aggregate probabilities of different inputs and outputs to a scattering event. This approach entails the use of [[virtual particles]] to mediate force-field interactions. Thus, instead of a continuous wave-particle interaction "naturally" evolving dynamically over time in the space of the differential wave equations, these interactions are modeled as discrete sequences of creation and annihilation of virtual particles.

* Each such particle is just a mode added or subtracted from the Fourier Fock space, with a given frequency corresponding to its energy (per the [[Planck relation]]). At higher energies, more such virtual particles can be involved in a given interaction. Critically, if particles are actually interacting, then the length scales involved presumably get quite close, causing the force strengths to go up in magnitude (per the $1/r$ nature of the forces), resulting in the potential for infinite sums of ever-higher energy virtual particles to be involved.

* Managing these infinite sums requires a [[renormalization]] procedure, where the infinities are swept away by using the empirically-observed _dressed_ values of particle masses and charges, and effectively working backward from these known values to estimate the _bare_ parameters that show up in the original wave equations. The seemingly unprincipled, empirical nature of the renormalization process has caused considerable concern for many physicists, but others have embraced it as an important constraint that evidently works to describe Nature.

Thus, overall, the Standard Model represents a powerful mathematical framework that is at once spectacularly successful for computing highly accurate predictions for the outcomes of all known quantum physics experiments, and yet remarkably "kludgy" in some ways. Much of the difficulty is associated with managing the discretized accounting required by the field-theory framework and the consequent need for the renormalization procedure. But there are also many aspects of the model that seem rather strange and inexplicable, especially in the context of the [[weak]] interactions with their parity violations and spontaneous symmetry breaking, and the nature of the [[generation]]s of masses. This situation leaves many scientists dissatisfied, despite the obvious success of the framework. All of the attempts to go beyond the Standard Model, e.g., by doubling-down on the principle of symmetry ("supersymmetry") have been unsuccessful to date.

## Components

{id="figure_particles" style="height:40em"}
![Elementary particles of the Standard Model. There are also antiparticle copies of all of the matter particles, with opposite charge values. All particles are actually represented by wave fields, and the particle is represented as frequency mode in these fields, corresponding to its energy. There are three colors of each quark, corresponding to the color charge property. Adapted from [MissMJ/Wikimedia](https://commons.wikimedia.org/wiki/File:Standard_Model_of_Elementary_Particles.svg).](media/fig_standard_model_particles.png)

The Standard Model is composed of massive [[fermion]] [[particle]]s, and force fields ([[#figure_particles]]), which are characterized in terms of [[boson]]s as force-carrying particles. However, all of these are represented using various forms of wave equations ([[Dirac]] for the fermions), with no explicit representation of discrete particles at the level of the underlying Lagrangian formulation. Particles enter through the discretized path-integral framework based on Feynman diagrams, and are operationalized by adding or subtracting corresponding frequency modes to the quantum [[field theory]] fields, which only represent the momentum-space (Fourier space) picture.

The complete specification of all the elements of the standard model is a bit complicated, because it varies depending on the way that the spontaneous symmetry breaking of the weak interaction is treated, and a few other complexities addressed in this [stackexchange](https://physics.stackexchange.com/questions/259393/number-of-degrees-of-freedom-in-the-standard-model-lagrangian) post. Following that logic there, we have:

* 7.5 Dirac fields for the fermions, per mass [[generation]] = 22.5 for all three generations: 1 electron, 1/2 for the left-handed neutrino (the right-handed does not exist), 3 for the up quark (3 colors), and 3 for the down quark (3 colors). Each Dirac field has 4 complex-valued state variables, so that is 90 total complex numbers, or 180 real-valued numbers.

* 28 force field unique degrees of freedom, which can be obtained by counting the bosonic particles: 2 x 9 = 18 for the zero-mass photons and gluons, and 3 x 3 generations for the massive weak bosons = 9, plus one Higgs boson. Because these are all second-order wave fields, the derivatives of these fields are also relevant state variables, so it is 56 total.

Thus, there are a grand total of **236 unique real-valued numbers** required to specify the total unique state of the Standard Model.

## Feynman diagrams

{id="figure_feynman" style="height:40em"}
![Elementary force interactions in the Standard Model. All Feynman diagrams in the model can be constructed from combinations of these vertices. **q** is any quark, **g** is a gluon, **X** is any charged particle, **γ** is a photon, **f** is any fermion, **m** is any particle with mass (with the possible exception of the neutrinos), $m_B$ is any boson with mass. In diagrams with multiple particle labels separated by /, one particle label is chosen. In diagrams with particle labels separated by | the labels must be chosen in the same order. For example, in the four boson electroweak case the valid diagrams are WWWW, WWZZ, WWγγ, WWZγ. Adapted from [Ponor/Wikimedia](https://commons.wikimedia.org/wiki/File:Standard_Model_%E2%80%93_All_Feynman_diagram_vertices.svg).](media/fig_standard_model_feynman.png)

[[#figure_feynman]] shows the elementary force interactions in the Standard Model, using [[Feynman diagram]]s. Each such diagram corresponds to [[#Lagrangian]] terms that must be added to capture the dynamics of a particular interaction (described in the next section), and it specifically does _not_ represent an actual spatial trajectory that one might observe in a given collider experiment. 

These diagrams reflect the particle nature of the wave-particle [[duality]], by discretizing the interactions in terms of specific incoming and outgoing particles, and computing scattering matricies (S-matrix) that reflect possible interactions that would take place at close range. These interactions are subject to energy thresholds and other conservation principles that reflect the known rest mass and other quantum properties of the relevant particles.

Although these diagrams do not represent spatial particle trajectories, there is nevertheless an implicit localization that reflects the particle nature, in contrast to the Fourier space representation of the quantum field, which has no spatial localization at all. It is this joint calculation across these complementary basis representations, with the localization remaining implicit in the interactions that are computed, that enables the framework to capture the relevant phenomena.

## Lagrangian

The Standard Model can be specified using a single equation, which is a [[Lagrangian]] expression for all of the interacting fields. It can be written compactly, in the electroweak (pre symmetry-breaking) form, as a sum of multiple factors. In our treatment, we are only concerned with the [[lepton]]s, not the [[quark]]s, so that cuts out several of the factors at the end. In general, there are dynamical terms that specify the update equations for each field, and mass terms that function like the mass drag factor in the [[Klein-Gordon]] equation, which come directly from coupling to the [[Higgs]] field.

The total Lagrangian is a sum of these terms:

$$
\mathcal{L}_{SM} = \mathcal{L}_{BW} + \mathcal{L}_l + \mathcal{L}_{lm} + \mathcal{L}_{nm} + \mathcal{L}_\phi + \mathcal{L}_G + \mathcal{L}_q + \mathcal{L}_{qmd} + \mathcal{L}_{qmu}
$$

Note: the following math was greatly facilitated by this [latex source](https://github.com/SodiumIodide/Standard-Model-Lagrangian).

### Leptons, electroweak

The electroweak model is derived in terms of a $B_\mu$ force field, known as a _weak hypercharge_ gauge field, that functions as the gauge boson vector potential, entirely analogous to the electromagnetic potential $A^\mu$ from [[Maxwell]]. This is supplemented with a _weak isospin_ field $\mathbf{W}_\mu$ that has 3 components, which will end up interacting with the Higgs field to gain mass. The actual post-symmetry-breaking W and Z weak fields, and the observed A electromagnetic field, are a mixture of the original B and W factors.

{id="eq_forces" title="force fields: EM, weak"}
$$
\mathcal{L}_{BW} = -\frac{1}{4} B_{\mu\nu}B^{\mu\nu} - \frac{1}{8}tr\left(\mathbf{W}_{\mu\nu}\mathbf{W}^{\mu\nu}\right)
$$

These terms:

$$
B{\mu\nu}=\partial_\mu B_\nu-\partial_\nu B_\mu
$$

essentially provide a more complicated way of writing the wave equation operating on the four-potential $B_\mu$, exactly as in Maxwell's equations. 

The expression for the $\mathbf{W}_{\mu\nu}$ is more complex, involving a 2x3 traceless Hermitian matrix..??

$$
\mathbf{W}_{\mu\nu} = \partial_\mu\mathbf{W}_\nu - \partial_\nu\mathbf{W}_\mu + ig_2 \frac{\left(\mathbf{W}_\mu\mathbf{W}_\nu - \mathbf{W}_\nu\mathbf{W}_\mu\right)}{2}
$$  

Then we have the lepton particle dynamical terms, which are implicitly extended to include all three [[generation]]s (electron, muon, tau), for the charged particle ($e$ = electron) and the uncharged neutrino, $\nu$. The $L$ and $R$ suffixes indicate the left and right chiral spin orientations -- there is no right spin version of the [[neutrino]], so it is only specified for the electron.

{id="eq_lepton" title="lepton dynamical terms"}
$$
\mathcal{L}_l = \left(\bar{\nu}_L,\bar{e}_L\right) \tilde{\sigma}^\mu iD_\mu \binom{\nu_L}{e_L} + \bar{e}_R\sigma^\mu iD_\mu e_R + \bar{\nu}_R\sigma^\mu iD_\mu\nu_R + \text{(h.c.)}
$$

where (h.c.) indicates to add the Hermitian conjugate (complex conjugate transposed) of the preceding terms. This is specifying a first-order version of the [[Dirac]] wave equations. In the second-order version, the Hermitian conjugates are multiplied directly with the factors shown above.

The local gauge transformation process introduces these covariant derivatives:

$$
D_\mu \binom{\nu_L}{e_L} = \left[\partial_\mu-\frac{ig_1}{2}B_\mu + \frac{ig_2}{2}\mathbf{W}_\mu\right]  \binom{\nu_L}{e_L}
$$

$$
D_\mu e_R = \left[\partial_\mu-ig_1B_\mu\right]e_R
$$

$$
D_\mu\nu_R = \partial_\mu\nu_R
$$

And then the lepton mass terms in coupling with the Higgs:

{id="eq_lepton-mass" title="lepton mass terms"}
$$
\mathcal{L}_{lm} = -\frac{\sqrt{2}}{\nu}\left[\left(\bar{\nu}_L,\bar{e}_L\right)\phi M^e e_R + \bar{e}_R\bar{M}^e\bar{\phi}\binom{\nu_L}{e_L}\right]
$$

plus a neutrino mass term (not officially part of the original Standard Model)

{id="eq_neutrino-mass" title="neutrino mass term"}
$$
\mathcal{L}_{nm} = \frac{\sqrt{2}}{\nu}\left[\left(-\bar{e}_L,\bar{\nu}_L\right)\phi^{*}M^\nu\nu_R + \bar{\nu}_R\bar{M}^\nu\phi^T\binom{-e_L}{\nu_L}\right]
$$

and the Higgs field dynamical and mass term, where the mass term is what causes the spontaneous symmetry breaking:

{id="eq_higgs" title="Higgs dynamical and mass term"}
$$
\mathcal{L}_h = \overline{\left(D_\mu\phi\right)}D^\mu\phi - \frac{m_h^2\left[\bar{\phi}\phi-\frac{\nu^2}{2}\right]^2}{2\nu^2}
$$

$$
D_\mu\phi = \left[\partial_\mu + \frac{ig_1}{2}B_\mu+\frac{ig_2}{2} \mathbf{W}_\mu\right] \phi
$$

$\phi$ is a 2-component complex Higgs field. Since $\mathcal{L}$ is SU(2) gauge invariant, a gauge can be chosen so $\phi$ has the form:

$$
\phi^T = \frac{\left(0,\nu+h\right)}{\sqrt{2}}
$$

$$
\phi_0^T = \text{expectation value of} \; \phi = \frac{\left(0,\nu\right)}{\sqrt{2}}
$$

Where $\nu$ is a real constant such that:

$$
\mathcal{L}_\phi = \overline{\left(\partial_\mu\phi\right)} \partial^\mu\phi - \frac{\mu\phi-m_h^2\left[\bar{\phi}\phi - \frac{\nu^2}{2}\right]^2}{2\nu^2}
$$

is minimized, and $h$ is a residual Higgs field
    
### Quarks, color force

This section describes the [[quark]] specific factors, which we are currently ignoring.

{id="eq_color" title="color force"}
$$
\mathcal{L}_G = -\frac{1}{2}tr\left(\mathbf{G}_{\mu\nu}\mathbf{G}^{\mu\nu}\right) 
$$

$\mathbf{G}_\mu$ is a $3\times3$ traceless Hermitian matrix.

{id="eq_quark" title="quark dynamical term"}
$$
\mathcal{L}_q = +\left(\bar{u}_L, \bar{d}_L\right)\tilde{\sigma}^\mu iD_\mu \binom{u_L}{d_L} + \bar{u}_R\sigma^\mu iD_\mu u_R + \bar{d}_R\sigma^\mu iD_\mu d_R + \text{(h.c.)}
$$

$$
D_\mu \binom{u_L}{d_L} = \left[\partial_\mu-\frac{ig_1}{6}B_\mu + \frac{ig_2}{2}\mathbf{W}_\mu+ig\mathbf{G}_\mu\right] \binom{u_L}{d_L}
$$

$$
D_\mu u_R = \left[\partial_\mu+\frac{i2g_1}{3}B_\mu + ig\mathbf{G}_\mu\right] u_r
$$

$$
D_\mu d_R = \left[\partial_\mu - \frac{ig_1}{3}B_\mu + ig\mathbf{G}_\mu\right] d_R
$$

{id="eq_quark-dsm" title="down, strange, bottom mass term"}
$$
\mathcal{L}_{qmd} = - \frac{\sqrt{2}}{\nu}\left[\left(\bar{u}_L,\bar{d}_L\right)\phi M^dd_R + \bar{d}_R\bar{M}^d\bar{\phi}\binom{u_L}{d_L}\right]
$$

{id="eq_quark-uct" title="up, charm, top mass term"}
$$
\mathcal{L}_{qmu} = - \frac{\sqrt{2}}{\nu}\left[\left(-\bar{d}_L,\bar{u}_L\right)\phi^{*}M^uu_R + \bar{u}_R\bar{M}^u\phi^T\binom{-d_L}{u_L}\right]
$$

And there you have it! 

As complicated as these equations are, they actually represent one of the most compact ways of writing them, with multiple implicit expansions required to capture the full model. See this [Wikipedia image](https://en.wikipedia.org/wiki/Mathematical_formulation_of_the_Standard_Model#/media/File:Formula_of_the_Standard_Model_of_particle_physics_-_Lagrangian_L_SM.jpg) for a much more complicated version of these equations.

## Pages

