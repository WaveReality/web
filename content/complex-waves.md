+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

<!--- TODO: WaveC, Schrodinger, WaveCDir, Weyl logic. -->
<!--- claude todo: try the idea of 2nd order complex wave?? -->

The _second-order_ [[wave]] equation is the prototypical version of a wave, which corresponds to macroscopic water waves and other such familiar phenomena. The core dynamic in such waves is the continuous bidirectional conversion of energy between wrinkles in _space_ (potential energy) and movement through _time_ (kinetic energy). However, standard quantum mechanics is based on [[Schrodinger]] waves, which obey a _first-order_ equation that uses the rotational properties of [[complex number]]s instead of the second-order integration of kinetic-energy to accomplish the core oscillatory property of waves.

The ability to capture oscillations through either a second-order kinetic system, or a first-order complex-number system, plays out in various ways throughout the larger space of quantum wave equations. Here, we develop a systematic understanding of this space, and the relative tradeoffs between these different wave functions.

Of primary interest is the extent to which the different wave equations tend to spread out over time (known as **dispersion**, which is also related to **diffraction**), because that relates to the [[epistemic]] vs _ontological_ aspect of what the wave is describing: as a wave propagates, the uncertainty in where it is going should generally increase, especially to the extent it exhibits [[stochastic motion]]. This is purely epistemic: our knowledge of it becomes less precise, but, under the [[pilot wave]] framework, a discrete particle should be localized precisely at all times, and its surrounding [[quantum wave]] should presumably _not_ spread out.

Thus, if we can find wave equations that differ in their tendency to spread out, this may have implications for the kinds of waves we might want to be using for a pilot-wave model. One interesting conclusion from this analysis is that the [[neutrino]] could be seen as a kind of necessary byproduct of a type of first-order electron waves that don't spread out as much as they do in the second-order wave function (described by the [[Weyl]] equation, which is what is necessary for the [[weak#electroweak]] framework). Furthermore, the Schrödinger wave function lies at the other extreme end of the spectrum in terms of its tendency to spread out, which would seem to have major implications for all of the abstract quantum math based on it.

Before proceeding, it is particularly useful to first read the [[harmonic oscillator]] page, which provides a simple and direct comparison between a second-order kinetic version of an oscillator, and a first-order complex-number version of the same oscillator dynamics. This harmonic oscillator has a single spatial value (i.e., within a single isolated [[cellular automaton]] cell), so its spatial dimensionality is much simpler than the wave, especially when it spreads over 3D space. That makes it easier to see the fundamental difference in the way the time dimension works, which is the key difference between the first and second-order wave equations.

The natural starting point with respect to dispersion is with a second-order wave equation configured instead to implement the **diffusion equation**. This is a very simple change in the equation, which also reveals the essential role of the second-order acceleration in driving the oscillatory behavior of the wave equation. You simply directly couple the spatial curvature (potential energy) directly to the _velocity_ (first order), without going through the acceleration (second order).

{id="eq_wave" title="wave equation"}
$$
\frac{\partial^2 \phi}{\partial t^2} = c^2 \frac{\partial^2 \phi}{\partial x^2}
$$

{id="eq_diffuse" title="diffusion equation"}
$$
\frac{\partial \phi}{\partial t} = c^2 \frac{\partial^2 \phi}{\partial x^2}
$$

When you run this equation (you can do it in the [[waves simulation]]) you see any disturbance melt away over time -- any "concentration" of wave "stuff" simply diffuses down into a uniform splat.

This is an interesting starting point, because the first-order Schrödinger equation is effectively the same as the diffusion equation, except it operates on complex state variables, and it sticks an $i$ in there to drive the rotation.

{id="eq_schrod" title="Schrödinger essentially"}
$$
\frac{\partial \chi}{\partial t} = -i \frac{\partial^2 \chi}{\partial x^2}
$$

where $\chi = \phi_a + i \phi_b$ is a complex-valued wave state with the two underlying values, which are just ordinary numbers. Although everyone somehow gets hung up on the $\phi_b$ number being "imaginary", it is just a regular number that happens to be located on an orthogonal plane to the "real" number, so that these two numbers can rotate around each other _in qualitatively the same way that velocity and the state position value do_ in the second-order version of the wave equation. 

Is the velocity in a second-order equation "imaginary"? No. It is just another degree of freedom -- a different value that you can use to make things oscillate. The connection between rotation and oscillation is basic trigonometry: the sine and cosine are waves that arise in any orthogonal basis coordinates as you rotate around a circle, and the complex plane just provides a 2D space for this rotation to occur.

To understand exactly what is happening in [[#eq_schrod], we can write it in terms of the two ordinary numbers. The key algebraic step is that you only keep the factors _without_ an $i$ for the $\phi_a$ factor, and those _with_ an $i$ for the $\phi_b$ factor. Furthermore, the convenient fact that $i^2 = -1$ makes the signs work out correctly for the rotation. The net result is that the $-i$ factor makes $\phi_a$ depend on $\phi_b$ and vice-versa -- it causes the two values to rotate into each other:

$$
\frac{\partial \phi_a}{\partial t} = - \frac{\partial^2 \phi_b}{\partial x^2}
$$

$$
\frac{\partial \phi_b}{\partial t} = \frac{\partial^2 \phi_a}{\partial x^2}
$$

From the perspective of the second-order wave equation ([[#eq_wave]]), the Schrödinger equation is a bit strange, because the spatial factor on the right-hand side is a _second-order_ spatial derivative (also written as $\nabla^2 \phi$ where $\nabla^2$ is the Laplacian).

The critical advantage of this second-order spatial derivative is that _it doesn't care about direction:_ a spatial disturbance in any direction surrounding a given point in the wave state will give rise to a change in that wave state, proportional to the magnitude of the disturbance. This means that this equation supports waves traveling in any direction, just like the standard second-order wave equation, because that directional property depends entirely on the nature of the spatial derivative factor, which is shared with the second-order wave equation.

However, the coupling of this second-order spatial factor directly to the first-order temporal derivative in the Schrödinger equation results in an extremely high level of dispersion

<!--- todo: directional wave, then helical. -->
