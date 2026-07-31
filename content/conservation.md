+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

A critical property of quantum wave functions is that they exhibit strict **conservation** of some overall quantity over time. In the case of the [[Schrodinger]] wave function, this quantity is interpreted as the overall **probability** of finding a particle in a given location, according to the standard [[Copenhagen]] interpretation. For the [[Klein-Gordon]] and [[Dirac]] functions, this quantity makes more sense as the **charge** of an associated particle such as an [[electron]], because the conserved quantity can be either positive or negative valued, just like charge (i.e., the anti-particle of the electron, the _positron_, has positive charge).

Here, we derive the relevant mathematical properties of this conserved quantity, which depends in general on the use of [[complex number]]s. The essential, intuitive reason for this is that the conserved quantity is like the **hypotenuse** of a triangle, i.e., the **radius** on a circle that the state of the system is continuously rotating through. If you only have one number, then that number will oscillate like the sine or cosine as the system oscillates. If you instead have the two numbers within a complex number, then you can compute this hypotenuse / radius as the square of these two numbers, which is what the **complex conjugate** function computes:

$$
|y|^2 = y y^* = (a + i b)(a - i b) = a^2 + b^2
$$

In addition to computing the total conserved value, we are also interested in computing the local **density** and **current** flow of this conserved value, through the use of a **continuity equation**, which in the case of a conserved charge value then allows one to couple to the electromagnetic wave functions ([[Maxwell]]).

Starting at the most general level, the mathematical definition of a conserved quantity is that the sum total (i.e., integral) of its value across all of space does not change. For the case of a complex-valued wave state $\phi$, this is the complex conjugate: $\phi \phi^*$. Therefore, the appropriate integral is:

$$
\int_{-\infty}^{+\infty} \phi(t,\vec{x}) \phi^*(t,\vec{x}) dx
$$

where we have expressed the state value as a continuous function of both time and spatial coordinates, which we subsequently drop for convenience.

To determine the conserved quantity, we set the rate of change of this integral to zero, which means that its value cannot change over time:

$$
\frac{\partial }{\partial t}\left[\int_{-\infty}^{+\infty} \phi \phi^* dx\right] = 0
$$

by propagating the temporal derivative into the integral, you get:

{id="eq_conserve-int" title="Conserved value"}
$$
\int_{-\infty}^{+\infty} \left( \frac{\partial \phi^*}{\partial t} \phi + \phi^* \frac{\partial \phi}{\partial t} \right) dx = 0
$$

For Schrödinger's equation, we have a definition of this first-order time differential directly within the wave function itself:

{id="eq_schrodinger" title="Schrödinger's equation"}
$$
i \hbar \frac{\partial {\phi}}{\partial t} = -\frac{\hbar^2}{2 m_0} \nabla^2 \phi + V \phi
$$

Therefore, the right-hand-side of this equation can be substituted directly, resulting in this definition of a **probability gradient**, which is the directional (vector) flow of the conserved probability across space and time:

$$
\vec{J} \equiv - \frac{i \hbar}{2 m_0} \left( \phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right)
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
\phi^* (\partial_{\mu} \partial^{\mu} + m_0^2) \phi - \phi (\partial_{\mu} \partial^{\mu} + m_0^2) \phi^* = 0
$$

where the KG wave equation can be written in four-vector notation as:

$$
(\partial_{\mu} \partial^{\mu} + m_0^2) \phi = 0
$$

because two of these four-derivatives gives you:

$$
\partial_{\mu} \partial^{\mu} = \frac{\partial^2}{\partial t^2} - \nabla^2
$$

Thus, the starting expression amounts to multiplying the standard KG wave equation by the complex conjugate $\phi^*$, and subtracting the opposite configuration, which is the KG wave equation operating on the conjugate variable $\phi^*$, multiplied by the wave state $\phi$.

If we take the first half of this expression, it is:

$$
\phi^* (\partial_{\mu} \partial^{\mu} + m_0^2) \phi
$$

$$
\partial_{\mu} (\phi^* \partial_{\mu} \phi) - (\partial_{\mu} \phi^*)(\partial^{\mu} \phi) + m_0^2 \phi^* \phi
$$

and for the opposite configuration:

$$
\phi (\partial_{\mu} \partial^{\mu} + m_0^2) \phi^*
$$

$$ 
\partial_{\mu} (\phi \partial_{\mu} \phi^*) - (\partial_{\mu} \phi)(\partial^{\mu} \phi^*) + m_0^2 \phi \phi^*
$$

So when you subtract them, the second and third terms in each expression are the same, and cancel out, leaving only the difference in the first terms:

$$
(\partial_{\mu} \phi^*)(\partial^{\mu} \phi) - (\partial_{\mu} \phi)(\partial^{\mu} \phi^*) = 0
$$

$$
\partial_{\mu} (\phi^* \partial^{\mu} \phi - \phi \partial^{\mu} \phi^*) = 0
$$

Now this is exactly what we were looking for, if we recognize that this is an expression where the four-derivative of something equals zero. That something must be the conserved four-current, $J^{\mu}$:

$$
\partial_{\mu} J^{\mu} = 0
$$

$$
\partial_{\mu} (\phi^* \partial^{\mu} \phi - \phi \partial^{\mu} \phi^*) = 0
$$

{id="eq_charge-current-4vec" title="charge-current four vector"}
$$ 
J^{\mu} = \phi^* \partial^{\mu} \phi - \phi \partial^{\mu} \phi^*
$$

We can then separate this into the charge ($J^0$) and current components, as:

$$
  \frac{J_0}{c} = \rho \equiv \frac{i \hbar e}{2m_0c^2} \left( \phi^* \frac{\partial \phi}{\partial t} - \phi \frac{\partial \phi^*}{\partial t} \right)
$$

and

$$
\vec{J} \equiv - \frac{i \hbar e}{2m_0} \left( \phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right)
$$

where the extra conversion constants are inherited from the analogous expression for Schrödinger's equation, and ensure that the KG equation reduces to it in the non-relativistic limit.

To actually compute these in a simulation, we need to again break down the complex field value $\phi$ into its two scalar components. We simplify $\hbar=1$ and $c=1$, but preserve the mass term, and use $a = \varphi_a$ and $b = \varphi_b$ for ease of calculation:

$$
\rho = \frac{ie}{2m_0} \left(\phi^* \frac{\partial \phi}{\partial t} - \phi \frac{\partial \phi^*}{\partial t} \right)
$$

$$
\phi^* \frac{\partial \phi}{\partial t} - \phi \frac{\partial \phi^*}{\partial t} = (a - ib) (\dot a + i \dot b) - (a + ib) (\dot a - i \dot b)
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
\rho = \frac{e}{m_0} (\varphi_b \dot \varphi_a - \varphi_a \dot \varphi_b)
$$

and for the spatial (current) terms, it boils down to the same kind of equation in the end:

$$
\vec{J} = -\frac{ie}{2m_0} \left(\phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* \right)
$$

$$
\phi^* \vec{\nabla} \phi - \phi \vec{\nabla} \phi^* = 2 i a \vec{\nabla} b - 2 i b \vec{\nabla} a
$$

$$
\vec{J} = \frac{e}{m_0} (a \vec{\nabla} b - b \vec{\nabla} a)
$$

$$
\vec{J} = \frac{e}{m_0} (\varphi_a \vec{\nabla} \varphi_b - \varphi_b \vec{\nabla} \varphi_a)
$$

These are the expressions that are used in [[Dirac]] to couple with the electromagnetic field.


