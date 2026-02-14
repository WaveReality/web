+++
bibfile = "mechphys.json"
+++

{id="figure_packet" style="height:20em"}
![A wave packet, which is a spatially localized wave disturbance that propagates through space as a coherent entity. The two values shown here are the state value and the first derivative --- the relationship between these two determines which direction the wave travels. This is our model for a particle. Mathematically, it can be constructed by multiplying a Gaussian function (normal bell-shaped distribution curve) times a sine wave.](media/fig_wave_packet_raw.png)

{id="figure_mass" style="height:20em"}
![The additional mass term $-m_0^2 \varphi$ in the Klein-Gordon (KG) wave equation "drags down" the wave in proportion to the height of the waves (i.e., amplitude away from zero, either positive or negative). This fights against the curvature of the wave, computed by $\nabla^2$. Higher frequency waves have higher curvature, and thus move faster than lower frequency waves.](media/fig_kg_mass_drag.png)

In beginning to explore the wave model of matter, we need to first establish a few basic concepts of what it would even mean for a particle to be described by a wave. The main idea is that a particle corresponds to a _wave packet_, which is a spatially localized wave disturbance ([[#figure_packet]]). It can act like a particle in that it is somewhat spatially localized, and moves as a coherent entity. If you zoomed out very far, and blurred your eyes, you could imagine that a wave packet would look like a tiny point particle.

Nevertheless, it fundamentally acts like a wave, in the sense that it is actually made of oscillations, and obeys a wave equation. As to what exactly the wave material is and what it means for observations, we'll postpone for later. At this point, we'll content ourselves with this level of description, and just start developing some new twists on the basic wave equation; we'll return to the thorny interpretational issues (e.g., wave-particle duality, probabilistic interpretation of the wave, etc), once we have a better sense of how these new waves behave.

Recall that the wave equation can be written as a second-order differential equation:

{id="eq_wave" title="standard wave equation"}
$$
\frac{\partial^2 {\varphi}}{\partial t^2} = c^2 \nabla^2 \varphi
$$

What if we add a single new term to this equation, where we subtract away some *mass* ($m_0$, a constant) from the Laplacian ($\nabla^2 \varphi$) curvature driving force term ([[#figure_mass]]):

{id="eq_mass" title="mass subtraction"}
$$
\frac{\partial^2 {\varphi}}{\partial t^2} = c^2 \left( \nabla^2 \varphi - \frac{m_0^2}{\hbar^2} \varphi \right)
$$

or, in somewhat simpler notation that we'll use more frequently:

{id="eq_kg" title="Klein-Gordon equation"}
$$
\frac{\partial^2 {\varphi}}{\partial t^2} {{=}} c^2 \left(\nabla^2 - \frac{m_0^2}{\hbar^2} \right) \varphi
$$

where "hbar" $\hbar = \frac{h}{2\pi}$ and $h$ is Planck's constant. As we elaborate below, the most natural interpretation of Planck's constant here is as a simple scaling term on the impact of mass on the matter wave dynamics. For this reason, we strongly argue that anything having to do with this constant reflects an interaction with matter waves, and for example it makes no sense to use *h* in characterizing the behavior of light waves, because they have no mass. There is no evidence of any such constant in [[Maxwell]]'s EM wave equations.

Thus, Einstein's creation of the photon with energy $E = h \nu$ is really just a calculational tool for representing the interaction between EM waves and matter in atomic systems, and it is the matter waves that impart the *h* constant, not the EM "photon".

This new equation ([[#eq_kg]]) is called the **Klein-Gordon (KG)** equation, named after Oskar Klein and Walter Gordon, who published the first papers on it ([[@Klein26]]; [[@Gordon27]]; see [[@Kragh84]] for a detailed history of this equation, which was actually discovered by many individuals, including Schrodinger). This equation captures a surprising number of important phenomena, as we detail next.

First, we'll introduce some variations on how to write this equation, which are all obviously identical to the KG equation given above, but highlight different features of it, as we'll see more later. Here's one such variation:

$$
\frac{\partial^2 {\varphi}}{\partial t^2} - c^2 \nabla^2\varphi = -\frac{c^2 m_0^2}{\hbar^2} \varphi
$$

and another:

$$
\left(\frac{\partial^2 {}}{\partial t^2} - c^2 \nabla^2 + \frac{c^2 m_0^2}{\hbar^2}\right) \varphi = 0
$$

These last two forms are useful for relating to the four-vector version of the wave equation, where we saw that:

- $\partial\_\mu \partial^\mu = \frac{\partial^2}{\partial t^2} - c^2 \nabla^2 $

so that the equation can be written:

- $\partial\_\mu \partial^\mu \varphi = - \frac{c^2 m_0^2}{\hbar^2} \varphi $

or:

- $\left(\partial\_\mu \partial^\mu + \frac{c^2 m_0^2}{\hbar^2}\right) \varphi = 0 $

To actually implement this KG equation in our cellular automaton model, we make one modification to the acceleration term, to subtract off the mass:

{id="eq_disc" title="discrete Klein-Gordon equation"}
$$
\ddot \varphi_i^{t+1} = c^2 \frac{3}{13}\sum\_{j \in N\_{26}} k_j (\varphi_j - \varphi_i) - \frac{c^2 m_0^2}{\hbar^2} \varphi_i
$$

## Variable Speeds: Momentum from Frequency

{id="figure_frequency" style="height:20em"}
![Relationship between frequency and speed in the Klein-Gordon (KG) wave function, which derives from competition between the "mass drag" and the overall curvature of the wave. Higher frequency waves have more curvature and thus move faster.](media/fig_kg_freq_speed.png)

One of the most important features of this KG equation is that waves now travel at _variable speeds_, instead of always moving at exactly the same speed (the speed of light). This speed now depends on the relationship between the curvature ($\nabla^2 \varphi$) and the squared-mass value $\frac{m_0^2}{\hbar^2} \varphi$. In essence, the mass "drags down" the wave propagation force conveyed by the local curvature, $\nabla^2 \varphi$. Therefore, to get the wave to move faster, you need more curvature, which is to say, a higher frequency wave, because higher frequency waves have more waves per unit length, and this means overall greater "curvature" ([[#figure_frequency]]).

This relationship between frequency $f$ of a wave and the momentum (velocity \* mass) of the particle that it describes is captured in one of the most basic equations of quantum physics:

- $p = \frac{h}{c} f $

where $p$ is the momentum, and $h$ is Planck's constant. This can also be written in terms of the wavelength $\lambda$, which is the inverse of the frequency:

- $f = \frac{c}{\lambda} $
- $\lambda = \frac{c}{f} $
- $\lambda f = c $

so that momentum is inversely proportional to the length of the waves:

- $p = \frac{h}{\lambda} $

Although it might be tempting to compute the velocity from the momentum expression given above (e.g., $p = m v$ so $v = \frac{p}{m}$), this is not accurate due to the effects of special relativity as we discuss in greater detail below. Instead, the appropriate equation that relates the momentum and the velocity is:

{id="eq_" title="relativistic momentum-velocity relationship"}
$$
p = \gamma m_0 v
$$

where $\gamma$ is the Lorentz factor as described in [[special relativity]]. When we make all the necessary substitutions and do a bit of algebra, we end up with this expression for the velocity of the "particle" as a function of wavelength, rest mass, and the relevant constants:

{id="eq_" title="velocity as function of wavelength, relativistically correct"}
$$
v = \frac{h c}{\sqrt{c^2 m_0^2 \lambda^2 + h^2}}
$$

See [[#velocity_derivation]] for all the algebraic steps in this derivation. In the exploration below we'll confirm this equation experimentally. The complexity of this equation for velocity versus the much simpler expressions for momentum indicate why momentum and not velocity is the natural quantity to deal with in relativistic wave functions.

Also, as we will explore in greater detail later, the momentum can be computed directly from the wave function in terms of the first-order spatial derivative or gradient:

{id="eq_" title="essence of momentum operator"}
$$
\vec{p} \propto \vec{\nabla} \varphi
$$

where this spatial gradient is again going to be greater overall as the wave frequency increases, as suggested by the quantum mechanical relationships above.

Again, the main point for now is just that introducing the mass term makes the relative curvature or frequency of the wave matter in determining the overall velocity or momentum of the wave packet that describes a particle. Without this mass term, all waves travel at the speed of light.

## Newtonian Mechanics: F = ma

We can intuitively see that this very simple modification to the wave equation captures all of classical (Newtonian) mechanics for a "particle" characterized by a wave according to this equation. In the absence of any external forces, the wave will propagate along at a constant velocity (Newton's first law of inertia), because the frequency of the wave does not decrease (and it would not change its overall direction of propagation). If a force is applied to this system, it will change the frequency of the oscillation of the wave, and thus result in a change in momentum, in accord with Newton's second law:

- $F = \frac{\partial \vec{p}}{\partial t} = m \frac{\partial \vec{v}}{\partial t} = m \vec{a} $

After a few more developments, we can make this relationship much more formally accurate and precise, by considering the overall energy and momentum relationships computed by the KG wave equation. We will see that Schrodinger's equation, which is the primary wave equation for basic quantum physics, captures classical Newtonian physics, and that the KG equation is a version of Schrodinger's equation that also takes into account special relativity, which is important when particles are moving very fast (i.e., relatively close to the speed of light).

Anticipating these results, and relying on intuition for now, we see that with one tiny addition, we now have an equation that can describe the motion of a massive particle through space (e.g., as a wave packet), in agreement with all the known physical laws (i.e., quantum physics and special relativity, which reduce in certain cases to the more familiar Newtonian mechanics).

## What is the Mass?

The value $m_0$ in the KG equation is the **rest mass** of the particle that it describes (the 0 subscript indicates "rest"). It is a fixed, constant value for a given type of particle, and thus we can easily imagine that it is just built into the equations for the state variables that characterize that type of particle.

For an electron, this mass is $9.1x10^{-31}$ kg, which is an extremely tiny weight. In the typical units that we use in our numerical models, the rest mass of the electron has the value of 3.0.

The value of $m_0$ also defines the **Compton wavelength** of a given particle, which is the wavelength of a particle at rest, due strictly to the rest mass. The formula for the Compton wavelength $\lambda_C$ is:

- $\lambda_C = \frac{h}{m_0 c} $

One other thing about the rest mass --- we can consider what happens when you set the rest mass of our particle to zero. You should see that we immediately recover the basic wave equation from before. Thus, it is immediately obvious that "particles" with a zero rest mass must move at the speed of light. This is a basic postulate of special relativity. Note also that it is impossible for a massive particle to travel at the speed of light, because it would have to have an infinitely high frequency, and this is not possible (even in a continuous spatial model).

## Smooth, Continuous Motion

As a continuous-valued second-order wave equation, the KG equation still has all of the advantages discussed earlier about making the underlying grid disappear. Furthermore, it easily and naturally allows us to describe continuously variable rates of speed, in terms of continuous variation in the frequency or wavelength of wave oscillation. Particles described by such an equation can also travel in any direction, because of the essentially perfect spatial symmetry of wave propagation exhibited by this equation.

In contrast, models with a discrete particle suffer from all manner of complexity in overcoming such problems. Thus, by describing a particle exclusively using this wave equation, we avoid many difficulties, and it remains to be seen if we encounter any other problems.

## Exploration of Klein-Gordon Waves

Now we explore the above properties of the Klein-Gordon wave equation, to get a concrete sense of how it works, before turning to a more extended discussion of how the KG wave equation explains the stretchy properties of matter implied by the Lorentz transformation and special relativity.

## Initial Conditions

Creating a moving wave packet that moves with a given velocity is a bit more complicated than for the simple wave equation, because we have more constraints to take into account.

## Schrodinger's Equation vs Klein-Gordon

The Klein-Gordon equation that we've been exploring is typically introduced as a strange and problematic alternative to the Schrodinger wave equation, which provides the cornerstone of standard quantum physics. As we saw, the KG equation is derived from the relativistic total energy, whereas we'll see here that Schrodinger's equation can be derived from *Newtonian* total energy, and thus is clearly not accurate for anything moving very fast (a significant fraction of the speed of light), or where new particles are created out of raw energy. The fact that it holds such strong sway in the field can be attributed to its strict conservation properties --- as normally interpreted, the Schrodinger equation conserves the total probability value, as it propagates through space. In contrast, the KG equation does not have such a strict conservation behavior.

Furthermore, the Schrodinger equation is a first-order wave equation, which has many advantages from an analytical perspective, even as it makes it very difficult for many people to understand, due to its reliance on [[complex numbers]]. In general, wave-like behavior can either be described by a second-order equation involving normal scalar variables (as we've been doing), or it can be described by a first-order equation involving complex numbers, exemplified by the Schrodinger equation.

In the first-order version, you have two variables for every one variable in the second-order one --- we'll see later that this fact allows us to use only four variables to represent an electron using a second-order wave equation, whereas the standard first-order Dirac equation requires eight. The general intuition is that a first-order wave equation involves motion as rotation among its complex variables, in addition to motion through space, whereas the second-order equation just has motion through space. This will be clearer as we examine Schrodinger's equation more closely.

Before we do so, you should review [[complex numbers]] if you are not completely familiar with them. Complex numbers are important for understanding [[Schrodinger]]'s equation (as described thereafter), and for subsequent developments of the KG equation, which requires complex numbers to represent a conserved charge value.

Using the total energy ([[Hamiltonian]]) approach, we can derive Schrodinger's equation, using the very same energy and momentum operators that we used in the derivation of the KG equation above. To remind, these operators are:

- **momentum operator:** $\hat{p} = -i \hbar \vec{\nabla} $
- **energy operator:** $\hat{E} = i \hbar \frac{\partial }{\partial t} $
- gradient operator: $\vec{\nabla} = \left(\frac{\partial {}}{\partial {x}}, \frac{\partial {}}{\partial {y}}, \frac{\partial}{\partial {z}}\right)$

Next, we need to define the total energy Hamiltonian. Instead of the relativistic total energy, we use the classical Newtonian expression for the kinetic energy of a particle, in terms of its velocity $\vec{v}$, just as we did in the simple wave energy calculation in [[waves]]:

- $K = \frac{1}{2} m_0 \vec{v}^2 = \frac{1}{2 m_0} \vec{p}^2$

The second form uses the Newtonian relationship of momentum to velocity (just $\vec{p} = m_0 \vec{v}$) --- because we have a momentum operator, we need to use this momentum form.

We also include a potential energy term that is a function of any kind of electrical or other force potential that the particle experiences. We won't deal much with such forces at this point, so we just call this potential energy $V$ for now, and focus on the kinetic energy. The total energy or Hamiltonian in abstract terms is just the kinetic energy $K$ plus this potential energy:

- $E = K + V$
- $E = \frac{1}{2 m_0} \vec{p}^2 + V $

We can now just apply our momentum and energy operators to these expressions, and the result is in fact:

- **Schrodinger's equation:** $i \hbar \frac{\partial {\phi}}{\partial t} = -\frac{\hbar^2}{2 m_0} \nabla^2 \phi + V \phi $

The net result is that we can conclude that Schrodinger's equation provides an accurate description of the flow of energy and momentum over time of a "particle" described by a wave, such that it obeys classical Newtonian physical laws.

Omitting various constants (factors of $h$) and any external force potential, Schrodinger's equation is:

- $i \frac{\partial {\phi}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \phi $

where $m_0$ is again the rest mass of the particle in question. This is clearly very similar to the basic second-order KG wave equation:

- $\frac{\partial^2 {\varphi}}{\partial t^2} = c^2 \nabla^2 \varphi - \frac{m_0^2}{\hbar^2} \varphi $

except that the temporal derivative is first-order, and mass enters in a different way. Nevertheless, the driving force is still the overall curvature of the wave, computed by $\nabla^2 \varphi$. As we noted above, the multiplication by the $i$ term causes things to rotate --- this rotation is key for making the first-order equation behave like a wave.

To see this effect more explicitly, we can write out Schrodinger's equation in terms of the two underlying scalar values:

- $i \frac{\partial {\varphi_a + i \varphi_b}}{\partial t} = - \frac{1}{2m_0} \nabla^2 (\varphi_a + i \varphi_b) $
- $-\frac{\partial {\varphi_b}}{\partial t} + \frac{\partial {i \varphi_a}}{\partial t} = -\frac{1}{2m_0} \nabla^2 \varphi_a - i \nabla^2 \varphi_b $

where $\varphi_a$ indicates a scalar state variable that is the $a$ component of $\phi$, and $\varphi_b$ is the $b$ component of $\phi$. Note that the derivatives operate separately on each of the two variables. At this point, we now can just separate all the terms that involve an $i$ from those that do not, to get update equations for each of the two variables. For the real-valued components (without the $i$):

- $-\frac{\partial {\varphi_b}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \varphi_a $
- $\frac{\partial {\varphi_b}}{\partial t} = \frac{1}{2m_0} \nabla^2 \varphi_a $

and for the imaginary components (dropping the $i$ now, because we no longer need it to keep the variables separated):

- $\frac{\partial {\varphi_a}}{\partial t} = - \frac{1}{2m_0} \nabla^2 \varphi_b $

In a discrete-space and time CA-like implementation, these equations would be written:

- $\dot {\varphi_a}\_i^{t+1} = - \frac{3}{26 m_0} \sum\_{j \in N\_{26}} k_j ({\varphi_b}\_j^t - {\varphi_b}\_i^t) $
- ${\varphi_a}\_i^{t+1} = {\varphi_a}\_i^t + \dot {\varphi_a}\_i^{t+1} $

and:

- $\dot {\varphi_b}\_i^{t+1} = \frac{3}{26 m_0}\sum\_{j \in N\_{26}} k_j ({\varphi_a}\_j^t - {\varphi_a}\_i^t) $
- ${\varphi_b}\_i^{t+1} = {\varphi_b}\_i^t + \dot {\varphi_b}\_i^{t+1} $

So, in the end, Schrodinger's equation really just boils down to two very simple differential equations. Interestingly, these equations are *coupled*, in the sense that it is the curvature of $\varphi_a$ that drives the change in $\varphi_b$, and vice-versa. This is the rotational aspect of the equation mentioned earlier, which is caused by the presence of the $i$ in the equation.

When you actually implement Schrodinger's equation on a computer using the update rules given above, the resulting system is numerically unstable. In other words, the resulting numbers quickly blow up to infinity. This is not due to any kind of numerical roundoff error from limited precision floating point numbers on the computer, but rather due to the way that changes in state values reverberate back and forth across the two scalar values. However, it is possible to overcome it relatively simply by just alternating the update: on one time step you compute one value, and on the next you update the other. This is what is done for illustrative purposes in the computer explorations.

The basic phenomenology of Schrodinger's equation is that wave packets propagate through space, with a speed that is proportional to $\nabla^2 \phi$, which in turn is proportional to the frequency of the wave. In other words, it describes exactly the same behavior as the KG equation, where particle speed is proportional to frequency.

One critical property of Schrodinger's equation (which our current scalar KG equation does not have) is that it preserves the overall magnitude of the $\phi$ state values across all of space, for all time. This is to say, if you compute the sum of $\phi \phi^\*$ for each point in space, this sum will remain the same across time under the Schrodinger equation. This conserved value is interpreted as a probability in standard quantum mechanics. For example, we can initialize the state with a localized wave packet (Figure 4.1) to represent the initial probability for the location and velocity of a particle (velocity being a function of the frequency of the wave packet). If we then apply the Schrodinger equation repeatedly, we can interpret the resulting $\phi \phi^\*$ values as the probability of the particle having moved to the corresponding location.

In other words, the wave packet defines a kind of "cloud of probability" for finding a discrete particle within its midst. However, these probabilities have different meanings in different scenarios, and it is notoriously difficult to come up with a intuitively sensible interpretation of what these probability clouds mean. We return to these issues later.

<!--- TODO: say more about conjugate thing and conserved probability. -->
<!---  -->
<!--- % todo: demos of schrodinger.. -->

