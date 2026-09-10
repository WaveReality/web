+++
Categories = ["Interpretations"]
bibfile = "mechphys.json"
+++

**Configuration space** is a critical element of standard quantum mechanics, in most of its various formulations, including the classical [[Hilbert space]] formulation evolving according to the [[Schrodinger]] wave equation, and _also_ the [[pilot-wave]] framework. It is the space defined by the **multiparticle** configuration of all the elements of relevance to a given experimental setup being analyzed. Critically, the size of this space increases **exponentially** as a function of the number of such elements, because it is effectively a **tensor product** space, i.e., the **outer product** across all elements.

For example, the basic [[wave]] equation for a single "particle" (e.g., using the [[Klein-Gordon]] or Schrödinger equation) can be represented using the [[cellular automaton]] (CA) framework with two real-valued numbers per each state (e.g., for position and velocity, or the two complex numbers in the Schrödinger version), using a 3D state space discretized at a specific resolution. If we use _R_ discrete cells per spatial dimension, the total size of this state space is:

{id="eq_n1" title="size of 1-particle state space"}
$$
N = 2R^{3}
$$

The exponential character of this space becomes apparent when we see that each particle or element requires its own 3 additional dimensions in this space, where _M_ represents the total number of such elements:

{id="eq_nm" title="size of M-dimensional state space"}
$$
N = 2R^{3M}
$$

What this means is that you effectively have different discrete CA cells for each possible combination of individual dimensions for each element. For example, there is a unique cell representing the wave function for the specific combination of position 23 on the _X_ axis of particle _A_ _and_ position 12 of the _Z_ axis of particle _B_. This is what an outer product means: all possible combinations are represented with their own distinct state value.

It should be clear that this space is entirely [[non-local]] across particles. Each possible spatial location in one particle is "crossed" with each possible spatial location in the other particle. This is why the standard formalism has no problem dealing with non-local entanglement across spatially-separated particles.

From a purely computational point of view, the outer-product representation is maximally expensive, and maximally expressive. It is essentially computationally universal, allowing the system to represent any and all interactions in any way that might be necessary. Thus, it is not unreasonable to suggest that this representation is a [[tools vs models|calculational tool]], consistent with various arguments in the literature ([[@Wallace21]]; [[@Myrvold15]]), but not without some debate ([[@Carroll21]]; [[@North12]]; [[@NeyAlbert13]]).

## Necessarily exponential?

The exponential size of configuration space represents a _fatal_ barrier to any kind of plausible CA model of physics, in the same way that non-local computation does. Therefore, we are strongly motivated to adopt virtually _any_ other kind of physical mechanism that might avoid the need for such a thing.

Interestingly, there are various lines of evidence that the exponential size is indeed excessive, exponentially so. First, [[@^PoulinQarrySommaEtAl11]] concluded that the exponential size of configuration space was unnecessary, because the actual physical time-dependent [[Hamiltonian]]s don't end up using the vast majority of the space. This appears to contradict a proof from [[@^Montina08]] showing that the exponential space is necessary, but they made much stronger assumptions in this proof, including the ability to transition from any state to any other via a finite number of quantum gates. One important difference is that the first paper used a weaker constraint (polynomial vs. finite). Furthermore, the second paper requires a Markovian state evolution mechanism, which they argue is the first assumption to question in challenging the proof. Thus, the full resolution at this abstract mathematical level appears to be inconclusive.

[[@^NorsenMarianOriols15]] analyzed the contributions of configuration space to a pilot-wave based framework. They concluded that indeed the configuration space contains a large amount of "redundant" information, and that even the simplest approximation for the inter-particle interaction terms does a reasonable (yet imperfect) job of capturing the behavior of the full configuration-space model. Exploration of higher-order terms in this approximation are ongoing ([[@Norsen22]]).

Another major constraint on both non-locality and configuration space comes from considerations of integrating [[gravity]] with quantum mechanics, which is famously difficult, and seemingly impossibly contradictory in many respects. At a very basic level, general relativity, like special relativity, is fundamentally local in spacetime, and it is unclear how the apparent non-locality of the quantum world could be reconciled. The ultraviolet divergence of the standard model, reflected in the [[renormalization]] procedures, also makes it unclear how much energy might be in a given localized region of space.

A further contradiction is that according to general relativity, the entropy inside a given region of spacetime is strictly limited, as a function of the surface area of the region, known as the _Bekenstein-Hawking_ bound ([[@Bekenstein81]]). This implies a corresponding limit in the number of degrees of freedom in the corresponding quantum state, which would be well below that corresponding to any kind of exponential configuration space. The interplay between this constraint and considerations from quantum [[field theory]] remains an active topic of investigation ([[@BoussoChandrasekaranShahbazi-Moghaddam20]]; [[@Witten22]]; [[@Giddings15]]; [[@Witten18]]; [[@Casini08]]). 

Overall, it is difficult to draw strong conclusions from these existing analyses, but it does not seem certain that the exponential size of configuration space is an absolute irrevocable requirement to capture the phenomenology of quantum physics. Therefore, given how toxic such a thing is to the CA-based approach, we can adopt a pragmatic approach to see if indeed we can succeed at this enterprise, as documented in the [[Spinfield Model]].

Ultimately, the most relevant data will come from [[quantum computer]]s, which are pushing the envelope on this issue, as their computational power derives directly from the exponential size of configuration space. See that page for relevant current status.

