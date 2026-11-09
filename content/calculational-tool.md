+++
Categories = ["Interpretations"]
bibfile = "mechphys.json"
+++

The differences between the standard [[Copenhagen]] interpretation vs. the [[pilot-wave]] model exemplify the broader distinctions between **calculational tools** vs. **physical models**. A calculational tool is a mathematical framework that makes it easy to compute relevant results, but the underlying processes implied by the nature of the calculations have no direct mapping onto underlying "real" physical properties of the system.

Because there is so much longstanding confusion and controversy about the quantum realm, it is essential to anchor this argument in a very well-understood domain, that provides a clear illustration of the relevant differences. The contrast between the **Coulomb** and **Lorenz** gauge representations of [[Maxwell]]'s equations of electromagnetic (EM) phenomena provides a paradigmatic example of the differences between a calculational tool versus a physical model.

The Coulomb gauge provides a convenient way to calculate the electrostatic force between two charged particles in terms of their respective charges ($q_1$, $q_2$) and distance _r_ between the two particles:

{id="eq_coulomb" title="Coulomb electrostatic force law"}
$$
F = k_e \frac{q_1 q_2}{r^2}
$$

But this is not a plausible physical model because the math requires the use of the distances between relevant particles, and their respective aggregate charges. Not only is this a (highly) non-local computation ("action at a distance"), Nature would presumably have to run this computation for all other bodies within some relevant radius, which ends up being exponentially explosive.

By contrast, a plausible physical model should compute EM forces directly from the collective effects of each individual charged particle everywhere, and propagate these forces via local mechanisms at the speed of light. Indeed, that is precisely what the [[Maxwell#the Lorenz gauge]] version of Maxwell's equations provides, which propagates EM forces in an entirely local manner via standard [[wave]] equations.

The seemingly obvious conclusion is:

> If you're trying to understand how Nature actually operates (i.e., in terms of a physical model), you definitely want to use the Lorenz gauge version, and not the Coulomb one.

But critically:

> Both gauges describe _exactly_ the same physics, so you can't run an experiment to tell you which gauge Nature is actually using.

Some people will conclude from this last point that it therefore is _entirely irrelevant_ which gauge Nature is using, because all you care about is having _any_ complete description of the physical phenomenon, and everything else is just "aesthetics". Furthermore, you should definitely have as many tools in your toolkit to conveniently solve different types of problems with a minimum of effort. Indeed, according to this criterion, the Coulomb gauge is typically much more convenient that the Lorenz gauge (although you can get into trouble by using this gauge incorrectly; [[@BrillGoodman67]]; [[@Jackson02]]; [[@Onoochin01]]).

> Thus, the core remaining question is why the **form** of the mathematical equations still matters, even if the end results of the calculation are the same?

In short, the form of the equations matter if you think of them as a description of a process that is actually taking place physically -- i.e., that the equations actually map onto corresponding physical processes in Nature. You could choose to ignore this, and just say "math is just math -- all I care about are the results!". But this elides the central question:

> Do you think that Nature actually exists as a physical process, "generating" all of the phenomena we observe with our senses and measuring devices?

If you do believe in such a thing (you don't have to -- it is a philosophical choice that every individual has to make for themselves), then there is a "logical" subsequent question as to what that physical process actually is. What is the best description we can provide of this process? This is nominally the goal of the field of physics, to provide just such a description.

Obviously, our description should **account for all the data**. But what happens if you have two different theories that both account for the _same_ data, but they differ in the _way_ that they do so? This is precisely the situation with the Coulomb and Lorenz gauge formulations. The straightforward idea here is that:

> A physical model should provide a set of equations that describe the fine-grained process by which the relevant physics emerges, not just the final results (which can be provided by a calculational tool).

Thus, in the case of the Lorenz gauge, this fine-grained process is a wave equation that can be implemented with an entirely local computation in 3D coordinates, with each point in 3D space reacting locally to the state of its neighboring points, as described in [[wave]] and by the [[wave simulation]].

By contrast, the Coulomb equations describe the non-local transmission of force instantaneously across arbitrary spatial distances, characterized by the _r_ distance factor in [[#eq_coulomb]]. If we would take this literally as a description of the way that Nature actually operates, then clearly there would need to be some accounting of how this _r_ factor could be available at each particle's location to drive the force that then causes them to accelerate, etc. Clearly, this is problematic when you consider each particle in the universe having to somehow have access to every other particle's distance and charge.

Thus, according to this line of reasoning, as descriptions of the physical processes of Nature, it seems reasonable to conclude that the Lorenz gauge description is much less problematic than the Coulomb one. The relevant information is available locally, and propagates naturally at the speed of light. This process of the electrostatic force propagation is an essential property of EM, and additional steps are required to accommodate it in the context of the Coulomb gauge formulation (known as the _retarded potentials_, which take into account the time lags associated with speed-of-light transmission).

If the EM waves propagate through anything other than empty space, then the nature of this EM force propagation is affected, in potentially very complex ways. Thus, as you consider increasingly complex situations, the Coulomb formulation becomes increasingly unwieldy, while the Lorenz gauge continues to function in precisely the same way.

The consistent nature of the Lorenz gauge formulation makes it a [[universal basis]] representation for EM. Any and all EM phenomena can be captured using this one set of local wave equations operating over the EM four-vector state $A^\mu$. Thus, there is no evident barrier to thinking of these wave dynamics as a reasonable description of the actual process that Nature somehow implements.

We have absolutely no idea _how_ it actually works at a level below these wave equations, because we have no way of accessing or measuring anything at that level, so we just have to stop there and say, whatever Nature is doing, it is well-described by these wave equations operating on a four-vector state at each point in 3D space.

A skeptic may say, well, if you're going to have to stop at that level, why not just stop at a level above that, and be happy with the Coulomb gauge description? At that point, I would just have to say that you can choose to believe whatever you want, but given that we actually do have the Lorenz gauge formulation, and it provides this universal, local process model, that seems to provide a better description of the Natural process.

If we didn't actually have the Lorenz gauge formulation, the limitations of the Coulomb gauge might motivate you to find a better model. Indeed, historically, many people were disturbed by the Newtonian gravitational version of the [[#eq_coulomb]] inverse-square law, because of its implied action-at-a-distance. Einstein's development of general relativity provided a Lorenz-gauge-like solution in this domain. Of course, it also makes different testable predictions relative to Newtonian gravitation, so everyone can accept it as a better theory on the basis of explanatory scope.

## Application to quantum physics

The situation in quantum physics is like the EM gauge case, in that we already have calculational tools in the form of the [[Standard Model]] which provide highly accurate, extremely well-tested accounts of the phenomenology of quantum physics. Thus, it is highly unlikely that we will find a new theoretical account that makes all of these same high-accuracy predictions, _and_ makes clear novel distinctive predictions. 

Instead, we can most likely hope for something akin to the Lorenz gauge: a framework that makes more sense as a universal basis representation for quantum physics, that does not require hand-coded custom representations for each different type of problem that is being solved.

Meanwhile, it seems reasonable to conclude that something like the Copenhagen measurement process, with its entirely non-local wave-function collapse postulate at an otherwise ill-defined moment of measurement, is an example of a calculational tool like the Coulomb gauge, and it should not be interpreted as a physical model for the underlying processes of Nature.  

Furthermore, the popularity of the otherwise fantastical _many worlds_ interpretation (MWI, [[@Everett57]]) is largely due to the fact that it avoids the measurement problem entirely. But at what cost? By forking off _an entire new Universe_ at each possible point of measurement!? This is supposed to be a more plausible physical model?

Therefore, to reiterate:

> People should not waste their time worrying about the physical plausibility of a calculational tool.

Specifically, with the availability of the Bohmian version of pilot-wave model, which _eliminates the inexplicable nature of the measurement process_ ([[@NorsenMarianOriols15]]), it does not make sense to continue to try to make sense of the Copenhagen framework as a physical model. For example, the pilot-wave framework makes intuitive sense of the otherwise-mysterious [[double-slit]] experiment results: the particle goes through one slit, while the wave goes through both and gives rise to the interference effects. Done. See [[@^BushFrumkinSaenz24]] for discussion of the physical, macroscopic realization of these and many other quantum phenomena, using vibrating oil droplets.

The Bohmian pilot-wave model is directly analogous to the Lorenz gauge framework for EM, while the Copenhagen interpretation is likewise analogous to the Coulomb gauge. Both produce entirely the same mathematical results (so you cannot run an experiment to determine which one Nature is using), but the pilot-wave model provides a more plausible physical model, while the Copenhagen framework is typically more convenient for calculating results. 

It is also important to emphasize that, unlike the Lorenz gauge of EM (or general relativity), the Bohmian pilot-wave model has a number of significant limitations as a plausible physical model. However, _all_ of the issues with the pilot-wave model are actually present in the standard [[Hilbert space]] formulation as well: it is just that the pilot-wave model makes these issues more obvious, by making stronger ontological commitments. For example, the incompatibility of the pilot-wave model with [[special relativity]] is also entirely present in the [[configuration space]] framework used in the standard Hilbert space formulation ([[@Bell75]]; [[@AharonovAlbert81]]; [[@Peres00]]; [[@PeresTerno04]]; [[@Norsen11]]).

Do we really want to reward the foundational evasiveness of the Copenhagen or MWI frameworks, instead of recognizing and confronting the fundamental issues in the standard framework head on? According to surveys, most working physicists endorse either the Copenhagen or MWI frameworks, while none endorsed the pilot-wave framework ([[@Tegmark10]]). Thus, it does appear that evasiveness is winning at this point.

See [[universal basis]] for a way of formalizing what makes a physical model more intuitively appealing as such, without appealing to macroscopic physical intuition or other such vague constraints. That page also has further discussion about the limitations of the standard frameworks for providing such a universal basis representation, and highlights the ways in which the existing [[Standard Model]] is incomplete, despite its ability to generate spectacularly precise predictions of various experimental results.

## The seduction of elegant models

An important corollary of the "don't confuse the math for the physics" mantra is: **don't be seduced by elegant mathematical models**. For example, there is an overarching framework based on _symmetry groups_ for organizing the elements of the [[Standard Model]]. Many people have tried to push this framework in various ways (e.g., "supersymmetry") to make novel predictions, all of which have failed to date. This perhaps suggests that this elegant framework of symmetry is not really what is driving the structure of fundamental physics: it works as far as it goes, but it does not appear to be a truly _generative_ principle that allows one to get deeper insight than the hard-won, empirically-based standard model. See [[conservation]] for more discussion.

Another potential example is the idea of using _only_ interacting wave equations and somehow avoiding the need for any hard little [[particle]]s entirely ([[@Hobson13]]; [[@Sebens22]]). Wave equations are unambiguously at the foundation of the [[Standard Model]], and even such an exquisitely simple equation like [[Klein-Gordon]] provides such incredible explanatory scope of physics phenomena. But this simplicity diminishes once you need to account for interactions with the EM force field, leading to the much more "complex" [[Dirac]] equations, via [[gauge theory]]. And then the [[weak#electroweak]] introduces even more mind-boggling complexity!

Evidently, Nature is just as complex as it needs to be, and putting a strong prior on mathematical simplicity and elegance ends up running the risk of searching under the proverbial street light. Indeed, the strong appeal of the Bohmian formulation of the pilot-wave theory is that it is mathematically extremely simple, and thus can be easily proven to be identical to the standard model. But we may need to venture out into rougher terrain to fix all of the remaining issues.

