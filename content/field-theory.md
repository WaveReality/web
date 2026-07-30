+++
bibfile = "mechphys.json"
+++

**Quantum field theory** (**QFT**) is a mathematical framework, also known as **second quantization**, that is essential for handling the higher-energy phenomena where particles are created and destroyed. The standard wave equations such as the [[Schrodinger]] equation can only describe the time-evolution of a particular configuration of particles, but there is no way to represent the creation and destruction of particles.

The [[Dirac]] equation improves over the Schrödinger equation in being fully relativistically correct, and can thus represent the high energy associated with massive particles moving near the speed of light, but it doesn't have a way of actually representing the remarkable fluidity of such particles at high energy, where new particles can be created out of pure energy (typically from the collision of other particles at high speeds, as in an accelerator).

{id="figure_fourier" style="height:20em"}
![The Fourier transform, which is the basis of the photon model in QED. A fourier transform converts a function from normal physical space into an orthogonal basis space of sine waves parameterized according to their amplitude, phase and frequency. No position parameter is retained in Fourier space, as the sine waves are infinite in extent. The QED model of the photon is, implausibly, one of these Fourier sine waves, similarly infinite in extent, without any physical localization.](media/fig_fourier_transform.png)

The mathematics of field theory are based on a **Fourier space** representation ([[#figure_fourier]]), where each "particle" is associated with a specific Fourier component, which is a sine wave having a specific frequency and phase. Adding or removing a particle amounts to just adding or removing a corresponding Fourier component, which is known mathematically as a **Fock space**.

This framework is called _second quantization_ because the Fock addition and subtraction operations are **discrete, all-or-nothing** events. However, as will become clear, the precise mathematical way in which these quantized operations proceed is highly implausible as a representation of the underlying physical processes. Nevertheless, the strong conservation laws require that all the participants in a given interaction (creation and destruction of particles) must be somehow managed at the same time, so that everything balances out in the end.

Fourier space is the **momentum space** representation complementary to _position space_: it is defined in terms of frequency and phase coordinates, where the frequency is proportional to the momentum of a particle, e.g., as captured in the [[Klein-Gordon]] equation. This Fourier representation is particularly convenient for keeping track of total energy, which is proportional to momentum and frequency, and allows one to more easily represent different types of particles splitting up some total amount of energy in different ways.

As captured in the Heisenberg [[uncertainty principle]], momentum and position are "conjugate" (complementary) variables, so when you represent everything in precisely in Fourier space, the position information is completely lost. 

* **Thus, "particles" in Fourier space have no positions: they spread across the entire space.**

Position information only arises in Fourier space in terms of the constructive and destructive interference effects of different phases and frequencies. Thus, representing a particle with some specificity of spatial position requires many different sine waves "working together" to add up in one part of space and cancel out in other parts of space.

Representing something with a fully precise spatial position requires an _infinite_ number of such sine waves, in the same way that representing a fully precise momentum (i.e, frequency) requires a continuum infinity of particle positions oscillating according to a specific precise frequency. This is again the Heisenberg uncertainty principle, and it is a basic property of [[wave]]s.

Once one understands these properties of Fourier space, it is clear that QFT is yet another [[tools vs models|calculational tool]] that makes it easier to compute the outcomes of particle accelerator experiments, but it does not represent the actual physical interactions among particles.

<!--- TODO: cloud chamber picture! -->

One of the main measurement devices in these accelerator experiments is the _cloud chamber_ (and more modern versions thereof), which show the tracks of the particles as they spew out of the accelerator. Inevitably, these show connected tracks where one particle flies out and then breaks apart into other tracks _at specific points in space_. There is effectively no way to represent this position-space nature of these processes within QFT. 

These particle tracks are instead consistent with the [[pilot-wave]] framework, where particles have a specific position at all times, even as they are exhibiting wave-like behavior that is well-captured by the QFT framework.

## Particle field space

The representation of particles in terms of simple [[harmonic oscillator]] variables at each point in space suggests a different type of field theory, where particles are excitations in this harmonic oscillator field. The point-like nature of particles (e.g., in the cloud chambers) suggests that these remain localized.

But clearly each point in space must have the capacity to represent any kind of particle. This is necessary to account for the fluidity of particle transmutations observed in high-energy accelerator experiments.

