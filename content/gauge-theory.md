+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

**Gauge theory** is a central tool in the [[Standard Model]], used to predict the full form of the [[weak]] and [[strong]] forces, and how they interact with their sources. These forces can be transformed in various ways that do not alter their overall effects, known as a **gauge invariance**. This reflects the essential and defining role that [[conservation]] laws have in shaping the forces of nature: each conservation law is associated with a corresponding type of invariance.

The principle of gauge invariance can be illustrated in the case of the electromagnetic force interaction with a "particle" defined by a complex wave state $\chi$. The classical [[Lagrangian]] field density can be defined as the kinetic energy minus a potential energy, where, in [[four-vector]] notation, the kinetic energy is the 2nd order temporal derivative of the wave function (i.e., the contravariant times the covariant):

$$
L = T - V
$$

$$
L_0(\chi(x), \partial^\mu\chi(x)) = \partial_\mu\chi^*\partial^\mu\chi - V(\chi^* \chi)
$$

## Global gauge invariance

The integral form of the Lagrangian action essentially involves a subtraction between the ending and starting action values. Therefore, adding or multiplying the wave state by a constant global factor has no effect on the resulting laws of motion. This is known as a **global gauge transformation** and it gives rise to a corresponding symmetry, in the form of a conserved current (see [[continuity]] for the derivation):

$$
j^\mu = \text{const.} \chi^* \partial^\mu \chi
$$

$$
\partial_\mu j^\mu = 0
$$

This form of global gauge invariance and its resulting conserved current symmetry constitute the $U(1)$ symmetry group. This group essentially means that the field dynamics have an associated strictly conserved value, that is thereby unaffected by global gauge transformations.

## Local gauge invariance

Instead of a global constant factor, it is also possible to introduce a _local_ factor that varies as a function of position, known as a **local gauge transformation**. To specify the value of this local transformation at every point, an _additional field_ of parameters is required, representing additional degrees of freedom in the physical system in question.

When this local gauge transformation is  applied, the resulting laws of motion are _not_ automatically preserved. However, one can change the definition of the _contravariant derivative_ $\partial^\mu \chi$ in a way that produces the same effects as the local gauge transformation.

This new form of the derivative defines a **minimal coupling** between the new field of parameters and the original state field, and mathematically determines how to update the Lagrangian and the resulting laws of motion to specify this interaction between the two fields.

In effect, introducing the local gauge transformation is the primary way to introduce a coupling between two wave fields. In the example we explore next, the parameters of the local gauge transformation are the spatial distributions of electric charge, which then serve as the source of the electrical potential field.

Critically, without the local gauge transformation and these additional charge parameters, the system could only represent a single global charge value.

Thus, the local gauge transformation represents physically-necessary additional degrees of freedom required to specify the nature of real physical systems. As such, the abstract symmetry properties are not really the primary motivation for introducing a local gauge transformation: instead it is really motivated by the physical reality of the system under consideration. Nevertheless, the gauge theory framework enables one to formalize the interactions in a way that evidently aligns with how it actually works in nature. The critical feature is that the approach guarantees the conservation (symmetry) of the source current. See this [stackexchange discussion](https://physics.stackexchange.com/questions/384978/why-do-can-we-impose-local-gauge-invariance) for an explanation that is consistent with this.

Another application of this approach in the [[weak]] interaction is to introduce the [[Higgs]] field to specify the masses of individual particles. Without this additional parameterization field, defined as a local gauge transformation, the [[Klein-Gordon]] wave function can only describe particles with exactly the same rest mass.

### Electromagnetic minimal coupling

To see how the local gauge transformation works in practice, we can use the example of the electromagnetic field coupled to the complex matter field, i.e, as in the case of the [[complex KG]] system.

Here is the local gauge transformation:

$$
\chi(x) \rightarrow U(x) \chi(x)
$$

$$
U(x) = e^{-i \alpha(x)}
$$

where $\alpha(x)$ is an arbitrary function specifying the phase of $U(x)$ at every point -- we will parameterize this with an additional field in a moment.

If you apply the covariant four-vector derivative to this transformed field, as required by the Lagrangian, you end up with an extra term:

$$
\partial_\mu \chi(x) \rightarrow U(x) \partial_\mu \chi(x) + \chi(x) \partial_\mu U(x)
$$

This extra term means that the system is not invariant to this local gauge transformation. No such system would be -- this is the whole point.

Now, you introduce a new **gauge field** that is going to represent the local transformation parameters, in the form of the electromagnetic vector potential field $A^\mu(x)$

$$
A^\mu(x) \rightarrow A^\mu(x) + \frac{1}{e} \partial^\mu \alpha(x)
$$ 

where $e$ is the **coupling constant** that will parameterize the strength of the interaction between this EM field and the original complex particle field.

Next, we define our own new **covariant derivative** in a way that will make all the math work out under the local gauge transformation:

$$
D_\mu \chi(x) \def [\partial_\mu + i e A^\mu(x)] \chi(x)
$$

which can then be locally transformed just like we did above:

$$
D_\mu \chi(x) \rightarrow U(x) D_\mu \chi(x)
$$

todo: finish!

## Local Gauge Invariance

TODO: old version, from Complex-KG:

If you look at it in the right way, the electromagnetic field can be seen as a way of canceling out an extra degree of freedom present in the complex KG wave field equations. As discussed in [[Maxwell]], the Lorenz gauge is an example of a situation where we introduced some extra constraints on the field variables, and this eliminated a degree of freedom in the electromagnetic equations, and also made them appear a lot simpler than they otherwise would.

This notion of a _gauge_ is very general: it just means that whenever you have some unconstrained values in your equations (i.e., values that can change without changing the observable results that you can measure in physics experiments), then you need to apply some kind of gauge to fix these variables. In the Lorenz gauge example, the extra degree of freedom comes from the fact that the observable variables are the force vectors $\vec{E}$ and $\vec{B}$, which are essentially derivatives of the underlying potentials $A_0$, $\vec{A}$. Thus, the raw values of the potentials can be moved up or down, and it won't change the slope of the fields, and therefore it won't change the observable force vectors.

However, it is not clear exactly how to reconcile such an argument with the fact that the potentials appear directly in our coupling equations, and also are observable in terms of the Arahnov-Bohm effect, as discussed elsewhere.

Nevertheless, here is the argument for the electromagnetic coupling being a form of local gauge invariance. If you just take our basic complex KG wave equation, you can get exactly the same overall behavior if you multiply the thing by a "phase transformation" (it is often said that gauge invariance should really be phase invariance) which is basically just a rotation along the complex axes. This is exactly the kind of rotation discussed above. The generic form of a rotation in complex numbers is to multiply by an exponential term:

$$
\chi(x) \rightarrow \exp \left(\frac{ie}{\hbar c} \chi \right) \chi(x)
$$

Where the $\chi$ term is the amount that you're rotating (the rest are just convenient constants for getting $\chi$ into the right units) --- think of it as some number of degrees of rotating the underlying $\phi_a$ value into $\phi_b$ (and vice-versa) for the complex number $\chi$.

If $\chi$ is independent of location ($x$), then it is just a constant and nothing happens. This is a *global* gauge/phase transformation, and it is not very interesting. However, if $\chi$ is now itself a function of location (i.e., $\chi(x)$ ), this is a *local* gauge transformation, and this is where the electromagnetic coupling comes in. If you have such a local variable, and you take the derivative of the resulting overall system that includes this locally-varying thing, you get this extra term for the derivative of $\chi(x)$ with respect to x:

$$
\partial_\mu \chi(x) \rightarrow \exp \left( \frac{ie}{\hbar c} \chi(x) \right) \left(\partial_\mu + \frac{ie}{\hbar c} \partial_\mu \chi(x) \right) \chi(x)
$$

So now your nice wave equation is a mess, and it varies from one place to another as a function of this $\partial_\mu \chi(x)$ term. So here is the trick: you basically just add this annoying term into the overall EM potential field (which is OK because such additions do not change the gradients and therefore do not affect observable EM field vectors):

$$
A_\mu(x) \rightarrow A_\mu(x) + \partial_\mu \chi(x)
$$

And then you just subtract this whole thing back out from your messy equation, and this gives you something just slightly less messy:

$$
\left(i \hbar \partial_\mu - \frac{e}{c} A_\mu\right)\chi
$$

So, this ends up being the same thing as the minimal coupling described earlier. Somehow, this whole process seems like a rather contrived way to end up with something that already made quite a bit of sense before hand. However, as noted earlier, the true payoff of such a procedure appears to come in addressing the weak and strong forces, which we leave for a future refinement of the model.

## Notes

* SU(2) is just the 3D rotation space -- spin!  So SU(2) x U(1) is pretty much determined by any spinning thing interacting with a massless radiative field. Not a general principle as much as a necessary fact, given that fermions spin..

* SU(3) is another story..

* Massive bosons are described by KG field -- what about range? need to investigate. theoretically it follows from that.

* Proca is 4-vector (vector boson) version of KG.


