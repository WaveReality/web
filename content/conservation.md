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

Mathematically, even though everything ultimately obeys local conservation laws, the specification of these laws can be divided into **global gauge invariant** conditions and local ones. Global invariants amount to transforming (translating or multiplying) the overall state of the system by _constant_ factors that apply the same way everywhere, while local ones emerge from [[gauge theory]], which involves introducing new _local_ transformation parameters at each point in spacetime. These local parameters represent a new _field_ of values, with the result specifying how two fields interact in a way that is locally conservative.

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

The final 3 entries in the table define the charge-like values conserved (via the local gauge invariance mechanism in [[gauge theory]]) by the 3 fundamental forces in the Standard Model. These are defined by a _unitary group_ (U) or _special unitary group_ (SU) of the given dimensionality. A unitary group just means that it is conservative, like multiplying by the unit 1. The U(1) group is basically Maxwell's equations interacting with a complex-valued matter / charge wave. The SU(2) group is defined by a 2x2 unitary matrix, which is the rotation group that is defined by a single parameter (the rotation angle), consistent with the idea that this group defines the property of [[spin]]. Thus, the [[weak]] force is fundamentally about spin, and the conserved charge-like value is weak isospin, although this conservation status is complicated by the [[Higgs]] field.

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

This minimal set of conservation laws is perhaps less "principled" than it might otherwise seem, which is consistent with the idea that there might be more "pragmatic" (mechanistic) factors at work, versus purely abstract principles such as symmetry.

For example, the U(1) interaction generated by the local gauge transformation is strongly motivated by the nature of electromagnetic potentials, which of course were well known in advance of the development of gauge theory. Furthermore, there remain some interesting mismatches between the gauge theory formalism and the actual nature of EM waves, which have 2 fewer degrees of freedom relative to the gauge theory version.

The SU(2) x U(1) interaction by itself could be a lot simpler, but instead it is greatly complicated by the [[Higgs]] mechanism to create the [[weak]] interaction, in terms of an integrated _electroweak_ framework that is then subject to spontaneous symmetry breaking. This whole apparatus is necessary because the only kind of field that can interact in a conservative (i.e., gauge invariant) way with a charge source field is one with _massless_ wave equations (i.e., Goldstone bosons; [[@GoldstoneSalamWeinberg62]]; [[@Goldstone61]]). If you have wave fields with a mass term (i.e., the [[Klein-Gordon]] equations), then it is no longer gauge invariant, so the whole enterprise breaks down.

The fact that the electroweak mechanism with spontaneous symmetry breaking actually works is kind of amazing, but it is far from a "clean" principled application of symmetry principles. There are various empirically-defined coupling factors and mixing angles that are needed to make it work, which all enter as external parameters of the theory. Furthermore, there are huge discrepancies between the predictions of the Higgs mechanism for the energy of the [[vacuum]], and associated with the mass of the Higgs boson, that continue to generate consternation about the overall framework.


