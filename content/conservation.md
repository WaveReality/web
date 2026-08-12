+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

> Nature is a meticulous accountant.

therefore:

> Physics is really just a branch of accounting.

The single most important property of fundamental physical laws is that they exhibit **conservation** properties, such as the **conservation of total energy**. Consider what a disaster it would be if this weren't the case: stuff would either drain away over time, leaving nothing left behind, or accumulate, filling every corner of space with too much energy!

Furthermore, these conservation laws need to be **local** in space and time, according to the principles of [[special relativity]] -- this is known as **Lorentz invariance** in this context. This is also consistent with the principles of the [[cellular automaton]] (CA) framework, where all interactions are local in space and time, defined over the local neighborhood of grid cells and discrete time update steps.

Thus, you can see that another way of saying something is conserved is to say that it is **invariant** -- it doesn't change. Indeed, the entire edifice of the [[Standard Model]] of quantum physics is built upon [[gauge theory]], which is a mathematical framework for deriving locally invariant laws. This framework starts with the [[Lagrangian]] and [[Hamiltonian]] expressions of the total energy in a system, and derives _laws of motion_ that are **locally gauge invariant**, which means they conserve some kind of value (which turns out to be **charge**) locally.

According to Emmy Noether's famous theorems ([[@Noether18]]), each conservation law implies a kind of **symmetry** in Nature. Although this seems very "deep" on the one hand, it also can seem fairly "trivial" or obvious on the other, because indeed the meaning of the word _symmetry_ in this context is essentially just _invariance_. For example, the symmetry (invariance) with respect to time is associated with the conservation of energy. Yeah, that's what it means for energy to not change over time. It is invariant over time. Time is right there in the definition of what energy is being conserved over. It isn't invariant over space!

Mathematically, even though everything ultimately obeys local conservation laws, the specification of these laws can be divided into **global gauge invariant** conditions and local ones. Global invariants amount to transforming (translating or multiplying) the overall state of the system by _constant_ factors that apply the same way everywhere, while local ones emerge from introducing new _local_ transformation parameters at each point in spacetime. These local parameters represent a new _field_ of values, with the result specifying how two fields interact in a way that is locally conservative.

The canonical example of the local gauge invariance is quantum electrodynamics ([[QED]]), where the two fields represent [[Maxwell]]'s electromagnetic (EM) field equations interacting with the [[Dirac]] field for [[electron]]s as the source of those EM fields. The locally conserved quantity from this interaction is electric charge.

{id="table_conserve" title="Physical Conservation Laws"}
| Conservation law      | Invariance              | Parameters        |
|-----------------------|-------------------------|-------------------|
| energy                | time translation        | 1: t              |
| linear momentum       | space translation       | 3: x,y,z          |
| angular momentum      | rotation                | 3: x,y,z          |
| boost                 | "velocity" (space-time) | 3: x,y,z          |
|-----------------------|-------------------------|-------------------|
| electric charge       | U(1) gauge invariance   | 1: V potential    |
| weak isospin          | SU(2) gauge invariance  | 1: weak potential |
| strong (color) charge | SU(3) gauge invariance  | 3: r,g,b          |

There are 7 total conservation laws that define the Standard Model, shown in [[#table_conserve]]. The first four are global spatial transformations associated with the geometry of special relativity (known as a **Minkowski space**) that preserve the space-time interval between events, known as the [Poincaré group](https://en.wikipedia.org/wiki/Poincar%C3%A9_group), which define the most basic conservation laws of physics, listed as the top . There is a whole branch of mathematics surrounding such **Lie groups** (pronounced "Lee"), which you can read about in the above Wikipedia link.

The final 3 entries in the table define the charge-like values conserved (via the local gauge invariance mechanism in [[gauge theory]]) by the 3 fundamental forces in the Standard Model. These are defined by a _unitary group_ (U) or _special unitary group_ (SU) of the given dimensionality. A unitary group just means that it is conservative, like multiplying by the unit 1. The SU(2) group is defined by a 2x2 unitary matrix, which is the rotation group that is defined by a single parameter (the rotation angle), consistent with the idea that this group defines the property of [[spin]]. Thus, the [[weak]] force is fundamentally about spin, and the conserved charge-like value is weak isospin, although this conservation status is complicated by the [[Higgs]] field.

As a bit of a simplification (but not too much of one):

> The 7 conservation / invariance laws represented in [[#table_conserve]] entirely determine the nature of fundamental physics!

Thus, the quotes at the top of the page capture an essential truth: physics is fundamentally about allowing _something_ "interesting" to happen, while keeping the accounting airtight. Things can change, and exert forces on other things, but only in ways that don't ruin everything for the next generation, and which are fundamentally consistent everywhere. Nature follows the golden rule, even if we don't.

## Change without change: waves, spin and phase

The seemingly strange properties of [[special relativity]], and the associated first four conservation laws in [[#table_conserve]], are all satisfied by the basic [[wave]] equation, which lies at the heart of the electromagnetic field equations defined by [[Maxwell]]. A pure, frictionless wave is basically _the_ way to move energy around from one place to another in a way that conserves total energy, and is completely local in space and time.

Thus, the fact that all of quantum physics is built upon the machinery of wave equations can be seen as a direct consequence of the fundamental conservation laws. And special relativity can likewise be seen as a consequence of the properties of waves.

The energy-conserving nature of waves derives from the fact that they are just vehicles for transmitting **perturbations** around -- the wave medium itself is "immaterial" (unlike the classical notion of the [[aether]]): what matters is the disturbance or perturbation that gets passed along from one place to the next as the waves propagate.

However, a simple scalar wave field does not support the ability to represent a conserved scalar quantity like electrical charge, in a way that could be localized into a confined region of space. To do that, one needs _two_ coordinated waves, where the _phase_ relationship between these two waves can then define something that acts like an electrical charge. Thus, a moving charge could be represented as a moving phase perturbation between two wave states.

Mathematically, one way to represent the coordinated phase relationship is via [[complex number]]s, which naturally leads to the phenomenon of [[spin]], where the real and imaginary components of a complex number are constantly rotating around the complex plane. Thus, in short order, we now encounter all of the major ingredients of the quantum wave functions, directly from the conservation laws.

The strange properties of the [[weak]] interactions are all tied up with dealing with spin, and the difficulties in dealing with mass in the context of the gauge theory framework. Once you introduce a mass factor into the wave equation, as in the [[Klein-Gordon]] wave, it immediately fixes a specific gauge to the wave -- it is no longer gauge invariant. The strange properties of the [[Higgs]] mechanism within the overall weak interaction framework manage this delicate balance between the conservative properties of gauge invariance, and the need to represent the mass of massive particles.

## Symmetry for symmetry's sake

In the above account, _conservation laws_ are primary, and symmetry is a natural consequence thereof. However, there have been a number of attempts to build on the idea that symmetry itself is the primary thing driving the laws of nature, in the form of various "supersymmetry" theories. However, none of the unique predictions from these theories have been supported by the considerable evidence accumulated to date, and most scientists have concluded that this is a failed effort. Thus, it looks like this "minimal" set of conservation laws is all that is really going on.

## Conserved charge and continuity equations

In the remainder of this page, we develop the mathematical understanding of how a complex-valued wave function can represent a conserved quantity. This approach is based on continuity equations, which are appropriate for understanding the conservation properties of a single "free particle" wave function. The [[gauge theory]] framework is needed when describing conservation laws involving the _interaction_ between two or more waves.

In the case of the [[Schrodinger]] wave function, the conserved quantity that can be extracted from the phase relationships is interpreted as the overall **probability** of finding a particle in a given location, according to the standard [[Copenhagen]] interpretation. For the [[Klein-Gordon]] and [[Dirac]] functions, this quantity makes more sense as the **charge** of an associated particle such as an [[electron]], because the conserved quantity can be either positive or negative valued, just like charge (i.e., the anti-particle of the electron, the _positron_, has positive charge).

Here, we derive the relevant mathematical properties of this conserved quantity, which depends in general on the use of [[complex number]]s. The essential, intuitive reason for this is that the conserved quantity is like the **hypotenuse** of a triangle, i.e., the **radius** on a circle that the state of the system is continuously rotating through. If you only have one number, then that number will oscillate like the sine or cosine as the system oscillates. If you instead have the two numbers within a complex number, then you can compute this hypotenuse / radius as the square of these two numbers, which is what the **complex conjugate** function computes:

$$
|y|^2 = y y^* = (a + i b)(a - i b) = a^2 + b^2
$$

In addition to computing the total conserved value, we are also interested in computing the local **density** and **current** flow of this conserved value, through the use of a **continuity equation**, which in the case of a conserved charge value then allows one to couple to the electromagnetic wave functions ([[Maxwell]]).

Starting at the most general level, the mathematical definition of a conserved quantity is that the sum total (i.e., integral) of its value across all of space does not change. For the case of a complex-valued wave state $\chi$, this is the complex conjugate: $\chi \chi^*$. Therefore, the appropriate integral is:

$$
\int_{-\infty}^{+\infty} \chi(t,\vec{x}) \chi^*(t,\vec{x}) dx
$$

where we have expressed the state value as a continuous function of both time and spatial coordinates, which we subsequently drop for convenience.

To determine the conserved quantity, we set the rate of change of this integral to zero, which means that its value cannot change over time:

$$
\frac{\partial }{\partial t}\left[\int_{-\infty}^{+\infty} \chi \chi^* dx\right] = 0
$$

by propagating the temporal derivative into the integral, you get:

{id="eq_conserve-int" title="Conserved value"}
$$
\int_{-\infty}^{+\infty} \left( \frac{\partial \chi^*}{\partial t} \chi + \chi^* \frac{\partial \chi}{\partial t} \right) dx = 0
$$

For Schrödinger's equation, we have a definition of this first-order time differential directly within the wave function itself:

{id="eq_schrodinger" title="Schrödinger's equation"}
$$
i \hbar \frac{\partial {\chi}}{\partial t} = -\frac{\hbar^2}{2 m_0} \nabla^2 \chi + V \chi
$$

Therefore, the right-hand-side of this equation can be substituted directly, resulting in this definition of a **probability gradient**, which is the directional (vector) flow of the conserved probability across space and time:

$$
\vec{J} \equiv - \frac{i \hbar}{2 m_0} \left( \chi^* \vec{\nabla} \chi - \chi \vec{\nabla} \chi^* \right)
$$

Another more general route for computing this probability gradient is to start with the **continuity equation** in terms of the _density_ $\rho$ and _current_ $\vec{J}$:

$$
\frac{\partial \rho}{\partial t} + \vec{\nabla} \cdot \vec{J} = 0
$$

$$
\frac{\partial \rho}{\partial t} = - \vec{\nabla} \cdot \vec{J}
$$

where $\vec{\nabla} \cdot$ is the _divergence_ of a vector quantity:

$$
\vec{\nabla} \cdot \vec{J} \equiv \frac{\partial J_x}{\partial x} + \frac{\partial J_y}{\partial y} + \frac{\partial J_z}{\partial z}
$$

(this is just the sum of the spatial derivatives along each spatial direction). We cover divergence in greater detail in [[Maxwell]] --- it represents the amount of new "stuff" accumulating in a given region of space from the flow in from its neighbors. If the total amount of stuff is to remain constant, then this increment needs to be offset by a change in the amount of stuff in that region itself, which is $\frac{\partial }{\partial t}\rho}$. This is what this equation captures.

This continuity relationship can be expressed in a [[four-vector]] (space-time) derivative notation, in terms of a single four-vector charge / current variable $J^{\mu}$:

$$
\partial_{\mu} J^{\mu} = \frac{\partial J^0}{\partial t} + \frac{\partial{J^1}}{\partial{x}} + \frac{\partial{J^2}}{\partial{y}} + \frac{\partial{J^3}}{\partial{z}}
$$

$$
\partial_{\mu} J^{\mu} = \frac{\partial J^0}{\partial t} + \vec{\nabla} \cdot  \vec{J}
$$

where $J^0 = c \rho$ (so, charge, like energy, is a _time-like_ quantity, whereas current is a space-like quantity). The continuity relationship in these terms is therefore:

$$
\partial_{\mu} J^{\mu} = 0
$$

This one expression succinctly captures the key idea that the time-like first element is trading-off against the three space-like elements to produce an overall conservation of current, very much in the same way that the basic wave equations involve a tradeoff between time and space derivatives.

From here, we can go back to [[#eq_conserve-int]] and establish the connection with the continuity equation, in the context of the KG wave function ([[@Greiner00]]; [[@Gingrich04]]), :

$$
\chi^* (\partial_{\mu} \partial^{\mu} + m_0^2) \chi - \chi (\partial_{\mu} \partial^{\mu} + m_0^2) \chi^* = 0
$$

where the KG wave equation can be written in four-vector notation as:

$$
(\partial_{\mu} \partial^{\mu} + m_0^2) \chi = 0
$$

because two of these four-derivatives gives you:

$$
\partial_{\mu} \partial^{\mu} = \frac{\partial^2}{\partial t^2} - \nabla^2
$$

Thus, the starting expression amounts to multiplying the standard KG wave equation by the complex conjugate $\chi^*$, and subtracting the opposite configuration, which is the KG wave equation operating on the conjugate variable $\chi^*$, multiplied by the wave state $\chi$.

If we take the first half of this expression, it is:

$$
\chi^* (\partial_{\mu} \partial^{\mu} + m_0^2) \chi
$$

$$
\partial_{\mu} (\chi^* \partial_{\mu} \chi) - (\partial_{\mu} \chi^*)(\partial^{\mu} \chi) + m_0^2 \chi^* \chi
$$

and for the opposite configuration:

$$
\chi (\partial_{\mu} \partial^{\mu} + m_0^2) \chi^*
$$

$$ 
\partial_{\mu} (\chi \partial_{\mu} \chi^*) - (\partial_{\mu} \chi)(\partial^{\mu} \chi^*) + m_0^2 \chi \chi^*
$$

So when you subtract them, the second and third terms in each expression are the same, and cancel out, leaving only the difference in the first terms:

$$
(\partial_{\mu} \chi^*)(\partial^{\mu} \chi) - (\partial_{\mu} \chi)(\partial^{\mu} \chi^*) = 0
$$

$$
\partial_{\mu} (\chi^* \partial^{\mu} \chi - \chi \partial^{\mu} \chi^*) = 0
$$

Now this is exactly what we were looking for, if we recognize that this is an expression where the four-derivative of something equals zero. That something must be the conserved four-current, $J^{\mu}$:

$$
\partial_{\mu} J^{\mu} = 0
$$

$$
\partial_{\mu} (\chi^* \partial^{\mu} \chi - \chi \partial^{\mu} \chi^*) = 0
$$

{id="eq_charge-current-4vec" title="charge-current four vector"}
$$ 
J^{\mu} = \chi^* \partial^{\mu} \chi - \chi \partial^{\mu} \chi^*
$$

We can then separate this into the charge ($J^0$) and current components, as:

$$
  \frac{J_0}{c} = \rho \equiv \frac{i \hbar e}{2m_0c^2} \left( \chi^* \frac{\partial \chi}{\partial t} - \chi \frac{\partial \chi^*}{\partial t} \right)
$$

and

$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \chi^* \vec{\nabla} \chi - \chi \vec{\nabla} \chi^* \right)
$$

where the extra conversion constants are inherited from the analogous expression for Schrödinger's equation, and ensure that the KG equation reduces to it in the non-relativistic limit.

To actually compute these in a simulation, we need to again break down the complex field value $\chi$ into its two scalar components. We simplify $\hbar=1$ and $c=1$, but preserve the mass term, and use $a = \phi_a$ and $b = \phi_b$ for ease of calculation:

$$
\rho = \frac{ie}{2m_0} \left(\chi^* \frac{\partial \chi}{\partial t} - \chi \frac{\partial \chi^*}{\partial t} \right)
$$

$$
\chi^* \frac{\partial \chi}{\partial t} - \chi \frac{\partial \chi^*}{\partial t} = (a - ib) (\dot a + i \dot b) - (a + ib) (\dot a - i \dot b)
$$

$$
= (a \dot a + i a \dot b - i b \dot a + b \dot b) - (a \dot a - i a \dot b + i b \dot a + b \dot b)
$$

$$
= 2 i a \dot b - 2 i b \dot a
$$

$$
\rho = \frac{e}{m_0} (b \dot a - a \dot b) 
$$

$$
\rho = \frac{e}{m_0} (\phi_b \dot \phi_a - \phi_a \dot \phi_b)
$$

and for the spatial (current) terms, it boils down to the same kind of equation in the end:

$$
\vec{J} = -\frac{ie}{2m_0} \left(\chi^* \vec{\nabla} \chi - \chi \vec{\nabla} \chi^* \right)
$$

$$
\chi^* \vec{\nabla} \chi - \chi \vec{\nabla} \chi^* = 2 i a \vec{\nabla} b - 2 i b \vec{\nabla} a
$$

$$
\vec{J} = \frac{e}{m_0} (a \vec{\nabla} b - b \vec{\nabla} a)
$$

$$
\vec{J} = \frac{e}{m_0} (\phi_a \vec{\nabla} \phi_b - \phi_b \vec{\nabla} \phi_a)
$$

These are the expressions that are used in [[Dirac]] to couple with the electromagnetic field.


