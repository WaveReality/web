+++
Name = "QED"
bibfile = "mechphys.json"
+++

The most successful quantum model in terms of generating precise empirical predictions is **quantum electrodynamics (QED)**, which introduces significantly different conceptual and mathematical frameworks compared to the relatively simple [[Hilbert space]] matrix mechanics and [[Schrodinger]] wave functions. There are two key conceptual elements to QED: the path integral and quantum field theory, which enable it to accurately represent the detailed interactions between an electron and the EM field.

The _path integral_ method developed by Richard Feynmann (originated by Dirac) works by iteratively enumerating all the possible events that might take place in an interaction between electron and photon, and computing the total probability of different outcomes by integrating across all these paths. This method is intuitively illustrated by the famous _Feynmann diagrams_ showing the different possible scenarios in each path. Critically, at this level, the key elements of the theory are the coupled [[Dirac]] equation for the electron and [[Maxwell]]'s equations for the EM field.

A central problem with this path integral framework is that the path integral sums diverge into infinity. A kind of mathematical "hack" (according to Feymann and others) called _renormalization_ was finally able to resolve this divergence. Despite its seemingly arbitrary "pragmatic" basis, renormalization proved successful and allowed the framework to be extended to other domains, involving the weak and strong forces. Clearly, a mathematical framework based on divergent infinite integrals is not a particularly promising basis for a plausible physical model, and is thus clearly a [[tools vs models|calculational tool]].

At higher energy levels, the possible paths include the spontaneous creation of new particles out of raw energy (i.e., _pair production_). Dealing with these kinds of events in an efficient way mathematically required an entirely new framework called **quantum field theory (QFT)**, which does away with the configuration-space representation based on a specific number of particles, and instead considers an entire spatially-extended field where particles can more easily be created and destroyed as "excitations of the field".

{id="figure_fourier" style="height:20em"}
![The Fourier transform, which is the basis of the photon model in QED. A fourier transform converts a function from normal physical space into an orthogonal basis space of sine waves parameterized according to their amplitude, phase and frequency. No position parameter is retained in Fourier space, as the sine waves are infinite in extent. The QED model of the photon is, implausibly, one of these Fourier sine waves, similarly infinite in extent, without any physical localization.](media/fig_fourier_transform.png)

The mathematics of this field are based on a _Fourier space_ representation ([[#figure_fourier]]), where each "particle" is associated with a specific Fourier component, which is a sine wave having a specific frequency and phase. Adding or removing a particle amounts to just adding or removing a corresponding Fourier component.

Fourier space is the _momentum space_ representation complementary to _position space_: it is defined in terms of frequency and phase coordinates, where the frequency is proportional to the momentum of a particle as we saw in the discussion of the original motivation for matter waves by de Broglie. This Fourier representation is particularly convenient for keeping track of total energy, which is proportional to momentum and frequency, and allows one to more easily represent different types of particles splitting up some total amount of energy in different ways.

As captured in the Heisenberg [[uncertainty principle]], momentum and position are "conjugate" (complementary) variables, so when you represent everything in precisely in Fourier space, the position information is completely lost. _Thus, "particles" in Fourier space have no positions: they spread across the entire space._

Position information only arises in Fourier space in terms of the constructive and destructive interference effects of different phases and frequencies. Thus, representing a particle with some specificity of spatial position requires many different sine waves "working together" to add up in one part of space and cancel out in other parts of space.

Representing something with a fully precise spatial position requires an _infinite_ number of such sine waves, in the same way that representing a fully precise momentum (i.e, frequency) requires a continuum infinity of particle positions oscillating according to a specific precise frequency. This is again the Heisenberg uncertainty principle, and it is a basic property of [[waves]].

Under the primary postulate of the [[pilot-wave]] framework, that particles have a specific position at all times, it is clear that this QFT Fourier space representation is exactly the wrong one for describing a physically accurate model. But, as usual, people have a difficult time recognizing that this QFT framework is just another [[tools vs models|calculational tool]] that is convenient for computing certain kinds of problems, and they end up thinking of the QFT particle representation as a physical model.

Specifically, QFT provides the only mathematically tractable model of a photon as some kind of discrete particle-like entity, but for all of the reasons explained above, such a photon is nothing like any other kind of particle we would recognize in the real world. It has infinite spatial extent, and is defined only in terms of its frequency and phase.

In summary, QED and QFT are amazing calculational tools that have been used to make some of the most accurate predictions in all of physics. But they have a number of obvious problems as physical models for how nature actually operates, especially relative to the assumptions of the [[pilot-wave]] framework. Nevertheless, the coupled [[Dirac]] and [[Maxwell]] fields at the heart of this theory are clearly the key basic ingredients, and it is possible that a more computationally-based approach, with the strong constraint of discrete electron particles in specific spatial locations, could provide the basis for an accurate physical model.

