+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

{id="figure_complex" style="height:30em"}
![Complex numbers are just a way of representing two real values with one expression, which is essentially a vector of two numbers. These two values are aligned along two separate orthogonal dimensions. The imaginary number $i$, where $i^2 = -1$, is what keeps these two values orthogonal --- the first value $a$ is along the real axis, and the second value $b$ is along the imaginary axis. The complex conjugate, $c^*$, is simply subtracting the imaginary part instead of adding it (i.e., it represents a reflection along the imaginary dimension). Multiplying $c c^*$ gives the squared magnitude of the vector, which is a single real-valued scalar number. It is the (squared) length of the hypotenuse of the vector. In the typical oscillatory use of complex numbers, they represent rotation around the complex plane, constrained to the (unit) circle as shown.](media/fig_complex_numbers.png)

The symbol $\chi$ is used to represent a complex-valued state variable:

$$
\chi \def a + i b
$$

$$
\def \phi_a + i \phi_b
$$

So, $\chi$ is composed of two separate real-valued numbers, designated $a$ and $b$ (or $\phi_a$ and $\phi_b$, to indicate that they are scalar state variables). A complex number is really just a way of representing a _vector_ of two separate real valued numbers, aligned along orthogonal dimensions, in an efficient and compact manner ([[#figure_complex]]). It is essential to appreciate that, despite the presence of the imaginary number $i$ (where $i^2 = -1$ or $i = \sqrt{-1}$), _all you ever really have is two real-valued numbers._ There is nothing "imaginary" or mysterious or spooky about the second number in a complex number: all the $i$ does is keep these two values separate from each other.

In the end, we will deconstruct all of our complex numbers into their real-valued components, and write purely real-valued expressions that determine their update rules. These expressions will be more complicated than the ones using complex numbers, but they are required for actually implementing the equations on the computer, and they also provide a more explicit and obvious indication of exactly what drives each value.

These are the conventions for state variables:

* $\phi$ = phi = simple scalar state value: one single real-valued number per cell.
* $\chi$ = chi = complex-number state value: two independent real-valued numbers per cell.
* $\psi$ = psi = full [[Dirac]] state value, which has two separate complex numbers, and captures the property of [[spin]]. We also use this for the [[Schrodinger]] wave state, as it is conventionally done, even though it is just a single complex number.

Here's a few interesting facts about complex numbers:

To do algebra on them, you just have to remember to _keep the real-values sorted separately from the imaginary ones,_ but otherwise treat them just like a pair of numbers:

$$
y = a + ib
$$

$$
z = c + id
$$

$$
y + z = (a + c) + i(b + d)
$$

$$
y z = a c + i a d + i b c - b d
$$

$$
= (ac - bd) + i (ad + bc)
$$

Notice that this multiplication rule is the same as $(a + b)(c + d) = ac + ad + bc + bd$, where you just multiply everything through, except that you end up with these $i$ terms crossing over, and when you multiply $ib$ and $id$, you end up with $i^2bd$, at which point the $i^2$ disappears into a $-1$ (i.e., it crosses over from the imaginary world into the real one).

As you should expect from [[#figure_complex]], adding two complex numbers is like adding two vectors, and multiplication is just like multiplying vectors. Complex numbers really are just a compact way of writing vectors!

Multiplication by $i$: If you multiply a complex number by $i$, then you basically switch the real and imaginary parts: the real moves to the imaginary position, and the imaginary becomes real:

$$
i(a + ib) = ia - b = -b + ia
$$

Geometrically, this is equivalent to rotating a vector by 90 degrees! If you do this four times, you'll end up back where you started (as you would expect by doing a 360).

The **complex conjugate** of a complex number is just that number with imaginary dimension inverted:

$$
y^* = a - i b
$$

This can also be written using a bar:
$$
\bar{y} = a - i b
$$

The primary use of such a thing is to find the magnitude of a complex number (i.e., the length of the vector that it represents), as:

$$
|y|^2 = y y^* = (a + i b)(a - i b)
$$

$$
= a^2 - iab + iab + b^2
$$

$$
= a^2 + b^2
$$

This should be recognizable as simply the pythagorean theorem for the squared length of the hypotenuse of a right triangle ($a^2 + b^2 = c^2$). Again, complex numbers have no mystery: they just represent a two-valued vector.

## Rotation by the exponential (Euler's formula)

One of the most useful tricks with complex numbers is that you can rotate them around the unit circle by multiplying by $e^{i\theta}$, where $\theta$ is the angle that it is rotated by (in radians). This is a result of Euler's formula:

$$
e^{i\theta} = \cos \theta + i \sin \theta
$$

Which is directly a rotation around the the complex plane: the real component is multiplied by the x-axis-like rotational factor, while the imaginary component gets the y-axis-like rotational factor. If $\theta$ is small, then cosine is close to 1 and sine is close to 0, and there isn't much rotation, etc.

The essential insight for _why_ Euler's formula is true, is that the derivative of the sine function is cosine, and vice-versa:

$$
\frac{d}{dx} \sin x = \cos x; \qquad \frac{d}{dx} \cos x = -\sin x
$$

and the natural exponential function $e^x$ is defined as the function whose derivative is itself!

$$
\frac{d}{dx} e^{x} = e^{x}
$$

So by putting the $i$ into this expression, Euler's formula emerges because the $i$ switches the real component into the imaginary one, and vice-versa:

$$
\frac{d}{dx} e^{i\theta} = \cos \theta + i \sin \theta = e^{i\theta}
$$

low and behold, the derivative is the same as the function itself!

Fundamentally, this deep connection between trigonometry and the natural exponential arises because the circle has this property, captured in the derivatives, that as you move along it, the _x_ and _y_ components are constantly accelerating and decelerating in just such a way as to keep the total radius constant. Thus, the slope (derivative) of one is equal to the value of the other, and vice-versa.

