+++
Name = "Dirac"
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

The **Dirac** wave function builds on the considerable progress made in the [[complex KG]] version of the [[Klein-Gordon]] (KG) wave function, which developed all of the tools needed to couple a [[matter wave]] that defines a **wave of charge** with the electromagnetic waves of [[Maxwell]]'s equations. This coupling of charge waves and EM waves has been pursued more recently in neoclassical self-coupled field theory ([[@JaynesCummings63]]; [[@CrispJaynes69]]; [[@BarutVanHuele85]]; [[@BarutDowling90]]; [[@Crisp96]]; [[@FinsterSmollerYau99a]]; [[@Radford03]]; [[@MasielloDeumensOhrn05]]).

This final step in our long journey of progressively more complicated wave functions, is to come to terms with the full glory of the [[electron]]. The standard model of physics has a set of parameters that are used to characterize the basic properties of the fundamental particles. One such property is the rest mass $m_0$, which we introduced in [[Klein-Gordon]] to make the waves slow down and move at variable speeds.

Another such property is electrical charge, which we were able to extract from our wave equation once we used a complex-valued state variable, having two independent scalar values within it, in [[complex KG]]. Furthermore, we found that this charge could come with two different signs, positive or negative. This turns out to be convenient, because electrons also come in a positive form, called a positron. The positron has the same mass as the electron, but just an opposite charge.

The third basic property of the electron is known as its **spin**. It is known as a spin $\frac{1}{2}$ particle, along with all of the other fundamental particles known (e.g., quarks). Unfortunately, our wave equations so far do not support this spin property, and so we'll need to do a little bit more work. However, once we're done, we'll find that our equations capture all of the fundamental properties of the electron: we should have a 100% complete description of it!

Actually there is one last thing, which is that the electron is a member of the [[lepton]] family, whereas the other fundamental particles are quarks, and they have other fundamental properties in addition to those carried by leptons. But, this is presumably because quarks live in some other set of state variables separate from the lepton state variables we're simulating here in our model. So, with that assumption, we might have captured everything about the electron.

So what exactly does it mean for an electron to have spin $\frac{1}{2}$? The quantum mechanical concept of spin is perhaps one of the most difficult to grasp. Sometimes people try to think of a little point particle spinning about like a top on its axis, but this doesn't actually fit the facts very well. In the end, the best strategy may be to just see what the equations we derive next actually do, and call that spin. Indeed, in the computer simulations, one can clearly see a spinning motion.

This spin is very much like the first-order Schrödinger equation dynamics, where the two different elements of the complex variable rotate into each other. We also saw this kind of rotation earlier in the complex coupled KG equation, where the electromagnetic potential introduced a rotation among the complex variables. In the present case, we're going to have four different variables, and they will all rotate amongst themselves to produce this mysterious spin.

Incidentally, quantum physics holds that photons (which we think of as wave packets of the electromagnetic field that we've already characterized above) have a spin of 1. Furthermore, the charged complex KG wave equation is described as having a spin of 0. This latter case makes sense to me, in that the two components of the complex number do not rotate into each other, and thus they do not spin at all. However, the electromagnetic field case is a bit more confusing, because as we saw, the four components of this field do not interact with each other in the basic wave equations either!

Therefore, it would seem that it should have a spin of 0 as well. Countering this are two considerations. First, the observable variables of the electric and magnetic fields $\vec{E}$ and $\vec{B}$, which are derived from these non-interacting electrical potentials, do rotate around each other as the wave propagates. Second, when these potentials interact with our charge wave, the do so in a way that ends up coupling (and rotating) the two independent scalar values in the complex number, and thus they impart some spin on our otherwise spinless particle.

Our first step is to introduce a new state variable $\psi$, to represent a field having four independent scalar values. Mathematically, this is defined as a vector of two complex numbers:

{id="eq_two-component" title="two-component complex wave state"}
$$
\psi = \begin{bmatrix} \phi_{1a} + i \phi_{1b} \\ \phi_{2a} + i \phi_{2b} \end{bmatrix}
$$

where the first complex number is denoted with a subscript $1$, and contains the two real-valued components $\phi_{1a}$ and $\phi_{1b}$, and the second has subscript $2$, and contains $\phi_{2a}$ and $\phi_{2b}$. So, the spin is going to amount to these four variables rotating amongst themselves.

How do we extend our basic complex-coupled KG equation to include this spin factor? Several authors in the literature have described a second-order version of the Dirac equation, which should look very familiar to you at this point, because it is essentially our current KG equation plus one additional spin term. One of the first references to such a thing comes from [[@^FeynmanGell-Mann58]], where they describe an equation that possesses all of the critical properties of the standard first-order Dirac equation, and note that it only requires four state variables instead of the eight required for the first-order equation. Indeed, Feynman states that he much prefers this form of the equation. This affection is presumably not widely shared, because there are relatively few other references to such an equation in the literature. Most of them come from a series of papers by Levere Hostler (e.g., [[@Hostler82]]; [[@Hostler83]]; [[@Hostler85]]).

The version of the equation described by [[@^FeynmanGell-Mann58]] is:

$$
\left[ \left(i {\nabla}_\mu - {A}_\mu\right)^2 + \vec{\sigma} \cdot \left(\vec{B} + i \vec{E} \right) \right] \psi = m_0^2 \psi
$$

where $\vec{\sigma}$ are the standard Pauli matricies that we'll describe in a moment. [[@^Hostler85]] describes a similar equation (which has the minus sign reversed in various places, but is otherwise the same):

$$
\left[ \left(-i \partial_\mu - e {A}_\mu\right)^2 + m_0^2 + e i \vec{\sigma} \cdot \left(\vec{E} + i \vec{B}\right) \right] \psi = 0
$$

It should be clear that the first squared term is just the complex KG equation coupled to the EM field. Therefore, we can write this equation in our current notation as:

$$
\left[\left(i \hbar \partial_\mu - \frac{e}{c}{A}_\mu \right)^2 + \frac{e}{c} \vec{\sigma} \cdot \left(\vec{B} + i \vec{E} \right) \right] \psi = m_0^2 c^2 \psi
$$

Now for the Pauli matricies $\vec{\sigma}$. This is a vector of values $(\sigma_x, \sigma_y, \sigma_z)$ that enter into a dot product with the complex-valued vector composed of the magnetic and electric field values $\vec{B}$ and $\vec{E}$:

{id="eq_pauli" title="Pauli matricies"}
$$
\vec{\sigma} = \left( \begin{bmatrix} 0 & 1\\
1 & 0 \end{bmatrix}, \begin{bmatrix} 0 & -i \\
i & 0 \end{bmatrix},  \begin{bmatrix} 1 & 0 \\
0 & -1 \end{bmatrix} \right)
$$

In the end, the net result of the dot product with an arbitrary vector $\vec{p}$ is:

$$
\vec{\sigma} \cdot \vec{p} = \begin{bmatrix} p_z & p_x - i p_y \\
p_x + i p_y & -p_z \end{bmatrix}
$$

So applied to our $\vec{B} + i\vec{E}$ vector, it is:

$$
\vec{\sigma} \cdot \left(\vec{B} + i \vec{E} \right) =
$$

$$
\begin{bmatrix} B_z + iE_z & B_x + iE_x - iB_y + E_y \\
B_x + iE_x + i B_y -E_y & -B_z - iE_z \end{bmatrix} =
$$

$$
\begin{bmatrix} B_z + iE_z & B_x + E_y + i(E_x - B_y) \\
B_x - E_y + i(E_x + B_y) & -B_z - iE_z \end{bmatrix}
$$

So, now we're getting some sense of how this works: different components of the electromagnetic field exert different forces on the different components of the $\psi$ state, producing a rotational effect.

This entire result then is multiplied by the two complex numbers in the $\psi$ state:

$$
\begin{bmatrix} B_z + iE_z & B_x + E_y + i(E_x - B_y) \\
B_x - E_y + i(E_x + B_y) & -B_z - iE_z \end{bmatrix} 
\times \begin{bmatrix} \phi_{1a} + i \phi_{1b} \\
\phi_{2a} + i \phi_{2b} \end{bmatrix}
$$

Which produces this for the first complex number:

$$
\phi_{1a} + i \phi_{1b} = (\phi_{1a} + i \phi_{1b})(B_z + iE_z) + (\phi_{2a} + i \phi_{2b})(B_x + E_y + i(E_x - B_y))
$$

which decomposes into the two scalar variables as:

$$
\phi_{1a} = \phi_{1a} B_z - \phi_{1b} E_z + \phi_{2a} (B_x + E_y) - \phi_{2b} (E_x - B_y)
$$

$$
\phi_{1b} = \phi_{1b} B_z + \phi_{1a} E_z + \phi_{2b} (B_x + E_y) + \phi_{2a} (E_x - B_y)
$$

And for the second complex number:

$$
\phi_{2a} + i \phi_{2b} = (\phi_{2a} + i \phi_{2b})(-B_z - iE_z) + (\phi_{1a} + i \phi_{1b})(B_x - E_y + i(E_x + B_y))
$$

which decomposes into:

$$
\phi_{2a} = -\phi_{2a} B_z + \phi_{2b} E_z + \phi_{1a} (B_x - E_y) - \phi_{1b} (E_x + B_y)
$$

$$
\phi_{2b} = -\phi_{2b} B_z - \phi_{2a} E_z + \phi_{1b} (B_x - E_y) + \phi_{1a} (E_x + B_y)
$$

So, with the minus sign flip that took place in the main equation, and the fact that all of the rest of the equation operates on each complex component of $\psi$ separately, without any mixing across components, the final update equations for this second-order Dirac equation are just the basic complex KG equations plus these terms:

{id="eq_dirac" title="Dirac functions in real, second-order form"}
$$
\begin{array}{lcl}
\ddot{\phi}_{1a} & = & \nabla^2 \phi_{1a} - m_0^2 \phi_{1a} + 2 e \left(A_0 \dot \phi_{1b} + \vec{A} \cdot \vec{\nabla} \phi_{1b} \right) +\\
& & e^2 \phi_{1a} \left(A_0^2 - \vec{A}^2 \right) + e \left( \phi_{1a} B_z - \phi_{1b} E_z + \phi_{2a} (B_x + E_y) - \phi_{2b} (E_x - B_y) \right)\\

\ddot \phi_{1b} & = & \nabla^2 \phi_{1b} - m_0^2 \phi_{1b} - 2 e \left(A_0 \dot \phi_{1a} + \vec{A} \cdot \vec{\nabla} \phi_{1a} \right) +\\
& & e^2 \phi_{1b} \left(A_0^2 - \vec{A}^2 \right) + e \left( \phi_{1b} B_z + \phi_{1a} E_z + \phi_{2b} (B_x + E_y) + \phi_{2a} (E_x - B_y) \right)\\

\ddot \phi_{2a} & = & \nabla^2 \phi_{2a} - m_0^2 \phi_{2a} + 2 e \left(A_0 \dot \phi_{2b} + \vec{A} \cdot \vec{\nabla} \phi_{2b} \right) +\\
& & e^2 \phi_{2a} \left(A_0^2 - \vec{A}^2 \right) + e \left( -\phi_{2a} B_z + \phi_{2b} E_z + \phi_{1a} (B_x - E_y) - \phi_{1b} (E_x + B_y) \right)\\

\ddot \phi_{2b} & = & \nabla^2 \phi_{2b} - m_0^2 \phi_{2b} - 2 e \phi_{2a} \left(A_0 \dot \phi_{2a} + \vec{A} \cdot \vec{\nabla} \phi_{2a} \right) +\\
& & e^2 \phi_{2b} \left(A_0^2 - \vec{A}^2 \right) + e \left( -\phi_{2b} B_z - \phi_{2a} E_z + \phi_{1b} (B_x - E_y) + \phi_{1a} (E_x + B_y) \right)\\
\end{array}
$$

Again, it is fundamentally the wave equation, plus three additional terms that characterize the interaction with the electromagnetic field. Note that, as with the mixing across complex components $\phi_a$ and $\phi_b$ that occurred in the previous version of the coupled KG equations, the mixing or spin across $\chi_1$ and $\chi_2$ occurs via the electromagnetic field interaction. This time, the vector force fields are now required for the coupling, requiring that we compute them from the potentials, as described earlier (involving the $\vec{\nabla}$ first-order gradient and, for the first time, the $\vec{\nabla} \times$ function, which is very similar in its discrete form to the $\vec{\nabla}$ function).

<!--- todo: run KG versions of the basic tests and see whether this is true!!! -->

Interestingly, although we need to continue the broken symmetry from the previous coupled-complex KG equation, where we use the current values of $\dot \phi_{1a}$ and $\dot \phi_{2a}$ to update the $\phi_{1b}$ and $\phi_{2b}$ variables, we apparently do not need to perform a similar symmetry breaking for the new couplings in this Dirac equation.

So, to answer the question of "what is spin?", we need only look at these equations. Spin, it seems, is this rotation of state values through the two complex variables in the $\psi$ state: $\chi_1$ and $\chi_2$. As is evident, this spinning occurs via interactions with the electromagnetic field vectors oriented along the three different spatial directions. The fact that, in our CA model we actually fix these directions according to the underlying cubic grid may seem strange and arbitrary. However, this does not mean that stuff can only spin along these fixed directions, anymore than it means that waves can only propagate in certain directions. By having different continuous values along these dimensions, any "direction" of spin relative to the underlying grid can occur.

Although somewhat complex, these equations should describe the entirety of the complexity of the electron's interaction with the electromagnetic field, which is to say, with other electrons and positive electric charges in the nucleus. Therefore, as we know, a huge proportion of the known complexity of the universe stems from the consequences of these basic equations. So, perhaps they do not look so complex in comparison.

<!--- todo: introduce coupled complex KG (CCKG) equation nomenclature % todo: also introduce covariant derivative term -->

One final thing to note is that the charge and current density equations from the previous version of the coupled complex KG equation still hold for this Dirac version, because these additional terms do not enter into the covariant derivative, and are therefore canceled out in the subtraction, just like the mass term. The actual numerical calculation changes only to accommodate the $\psi$ state value instead of the single complex $\chi$ state. The charge and current equations are:

$$
\rho = \frac{\hbar e}{m_0 c^2} \left((\phi_{1b} \dot \phi_{1a} - \phi_{1a} \dot \phi_{1b}) + (\phi_{2b} \dot \phi_{2a} - \phi_{2a} \dot \phi_{2b})\right) - \frac{e^2}{m_0 c^2} A_0 (\phi_{1a}^2 + \phi_{1b}^2 + \phi_{2a}^2 + \phi_{2b}^2)
$$

$$
\vec{J} = \frac{\hbar e}{m_0 c^2} \left((\phi_{1a} \vec{\nabla} \phi_{1b} - \phi_{1b} \vec{\nabla} \phi_{1a}) + (\phi_{2a} \vec{\nabla} \phi_{2b} - \phi_{2b} \vec{\nabla} \phi_{2a})\right) - \frac{e^2}{m_0 c^2} \vec{A} (\phi_{1a}^2 + \phi_{1b}^2 + \phi_{2a}^2 + \phi_{2b}^2)
$$

## First-order version

{id="eq_dirac-first" title="first-order Dirac equation"}
$$
(i \gamma^\mu \partial_\mu - m)\psi = 0
$$

TODO: there is still the important issue of converting back and forth between the 2nd order and 1st order versions of the state. Also, it is important that the 2nd order does not have spin factors operating among the components directly, as the 1st order version does -- they only show up in the coupling to EM. This is weird. The thing doesn't actually spin on its own in 2nd order form.

more refs: [[@Brown58]], [[@Tonin59]], [[@Marx67]], [[@Marx70]], [[@Case57]] <- weyl!!, [[@Diaz-CruzLopezMeza-AldamaEtAl15]], [[@KibblePolkinghorne58]], [[@BarutMullen62]], [[@BabinFigotin14]], [[@Cardoso93]], [[@Veblen33]], [[@DreinerHaberMartin10]]

[[@BilenkyPetcov87]] <- neutrinos

