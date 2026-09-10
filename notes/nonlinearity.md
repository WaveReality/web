+++
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

Key point:

* configuration space is not about nonlinear interactions, but rather about *correlations*.

* Montina08 is the key paper proving that exponential config space is necessary to represent the correlations, even allowing for nonlinear interactions! The key assumption is Markovian-ness. Need to understand this better!!  But they definitely say that it could be removed.  Stochastic motion + non-local wave influences seems like it would really change things.

* 3N dimensions -> $R^{3N}$ numbers is key about exponential explosion.

## Claude chat

### Nonlinear SM vs linear Hilbert space

> Prompt: "the nonlinear interaction terms in the weak and strong forces of the standard model make the linear hilbert space model inaccurate"

Nonlinearity in the field equations and linearity of the state space are two different things, and the Standard Model has the first without giving up the second.

#### Where the nonlinearity actually lives

In Yang–Mills theory the field strength is $F_{\mu\nu} = \partial_\mu A_\nu - \partial_\nu A_\mu + ig[A_\mu, A_\nu]$, so $\text{tr}\,F^2$ contains cubic and quartic terms in $A$. The classical equations of motion are genuinely nonlinear PDEs — gluons couple to gluons, W's couple to W's. That's why QCD confines and why non-abelian theories have instantons, sphalerons, and a nonperturbative vacuum structure that abelian theories don't.

But that nonlinearity is in the *operator* equations, i.e. in the Heisenberg-picture field equations and equivalently in the Hamiltonian's functional dependence on the field operators. It is not nonlinearity in the map that evolves states.

#### Why the state space stays linear

Evolution is still:

$$
|\psi(t)\rangle = e^{-iHt}|\psi(0)\rangle
$$

and $e^{-iHt}$ is a linear operator no matter how high the powers of $\hat A$ inside $H$ are. Superposition holds exactly: if $|\psi_1\rangle$ and $|\psi_2\rangle$ each evolve some way, $\alpha|\psi_1\rangle + \beta|\psi_2\rangle$ evolves to $\alpha U|\psi_1\rangle + \beta U|\psi_2\rangle$, with no cross terms.

The cleanest illustration is the anharmonic oscillator, $H = p^2/2m + \lambda x^4$. Classical trajectories are nonlinear and chaotic-adjacent; the quantum evolution is exactly linear and unitary. "Interacting" and "nonlinear dynamics on Hilbert space" are simply not the same predicate. Feynman diagrams beyond tree level are precisely the perturbative bookkeeping for a nonlinear $H$ acting linearly on states.

Worth adding: fundamental nonlinearity in state evolution is not a free parameter you can quietly turn on. Weinberg-type nonlinear modifications of QM generically allow superluminal signaling and are experimentally bounded to extraordinary precision (NMR and atomic-clock tests put limits many orders of magnitude below typical energy scales). Nonlinear equations that *do* appear in practice — Gross–Pitaevskii, Hartree–Fock, semiclassical gravity's $G_{\mu\nu} = 8\pi\langle T_{\mu\nu}\rangle$ — are mean-field *approximations* to an underlying linear theory, and their nonlinearity is an artifact of the truncation.

#### The real problems with the naive Hilbert-space picture

There are genuine reasons to distrust "the Standard Model = a vector in a Hilbert space," but they're different from the one you named:

- **Haag's theorem.** The interacting field theory is not unitarily equivalent to the free Fock-space representation. The Fock space you build from free particles is *the wrong space* for the interacting theory, which is one reason canonical perturbation theory is formally ill-defined and needs renormalization to make sense.
- **No rigorous construction.** Nobody has constructed the Hilbert space of 4D Yang–Mills with a mass gap — that's a Clay Millennium Problem. In that sense the Standard Model's Hilbert space is a physicist's assumption, not a theorem, and the self-interaction is exactly why.
- **Indefinite metric.** Covariant gauge quantization works in a space with negative-norm states; the physical Hilbert space appears only as BRST cohomology. So the "Hilbert space" of perturbative gauge theory isn't a Hilbert space until you quotient.
- **Confinement.** Asymptotic states aren't quark and gluon states at all, so the natural particle basis doesn't exist non-perturbatively.
- **Infrared structure.** In gauge theories with massless fields, soft radiation makes the naive S-matrix elements vanish; you need coherent states or dressed asymptotics, and the vacuum has a large degeneracy labeled by asymptotic charges.

The algebraic QFT response to all of this is to take the net of local operator algebras as primary and let Hilbert space representations be derived objects (via GNS), since inequivalent representations proliferate. That's a real demotion of the Hilbert space — but the algebras are still linear, and the states are still linear functionals on them.

So: if the claim is "self-interaction breaks linearity of quantum mechanics," no. If it's "the textbook Fock-space picture is not a valid description of the interacting Standard Model," that's correct and mathematically serious — just for reasons of inequivalent representations and non-perturbative construction rather than nonlinearity per se.

### Exponential configuration space

> Prompt: "does the use of exponential configuration space in the Hilbert space capture the nonlinear interaction terms?"

Yes, in a precise and rather satisfying mechanical sense — and no in the sense that matters for constructing the Standard Model. Both halves are worth separating.

#### The sense in which it does

The exponential functor:

$$
\Gamma(\mathfrak{h}) = e^{\mathfrak{h}} = \bigoplus_n \mathrm{Sym}^n \mathfrak{h}$ (or $\bigoplus_n \Lambda^n\mathfrak{h}
$$

for fermions) is exactly the structure that lets a *polynomial* — hence nonlinear — expression in the fields act as a *linear* operator. The trick is the grading. Fields are linear in $a, a^\dagger$, so a cubic vertex is something like $a^\dagger a^\dagger a$, which is a perfectly linear map that happens to send $\mathrm{Sym}^n \to \mathrm{Sym}^{n+1}$. The three-gluon vertex is literally the statement that a linear operator on the exponential space connects the 1-gluon sector to the 2-gluon sector. Degree-$k$ nonlinearity becomes an operator hopping $k$ grades. Nothing about linearity is strained; the nonlinearity is re-encoded as particle-number non-conservation.

The two structural properties of $\Gamma$ do real work here:

$$
\Gamma(\mathfrak{h}_1 \oplus \mathfrak{h}_2) \cong \Gamma(\mathfrak{h}_1)\otimes\Gamma(\mathfrak{h}_2), \qquad \langle \varepsilon(f), \varepsilon(g)\rangle = e^{\langle f,g\rangle}
$$

The first turns addition of modes into tensoring of state spaces, which is why locality and mode decomposition cohere. The second is the exponential-vector (coherent state) structure — and this is where your intuition is genuinely vindicated: the nonlinear classical Yang–Mills equations are recovered as the leading-order evolution of coherent states under the linear quantum dynamics. The exponential vectors are precisely the bridge on which linear Hilbert-space evolution projects down to nonlinear classical field equations. So the exponential structure doesn't hide the nonlinearity; it's the place where you can see where it went.

There's a dual picture worth knowing. Via the Wiener–Itô–Segal isomorphism, $\Gamma(\mathfrak{h}) \cong L^2(\mathcal{S}', d\mu_0)$ with $\mu_0$ Gaussian, and the $n$-particle sectors are the Wick/Hermite polynomials. In that language, adding a nonlinear interaction means tilting the measure: $d\mu \propto e^{-\int \lambda\varphi^4} d\mu_0$. The nonlinearity is absorbed into the *measure*, not into the linearity of the space. You still get a Hilbert space $L^2(d\mu)$; it's just a different one.

#### The sense in which it does not

That last sentence is the whole problem. $\Gamma$ is intrinsically a *quasi-free* construction — it's the second quantization of a one-particle space equipped with a Gaussian state. The interacting vacuum is not Gaussian, and the tilted measure is generically not absolutely continuous with respect to $\mu_0$, so $L^2(d\mu) \not\cong \Gamma(\mathfrak{h})$ unitarily. That *is* Haag's theorem, seen from the measure-theoretic side.

Concretely, how far the exponential space gets you is dimension-dependent:

- **$d = 2, 3$:** Wick-ordered $:\!\varphi^4\!:$ is a densely defined self-adjoint operator on Fock space, and Glimm–Jaffe constructed the theory. The exponential space really does carry the nonlinear interaction.
- **$d = 4$:** it does not. $:\!\varphi^4\!:$ isn't a well-defined operator on Fock space, which is why you need renormalization rather than just Wick ordering — and $\varphi^4_4$ is now known to be trivial (Aizenman–Duminil-Copin), so the target theory isn't there to be constructed anyway.
- **Yang–Mills in $d=4$:** the object you'd want to build the exponential space *on* — the physical one-particle space — doesn't exist perturbatively. Confinement means asymptotic states are hadrons, not the gluons whose self-coupling generated the nonlinearity in the first place. The Fock space of gluons is the wrong exponential space by construction.

There's a partial rescue: extend $\Gamma(\mathfrak{h})$ to a Gelfand triple (white-noise / Hida distribution analysis), and Wick powers $:\!\varphi^n\!:$ become continuous maps from test vectors into distribution vectors. This is genuinely useful, but note what it concedes — the interaction is then a *sesquilinear form* on a dense subspace rather than an operator on the Hilbert space. The exponential space hosts the nonlinearity only in a weakened sense.

#### Where that leaves it

The exponential structure is what makes nonlinear interaction *representable* by linear operators at all, and it's also where the classical nonlinear equations reappear in the $\hbar \to 0$ limit. What it doesn't do is guarantee that the interacting theory lives in that particular exponential space. Haag–Ruelle scattering theory says the correct picture is: the interacting Hilbert space is Fock-like *asymptotically*, built on physical dressed particles, given a mass gap and asymptotic completeness — while at finite times and short distances it isn't an exponential space of the Lagrangian's fields.

So: right structure for the kinematics of nonlinearity, wrong space for the dynamics. The failure mode is representation-theoretic, not a failure of linearity.

### Feynman diagrams as conditioning

> Prompt: "does the use of Feynman diagrams to configure the terms in an s-matrix integral constitute a kind of nonlinear operation, in that each such diagram holds for a specific type of interactions, and so you are already conditioning on energy levels and particle types in solving for a specific type of interaction? This is in comparison to a hypothetical alternative situation where the exact same set of equations is used for all possible calculations of the standard model.

No — extracting a specific process is a linear operation, and in a fairly strict sense. But your hypothetical alternative *exists*, it's the more fundamental formulation, and it is nonlinear — just not in the state space. That contrast is the interesting part.

#### Why picking a process isn't a nonlinear step

$S$ is one fixed unitary operator on one fixed space. Computing $e^+e^- \to \mu^+\mu^-$ means evaluating $\langle f|S|i\rangle$ for particular $|i\rangle, |f\rangle$. Reading off a matrix element of a fixed linear operator in a chosen basis is the definition of a linear operation; it's no more nonlinear than writing $M_{23}$ for a matrix. Choosing external legs conditions the *input*, not the map.

Two further points cut against the "conditioning" reading:

**The diagrams aren't chosen, they're enumerated.** The Feynman rules are derived once from $\mathcal{L}$ — vertices, propagators, symmetry factors, all fixed before any process is contemplated. Given external legs and an order in the coupling, the diagram set is generated algorithmically; that's literally what FeynArts and MadGraph do. No physics judgment enters the enumeration. Wick's theorem does the work, and diagrams are its combinatorial bookkeeping.

**Processes are not independent of each other.** Unitarity of the single operator $S$ ties them together: the optical theorem and the Cutkosky cutting rules mean the imaginary part of a loop amplitude in one process is fixed by tree amplitudes of *other* processes. If each diagram were a separate conditioned calculation, that couldn't hold. Crossing symmetry does the same thing — one analytic function covers $s$-, $t$-, and $u$-channel processes in different kinematic regions. The interlocking is a signature of linearity, not of case-by-case treatment.

#### Your hypothetical alternative, and its actual nonlinearity

There genuinely is a "same equations for everything" formulation. Three equivalent faces of it:

- **The generating functional** $Z[J] = \int \mathcal{D}\phi\, e^{iS[\phi] + i\int J\phi}$. Every correlator, every process, is a functional derivative of this one object.
- **The effective action** $\Gamma[\phi]$, whose stationarity condition $\delta\Gamma/\delta\phi = 0$ is the *exact* quantum equation of motion — the direct quantum successor to the nonlinear classical Yang–Mills equation.
- **The Dyson–Schwinger tower**, e.g. schematically $S^{-1} = S_0^{-1} - \Sigma[S, D, \Gamma_3]$, and the functional RG's Wetterich equation $\partial_k \Gamma_k = \tfrac12 \mathrm{Tr}\big[(\Gamma_k^{(2)} + R_k)^{-1}\partial_k R_k\big]$ — a single exact equation.

And these are **nonlinear** — the propagator appears inside its own self-energy, the vertex depends on the propagators which depend on the vertex, and the Wetterich equation has $\Gamma_k$ inverted on the right-hand side. So the answer to your question inverts nicely: the universal, non-conditioned formulation is the nonlinear one, and Feynman diagrams are what you get by *iteratively solving that nonlinear system order by order*. Diagram topologies are the iteration terms. Perturbation theory doesn't introduce nonlinearity by fragmenting a linear problem — it removes nonlinearity by linearizing a nonlinear functional problem one order at a time.

That's also why the nonlinear form is not redundant: dynamical mass generation and chiral symmetry breaking are solutions of the DSE tower that no finite set of diagrams produces, because they're non-analytic in the coupling. The self-referential structure is doing something the iterated expansion cannot.

#### What your instinct is correctly detecting

The regime-dependence you're pointing at is real; it's just not nonlinearity. Where genuine case-by-case choice enters:

- **Scale choice.** Exact amplitudes are $\mu$-independent; truncated ones are not, so you set $\mu \sim Q$ to avoid large logs. This is conditioning on the energy regime, and it's unavoidable at finite order. Probably the strongest version of your point.
- **Different Lagrangians in different regimes.** Chiral perturbation theory, HQET, SCET, NRQCD, SMEFT — genuinely different degrees of freedom and expansion parameters. There is no single *practical* calculational framework spanning all scales of the Standard Model.
- **The series doesn't converge.** Perturbative expansions are asymptotic, with $n!$ growth, renormalons, and Dyson's instability argument for QED. So the "universal formula" $\langle 0|T e^{-i\int H_I}|0\rangle$ is formal, and extracting numbers requires resummation choices made per observable.
- **Non-perturbative sectors.** Instanton effects go as $e^{-1/g^2}$ and are invisible at every order; hadronic matrix elements need lattice QCD.
- **Asymptotic states.** You can't compute an S-matrix element for quarks and gluons at all — only IR-safe observables, with jet definitions and factorization theorems that carry process-dependent structure.

The accurate word for all of this is **non-uniform** or **scheme-dependent**, not nonlinear. The distinction matters because it localizes the difficulty: nonlinearity would break superposition and let you signal superluminally; non-uniformity of approximation schemes breaks none of the theory's structure, it just means we lack a single convergent algorithm.

So the layered picture: linear on Hilbert space, nonlinear in correlator/coupling space, non-uniform in calculational practice. Three different levels, and the nonlinearity you keep (correctly) sensing lives in the middle one.

> Prompt: "so is the nonlinearity in correlator space physically meaningful?"

Yes — and by the strictest test available: it changes which solutions exist, and it produces numbers that have been measured.

#### Bifurcation, and most of the mass of ordinary matter

The quark gap equation is a nonlinear fixed-point problem — the propagator appears inside its own self-energy. Below a critical coupling, the only solution is $m = 0$. Above it, a second branch appears with $m \neq 0$. That's a bifurcation, and it is the origin of dynamical chiral symmetry breaking: $\langle \bar q q\rangle \approx -(250\,\text{MeV})^3$, a constituent quark mass of a few hundred MeV from nearly massless inputs, and roughly 99% of the proton's mass having nothing to do with Higgs Yukawa couplings. A linear equation has no mechanism for "a qualitatively new solution appears past a threshold." This one does, and the answer is on your bathroom scale.

#### Dimensional transmutation

$\beta(g) = -b_0 g^3 + b_1 g^5 + \dots$ is a nonlinear ODE, and integrating it converts a dimensionless coupling into a mass scale: $\Lambda_{\text{QCD}} \approx 200\text{–}300$ MeV, generated from a theory whose Lagrangian contains no scale at all. Asymptotic freedom is the sign of the leading nonlinear term, and it traces directly back to the gauge self-interaction you started this conversation with — the cubic and quartic gluon vertices are what make $b_0 > 0$. The resulting curve, $\alpha_s(M_Z) \approx 0.118$ running to $\approx 0.3$ near 2 GeV, has been checked across more than two orders of magnitude in energy. You are looking at the experimentally verified solution of a nonlinear equation.

#### Fixed points, and exponents measured to six digits

RG flow *is* a nonlinear dynamical system on the space of couplings, and the whole vocabulary transfers: fixed points, basins of attraction, relevant and irrelevant directions as eigenvalues of the linearization about the nonlinear flow. Universality — why wildly different microscopic systems share critical exponents — is the statement that they flow to the same fixed point. Those exponents are measured in condensed matter to several significant figures. The conformal bootstrap pins the 3D Ising exponents to ~6 digits using *only* nonlinear consistency conditions (crossing, unitarity, analyticity), with no Lagrangian at all. Hard to call that formal bookkeeping.

Relatedly, unitarity itself is a nonlinear constraint: $S^\dagger S = 1$ is quadratic in amplitudes. The optical theorem, $\text{Im}\,\mathcal{M}(s) \propto \sigma_{\text{tot}}$, is a nonlinear relation among correlators that you can read off a detector. Dispersive positivity bounds built from it are currently used to carve out allowed regions of SMEFT Wilson coefficients — nonlinearity in correlator space constraining what new physics is possible.

And a live example: whether the electroweak vacuum is stable or metastable is entirely a question about nonlinear flow — whether the Higgs quartic $\lambda$ dips below zero around $10^{10}$–$10^{11}$ GeV. Current measurements put us near the metastable boundary. That's a physically consequential answer determined by the nonlinear structure and nothing else.

#### The honest caveats

Not every nonlinear equation in the tower carries physical content. Individual off-shell Green's functions — the gluon propagator most notoriously — are gauge-dependent, and $\Gamma[\phi]$ shifts under field redefinitions. $\beta$ beyond two loops is scheme-dependent. So the correct claim is that the *nonlinear structure* is physical, while any particular nonlinear equation may be partly convention. The invariant content shows up in the places listed above: existence and location of bifurcations, fixed points and their exponents, generated scales, and relations among on-shell observables.

#### The frame worth keeping

Nonlinearity is conserved under quantization, not eliminated. It gets relocated — out of the state space, where it would break superposition and permit superluminal signaling, and into the space of correlation functions and couplings, where it instead generates mass scales, drives symmetry breaking, and sets phase structure. Your original intuition that the Standard Model's self-interactions must have some non-linear consequence for the formalism was right. It just lands one level up from where you first put it, and it earns its keep there.

> Prompt: "So what is the implication of all this for the physical reality of the configuration space that is used in the Hilbert state space -- if it is only there to allow a linear framework to capture the fundamental nonlinearity of the underlying standard model, then perhaps there could be an alternative mathematical formulation that does not require this exponentially-large configuration space?"

Two claims are bundled in your question, and they have opposite answers. The *linearity* looks close to non-negotiable. The *exponential size* is very likely not physically real — but for reasons that have nothing to do with the self-interaction, and none of the existing alternatives eliminate what the space encodes. Let me also flag a premise worth adjusting: the exponential structure isn't there to accommodate nonlinearity. Free QED needs a Hilbert space, and so does a single spin. The exponential grading comes from multiparticle statistics, and the nonlinearity — as we traced — didn't get absorbed by inflating the state space, it got relocated to correlator space. So the space isn't scaffolding erected for the interactions.

#### QFT already demotes configuration space

The "physical reality of the 3N-dimensional configuration space" debate is a non-relativistic quantum mechanics problem. In QFT the primitives are operator-valued fields attached to regions of 4D spacetime; particle number is not fundamental, and the Fock/configuration-space description is a derived, representation-dependent artifact — as Haag's theorem already forces on us. Local algebras in QFT are type III$_1$ von Neumann factors, which means there is no tensor factorization into "here" and "there," no local number operator, and no well-posed question "how many particles are in this box." The exponentially large configuration space has already lost its claim to being the furniture of the world, independently of anything you or I might prefer.

#### Evidence the size is genuinely redundant

- **Holographic bounds.** The Bekenstein–Bousso bound says the degrees of freedom in a region scale with its bounding *area*, not its volume. Quantum field theory's Hilbert space is vastly larger than that — it overcounts, and gravity knows it. This is the strongest physical argument that the exponential space is not real.
- **Physical states are exponentially thin.** Ground states of local Hamiltonians obey area-law entanglement and are parameterized by polynomially many numbers; that's why DMRG and tensor networks work. And states reachable by polynomial-time evolution form a measure-zero subset of Hilbert space. Almost the entire space is never visited by anything.
- **Algebra type is not fixed.** Recent work (Leutheusser–Liu, Witten, and the Type II$_\infty$ story in black hole thermodynamics) has the algebra type itself changing with $G_N$ and with the observer's description — Hilbert space looking emergent and approximate rather than given.

#### Formulations that don't build it

These aren't hypothetical; several are how frontier calculations are actually done.

- **Algebraic QFT.** Net of local algebras primary, state as a functional on the algebra, Hilbert space derived via GNS when you want it. This is the mainstream answer to "the Hilbert space isn't fundamental."
- **Euclidean functional integral / lattice.** A measure on field configurations, correlators as the data, Hilbert space *reconstructed* afterwards via Osterwalder–Schrader if you care. Lattice QCD computes the proton mass to a few percent without a Hilbert space appearing anywhere in the algorithm.
- **On-shell amplitude methods.** BCFW recursion, generalized unitarity, the amplituhedron: observables from locality and unitarity as *outputs* of positive geometry rather than inputs, with no Lagrangian and no state space. The honest limitation is that this is most developed for planar $\mathcal{N}=4$ and gravity, computes scattering rather than general finite-time dynamics, and is not yet a formulation of the Standard Model.
- **Bootstrap.** Spectrum and exponents from symmetry, unitarity, and crossing alone.
- **Dyson–Schwinger / functional RG** — the nonlinear formulation taken as primary, which is exactly the thing you were reaching for.

#### What can't be dropped, and the price

Every one of those relocates the content rather than removing it — into an algebra's noncommutativity, into a measure, into positive geometry, into entanglement structure. And three constraints bite hard on any replacement:

1. **Tsirelson's bound.** Bell correlations exceed local-realistic limits by exactly the amount Hilbert space predicts. Reproduce that and you've reproduced the noncommutative tensor structure.
2. **Quantum computational advantage.** This is the sharpest one. If the exponential space were mere formal scaffolding compressible to polynomially many real parameters in general, Shor's algorithm shouldn't work. The dimension appears to be doing physical labor.
3. **Reconstruction theorems.** Gleason, and the operational axiomatizations (Hardy; Chiribella–D'Ariano–Perinotti), *derive* the Hilbert-space formalism — tensor products included — from a handful of operational requirements. The formalism isn't a modeling choice someone made; it's forced by the probability structure.

#### My read

The exponentially large configuration space is best understood as a redundant but very efficient chart, not as the manifold. It's over-large by the holographic bound's measure, most of it is unreachable, and QFT already treats it as derived. So "not fundamental" is a defensible and I'd say probably correct position, and quantum gravity is where the seams show.

But the direction of the eventual replacement almost certainly isn't "nonlinear instead of linear." The candidates on the table — algebras, geometry, entanglement, measures — are all ways of keeping exact linearity and superposition while dropping the assumption that a single big vector space is the primary object. Your instinct that something is over-built is, I think, right. The over-built thing is the *dimensionality*, and the culprit is gravity's holographic bound, not the weak and strong forces.

### Lit

> Prompt: "what are some papers that discuss this topic most directly?"

Verified against arXiv/journal listings. Organized by the threads we went through.

#### Is the exponentially large space physically real?

The single most direct paper for your question:

- **Poulin, Qarry, Somma, Verstraete, "Quantum simulation of time-dependent Hamiltonians and the convenient illusion of Hilbert space"** — [arXiv:1102.1360](https://arxiv.org/abs/1102.1360), PRL **106**, 170501 (2011). Proves that states reachable by *any* time-dependent local Hamiltonian in polynomial time occupy an exponentially small volume of Hilbert space. The title is your thesis, stated by people who then made it quantitative.
- **Carroll, "Reality as a Vector in Hilbert Space"** — [arXiv:2103.09780](https://arxiv.org/abs/2103.09780). The strongest defense of the opposite position: the vector *is* the ontology, and space and fields are emergent. Worth reading precisely because it's the view you're pushing against, argued carefully.
- **Giddings, "Hilbert space structure in quantum gravity: an algebraic perspective"** — [arXiv:1503.08207](https://arxiv.org/abs/1503.08207). Argues that factorization is already problematic in ordinary QFT, and that the algebra of observables should be primary.

#### Hilbert space as derived rather than fundamental

- **Witten, "Notes on Some Entanglement Properties of Quantum Field Theory"** — [arXiv:1803.04993](https://arxiv.org/abs/1803.04993), Rev. Mod. Phys. **90**, 045003. The best entry point to type III$_1$ algebras, Reeh–Schlieder, and why "how many particles are in this region" is not a well-posed question. If you read one thing on the algebraic demotion of Hilbert space, read this.
- **Leutheusser & Liu** — [arXiv:2110.05497](https://arxiv.org/abs/2110.05497) and [arXiv:2112.12156](https://ui.adsabs.harvard.edu/abs/2023PhRvD.108h6020L/abstract) (both PRD 2023). Type III$_1$ structure *emerging* from a type I boundary algebra in the large-$N$ limit — algebra type as a derived, limit-dependent property.
- **Witten, "Gravity and the Crossed Product"** — [arXiv:2112.12828](https://arxiv.org/abs/2112.12828), JHEP 10 (2022) 008, and **Chandrasekaran, Penington, Witten, "Large $N$ Algebras and Generalized Entropy"** — [arXiv:2209.10454](https://arxiv.org/abs/2209.10454), JHEP 04 (2023) 009. Where the algebra type changes with $G_N$ and entropy becomes a property of the algebra rather than of a state in a fixed space.

#### Why the configuration/Fock space is the wrong space for interacting fields

- **Earman & Fraser, "Haag's Theorem and its Implications for the Foundations of Quantum Field Theory"** — Erkenntnis **64** (2006) 305, [free preprint](https://philsci-archive.pitt.edu/2673/). Careful, readable, and exactly on the point that the interacting theory doesn't live in the exponential space you build for the free one.
- **Miller, "Haag's Theorem, Apparent Inconsistency, and the Empirical Adequacy of Quantum Field Theory"** — [BJPS **69** (2018)](https://www.journals.uchicago.edu/doi/abs/10.1093/bjps/axw029). The rebuttal: why the theorem doesn't undermine QFT's empirical success.

#### The nonlinearity that is physically real

- **Roberts & Williams, "Dyson-Schwinger equations and their application to hadronic physics"** — [arXiv:hep-ph/9403224](https://arxiv.org/abs/hep-ph/9403224), Prog. Part. Nucl. Phys. **33** (1994) 477. The standard review; the gap-equation bifurcation is worked out explicitly.
- **Alkofer & von Smekal, "The infrared behaviour of QCD Green's functions"** — [Phys. Rept. **353** (2001) 281](https://ui.adsabs.harvard.edu/abs/2001PhR...353..281A/abstract). Confinement, dynamical symmetry breaking, *and* the positivity violation of transverse gluons — i.e. exactly where the naive Hilbert-space interpretation of the fields fails.
- **Berges, Tetradis, Wetterich, "Non-Perturbative Renormalization Flow in QFT and Statistical Physics"** — [arXiv:hep-ph/0005122](https://arxiv.org/abs/hep-ph/0005122), Phys. Rept. **363** (2002) 223. The exact nonlinear functional flow equation, treated as the primary object.
- **Roberts et al., "Empirical Consequences of Emergent Mass"** — [arXiv:2009.04011](https://arxiv.org/pdf/2009.04011). Ties the nonlinear structure to measured hadron properties.

#### Why you can't simply make the state evolution nonlinear

- **Weinberg**, Annals Phys. **194** (1989) 336 and PRL **62**, 485 (1989) — the nonlinear generalization, proposed as a foil for precision tests.
- **Gisin**, Phys. Lett. A **143** (1990) 1 — superluminal signalling.
- **Polchinski**, PRL **66**, 397 (1991) — either EPR signalling or signalling between branches. See also the [later comment](https://arxiv.org/abs/quant-ph/0012041).
- **Bollinger et al.** — the $^9$Be$^+$ hyperfine test, bounding nonlinear corrections at $4\times10^{-27}$ of the binding energy per nucleon (PRL **63**, 1031, 1989 — journal page from memory, the experiment and bound are confirmed).
- **Kaplan & Rajendran, "Causal framework for nonlinear quantum mechanics"** — [PRD **105**, 055002](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.105.055002). The modern attempt to have nonlinearity without signalling; useful for seeing how expensive it is.

#### What forces Hilbert space in the first place

- **Hardy, "Quantum Theory From Five Reasonable Axioms"** — [arXiv:quant-ph/0101012](https://arxiv.org/abs/quant-ph/0101012).
- **Chiribella, D'Ariano, Perinotti, "Informational derivation of quantum theory"** — [arXiv:1011.6451](https://arxiv.org/abs/1011.6451), PRA **84**, 012311. Derives the formalism, tensor products included, from operational principles. Relevant because it shows the space isn't a modeling choice you can trade away for free.

#### Formalisms that never build the state space

- **Arkani-Hamed & Trnka, "The Amplituhedron"** — [arXiv:1312.2007](https://arxiv.org/abs/1312.2007), JHEP 10 (2014) 030. Locality and unitarity as *outputs* rather than inputs; the opening paragraphs make your argument better than most philosophy papers do.
- **El-Showk et al., "Solving the 3D Ising Model with the Conformal Bootstrap"** — [arXiv:1203.6064](https://arxiv.org/abs/1203.6064), and the review **Poland, Rychkov, Vichi** — [arXiv:1805.04405](https://arxiv.org/abs/1805.04405), Rev. Mod. Phys. **91**, 015002.

#### The configuration-space-realism debate itself

- **Wallace & Timpson, "Quantum Mechanics on Spacetime I: Spacetime State Realism"** — [BJPS **61** (2010) 697](https://philsci-archive.pitt.edu/4621/1/ststaterealism.pdf). Takes states-of-regions as fundamental instead of a vector on configuration space — the philosophical counterpart of the algebraic move.
- **Swanson, "How to Be a Relativistic Spacetime State Realist"** — [PDF](https://personal.lse.ac.uk/robert49/PPB/pdf/swanson2020str.pdf). Fixes the relativistic problems in the above using algebraic QFT.

If you want three to start: Poulin et al. for the quantitative version of your intuition, Witten's entanglement notes for why the field-theoretic Hilbert space is already not what the textbooks suggest, and Earman & Fraser for the cleanest statement of why the exponential space fails specifically once interactions are on.

**Sources:** all identifiers above were checked against arXiv and journal listings during this response, except the two explicitly flagged as recalled (Alkofer–von Smekal's arXiv number, which I've replaced with the journal reference, and the Bollinger PRL page number).

### More explicit statement of the question

Prompt: "Just to reiterate a bit on the discussion topic: I'm trying to understand the relationship between two different levels of description of the same phenomena. At the abstract level, the standard usage of the Hilbert space and Schrodinger equation in the 'old / nonrelativistic' version of quantum mechanics, which is the basis for most of the debate about the physical interpretation of quantum mechanics, models interactions among particles using the linear time-evolution of the Schrodinger equation operating over configuration space. By contrast, at the detailed level, you have the full complexity of the SM. So, the question is: is the linear Schrodinger equation somehow able to capture all the nonlinearities of the SM because it can use the extra degrees of freedom of the configuration space? In other words, the exponential size of the configuration space allows you to essentially have unique state terms for the interactions, which allow it to capture any kind of interaction because you have distinct degrees of freedom available in the state space to do so, even though you're just using a linear function (Schrodinger) to evolve the state over time. But, again to contrast with the SM: this more detailed picture of the 'actual physics' shows that the interactions are actually nonlinear, and this nonlinearity is mathematically challenging, but nevertheless, if you could actually solve the nonlinear math directly, you would not need the excessive (largely unused) extra degrees of freedom of a configuration space. Therefore, if all of this is true, then the configuration space is really just an artifact of using the linear Schrodinger equation to try to capture something that is actually fundamentally nonlinear. And therefore, everyone who gets upset about the exponential size of the configuration space is kind of missing the point. The 'actual nonlinear physics' doesn't actually need configuration space in the first place.  Does the structure of this argument make sense? Has someone articulated this point clearly in the literature?"

Yes, the structure makes sense — it's stated precisely enough to be tested, which most versions of this intuition aren't. And it fails, at one identifiable premise, with a theorem aimed almost directly at the counterfactual. Let me separate the two.

#### The load-bearing premise

The claim doing the work is: *the exponential size of the state space is what allows a linear operator to represent nonlinear interactions.*

Take N free, non-interacting spin-½ particles. The Hamiltonian is $H = \sum_i h_i$ — no interaction terms at all, nothing nonlinear anywhere, not even a two-body potential. The state space is still $2^N$-dimensional, not $2N$-dimensional. And the singlet state lives in there, violating Bell inequalities.

So the exponential space is fully present when there is nothing whatsoever to linearize. That's a direct counterexample: the exponentiality cannot be an artifact of linearizing interactions, because it survives the removal of all interactions.

The correct diagnosis is that the exponential size comes from the **tensor product structure** — the fact that joint states are not determined by their marginals. It is about *correlation*, not interaction. Interactions are what *generate* entanglement dynamically, but they don't create the space; the space is there in the kinematics, before any dynamics.

#### A second, independent problem with the same premise

Your picture implies something like: distinct interaction types get distinct degrees of freedom, so the state space acts as a lookup table that a linear map can route through. But look at what the Hamiltonian actually is. For N particles with two-body interactions, $H = \sum_i h_i + \sum_{i<j} V_{ij}$ — **polynomially many terms, each acting on at most two particles.** The interaction is specified by poly-many numbers. It's a sparse, local operator on an exponentially large space.

If the space were providing room to encode interactions, $H$ would need exponentially many parameters. It doesn't. The description of the interaction is compact; it's the *state* that's exponential. Those are different objects, and only one of them is big.

#### The counterfactual is a theorem, and it goes the other way

Your C2 — "if you could solve the nonlinear math directly, you wouldn't need the extra degrees of freedom" — is precisely the question **Montina, "Exponential complexity and ontological theories of quantum mechanics"** ([arXiv:0711.4770](https://arxiv.org/abs/0711.4770), PRA **77**, 022104) settles. He proves that any ontological Markovian theory reproducing quantum dynamics needs at least $2N-2$ continuous variables, where $N$ is the Hilbert space dimension — so exponentially many in system size.

Crucially, **the theorem doesn't care whether the underlying dynamics is linear.** You are allowed your nonlinear equations. You still need exponentially many variables. The exponential resource is not scaffolding erected by the linear formalism; it's a lower bound on *any* description that reproduces the phenomena. The precursor is **Hardy, "Quantum ontological excess baggage"** (Stud. Hist. Phil. Mod. Phys. **35** (2004) 267), and the complementary result is **Pusey, Barrett, Rudolph, "On the reality of the quantum state"** ([Nature Physics **8**, 475](https://www.nature.com/articles/nphys2309)), which blocks the escape of treating the big state as mere ignorance about a small one.

Honest caveat: Montina's theorem assumes Markovian dynamics. Non-Markovian ontological models can evade the bound — but they buy it by paying in memory, which is not obviously cheaper, and it's a narrow door rather than a refutation.

And separately, empirically: any description with a small, local state is a local hidden variable theory. Classical nonlinear field theory is exactly that. Loophole-free Bell experiments closed that in 2015.

#### The level-mixing

There's also a problem in how the two levels are related. Nonrelativistic QM's configuration space is not capturing the SM's nonlinearity — it's the fixed-particle-number sector, with interactions inserted by hand as potentials $V(x_i - x_j)$ that are the low-energy residue of field-theoretic exchange. The gauge self-coupling that generates the SM's nonlinearity doesn't survive to that level at all. Meanwhile in QFT proper, the state space is Fock space and — as we traced — the nonlinearity went into *correlator* space, not into the state space's dimensionality. So the two levels don't stand in the relationship the argument needs.

#### What survives, and who has said it

Your conclusion is partly right, by a different route. The claim that the interpretive debate over-indexes on a 3N-dimensional configuration space that is an artifact — **that** is a live and well-defended position:

- **Myrvold, "What is a wavefunction?"** Synthese **192** (2015) 3247. Argues configuration space and the wavefunction can be real but not *fundamental*, with careful treatment of the relativistic case.
- **Wallace, "Against wavefunction realism"** ([PDF](https://files01.core.ac.uk/download/pdf/195905612.pdf), in *Current Controversies in Philosophy of Science*). Configuration space is one representation among infinitely many; nothing privileges it.
- **Wallace & Timpson**, spacetime state realism (cited earlier).

But note what they say the artifact is an artifact *of*: the nonrelativistic approximation — fixed particle number, a privileged position basis — not of linearization. Same conclusion, different and defensible mechanism.

As for the nonlinearity version specifically: I don't know of anyone who has argued it, and I think that's because it doesn't survive the free-particle counterexample. The nearest neighbors are 't Hooft's cellular-automaton interpretation and Wetterich's "quantum mechanics from classical statistics" — both trying to recover quantum linearity as emergent from a smaller underlying description — and both run directly into Bell, Montina, and PBR, which is where the arguments about them are actually fought.

#### The version I'd defend

The exponential state space is not an artifact of linearizing nonlinear interactions; it's the cost of reproducing quantum correlations, and it's provably not compressible by switching to nonlinear underlying dynamics. But the *particular* exponential space of nonrelativistic configuration-space QM — the one the interpretation literature fixates on — genuinely is an artifact of the nonrelativistic limit, and QFT already dissolves it. And the space is over-large for an entirely separate reason we discussed: the holographic bound. Three distinct claims, one of which is yours and two of which give you most of what you wanted.

They don't contradict each other, but your diagnosis of where the load sits is right — and better than that, **Montina states the assumption explicitly and then concedes your objection himself**, three years before Poulin et al. I pulled both papers to check.

### Poulin et al vs. Montina

the paper by Poulin, Quarry, Somma & Verstraete (2011) appears to contradict the one by Montina (2008), where the former says that configuration space is exponentially too large, and Montina says that this exponential space is necessary under the Markovian assumptions. It appears that Poulin et al make use of an assumption based on the time evolution of physically realistic Hamiltonians, which greatly prunes the size of the space, while Montina instead requires that you be able to get to any part of the space from any other part of the space, which seems like an extremely strong and unwarranted assumption, that directly contradicts that of Poulin et al. Are there any papers that have attempted to reconcile or otherwise discuss this contradiction?

#### Montina's assumption, verbatim

> "We assume that every unitary evolution is physically attainable. This hypothesis rests on the fact that, in quantum computers, every unitary evolution is in principle feasible by means of a finite number of quantum gates, that have a physical implementation."

So it's an attainability assumption on the *dynamics*, not quite ergodicity on the space — though for your purposes they amount to the same thing. One correction worth making, because it sharpens your point rather than weakening it: the assumption isn't *unwarranted*. Solovay–Kitaev guarantees every unitary is approximable by finitely many gates, so as a statement about quantum mechanics in principle it's true. The gap is that **"finite" is not "polynomial"** — and that is precisely the word Poulin et al. change. Montina's justification is resource-blind; Poulin et al. impose a resource bound. So the right criticism isn't that Montina assumed something false, it's that he assumed something *unbounded in resources*.

#### Montina's own discussion of your objection

From the concluding section:

> "Note that the subset of available quantum states can be parametrized with a number of variables which grows with the number of quantum gates, i.e., with the physical size. If this subset is dense in the Hilbert space or its parametrization is computational hard, then our theorem supports the conjecture on the polynomial noncomputability, **but these conditions cannot be assumed a priori.**"

That is your objection, granted in advance. He explicitly acknowledges that the physically available subset admits a parameterization growing only with physical size, and states that his theorem bites only under two further conditions he declines to assume. He also names which hypothesis he thinks should go: "In our opinion, the Markovian property is the only one sacrifiable."

So the correct reading is that Montina proved a **conditional**, not an unconditional exponential bound on physics. Poulin et al. then established that the antecedent's first branch — density in Hilbert space — fails. Both results stand.

#### Where the real question now sits

This matters for your larger argument, so it's worth being precise about what Poulin et al. actually deliver. Their conclusion:

> "all physical states live on a tiny submanifold, and that manifold can easily be parameterized by all poly-sized quantum circuits. This raises the question whether it makes sense to describe many-body quantum systems as vectors in a linear Hilbert space."

Note what the parameterization *is*: a poly-sized quantum circuit. That's a compact description — and extracting measurement statistics from it is BQP-hard. Which is exactly Montina's second branch: *"or its parametrization is computational hard."* Poulin et al. closed the density branch and left the hardness branch wide open, and it's the branch that's now widely believed to hold.

So the two papers together don't produce a contradiction; they reduce the whole dispute to a complexity-theoretic conjecture. If the physical corner's parameterization is computationally hard, the exponential resource is real in the only sense that matters operationally, even though the corner is geometrically thin. Your "largely unused degrees of freedom" survives as a claim about *volume*, and dies as a claim about *usable compression*.

#### Has anyone written the reconciliation?

I checked: **Poulin et al. cite Montina nowhere, and never mention ontological models or hidden variables at all.** The two results live in disjoint literatures — quantum many-body simulation on one side, quantum foundations on the other — which is presumably why the apparent clash has gone largely unremarked. I'm not aware of a paper that takes it on directly. What does engage the substance:

- **Montina's own discussion section** ([arXiv:0711.4770](https://arxiv.org/abs/0711.4770), final two paragraphs) — the primary source, and the closest thing to a reconciliation that exists.
- **Aaronson, "Multilinear Formulas and Skepticism of Quantum Computing"** — [quant-ph/0311039](https://arxiv.org/abs/quant-ph/0311039). This is the nearest thing to a paper *about* your tension, and it predates both. The "Sure/Shor separator" is exactly the project of asking whether the states we've actually seen admit compact descriptions while Shor-algorithm states don't, and he proves multilinear-formula lower bounds showing some physically relevant states resist compression.
- **Aaronson, "Quantum Computing and Hidden Variables"** — [quant-ph/0408035](https://arxiv.org/pdf/quant-ph/0408035), PRA **71**, 032325. Montina's own reference [20]; where hidden-variable dimension meets complexity theory. Montina discusses his relation to it directly.
- **Montina's later reframing in terms of communication complexity** — [arXiv:1206.2961](https://arxiv.org/abs/1206.2961) (PRL **109**, 110501) and [arXiv:1301.4441](https://arxiv.org/abs/1301.4441) (PRL **111**, 160502). Replacing "number of ontic variables" with "bits that must be communicated" is the resource-sensitive version of his own question, and is effectively his answer to the objection you're raising.
- **Harrigan & Spekkens** ontological-models framework, Found. Phys. **40**, 125 (2010) — where assumptions like Montina's become explicit and comparable across results.

#### Bottom line

You found a real gap in the *rhetoric* around these two results, not a contradiction between them. Montina's exponential bound is conditional on unbounded-resource attainability; Poulin et al. remove that and get a thin manifold; the thin manifold's parameterization is a quantum circuit nobody can classically evaluate. If you want the version of your thesis that has actual teeth, it's now a complexity claim rather than a geometric one — and Aaronson's Sure/Shor separator is the literature where that fight is held.

**Sources:** Montina quotes extracted from the [arXiv PDF](https://arxiv.org/pdf/0711.4770); Poulin et al. quotes and the absence of any Montina citation verified from the [arXiv PDF](https://arxiv.org/pdf/1102.1360). Aaronson identifiers confirmed via arXiv listings.

