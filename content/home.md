+++
URL = ""
Title = ""
bibfile = "mechphys.json"
+++

<img src="media/icon.png" style="width:128px;height:128px;align-self:center">

This is a work-in-progress wiki-like collection of documentation in support of the development of a computational model of the phenomenology of quantum electrodynamics (QED), based on the coupled [[Dirac]]-[[Maxwell]] wave functions, along with discrete [[electron]] particles, consistent with the [[pilot-wave]] framework of Bohm and de Broglie. This computational model is based on the [[cellular-automaton]] framework, which is arguably the simplest way that physics could autonomously emerge in parallel, everywhere in the universe, all at once.

The primary goal of this project is to better understand the basic physics of electrons interacting with the electromagnetic field, and to try to sort through some of the notorious paradoxes and conceptual challenges that lie at the heart of quantum mechanics (QM). There is an easy-to-use GUI-based simulation software package (note: currently under development) that allows one to interactively explore various physics models, providing a concrete and hands-on level of understanding. This provides a different and potentially valuable set of tools for someone trying to learn more about how quantum physics actually works, which may result in quicker and deeper understanding than staring at equations :)

From an _ontological_ perspective (i.e., in terms of what we think is _actually real_), there are two seemingly-conflicting entities at the heart of QM: _waves_ and _particles_ (i.e., the "wave-particle duality"). Much of the confusion and paradox in QM is tied up with this duality. The overall approach here follows the _de Broglie-Bohm pilot-wave_ framework in embracing the simultaneous and interacting reality of both particles and waves, and we spend a lot of time and effort understanding the seemingly "magical" properties of [[waves]] in particular.

At one point, I was hopeful that somehow one could do away with particles entirely, given that wave equations can account for so much of the known phenomena in physics, including the mind-bending space-time distortions of special relativity, and the complex dynamics of electricity and magnetism (EM). However, pure waves end up being just a bit too "squishy" and tend to ooze out all over the place. Particles provide a kind of hard "skeleton" that is particularly important for maintaining the strict conservation laws that are so central in physics: Nature is above all a meticulous accountant, it seems. For example, there are strong conservation laws for charge, spin, and particle numbers of various sorts, along with the ubiquitous conservation of total energy.

Furthermore, the [[Pauli exclusion principle]] prevents there from being two of the same _fermions_ (spin 1/2 particles like electrons and quarks) in the same quantum state, which in the pilot-wave framework means being in the same place at the same time with the same spin. This suggests from our computational, [[cellular-automaton]] perspective that there is some kind of underlying constraint like "slots" in a lattice for holding at most one of each type of particle. This is both a welcome simplification for our models of these particles, and a tantalizing suggestion that this computational perspective might provide some unique insights into the underlying nature of the physical world.

This computational perspective also provides an interesting motivation for the need for waves. If you just have a simple discrete point-like particle sitting in some kind of lattice-like grid, it is very difficult to implement realistic force-field interactions among such particles, especially when using other discrete particles like "photons" to mediate these interactions. The purely particle picture of an electron constantly spewing baseball-like photons out in all directions to hit other electrons is very difficult to sustain. How does such a scheme ever achieve any kind of smooth field-like coverage of space using discrete point-like entities?  How many balls per femtosecond does it have to spew?  How do they manage to spread out uniformly over space and time, while properly conveying the dynamic interactions among the magnetic and electric aspects of the wave functions?

Instead, it is far more straightforward to use [[maxwell|EM]] wave equations to model the force field interactions among electrons. However, the ability of a discrete localized electron to "sense" such a force field as a distributed wave remains problematic: EM waves that influence electrons are widely distributed things, and small discrete samples at one point of a wave would not provide the proper net influence that the physical laws require. Thus, it works much better for the electron to also have its own wave field that is directly coupled with the EM wave field. In effect, the electron's wave field acts like a kind of antenna that senses and responds to the EM forces, and then conveys the results to shape the unfolding trajectory of the discrete particle through space and time, as captured in the pilot-wave model.

We will ultimately implement this wave-particle model through coupled [[Dirac]] wave functions for the electron and [[Maxwell]]'s equations for EM, with the Dirac wave providing the guiding [[pilot-wave]] for a discrete electron particle localized within a cubic lattice grid. We call this **Wave Electrodynamics** or **WELD**, which attempts to explain the same phenomena as _quantum electrodynamics_ ([[QED]]).

The waves in this model are all implemented using the same cubic lattice grid that the discrete electron particles live on, where local neighborhood interactions among the lattice cells implement a highly spatially symmetric form of the _Laplacian_ spatial gradient function at the core of the wave function. In short, the entire model is essentially an elaborate form of [[cellular automaton]] (CA), which has many appealing properties as the simplest-possible framework for a physical system, as advocated by a number of theorists over the years, including John Von Neumann, Stanislaw Ulam, [[@^Zuse69]], [[@^FredkinToffoli82]], [[@^Fredkin90]], [[@^tHooft05]] (2015); [[@Wolfram97]]).

One of the primary challenges of this CA framework is reconciling the local interactions among neighboring cells, which so naturally produces a relativistic speed-of-light limit (one time step update per lattice cell), with the now irrefutable evidence for some kind of non-locality in quantum physics. Recent work within the pilot-wave framework has helped to significantly clarify the nature of these non-local interactions, and the broader conflict that they actually pose for all of QM, despite many attempts to downplay these issues from within the standard QM frameworks ([[@DurrGoldsteinNorsenEtAl14]]).

Without having any definitive results at this point, the approach taken here is to push the current WELD framework along as far as possible, addressing the many significant outstanding problems, and use the resulting computational models to further understand the nature of these non-local interactions. Given that electron particles are strongly interacting with distributed wave functions in this framework, it is not inconceivable that the necessary non-local interactions may emerge from these distributed wave interactions in a way that remains compatible with observed data. For example, it is already clear that the time step needed for updating the wave function computations is actually at least twice the speed-of-light rate in a simple CA, raising the possibility of supra-luminal interactions at the lattice level, while still maintaining fully relativistic dynamics in the EM wave function propagation.

In summary, the WELD approach represents an exploratory computational-modeling approach that could afford important insights into the otherwise puzzling nature of the quantum world. We can already see from the above intuitive arguments that the central wave-particle duality of the quantum world actually makes good sense from the lens of trying to implement these systems in a computational model. One advantage of the computational modeling approach is that it allows one to explore more complex, nonlinear interactions that can be difficult to analyze mathematically. 

Indeed, we will also see many cases where the simplified mathematical frameworks that have been developed to perform useful quantum calculations have been mistaken for physical models of the underlying processes, giving rise to most (if not all) of the apparent paradoxes in the quantum world. By extrapolating from clear examples in other domains, and the considerable progress by a number of researchers in developing a plausible quantum physical models, it is not unreasonable to be optimistic that a fully sensible, intuitively comprehensible physical model of the quantum world is within reach.

The content is organized into two major categories: background, and specific wave functions used in the models, with the following sequence of pages within each category providing a progressive, systematic treatment.

## Background

* [[History]] reviews the history and basic phenomenology of quantum physics (QM).

* [[Tools vs models]] identifies a critical distinction between calculational tools versus physical models, that are intended to provide a "mechanistic", autonomous model of how physics might _actually_ operate. There are typically many different ways to calculate a prediction, but presumably physics is not strategically selecting different calculational tools based on different configurations of elements --- it must be just doing one consistent thing across all of space and time. That thing is what we seek to understand here.

* The contrast between the [[Hilbert space]] formalism vs. [[pilot-wave]] models provides a nice example of this difference between an efficient calculational tool vs. a physical model (respectively). Critically the pilot wave framework requires that the mysterious wave functions of quantum physics _actually exist_ as a real physical thing, which is a central pillar of the current approach.
    
* [[Epistemic vs ontic]] uncertainty discusses a critical distinction between sources of uncertainty that are due to limits on our knowledge of the exact state of an underlying physical system, versus _true_ stochasticity that reflects a fundamental underlying physical process that is _inherently_ stochastic. The present framework requires this latter form of _ontic_ (also known as _aleatoric_) stochasticity to produce physically-accurate behavior of discrete particles like an electron within the CA framework.
 
* [[Contextual]] variables, as contrasted with "real" variables, and their role in understanding the phenomenology of QM.

* [[Semiclassical]] models, that combine a classical treatment of the EM field according to [[Maxwell]]'s equations, with a quantum treatment of the [[electron]]. This is the same approach used in this framework.

* [[QED]] provides an overview of _quantum electrodynamics_ which provides a highly accurate _description_ of the relevant phenomenology of interest. However, we argue that QED is a calculational tool, not a plausible physical model.

* The [[Zero-point]] field and _stochastic electrodynamics_ which provides an alternative formulation of QED, which the present framework shares some important similarities.

* The [[cellular automaton]] (CA) framework for representing space, time and local interactions.

* The concept of a luminiferous [[aether]], which was theoretically disproven by the famous Michelson-Morely experiment of 1887, but in fact is entirely consistent with a privileged reference frame as required by the CA framework, as long as that reference frame obeys the critical time / space distortion properties of [[special relativity]] (which of course it must, to be consistent with well-established empirical data).

* The tricky problem of [[non-locality]] in QM, which appears to strongly violate the local nature of the CA framework. We suggest based on our computational model that quantum wave propagation requires a supraliminal (faster than light-speed) update rate, which is potentially able to account for the observed data. The apparently instantaneous nature of "quantum collapse" has long been a murky area of standard quantum formulations, which adopt a purely epistemic understanding of the quantum wave function, and thus are theoretically not constrained by the idea that these waves must propagate through local mechanisms.

## The WELD model

* [[Waves]] provides the essential foundation for understanding the phenomenology and mathematical formulation of waves, within the CA framework.

* [[Maxwell]]'s equations for the electromagnetic (EM) force field, as a classical field, implemented through CA-based wave functions.

* [[Matter waves]] discusses the general idea of a quantum wave equation that captures something about the properties of massive particles: but what is it actually representing? In the standard QM frameworks, it represents the _epistemic_ probability of finding a discrete particle at a given location at a given point in time. In the [[pilot wave]] framework adopted here, it represents a physically real wave permeating space, that guides the movement of a discrete particle.

* The [[Klein-Gordon]] (KG) equation provides the simplest version of a relativistically-accurate quantum matter wave function, which can be seen as a kind of second-order version of the much more widely discussed [[Schrodinger]] wave function.

* The [[Dirac]]] wave function builds on the KG equation to capture the phenomenon of _spin_, and provides a physically complete description of the quantum dynamics of a particle like the electron. This is what drives all the amazing predictive accuracy of the [[QED]] framework within the standard model of physics, and is what we hypothesize exists as a real physical wave.

* Finally, the [[electron]] is modeled as a discrete particle that moves with _intrinsic_ (_ontic_) stochasticity under the influence of the Dirac wave function.

