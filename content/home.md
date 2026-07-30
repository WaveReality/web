+++
URL = ""
Title = ""
bibfile = "mechphys.json"
+++

<img src="media/icon.png" style="width:128px;height:128px;align-self:center">

**Wave reality** is dedicated to exploring the idea that the quantum wave function is _real_, and not just a description of our state of [[epistemic vs ontic|epistemological]] ignorance. The reality of the wave function is strongly indicated by the classic [[double-slit]] experiment results, where some kind of spatially-distributed wave-like interference phenomenon seems to be influencing the trajectories of discrete particles. In addition, there are increasingly strong theoretical and empirical attempts to show that a purely epistemic account contradicts quantum theory ([[@PuseyBarrettRudolph12]], [[@RingbauerDuffusBranciardEtAl15]]).

The relationship between waves and particles is the central conundrum in quantum physics, and there have been attempts to resolve these issues by somehow eliminating either the particles or the waves. For example, if you carefully examine the mathematical foundations of quantum physics, you find that there are only waves (or wave-like fields), suggesting that particles can be somehow entirely eliminated ([[@Hobson13]]; [[@Sebens22]]).

However, this approach fails to provide explanations for the many particle-like phenomena that have been reliably established, and are implicit in the measurement process of standard quantum frameworks. Thus, even though the equations look like waves, they are used in a way that ultimately describes the behavior of particles. The opposite view, that there are only particles, is even less tenable: as argued above, there are many strong indications that wave phenomena are required (see also [[@Hobson13]]).

The [[pilot-wave]] framework of de Broglie and Bohm ([[@Bohm52]]; [[@Norsen22a]]) instead embraces the **wave-particle** [[duality]] fully, by positing a physically real quantum wave function guiding a discrete particle around. This framework naturally and intuitively explains all of the otherwise strange phenomena in quantum physics. Thus, we adopt this overall framework, and furthermore show that a particular interaction among waves and particles is actually an important _feature_ of how physics must work to "solve" fundamental problems, rather than some kind of paradoxical bug that we just have to somehow learn to swallow.

This website contains a work-in-progress wiki-like collection of documentation in support of the development of a computational model of the phenomenology of quantum electrodynamics ([[QED]]), starting with the coupled [[Dirac]] -- [[Maxwell]] wave functions, along with discrete [[electron]] [[stochastic particles]], consistent with the [[pilot-wave]] framework. This computational model is based on the [[cellular automaton]] framework, which is arguably the simplest way that physics could autonomously emerge in parallel, everywhere in the universe, all at once.

The primary goal of this project is to better understand the basic physics of electrons interacting with the electromagnetic field, and to try to sort through some of the notorious paradoxes and conceptual challenges that lie at the heart of quantum mechanics (QM). An easy-to-use GUI-based [[waves simulator]] is integrated into this content, which allows one to interactively explore various physics models, providing a concrete and hands-on level of understanding. This provides a different and potentially valuable set of tools for someone trying to learn more about how quantum physics actually works, which may result in quicker and deeper understanding than staring at equations :)

This project is fully committed to **accounting for all the empirical data** --- there is no point in coming up with an elegant theory that is obviously false. We cannot impose our own aesthetic preferences onto Nature, and must humbly accept any irrefutable facts that have been reliably empirically established. However, there is a huge underexplored space of possible physical mechanisms that could account for all the data, while avoiding the kinds of obvious conceptual paradoxes that pervade the standard interpretations of quantum physics.

For example, Einstein famously rejected quantum mechanics because "God does not play dice with the universe" --- that seems just a bit presumptuous. Indeed, a key element of the framework we develop here requires true randomness in the form of [[stochastic particles]], and this one concession of nondeterminism can be traded in to retain several other more important (in my opinion) principles, such as a fundamental **locality** to physical mechanisms.

A critical conceptual foundation of this approach is to recognize the distinction between **calculational tools** versus **physical models** ([[tools vs models]]). The widespread failure to understand this distinction underlies many of the apparently intractable puzzles in quantum physics, that can be seen as arising from specific choices of calculational tools. There are typically many different ways to calculate a prediction of an experimental result, but presumably Nature is not strategically selecting different calculational tools based on different configurations of elements.

A fundamental assumption here is that Nature must be doing one consistent thing using one set of physical mechanisms, uniformly and consistently across space and time. That thing is what we seek to understand here.

## Standard interpretations and quantum non-locality

The [[Copenhagen]] interpretation of QM, developed by Niels Bohr and Werner Heisenberg in the 1920's (see [[history]]), is the source of most of the apparent paradoxes and conundrums associated with quantum physics. At the heart of this is the interpretation of the wave function as a purely [[epistemic vs ontic|epistemological]], non-physical entity. If it is not actually real, then it isn't subject to the kinds of constraints that we might otherwise expect it to have, like being local in space and time, which is _required_ by the well-established phenomena of [[special relativity]].

The central premise of the Copenhagen interpretation is that there is a special kind of process called _measurement_ that somehow causes quantum wave functions to _collapse_ down to a single point, with the wave function defining the _probability_ of finding a discrete particle at any given point. Exactly how a potentially widely-distributed wave might somehow gather up all of its far-flung bits and shrink down into one randomly-chosen point, in an instantaneous, manifestly non-local process, is entirely beyond the explanatory scope of the theory.

In the face of these kinds of obviously non-physical aspects of this framework, the standard answer is to "shut up and calculate". This is the hallmark of a calculational tool, and as such, it seems prudent to consider this framework as such, and we will not spend any further effort here probing its fundamental strangeness.

And yet, the Copenhagen interpretation remains the most popular interpretation according to informal surveys of working physicists ([[@Tegmark98]]; [[@SchlosshauerKoflerZeilinger13]]). The next-most popular interpretation after Copenhagen according those surveys is the _many-worlds_ interpretation originated by [[@^Everett57]], which postulates that the entire universe splits at each measurement event. This nominally avoids the need for wave function collapse, but at what cost? An infinite accumulation of new universes spawning everywhere? This is so completely physically implausible that it just defies belief that so many people could even contemplate such a theory, just because it simplifies the math.

Relatively few / none of the respondents in these surveys endorsed the [[pilot-wave]] approach, despite the fact that it eliminates almost all of the strange paradoxes and counterintuitive ideas advanced in the more widely-accepted interpretations. A common objection to the pilot wave framework is that it is defined over [[configuration space]], which grows exponentially in the number of particles. However, every other framework is _also_ defined over configuration space, or it just pushes this exponential explosion into the forking of new universes in the case of the many-worlds framework.

In effect, the other interpretations merely avoid confronting the difficulties of this exponential configuration space by denying its reality in one way or another, while still requiring it to calculate how the physics actually works. Indeed, the entire foundation of the Copenhagen interpretation is willful denial of reality: reality doesn't exist until it is somehow "measured". And there is no explanation for what actually constitutes a measurement, and how instantaneous collapse could possibly be made compatible with special relativity.

From this perspective, the single most important unsolved problem in quantum physics is to eliminate the [[configuration space]] representation, and somehow derive a fully relativistic and local wave-particle [[duality]] where the nonlinear interactions that are otherwise captured in the calculational tool of configuration space naturally emerge. This new framework is likely to be fundamentally nondeterministic (see [[stochastic particles]]) and perhaps requires superliminal, but still finite, wave propagation speeds for the quantum waves. It is increasingly clear that the various existing "no-go" theorems such as Bell's inequalities do not exclude this broader space of possible mechanisms.

## Developing the model

Following from the above background (including linked pages above), the following is a suggested sequence of how to proceed through this content:

* [[Contextual]] variables, as contrasted with "real" variables, and their role in understanding the phenomenology of QM.

* [[Semiclassical]] models, that combine a classical treatment of the EM field according to [[Maxwell]]'s equations, with a quantum treatment of the [[electron]]. This is the same approach used in this framework.

* [[QED]] provides an overview of _quantum electrodynamics_ which provides a highly accurate _description_ of the relevant phenomenology of interest. However, we argue that QED is a calculational tool, not a plausible physical model.

* The [[Zero-point]] field and _stochastic electrodynamics_ which provides an alternative formulation of QED, which the present framework shares some important similarities.

* The concept of a luminiferous [[aether]], which was theoretically disproven by the famous Michelson-Morely experiment of 1887, but in fact is entirely consistent with a privileged reference frame as required by the CA framework, as long as that reference frame obeys the critical time / space distortion properties of [[special relativity]] (which of course it must, to be consistent with well-established empirical data).

* [[Waves]] provides the essential foundation for understanding the phenomenology and mathematical formulation of waves, within the CA framework.

* [[Maxwell]]'s equations for the electromagnetic (EM) force field, as a classical field, implemented through CA-based wave functions.

* [[Matter waves]] discusses the general idea of a quantum wave equation that captures something about the properties of massive particles: but what is it actually representing? In the standard QM frameworks, it represents the _epistemic_ probability of finding a discrete particle at a given location at a given point in time. In the [[pilot-wave]] framework adopted here, it represents a physically real wave permeating space, that guides the movement of a discrete particle.

* The [[Klein-Gordon]] (KG) equation provides the simplest version of a relativistically-accurate quantum matter wave function, which can be seen as a kind of second-order version of the much more widely discussed [[Schrodinger]] wave function.

* In [[complex KG]], we apply the KG wave functions to a complex wave state, which enables a conserved quantity that acts like electrical charge to be computed. Although the [[complex number]]s involved seem mysterious, we see that they are just a mathematical convenience, and we can eliminate them entirely in our second-order wave functions, which are computed using only real-valued numbers. This complex KG system can be directly coupled to [[Maxwell]]'s equations, getting very close to the ultimate goal of capturing all of the properties of an [[electron]].

* The final step in the wave function development sequence is the second-order [[Dirac]] wave function,  which builds on the complex KG system, and captures the phenomenon of _spin_. This then provides a physically complete description of the quantum dynamics of a particle like the electron. This is what drives all the amazing predictive accuracy of the [[QED]] framework within the standard model of physics, and is what we hypothesize exists as a real physical wave.

* Finally, the [[electron]] is modeled as a discrete [[stochastic particles|stochastic particle]] that moves with _intrinsic_ (_ontic_) noise under the influence of the Dirac wave function.

