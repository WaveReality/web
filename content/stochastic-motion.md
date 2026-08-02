+++
Categories = ["Spinfield Model"]
bibfile = "mechphys.json"
+++

The [[pilot-wave]] model posits a discrete, localized particle whose motion is influenced by a distributed, physically real [[quantum wave]]. In the [[cellular automata]] (CA) framework, a discrete localized particle lives entirely within one CA cell.

An obvious problem with this notion of something being contained entirely within a discrete cell is that it becomes challenging to imagine how it might ever move to another such cell. Such a move would have to happen in a discrete jump, creating a major discontinuity in the overall wave state, and potentially making the particle trajectory seemingly discontinuous and anisotrophic. 

The only way to overcome those difficulties is to use **stochastic** discrete jumps, such that, on longer time averages, the timing and spatial distribution of such jumps smooths out into a continuous, isotrophic distribution. Thus, contrary to Einstein's oft-cited objection that "God does not play dice with the universe", in fact it seems that an essential form of randomness is _necessary_ for discrete particles to move in a physically plausible manner.

Furthermore, although the continuum limit is mathematically approachable through the tools of calculus, it is problematic from a physics perspective due to the nearly-infinite field strengths (and thus energies) that would be present in the immediate vicinity of a charged particle. This _ultraviolet catastrophe_ is a recurring theme throughout the [[history]] of quantum physics, and it is nicely resolved through the use of the discrete cubic lattice of the CA framework.

## Brownian motion and the Schrodinger equation

The seminal work in analyzing stochastic discrete particle motion within the context of quantum physics was done by [[@^Nelson66]], who showed that a form of stochastic discrete particle motion actually results in the [[Schrodinger]] wave function in the continuous time-average limit. Interestingly, this work builds on the original work by Einstein on Brownian random-walk motion, back in 1905. Subsequent work has developed these ideas in multiple ways ([[@Cufaro-PetroniVigier83]]; [[@CufaroPetroniVigier79]]; [[@Ord96]]; [[@^Sciarretta18]]; [[@^Sciarretta21]] and others reviewed therein).

Critically none of this existing work involves an integrated wave-particle [[duality]]; it focuses exclusively on the time-average distributions of discrete particle motion. Thus, unlike the [[pilot-wave]] framework, the quantum wave function in a purely stochastic particle model is entirely [[epistemic vs ontic|epistemic]]: it just describes the expected value of a discrete particle's random walk trajectories over time. There is no physical reality to such a wave. By contrast, our goal here is to derive an integration of discrete particle motion with quantum wave functions.

{id="figure_pf-origin" style="height:20em"}
![Stochastic origin of quantum momentum / frequency relationship. The momentum on the left is 0.5c while on the right is 0. The distribution of position is on the vertical axis, while time is on the horizontal axis, with each point centered at the origin in the center (i.e., the temporal autocorrelation function). The variance on the left is half of that on the right.](media/fig_asmom5_0_autoc.png)

One key intuition for why discrete particle motion naturally exhibits quantum wave-like behavior is that a slow drift rate produces a wide cloud of space where particle could be, corresponding to a long wavelength in the probability cloud that the Schrödinger wave function describes. However, when the particle has high momentum (velocity), it moves more deterministically in a given direction, resulting in a narrower range of variance around the particle's mean trajectory, resulting in a narrower effective wavelength ([[#figure_pf-origin]]).

While this pure particle-based approach is appealing in its simplicity, it does not appear to provide an explanation for phenomena such as the [[double-slit]] experiment, where somehow a particle can interfere _with itself_, but only if the other slit is open. Furthermore, it cannot be the case that these interference effects only arise in the rare cases when a discrete particle happens to wander so aimlessly as to go through both slits somehow.

Furthermore, discrete particles require some other continuous state values embedded in the same discrete lattice to drive the _probabilities_ underlying their stochastic behavior. The nature of such values is generally unaddressed in existing frameworks. 

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

TODO: It would be useful to derive a relativistic version of these individual motion equations, where the effective mass, and thus momentum, increases with velocity. This is what the wave function naturally does.

TODO: [[Weyl]] and [[Dirac]] wave couples spin with direction as a helical thing.

## Other stochastic motion frameworks

### Heat bath models

There is a literature on coupling of a stochastic particle with a "heat bath", somewhat like the [[zero-point]] field, and trying to understand the aggregate behavior of such a system. [[@^DunkelHanggi05a]], [[@DunkelHanggi05]] provide a relatively accessible treatment, building on foundational work ([[@Dudley65]], [[@Dudley73]], [[@GuerraRuggiero78]], [[@Nakagomi88]]). This all builds on Langevin equations, which are stochastic equations of motion, with connections to Ornstein-Ullenbach and Fokker-Planck etc. The specific restriction to heat bath dynamics vs. some kind of other intrinsic stochastic process is perhaps overly restrictive, but they nevertheless have a four-vector representation that seems to involve a conservation of energy between the time and momentum factors, which is really the essential calculus for the SHO model.

One key "no-go" finding from [[@^Dudley65]] is that a purely Markovian position-based system doesn't capture particle motion -- you _need_ an additional momentum / velocity vector as part of the state. This is definitely key.

* wave function trades energy against momentum -- momentum is $\gamma m0 \nu$, $E^2 = p^2 + m0^2$ so the 1 in above eq is like m0^2 -- not clear where the 1/2 comes from but whatever.
* key idea that m0 is the internal motion of the particle rotating through spin, so need to just have that always going on as an "anchor", and then there are these extra $\nu_{\mu}$ factors where the x^2 + y^2 + z^2 hypotenuse of the momentum-velocity is v^2 relative to c^2 -- i.e., need to constrain total length to c^2.

* probabilistically when it stays still it rotates the internal spin. if it never stays still it never rotates the spin. the nutrino very rarely rotates the spin, but does sometimes. spin coupling couples the two helicies of the Weyl. need to go back to that. are there 2 neutrinos trapped inside one electron??

### Quantum cellular automata and random walks

There is a growing literature on quantum cellular automata (QCA) and the quantum random walk (QW) model, going back to the influential papers by [[@^Feynman82]], [[@^AharonovDavidovichZagury93]], and [[@^Bialynicki-Birula94]]. This approach has become somewhat more active recently due to the practical application of these models for designing quantum algorithms ([[@ChildsCleveDeottoEtAl03]]; [[@ChiribellaDArianoPerinotti11]]; [[@DAriano17]]; [[@Kempe09]]). Although this work shares many essential properties in common with the present approach, it is not directly applicable, due to some differences in basic assumptions.

The central idea behind the random walk model proposed originally by [[@^AharonovDavidovichZagury93]] is that a shift operator that translates the quantum state in either the left or right direction along a discrete 1D lattice can be driven by a quantum measurement process on a _different_ element of the quantum state (e.g., the spin), which thus would produce a random sample according to the usual probabilistic quantum theory. This "internal" source of randomness will thus produce an emergent random walk dynamic as the state evolves (translates) over time, in the same way that thermal noise produces Brownian random walk motion as originally analyzed by Einstein.

The results from these analyses show that standard wave functions such as the [[Dirac]], [[Maxwell]] and [[Weyl]] functions emerge from these random walk processes. A closely related analysis in 3D by [[@^Bialynicki-Birula94]] showed the emergence of the Weyl and Maxwell wave functions in the context of a cellular automaton update rule with _unitary_ update rules.

The basic intuition is that these unitary update rules cause the propagation of a wave-like pattern at the speed of light (one unit cell per unit time), and if multiple internal cell states are present, this propagation can also include the [[spin]] property where the state rotates through these internal states as it also propagates. This is the essential feature of the [[Weyl]] wave functions, which describe a "pure spin" particle such as a massless [[neutrino]] that travels at the speed of light while spinning in one fixed helical rotation.

Critically, these approaches do _not_ provide a model of a discrete particle moving with graded momentum (velocities) in an isotrophic manner along a cubic grid, which is what we develop here, building on the approach originated by [[@Nelson66]] in analyzing single-particle Brownian motion. We use equations of motion initially developed by [[@^Sciarretta18]].

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


