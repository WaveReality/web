+++
bibfile = "mechphys.json"
+++

Compared to the [[Hilbert space]] approach, the explicit use of the **Schrödinger wave equation** represents an increased level of commitment to the details involved in the dynamics of the wave updating, its frequency and phase characteristics, and how it spreads out over time. Schrödinger's wave equation captures basic non-relativistic Newtonian physics in a simple linear, first-order framework, and can be derived from a [[Hamiltonian]] representing the total energy of the system, which is strictly conserved over time. It captures the fundamental relationships between momentum and wave frequency at the heart of quantum physics, as discussed in [[Klein-Gordon]].

However, it has a rather simplistic treatment for how forces affect charged particles in terms of overall scalar potentials, and says nothing in detail about how electric charge generates the EM wave field (or photons for that matter), or the detailed way in which different particles might interact with each other. Indeed, because the Schrödinger wave equation is linear, it is incapable of capturing particle interactions, because the waves simply superpose (additively combine) past each other, without impacting each other at all.

Thus, in order to capture relevant interactions, the Schrödinger wave equation is applied to a multi-dimensional  [[configuration space]] representation that is essentially equivalent to the state space representation in matrix mechanics. For example, if there are two interacting particles, then they each get their own set of 3D dimensional coordinates within this configuration space, and the entire wave function evolves over time so as to conserve the overall energy / probability represented in the configuration space.

As noted above, this configuration space is entirely [[non-locality|non-local]] by its very construction, representing at each instant of time the entire configuration of the system, regardless of how far apart any of the particles might be. Interestingly, exactly such a configuration space model is used in _classical_ applications of the Hamiltonian framework, and yet somehow its use there is widely recognized as just being a calculational tool.

In summary, the high-dimensional non-local configuration space is very different from anything anyone would recognize as actual 3D physical space. Nevertheless, one of the most striking and challenging results from these standard QM models is that the non-local effects that they predict actually do appear to be empirically validated. Thus, a significant challenge remains to understand the underlying physical nature of these effects, and how they can occur without violating everything else we have come to regard as strict physical laws, specifically the speed-of-light constraints of special relativity ([[@DurrGoldsteinNorsenEtAl14]]).

Also, while the dimensionality of configuration space increases linearly in the number of particles involved, the underlying computational complexity of the space grows exponentially, and quickly becomes computationally intractable for even relatively moderately-sized such spaces. This is precisely what makes quantum computers so attractive. Nevertheless, it remains unclear how Nature might get around such prohibitive exponential scaling problems, in whatever computation it is performing.

## Schrodinger's equation

Using the total energy ([[Hamiltonian]]) approach, we can derive Schrödinger's equation, using the same energy and momentum operators that we used in the derivation of the [[Klein-Gordon]] equation (strongly recommend reading that page first, for the introductory treatment of this approach). To remind, these operators are:

{id="eq_momentum" title="momentum operator"}
$$
\hat{p} = -i \hbar \vec{\nabla}
$$

{id="eq_energy" title="energy operator"}
$$
\hat{E} = i \hbar \frac{\partial }{\partial t}
$$

{id="eq_gradient" title="gradient operator"}
$$
\vec{\nabla} = \left(\frac{\partial {}}{\partial {x}}, \frac{\partial {}}{\partial {y}}, \frac{\partial}{\partial {z}}\right)
$$

Next, we need to define the total energy Hamiltonian. Instead of the relativistic total energy, we use the classical Newtonian expression for the kinetic energy of a particle, in terms of its velocity $\vec{v}$, just as we did in the simple wave energy calculation in [[wave]]s:

{id="eq_kinetic" title="kinetic energy of particle"}
$$
K = \frac{1}{2} m_0 \vec{v}^2 = \frac{1}{2 m_0} \vec{p}^2
$$

The second form uses the Newtonian relationship of momentum to velocity (just $\vec{p} = m_0 \vec{v}$) --- because we have a momentum operator, we need to use this momentum form.

We also include a potential energy term that is a function of any kind of electrical or other force potential that the particle experiences. We won't deal much with such forces at this point, so we just call this potential energy $V$ for now, and focus on the kinetic energy. The total energy or Hamiltonian in abstract terms is just the kinetic energy $K$ plus this potential energy:

{id="eq_kv" title="kinetic and potential energy"}
$$
E = K + V
$$

$$
E = \frac{1}{2 m_0} \vec{p}^2 + V
$$

We can now just apply our momentum and energy operators to these expressions, and the result is in fact:

{id="eq_schrodinger" title="Schrödinger's equation"}
$$
i \hbar \frac{\partial {\phi}}{\partial t} = -\frac{\hbar^2}{2 m_0} \nabla^2 \phi + V \phi
$$

The net result is that we can conclude that Schrödinger's equation provides an accurate description of the flow of energy and momentum over time of a "particle" described by a wave, such that it obeys classical Newtonian physical laws. Note that in comparison with the KG equation, there is no speed-of-light factor $c$ in this equation, consistent with its non-relativistic nature.

Omitting various constants (factors of $h$) and any external force potential, Schrödinger's equation is:

{id="eq_schrodinger" title="Schrödinger's equation, essence"}
$$
i \frac{\partial {\phi}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \phi
$$

where $m_0$ is again the rest mass of the particle in question. This is clearly very similar to the basic second-order KG wave equation:

{id="eq_KG" title="Klein-Gordon equation"}
$$
\frac{\partial^2 {\varphi}}{\partial t^2} = c^2 \nabla^2 \varphi - \frac{m_0^2}{\hbar^2} \varphi
$$

except that the temporal derivative is first-order, and mass enters in a different way. Nevertheless, the driving force is still the overall curvature of the wave, computed by $\nabla^2 \varphi$. As we noted above, the multiplication by the $i$ term causes things to rotate --- this rotation is key for making the first-order equation behave like a wave.

To see this effect more explicitly, we can write out Schrödinger's equation in terms of the two underlying scalar values:

$$
i \frac{\partial ({\varphi_a + i \varphi_b})}{\partial t} = - \frac{1}{2m_0} \nabla^2 (\varphi_a + i \varphi_b)
$$

$$
-\frac{\partial {\varphi_b}}{\partial t} + \frac{\partial {i \varphi_a}}{\partial t} = -\frac{1}{2m_0} \nabla^2 \varphi_a - i \nabla^2 \varphi_b
$$

where $\varphi_a$ indicates a scalar state variable that is the $a$ component of $\phi$, and $\varphi_b$ is the $b$ component of $\phi$. Note that the derivatives operate separately on each of the two variables. At this point, we now can just separate all the terms that involve an $i$ from those that do not, to get update equations for each of the two variables. For the real-valued components (without the $i$):

$$
-\frac{\partial {\varphi_b}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \varphi_a
$$

$$
\frac{\partial {\varphi_b}}{\partial t} = \frac{1}{2m_0} \nabla^2 \varphi_a
$$

and for the imaginary components (dropping the $i$ now, because we no longer need it to keep the variables separated):

$$
\frac{\partial {\varphi_a}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \varphi_b
$$

In a discrete-space and time [[cellular automaton]] implementation, these equations would be written:

$$
\dot {\varphi_a}_i^{t+1} = - \frac{3}{26 m_0} \sum_{j \in N_{26}} k_j ({\varphi_b}_j^t - {\varphi_b}_i^t)
$$

$$
{\varphi_a}_i^{t+1} = {\varphi_a}_i^t + \dot {\varphi_a}_i^{t+1}
$$

and:

$$
\dot {\varphi_b}_i^{t+1} = \frac{3}{26 m_0}\sum_{j \in N_{26}} k_j ({\varphi_a}_j^t - {\varphi_a}_i^t)
$$

$$
{\varphi_b}_i^{t+1} = {\varphi_b}_i^t + \dot {\varphi_b}_i^{t+1}
$$

So, in the end, Schrödinger's equation really just boils down to two very simple differential equations. Interestingly, these equations are *coupled*, in the sense that it is the curvature of $\varphi_a$ that drives the change in $\varphi_b$, and vice-versa. This is the rotational aspect of the equation mentioned earlier, which is caused by the presence of the $i$ in the equation.

When you actually implement Schrödinger's equation on a computer using the update rules given above, the resulting system is numerically unstable. In other words, the resulting numbers quickly blow up to infinity. This is not due to any kind of numerical roundoff error from limited precision floating point numbers on the computer, but rather due to the way that changes in state values reverberate back and forth across the two scalar values. However, it is possible to overcome it relatively simply by just alternating the update: on one time step you compute one value, and on the next you update the other. This is what is done for illustrative purposes in the computer explorations.

The basic phenomenology of Schrödinger's equation is that wave packets propagate through space, with a speed that is proportional to $\nabla^2 \phi$, which in turn is proportional to the frequency of the wave. In other words, it describes exactly the same behavior as the KG equation, where particle speed is proportional to frequency.

One critical property of Schrödinger's equation (which the scalar [[Klein-Gordon]] equation does not have) is that it preserves the overall magnitude of the $\phi$ state values across all of space, for all time. This is to say, if you compute the sum of $\phi \phi^*$ for each point in space, this sum will remain the same across time under the Schrödinger equation. This conserved value is interpreted as a probability in standard quantum mechanics.

For example, we can initialize the state with a localized wave packet (see [[matter waves]]) to represent the initial probability for the location and velocity of a particle (velocity being a function of the frequency of the wave packet). If we then apply the Schrödinger equation repeatedly, we can interpret the resulting $\phi \phi^*$ values as the probability of the particle having moved to the corresponding location.

In other words, the wave packet defines a kind of "cloud of probability" for finding a discrete particle within its midst. However, these probabilities have different meanings in different scenarios, and it is notoriously difficult to come up with a intuitively sensible interpretation of what these probability clouds mean (see [[Copenhagen]] for discussion).

## Explorations

See [[Schrodinger 1D Simulation]] and [[Schrdodinger 3D Simulation]].

