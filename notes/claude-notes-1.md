This is a recording of a chat with Claude Opus 5 about Hilbert space, bell inequality, pilot wave models, etc. The results were used to inform the writing of bell.md, pilot-wave.md, configuration-space.md and a few others. This took place in the week of Sept 14, 2026.

Some key points:

* configuration space is not about nonlinear interactions, but rather about *correlations*.

* Montina08 is the key paper proving that exponential config space is necessary to represent the correlations, even allowing for nonlinear interactions! The key assumption is Markovian-ness. Need to understand this better!!  But they definitely say that it could be removed.  Stochastic motion + non-local wave influences seems like it would really change things.

* 3N dimensions -> $R^{3N}$ numbers is key about exponential explosion.

## Nonlinear SM vs linear Hilbert space

> Prompt: "the nonlinear interaction terms in the weak and strong forces of the standard model make the linear hilbert space model inaccurate"

Nonlinearity in the field equations and linearity of the state space are two different things, and the Standard Model has the first without giving up the second.

### Where the nonlinearity actually lives

In Yang–Mills theory the field strength is $F_{\mu\nu} = \partial_\mu A_\nu - \partial_\nu A_\mu + ig[A_\mu, A_\nu]$, so $\text{tr}\,F^2$ contains cubic and quartic terms in $A$. The classical equations of motion are genuinely nonlinear PDEs — gluons couple to gluons, W's couple to W's. That's why QCD confines and why non-abelian theories have instantons, sphalerons, and a nonperturbative vacuum structure that abelian theories don't.

But that nonlinearity is in the *operator* equations, i.e. in the Heisenberg-picture field equations and equivalently in the Hamiltonian's functional dependence on the field operators. It is not nonlinearity in the map that evolves states.

### Why the state space stays linear

Evolution is still:

$$
|\psi(t)\rangle = e^{-iHt}|\psi(0)\rangle
$$

and $e^{-iHt}$ is a linear operator no matter how high the powers of $\hat A$ inside $H$ are. Superposition holds exactly: if $|\psi_1\rangle$ and $|\psi_2\rangle$ each evolve some way, $\alpha|\psi_1\rangle + \beta|\psi_2\rangle$ evolves to $\alpha U|\psi_1\rangle + \beta U|\psi_2\rangle$, with no cross terms.

The cleanest illustration is the anharmonic oscillator, $H = p^2/2m + \lambda x^4$. Classical trajectories are nonlinear and chaotic-adjacent; the quantum evolution is exactly linear and unitary. "Interacting" and "nonlinear dynamics on Hilbert space" are simply not the same predicate. Feynman diagrams beyond tree level are precisely the perturbative bookkeeping for a nonlinear $H$ acting linearly on states.

Worth adding: fundamental nonlinearity in state evolution is not a free parameter you can quietly turn on. Weinberg-type nonlinear modifications of QM generically allow superluminal signaling and are experimentally bounded to extraordinary precision (NMR and atomic-clock tests put limits many orders of magnitude below typical energy scales). Nonlinear equations that *do* appear in practice — Gross–Pitaevskii, Hartree–Fock, semiclassical gravity's $G_{\mu\nu} = 8\pi\langle T_{\mu\nu}\rangle$ — are mean-field *approximations* to an underlying linear theory, and their nonlinearity is an artifact of the truncation.

### The real problems with the naive Hilbert-space picture

There are genuine reasons to distrust "the Standard Model = a vector in a Hilbert space," but they're different from the one you named:

- **Haag's theorem.** The interacting field theory is not unitarily equivalent to the free Fock-space representation. The Fock space you build from free particles is *the wrong space* for the interacting theory, which is one reason canonical perturbation theory is formally ill-defined and needs renormalization to make sense.
- **No rigorous construction.** Nobody has constructed the Hilbert space of 4D Yang–Mills with a mass gap — that's a Clay Millennium Problem. In that sense the Standard Model's Hilbert space is a physicist's assumption, not a theorem, and the self-interaction is exactly why.
- **Indefinite metric.** Covariant gauge quantization works in a space with negative-norm states; the physical Hilbert space appears only as BRST cohomology. So the "Hilbert space" of perturbative gauge theory isn't a Hilbert space until you quotient.
- **Confinement.** Asymptotic states aren't quark and gluon states at all, so the natural particle basis doesn't exist non-perturbatively.
- **Infrared structure.** In gauge theories with massless fields, soft radiation makes the naive S-matrix elements vanish; you need coherent states or dressed asymptotics, and the vacuum has a large degeneracy labeled by asymptotic charges.

The algebraic QFT response to all of this is to take the net of local operator algebras as primary and let Hilbert space representations be derived objects (via GNS), since inequivalent representations proliferate. That's a real demotion of the Hilbert space — but the algebras are still linear, and the states are still linear functionals on them.

So: if the claim is "self-interaction breaks linearity of quantum mechanics," no. If it's "the textbook Fock-space picture is not a valid description of the interacting Standard Model," that's correct and mathematically serious — just for reasons of inequivalent representations and non-perturbative construction rather than nonlinearity per se.

## Exponential configuration space

> Prompt: "does the use of exponential configuration space in the Hilbert space capture the nonlinear interaction terms?"

Yes, in a precise and rather satisfying mechanical sense — and no in the sense that matters for constructing the Standard Model. Both halves are worth separating.

### The sense in which it does

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

### The sense in which it does not

That last sentence is the whole problem. $\Gamma$ is intrinsically a *quasi-free* construction — it's the second quantization of a one-particle space equipped with a Gaussian state. The interacting vacuum is not Gaussian, and the tilted measure is generically not absolutely continuous with respect to $\mu_0$, so $L^2(d\mu) \not\cong \Gamma(\mathfrak{h})$ unitarily. That *is* Haag's theorem, seen from the measure-theoretic side.

Concretely, how far the exponential space gets you is dimension-dependent:

- **$d = 2, 3$:** Wick-ordered $:\!\varphi^4\!:$ is a densely defined self-adjoint operator on Fock space, and Glimm–Jaffe constructed the theory. The exponential space really does carry the nonlinear interaction.
- **$d = 4$:** it does not. $:\!\varphi^4\!:$ isn't a well-defined operator on Fock space, which is why you need renormalization rather than just Wick ordering — and $\varphi^4_4$ is now known to be trivial (Aizenman–Duminil-Copin), so the target theory isn't there to be constructed anyway.
- **Yang–Mills in $d=4$:** the object you'd want to build the exponential space *on* — the physical one-particle space — doesn't exist perturbatively. Confinement means asymptotic states are hadrons, not the gluons whose self-coupling generated the nonlinearity in the first place. The Fock space of gluons is the wrong exponential space by construction.

There's a partial rescue: extend $\Gamma(\mathfrak{h})$ to a Gelfand triple (white-noise / Hida distribution analysis), and Wick powers $:\!\varphi^n\!:$ become continuous maps from test vectors into distribution vectors. This is genuinely useful, but note what it concedes — the interaction is then a *sesquilinear form* on a dense subspace rather than an operator on the Hilbert space. The exponential space hosts the nonlinearity only in a weakened sense.

### Where that leaves it

The exponential structure is what makes nonlinear interaction *representable* by linear operators at all, and it's also where the classical nonlinear equations reappear in the $\hbar \to 0$ limit. What it doesn't do is guarantee that the interacting theory lives in that particular exponential space. Haag–Ruelle scattering theory says the correct picture is: the interacting Hilbert space is Fock-like *asymptotically*, built on physical dressed particles, given a mass gap and asymptotic completeness — while at finite times and short distances it isn't an exponential space of the Lagrangian's fields.

So: right structure for the kinematics of nonlinearity, wrong space for the dynamics. The failure mode is representation-theoretic, not a failure of linearity.

## Feynman diagrams as conditioning

> Prompt: "does the use of Feynman diagrams to configure the terms in an s-matrix integral constitute a kind of nonlinear operation, in that each such diagram holds for a specific type of interactions, and so you are already conditioning on energy levels and particle types in solving for a specific type of interaction? This is in comparison to a hypothetical alternative situation where the exact same set of equations is used for all possible calculations of the standard model.

No — extracting a specific process is a linear operation, and in a fairly strict sense. But your hypothetical alternative *exists*, it's the more fundamental formulation, and it is nonlinear — just not in the state space. That contrast is the interesting part.

### Why picking a process isn't a nonlinear step

$S$ is one fixed unitary operator on one fixed space. Computing $e^+e^- \to \mu^+\mu^-$ means evaluating $\langle f|S|i\rangle$ for particular $|i\rangle, |f\rangle$. Reading off a matrix element of a fixed linear operator in a chosen basis is the definition of a linear operation; it's no more nonlinear than writing $M_{23}$ for a matrix. Choosing external legs conditions the *input*, not the map.

Two further points cut against the "conditioning" reading:

**The diagrams aren't chosen, they're enumerated.** The Feynman rules are derived once from $\mathcal{L}$ — vertices, propagators, symmetry factors, all fixed before any process is contemplated. Given external legs and an order in the coupling, the diagram set is generated algorithmically; that's literally what FeynArts and MadGraph do. No physics judgment enters the enumeration. Wick's theorem does the work, and diagrams are its combinatorial bookkeeping.

**Processes are not independent of each other.** Unitarity of the single operator $S$ ties them together: the optical theorem and the Cutkosky cutting rules mean the imaginary part of a loop amplitude in one process is fixed by tree amplitudes of *other* processes. If each diagram were a separate conditioned calculation, that couldn't hold. Crossing symmetry does the same thing — one analytic function covers $s$-, $t$-, and $u$-channel processes in different kinematic regions. The interlocking is a signature of linearity, not of case-by-case treatment.

### Your hypothetical alternative, and its actual nonlinearity

There genuinely is a "same equations for everything" formulation. Three equivalent faces of it:

- **The generating functional** $Z[J] = \int \mathcal{D}\phi\, e^{iS[\phi] + i\int J\phi}$. Every correlator, every process, is a functional derivative of this one object.
- **The effective action** $\Gamma[\phi]$, whose stationarity condition $\delta\Gamma/\delta\phi = 0$ is the *exact* quantum equation of motion — the direct quantum successor to the nonlinear classical Yang–Mills equation.
- **The Dyson–Schwinger tower**, e.g. schematically $S^{-1} = S_0^{-1} - \Sigma[S, D, \Gamma_3]$, and the functional RG's Wetterich equation $\partial_k \Gamma_k = \tfrac12 \mathrm{Tr}\big[(\Gamma_k^{(2)} + R_k)^{-1}\partial_k R_k\big]$ — a single exact equation.

And these are **nonlinear** — the propagator appears inside its own self-energy, the vertex depends on the propagators which depend on the vertex, and the Wetterich equation has $\Gamma_k$ inverted on the right-hand side. So the answer to your question inverts nicely: the universal, non-conditioned formulation is the nonlinear one, and Feynman diagrams are what you get by *iteratively solving that nonlinear system order by order*. Diagram topologies are the iteration terms. Perturbation theory doesn't introduce nonlinearity by fragmenting a linear problem — it removes nonlinearity by linearizing a nonlinear functional problem one order at a time.

That's also why the nonlinear form is not redundant: dynamical mass generation and chiral symmetry breaking are solutions of the DSE tower that no finite set of diagrams produces, because they're non-analytic in the coupling. The self-referential structure is doing something the iterated expansion cannot.

### What your instinct is correctly detecting

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

### Bifurcation, and most of the mass of ordinary matter

The quark gap equation is a nonlinear fixed-point problem — the propagator appears inside its own self-energy. Below a critical coupling, the only solution is $m = 0$. Above it, a second branch appears with $m \neq 0$. That's a bifurcation, and it is the origin of dynamical chiral symmetry breaking: $\langle \bar q q\rangle \approx -(250\,\text{MeV})^3$, a constituent quark mass of a few hundred MeV from nearly massless inputs, and roughly 99% of the proton's mass having nothing to do with Higgs Yukawa couplings. A linear equation has no mechanism for "a qualitatively new solution appears past a threshold." This one does, and the answer is on your bathroom scale.

### Dimensional transmutation

$\beta(g) = -b_0 g^3 + b_1 g^5 + \dots$ is a nonlinear ODE, and integrating it converts a dimensionless coupling into a mass scale: $\Lambda_{\text{QCD}} \approx 200\text{–}300$ MeV, generated from a theory whose Lagrangian contains no scale at all. Asymptotic freedom is the sign of the leading nonlinear term, and it traces directly back to the gauge self-interaction you started this conversation with — the cubic and quartic gluon vertices are what make $b_0 > 0$. The resulting curve, $\alpha_s(M_Z) \approx 0.118$ running to $\approx 0.3$ near 2 GeV, has been checked across more than two orders of magnitude in energy. You are looking at the experimentally verified solution of a nonlinear equation.

### Fixed points, and exponents measured to six digits

RG flow *is* a nonlinear dynamical system on the space of couplings, and the whole vocabulary transfers: fixed points, basins of attraction, relevant and irrelevant directions as eigenvalues of the linearization about the nonlinear flow. Universality — why wildly different microscopic systems share critical exponents — is the statement that they flow to the same fixed point. Those exponents are measured in condensed matter to several significant figures. The conformal bootstrap pins the 3D Ising exponents to ~6 digits using *only* nonlinear consistency conditions (crossing, unitarity, analyticity), with no Lagrangian at all. Hard to call that formal bookkeeping.

Relatedly, unitarity itself is a nonlinear constraint: $S^\dagger S = 1$ is quadratic in amplitudes. The optical theorem, $\text{Im}\,\mathcal{M}(s) \propto \sigma_{\text{tot}}$, is a nonlinear relation among correlators that you can read off a detector. Dispersive positivity bounds built from it are currently used to carve out allowed regions of SMEFT Wilson coefficients — nonlinearity in correlator space constraining what new physics is possible.

And a live example: whether the electroweak vacuum is stable or metastable is entirely a question about nonlinear flow — whether the Higgs quartic $\lambda$ dips below zero around $10^{10}$–$10^{11}$ GeV. Current measurements put us near the metastable boundary. That's a physically consequential answer determined by the nonlinear structure and nothing else.

### The honest caveats

Not every nonlinear equation in the tower carries physical content. Individual off-shell Green's functions — the gluon propagator most notoriously — are gauge-dependent, and $\Gamma[\phi]$ shifts under field redefinitions. $\beta$ beyond two loops is scheme-dependent. So the correct claim is that the *nonlinear structure* is physical, while any particular nonlinear equation may be partly convention. The invariant content shows up in the places listed above: existence and location of bifurcations, fixed points and their exponents, generated scales, and relations among on-shell observables.

### The frame worth keeping

Nonlinearity is conserved under quantization, not eliminated. It gets relocated — out of the state space, where it would break superposition and permit superluminal signaling, and into the space of correlation functions and couplings, where it instead generates mass scales, drives symmetry breaking, and sets phase structure. Your original intuition that the Standard Model's self-interactions must have some non-linear consequence for the formalism was right. It just lands one level up from where you first put it, and it earns its keep there.

## Linearity needed for capturing underlying nonlinearity?

> Prompt: "So what is the implication of all this for the physical reality of the configuration space that is used in the Hilbert state space -- if it is only there to allow a linear framework to capture the fundamental nonlinearity of the underlying standard model, then perhaps there could be an alternative mathematical formulation that does not require this exponentially-large configuration space?"

Two claims are bundled in your question, and they have opposite answers. The *linearity* looks close to non-negotiable. The *exponential size* is very likely not physically real — but for reasons that have nothing to do with the self-interaction, and none of the existing alternatives eliminate what the space encodes. Let me also flag a premise worth adjusting: the exponential structure isn't there to accommodate nonlinearity. Free QED needs a Hilbert space, and so does a single spin. The exponential grading comes from multiparticle statistics, and the nonlinearity — as we traced — didn't get absorbed by inflating the state space, it got relocated to correlator space. So the space isn't scaffolding erected for the interactions.

### QFT already demotes configuration space

The "physical reality of the 3N-dimensional configuration space" debate is a non-relativistic quantum mechanics problem. In QFT the primitives are operator-valued fields attached to regions of 4D spacetime; particle number is not fundamental, and the Fock/configuration-space description is a derived, representation-dependent artifact — as Haag's theorem already forces on us. Local algebras in QFT are type III$_1$ von Neumann factors, which means there is no tensor factorization into "here" and "there," no local number operator, and no well-posed question "how many particles are in this box." The exponentially large configuration space has already lost its claim to being the furniture of the world, independently of anything you or I might prefer.

### Evidence the size is genuinely redundant

- **Holographic bounds.** The Bekenstein–Bousso bound says the degrees of freedom in a region scale with its bounding *area*, not its volume. Quantum field theory's Hilbert space is vastly larger than that — it overcounts, and gravity knows it. This is the strongest physical argument that the exponential space is not real.
- **Physical states are exponentially thin.** Ground states of local Hamiltonians obey area-law entanglement and are parameterized by polynomially many numbers; that's why DMRG and tensor networks work. And states reachable by polynomial-time evolution form a measure-zero subset of Hilbert space. Almost the entire space is never visited by anything.
- **Algebra type is not fixed.** Recent work (Leutheusser–Liu, Witten, and the Type II$_\infty$ story in black hole thermodynamics) has the algebra type itself changing with $G_N$ and with the observer's description — Hilbert space looking emergent and approximate rather than given.

### Formulations that don't build it

These aren't hypothetical; several are how frontier calculations are actually done.

- **Algebraic QFT.** Net of local algebras primary, state as a functional on the algebra, Hilbert space derived via GNS when you want it. This is the mainstream answer to "the Hilbert space isn't fundamental."
- **Euclidean functional integral / lattice.** A measure on field configurations, correlators as the data, Hilbert space *reconstructed* afterwards via Osterwalder–Schrader if you care. Lattice QCD computes the proton mass to a few percent without a Hilbert space appearing anywhere in the algorithm.
- **On-shell amplitude methods.** BCFW recursion, generalized unitarity, the amplituhedron: observables from locality and unitarity as *outputs* of positive geometry rather than inputs, with no Lagrangian and no state space. The honest limitation is that this is most developed for planar $\mathcal{N}=4$ and gravity, computes scattering rather than general finite-time dynamics, and is not yet a formulation of the Standard Model.
- **Bootstrap.** Spectrum and exponents from symmetry, unitarity, and crossing alone.
- **Dyson–Schwinger / functional RG** — the nonlinear formulation taken as primary, which is exactly the thing you were reaching for.

### What can't be dropped, and the price

Every one of those relocates the content rather than removing it — into an algebra's noncommutativity, into a measure, into positive geometry, into entanglement structure. And three constraints bite hard on any replacement:

1. **Tsirelson's bound.** Bell correlations exceed local-realistic limits by exactly the amount Hilbert space predicts. Reproduce that and you've reproduced the noncommutative tensor structure.
2. **Quantum computational advantage.** This is the sharpest one. If the exponential space were mere formal scaffolding compressible to polynomially many real parameters in general, Shor's algorithm shouldn't work. The dimension appears to be doing physical labor.
3. **Reconstruction theorems.** Gleason, and the operational axiomatizations (Hardy; Chiribella–D'Ariano–Perinotti), *derive* the Hilbert-space formalism — tensor products included — from a handful of operational requirements. The formalism isn't a modeling choice someone made; it's forced by the probability structure.

### My read

The exponentially large configuration space is best understood as a redundant but very efficient chart, not as the manifold. It's over-large by the holographic bound's measure, most of it is unreachable, and QFT already treats it as derived. So "not fundamental" is a defensible and I'd say probably correct position, and quantum gravity is where the seams show.

But the direction of the eventual replacement almost certainly isn't "nonlinear instead of linear." The candidates on the table — algebras, geometry, entanglement, measures — are all ways of keeping exact linearity and superposition while dropping the assumption that a single big vector space is the primary object. Your instinct that something is over-built is, I think, right. The over-built thing is the *dimensionality*, and the culprit is gravity's holographic bound, not the weak and strong forces.

## Literature on this

> Prompt: "what are some papers that discuss this topic most directly?"

Verified against arXiv/journal listings. Organized by the threads we went through.

### Is the exponentially large space physically real?

The single most direct paper for your question:

- **Poulin, Qarry, Somma, Verstraete, "Quantum simulation of time-dependent Hamiltonians and the convenient illusion of Hilbert space"** — [arXiv:1102.1360](https://arxiv.org/abs/1102.1360), PRL **106**, 170501 (2011). Proves that states reachable by *any* time-dependent local Hamiltonian in polynomial time occupy an exponentially small volume of Hilbert space. The title is your thesis, stated by people who then made it quantitative.
- **Carroll, "Reality as a Vector in Hilbert Space"** — [arXiv:2103.09780](https://arxiv.org/abs/2103.09780). The strongest defense of the opposite position: the vector *is* the ontology, and space and fields are emergent. Worth reading precisely because it's the view you're pushing against, argued carefully.
- **Giddings, "Hilbert space structure in quantum gravity: an algebraic perspective"** — [arXiv:1503.08207](https://arxiv.org/abs/1503.08207). Argues that factorization is already problematic in ordinary QFT, and that the algebra of observables should be primary.

### Hilbert space as derived rather than fundamental

- **Witten, "Notes on Some Entanglement Properties of Quantum Field Theory"** — [arXiv:1803.04993](https://arxiv.org/abs/1803.04993), Rev. Mod. Phys. **90**, 045003. The best entry point to type III$_1$ algebras, Reeh–Schlieder, and why "how many particles are in this region" is not a well-posed question. If you read one thing on the algebraic demotion of Hilbert space, read this.
- **Leutheusser & Liu** — [arXiv:2110.05497](https://arxiv.org/abs/2110.05497) and [arXiv:2112.12156](https://ui.adsabs.harvard.edu/abs/2023PhRvD.108h6020L/abstract) (both PRD 2023). Type III$_1$ structure *emerging* from a type I boundary algebra in the large-$N$ limit — algebra type as a derived, limit-dependent property.
- **Witten, "Gravity and the Crossed Product"** — [arXiv:2112.12828](https://arxiv.org/abs/2112.12828), JHEP 10 (2022) 008, and **Chandrasekaran, Penington, Witten, "Large $N$ Algebras and Generalized Entropy"** — [arXiv:2209.10454](https://arxiv.org/abs/2209.10454), JHEP 04 (2023) 009. Where the algebra type changes with $G_N$ and entropy becomes a property of the algebra rather than of a state in a fixed space.

### Why the configuration/Fock space is the wrong space for interacting fields

- **Earman & Fraser, "Haag's Theorem and its Implications for the Foundations of Quantum Field Theory"** — Erkenntnis **64** (2006) 305, [free preprint](https://philsci-archive.pitt.edu/2673/). Careful, readable, and exactly on the point that the interacting theory doesn't live in the exponential space you build for the free one.
- **Miller, "Haag's Theorem, Apparent Inconsistency, and the Empirical Adequacy of Quantum Field Theory"** — [BJPS **69** (2018)](https://www.journals.uchicago.edu/doi/abs/10.1093/bjps/axw029). The rebuttal: why the theorem doesn't undermine QFT's empirical success.

### The nonlinearity that is physically real

- **Roberts & Williams, "Dyson-Schwinger equations and their application to hadronic physics"** — [arXiv:hep-ph/9403224](https://arxiv.org/abs/hep-ph/9403224), Prog. Part. Nucl. Phys. **33** (1994) 477. The standard review; the gap-equation bifurcation is worked out explicitly.
- **Alkofer & von Smekal, "The infrared behaviour of QCD Green's functions"** — [Phys. Rept. **353** (2001) 281](https://ui.adsabs.harvard.edu/abs/2001PhR...353..281A/abstract). Confinement, dynamical symmetry breaking, *and* the positivity violation of transverse gluons — i.e. exactly where the naive Hilbert-space interpretation of the fields fails.
- **Berges, Tetradis, Wetterich, "Non-Perturbative Renormalization Flow in QFT and Statistical Physics"** — [arXiv:hep-ph/0005122](https://arxiv.org/abs/hep-ph/0005122), Phys. Rept. **363** (2002) 223. The exact nonlinear functional flow equation, treated as the primary object.
- **Roberts et al., "Empirical Consequences of Emergent Mass"** — [arXiv:2009.04011](https://arxiv.org/pdf/2009.04011). Ties the nonlinear structure to measured hadron properties.

### Why you can't simply make the state evolution nonlinear

- **Weinberg**, Annals Phys. **194** (1989) 336 and PRL **62**, 485 (1989) — the nonlinear generalization, proposed as a foil for precision tests.
- **Gisin**, Phys. Lett. A **143** (1990) 1 — superluminal signalling.
- **Polchinski**, PRL **66**, 397 (1991) — either EPR signalling or signalling between branches. See also the [later comment](https://arxiv.org/abs/quant-ph/0012041).
- **Bollinger et al.** — the $^9$Be$^+$ hyperfine test, bounding nonlinear corrections at $4\times10^{-27}$ of the binding energy per nucleon (PRL **63**, 1031, 1989 — journal page from memory, the experiment and bound are confirmed).
- **Kaplan & Rajendran, "Causal framework for nonlinear quantum mechanics"** — [PRD **105**, 055002](https://journals.aps.org/prd/abstract/10.1103/PhysRevD.105.055002). The modern attempt to have nonlinearity without signalling; useful for seeing how expensive it is.

### What forces Hilbert space in the first place

- **Hardy, "Quantum Theory From Five Reasonable Axioms"** — [arXiv:quant-ph/0101012](https://arxiv.org/abs/quant-ph/0101012).
- **Chiribella, D'Ariano, Perinotti, "Informational derivation of quantum theory"** — [arXiv:1011.6451](https://arxiv.org/abs/1011.6451), PRA **84**, 012311. Derives the formalism, tensor products included, from operational principles. Relevant because it shows the space isn't a modeling choice you can trade away for free.

### Formalisms that never build the state space

- **Arkani-Hamed & Trnka, "The Amplituhedron"** — [arXiv:1312.2007](https://arxiv.org/abs/1312.2007), JHEP 10 (2014) 030. Locality and unitarity as *outputs* rather than inputs; the opening paragraphs make your argument better than most philosophy papers do.
- **El-Showk et al., "Solving the 3D Ising Model with the Conformal Bootstrap"** — [arXiv:1203.6064](https://arxiv.org/abs/1203.6064), and the review **Poland, Rychkov, Vichi** — [arXiv:1805.04405](https://arxiv.org/abs/1805.04405), Rev. Mod. Phys. **91**, 015002.

### The configuration-space-realism debate itself

- **Wallace & Timpson, "Quantum Mechanics on Spacetime I: Spacetime State Realism"** — [BJPS **61** (2010) 697](https://philsci-archive.pitt.edu/4621/1/ststaterealism.pdf). Takes states-of-regions as fundamental instead of a vector on configuration space — the philosophical counterpart of the algebraic move.
- **Swanson, "How to Be a Relativistic Spacetime State Realist"** — [PDF](https://personal.lse.ac.uk/robert49/PPB/pdf/swanson2020str.pdf). Fixes the relativistic problems in the above using algebraic QFT.

If you want three to start: Poulin et al. for the quantitative version of your intuition, Witten's entanglement notes for why the field-theoretic Hilbert space is already not what the textbooks suggest, and Earman & Fraser for the cleanest statement of why the exponential space fails specifically once interactions are on.

**Sources:** all identifiers above were checked against arXiv and journal listings during this response, except the two explicitly flagged as recalled (Alkofer–von Smekal's arXiv number, which I've replaced with the journal reference, and the Bollinger PRL page number).

## More explicit statement of the question

Prompt: "Just to reiterate a bit on the discussion topic: I'm trying to understand the relationship between two different levels of description of the same phenomena. At the abstract level, the standard usage of the Hilbert space and Schrodinger equation in the 'old / nonrelativistic' version of quantum mechanics, which is the basis for most of the debate about the physical interpretation of quantum mechanics, models interactions among particles using the linear time-evolution of the Schrodinger equation operating over configuration space. By contrast, at the detailed level, you have the full complexity of the SM. So, the question is: is the linear Schrodinger equation somehow able to capture all the nonlinearities of the SM because it can use the extra degrees of freedom of the configuration space? In other words, the exponential size of the configuration space allows you to essentially have unique state terms for the interactions, which allow it to capture any kind of interaction because you have distinct degrees of freedom available in the state space to do so, even though you're just using a linear function (Schrodinger) to evolve the state over time. But, again to contrast with the SM: this more detailed picture of the 'actual physics' shows that the interactions are actually nonlinear, and this nonlinearity is mathematically challenging, but nevertheless, if you could actually solve the nonlinear math directly, you would not need the excessive (largely unused) extra degrees of freedom of a configuration space. Therefore, if all of this is true, then the configuration space is really just an artifact of using the linear Schrodinger equation to try to capture something that is actually fundamentally nonlinear. And therefore, everyone who gets upset about the exponential size of the configuration space is kind of missing the point. The 'actual nonlinear physics' doesn't actually need configuration space in the first place.  Does the structure of this argument make sense? Has someone articulated this point clearly in the literature?"

Yes, the structure makes sense — it's stated precisely enough to be tested, which most versions of this intuition aren't. And it fails, at one identifiable premise, with a theorem aimed almost directly at the counterfactual. Let me separate the two.

### The load-bearing premise

The claim doing the work is: *the exponential size of the state space is what allows a linear operator to represent nonlinear interactions.*

Take N free, non-interacting spin-½ particles. The Hamiltonian is $H = \sum_i h_i$ — no interaction terms at all, nothing nonlinear anywhere, not even a two-body potential. The state space is still $2^N$-dimensional, not $2N$-dimensional. And the singlet state lives in there, violating Bell inequalities.

So the exponential space is fully present when there is nothing whatsoever to linearize. That's a direct counterexample: the exponentiality cannot be an artifact of linearizing interactions, because it survives the removal of all interactions.

The correct diagnosis is that the exponential size comes from the **tensor product structure** — the fact that joint states are not determined by their marginals. It is about *correlation*, not interaction. Interactions are what *generate* entanglement dynamically, but they don't create the space; the space is there in the kinematics, before any dynamics.

### A second, independent problem with the same premise

Your picture implies something like: distinct interaction types get distinct degrees of freedom, so the state space acts as a lookup table that a linear map can route through. But look at what the Hamiltonian actually is. For N particles with two-body interactions, $H = \sum_i h_i + \sum_{i<j} V_{ij}$ — **polynomially many terms, each acting on at most two particles.** The interaction is specified by poly-many numbers. It's a sparse, local operator on an exponentially large space.

If the space were providing room to encode interactions, $H$ would need exponentially many parameters. It doesn't. The description of the interaction is compact; it's the *state* that's exponential. Those are different objects, and only one of them is big.

### The counterfactual is a theorem, and it goes the other way

Your C2 — "if you could solve the nonlinear math directly, you wouldn't need the extra degrees of freedom" — is precisely the question **Montina, "Exponential complexity and ontological theories of quantum mechanics"** ([arXiv:0711.4770](https://arxiv.org/abs/0711.4770), PRA **77**, 022104) settles. He proves that any ontological Markovian theory reproducing quantum dynamics needs at least $2N-2$ continuous variables, where $N$ is the Hilbert space dimension — so exponentially many in system size.

Crucially, **the theorem doesn't care whether the underlying dynamics is linear.** You are allowed your nonlinear equations. You still need exponentially many variables. The exponential resource is not scaffolding erected by the linear formalism; it's a lower bound on *any* description that reproduces the phenomena. The precursor is **Hardy, "Quantum ontological excess baggage"** (Stud. Hist. Phil. Mod. Phys. **35** (2004) 267), and the complementary result is **Pusey, Barrett, Rudolph, "On the reality of the quantum state"** ([Nature Physics **8**, 475](https://www.nature.com/articles/nphys2309)), which blocks the escape of treating the big state as mere ignorance about a small one.

Honest caveat: Montina's theorem assumes Markovian dynamics. Non-Markovian ontological models can evade the bound — but they buy it by paying in memory, which is not obviously cheaper, and it's a narrow door rather than a refutation.

And separately, empirically: any description with a small, local state is a local hidden variable theory. Classical nonlinear field theory is exactly that. Loophole-free Bell experiments closed that in 2015.

### The level-mixing

There's also a problem in how the two levels are related. Nonrelativistic QM's configuration space is not capturing the SM's nonlinearity — it's the fixed-particle-number sector, with interactions inserted by hand as potentials $V(x_i - x_j)$ that are the low-energy residue of field-theoretic exchange. The gauge self-coupling that generates the SM's nonlinearity doesn't survive to that level at all. Meanwhile in QFT proper, the state space is Fock space and — as we traced — the nonlinearity went into *correlator* space, not into the state space's dimensionality. So the two levels don't stand in the relationship the argument needs.

### What survives, and who has said it

Your conclusion is partly right, by a different route. The claim that the interpretive debate over-indexes on a 3N-dimensional configuration space that is an artifact — **that** is a live and well-defended position:

- **Myrvold, "What is a wavefunction?"** Synthese **192** (2015) 3247. Argues configuration space and the wavefunction can be real but not *fundamental*, with careful treatment of the relativistic case.
- **Wallace, "Against wavefunction realism"** ([PDF](https://files01.core.ac.uk/download/pdf/195905612.pdf), in *Current Controversies in Philosophy of Science*). Configuration space is one representation among infinitely many; nothing privileges it.
- **Wallace & Timpson**, spacetime state realism (cited earlier).

But note what they say the artifact is an artifact *of*: the nonrelativistic approximation — fixed particle number, a privileged position basis — not of linearization. Same conclusion, different and defensible mechanism.

As for the nonlinearity version specifically: I don't know of anyone who has argued it, and I think that's because it doesn't survive the free-particle counterexample. The nearest neighbors are 't Hooft's cellular-automaton interpretation and Wetterich's "quantum mechanics from classical statistics" — both trying to recover quantum linearity as emergent from a smaller underlying description — and both run directly into Bell, Montina, and PBR, which is where the arguments about them are actually fought.

### The version I'd defend

The exponential state space is not an artifact of linearizing nonlinear interactions; it's the cost of reproducing quantum correlations, and it's provably not compressible by switching to nonlinear underlying dynamics. But the *particular* exponential space of nonrelativistic configuration-space QM — the one the interpretation literature fixates on — genuinely is an artifact of the nonrelativistic limit, and QFT already dissolves it. And the space is over-large for an entirely separate reason we discussed: the holographic bound. Three distinct claims, one of which is yours and two of which give you most of what you wanted.

They don't contradict each other, but your diagnosis of where the load sits is right — and better than that, **Montina states the assumption explicitly and then concedes your objection himself**, three years before Poulin et al. I pulled both papers to check.

### Poulin et al vs. Montina

the paper by Poulin, Quarry, Somma & Verstraete (2011) appears to contradict the one by Montina (2008), where the former says that configuration space is exponentially too large, and Montina says that this exponential space is necessary under the Markovian assumptions. It appears that Poulin et al make use of an assumption based on the time evolution of physically realistic Hamiltonians, which greatly prunes the size of the space, while Montina instead requires that you be able to get to any part of the space from any other part of the space, which seems like an extremely strong and unwarranted assumption, that directly contradicts that of Poulin et al. Are there any papers that have attempted to reconcile or otherwise discuss this contradiction?

### Montina's assumption, verbatim

> "We assume that every unitary evolution is physically attainable. This hypothesis rests on the fact that, in quantum computers, every unitary evolution is in principle feasible by means of a finite number of quantum gates, that have a physical implementation."

So it's an attainability assumption on the *dynamics*, not quite ergodicity on the space — though for your purposes they amount to the same thing. One correction worth making, because it sharpens your point rather than weakening it: the assumption isn't *unwarranted*. Solovay–Kitaev guarantees every unitary is approximable by finitely many gates, so as a statement about quantum mechanics in principle it's true. The gap is that **"finite" is not "polynomial"** — and that is precisely the word Poulin et al. change. Montina's justification is resource-blind; Poulin et al. impose a resource bound. So the right criticism isn't that Montina assumed something false, it's that he assumed something *unbounded in resources*.

### Montina's own discussion of your objection

From the concluding section:

> "Note that the subset of available quantum states can be parametrized with a number of variables which grows with the number of quantum gates, i.e., with the physical size. If this subset is dense in the Hilbert space or its parametrization is computational hard, then our theorem supports the conjecture on the polynomial noncomputability, **but these conditions cannot be assumed a priori.**"

That is your objection, granted in advance. He explicitly acknowledges that the physically available subset admits a parameterization growing only with physical size, and states that his theorem bites only under two further conditions he declines to assume. He also names which hypothesis he thinks should go: "In our opinion, the Markovian property is the only one sacrifiable."

So the correct reading is that Montina proved a **conditional**, not an unconditional exponential bound on physics. Poulin et al. then established that the antecedent's first branch — density in Hilbert space — fails. Both results stand.

### Where the real question now sits

This matters for your larger argument, so it's worth being precise about what Poulin et al. actually deliver. Their conclusion:

> "all physical states live on a tiny submanifold, and that manifold can easily be parameterized by all poly-sized quantum circuits. This raises the question whether it makes sense to describe many-body quantum systems as vectors in a linear Hilbert space."

Note what the parameterization *is*: a poly-sized quantum circuit. That's a compact description — and extracting measurement statistics from it is BQP-hard. Which is exactly Montina's second branch: *"or its parametrization is computational hard."* Poulin et al. closed the density branch and left the hardness branch wide open, and it's the branch that's now widely believed to hold.

So the two papers together don't produce a contradiction; they reduce the whole dispute to a complexity-theoretic conjecture. If the physical corner's parameterization is computationally hard, the exponential resource is real in the only sense that matters operationally, even though the corner is geometrically thin. Your "largely unused degrees of freedom" survives as a claim about *volume*, and dies as a claim about *usable compression*.

### Has anyone written the reconciliation?

I checked: **Poulin et al. cite Montina nowhere, and never mention ontological models or hidden variables at all.** The two results live in disjoint literatures — quantum many-body simulation on one side, quantum foundations on the other — which is presumably why the apparent clash has gone largely unremarked. I'm not aware of a paper that takes it on directly. What does engage the substance:

- **Montina's own discussion section** ([arXiv:0711.4770](https://arxiv.org/abs/0711.4770), final two paragraphs) — the primary source, and the closest thing to a reconciliation that exists.
- **Aaronson, "Multilinear Formulas and Skepticism of Quantum Computing"** — [quant-ph/0311039](https://arxiv.org/abs/quant-ph/0311039). This is the nearest thing to a paper *about* your tension, and it predates both. The "Sure/Shor separator" is exactly the project of asking whether the states we've actually seen admit compact descriptions while Shor-algorithm states don't, and he proves multilinear-formula lower bounds showing some physically relevant states resist compression.
- **Aaronson, "Quantum Computing and Hidden Variables"** — [quant-ph/0408035](https://arxiv.org/pdf/quant-ph/0408035), PRA **71**, 032325. Montina's own reference [20]; where hidden-variable dimension meets complexity theory. Montina discusses his relation to it directly.
- **Montina's later reframing in terms of communication complexity** — [arXiv:1206.2961](https://arxiv.org/abs/1206.2961) (PRL **109**, 110501) and [arXiv:1301.4441](https://arxiv.org/abs/1301.4441) (PRL **111**, 160502). Replacing "number of ontic variables" with "bits that must be communicated" is the resource-sensitive version of his own question, and is effectively his answer to the objection you're raising.
- **Harrigan & Spekkens** ontological-models framework, Found. Phys. **40**, 125 (2010) — where assumptions like Montina's become explicit and comparable across results.

### Bottom line

You found a real gap in the *rhetoric* around these two results, not a contradiction between them. Montina's exponential bound is conditional on unbounded-resource attainability; Poulin et al. remove that and get a thin manifold; the thin manifold's parameterization is a quantum circuit nobody can classically evaluate. If you want the version of your thesis that has actual teeth, it's now a complexity claim rather than a geometric one — and Aaronson's Sure/Shor separator is the literature where that fight is held.

**Sources:** Montina quotes extracted from the [arXiv PDF](https://arxiv.org/pdf/0711.4770); Poulin et al. quotes and the absence of any Montina citation verified from the [arXiv PDF](https://arxiv.org/pdf/1102.1360). Aaronson identifiers confirmed via arXiv listings.

## Bell example

> take a maximally-entangled bell state (https://en.wikipedia.org/wiki/Bell_state) and, by explicitly writing out the full tensor product space for both qbits, show in full detail, first using symbolic equations and then with specific numbers, how the measurement process produces the CHSH inequality values

Everything below is verified numerically ([chsh.py](/private/tmp/claude-503/-Users-oreilly-Library-Application-Support-Claude-scratch-workspaces-c96e154b-9b88-4167-ac69-883556e884c2-e282c857-eff4-400f-a680-03d505db5460-scratch-2026-09-09-354f52/773531c2-447b-4f3a-a2bc-22d8c080a363/scratchpad/chsh.py)).

### 1. The single-qubit space

Each qubit lives in $\mathcal{H} = \mathbb{C}^2$ with orthonormal basis

$$|0\rangle = \begin{pmatrix}1\\0\end{pmatrix}, \qquad |1\rangle = \begin{pmatrix}0\\1\end{pmatrix}$$

### 2. The full tensor product space, written out

$\mathcal{H}_A \otimes \mathcal{H}_B = \mathbb{C}^2 \otimes \mathbb{C}^2 = \mathbb{C}^4$. Using the Kronecker convention $\binom{a}{b} \otimes \binom{c}{d} = (ac,\ ad,\ bc,\ bd)^T$, the four product basis vectors are:

$$|00\rangle = \begin{pmatrix}1\\0\\0\\0\end{pmatrix}\quad |01\rangle = \begin{pmatrix}0\\1\\0\\0\end{pmatrix}\quad |10\rangle = \begin{pmatrix}0\\0\\1\\0\end{pmatrix}\quad |11\rangle = \begin{pmatrix}0\\0\\0\\1\end{pmatrix}$$

Index ordering: component $k$ corresponds to $|i\rangle_A|j\rangle_B$ with $k = 2i + j$. **This is the crucial structure** — the space is 4-dimensional, not $2+2 = 4$-as-a-direct-sum. A general state $\sum_{ij} c_{ij}|ij\rangle$ needs 4 complex amplitudes, while a product state $(\alpha|0\rangle + \beta|1\rangle)\otimes(\gamma|0\rangle+\delta|1\rangle)$ has only 4 parameters constrained to the rank-1 surface $c_{00}c_{11} = c_{01}c_{10}$.

### 3. The Bell state

$$|\Phi^+\rangle = \frac{1}{\sqrt{2}}\left(|00\rangle + |11\rangle\right) = \frac{1}{\sqrt 2}\begin{pmatrix}1\\0\\0\\1\end{pmatrix}$$

**Verify it is entangled.** Write the amplitudes as a matrix $C_{ij}$:

$$C = \frac{1}{\sqrt 2}\begin{pmatrix}1 & 0\\ 0 & 1\end{pmatrix}, \qquad \det C = \tfrac12 \neq 0$$

A product state has $\det C = 0$ (rank 1). Here the Schmidt rank is 2, so $|\Phi^+\rangle$ cannot be written as $|\psi\rangle_A \otimes |\chi\rangle_B$ for *any* single-qubit states. The extra dimensions of the tensor product are doing real work — this state exists only because the space is a product, not a sum.

### 4. The measurement observables

Alice and Bob each measure a spin component in the $x$–$z$ plane, at angles $\alpha$ and $\beta$. The observable is $\hat n(\theta)\cdot\vec\sigma$ with $\hat n = (\sin\theta, 0, \cos\theta)$:

$$A(\alpha) = \sin\alpha\,\sigma_x + \cos\alpha\,\sigma_z = \begin{pmatrix}\cos\alpha & \sin\alpha \\ \sin\alpha & -\cos\alpha\end{pmatrix}$$

and identically $B(\beta)$. Each has eigenvalues $\pm 1$ (since $A^2 = I$, $\operatorname{tr}A = 0$), with eigenvectors

$$|{+}_\alpha\rangle = \begin{pmatrix}\cos\frac{\alpha}{2}\\[2pt] \sin\frac{\alpha}{2}\end{pmatrix},\qquad |{-}_\alpha\rangle = \begin{pmatrix}-\sin\frac{\alpha}{2}\\[2pt] \cos\frac{\alpha}{2}\end{pmatrix}$$

Check: $A(\alpha)|{+}_\alpha\rangle = \binom{\cos\alpha\cos\frac\alpha2 + \sin\alpha\sin\frac\alpha2}{\sin\alpha\cos\frac\alpha2 - \cos\alpha\sin\frac\alpha2} = \binom{\cos(\alpha/2)}{\sin(\alpha/2)} = +|{+}_\alpha\rangle$ ✓

The spectral projectors are

$$P^A_\pm(\alpha) = |{\pm}_\alpha\rangle\langle{\pm}_\alpha| = \tfrac12\begin{pmatrix}1\pm\cos\alpha & \pm\sin\alpha\\ \pm\sin\alpha & 1\mp\cos\alpha\end{pmatrix}, \qquad A(\alpha) = P^A_+ - P^A_-$$

### 5. The joint measurement on $\mathbb{C}^4$

Alice measuring $\alpha$ **and** Bob measuring $\beta$ is a single four-outcome projective measurement on the joint space, with rank-1 projectors

$$\Pi_{s_A s_B} = P^A_{s_A}(\alpha) \otimes P^B_{s_B}(\beta), \qquad s_A, s_B \in \{+,-\}$$

These are four orthogonal rank-1 projectors summing to $I_4$. Outcome probabilities by the Born rule:

$$P(s_A, s_B \mid \alpha,\beta) = \langle \Phi^+|\,\Pi_{s_A s_B}\,|\Phi^+\rangle = \left|\left(\langle s_A{}_\alpha| \otimes \langle s_B{}_\beta|\right)|\Phi^+\rangle\right|^2$$

### The four amplitudes

Using $\langle {+}_\alpha|0\rangle = \cos\frac\alpha2$, $\langle{+}_\alpha|1\rangle = \sin\frac\alpha2$, $\langle{-}_\alpha|0\rangle = -\sin\frac\alpha2$, $\langle{-}_\alpha|1\rangle = \cos\frac\alpha2$, and $|\Phi^+\rangle = \frac{1}{\sqrt2}(|00\rangle + |11\rangle)$:

$$\langle {+}_\alpha {+}_\beta|\Phi^+\rangle = \tfrac{1}{\sqrt2}\left[\cos\tfrac\alpha2\cos\tfrac\beta2 + \sin\tfrac\alpha2\sin\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\cos\tfrac{\alpha-\beta}{2}$$

$$\langle {+}_\alpha {-}_\beta|\Phi^+\rangle = \tfrac{1}{\sqrt2}\left[-\cos\tfrac\alpha2\sin\tfrac\beta2 + \sin\tfrac\alpha2\cos\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\sin\tfrac{\alpha-\beta}{2}$$

$$\langle {-}_\alpha {+}_\beta|\Phi^+\rangle = \tfrac{1}{\sqrt2}\left[-\sin\tfrac\alpha2\cos\tfrac\beta2 + \cos\tfrac\alpha2\sin\tfrac\beta2\right] = -\tfrac{1}{\sqrt2}\sin\tfrac{\alpha-\beta}{2}$$

$$\langle {-}_\alpha {-}_\beta|\Phi^+\rangle = \tfrac{1}{\sqrt2}\left[\sin\tfrac\alpha2\sin\tfrac\beta2 + \cos\tfrac\alpha2\cos\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\cos\tfrac{\alpha-\beta}{2}$$

Writing $\delta \equiv \alpha - \beta$, the joint probabilities are

$$\boxed{P(\pm,\pm) = \tfrac12\cos^2\tfrac{\delta}{2}, \qquad P(\pm,\mp) = \tfrac12\sin^2\tfrac\delta2}$$

**Normalization:** $\cos^2\frac\delta2 + \sin^2\frac\delta2 = 1$ ✓

**Marginals (no-signalling):**
$$P(A = +) = P(+,+) + P(+,-) = \tfrac12\cos^2\tfrac\delta2 + \tfrac12\sin^2\tfrac\delta2 = \tfrac12$$

Alice's marginal is $\tfrac12$ *regardless of $\beta$* — Bob's setting choice is invisible to Alice. All the structure is in the correlations, none in the marginals. This is why entanglement doesn't permit signalling.

### The correlation function

$$E(\alpha,\beta) = \langle \Phi^+|A(\alpha)\otimes B(\beta)|\Phi^+\rangle = P(+,+) + P(-,-) - P(+,-) - P(-,+)$$

$$= \cos^2\tfrac\delta2 - \sin^2\tfrac\delta2 = \boxed{\cos(\alpha - \beta)}$$

A cross-check without projectors: for $|\Phi^+\rangle = \frac{1}{\sqrt2}\sum_i|ii\rangle$ one has $\langle\Phi^+|M\otimes N|\Phi^+\rangle = \frac12\operatorname{tr}(MN^T)$, and $\operatorname{tr}\!\big(A(\alpha)B(\beta)\big) = 2\cos(\alpha-\beta)$. ✓

### 6. CHSH and the classical bound

With two settings each ($a, a'$ for Alice; $b, b'$ for Bob):

$$S = E(a,b) - E(a,b') + E(a',b) + E(a',b')$$

**Local hidden variable bound.** If each run carries predetermined values $A_a, A_{a'}, B_b, B_{b'} \in \{-1,+1\}$ fixed by some $\lambda$, then for each $\lambda$:

$$S(\lambda) = A_a(B_b - B_{b'}) + A_{a'}(B_b + B_{b'})$$

Since $B_b, B_{b'} = \pm1$: either $B_b = B_{b'}$, making the first bracket $0$ and the second $\pm2$; or $B_b = -B_{b'}$, making the second $0$ and the first $\pm2$. Either way $|S(\lambda)| \le 2$, and averaging over $\lambda$ preserves it:

$$\boxed{|S_{\text{LHV}}| \le 2}$$

**Quantum prediction.** With $E = \cos(\alpha-\beta)$:
$$S = \cos(a-b) - \cos(a-b') + \cos(a'-b) + \cos(a'-b')$$

Choosing $a = 0,\ a' = \frac\pi2,\ b = \frac\pi4,\ b' = \frac{3\pi}{4}$ makes all four terms contribute $+\frac{\sqrt2}{2}$:
$$S = 4\cdot\tfrac{\sqrt2}{2} = 2\sqrt2 \approx 2.8284$$

### 7. Now with explicit numbers

#### The four observables

```
A  = A(0°)   = σ_z          = [[ 1.000000,  0.000000],
                               [ 0.000000, -1.000000]]

A' = A(90°)  = σ_x          = [[ 0.000000,  1.000000],
                               [ 1.000000,  0.000000]]

B  = B(45°)  = (σ_x+σ_z)/√2 = [[ 0.707107,  0.707107],
                               [ 0.707107, -0.707107]]

B' = B(135°) = (σ_x−σ_z)/√2 = [[-0.707107,  0.707107],
                               [ 0.707107,  0.707107]]
```

Each has eigenvalues exactly $\{+1, -1\}$.

#### The four 4×4 tensor products

Shown multiplied by $\sqrt2$ so the entries are integers. Acting on $|\Phi^+\rangle = \frac{1}{\sqrt2}(1,0,0,1)^T$:

```
√2·(A ⊗ B)                    √2·(A ⊗ B')
[ 1   1   0   0]              [-1   1   0   0]
[ 1  -1   0   0]              [ 1   1   0   0]
[ 0   0  -1  -1]              [ 0   0   1  -1]
[ 0   0  -1   1]              [ 0   0  -1  -1]
⟨Φ⁺|·|Φ⁺⟩ = +0.707107         ⟨Φ⁺|·|Φ⁺⟩ = −0.707107

√2·(A' ⊗ B)                   √2·(A' ⊗ B')
[ 0   0   1   1]              [ 0   0  -1   1]
[ 0   0   1  -1]              [ 0   0   1   1]
[ 1   1   0   0]              [-1   1   0   0]
[ 1  -1   0   0]              [ 1   1   0   0]
⟨Φ⁺|·|Φ⁺⟩ = +0.707107         ⟨Φ⁺|·|Φ⁺⟩ = +0.707107
```

Worked example for $A \otimes B$, by hand. Let $M = \frac{1}{\sqrt2}\begin{psmallmatrix}1&1&0&0\\1&-1&0&0\\0&0&-1&-1\\0&0&-1&1\end{psmallmatrix}$ and $v = \frac{1}{\sqrt2}(1,0,0,1)^T$:

$$Mv = \tfrac{1}{\sqrt2}\cdot\tfrac{1}{\sqrt2}\begin{pmatrix}1\cdot1 + 1\cdot 0\\ 1\cdot1 - 1\cdot0 \\ -1\cdot 0 - 1\cdot 1\\ -1\cdot0 + 1\cdot1\end{pmatrix} = \tfrac12\begin{pmatrix}1\\1\\-1\\1\end{pmatrix}$$

$$v^T(Mv) = \tfrac{1}{\sqrt2}(1,0,0,1)\cdot\tfrac12(1,1,-1,1)^T = \tfrac{1}{2\sqrt2}(1 + 1) = \tfrac{1}{\sqrt2} = 0.707107 \;\checkmark$$

#### All joint probabilities

| settings | $\delta = \alpha-\beta$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | sum | $E$ |
|---|---|---|---|---|---|---|---|
| $(a,b) = (0°,45°)$ | $-45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |
| $(a,b') = (0°,135°)$ | $-135°$ | 0.073223 | 0.426777 | 0.426777 | 0.073223 | 1.000000 | $-0.707107$ |
| $(a',b) = (90°,45°)$ | $+45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |
| $(a',b') = (90°,135°)$ | $-45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |

Every marginal is exactly $0.500000$ for both parties in all four rows.

Sample check of row 1: $\frac12\cos^2(-22.5°) = \frac12(0.923880)^2 = \frac12(0.853553) = 0.426777$ ✓, and $\frac12\sin^2(-22.5°) = \frac12(0.382683)^2 = 0.073223$ ✓.

#### Assembling $S$

$$S = (+0.707107) - (-0.707107) + (0.707107) + (0.707107)$$
$$= 0.707107 \times 4 = \mathbf{2.828427}$$

$$2\sqrt2 = 2.828427\ \checkmark \qquad\text{versus the classical bound } 2$$

Violation ratio: $2\sqrt2 / 2 = \sqrt2 \approx 1.414$.

### 8. Why exactly $2\sqrt2$

Form the CHSH operator on $\mathbb{C}^4$:

$$\hat S = A\otimes B - A\otimes B' + A'\otimes B + A'\otimes B'$$

Applying it to the Bell state numerically gives $\hat S|\Phi^+\rangle = 2\sqrt2\,|\Phi^+\rangle$ — **the Bell state is an exact eigenvector of the CHSH operator at these settings**, with eigenvalue $2\sqrt2$. That's why the violation is maximal: you can't do better than an eigenvalue, and Tsirelson's bound says no eigenvalue is larger. The one-line reason (Landau's identity):

$$\hat S^2 = 4\,I \;+\; [A,A']\otimes[B,B']$$

Since $\|A\| = \|A'\| = 1$, $\|[A,A']\| \le 2$, so $\|\hat S^2\| \le 4 + 4 = 8$ and $\|\hat S\| \le 2\sqrt2$. If $A$ and $A'$ commuted — the classical case, where both have simultaneous definite values — the commutator term vanishes and you recover $|S| \le 2$ exactly. **The violation is precisely the non-commutativity of the two measurement settings, amplified through the tensor product structure.**

---

Want this as a properly typeset page (real matrices, MathJax, the probability derivation laid out side by side)? One line and I'll publish it as an artifact.

## Bell non-entangled case

> thanks! do this same thing now for a non-entangled bell state, to show explicitly the difference.

Small terminology note first: all four Bell states are *maximally* entangled by definition — there's no unentangled one. So I'll use the two natural foils: a **pure product state**, and the **separable mixture with the same marginals as $|\Phi^+\rangle$**, which is the sharper comparison. Let me verify both numerically.

All verified ([chsh2.py](/private/tmp/claude-503/-Users-oreilly-Library-Application-Support-Claude-scratch-workspaces-c96e154b-9b88-4167-ac69-883556e884c2-e282c857-eff4-400f-a680-03d505db5460-scratch-2026-09-09-354f52/773531c2-447b-4f3a-a2bc-22d8c080a363/scratchpad/chsh2.py)). Same machinery as before — same $\mathbb{C}^4$, same observables $A(\alpha)$, same angles $a=0°, a'=90°, b=45°, b'=135°$ — only the state changes.

### Case A: the pure product state $|0\rangle \otimes |0\rangle$

#### The state in the full tensor product space

$$|\psi_{\text{prod}}\rangle = |0\rangle_A \otimes |0\rangle_B = \begin{pmatrix}1\\0\end{pmatrix}\otimes\begin{pmatrix}1\\0\end{pmatrix} = \begin{pmatrix}1\\0\\0\\0\end{pmatrix}$$

**Entanglement test.** The amplitude matrix is

$$C = \begin{pmatrix}1 & 0\\0&0\end{pmatrix}, \qquad \det C = 0$$

Schmidt rank 1 — a genuine product state. Contrast with $|\Phi^+\rangle$, where $\det C = \frac12$.

#### The correlation factorizes — symbolically

This is the whole difference, in one line. Because the state is a product, the expectation of any product observable splits:

$$E(\alpha,\beta) = \langle 0|\otimes\langle 0|\;A(\alpha)\otimes B(\beta)\;|0\rangle\otimes|0\rangle = \underbrace{\langle 0|A(\alpha)|0\rangle}_{\text{Alice alone}}\cdot\underbrace{\langle 0|B(\beta)|0\rangle}_{\text{Bob alone}}$$

With $\langle 0|A(\alpha)|0\rangle = A(\alpha)_{00} = \cos\alpha$:

$$\boxed{E_{\text{prod}}(\alpha,\beta) = \cos\alpha\,\cos\beta} \qquad\text{versus}\qquad E_{\Phi^+}(\alpha,\beta) = \cos(\alpha-\beta)$$

**This is the structural difference.** $\cos\alpha\cos\beta$ is a *separable* function — a product of a function of $\alpha$ alone and a function of $\beta$ alone. $\cos(\alpha-\beta)$ is not; it cannot be factored, because $\cos(\alpha-\beta) = \cos\alpha\cos\beta + \sin\alpha\sin\beta$ has two terms. That second term $\sin\alpha\sin\beta$ is the entire Bell violation.

#### The joint probabilities factorize too

$$P(s_A,s_B) = \left|\langle s_A{}_\alpha|0\rangle\right|^2 \cdot \left|\langle s_B{}_\beta|0\rangle\right|^2 = P(s_A)\cdot P(s_B)$$

with $P(A{=}+) = \cos^2\frac\alpha2$, $P(B{=}+) = \cos^2\frac\beta2$. The outcomes are **statistically independent** — the joint distribution is literally the product of the marginals, for every setting pair. There is nothing to correlate.

#### Numbers at the Bell-optimal angles

| settings | $E$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | $P(A{=}{+})$ | $P(B{=}{+})$ |
|---|---|---|---|---|---|---|---|
| $(a,b)=(0°,45°)$ | $+0.707107$ | 0.853553 | 0.146447 | 0.000000 | 0.000000 | 1.0000 | 0.8536 |
| $(a,b')=(0°,135°)$ | $-0.707107$ | 0.146447 | 0.853553 | 0.000000 | 0.000000 | 1.0000 | 0.1464 |
| $(a',b)=(90°,45°)$ | $\;\;0.000000$ | 0.426777 | 0.073223 | 0.426777 | 0.073223 | 0.5000 | 0.8536 |
| $(a',b')=(90°,135°)$ | $\;\;0.000000$ | 0.073223 | 0.426777 | 0.073223 | 0.426777 | 0.5000 | 0.1464 |

Factorization check on row 3: $P(+,+) = 0.5 \times 0.853553 = 0.426777$ ✓ exactly.

$$S = 0.707107 - (-0.707107) + 0 + 0 = \mathbf{1.414214} = \sqrt2$$

Note that Alice's marginal here *does* depend on her own setting ($1.0$ vs $0.5$) — unlike the Bell state, where it was always $0.5$. That's a giveaway that this isn't a fair comparison. Hence:

### Case B: the separable mixture (the sharp foil)

$$\rho_{\text{mix}} = \tfrac12|00\rangle\langle00| + \tfrac12|11\rangle\langle11|$$

This is the *decohered* Bell state: identical local marginals, identical perfect correlation in the computational basis, but classically correlated rather than entangled. Half the pairs are $|00\rangle$, half are $|11\rangle$, decided by a coin flip at the source. Exactly the "local hidden variable" picture.

#### Density matrices side by side

```
ρ_Bell = |Φ⁺⟩⟨Φ⁺|                  ρ_mix = ½|00⟩⟨00| + ½|11⟩⟨11|
  [0.5   0    0   0.5]               [0.5   0    0    0 ]
  [ 0    0    0    0 ]               [ 0    0    0    0 ]
  [ 0    0    0    0 ]               [ 0    0    0    0 ]
  [0.5   0    0   0.5]               [ 0    0    0   0.5]
         ↑            ↑
    the coherences — the only difference
```

**The entire distinction is those two off-diagonal $0.5$'s.** Same diagonal, same probabilities in the $z$ basis, same reduced states:

$$\operatorname{Tr}_B \rho_{\text{Bell}} = \operatorname{Tr}_B \rho_{\text{mix}} = \begin{pmatrix}0.5 & 0\\0&0.5\end{pmatrix} = \tfrac{I}{2}$$

So no local measurement on either qubit alone can tell them apart. Only the *correlations across settings* can.

#### Peres–Horodecki (PPT) test

Partial transpose on $B$:

```
ρ_Bell^{T_B} =                     ρ_mix^{T_B} =
  [0.5   0    0    0 ]               [0.5   0    0    0 ]
  [ 0    0   0.5   0 ]               [ 0    0    0    0 ]
  [ 0   0.5   0    0 ]               [ 0    0    0    0 ]
  [ 0    0    0   0.5]               [ 0    0    0   0.5]
eigenvalues: ½, ½, +½, −½            eigenvalues: ½, 0, 0, ½
       NEGATIVE ⇒ ENTANGLED                 all ≥ 0 ⇒ SEPARABLE
```

The middle block $\begin{psmallmatrix}0&0.5\\0.5&0\end{psmallmatrix}$ has eigenvalues $\pm\frac12$. That single negative eigenvalue is a certificate of entanglement.

#### The correlation function

$$E_{\text{mix}}(\alpha,\beta) = \operatorname{Tr}\!\left[\rho_{\text{mix}}\,A(\alpha)\otimes B(\beta)\right] = \tfrac12\langle00|A\otimes B|00\rangle + \tfrac12\langle11|A\otimes B|11\rangle$$

$$= \tfrac12\big[\cos\alpha\cos\beta\big] + \tfrac12\big[(-\cos\alpha)(-\cos\beta)\big] = \boxed{\cos\alpha\,\cos\beta}$$

Same separable form as the pure product state.

#### Numbers

| settings | $E_{\text{mix}}$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | $P(A{=}{+})$ |
|---|---|---|---|---|---|---|
| $(a,b)=(0°,45°)$ | $+0.707107$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 0.5000 |
| $(a,b')=(0°,135°)$ | $-0.707107$ | 0.073223 | 0.426777 | 0.426777 | 0.073223 | 0.5000 |
| $(a',b)=(90°,45°)$ | $\;\;0.000000$ | 0.250000 | 0.250000 | 0.250000 | 0.250000 | 0.5000 |
| $(a',b')=(90°,135°)$ | $\;\;0.000000$ | 0.250000 | 0.250000 | 0.250000 | 0.250000 | 0.5000 |

$$S = 0.707107 - (-0.707107) + 0 + 0 = \mathbf{1.414214} = \sqrt2$$

### The diagnostic: where exactly they diverge

Put the two states' correlations side by side:

| settings | $E$ for $|\Phi^+\rangle$ | $E$ for $\rho_{\text{mix}}$ | same? |
|---|---|---|---|
| $(a,b) = (0°,45°)$ | $+0.707107$ | $+0.707107$ | ✅ **identical** |
| $(a,b') = (0°,135°)$ | $-0.707107$ | $-0.707107$ | ✅ **identical** |
| $(a',b) = (90°,45°)$ | $+0.707107$ | $\;\;0.000000$ | ❌ |
| $(a',b') = (90°,135°)$ | $+0.707107$ | $\;\;0.000000$ | ❌ |
| | $S = 2.828427$ | $S = 1.414214$ | |

Look at the first two rows: for Alice's setting $a = 0°$ (i.e. $A = \sigma_z$), the mixture reproduces the Bell state's joint probabilities **exactly, to every decimal place** — not just the correlation but all four $P(s_A,s_B)$. The classical mixture is a perfect impostor as long as Alice measures $\sigma_z$.

It fails the moment Alice switches to $a' = 90°$, where $A' = \sigma_x$. $\sigma_x$ is off-diagonal, so its expectation reads the *coherences* — precisely the matrix elements that $\rho_{\text{mix}}$ lacks. Both of those rows collapse to $E = 0$, killing half the CHSH sum.

**The single cleanest statement of the difference.** Measure both parties along the *same* in-plane angle $\theta$:

$$E_{\Phi^+}(\theta,\theta) = \cos(\theta-\theta) = 1 \quad\text{for every }\theta$$
$$E_{\text{mix}}(\theta,\theta) = \cos^2\theta \quad\Rightarrow\quad 1 \text{ at } 0°,\; 0.5 \text{ at } 45°,\; 0 \text{ at } 90°$$

The Bell state is perfectly correlated in *every* basis. The mixture is perfectly correlated in *one* basis and uncorrelated in the conjugate one. That basis-independence is what has no classical analogue, and CHSH is the device that detects it by querying two incompatible bases at once.

### Why no separable state can ever reach 2

For $\rho_{\text{mix}}$ the local hidden variable model is not hypothetical — write it down:

- $\lambda \in \{0, 1\}$, each with probability $\tfrac12$ (which pair was emitted).
- Given $\lambda$, Alice outputs $+1$ with probability $|\langle{+}_\alpha|\lambda\rangle|^2$, using only $\alpha$ and $\lambda$.
- Given $\lambda$, Bob independently outputs $+1$ with probability $|\langle{+}_\beta|\lambda\rangle|^2$, using only $\beta$ and $\lambda$.

Then $E(\alpha,\beta) = \frac12[\cos\alpha\cos\beta] + \frac12[(-\cos\alpha)(-\cos\beta)] = \cos\alpha\cos\beta$ ✓ — reproduced exactly, with no communication. Since an LHV model exists, the $|S| \le 2$ proof from before applies verbatim.

Maximizing over *all* settings for $E = \cos\alpha\cos\beta$:

$$S = \cos a\,(\cos b - \cos b') + \cos a'\,(\cos b + \cos b')$$
$$\max S = |\cos b - \cos b'| + |\cos b + \cos b'| = 2\max(|\cos b|, |\cos b'|) \le 2$$

The bound $2$ is reachable (e.g. $a = 0°, b = 0°, b' = 180°$) — but only by the degenerate choice $B' = -B$, where Bob isn't really using two independent settings. Saturated, never exceeded.

### Summary

| | $|\Phi^+\rangle$ | $|0\rangle|0\rangle$ | $\rho_{\text{mix}}$ |
|---|---|---|---|
| Schmidt rank / separable? | 2, **entangled** | 1, product | separable (mixed) |
| $\det C$ or PPT | $-\tfrac12$ eigenvalue | — | PPT, all $\ge 0$ |
| local marginals | $I/2$ | $|0\rangle\langle0|$ | $I/2$ |
| $E(\alpha,\beta)$ | $\cos(\alpha-\beta)$ | $\cos\alpha\cos\beta$ | $\cos\alpha\cos\beta$ |
| factorizes? | **no** | yes | yes |
| $E(\theta,\theta)$ | $1$ for all $\theta$ | $\cos^2\theta$ | $\cos^2\theta$ |
| $S$ at Bell angles | $2\sqrt2 = 2.828$ | $\sqrt2 = 1.414$ | $\sqrt2 = 1.414$ |
| $S$ maximized | $2\sqrt2$ | $2$ | $2$ |
| LHV model exists? | **no** | yes | yes |

The one-sentence version: entanglement shows up as a correlation function that **cannot be factored into a function of Alice's setting times a function of Bob's**, and CHSH is exactly the linear combination that is bounded by 2 for every factorizable correlation and reaches $2\sqrt2$ for the non-factorizable one.

### Bell superposition with classical wave

> Isn't it pretty obvious then that the entire key to the bell inequality effects is superposition? how could you possibly expect to reproduce the effects without superposition? but can't you have superposition in a classical wave? has anyone tried to capture the bell inequality effects with some kind of local realist wave that supports superposition?

Your instinct is half right, and the half that's wrong is the interesting half.

### Superposition alone does nothing

Take $|+\rangle|+\rangle$ where $|+\rangle = \frac{1}{\sqrt2}(|0\rangle+|1\rangle)$. Both qubits are *maximally* superposed in the computational basis. Running it through the identical machinery:

| state | $|\det C|$ | $S$ at Bell angles | max $S$ over all settings |
|---|---|---|---|
| $\|0\rangle\|0\rangle$ | 0 | 1.414214 | **2.000000** |
| $\|+\rangle\|+\rangle$ — fully superposed | 0 | 1.414214 | **2.000000** |
| $\|+\rangle\|0\rangle$ | 0 | 0.000000 | **2.000000** |
| arbitrary product $(20°)\otimes(70°)$ | 0 | −0.245576 | **2.000000** |
| $\|\Phi^+\rangle$ | 0.5 | 2.828427 | **2.828427** |

Every product state caps at exactly 2.000000 no matter how much superposition you pour into it. The reason is the factorization from last time: $E = \langle A\rangle\langle B\rangle$ holds for *any* product state, superposed or not, and a factorizable correlation is an LHV correlation.

### Why "superposition" can't be the right concept

It's **basis-dependent**. Write $|00\rangle$ in the $x$ basis:

$$|00\rangle = \tfrac12\big(|{+}{+}\rangle + |{+}{-}\rangle + |{-}{+}\rangle + |{-}{-}\rangle\big)$$

A four-term superposition — and still a product state, still $S \le 2$. Conversely $|\Phi^+\rangle$ is a superposition in *every* local basis; no change of local coordinates turns it into a product. **Superposition is a property of a state relative to a basis you chose. Entanglement is a property of the state.** That's why $\det C$ (Schmidt rank) is the invariant and "is it a superposition" isn't.

The precise thing is: **superposition that cannot be factored across the tensor product.**

### And there's a second necessary ingredient

Even with a Bell state, keep the state and make Alice's two settings compatible:

| Alice's settings | $\max S$ |
|---|---|
| $a = 0°,\ a' = 0°$ (identical, $[A,A']=0$) | **2.000000** |
| $a = 0°,\ a' = 30°$ | 2.449490 |
| $a = 0°,\ a' = 60°$ | 2.732051 |
| $a = 0°,\ a' = 90°$ (maximally incompatible) | **2.828427** |

With commuting settings the Bell state violates nothing. This is the Landau identity made numerical: $\hat S^2 = 4I + [A,A']\otimes[B,B']$. You need entanglement **and** non-commuting local measurements; either one alone is worth exactly 2.

### Yes, classical waves superpose — and people have built the analogue

This is a real and active field, not a gap in the literature. Classical optics gives you superposition for free (Maxwell's equations are linear), and you can form a **tensor product of two degrees of freedom of a single beam** — polarization $\otimes$ transverse spatial mode. States of that beam that are non-separable across those two DOFs exist, and they **do violate a CHSH-form inequality**, with classical light, no quantization needed.

Spreeuw called this "a classical analogy of entanglement" (1998); Eberly, Qian and collaborators developed it extensively; Karimi and Boyd wrote a *Science* perspective on it under the title "Classical entanglement?"

What it demonstrates is precise and worth absorbing: **the CHSH algebra is a consequence of tensor-product structure alone, and has nothing to do with quantum mechanics per se.** What it does *not* give you is nonlocality — because polarization and spatial mode live in the *same beam, in the same place*. You cannot hand one to Alice in Vienna and the other to Bob in Delft. Bell's theorem is about spacelike-separated measurements; this construction has the mathematics without the separation, so nothing nonlocal is being shown.

### Why a local-realist wave theory can't work

Here's the thing that makes the attempts futile in advance: **Bell's theorem doesn't care what the hidden variable is.** $\lambda$ can be a scalar, a vector, a classical field configuration, an infinite-dimensional wave, an entire universe's worth of structure. The theorem never opens the box. It assumes only:

1. **Factorizability:** $P(a,b \mid \alpha,\beta,\lambda) = P(a\mid\alpha,\lambda)\,P(b\mid\beta,\lambda)$
2. **Measurement independence:** $\rho(\lambda)$ doesn't depend on $\alpha,\beta$
3. Definite $\pm1$ outcomes

From those three, $|S| \le 2$ follows by the two-line argument I gave earlier. Richness of $\lambda$ buys you exactly nothing, because what's bounded is the *factorization of the probability*, not the complexity of the wave. A classical wave superposes, but it superposes *locally* — its amplitude at Alice's site is a function of Alice's site.

### The attempts, and how each one pays

- **Stochastic electrodynamics** (Marshall, Boyer, de la Peña & Cetto): classical EM plus a random zero-point field. Genuinely reproduces the harmonic-oscillator ground state, blackbody spectrum, Casimir effect. Cannot violate Bell.
- **Local realistic optical models** (Santos, Marshall; Khrennikov's "prequantum classical statistical field theory"): classical wave fields aiming at quantum optics. They reproduce a lot — and stop at Bell.
- **Detection-loophole models** (Pearle 1970; Clauser–Horne; Gisin & Gisin): these *do* reproduce apparent CHSH violations with local classical waves — by exploiting unfair sampling. Below a critical detector efficiency ($\eta \approx 82.8\%$ for maximally entangled CHSH, $2/3$ with Eberhard's non-maximally entangled version) a local model can fake any violation. This is the one place classical waves genuinely succeeded, and it's exactly what the loophole-free experiments of 2015 closed.
- **Hydrodynamic pilot waves** (Couder & Fort's walking droplets; Bush's program): a real classical wave-particle system reproducing single-slit statistics, tunneling analogues, quantized orbits. Bush and collaborators are explicit that it cannot produce Bell violations or multiparticle entanglement — because the droplet's wave lives in physical space.
- **'t Hooft's cellular automaton interpretation:** escapes by denying assumption 2 (superdeterminism).
- **Bohmian mechanics:** *is* a wave theory with superposition and definite outcomes — and is explicitly, irreducibly nonlocal. It's the proof of concept that if you want a realist wave theory reproducing QM, locality is the price.
- **Everett:** keeps locality, pays by giving up assumption 3 (single definite outcomes).

Every one of these is a receipt for one of Bell's three assumptions.

### The structural reason — and it connects to where we started

A classical wave and a quantum state superpose in spaces that grow differently:

$$\text{$N$ classical fields on } \mathbb{R}^3: \quad \dim \sim N \times (\text{field DOFs}) \qquad \textbf{additive}$$
$$\text{$N$ qubits}: \quad \dim = 2^N \qquad\qquad\qquad\qquad\quad\ \textbf{multiplicative}$$

Two classical waves in the same room give you a wave in *one* 3-dimensional space with two excitations. Two qubits give you $\mathbb{C}^2\otimes\mathbb{C}^2$, where $|\Phi^+\rangle$ lives in the part that isn't any product. Classical waves have superposition without the tensor product, and $\det C = 0$ is a theorem for them, not an accident.

Which is the payoff for the thread we've been on. You proposed earlier that the exponentially large configuration space might be redundant scaffolding — that "the actual physics doesn't need configuration space." Bell's theorem, and the 2015 loophole-free experiments, are the experimental test of exactly that proposition, and it came back negative. A classical wave in ordinary 3-space *is* the theory in which the configuration space has been dispensed with. It caps at 2. Nature gives 2.82.

The tensor product isn't bookkeeping. It's the thing being measured.

### Pilot wave has both position and momentum well-defined

**Position is the only beable. Momentum is not an independent variable — but it *is* continuously defined, because it's derived from the wavefunction.** That's the setup for the tension you're pointing at, and it's a real one with a real answer.

### What actually exists in the theory

Bohmian mechanics has exactly two things:

1. **Particle positions** $Q_k(t) \in \mathbb{R}^3$, one per particle — the configuration $Q = (Q_1,\dots,Q_N) \in \mathbb{R}^{3N}$.
2. **The wavefunction** $\psi(q_1,\dots,q_N,t)$, which lives on the $3N$-dimensional configuration space and obeys the ordinary Schrödinger equation.

(Note the split, which is the wavefunction-realism problem from earlier in this conversation showing up as an ontological commitment: particles in 3-space, wave on configuration space.)

Momentum is *not* on that list. But the **guidance equation** assigns a velocity:

$$\frac{dQ_k}{dt} = \frac{\hbar}{m_k}\,\mathrm{Im}\!\left[\frac{\nabla_k \psi}{\psi}\right]_{q=Q} \;=\; \frac{\nabla_k S(Q,t)}{m_k}, \qquad \psi = Re^{iS/\hbar}$$

This is a first-order ODE, so given $\psi$ and an initial configuration the trajectory is determined, smooth, and unique (global existence proved by Berndl et al.; the only trouble spots are nodes where $\psi = 0$ and the velocity field is singular, and those are measure-zero for initial conditions).

So yes: **each particle has a definite position and a definite velocity at every instant, both continuously defined.** Your question is exactly the right one.

### Why this doesn't violate uncertainty

Three independent things are going on, and all three matter.

#### 1. Bohmian momentum ≠ measured momentum

This is the main answer. In Bohmian mechanics, **only position measurements are faithful.** Every other "measurement" is a physical interaction whose *outcome* is produced by the apparatus coupling, not read off a pre-existing value. Momentum measurement doesn't reveal $m\,dQ/dt$; it produces a number distributed as $|\tilde\psi(p)|^2$.

The cleanest illustration is the **hydrogen ground state**. $\psi_{100} \propto e^{-r/a}$ is real, so $S = \text{const}$, so $\nabla S = 0$, so:

$$\boxed{\text{The Bohmian electron in the hydrogen ground state is at rest. } v = 0 \text{ exactly.}}$$

It just sits there. Yet a momentum measurement on that state returns a spread with $\langle p^2\rangle = \hbar^2/a^2 \neq 0$ — as it must, since $\langle T\rangle = 13.6$ eV. The Bohmian momentum (exactly zero) and the measured momentum (broadly distributed) are simply **different quantities**. The same holds for every real wavefunction: all real bound states have motionless Bohmian particles.

So there's no conflict, because the theory never claims a momentum measurement reveals $m\dot Q$. This is Bohmian mechanics being explicitly **contextual**, which is also how it lives with Kochen–Specker.

#### 2. The uncertainty relation is a theorem about ψ, and BM has the same ψ

$\sigma_x \sigma_p \ge \hbar/2$ is a mathematical fact about a function and its Fourier transform. Bohmian mechanics uses the identical $\psi$ evolving by the identical Schrödinger equation, and reproduces the Born rule exactly. So it reproduces every statistical prediction of QM, uncertainty relations included, automatically.

What changes is the *reading*. In standard QM, uncertainty is often glossed as "the particle has no definite position and momentum." In BM it's a statement about **the distribution of measurement outcomes**, with no implication about what exists. The particle has both; you just can't get at them.

#### 3. Quantum equilibrium — and why "can't get at them" is principled

Why can't you beat it by cleverer measurement? Because of the **quantum equilibrium hypothesis**: the distribution of configurations over an ensemble is $\rho = |\psi|^2$. The guidance equation preserves this (equivariance), and Dürr, Goldstein & Zanghì argued it's *typical* — overwhelmingly likely given the theory's own dynamics.

Their **absolute uncertainty** result then shows that an observer made of Bohmian particles, themselves in equilibrium, cannot learn $Q$ more precisely than $|\psi|^2$ permits. The limitation is a theorem about what subsystems can know, not an axiom.

And this has a striking corollary, which I think is the most interesting thing in the neighborhood: **if you could prepare a non-equilibrium ensemble, $\rho \neq |\psi|^2$, you could violate the uncertainty relations and signal superluminally.** In Bohmian mechanics the uncertainty principle is *contingent* — a feature of the equilibrium state of the universe, not a law. Antony Valentini has built a research program on exactly this (relic non-equilibrium from the early universe, "subquantum" information processing). Whether that's a virtue or a red flag is a matter of taste, but it's a genuine structural difference from orthodox QM, and in principle testable.

### Where the two momenta do agree

Worth knowing, because it's not arbitrary: in **scattering and time-of-flight** setups, the asymptotic Bohmian velocity $Q(t)/t$ as $t\to\infty$ *does* converge to a value distributed as $|\tilde\psi(p)|^2$. So for free particles measured far downstream — which is how momentum is usually measured — the Bohmian momentum and the measured momentum coincide. The divergence appears for confined and bound states, where "momentum measurement" is a much more invasive operation.

### Two footnotes

- **Spin is not a beable.** In standard BM, only positions are; spin lives entirely in the spinor wavefunction, and a Stern–Gerlach "spin measurement" is really a position measurement of where the particle lands. Contextuality again.
- **The trajectories aren't Newtonian.** Rewriting the guidance equation in second-order form gives $m\ddot Q = -\nabla(V + \mathcal{Q})$ with the quantum potential $\mathcal{Q} = -\frac{\hbar^2}{2m}\frac{\nabla^2 R}{R}$. That extra term is what produces interference fringes from definite trajectories — and it depends on the *shape* of $\psi$, not its magnitude, which is where the nonlocality enters.

### Summary

Position and momentum are both continuously and definitely defined along every Bohmian trajectory. The uncertainty principle survives because (a) the momentum a measurement returns is not the momentum the particle has, (b) the statistical predictions come from $|\psi|^2$ and $|\tilde\psi|^2$, which BM reproduces exactly, and (c) quantum equilibrium makes the configuration unknowable beyond $|\psi|^2$ as a theorem rather than a postulate. The price, as always with Bohm, is paid elsewhere: the guidance equation for particle $k$ depends instantaneously on the positions of all the others, so the theory is explicitly nonlocal — which is precisely the bill that Bell's theorem, from the last exchange, says every realist wave theory has to pay.

## Position basis is privileged

> what about the privileged nature of the position basis in bohmian mechanics then -- does that predict that position can be subject to entanglement or superposition? isn't that at variance with standard QM theory?


No — and the confusion is worth unpicking, because BM contains *both* things at once.

### BM has a fully superposed, fully entangled wavefunction

The wavefunction in Bohmian mechanics is **the same $\psi$ as in standard QM**, obeying the same Schrödinger equation, with the same superpositions and the same entanglement in position. Nothing is removed. What BM *adds* is a definite configuration $Q$ riding along, guided by that $\psi$.

So the ontology has two levels:

| | superposed? | entangled? |
|---|---|---|
| $\psi(q_1,\dots,q_N,t)$ on $\mathbb{R}^{3N}$ | **yes** — exactly as in QM | **yes** — exactly as in QM |
| configuration $Q \in \mathbb{R}^{3N}$ | no — one definite point | n/a |

"Position is privileged" means *position is the variable that gets a definite value*, not *position can't be superposed*.

### In fact BM **requires** position superposition to work

Double slit: the wavefunction passes through **both** slits — a superposition in position — while the particle passes through one. The interference pattern is produced by the quantum potential $\mathcal{Q} = -\frac{\hbar^2}{2m}\frac{\nabla^2 R}{R}$, which depends on the *shape* of the superposed $\psi$. Close one slit and the wave changes, the quantum potential changes, and the trajectories redistribute.

**If position couldn't be superposed, Bohmian mechanics would predict no interference at all.** The superposition isn't tolerated by the theory; it's the engine.

Same for entanglement. The EPR–Bohm correlations in BM come entirely from $\psi$ being entangled on configuration space: the guidance equation for particle 1,

$$\dot Q_1 = \frac{\nabla_1 S(Q_1, Q_2, t)}{m_1}$$

depends on $Q_2$ *because* $\psi$ doesn't factorize. **The nonlocality of Bohmian mechanics just is the entanglement of the wavefunction in position.** Take that away and you'd lose the thing that lets BM pass Bell tests.

### So what does "privileged basis" actually mean?

Three specific things, none of which is "position isn't superposed":

1. **Only position is a beable.** The extra variable is the configuration; nothing else gets a definite value.
2. **All measurements reduce to position measurements.** Spin, momentum, energy — each is a position measurement of a pointer, with the outcome produced contextually by the interaction.
3. **The dynamics is written in the position representation.** But $\psi$ itself is still just a Hilbert space vector — you can Fourier transform it freely, and BM doesn't care.

### Is it at variance with standard QM?

**Empirically, no — and this is a theorem, not a hope.** The *equivariance* result: if $\rho = |\psi|^2$ at one time, the guidance equation preserves it for all later times. Combined with quantum equilibrium, non-relativistic BM reproduces every prediction of non-relativistic QM exactly. It isn't a rival theory; it's the same theory with an added ontology.

The only way it could differ is quantum non-equilibrium ($\rho \neq |\psi|^2$), which is Valentini's program and has never been observed.

**Interpretively, yes** — BM asserts a definite position between measurements, which Copenhagen-flavored QM declines to assert. That's a real difference in what's claimed to exist, but not a contradiction with the formalism, since BM contains the whole formalism.

### Two reasons the privileging is less arbitrary than it sounds

- **Bell's observation:** in the last analysis, every actual measurement result is a position — a pointer deflection, a spot on a screen, a localized detector click, ink on paper. BM arguably just makes explicit what was already operationally true of every laboratory outcome.
- **Decoherence einselects position anyway.** Zurek's pointer-basis analysis shows that for macroscopic objects the environmentally stable basis *is* approximately position — because interactions in nature are position-dependent, $V(x_i - x_j)$. So position's special status is selected by the form of the Hamiltonian, not chosen by fiat. BM promotes a dynamically favored basis to a fundamental one.

This is also how BM dissolves measurement: once branches decohere they stop overlapping in configuration space, so the branch not containing $Q$ ceases to influence the trajectory. The other branches still exist in $\psi$ — they're "empty waves" — but they're causally inert. Effective collapse without a collapse postulate.

### Where the privileging *does* cause real trouble

The superposition worry isn't the problem. These are:

1. **Relativity.** The instantaneous dependence of $\dot Q_1$ on $Q_2$ needs a preferred foliation of spacetime to be well defined. Quantum equilibrium hides it — no signalling, so no observable violation of Lorentz invariance — but the fundamental theory has a preferred slicing that relativity says shouldn't exist. This is the most serious objection to BM, and it's exactly Bell's-theorem's bill coming due.
2. **Quantum field theory.** With particle creation and annihilation, "the positions of the particles" isn't well defined. You have to pick a new beable: Bell's lattice fermion-number density, the Dürr–Goldstein–Tumulka–Zanghì Bell-type QFT with stochastic jumps, or field-configuration beables (Struyve, Valentini). The privileged position basis doesn't straightforwardly survive the move to QFT — it has to be replaced.
3. **Photons.** There's no satisfactory photon position operator (the Newton–Wigner problems), so a particle-position ontology sits badly with the EM field; bosons are usually handled with field beables instead.

### In one line

Bohmian mechanics doesn't say position can't be superposed or entangled — it keeps the superposed, entangled wavefunction in full and adds definite positions underneath it. It needs the superposition (for interference) and it needs the entanglement (for Bell correlations). The genuine costs of privileging position are the preferred foliation and the awkward extension to QFT, not any conflict with superposition.

## Photons don't have position operator, how do we measure them?

> if every measurement is in position space, and photons don't have a position operator, then how can we measure quantum properties of photons? and what are the best references for the position operator issues of the photon?

Two answers: the measurement one is simpler than it looks, and the no-go theorems are older and sharper than most treatments suggest.

### How we measure photons without a photon position operator

**A photodetector click is not a position measurement of the photon.** It's an excitation event in the *detector* — a photoelectron ejected from a cathode, an avalanche in an APD, a silver grain reduced in film, an atom ionized. Those are all matter degrees of freedom with perfectly good position operators. What's localized is the detector's response, and the detector is made of electrons.

Glauber's photodetection theory makes this precise: the click rate at a detector at $\mathbf{r}$ is

$$w^{(1)}(\mathbf r, t) \;\propto\; \big\langle \hat E^{(-)}(\mathbf r,t)\,\hat E^{(+)}(\mathbf r,t)\big\rangle$$

a normally-ordered correlation of **local field operators** at the detector's location. Note what this is *not*: it is not $|\langle \mathbf r|\psi\rangle|^2$ for any photon position eigenstate. It's a field expectation, and field operators $\hat E(x)$ are perfectly local — they're the local observables of QED, attached to spacetime regions, no problem at all.

That's the whole resolution, and it's worth stating sharply:

> **The electromagnetic field is local. The photon is not localizable.** These are consistent because "photon" is a global/asymptotic particle concept, while the field operators are the local ones.

This is the same structure we hit at the very start of this conversation — no local number operator, type III$_1$ algebras, particle number not being a local property. The photon position problem is the most acute special case of that general fact, not a separate pathology.

**Second point:** the modern formalism replaces the projection-valued measure with a **POVM**. There is no *sharp* (projective, covariant) localization observable for the photon — but there are perfectly good *unsharp* ones, and those describe real detectors, which have finite spatial resolution anyway. Amrein constructed these for massless particles in 1969; Kraus, and Busch–Grabowski–Lahti, developed the general unsharp-localization framework. Nothing operational is lost.

**Third point, and the practical one:** almost nothing we measure about photons is position. Polarization (a clean 2-dimensional internal space with ordinary self-adjoint operators), photon number $\hat n$, quadratures $\hat X, \hat P$, phase, time-bin, orbital angular momentum. **Every photonic Bell test uses polarization or time-bin encoding**, not position. The no-go theorems never come into contact with the experiments.

And Bell's dictum survives intact: the pointer that finally has a definite position is the *apparatus*, which is matter. Photons carry the correlations; electrons record them.

### What exactly goes wrong

Four distinct results, often conflated:

1. **Pryce (1948).** The natural candidate photon position operators have **non-commuting components**, $[\hat x_i, \hat x_j] \neq 0$, so they fail a basic requirement for a position observable.
2. **Newton & Wigner (1949).** Localized states are uniquely determined for elementary systems of **nonzero** mass and any spin. The construction fails for massless particles of helicity $> \tfrac12$ — *given* that the localized states are assumed spherically symmetric. That caveat is load-bearing; it's the crack Hawton's program works in.
3. **Wightman (1962).** The canonical theorem: no covariant projection-valued localization observable exists for massless particles of helicity $\ge 1$. This is the one to cite.
4. **Hegerfeldt (1974, 1998).** Any state with **positive energy** that is strictly localized in a bounded region instantly develops tails everywhere. Crucially this needs neither relativity nor translation invariance — positivity of the energy spectrum alone suffices. So strict localization is in tension with positive energy *generally*, for massive particles too. This one isn't a photon problem at all; it's a quantum theory problem.

The physical root causes for the photon specifically: transversality ($\nabla\cdot\mathbf E = 0$) is a **nonlocal constraint** on the mode space, and masslessness means there's no rest frame in which spin separates covariantly from orbital angular momentum.

---

# References

### The classic no-go results

- **A. S. Wightman**, "On the Localizability of Quantum Mechanical Systems," [*Rev. Mod. Phys.* **34**, 845 (1962)](https://link.aps.org/doi/10.1103/RevModPhys.34.845). **Start here for the theorem.** The definitive statement, including the helicity $\ge 1$ obstruction.
- **T. D. Newton & E. P. Wigner**, "Localized States for Elementary Systems," [*Rev. Mod. Phys.* **21**, 400 (1949)](https://link.aps.org/doi/10.1103/RevModPhys.21.400). The origin of the whole subject.
- **M. H. L. Pryce** (1948), on the mass-centre in relativistic field theory — where the non-commuting components were first found.
- **W. O. Amrein**, "Localizability for Particles of Mass Zero," *Helv. Phys. Acta* **42**, 149–190 (1969). The massless case specifically, **and** the constructive POVM alternative. Under-cited and directly on your question.
- **G. C. Hegerfeldt**, "Instantaneous spreading and Einstein causality in quantum theory," [*Ann. Phys. (Leipzig)* **7**, 716 (1998)](https://onlinelibrary.wiley.com/doi/10.1002/andp.199851007-817), and the original 1974 paper. Read this to see that the problem is broader than photons.

### Best modern reviews

- **I. Bialynicki-Birula**, "Photon Wave Function," *Progress in Optics* **36**, ed. E. Wolf (Elsevier, 1996), p. 245. **The best entry point** — clear, balanced, and constructive about what a photon wavefunction can and cannot be (the Riemann–Silberstein formulation).
- **O. Keller**, "On the theory of spatial localization of photons," [*Phys. Rep.* **411**, 1 (2005)](https://ui.adsabs.harvard.edu/abs/2005PhR...411....1K/abstract). **The most comprehensive treatment**, ~230 pages, with the near-field and light–matter-interaction angle.
- **A. Peres & D. R. Terno**, "Quantum Information and Relativity Theory," [*Rev. Mod. Phys.* **76**, 93 (2004)](https://arxiv.org/pdf/quant-ph/0212023v2). The section on photon localization is the crispest short statement of the operational/POVM resolution.
- **I. Bialynicki-Birula & Z. Bialynicka-Birula**, "Why photons cannot be sharply localized," *Phys. Rev. A* **79**, 032112 (2009).

### The dissenting constructive line (still live)

- **M. Hawton**, "Photon position operator with commuting components," *Phys. Rev. A* **59**, 954 (1999); Hawton & Baylis, *Phys. Rev. A* **64**, 012101 (2001). Constructs an operator with commuting components by relaxing the spherical-symmetry/rotational-covariance assumption that Newton–Wigner built in.
- The live dispute: [arXiv:2203.14555](https://arxiv.org/abs/2203.14555) (a comment arguing Hawton's operator lacks the required properties) and [arXiv:2205.04791](https://arxiv.org/pdf/2205.04791) ("Construction of a photon position operator with commuting components from natural axioms," which finds a canonical operator unitarily equivalent to Hawton's). Read these two together — they show exactly which axiom each side is willing to drop.
- **Hodgson, Southall, Purdy & Beige**, ["Local Photons"](https://www.frontiersin.org/journals/photonics/articles/10.3389/fphot.2022.978855/full), *Front. Photonics* (2022), for a recent take.

### The measurement side

- **R. J. Glauber**, "The Quantum Theory of Optical Coherence," *Phys. Rev.* **130**, 2529 (1963). The photodetection theory that makes "what a detector actually measures" precise. (Nobel, 2005.)
- **L. Mandel**, "Configuration-Space Photon Number Operators in Quantum Optics," *Phys. Rev.* **144**, 1071 (1966). The coarse-grained photon number operator for a finite volume — valid when the volume is large compared to a wavelength, which is precisely the honest statement of how far localization goes.
- **Mandel & Wolf**, *Optical Coherence and Quantum Optics* (CUP, 1995), Ch. 12 — textbook treatment.
- **Busch, Grabowski & Lahti**, *Operational Quantum Physics* — the general POVM/unsharp-observable framework.

If you read three: **Wightman** for the theorem, **Bialynicki-Birula** for the physics, **Keller** for everything else.

**Sources:** Wightman (1962), Newton–Wigner (1949), Amrein (1969), Bialynicki-Birula (1996), Keller (2005), Peres–Terno (2004), Hegerfeldt (1998), and the Hawton controversy papers were verified against journal/arXiv listings just now. Volume/page numbers for Pryce (1948), Glauber (1963), Mandel (1966), Bialynicki-Birula PRA **79** (2009), and Hawton PRA **59** (1999) are from memory — the papers and their content are confirmed, but double-check those specific page numbers if you're citing them formally.

## Field-based Bell model

> given that the photon is best represented as a distributed field-based entity, not some kind of localizable particle-like entity, has anyone actually tried to construct a configuration space of two such distributed photons in a Bell-type experiment? how would the measurement of polarization of one of these actually have the appropriate effects on the distributed field-level state of the other photon, to produce the expected correlations at the other detector?

Yes — completely standard, and the object has a name: the **two-photon amplitude**, or "biphoton wavefunction." It's routinely computed and routinely measured. But the way it evades the position-operator problem is the interesting part, and there's a deeper version of your question with a startling answer.

### The construction

Start in momentum–helicity space, where everything is clean. An SPDC source produces

$$|\psi\rangle = \sum_{\lambda_1\lambda_2}\int\! d^3k_1\, d^3k_2\; C(\mathbf k_1\lambda_1;\mathbf k_2\lambda_2)\; \hat a^\dagger_{\mathbf k_1\lambda_1}\hat a^\dagger_{\mathbf k_2\lambda_2}|0\rangle$$

**That joint amplitude $C$ is your configuration space.** It's a function on (wavevector ⊗ helicity)², i.e. the $\mathrm{Sym}^2$ sector of photon Fock space — precisely the exponential/tensor-product structure from the start of this conversation. Nothing about the photon being "distributed" interferes with it, because the construction never asked for photon positions.

To get a position-space object, use **local field operators** instead of position eigenstates:

$$\Psi(\mathbf r_1,t_1;\mathbf r_2,t_2) \;=\; \langle 0|\,\hat E^{(+)}(\mathbf r_2,t_2)\,\hat E^{(+)}(\mathbf r_1,t_1)\,|\psi\rangle$$

This is a genuine function of two spacetime points — a two-photon configuration space on $\mathbb{R}^3\times\mathbb{R}^3$. And the coincidence rate at two detectors is Glauber's second-order correlation:

$$G^{(2)}(1,2) = \big|\Psi(1,2)\big|^2$$

**This is the key move.** $\Psi$ is a *detection amplitude*, not a photon position wavefunction. It's built from $\hat E^{(+)}$, which is a perfectly good local operator, so Wightman's theorem never bites. You get a Born rule on a two-point configuration space without ever positing a photon position operator.

And it's not formal: this object is measured directly in ghost imaging, Hong–Ou–Mandel interferometry, and transverse-correlation experiments. Law and Eberly computed its Schmidt decomposition and found transverse Schmidt numbers in the hundreds — the biphoton is *enormously* entangled in its spatial degrees of freedom alone.

### The wrinkle your question anticipates

Here's where "distributed" actually bites, and it's a real experimental problem rather than a philosophical one. The full state is

$$|\Psi\rangle = \tfrac{1}{\sqrt2}\Big(|H\rangle_1|V\rangle_2\,\big|\phi_{HV}(\mathbf r_1,\mathbf r_2)\big\rangle \;-\; |V\rangle_1|H\rangle_2\,\big|\phi_{VH}(\mathbf r_1,\mathbf r_2)\big\rangle\Big)$$

If $\phi_{HV} \neq \phi_{VH}$ — if the spatial/spectral distributions differ between the two polarization terms — then **the distributed field degrees of freedom carry which-path information**, and the polarization entanglement is degraded or destroyed. The two terms no longer interfere.

This is exactly what happens in type-II SPDC because of birefringent walk-off, and it's why real sources need compensating crystals (Kwiat et al.'s 1995 design). You have to *engineer the distributed parts to be identical* so that tracing them out leaves a clean polarization Bell state:

$$\phi_{HV} = \phi_{VH} = \phi \quad\Longrightarrow\quad |\Psi\rangle = |\Psi^-\rangle_{\text{pol}} \otimes |\phi\rangle_{\text{spatial}}$$

So the field-distributed nature isn't neutral scenery. It's an active threat to the entanglement that experimentalists spend real effort neutralizing.

### How measuring one "affects" the other

Three statements, in order of importance:

**1. Nothing propagates. This is a theorem.** QED satisfies **microcausality**: $[\hat E_i(x), \hat E_j(y)] = 0$ for spacelike-separated $x,y$. Alice's polarizer and detector are built from operators that commute with everything at Bob's location. Bob's reduced density matrix, and therefore every local observable and every single-detector count rate $G^{(1)}(\mathbf r_2)$, is **provably unchanged** by anything Alice does — including her choice of setting. No-signalling isn't an assumption bolted on; it falls out of the field commutators.

**2. What changes is the conditional description, and nothing spatial changes at all.** Take the clean state above. Alice projects photon 1 onto $|\alpha\rangle = \cos\alpha|H\rangle + \sin\alpha|V\rangle$:

$$\langle\alpha|_1\,|\Psi^-\rangle \;\propto\; \cos\alpha\,|V\rangle_2 - \sin\alpha\,|H\rangle_2 \;=\; |\alpha^\perp\rangle_2$$

So the conditional amplitude for photon 2 is

$$\Psi_{\text{cond}}(\mathbf r_2) \;\propto\; \phi(\cdot,\mathbf r_2)\otimes|\alpha^\perp\rangle$$

**Look at what did and didn't change.** The spatial mode function $\phi$ is *untouched* — photon 2's distribution over space is identical before and after. The only thing that updates is which polarization component of the already-existing joint amplitude survives the conditioning. There is no mechanism needed for "reaching out and twisting a distributed field." The correlation was written into $C(\mathbf k_1\lambda_1;\mathbf k_2\lambda_2)$ at the crystal.

**3. So what's nonlocal?** Not any influence, and not the field. What's nonlocal is that $C$ is a **single function on the product space that does not factorize** — exactly the $\det C \neq 0$ condition from the CHSH calculation, now living in an infinite-dimensional mode space. Bell's theorem forbids *explaining* the resulting correlations by locally pre-assigned values. It does not require anything to travel.

### The distinction that resolves the apparent paradox

QED is simultaneously:

- **Local** in the microcausality sense — spacelike operators commute, no signalling, exactly.
- **Nonlocal** in the Bell sense — correlations are not locally causal, $P(a,b|\alpha,\beta,\lambda) \neq P(a|\alpha,\lambda)P(b|\beta,\lambda)$.

These are different predicates and both hold. Most confusion about "how does the photon know" dissolves once you see that the second kind of nonlocality requires no carrier for the first kind.

### The deepest version of your question

You asked whether anyone has done this properly at the field level, with distributed entities rather than particles. They have — and the answer is stronger than you'd expect.

**Summers and Werner proved that the *vacuum* of a relativistic quantum field theory violates Bell's inequalities — maximally, and generically** — for local algebras associated with spacelike-separated regions. No photons. No source. No pair production. The field's ground state already contains maximally Bell-violating correlations across every spacelike cut.

This is the algebraic counterpart of the Reeh–Schlieder theorem (Witten's notes, cited earlier): vacuum entanglement between adjacent regions is not merely present but UV-divergent. So the question "how do two distributed field entities come to be correlated" has a somewhat deflationary answer: **in QFT, being correlated across spacelike separation is the generic condition of the field, and the biphoton experiment is just a particularly clean way to harvest and address it.**

---

### References

**The two-photon amplitude, worked out explicitly**
- **Walborn, Monken, Pádua & Souto Ribeiro**, "Spatial correlations in parametric down-conversion," [*Phys. Rep.* **495**, 87 (2010)](https://arxiv.org/abs/1010.1236). **The best single reference for your question** — the transverse two-photon amplitude in full, with Bell violations and imaging applications.
- **Law & Eberly**, "Analysis and Interpretation of High Transverse Entanglement in Optical Parametric Down Conversion," [*PRL* **92**, 127903 (2004)](https://link.aps.org/doi/10.1103/PhysRevLett.92.127903). Schmidt decomposition of the biphoton — how entangled the distributed part actually is.
- **Smith & Raymer**, "Photon wave functions, wave-packet quantization of light, and coherence theory," [*New J. Phys.* **9**, 414 (2007)](https://arxiv.org/pdf/0708.0831). Multi-photon wavefunctions done carefully, including the two-photon case.
- **Rubin, Klyshko, Shih & Sergienko**, "Theory of two-photon entanglement in type-II optical parametric down-conversion," *Phys. Rev. A* **50**, 5122 (1994) — the source theory.
- **Glauber** (1963) for why $G^{(2)}$ is the right object.

**Bell violation at the field level**
- **Summers & Werner**, "The vacuum violates Bell's inequalities," *Phys. Lett. A* **110**, 257 (1985); "Maximal violation of Bell's inequalities is generic in quantum field theory," *Commun. Math. Phys.* **110**, 247 (1987); and ["Bell's inequalities and quantum field theory. I. General setting,"](https://pubs.aip.org/aip/jmp/article/28/10/2440/1031086/Bell-s-inequalities-and-quantum-field-theory-I) *J. Math. Phys.* **28**, 2440 (1987). **This is the answer to "has anyone done it properly."**
- **Witten**, [arXiv:1803.04993](https://arxiv.org/abs/1803.04993) — Reeh–Schlieder and why vacuum entanglement is unavoidable.

**Sources:** Walborn et al., Law & Eberly, and the Summers–Werner papers were verified against journal/arXiv listings just now. The Rubin–Klyshko–Shih–Sergienko and Kwiat et al. (1995) volume/page numbers are from memory — content confirmed, citations worth double-checking if you're using them formally.

## Isolated detectors at A and B

> so why exactly can you isolate A and B detectors as separable subspaces? just because they were never entangled? but why don't they then *become* entangled by operating on this entangled two-photon state?


You've put your finger on a real conflation — and then asked exactly the right follow-up. Let me separate the two senses of "separable," because the answer is that **they do become entangled, and it doesn't matter.**

### Two different things called "separable"

| | meaning | which one Bell needs |
|---|---|---|
| **separable subspaces** | the Hilbert space *factorizes*, $\mathcal H = \mathcal H_A \otimes \mathcal H_B$ — a statement about degrees of freedom | ✅ this one |
| **separable state** | the state is a product / unentangled mixture — a statement about $\rho$ | ❌ explicitly false here |

The factorization is **kinematic**. It says Alice's degrees of freedom and Bob's are *different degrees of freedom*, which is true whether or not they're correlated. The whole point of the experiment is that the state on that factorized space is maximally entangled. So "you can factorize" was never a claim that nothing is entangled.

### They do become entangled — here's the chain

When Alice's detector interacts with photon 1, the von Neumann chain runs:

$$\tfrac{1}{\sqrt2}\big(|H\rangle_1|V\rangle_2 - |V\rangle_1|H\rangle_2\big)\otimes|D_A^{0}\rangle \;\xrightarrow{\;U_A\;}\; \tfrac{1}{\sqrt2}\big(|H\rangle_1|V\rangle_2|D_A^{H}\rangle - |V\rangle_1|H\rangle_2|D_A^{V}\rangle\big)$$

After this, **Alice's detector is entangled with Bob's photon.** Fully. Then it entangles with her amplifier, her computer, the air in her lab, her retina. Decoherence propagates it outward. All true, all standard, and none of it breaks anything.

### Why it doesn't matter: local unitaries cannot create entanglement across the cut

$U_A$ acts on $\mathcal{H}_1 \otimes \mathcal{H}_{D_A}$ — both inside Alice's factor. So on the full space it's $U_A \otimes I_B$. And:

$$\rho_B' = \operatorname{Tr}_A\!\left[(U_A\otimes I)\,\rho\,(U_A^\dagger\otimes I)\right] = \operatorname{Tr}_A\!\left[\rho\, (U_A^\dagger U_A\otimes I)\right] = \operatorname{Tr}_A[\rho] = \rho_B$$

by cyclicity of the partial trace. **Bob's reduced state is exactly invariant** — same computation that gives no-signalling. And since the entanglement entropy across the cut is $S(\rho_B)$, the A|B entanglement is *exactly conserved*.

So what actually happened: Alice's detector **inherited** correlation that photon 1 already had with photon 2. It didn't create new A–B entanglement; it redistributed existing entanglement *within* Alice's factor, from "photon 1" to "photon 1 + detector + lab." The total across the cut is untouched.

That's the answer to your question. Entanglement with Bob's side gets passed around inside Alice's side freely. It never grows, because growing it would require an operation spanning the cut — which spacelike separation forbids.

### So the factorization survives, trivially

Just make $\mathcal{H}_A$ bigger:

$$\mathcal{H} = \underbrace{\big(\mathcal H_{\text{photon 1}}\otimes \mathcal H_{\text{polarizer}}\otimes\mathcal H_{\text{detector}}\otimes\mathcal H_{\text{lab}}\big)}_{\mathcal{H}_A}\;\otimes\;\underbrace{\big(\cdots\big)}_{\mathcal{H}_B}$$

Alice's outcome, however complicated the apparatus, is still represented by an operator of the form $M_A \otimes I_B$. CHSH needs only two things: a factorization by *location*, and observables that are local with respect to it. Neither is disturbed by entanglement.

### What the assumption really is

Stated precisely, it isn't "A and B are unentangled." It's:

$$\big[\mathcal{A}(O_A),\,\mathcal{A}(O_B)\big] = 0 \qquad \text{for spacelike-separated regions } O_A, O_B$$

**Microcausality** — a theorem of relativistic QFT, not a hypothesis about states. Everything Alice can build sits in $\mathcal{A}(O_A)$; everything in $\mathcal{A}(O_B)$ commutes with it.

### Where the factorization genuinely does fail

Your instinct that something is shaky here is right — just not at the entanglement step. Three places it actually breaks:

**1. Identical particles.** Photons are bosons; "photon 1" and "photon 2" are not genuine subsystems. The state space is $\mathrm{Sym}^2(\mathfrak h)$, not $\mathfrak h \otimes \mathfrak h$, and the particle-labelled factorization is a fiction that only works when the wavepackets are in distinguishable modes. **The honest factorization is by *mode*, not by particle** — the modes propagating toward Alice versus those toward Bob. That's what "isolating A and B" actually means, and it's why the analysis is valid even though the photons are indistinguishable. (There's a real literature on entanglement of identical particles — Ghirardi & Marinatto, Wiseman & Vaccaro — precisely about not getting this wrong.)

**2. In QFT, the tensor factorization strictly doesn't exist.** As we discussed at the start of this conversation, local algebras are **type III$_1$ factors**, which admit *no* tensor product decomposition. There is no "Alice's Hilbert space" in rigorous QFT. What survives is exactly the commuting-algebra statement above — which is all Bell's argument ever needed.

This is precisely why **Summers and Werner had to reformulate Bell's theorem algebraically**: the naive $\mathcal{H}_A \otimes \mathcal{H}_B$ isn't available in relativistic field theory, so they worked with commuting von Neumann algebras instead. Your question is the reason that paper exists.

**3. Gauge theories.** Gauss's law obstructs factorizing the physical Hilbert space by region — edge modes, and the well-known ambiguities in defining entanglement entropy in gauge theories.

### The closing thought

Alice's detector entangling with Bob's photon isn't a bug — it's the *measurement*. Without it there'd be no record and no outcome. And notice the final step: the CHSH value is computed from the **joint statistics of the two records**, which requires physically bringing the records together at subluminal speed. That's the last piece of why all this nonlocal correlation never becomes a signal. The entanglement spreads freely inside each wing; the correlation only becomes visible where the light cones meet.

## Bell ref on privileged position basis

> what is the best reference for Bell's position space point?

**J. S. Bell, "On the impossible pilot wave," *Foundations of Physics* **12**, 989–999 (1982)** — the passage is on **p. 996**, Section 3 ("Morals"), as the second of three morals. Reprinted as Chapter 17 of *Speakable and Unspeakable in Quantum Mechanics* (Cambridge, 1987), pp. 159–168.

The exact wording, from the paper:

> "The second moral is that in physics the only observations we must consider are position observations, if only the positions of instrument pointers. It is a great merit of the de Broglie-Bohm picture to force us to consider this fact. If you make axioms, rather than definitions and theorems, about the 'measurement' of anything else, then you commit redundancy and risk inconsistency."

Free full text: [unicamp.br/~chibeni/textosdidaticos/bell-1982.pdf](https://unicamp.br/~chibeni/textosdidaticos/bell-1982.pdf)

### One correction to what I told you earlier

I rendered this as "in the last analysis, all measurements are position measurements." That's a common paraphrase but it overstates Bell. His actual claim is **"the only observations we *must* consider"** — a claim about *sufficiency*, not about metaphysics. He isn't saying nothing else is ever measured; he's saying you never need to *axiomatize* anything beyond position, and that axiomatizing more invites redundancy and inconsistency. That's a weaker and more defensible thesis, and worth quoting exactly if you're going to use it.

The immediately following paragraph is the famous complaint about the word "measurement" itself — which he expanded into **"Against 'measurement'," *Physics World* **3**(8), 33 (1990)**, reprinted as ch. 23 of the second edition of *Speakable and Unspeakable*. That's the better citation if what you want is the critique of measurement-as-primitive rather than the position point specifically.

### The technical follow-ups

Bell states the moral but doesn't develop it. The people who did:

- **Dürr, Goldstein & Zanghì, "Quantum Equilibrium and the Role of Operators as Observables in Quantum Theory,"** [arXiv:quant-ph/0308038](https://arxiv.org/abs/quant-ph/0308038), *J. Stat. Phys.* **116**, 959 (2004). The full technical development: how POVMs and self-adjoint operators *emerge* from position measurements on apparatus, rather than being postulated. This is the paper that turns Bell's moral into a theorem.
- **Dürr, Goldstein & Zanghì, "Naive Realism About Operators,"** [arXiv:quant-ph/9601013](https://arxiv.org/html/quant-ph/9601013v1), *Erkenntnis* **45**, 379 (1996). The shorter, more readable polemical version — on the error of thinking an operator-observable names a property being revealed.
- **Goldstein**, ["Bohmian Mechanics"](https://plato.stanford.edu/entries/qm-bohm/), *Stanford Encyclopedia of Philosophy*, §on measurement — the best free overview, and it cites the Bell passage in exactly this context.

For the photon case we were discussing, the relevant tension is that Bell's moral was formulated for non-relativistic particle mechanics, where positions are beables. Extending it to QED is precisely where the position-operator problem and the Bell-type QFT constructions (Dürr–Goldstein–Tumulka–Zanghì) come in.

## Order of measurements in Bell

> Isn't there a basic ambiguity in the Bell tests about which measurement happens first? are there incompatible angle settings where it actually matters which one goes first, or given the ultimate dependence on the density matrix, this really doesn't matter? but logically, according to the contextual interpretation, if the measurement is actually changing the entire wave state to the new basis associated with a measurement, it seems like there should be cases where this matters? and especially if it would somehow be simultaneous -- what would that look like?

Sharp question, and it has a clean answer plus one piece that genuinely *is* unresolved — just not where you'd expect.

### The order doesn't matter, and it's exact

Three ways to compute the joint outcome probability for the Bell state, Alice at 0°, Bob at 45°:

| outcome | Alice measures first | Bob measures first | single joint PVM |
|---|---|---|---|
| $(+,+)$ | 0.426776695 | 0.426776695 | 0.426776695 |
| $(+,-)$ | 0.073223305 | 0.073223305 | 0.073223305 |
| $(-,+)$ | 0.073223305 | 0.073223305 | 0.073223305 |
| $(-,-)$ | 0.426776695 | 0.426776695 | 0.426776695 |

Identical to every digit. The reason, checked numerically:

$$\big\|\,[\,P^A_a\otimes I,\; I\otimes P^B_b\,]\,\big\| = 0.0$$

And the one-line proof, using the Lüders rule with Alice first:

$$p(a,b) = \langle\psi|(P^A_a\otimes I)(I\otimes P^B_b)(P^A_a\otimes I)|\psi\rangle \;\overset{\text{commute}}{=}\; \langle\psi|(P^A_a\otimes I)^2(I\otimes P^B_b)|\psi\rangle \;\overset{P^2=P}{=}\; \langle\psi|P^A_a\otimes P^B_b|\psi\rangle$$

Swapping the roles gives the same expression. Both sequential orders collapse to the symmetric joint projector — which is exactly the $\Pi_{ab} = P^A_a \otimes P^B_b$ I used in the original CHSH derivation without comment. So the "simultaneous" description isn't a fourth case; it's the thing both orders reduce to.

**This has to be true, and it's why microcausality exists.** For spacelike separation, different inertial frames genuinely disagree about which measurement came first — there is no frame-invariant fact about the order. If predictions depended on order, QM would contradict relativity outright. The commutation of spacelike-separated observables is precisely the structural feature that prevents it.

### Your basis-change intuition is right — it just lives inside each wing

Non-commuting measurements on the *same* system are emphatically order-dependent. Starting from $|0\rangle$:

| sequence | $(+,+)$ | $(+,-)$ | $(-,+)$ | $(-,-)$ | $P(\text{first} = +)$ |
|---|---|---|---|---|---|
| $\sigma_z$ then $\sigma_x$ | 0.50 | 0.50 | 0.00 | 0.00 | **1.00** |
| $\sigma_x$ then $\sigma_z$ | 0.25 | 0.25 | 0.25 | 0.25 | **0.50** |

Wildly different, because $\|[P_{z+}, P_{x+}]\| = 0.5 \neq 0$. So the phenomenon you're imagining is completely real.

It just never occurs across the A|B cut in a Bell test, because of how the experiment is structured:

- Alice's two settings $A$ and $A'$ **are** mutually non-commuting — $[\sigma_z, \sigma_x] \neq 0$ — but she performs **exactly one of them per run**. She never sequences them.
- Alice's observable and Bob's observable **always commute**, regardless of their angles. $A(\alpha)\otimes I$ and $I \otimes B(\beta)$ commute for *every* $\alpha, \beta$ — it's automatic from the tensor structure, not a special property of particular angles.

So there are no "incompatible angle settings where it matters." Incompatibility of angles is a within-wing property; the cross-wing commutator is identically zero for all angles.

### The elegant version

Look again at the Landau identity from the CHSH derivation:

$$\hat S^2 = 4\,I + [A,A']\otimes[B,B']$$

The only commutators that appear anywhere are $[A,A']$ and $[B,B']$ — **each one entirely inside a single wing**. The cross-wing commutator never appears, because it's zero. So the two facts do two different jobs:

- **Non-commutativity *within* each wing** is what lifts $S$ from 2 to $2\sqrt2$. It powers the violation.
- **Commutativity *across* the wings** is what makes the whole thing order-independent and relativistically consistent.

Remove the first and you get a classical theory. Remove the second and you get signalling and a contradiction with relativity. The Bell violation lives precisely in the gap between them.

### Where the ambiguity is genuinely real

Here's the part where your instinct lands on something unresolved. **The statistics are frame-independent. The collapse narrative is not.**

If you treat state reduction as a physical process occurring on a spacelike hypersurface, then in Alice's frame the state collapses at Alice's measurement event, and in Bob's frame at Bob's. Those are two different histories of the state, with two different intermediate states, agreeing only on the final joint statistics.

Aharonov and Albert worked this out in detail, and their conclusion is worth stating exactly: the covariance of relativistic quantum theory resides **exclusively in the experimental probabilities, and not in the underlying quantum states**. There is no covariant state history.

- **Aharonov & Albert**, "Can we make sense out of the measurement process in relativistic quantum mechanics?" [*Phys. Rev. D* **24**, 359 (1981)](https://www.semanticscholar.org/paper/Can-we-make-sense-out-of-the-measurement-process-in-Aharonov-Albert/ad2fd268c671f6e41e18562eddae600b0f313214); and *Phys. Rev. D* **29**, 223 (1984).
- **Hellwig & Kraus**, "Formal description of measurements in local quantum field theory," *Phys. Rev. D* **1**, 566 (1970) — proposes collapse on the *past light cone* to restore covariance. Covariant, but with its own costs.
- **Peres & Terno**, [*Rev. Mod. Phys.* **76**, 93 (2004)](https://arxiv.org/pdf/quant-ph/0212023v2) — the section arguing there is no such thing as a covariant wavefunction; the state is bookkeeping relative to a chosen foliation.

### The moral

If the order of collapse has no frame-invariant answer *and* no empirical consequence, the natural reading is that **collapse isn't a physical event** — it's an update of a description. Every major interpretation takes this as a data point and pays differently: Everett denies collapse happens at all; relational QM makes states observer-relative by construction; Bohmian mechanics bites the bullet and adopts a preferred foliation, which is exactly the relativity bill we discussed several turns back; consistent histories restricts which narratives are admissible.

So: no ambiguity in the predictions, a real and permanent ambiguity in the story, and that asymmetry is itself one of the more informative facts in the subject.

## Unscrambling the omelette

> has anyone used the hydrodynamic quantum analogs framework (HQA) to "unscramble the omelette" of the schrodinger wave function, which includes both epistemological and ontological components. the epistemological component is basically the distribution of starting locations and other parameters, in the pilot wave / HQA framework, while the ontological component is the Compton-wavelength, physically real quantum wave generated by the particle itself. The Bohmian version doesn't distinguish these two, and the Schrodinger wave function includes a spreading diffusion component that reflects an epistemic increase in uncertainty over time. However, in the HQA model, the double-solution generation of the quantum wave by the particle itself keeps the resulting ontological wave from diffusing out like that, so it preserves a more localized wave with a stable overall density radius.

Yes — and you've essentially reconstructed de Broglie's original program from scratch. The "unscrambling" you describe is **the double solution**, and there's an active HQA-inspired implementation of exactly it.

### The program you're describing has a name

De Broglie proposed at the 1927 Solvay conference a *théorie de la double solution* with two distinct waves:

- **$u$** — a real, physically existing wave with a singularity (or soliton-like concentration) at the particle. **Ontological.** Not normalizable.
- **$\psi$** — a smooth, normalizable, statistical wave. **Epistemic.** No independent physical existence.

The two are locked by the "harmony of phases": $u$ and $\psi$ share the same phase, which is what makes the particle's motion track $\nabla S$.

He couldn't make it work and presented only the reduced pilot-wave version at Solvay — and that's the version Bohm revived in 1952. So your reading of the history is right: **Bohm inherited the collapsed form, and the split was the thing that got lost.**

### The direct HQA realization: Dagan & Bush's "hydrodynamic quantum field theory"

This is the paper you want:

**Y. Dagan & J. W. M. Bush, "Hydrodynamic quantum field theory: the free particle,"** [*Comptes Rendus Mécanique* **348**, 555–571 (2020)](https://comptes-rendus.academie-sciences.fr/mecanique/item/CRMECA_2020__348_6-7_555_0/) ([free PDF](https://thales.mit.edu/bush/wp-content/uploads/2021/04/Dagan2020.pdf)).

Its stated aim is to revisit de Broglie's double solution in light of the walking-droplet system, and its ingredients are precisely the ones you listed:

- the particle carries an **internal oscillation at the Compton frequency** ($\omega_c = mc^2/\hbar$ — de Broglie's internal clock, where rest-mass energy is exchanged with field energy);
- that oscillation **sources a real scalar field** satisfying the **Klein–Gordon equation**;
- the particle self-propels by **resonant interaction with its own generated wave**.

The follow-up works out the pilot-wave's form: **M. Durey & J. W. M. Bush, "Hydrodynamic Quantum Field Theory: The Onset of Particle Motion and the Form of the Pilot Wave,"** [*Front. Phys.* **8**, 300 (2020)](https://www.frontiersin.org/journals/physics/articles/10.3389/fphy.2020.00300/pdf) — there the internal vibration is at *twice* the Compton frequency.

**Your dispersion point is exactly right, and it was de Broglie's own motivation.** A free Schrödinger packet spreads as $\sigma(t) = \sigma_0\sqrt{1 + (\hbar t/2m\sigma_0^2)^2}$ — fine for a probability distribution, catastrophic for a physically real electron. De Broglie's response was to make $u$ **nonlinear**, so it could be a non-dispersing soliton (his 1956 *Une tentative d'interprétation causale et non linéaire de la mécanique ondulatoire*, and *Non-linear Wave Mechanics*, 1960). In HQFT the field is *driven* by a persistently oscillating source rather than freely evolving, so it settles into a quasi-steady localized form at the Compton scale. Same goal, different mechanism.

### The epistemic half is also done — and it's the more convincing half

The other side of your unscrambling is showing that the *statistics* emerge rather than being fundamental. That's the corral result:

**D. M. Harris, J. Moukhtar, E. Fort, Y. Couder & J. W. M. Bush, "Wavelike statistics from pilot-wave dynamics in a circular corral,"** [*Phys. Rev. E* **88**, 011001(R) (2013)](https://link.aps.org/doi/10.1103/PhysRevE.88.011001) ([PDF](https://math.mit.edu/sites/bush/wp-content/uploads/2013/07/Harris-Corrals-2013.pdf)).

A single droplet walks chaotically in a corral; the histogram of its position over long times converges to a distribution prescribed by the Faraday wave mode of the cavity. So you get a QM-like probability density as an **emergent, ergodic property of deterministic chaotic dynamics** — not a postulate. That's the epistemic component derived rather than assumed, which is precisely the structural role $|\psi|^2$ plays in Bohmian mechanics as quantum equilibrium.

### Sharpening your claim about Bohm

One refinement worth making, because it makes your point stronger rather than weaker. Bohmian mechanics *does* separate ontic from epistemic — just along a different cut:

| | separated? |
|---|---|
| particle configuration $Q$ (ontic) vs. ensemble distribution $\rho$ (epistemic) | **yes** — that's quantum equilibrium, and Valentini's non-equilibrium program is about prying $\rho$ apart from $\|\psi\|^2$ |
| guiding wave (ontic role) vs. probability amplitude (epistemic role) | **no** — one $\psi$ does both jobs |

So Bohm separates *particle from distribution* but **merges guiding-wave with probability-amplitude**. De Broglie's $u/\psi$ split separates both. That second merge is the one you're pointing at, and it's real.

### The hard part — and why de Broglie gave up

The difficulty was never *making* the split. It's **keeping $u$ and $\psi$ locked** so that the ensemble statistics of the $u$-guided particle dynamics actually reproduce $|\psi|^2$, in general, for all potentials and all times. De Broglie's phase-harmony condition was supposed to enforce it and he couldn't make it rigorous — which is why he abandoned the program in 1927 and didn't return to it for thirty years.

HQFT hasn't solved this either. It's a *model inspired by* the hydrodynamic system, not a derivation of quantum mechanics; it doesn't yet recover the Schrödinger equation, let alone anything field-theoretic.

### The honest limits

Three things you should weigh before investing:

1. **HQA cannot reproduce Bell violations or multiparticle entanglement.** Bush is explicit about this in the reviews. It's a *local* classical system, so per the earlier part of our conversation, this is a theorem, not a technical gap. That caps what the analogy can ever demonstrate.
2. **The diffraction result didn't replicate.** Andersen et al., "Double-slit experiment with single wave-driven particles and its relation to quantum mechanics," *Phys. Rev. E* **92**, 013006 (2015), failed to reproduce the Couder–Fort single-particle diffraction pattern. The original result is no longer regarded as robust; Pucci et al. (*JFM*, 2018) did a more careful single-slit study. This was the single most quantum-like claim in the field, and it weakened considerably.
3. **HQFT is early-stage** — suggestive structure, not a competing theory.

### References

- **Bush & Oza, "Hydrodynamic quantum analogs,"** [*Rep. Prog. Phys.* **84**, 017001 (2021)](https://thales.mit.edu/bush/wp-content/uploads/2021/04/BushOza-ROPP.pdf). **Start here** — the comprehensive review, and it explicitly frames the system as the first macroscopic realization of de Broglie's proposed dynamics.
- **Dagan & Bush (2020)** and **Durey & Bush (2020)**, above — the HQFT papers, which are the direct answer to your question.
- **Colin, Durt & Willox, "de Broglie's double solution program: 90 years later,"** [arXiv:1703.06158](https://arxiv.org/abs/1703.06158), *Ann. Fond. Louis de Broglie* **42**, 19 (2017). **The definitive review of the double-solution program itself**, including the $u$-wave / $\psi$-wave ontological-vs-statistical distinction you're asking about — done independently of HQA, which makes it a useful cross-check on whether the split is coherent on its own terms.
- **Durt**, ["Revisiting de Broglie's double solution program"](https://fondationlouisdebroglie.org/AFLB-491/1036Durt.pdf), and **"Testing de Broglie's Double Solution in the Mesoscopic Regime,"** [*Found. Phys.* (2022)](https://link.springer.com/article/10.1007/s10701-022-00626-1) — on whether the split has testable consequences, which is the question that decides whether it's physics or bookkeeping.
- **Harris et al. (2013)** — the emergent-statistics half.

If you read two: **Colin–Durt–Willox** for whether the ontic/epistemic split is coherent, and **Dagan–Bush** for the HQA implementation of it.

## Darrow 2024

> does HQFT reproduce the Schrödinger equation in any limit?

Good question to press on — I told you last turn it doesn't, and that's right for HQFT as published but wrong as of 2024. There's a precise result I should have flagged.

### Dagan & Bush's HQFT (2020): no, and they say so

From their own conclusions:

> "Particular attention will be given to coupling this dynamical picture to the Ensemble Interpretation of quantum mechanics, and **assessing whether the statistical behavior of an ensemble of particles evolving according to the dynamics defined herein will be consistent with the standard formulation of quantum mechanics.**"

That's stated as future work. What they actually *derive* is narrower: the vibrating particle destabilizes into translation with $\bar p = \hbar k$ — the de Broglie relation emerges rather than being assumed. Everything about quantum statistics is explicitly conjectural in that paper ("one can envisage how the former may give rise to statistical behavior consistent with...").

And Durey & Bush's "non-relativistic limit" is not a Schrödinger limit — it means $|v| \ll c$, yielding a reduced *trajectory* equation. Different thing.

### But the successor framework does — exactly, and recently

**Darrow & Bush** ([*Symmetry* **16**, 149 (2024)](https://www.mdpi.com/2073-8994/16/2/149), [arXiv:2408.06972](https://arxiv.org/abs/2408.06972)) reformulated the whole thing as a Lorentz-covariant Lagrangian family coupling a Klein–Gordon field to a relativistic point particle — local, variational, two-way coupled. Then:

**D. Darrow, "Convergence to Bohmian mechanics in a de Broglie-like pilot-wave system,"** [arXiv:2408.05396](https://arxiv.org/pdf/2408.05396) (2024), proves a theorem. Writing the field as $\phi = \psi e^{-i\omega_c t}$ (factoring out the Compton oscillation):

> **Claim.** … In the non-relativistic limit $u \ll c$, $|\dot\psi| \ll \omega_c|\psi|$, the quantities $\psi$ and $q_p$ converge to $\psi^B$ and $q_p^B$ [the free Bohmian wavefunction and trajectory], respectively.

And he emphasizes: *"This result holds in a highly general setting of the pilot-wave framework, with arbitrary domain, initial conditions, and potential."*

So yes — a local, classical, relativistic field theory in ordinary 3-space converges **exactly** to single-particle Bohmian mechanics, with $\psi$ obeying the Schrödinger equation, in the non-relativistic limit.

### And the decomposition is precisely your unscrambling

This is the part you'll want. Darrow's own description of the limit:

> "the system converges exactly to single-particle Bohmian mechanics, where **the guiding wave (less a particle-centred wavepacket) plays the role of the Schrödinger wavefunction.**"

The field splits into exactly the two pieces you proposed:

| piece | role |
|---|---|
| a **particle-centred wavepacket** at the Compton scale, carried along with the particle, non-dispersing | de Broglie's **$u$** — ontological |
| the **remainder**, which converges to $\psi^B$ obeying Schrödinger | de Broglie's **$\psi$** — statistical/guiding |

And the mechanism that makes the limit work is exactly the decoupling you'd want: in this limit *"the particle continues to carry a Compton-scale wavepacket… but this wavepacket no longer radiates energy into the rest of the field."* The ontic wave stops feeding back into the statistical one, and the statistical one then obeys Schrödinger cleanly.

So the "omelette" does get unscrambled, and the theorem tells you the exact condition under which the two components separate.

### Three caveats that matter

**1. It's a *different coupling* from the HQFT/double-solution one.** Within the Darrow–Bush family, one coupling choice reproduces de Broglie's double-solution behaviours (the HQFT branch), and a *different* one gives the Bohmian limit. For the HQFT coupling itself, Darrow only conjectures: *"We conjecture that $\psi$ still converges to the Schrödinger wavefunction with such a coupling, and that the particle oscillates rapidly about a mean Bohmian trajectory."* So: **proven for the Bohmian coupling, conjectured for the HQFT coupling.**

**2. Single particle only.** The theorem says "single-particle Bohmian mechanics." No multiparticle sector, no entanglement, no configuration space.

**3. It recovers Bohmian mechanics, which for one particle is empirically identical to QM** — so this is a genuine recovery of single-particle quantum dynamics, not a weaker analogy. But it inherits Bohm's scope, not more.

### Where this leaves the program — and the through-line of this whole conversation

The ceiling is now located precisely, which is the valuable thing:

- **One particle, 3-space, local classical field theory:** fully recovers Schrödinger + Bohmian trajectories. Exactly. No configuration space needed.
- **Two entangled particles:** blocked, and not by a technical gap. A local classical field theory in 3-space caps at $S = 2$; nature gives $2\sqrt2$.

That's the sharpest available test of the thesis you opened with — that the exponential configuration space might be dispensable scaffolding. A serious research program has now shown it *is* dispensable for one particle, in full rigour, and the wall it hits at two particles is the Bell wall. Which locates the necessity of configuration space exactly where the earlier analysis said it lives: not in nonlinearity, not in single-particle dynamics, but in the tensor product.

**References**
- **Dagan & Bush**, [*C. R. Mécanique* **348**, 555 (2020)](https://thales.mit.edu/bush/wp-content/uploads/2021/04/Dagan2020.pdf) — HQFT; no Schrödinger derivation, question left open.
- **Darrow & Bush**, [*Symmetry* **16**, 149 (2024)](https://www.mdpi.com/2073-8994/16/2/149) — the Lorentz-covariant Lagrangian family.
- **Darrow**, [arXiv:2408.05396](https://arxiv.org/abs/2408.05396) (2024) — **the convergence theorem. This is your answer.**
- Prior related: **Sutherland** showed a certain coupling could be gauged away to match Bohmian mechanics; Darrow notes his mechanism is different and not gauge-removable.

## Separate collapse of A and B?

> going back to the Bell measurement ordering issue, the way I understand your answer is that it is assumed that the entire joint $A \otimes B$ quantum state does NOT collapse during the measurement -- only the A and B parts separately collapse, and that is why the order doesn't matter. Is this correct?

No — that's the one reading to discard. The joint state collapses **completely**, including Bob's half. Here it is numerically:

### Alice's measurement alone destroys the entanglement of the whole state

| | state vector | $\vert\det C\vert$ |
|---|---|---|
| initial $\vert\Phi^+\rangle$ | $(0.707107,\ 0,\ 0,\ 0.707107)$ | **0.5** — entangled |
| after **Alice only** ($\sigma_z$, got $+$) | $(1,\ 0,\ 0,\ 0)$ | **0.0** — product |

Alice measures her qubit, and the *joint* state goes from maximally entangled to $|00\rangle$. Bob's qubit goes from maximally mixed ($\rho_B = I/2$) to the definite pure state $|0\rangle$. That's about as global and as nonlocal as collapse gets. Nothing is confined to Alice's factor.

### The actual reason order doesn't matter

It's commutativity of the projectors, full stop:

$$(I\otimes P^B_b)(P^A_a\otimes I) \;=\; (P^A_a\otimes I)(I\otimes P^B_b) \;=\; P^A_a \otimes P^B_b$$

Both operators act on the **entire** joint space. Both collapse the **entire** joint state. They just happen to commute, so composing them in either order lands in the same place.

### And your original intuition was right about one thing — the intermediate states genuinely differ

| | state |
|---|---|
| after **Alice only** | $(1,\ 0,\ 0,\ 0)$ |
| after **Bob only** | $(0.853553,\ 0.353553,\ 0.353553,\ 0.146447)$ |
| **identical?** | **No** |

| | state |
|---|---|
| **Alice then Bob** | $(0.92388,\ 0.382683,\ 0,\ 0)$ |
| **Bob then Alice** | $(0.92388,\ 0.382683,\ 0,\ 0)$ |
| **identical?** | **Yes** |

Both orders end at $|0\rangle_A \otimes |{+}_{45°}\rangle_B$, but they pass through *different* intermediate states. So there **is** an order-dependent fact — it lives entirely in the state history, and nothing observable depends on it.

### The three levels, which is what your question was really probing

| quantity | changes when Alice measures? |
|---|---|
| global state vector | **yes** — entangled → product, instantly, across the whole system |
| Bob's **conditional** $\rho_B$ (given Alice's outcome) | **yes** — $I/2 \to \vert0\rangle\langle0\vert$ |
| Bob's **unconditional** $\rho_B$ (averaged over Alice's outcomes) | **no** — stays exactly $I/2$ |
| joint statistics | **order-independent** |

Verified: conditional $\rho_B = \begin{psmallmatrix}1&0\\0&0\end{psmallmatrix}$, unconditional $\rho_B = \begin{psmallmatrix}0.5&0\\0&0.5\end{psmallmatrix}$.

**That third row is the one doing the work.** What's confined to Alice's side isn't the collapse — it's the *accessibility* of the collapse. Bob's description changed the instant Alice measured, but to notice it he'd need Alice's outcome, and that bit travels at $\le c$. Without it, he averages over her outcomes and sees $I/2$ forever.

So the correct summary is the inverse of your reading:

> The collapse is global and nonlocal. What is local is only what anyone can *detect* of it.

### Which brings back the Aharonov–Albert point with more force

Since the intermediate states differ between the two orders, and since for spacelike separation there is no frame-invariant fact about which order occurred, there are **two genuinely different state histories**, each correct in its own frame, converging on the same endpoint and the same statistics.

That's the real residue of your original question, and it's exactly why Aharonov and Albert concluded that the covariance of relativistic quantum theory resides in the experimental probabilities and not in the underlying states. Your instinct that "there should be cases where the order matters" was correct — it matters for the narrative, and the narrative turns out not to be frame-invariant, which is a decent argument that it isn't physical.

## Projector ordering is due to tensor product separation of A and B

> Ok, so is the underlying mathematical reason that the projector operators commute is that the A and B elements live in different cells in the outer product tensor space? so, mathematically, when A rotates the entire space, the rotation on B is somehow "information preserving", whereas A has itself rotated into the eigenvector of the outcome? Can you show explicitly what happens to the full tensor product space after each order of measurement (A then B, B then A). Also, Aharonov & Albert don't specifically talk about the Bell test situation -- what paper(s) have specifically applied this order-dependent logic to the Bell test situation?

### 1. Yes — and the precise statement is one line

The tensor product multiplies slot-wise: $(M\otimes N)(M'\otimes N') = (MM')\otimes(NN')$. So

$$(P^A_a\otimes I)(I\otimes P^B_b) = (P^A_a\cdot I)\otimes(I\cdot P^B_b) = P^A_a\otimes P^B_b$$
$$(I\otimes P^B_b)(P^A_a\otimes I) = (I\cdot P^A_a)\otimes(P^B_b\cdot I) = P^A_a\otimes P^B_b$$

In indices: $(P^A\otimes I)_{(ij),(kl)} = P^A_{ik}\,\delta_{jl}$ and $(I\otimes P^B)_{(ij),(kl)} = \delta_{ik}\,P^B_{jl}$. Composing sums over the middle index pair and gives $P^A_{ik}P^B_{jl}$ either way. Each operator writes only on *its own* index slot and leaves a Kronecker delta on the other, so they can't collide.

Note what this does **not** depend on: the state, the entanglement, or the angles. It's purely kinematic — a property of the tensor product, true for every state and every pair of settings.

### 2. Two corrections to the "rotation" picture

**(a) A projector is not a rotation.** $P^A_+\otimes I$ is not unitary and not invertible — it's a rank-2 projector on $\mathbb{C}^4$. It *destroys* information; that's what collapse is.

**(b) Nothing at all happens to Bob's slot.** It's not an "information-preserving rotation" — the operator's B-factor is literally $I$. No operation.

And yet Bob's state changes. That's worth stating flatly, because it's the whole content of entanglement:

> The operator acts as the **identity** on Bob's factor. Bob's state changes **anyway** — because the projection on Alice's side selects which term of the entangled superposition survives, and that surviving term carries a definite B-part with it.

### 3. The explicit walkthrough

Basis order $\{|00\rangle, |01\rangle, |10\rangle, |11\rangle\}$. Alice measures $\sigma_z$ (0°), Bob at 45°, both get $+$.

```
P^A_+ ⊗ I                          I ⊗ P^B_+
[1.000  0      0      0    ]       [0.853553  0.353553  0         0       ]
[0      1.000  0      0    ]       [0.353553  0.146447  0         0       ]
[0      0      0      0    ]       [0         0         0.853553  0.353553]
[0      0      0      0    ]       [0         0         0.353553  0.146447]

keeps the |0⟩_A BLOCK entirely,     BLOCK-DIAGONAL: the identical 2×2
identity within each block           operation inside every block
```

**This is the geometric reason they commute.** One operator chooses *which block survives* and is blind to what's inside a block. The other acts *identically inside every block* and is blind to which block it's in. Block-selection and within-block-action can't interfere.

Commutator: $\max|\,[\cdot,\cdot]\,| = 0.0$ exactly.

#### Order 1: Alice then Bob

| step | $\vert00\rangle$ | $\vert01\rangle$ | $\vert10\rangle$ | $\vert11\rangle$ | prob |
|---|---|---|---|---|---|
| start $\vert\Phi^+\rangle$ | 0.707107 | 0 | 0 | 0.707107 | |
| $P^A_+\otimes I$ applied | 0.707107 | 0 | 0 | **0** | 0.500000 |
| *renormalized* → $\vert00\rangle$ | 1.000000 | 0 | 0 | 0 | |
| $I\otimes P^B_+$ applied | 0.853553 | 0.353553 | 0 | 0 | 0.853553 |
| *renormalized* | **0.923880** | **0.382683** | **0** | **0** | |

#### Order 2: Bob then Alice

| step | $\vert00\rangle$ | $\vert01\rangle$ | $\vert10\rangle$ | $\vert11\rangle$ | prob |
|---|---|---|---|---|---|
| start $\vert\Phi^+\rangle$ | 0.707107 | 0 | 0 | 0.707107 | |
| $I\otimes P^B_+$ applied | 0.603553 | 0.250000 | 0.250000 | 0.103553 | 0.500000 |
| *renormalized* | 0.853553 | 0.353553 | 0.353553 | 0.146447 | |
| $P^A_+\otimes I$ applied | 0.853553 | 0.353553 | **0** | **0** | 0.853553 |
| *renormalized* | **0.923880** | **0.382683** | **0** | **0** | |

**The detail worth noticing:** the *unnormalized* vector after both steps is byte-identical in the two orders — $(0.853553,\,0.353553,\,0,\,0)$. That's just $(P^A_+\otimes P^B_+)|\Phi^+\rangle$, the single operator both paths are secretly computing. The intermediate states differ ($|00\rangle$ versus $(0.853553, 0.353553, 0.353553, 0.146447)$); the composite doesn't.

Joint probability $0.5 \times 0.853553 = 0.426777$ both ways, matching the one-shot $\langle\Phi^+|P^A_+\otimes P^B_+|\Phi^+\rangle = 0.426777$.

### 4. Papers applying this specifically to Bell tests

You're right that Aharonov–Albert is general relativistic measurement theory. Here's the Bell-specific literature.

**Theory**

- **Tim Maudlin, *Quantum Non-Locality and Relativity: Metaphysical Intimations of Modern Physics*** (Blackwell 1994; 3rd ed. Wiley-Blackwell 2011). **The book-length treatment of exactly your question.** Develops *hyperplane dependence* — the result that with three mutually spacelike measurements on an entangled state, the quantum predictions depend on the hyperplane, so you cannot consistently conjoin predictions from distinct measurements. This is your ordering problem taken to its sharpest form.
- **Asher Peres, "Classical interventions in quantum systems. II. Relativistic invariance,"** [*Phys. Rev. A* **61**, 022117 (2000)](https://arxiv.org/abs/quant-ph/9906034). The technical version: spacelike interventions appear as *different sequences of quantum jumps* in different Lorentz frames, and he derives the conditions on the operators that force both sequences to yield identical observables. This is precisely the calculation above, generalized.
- **Wayne Myrvold, "On peaceful coexistence: is the collapse postulate incompatible with relativity?"** *Stud. Hist. Phil. Mod. Phys.* **33**, 435 (2002) — [free PDF](https://publish.uwo.ca/~wmyrvold/PeacefulCoexistence.pdf). Argues the conflict is only apparent, via the transformation properties of states.

**Experiment — and this is the part you'll enjoy**

Somebody built a theory in which the order *does* matter, and then tested it.

- **Suarez & Scarani (1997)** proposed **multisimultaneity**: a causal model assigning a real time-ordering to every pair of events. Its distinctive prediction is that the correlations should **disappear** in the "**before–before**" configuration — where both analyzers are in relative motion such that *each one, in its own rest frame, selects its output first*.
- **Zbinden, Brendel, Gisin & Tittel**, "Experimental test of nonlocal quantum correlation in relativistic configurations," *Phys. Rev. A* **63**, 022111 (2001).
- **Stefanov, Zbinden, Gisin & Suarez**, "Quantum correlations with spacelike separated beam splitters in motion: experimental test of multisimultaneity," *Phys. Rev. Lett.* **88**, 120404 (2002). Acousto-optic modulators as *moving* beam-splitters — the Doppler shift of the reflected photons defines an unambiguous rest frame in which "before or after" is well posed — with interferometers separated by 55 m.

**Result: the correlations persist, exactly as standard QM predicts. Multisimultaneity is falsified.** The before-before experiments rule out nonlocal models with a deterministic, time-ordered nonlocal part.

So the empirical answer to "does the order matter in a Bell test" is not merely "the formalism says no" — it's "a rival model said yes, it was built and run, and the answer came back no."

**Sources:** Maudlin, Myrvold, Peres, and Stefanov et al. (PRL **88**, 120404) verified against listings just now. The Suarez–Scarani (*Phys. Lett. A* **232**, 9) and Zbinden et al. (*PRA* **63**, 022111) volume/page numbers are from memory — the papers, dates and content are confirmed.



