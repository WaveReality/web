+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

**Gauge theory** is the _only_ mathematical framework that can tell us how a wave field that describes a [[particle]] (e.g., the [[Dirac]] field) should interact or _couple_ with a _force_ wave field, in a way that results in the [[conservation]] of overall energy and of the _charge density_ (or probability density) of the particle field. Remarkably, there is really only one way to do this, and furthermore, this only way of doing it _requires_ that the force field have _zero mass_, like the simple second-order wave equation, and [[Maxwell]]'s equations for the electromagnetic field.

This result was proven in the early 1960's by Goldstone and collaborators ([[@Goldstone61]]; [[@GoldstoneSalamWeinberg62]]), and the fact that the force-field [[boson]]s must be massless gives rise to the term **Goldstone boson**, which just refers to a zero-mass force field (in the context of it being required by gauge theory). The other relevant terms to know here are **Abelian** and **non-Abelian**, which are technical terms from group theory in mathematics, but in physics they amount to the difference between a force-field that doesn't interact with itself (like the EM field), versus one that does.

The context for all of this work at the time in the 1960's was that both the [[weak]] and [[strong]] forces seemed to be both _non-Abelian_ and massive, because they were known to only operate over a very short range, which is consistent with a massive field. 

> But if these force fields are massive, then according to Goldstone's theorem, they can't also be conservative!

This was _the_ major foundational problem in the field at the time, because everyone recognized how essential the principle of conservation is.

The solution to this problem arose via the [[Higgs]] mechanism and the full [[weak#electroweak]] theory developed starting with [[@^Weinberg67]]. But interestingly, the core mechanism behind this framework is actually already present in the most basic version of the Abelian gauge theory that we review below, and in some ways all the attention on the Higgs mechanism tends to distract from the essential nature of this solution. This solution was sketched in qualitative form already by [[@^Anderson63]], based on inspiration from superconductors and other non-quantum systems (see [[@^Witten16]] for an accessible history of these developments).

The critical phenomenon is:

> When the force and particle fields are coupled, there is a bidirectional **self-interaction** between these fields that ends up giving the force field an effective mass, in a way that depends on the strength of the force field! 

This is an essential feature of gauge coupling, which is often overlooked in the [[QED]] (electron, EM fields) case, because the strength of this coupling between the two fields is relatively weak, on the scale of $\alpha = \frac{1}{137}$. So the EM field does not gain very much mass from this interaction, and it can generally be neglected. But if the coupling strength is significantly larger, then this fundamentally nonlinear effect can become much more significant.

This self-interaction mass effect is only present in the region where the particle field has a non-zero magnitude, because the effect derives from the _charge current density_ of the field, which is a function of the field magnitude. This current density is what drives the force field (this is the source term of the standard Maxwell EM equations), and it represents one direction of the self-interaction. The other direction comes from an _additional_ effect on the current arising from the field forces that this current itself generates. It is this additional effect that ends up generating the mass of the force fields in the electroweak framework.

The Higgs field plays the role of the particle source in the electroweak system, and it ends up providing the equivalent of a uniformly distributed charge current density throughout the universe. Furthermore, because the electroweak system is non-Abelian, there are no less than 4 different force fields (some of which mutually interact with each other), and the way that the Higgs mass is distributed effectively creates one special configuration of these fields that actually remains massless, because the self-interaction terms end up canceling each other out, such that the current source density ends up being zero for such a configuration.

The result is that any waves that have this particular configuration will propagate over a long range, because they are effectively massless, while all the massive configurations will end up decaying over a relatively short range. Thus, the Higgs field acts much like a filter: the only thing that gets through is light itself (EM radiation over the effectively massless field configuration), while everything else is effectively the massive weak bosonic fields $W^\pm$ and $Z^0$.

In short, the inevitable mechanics of gauge coupling (the only possible way of doing things), which absolutely requires massless force fields, also happens to provide exactly the mechanism needed to generate an _effective_, _dynamic_ type of mass, which also happens to allow exactly one effective force field to remain massless and long range, while the rest are massive and short-range. And all of this fits _precisely_ with the empirical facts of the physical world we happen to live in. It is truly a mind-blowing situation, all stemming directly from the basic requirement of having force and particle fields interact in a conservative way. Hopefully that provides sufficient motivation to see how this all plays out!

## Coupling between electron and electromagnetic fields

The principle of gauge invariance can be illustrated in the case of the electromagnetic force interaction with the electron represented in a simplified form by a complex wave state $\chi$ obeying the [[complex KG]] wave equation (which is the simplest kind of wave state that can represent a conserved charge value). You should read the first part of that page, and then come back here where directed, when the EM coupling is addressed. You should also read the [[Lagrangian]] page, as that provides the mathematical framework upon which this is based (although we do cover enough to get by if you are impatient).

The analysis starts with a Lagrangian defined as the kinetic energy minus a potential energy for this electron field, where, in [[four-vector]] notation, the kinetic energy is the 2nd order temporal derivative of the wave function (i.e., the contravariant times the covariant):

$$
\mathcal{L} = T - V
$$

$$
\mathcal{L} = \partial_\mu\chi^*\partial^\mu\chi - V(\chi^* \chi)
$$

### Global gauge invariance

The integral form of the Lagrangian action essentially involves subtracting the starting from the ending action values. Therefore, adding a constant global factor $G$ has no effect on the resulting laws of motion. 

{id="eq_const" title="global (constant) gauge transformation"}
$$
\chi \rightarrow \chi' = G \chi
$$

This **global gauge transformation** gives rise to a corresponding "symmetry" or invariance, which really means that something is unchanging or _conserved_. In this case, there is a conserved _current_ (see [[continuity]] for the algebraic derivation), which critically involves multiplying the partial derivative of the field times the complex-conjugate of the field value, at every point:

{id="eq_current" title="conserved current"}
$$
j^\mu = G \chi^* \partial_\mu \chi
$$

where the conserved nature of the current means that its change over time across the whole field is 0:
$$
\partial_\mu j^\mu = 0
$$

In more advanced versions (such as the electroweak case), you can introduce multi-component wave fields, and that allows the introduction of a matrix in the definition of the current, which thereby determines the kind of symmetry group that the system obeys. This is basically where the Abelian vs. non-Abelian distinction arises. For this case, with just a single-component complex field, it corresponds to the simplest $U(1)$ symmetry group.

### Local gauge invariance

So far, we have only defined our electron, and shown that it can have a conserved current. But this current doesn't _do_ anything yet. This is just a free electron with no force field to interact with. The way that the force field is introduced in gauge theory seems very strange and somewhat arbitrary, and most treatments just plow right ahead without properly motivating the steps that come next. The whole point of this next step is to introduce the coupling with the force field in a way that still conserves the global current.

The way you do this is to first introduce a _local_ factor (which ends up being the force field) that varies as a function of position and time, known as a **local gauge transformation**, and then observe that doing so really messes up your conserved current. This is because, unlike the case with the global transformation, this local transformation can change the slope of the wave field at any given point, and it therefore changes the partial derivative that enters into the current factor.

So what are you going to do to fix this mistake that you just made by introducing this force field? Well, you're going to change the definition of the derivative in precisely the way necessary to undo the mistake, and restore your conserved current! This new form of the derivative defines a **minimal coupling** between the new force field and the particle field, and mathematically determines how to update the Lagrangian and the resulting laws of motion to specify this interaction between the two fields.

{id="eq_local" title="local gauge transformation"}
$$
\chi(x) \rightarrow \chi(x)' = G(x) \chi(x)
$$

$$
G(x) = e^{-i \alpha(x)}
$$

where $\alpha(x)$ is an arbitrary function specifying the phase of $G(x)$ at every point -- we will parameterize this with an additional field in a moment.

If you apply the covariant four-vector derivative to this transformed field, as required by the Lagrangian, you end up with an extra term:

$$
\partial_\mu \chi(x) \rightarrow U(x) \partial_\mu \chi(x) + \chi(x) \partial_\mu G(x)
$$

This extra term means that the system is not invariant to this local gauge transformation, as we expected -- no such system would be.

Now, you introduce a new **gauge field** that is going to represent the local transformation parameters, in the form of the electromagnetic vector potential field $A_\mu(x)$

$$
A_\mu(x) \rightarrow A_\mu(x) + \frac{1}{e} \partial_\mu \alpha(x)
$$ 

where $e$ is the **coupling constant** that will parameterize the strength of the interaction between this EM field and the original complex particle field.

Next, we define our own new **covariant derivative** in a way that will make all the math work out under the local gauge transformation (dropping the explicit factors of x from now on):

{id="eq_kgc_covd" title="covariant derivative for a charged scalar"}
$$
D_\mu \chi = \left(\partial_\mu - i \tfrac{e}{\hbar c} A_\mu\right)\chi
$$

$A_\mu$ obeys a wave equation, which is driven by the current generated by $\chi$:

{id="eq_kgc_eom" title="the potential is driven by the current"}
$$
\square\, A_\mu = \mu_0\, j_\mu
$$

And we plug in our covariant derivative into the current definition from above, ending up with this:

$$
j_\mu = -i\tfrac{e}{\hbar}\left[\chi^* D_\mu \chi - (D_\mu \chi)^* \chi\right]
$$

Expanding the covariant derivative splits the current into two pieces:

{id="eq_kgc_split" title="the two halves of the current"}
$$
j_\mu = \underbrace{-i\tfrac{e}{\hbar}\left[\chi^*\partial_\mu\chi - (\partial_\mu\chi^*)\chi\right]}_{\text{convection}}
\;-\; \underbrace{\tfrac{2e^2}{\hbar^2 c}\,|\chi|^2 A_\mu}_{\text{proportional to } A}
$$

The first term is conventionally described as a _convection_ factor, and it corresponds to the original charge density that we started out with before introducing the covariant derivative. It depends on how $\chi$'s phase varies.

The second term is where the force field $A_\mu$ is generating its own contribution to the current -- the feedback factor for the force to the current. Critically, you can move this to the left hand side of the wave equation, and recognize that it operates just like the mass term in the KG equation!

{id="eq_kgc_mass" title="a current that follows A is a mass term"}
$$
\square\, A_\mu + \mu_0 \tfrac{2e^2}{\hbar^2 c}\,|\chi|^2 A_\mu = 0
$$

That is a Klein-Gordon equation for $A_\mu$, with $m^2 \propto |\chi|^2$.

Interestingly, this "dynamic mass" of the EM field is not some kind of crazy unknown exotic property. Instead, this mass expression:

$$
\sqrt{\mu_0 2e^2|\chi|^2/\hbar^2 c}
$$

is known as the **plasma frequency** for a particle field that is filled with a relatively uniform density of charged currents, so that it behaves more like a fluid than discrete particles.

Light below this plasma frequency does not propagate at all, which is why radio waves reflect off the ionosphere and why metals are shiny.

Furthermore, a superconductor represents an intermediate case, where there is a uniform and persistent current, so the EM field acquires a real mass inside it. That mass is the inverse London penetration depth, and its consequence is the Meissner effect. This was one of the phenomena that inspired [[@^Anderson63]] in thinking about how one might solve the central mystery of how to have a massive force field without imposing a mass directly.

## Symmetry groups

* SU(2) is just the 3D rotation space -- spin!  So SU(2) x U(1) is pretty much determined by any spinning thing interacting with a massless radiative field. Not a general principle as much as a necessary fact, given that fermions spin..

* SU(3) is another story..

* Proca is 4-vector (vector boson) version of KG.

