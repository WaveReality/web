+++
Categories = ["Interpretations"]
bibfile = "mechphys.json"
+++

{id="figure_ca2d" style="height:20em"}
![Illustration of a simple 2-dimensional cellular automaton: space is divided into regular square cells (a uniform, regular tiling of space), and neighboring states interact by influencing the state update. Time updates synchronously, setting the fastest rate of propagation as cell width / time update.](media/fig_ca_2d.png)

{id="figure_cubes" style="height:40em"}
![Neighborhood interactions in regular cubic tiling of space in three-dimensions --- these interactions are used to compute the wave equation locally.](media/fig_space_cubes_fec_lapl.png)

The overall approach taken here is to determine if a specific framework for _implementing physics_ is able to account for all of the known phenomena described by the [[Standard Model]]. At the broadest level, the motivation is that Nature is some kind of fundamental physical _process_ that is _happening everywhere, autonomously_ producing everything in the universe. We seek a description of this process and the "state" that it operates on.

There is of course no guarantee that we can succeed in this mission, but neither is it clear that such an approach is doomed to fail (despite several issues that would appear to contradict this assertion). The goal is simply to see how far we can get, and what issues we encounter in the process.

Why doesn't the Standard Model itself provide the desired level of description? Because it requires extensive expert knowledge to perform detailed _analytical_ computations, which themselves are ultimately approximate due to the infinite sums resulting from the [[renormalization]] procedure.

By contrast, [[Maxwell]]'s equations for electromagnetic (EM) radiation in the Lorenz gauge provide an entirely local, autonomous, _mechanistic_, model of wave propagation that doesn't require any human intervention or expertise to produce the resulting physics. It can just _happen_ like that autonomously, everywhere in space. See [[tools vs models]] for more on this distinction.

This satisfying level of understanding is what led the physicists in the late 1800's to believe that physics was nearly solved. Critically, they hypothesized the presence of the [[aether]] as a kind of physical substrate for these EM waves, which was then invalidated by the Michaelson-Morley experiment. This level of thinking was taking things one step too far, however. We don't need to impose any kind of macroscopic, intuitive mechanism underlying the basic physical mechanisms. 

The whole point is that, at the most fundamental level, there are just fundamental mechanisms that we can describe, but, because they are fundamental, it is pointless to try to then impose some further "steampunk" kind of gears and fluids underlying these fundamental mechanisms. You have to stop _somewhere_. And that description must be compatible with all known physical phenomena per the Standard Model.

## The computabilty constraint

So what kind of principled constraints can we impose on the kind of description we seek? One fundamental constraint is that it should be **computable**. That is, the resulting theory should be able to be implemented on a universal computational device (i.e., a Turing machine), and it should run according to a **fixed program** over some kind of well-defined **state** variables. The fixed program may involve fundamentally stochastic processes, but it should not have any internal loops with a possibility of non-deterministic stopping behavior. There are obvious implications of this for [[quantum computer]]s.

These kinds of considerations have led to the computational framework known as a **cellular automaton (CA)**, which has been investigated as a basis for fundamental physics modeling since the 1950s. A CA consists of a regular, uniform division of space into **discrete cells**, each of which has one or more _state_ values, and each cell interacts only with its nearest neighbors (i.e., locally) to update its state value over time ([[#figure_ca2d]]; [[#figure_cubes]]).

Such a system was first described by Stanislaw Ulam in 1950, and it provides the simplest kinds of answers to fundamental questions about space, time, and the basic nature of physical laws ([[@VonNeumannBurks66]]; [[@Zuse70]]; [[@FredkinToffoli82]]; [[@Feynman82]]; [[@Fredkin90]]; [[@Hooft15]]). Space is _real_ and fundamental in the form of the underlying cells --- it isn't just an empty [[vacuum]] or a mathematical continuum.

The discretization of space, as contrasted with a true continuum, is essential for computability. It is impossible to perform finite computation on fully continuous space / time, because that would require an infinite amount of computation per time step. This is effectively Zeno's paradox. The only way that the continuum limit can be used is with analytical mathematics. Any computational implementation inevitably requires a discretization. In terms of the levels of infinities associated with the Cantor sets, a discrete space corresponds to the lowest level of infinity associated with the integer number line.

One still has an infinity to deal with, and this is plenty mind-blowing all by itself: space and time continuing infinitely in all directions, forever. But at least the further difficulty of an infinity of space or time _within_ any given segment, which is required for a truly continuous dimension, can be avoided. One could reasonably argue that the infinity of space and time is more plausible than the notion of an edge, as in the old flat Earth models and the end of the world.

The physical phenomenon of discrete point-like elementary particles also suggests the need for an **ultraviolet cutoff** at some distance scale: infinitely small point-like particles predict infinite field amplitudes in their immediate neighborhood. Instead, a discrete lattice with a given grid spacing distance provides a tractable, non-divergent mechanism for such discrete particles. This potentially resolves deep problems with the Standard Model surrounding [[renormalization]] and the nature of the [[back reaction]]. See also [[configuration space]] for computability and other arguments against the use of exponentially-large state spaces.

Time emerges naturally in its unique unidirectionality within the CA framework, simply as a discrete rate of change in the state values. Furthermore, the ratio of discrete spatial cell width to discrete rate of state update provides a natural upper limit to the rate at which anything can propagate within this system: i.e., the **speed of light** in a vacuum. Thus, this principal postulate of special relativity that light has a fixed upper speed limit emerges as a necessary consequence of more fundamental assumptions about the nature of space and time in the CA framework. 

Furthermore, the basic [[wave]] equation can be computed using a simple local neighborhood interaction among cells in a CA-like system, and [[Maxwell]]'s equations for the electromagnetic field and [[Dirac]]'s equation for the quantum wave function of an electron can be computed using primarily this basic wave equation. This framework predicts all of the strange properties of [[special relativity]] in a principled manner based on the discretization of space and time, together with wave dynamics. 

Note that these wave-based equations do require real-valued state variables, which is a departure from the simplest form of CA that only employs simple discrete state values. The continuously varying nature of electromagnetic radiation (EM) (i.e., light) provides strong evidence that the states of nature do have this continuous-valued nature. This use of real-valued state makes our CA framework essentially identical to a **finite element model**. Related techniques are used in the **lattice quantum chromodynamics** approach to numerically simulating the strong force between [[quarks]].

In addition to real-valued state, we require a stochastic process to randomly choose the next location for a discrete particle to move (see [[stochastic motion]]). 

Ultimately, the model needs to _work_ to explain the available data, and our intuitions about "mechanistic" level plausibility are secondary concerns: if these intuitions align with reality in a way that makes everything work, it is obviously great, but you cannot let them stand in the way of making progress.

To summarize, here are the critical constraints imposed by the CA framework that we adopt:

* Space and time are discrete.

* State can be continuous-valued.

* State is updated discretely at each time step, in _parallel_ across all cells, using only _locally available_ information within each cell and from its 26 neighbors in 3D space.

This parallel and local nature of the state update imposes strong constraints on the nature of the available mechanisms, in ways that end up aligning with known physical properties.

## The Game of Life

The CA framework was popularized in its two-dimensional form in _the game of Life_ by John Conway (described by [[@Gardner70]]). As anyone who has seen this system in operation knows, it is capable of producing remarkable complexity from such simple, local, deterministic rules.

In this CA (widely available as a screensaver), there is a two-dimensional grid of square cells, with each cell having a single binary state value (0 = "dead" and 1 = "alive"). This state value updates in discrete, simultaneous steps as a function of the state values in the 8 neighbors of each cell:

* If the sum of the neighbors' states is > 3 or < 2, then the cell is dead (0) on the next time step (from "overcrowding" or "loneliness", respectively).

* If it has exactly 3 live neighbors and is currently dead, then it is "born" and goes to 1.

* If it was already "alive" then it remains so, only if it has 2-3 living neighbors.

## Physical implications

One specific implication of the CA framework is that it unambiguously establishes the position basis as primary for representing discrete massive particles, consistent with the [[pilot-wave]] framework. In addition, the [[Pauli exclusion principle]] is strongly suggestive of a discrete CA-like state. This principle posits that only one _fermion_ (electron, quark, etc, with a quantum spin of 1/2) can occupy the same quantum state, including position, at a time.

Thus, the underlying CA state representation only needs to be able to hold one of each particle type, which eliminates the difficult problem of having to represent a variable number of such particles at each location. In other words, the "memory allocation" for each cell is constant, regardless of what kind of matter or energy might be present. This is not the case for the [[configuration space]] used in standard quantum frameworks, which has the hallmark of a calculational tool in that it is always constructed for each specific problem being solved.

This exclusion principle does not apply to _boson_ particles, and would thus require an indefinite number of memory slots to represent (along with all the other difficulties involved in the particle picture for force fields). Thus, the CA framework, along with a number of other considerations, strongly supports the [[semiclassical]] picture, where quantum particles interact with a classical purely wave-based EM state propagating according to [[Maxwell]]'s equations.

The notion of _autonomy_ in a CA is also particularly important as a physical model: the CA is entirely self-contained and can just plug away forever, running the same exact local laws every time step. By contrast, most calculational tools used in physics require a specific setup and different computational steps depending on exactly what situation is being modeled: they are far from "autonomous" in the sense of a CA.

When you look at the examples of plausible physical models ([[tools vs models]]), they all have this same autonomous character: e.g., general relativity and Maxwell's equations in the Lorenz gauge can just be configured with a starting state and then everything can evolve autonomously from there.

In summary, the CA framework is simple, elegant, and consistent with the most basic facts of physics. If one could develop a viable physical theory within the general confines of this framework, it would provide a uniquely simple and satisfying model of how nature works.

However, two important objections are typically raised about such a framework: Isn't it just like the [[aether]] that was so famously rejected by the Michelson-Morley experiment? If it has purely local interactions, how could it possibly account for the apparent [[non-locality]] of QM? See those pages for further discussion.

