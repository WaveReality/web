+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

The **four-vector** framework provides an especially compact way of representing the mathematics of wave functions and associated quantities, in terms of a single **space-time-vector** with the first element representing time, and the remaining three elements representing 3D space. This notation was initially developed by Minkowski for dealing with Einstein's [[special relativity]] theory, where space and time expand and contract together, and in some sense can be converted into each other.

The essence of the basic [[waves#wave]] function is indeed this dynamic interrelationship between the curvature of space (i.e., as the **momentum** factor in the [[Hamiltonian]]) and the resulting acceleration in time (i.e., the **energy** factor in the Hamiltonian), which is why special relativity can be seen as a consequence of the wave-like nature of physics:

{id="eq_wave" title="the wave equation"}
$$
\frac{\partial^2 y}{\partial t^2} = c^2 \frac{\partial^2 y}{\partial x^2}
$$

This interrelationship can be described as the **covariance** or _trading-off_ of time against space, when you rearrange the terms and set them equal to zero, and see that any change in space has to be _compensated for_ by a change in time, and vice-versa:

$$
\frac{\partial^2 y}{\partial t^2} - c^2 \frac{\partial^2 y}{\partial x^2} = 0
$$

The four-vector notation allows us to express this interrelationship in a particularly compact manner.

For the point $a$ in space-time, the **covariant** form of the four-vector is defined as:

{id="eq_cov" title="covariant four-vector"}
$$
{a}_\mu = (t,-x,-y,-z) = (a_t,-a_x,-a_y,-a_z) = (a_0,-a_1,-a_2,-a_3)
$$

The little $\mu$ (Greek "mu") subscript goes from $0..3$ in counting out the different items in the vector, as shown. The time and space coordinates have different signs here in a way that directly matches their relationship in the wave equation, capturing their _covariant_ nature.

In the simplest form, we assume that $c=1$, so you can just write a bare $t$. The more general form requires converting units of time into units of space by multiplying by the speed of light, so you need to write $ct$ instead of just $t$.

The other form of the four-vector is the **contravariant** form, which is indicated by the use of **superscripts** instead of subscripts: 

{id="eq_con" title="contravariant four-vector"}
$$
{a}^\mu = (t,x,y,z) = (a^t,a^x,a^y,a^z) = (a^0,a^1,a^2,a^3)
$$

The multiplication (**dot product**) of two four-vectors is defined in terms of these covariant and contravariant forms:

{id="eq_dot" title="four-vector dot product"}
$$
a \cdot b \equiv {a}^\mu {b}_\mu = {a}_\mu {b}^\mu \equiv a^0 b_0 + a^1 b_1 + a^2 b_2 + a^3 b_3 = a^t b^t - a^x b^x - a^y b^y - a^z b^z = \sum^3_{\mu = 0} {a}_\mu {b}^\mu
$$

And the derivative of a four-vector can also be defined. Just like four-vectors themselves, there are two versions, a covariant and a contravariant, where, potentially confusingly, the superscript / subscript relationship is _flipped_ for the derivatives!

The _covariant derivative_ doesn't have any minus signs:

{id="eq_cov-deriv" title="covariant derivative"}
$$
\partial_\mu \equiv \frac{\partial {}}{\partial ^\mu} \equiv \left(\frac{\partial {}}{\partial {a^0}},\frac{\partial {}}{\partial {a^1}},\frac{\partial {}}{\partial {a^2}},\frac{\partial {}}{\partial {a^3}}\right) \equiv \left(\frac{\partial {}}{\partial {t}}, \vec{\nabla} \right)
$$

where the $\vec{\nabla}$ symbol represents the spatial _gradient_ operator:

{id="eq_grad" title="spatial gradient"}
$$
\vec{\nabla} \equiv \left(\frac{\partial {}}{\partial {x}}, \frac{\partial {}}{\partial {y}}, \frac{\partial {}}{\partial {z}}\right)
$$

The *contravariant derivative* is the same deal, except it now has the minus signs:

{id="eq_con-deriv" title="contravariant derivative"}
$$
\partial^\mu \equiv \frac{\partial {}}{\partial _\mu} \left(\frac{\partial {}}{\partial {a^0}},-\frac{\partial {}}{\partial {a^1}},-\frac{\partial {}}{\partial {a^2}},- \frac{\partial {}}{\partial {a^3}}\right) \equiv \left(\frac{\partial {}}{\partial t}, -\vec{\nabla} \right)
$$

Now, finally, for the payoff. If you take the second-order derivatives of a four-vector, you combine the vector multiplication rules with the derivative equations to get the following:

$$
\partial_\mu \partial^\mu = \frac{\partial {}}{\partial t} \frac{\partial {}}{\partial t} - \frac{\partial {}}{\partial {x}} \frac{\partial {}}{\partial {x}} - \frac{\partial {}}{\partial {y}} \frac{\partial {}}{\partial {y}} - \frac{\partial {}}{\partial {z}} \frac{\partial {}}{\partial {z}}
$$

$$
= \frac{\partial^2 {}}{\partial t^2} - \frac{\partial^2 {}}{\partial {x}^2} - \frac{\partial^2 {}}{\partial {y}^2} - \frac{\partial^2 {}}{\partial {z}^2}
$$

$$
= \frac{\partial^2 {}}{\partial t^2} - \nabla^2
$$

So we can now say that the basic wave equation is obtained by setting the second-order four-vector derivative to zero:

$$
\partial_\mu \partial^\mu s = 0
$$

$$
\left(\frac{\partial^2 {}}{\partial t^2} - \nabla^2 \right) s = 0
$$

$$
\frac{\partial^2 {s}}{\partial t^2} - \nabla^2 s = 0
$$

$$
\frac{\partial^2 {s}}{\partial t^2} = \nabla^2 s
$$

Although this is equivalent to our basic wave equation, this way of computing the math, with time and space included in the same overall derivatives, will simplify calculations. For example, all of [[Maxwell]]'s equations for the electromagnetic field can be expressed in a single compact expression:

{id="eq_maxwell" title="Maxwell's equation"}
$$
\partial_\mu \partial^\mu A^\mu = - k^\mu J^\mu
$$

Notice that here the second-order derivative has a "source" term (instead of being $=0$), which acts like a driving force on the waves: it represents the charge and currents that drive the electromagnetic field.

Finally, we introduce just two more items of terminology. First, sometimes we'll need to convert a contravariant four-vector into a covariant four-vector, and we can do this using something called the **metric tensor**, which has two equivalent forms (they differ for general relativity, but not for [[special relativity]]):

$$
g_{\mu\nu} = g^{\mu\nu} = \begin{bmatrix}
1 & 0  & 0  & 0 \\
0 & -1 & 0  & 0 \\
0 & 0  & -1 & 0 \\
0 & 0  & 0  & -1\\
\end{bmatrix}
$$

To convert from one form of four-vector to another, you just multiply (we arbitrarily choose $g^{\mu\nu}$ here):

$$
a^\mu = g^{\mu\nu} a_\mu
$$

$$
a_\mu = g^{\mu\nu} a^\mu
$$

Finally, as if we needed an even simpler version of the wave equation (and one more symbol to memorize), the **d'Alembertian** $\sqcap$ (note: $\sqcap$ should actually just be a square box, but we don't have that available for technical reasons):

$$
\sqcap \equiv \frac{\partial^2 {}}{\partial t^2} - \nabla^2 = \partial_\mu \partial^\mu
$$

allows you to write the wave equation in the simplest possible way, as:

$$
\sqcap s = 0
$$

