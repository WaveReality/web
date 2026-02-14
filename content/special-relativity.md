+++
bibfile = "mechphys.json"
+++

{id="figure_contract" style="height:20em"}
![Intuitive explanation of Lorentz contraction of space that occurs as something moves faster, in terms of the relationship between wave frequency and speed in the Klein-Gordon wave equation. Because faster movement is associated with higher frequency and shorter wavelength, the system contracts in the direction of motion as it speeds up. Thus, any measurements made in the faster system will have their basic constituents, including yard sticks and everything else, shrunk in this way. This is one of the main effects of special relativity.](media/fig_wave_contract.png)

{id="figure_lorentz" style="height:20em"}
![The Lorentz Transformation, a central property of special relativity, which causes length in the direction of motion to shrink and time to expand (dilate) as a function of relative speed, in just such a way as to preserve the observed speed of light regardless of how fast one is going. The matter wave equation exhibits exactly this behavior, which completely masks any fixed matrix in which such waves might be implemented. In this example, a speeding light ray is observed at a given time t by an observer in a train speeding along at 86.6% of the speed of light $(c \approx 3.0x10^8)$, and by "us" sitting in a stopped train (on a siding presumably). All the measurements in black are what we observe in this stopped reference frame, while those in red are what the speeding train guy observes. If we wait 100 nanoseconds (ns) ($1x10^{-7}$ seconds --- 100 times slower than the clock rate on a 1Ghz computer chip), then this light ray will have moved 30 meters. However, from our stopped perspective, the speeding train will be partially keeping up with the light ray, so that it will appear to have traveled only 4m relative to the moving train. Thus, in this stopped reference frame, where 100ns have passed for this light to appear to have traveled 4m, we might naively assume that someone on the speeding train would measure the speed of light as only $4x10^{7} m/s$ --- oops!  But the Lorentz transformations of length and time exactly compensate. The length of the train in the direction of motion shrinks in half, so that people on the train measure the 4m in the stopped reference frame as 8m in the moving reference frame --- twice as long. Furthermore, time moves more slowly for the speeding train, such that the 100ns in our reference frame is measured as only 50ns in the speeding train reference frame (at a static reference point in the speeding train, which is the very back of the coal tender in this example). The measurement of time is very strange in special relativity, because what is observed as occurring at the same time (simultaneity) across different reference frames depends on both time ''and location''. Thus, when the light ray is measured at 8m ahead of the back of the coal tender, this registers as only 26.8ns of elapsed time!  If you divide this 8m by that amount of time, it comes out to exactly the same speed of light as in the stopped frame. The time transformation equation is: $t' = \gamma (t-vx/c^2)$ and the position transformation is: $x' = \gamma (x-vt)$, where t' and x' are as measured on the speeding train and t, x are on the stopped one, and $\gamma = 1/\sqrt{1-v^2/c^2}$.](media/fig_lorentz_xform.png)

**Special relativity** is the first version of Einstein's famous relativity theories, published in 1904. Interestingly, most of the mathematical principles and phenomena were already well established prior to this publication. What the theory added was a comprehensive new framework for understanding these otherwise disparate phenomena.

Apparently Einstein was motivated to come up with special relativity in part by thinking about what it would be like to catch up with a beam of light. Turns out you can't: light always moves at the same speed away from you, no matter how fast you are going. Thus, the [[waves]] described by the basic wave equation are always speeding along at the same speed, and are relativistic in a fairly straightforward but also somewhat uninteresting way: they just cruise around at the same speed of light, and not much else can be said for them.

In contrast, things that have a non-zero rest mass are subject to three major effects as their velocities increase ([[#figure_contract]], [[#figure_lorentz]]; all of these are as measured relative to a static observer watching such a thing whiz by --- in our model, we can conveniently use the underlying grid reference frame):

* They shrink along the direction of travel (Lorentz contraction).
* Time slows down (time dilation).
* Apparent (relativistic) mass increases.

Almost miraculously, all of these effects can be derived directly from the [[Klein-Gordon]] equation. Intuitively, the contraction of space along the direction of travel occurs because the length of a given thing is measured in terms of the wavelengths of the underlying particles that compose it. As these things move with greater velocity, their frequency increases and their wavelengths decrease, and this results in a shorter overall length (as viewed from an outside observer).

Similarly, the increase in apparent mass comes from the increased difficulty in further accelerating the system. Mass is measured in terms of the inertia of a system: its resistance to further acceleration. It happens that the amount of force required to increase the frequency of a given wave packet further is proportional to its existing frequency, so that it becomes harder and harder to accelerate something as it moves faster and faster. This is manifest as an increase in the apparent mass of the particle. The increase in apparent mass of the particle also has the effect of causing the apparent time to slow down: time is fundamentally measured as a rate of change, which depends on accelerations, which are affected by the increase in apparent mass.

The very stretchiness and malleability of space and time postulated by special relativity is a very strong argument in favor of waves over particles. Waves are inherently stretchy and malleable (as just invoked above to explain the key features of special relativity), whereas particles as hard discrete points just don't seem to fit with this picture. What would it mean for a point particle to contract along its direction of travel?

One of the main puzzles that people often have with special relativity is understanding how light can appear to move at the same fixed speed, when an observer is screaming along at nearly the speed of light. You would think that you would catch up to a photon (as in Einstein's motivating thought experiment), and even if you couldn't see it actually standing still, it sure should seem to move more slowly relative to a speeding observer than to someone sitting still.

The trick is, as an observer speeds up, the time dilation and space contraction effects conspire to alter the perception of the relative motion of other things, as illustrated in [[#figure_lorentz]]. The wavelength and velocity of the moving light wave never actually change (they are always the same relative to the underlying reference frame of the grid), but relative to the moving observer, they appear to be much faster. This is because the observer's rulers have shrunk, and their clocks have slowed down, so that the same light wave appears to be moving much faster to them.

In contrast, a static observer watching the moving observer chasing after the light beam will perceive the two as having a smaller relative velocity difference. These changes occur in exactly the right way to make the light wave appear to travel at the speed of light to the moving observer (which in fact it is).

## Details

{id="figure_lorentz2" style="height:20em"}
![Illustration of Lorentz transformation between a static reference frame F (which is the same reference frame as the page, for example), and a frame F' moving at velocity $v$ relative to F. Both frames contain a rigid rod, which when F' was at rest relative to F were the same length $l$, as indicated by the ruler marks. These rods now appear to be of different lengths in the two frames --- from the perspective of F, the rod in F' appears to have shrunk to a shorter length $l'$. Interestingly, from the perspective of an observer moving along with the F' reference frame, its rod appears to be of length $l$ (same as when it was stationary), and the other rod in F appears to have shrunk.](media/fig_lorentz_xform_small.png)

{id="figure_lorentz-factor" style="height:20em"}
![The amount of shrinkage as a function of velocity $v$ is determined by the Lorentz factor $\gamma$, which is plotted here (in natural units where the speed of light $c = 1$). Not much happens until you get very close to the speed of light (e.g., above 90% or .9).](media/fig_lorentz_factor.png)

{id="figure_lorentz-coord" style="height:20em"}
![Lorentz coordinate transformations, plotted in one spatial coordinate (x) and time (vertical axis). In the resting frame F, two points (think of them as marbles) just sit motionless, and thus form vertical trajectories through increasing time. When these are transformed into a reference frame F' (aligned at point (t{{=}}0,x{{=}}0) with F) moving to the left at $v {{=}} -.866$ (Lorentz factor $\gamma {{=}} 2$), several interesting features of the Lorentz transformation are evident. If we follow the marble that was originally located at (0,x) as it sits in frame F for 3 time steps, we see that it appears to move to the right in frame F', in the opposite direction of F' motion. Furthermore, because of the Lorentz factor, the 3 seconds in frame F amount to 6 seconds in F', and the distance it should travel due to the relative motion, computed in frame F ($-vt {{=}} .866 \times 3 {{=}} 2.598$) corresponds to 5.196 in frame F'. The second marble reveals a critical and somewhat counter-intuitive effect, where two events that are {\em simultaneous} in frame F (i.e., t{{=}}0 for both of these), occur at ''different times'' in frame F'. Specifically the second marble sitting at rest at x{{=}}2 at t{{=}}0 is not &quot;encountered&quot; by the moving frame F' until 3.464 time seconds later (in F' time units), due to it being offset in space from the first marble. It takes the frame F' $(-vx {{=}} .866 * 2 {{=}} 1.732)$ time units to get to this second point (in the units of the F frame), and when this gets subject to the time dilation effect, you end up with the 3.464. From this starting point for the second marble, the same time and space increments as for the first marble occur for the subsequent point 3 time units later.](media/fig_lorentz_xform_coord.png)

{id="figure_lorentz-coord-c" style="height:20em"}
![An alternative situation to the previous Figure, where the points in frame F are now moving at the speed of light (indicated by their slopes being 1; they are now photons instead of marbles). The Lorentz conversions preserve these slopes, so that the speeds are still 1 in the F' frame. Thus, the speed of light is always the same to all observers.](media/fig_lorentz_xform_coord_c.png)

Here we provide a more thorough treatment of **special relativity**, including the math needed to transform between different reference frames. We never actually have to use this in our numerical simulations, because all of our wave equations are automatically consistent with special relativity (i.e., they are **manifestly covariant**). But it is important to understand how all this works in any case.

The mathematics of special relativity involves the computation of relative distance and time measurements in different reference frames that are moving relative to each other ([[#figure_lorentz2]]). As illustrated, a rigid rod of the same length $l$ can appear shortened if it is moving along at a high velocity $v$ relative to a static reference frame F (i.e., not moving relative to this page of paper). We denote the moving rod's reference frame F'.

The factor for transforming between reference frames is the **Lorentz factor** gamma:

{id="eq_factor" title="Lorentz factor"}
$$
\gamma = \frac{1}{\sqrt{1 - \frac{v^2}{c^2}}}
$$

The shape of this function is shown in [[#figure_lorentz-factor]]. Because the velocity enters into this function as a squared term, the function has a parabolic shape, such that not much happens until the velocity gets very close to the speed of light.

For the situation illustrated in , the moving rod in frame F' appears shortened in the static frame F by a factor of $\frac{1}{\gamma}$:

{id="eq_contract" title="Lorentz contraction"}
$$
l' = \frac{l}{\gamma}
$$

Similarly, the duration of a given fixed interval $t$ (e.g., a second) in the moving reference frame F' appears longer from the perspective of F by the same factor:

{id="eq_dilate" title="Time dilation"}
$$
t' = \gamma t
$$

It is important to keep in mind that in these equations, the $l'$ and $t'$ refer to what something in the moving F' reference frame looks like to someone in the static frame, F, relative to these same quantities as measured in the static frame ($l$ and $t$). But this assumes that we have previously established in a common reference frame that the rod in F' actually has the same length as the one in F (and the second similarly has the same duration).

An alternative (and more conventional, but sometimes more confusing) way of using the prime notation is to directly convert between coordinate systems of the two reference frames, where the primed and unprimed cases both refer to _the exact same event_ from the two different perspectives. In this case $l'$ would refer to how an observer {\em in F'} would measure the rod length, whereas $l$ refers to what someone in F would measure _for the very same rod that is moving in F'_, not for the "standard length" of the rod when both reference frames where static.

In this case, things are exactly flipped, and we would say that $l'$ is the original rod length (say 1 meter), because it is in the F' reference frame that the rod is not moving, whereas in the F reference frame, the rod has shrunk to a shorter apparent length $l$ (say .5 meters, if $v=.866$). Similarly, the second measured in F' ($t'$) corresponds to a longer time interval $t$ in F (e.g., 2 seconds). It is definitely complicated to keep track of all this moving back and forth between reference frames!

In this way of doing things, in frame F, we designate an event as occurring at space-time location $(t,x,y,z)$ (this is a **four-vector** or **space-time coordinate** in _Minkowski_ space, which we introduced at the end of [[Maxwell]]. In frame F', this same event has coordinates $(t',x',y',z')$, where the two coordinate systems are aligned such that the origin (0,0,0,0) is the same in both. As before, we specify that the relative velocity $v$ between the two frames is entirely along the $x$ axis, for simplicity. We can compute these F' coordinates directly from our F coordinates, using Lorentz transformations (again in natural units where $c=1$, and $v$ goes between 0 and 1):

* $t' = \gamma (t - vx) $
* $x' = \gamma (x - vt) $
* $y' = y $
* $z' = z $

[[#figure_lorentz-coord]] shows where these formulas come from, and demonstrates their concrete application. The $-vx$ term in the transformation of time, and the $-vt$ term in the transformation of the $x$ coordinate, are critical and possibly counter-intuitive factors in these equations. As the figure explains, the $-vt$ is more sensible, as it simply reflects the relative motion of the reference frame over time. The $-vx$ term is somewhat less intuitive, but it reflects the fact that two events separated by some spatial distance will not be experienced as simultaneous in a moving reference frame! Instead, the spatial separation turns into a temporal separation due to the relative motion of the frame.

One can use these equations to compute the transformed velocity of an object moving at velocity $w$ along the x axis in frame F as it would appear to an observer in frame F'. We can compute this by taking velocity as distance over time, and noting that in time $t$ in frame F, the object will travel a distance $x = wt$. Thus, we need to compute the transformed velocity as distance over time in the F' frame, as:

- $w' = \frac{x'}{t'} $
- $= \frac{\gamma (x - v t)}{\gamma (t - v x)} $
- $= \frac{\gamma (wt - v t)}{\gamma (t - v w t)} $
- $= \frac{\gamma t (w - v)}{\gamma t (1 - v w)} $
- $= \frac{w - v}{1 - wv} $

(the intermediate steps are included to make everything as clear as possible, tracking the substitution of $x$ with $wt$).

Interestingly, if both velocities $w$ and $v$ are small (relative to the speed of light, $c=1$), then the denominator is close to 1, and you get the more intuitive result that the relative velocities just add (or subtract, as the case may be). This is known as the **Galilean transformation**, which holds for non-relativistic speeds, and is intuitive to most people. However, as either of these speeds increase, the denominator starts to get smaller, and the relative velocities do not add linearly. Here are several illustrative examples of how this equation works:

|          |         |           |                         |
|----------|---------|-----------|-------------------------|
| w = 0    | w' = -v |           | only relative motion    |
| w = v    | w' = 0  |           | same speed              |
| w = .2   | v = .1  | w' = .102 | slightly faster than .1 |
| w = -.1  | v = .1  | w' = .199 | slightly slower than .2 |
| w = .9   | v = .1  | w' = .87  | much faster than .8     |
| w = -.9  | v = .1  | w' = .99  | much slower than 1.1    |
| w = 1.0  | v = .1  | w' = 1.0  | speed of light is same  |
| w = -1.0 | v = .1  | w' = 1.0  | speed of light is same  |

One other important point to be made about special relativity is that there is an invariant calculation that can be made on the coordinates measured in any given reference frame, which will yield the same result as in any other reference frame. This is a kind of distance metric between two points $(t_1, x_1, y_1, z_1)$ and $(t_2, x_2, y_2, z_2)$ (both of which must be measured in the same reference frame):

- $ds^2 = (t_1 - t_2)^2 - (x_1 - x_2)^2 - (y_1 - y_2)^2 - (z_1 - z_2)^2 $

You can try this out on the coordinates in for the same points in the different reference frames: you'll get the same results for either (with small differences due to round-off errors). Interestingly, when you try it for the second case with the photons moving at the speed of light, you see that this value is always 0. Thus, in effect, the distance in space is equal to the distance in time. If we just have motion in one spatial dimension (e.g., $x$), this is clear:

- $(t_1 - t_2)^2 - (x_1 - x_2)^2 = 0 $
- $(t_1 - t_2)^2 = (x_1 - x_2)^2 $

This is exactly the same time = space relationship that we discussed earlier for the basic wave equation. Perhaps now the fundamental importance of this for relativity is clearer.

This distance metric also tells you in what way two events are separated. If the distance metric is positive, then the two points are separated by time, and if it is negative, they are separated by space.

Interestingly, the Lorentz transformations can be derived directly from this distance metric, and this distance metric can in turn be derived from the principle that the speed of light is a constant in any reference frame. Therefore, all of these effects are really just different manifestations of the same basic set of constraints, which are captured directly and automatically in the wave equations.

One additional property of special relativity has to do with the relationship between energy and momentum. If an object is at rest, its only energy is that associated with its rest mass, according to the famous equation:

- $E_0 = m_0 c^2 $

As the object moves faster, it gains energy in proportion to the Lorentz factor $\gamma$:

- $E = \gamma m_0 c^2 $

Furthermore, the relativistic momentum of this object is just the Lorentz factor times the standard Newtonian definition of momentum for motion at velocity $\vec{v}$:

- $\vec{p} = \gamma m_0 \vec{v} $

These two variables can be related in the relativistic energy-momentum equation:

- $E^2 = \vec{p}^2 c^2 + (m_0 c^2)^2 $
- $E = \sqrt{\vec{p}^2 c^2 + (m_0 c^2)^2} $

which is different way of expressing the energy of the system. It is very sensible, in that the two main contributors to energy are the momentum (i.e., kinetic energy) and the rest energy. If you set $c=1$, it is even simpler:

- $E^2 = \vec{p}^2+ m_0^2 $
- $E = \sqrt{\vec{p}^2+ m_0^2} $

You can think of the rest energy and momentum as two legs of a right triangle, such that the total energy is the hypotenuse, according to the pythagorean theorem.

In the limit of a slow velocity (relative to the speed of light), this expression approaches the Newtonian expression for kinetic energy (plus the rest mass energy):

- $E \approx \frac{1}{2}m_0 \vec{v}^2 + m_0 c^2 $

Interestingly, the Klein-Gordon equation can be derived directly from this equation!

Finally, we note that the warping of space and time that produces gravitational effects according to general relativity enters into our model in a very different way, as described later. Thus, in our model, special and general relativity result from two very different mechanisms.

## Relativistic Momentum and Velocity

This page shows all the mathematical steps needed to go from the relativistic relationship between momentum and velocity, to an expression for the velocity as a function of the wavelength $\lambda$, rest mass $m_0$, and the other standard constants (c, h). The $\gamma$ is the Lorentz factor used throughout special relativity.

- $p = \gamma m_0 v$ (starting equation)
- $v = \frac{1}{\gamma m_0} p$ (flip to solve for v)
- $v = \frac{\sqrt{1 - \frac{v^2}{c^2}}} {m_0} \frac{h}{\lambda}$ (substitute definition of p in terms of wavelength)
- $v^2 = \frac{(1 - \frac{v^2}{c^2})h^2} {m_0^2 \lambda^2}$ (square everything)
- $v^2 = \frac{h^2} {m_0^2 \lambda^2} - \frac{v^2 h^2} {c^2 m_0^2 \lambda^2}$ (separate terms)
- $v^2 + \frac{v^2 h^2} {c^2 m_0^2 \lambda^2} = \frac{h^2} {m_0^2 \lambda^2}$ (consolidate v's on lhs)
- $\frac{v^2 c^2 m_0^2 \lambda^2 + v^2 h^2} {c^2 m_0^2 \lambda^2} = \frac{h^2} {m_0^2 \lambda^2}$ (more consolidation)
- $\frac{v^2 c^2 m_0^2 \lambda^2 + v^2 h^2} {c^2} = h^2$ (eliminate redundant denominators)
- $\frac{v^2 (c^2 m_0^2 \lambda^2 + h^2)} {c^2} = h^2$ (pull out v term uniquely)
- $v^2 = \frac{h^2 c^2}{c^2 m_0^2 \lambda^2 + h^2}$ (multiply by inverse of v factor to move to other side)
- $v = \frac{h c}{\sqrt{c^2 m_0^2 \lambda^2 + h^2}}$ (final expression --- take square root)


