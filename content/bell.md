+++
Categories = ["Interpretations"]
bibfile = "mechphys.json"
+++

The definitive test of the [[non-locality]] of quantum physics is provided by a framework known as **Bell's inequality**, developed by [[@^Bell64]] (see [[@Bell04]]; [[@Bell66]]; [[@Bell81]] and [[@ClauserHorneShimonyEtAl69]]). The experimental tests of this logic have confirmed the quantum predictions, in ways that have also dealt with a number of possible loopholes in the logic ([[@AspectDalibardRoger82]]; [[@GiustinaVersteeghWengerowskyEtAl15]]; [[@HensenBernienDreauEtAl15]]; [[@ShalmMeyer-ScottChristensenEtAl15]]).

{id="figure_bert" style="height:25em"}
![Illustration of the standard intuition about the local hidden variable version of the Bell's inequality test, provided by J.S. Bell in his 1981 paper ([[@Bell81]]). Bell's friend Bertlemann is eccentric and always wears different color socks. Thus, if you happen to see one sock, you know the other is a different color, without having to look at it. This is perfectly sensible. However, the challenging thing about Bell's test is that your "measurement" of the color of the first sock actually _determines_ what you see: it doesn't actually _have_ a specific color until this measurement, because the socks _must_ be in a **superposition** (entangled) state to see the effect. Once you make that measurement, it "magically" influences the result that you get when you measure the other sock! This apparent non-local communication of the results of the first measurement, effectively changing the color of the sock on the other foot ("mid-step" as it were), is what boggles the mind. However, one could argue that this is merely revealing the fact that we have no idea what it means for something to be in a state of superposition, or how a measurement actually works to collapse this superposition into a distinct result.](media/fig_bertlmanns_socks.png)

The Bell's inequality framework involves computing the statistics of measurements of an **entangled** quantum system of two particles, which means that the quantum wave is in a state of **superposition**. While we can easily represent this superposition in the mathematics of the [[Hilbert space]], as we show below, it is unclear what the actual physical implication of such a superposition state is. People tend to think of it as an [[epistemic]] uncertainty, where somehow the system "really" is in some specific underlying state, but we just don't yet know what that state is.

As illustrated in [[#figure_bert]] from [[@^Bell81]], this epistemic interpretation corresponds to a **local hidden variable** theory, which was the point of the original thought experiment by [[@^EinsteinPodolskyRosen35]] (EPR) that stimulated Bell's work. Specifically, under the epistemic interpretation, the measurement process merely _reveals_ what the unknown underlying state of the system was, thereby removing our lack of knowledge. The point of Bell's framework is to show that _if_ this was actually the case, then certain experimental results should follow. By contrast, the standard quantum math actually shows qualitatively different results.

As described in [[#figure_bert]], the [[contextual]] property of quantum mechanics predicts that the process of measurement is actually changing the underlying quantum state, effectively _creating_ the resulting measurement outcome on the fly in the process of the measurement itself, rather than revealing some kind of unknown but fully pre-existing state. And because of the entangled nature of the two particles, the further prediction is that this creation process operating on one measurement of the overall quantum state should also affect the other measurement! In effect, measuring the color "pink" on one of Bertlemann's feet actually _creates_ on the fly the color of the other sock!

{id="figure_bell-packets" style="height:15em"}
![Illustration of the standard Bell test setup from the pilot-wave perspective of particles interacting with wave packets. Somehow, the waves are in a state of superposition as a result of entanglement, and this causes the results of measurement at A to influence the outcome at B. Or is it the other way around? This is a fundamental ambiguity in the framework.](media/fig_bell_packets.png)

The standard language for describing this logic is that Alice (A) and Bob (B), who are situated in labs that could be rather far apart, each make their own separate measurements on the entangled quantum state. When they later compare their results, they find that their results for events at the same point in time exhibit significant _correlations_, which reflect the influence of the measurement at A affecting the result at B (or is it vice-versa? This is a fundamental ambiguity in the framework!).

Somewhat unfortunately, the existing literature on this topic somehow seems to be missing a critical point, because it has been so focused on the local hidden variable alternative. Specifically, it seems clear that nobody really has any idea _how_ (in a physical model sense) the measurement process actually creates a measurement outcome from a superposition state (presumably because everyone pictures it in the epistemic sense). Mathematically, the key point is that the **measurement rotates the entire state space** so that it then becomes aligned with the measurement outcome. Because this rotation affects the entire space, it means that the measurement at B ends up being correlated with what happened at A.

But how does this actually work? And which side causes the rotation -- A or B? What if they each make conflicting rotations -- who goes first? How is there not an obvious violation of the speed of light as required by [[special relativity]] if this rotation can affect the entire state instantaneously, no matter how widely spread it is?

These are precisely the difficulties associated with the measurement process in the [[Copenhagen]] interpretation itself: how can it take a spatially-distributed wave function, and instantaneously collapse it down into a single particle-like point at the moment of measurement? The same non-local, instantaneous kind of assumptions are impossibly present for any kind of measurement in that framework.

Thus, to provide a more intuitive understanding of what might be happening at a physical level in the context of measurement, we need to examine how the [[pilot wave]] account explains the Bell's experiment results. In this framework, the particles retain a concrete hidden-variable like 3D position at all times, and all of the quantum effects arise from interactions with the wave field. How can that produce the observed correlations? Unfortunately, it turns out that the standard pilot wave framework _only_ tracks position in the discrete particles, while spin, which is the prime variable of interest in the Bell inequality paradigm, is purely in the wave state. So it doesn't really tell us anything. Except maybe that we really don't know what [[spin]] actually is.

Here are a few other relevant high-level points.

* Superposition is entirely a wave-specific phenomenon. Any kind of classical [[wave]] automatically exhibits superposition. And yet it still looks a bit like magic to see the superposition in effect in a classical wave.

* Extensive work has shown that classical waves, and corresponding optical wave phenomena, exhibit entanglement-like phenomena when put into a state of superposition ([[@Spreeuw98]]; [[@QianEberly11]]; [[@QianLittleHowellEtAl15]]; [[@KarimiBoyd15]]; [[@Khrennikov06]]).

* However, superposition alone is insufficient. You also need the outer-product nature of [[configuration space]], which keeps the wave state of the two different particles _separated_, even as they are in a state of superposition. So the superposition is fundamentally about the possible state of _each_ particle separately, not about both particles being somehow superposed directly on top of each other.

We return to these issues after exploring a fully worked-out example using the standard [[Hilbert space]] formalism, to get a better sense of what is happening at a mathematical level.

## Example

We'll use the modern terminology from [[quantum computing]], in terms of a **qubit**, which has a quantum binary state. In the actual experiments, this is typically either the polarization direction of a [[photon]], or the [[spin]] direction of a [[fermion]].

The _ket_ states for each qubit are denoted:

$$
|0\rangle = \begin{pmatrix}1\\0\end{pmatrix}, \qquad |1\rangle = \begin{pmatrix}0\\1\end{pmatrix}
$$

The full tensor product state space for the two-qubit space is thus a 4 dimensional vector, with the subscripts representing the 0 or 1 basis for each qubit of _a_ and _b_, which is a 2x2 matrix that is then stretched out into a 1d vector using the _Kronecker_ convention:

$$
|a \rangle \langle b | = \begin{pmatrix} a_0 b_0 & a_0 b_1 \\ a_1 b_0 & a_1 b_1 \end{pmatrix} = \begin{pmatrix} a_0 b_0 \\ a_0 b_1 \\ a_1 b_0 \\ a_1 b_1 \end{pmatrix}
$$

Thus, the four distinct pure outcome states in this outer-product state are:

$$
|00\rangle = \begin{pmatrix}1\\0\\0\\0\end{pmatrix}\quad |01\rangle = \begin{pmatrix}0\\1\\0\\0\end{pmatrix}\quad |10\rangle = \begin{pmatrix}0\\0\\1\\0\end{pmatrix}\quad |11\rangle = \begin{pmatrix}0\\0\\0\\1\end{pmatrix}
$$

The maximally entangled state of this system $\psi_\otimes$ is a **superposition** of the case where either qubit is either on or off:

$$
|\psi_\otimes\rangle = \frac{1}{\sqrt{2}}\left(|00\rangle + |11\rangle\right) = \frac{1}{\sqrt 2}\begin{pmatrix}1\\0\\0\\1\end{pmatrix}
$$

You can verify it is entangled by computing the determinant of the 2x2 matrix version of the state -- if this is non-zero, then it is entangled. The determinant $\det$ is the product of the diagonal elements minus the off-diagonal ones:

$$
C = \begin{pmatrix}a & b\\ c & d\end{pmatrix}$, \qquad det C = ad - bc
$$

$$
C_\otimes = \begin{pmatrix}1 & 0\\ 0 & 1\end{pmatrix}, \qquad \det C = 1 \neq 0
$$

As explained in [[hilbert-space#measurement]], the **density operator** $\rho$ contains all the information about all possible measurements on the given quantum state. For this entangled case, it is:

$$
\rho_\otimes = |\psi_\otimes\rangle \langle \psi_\otimes | = \begin{pmatrix} 0.5 & 0 & 0 & 0.5 \\ 0 & 0 & 0 & 0 \\ 0 & 0 & 0 & 0 \\ 0.5 & 0 & 0 & 0.5 \end{pmatrix}
$$

The critical point is that there are off-diagonal elements here, which directly capture the presence of a correlation between the probability of a $|00\rangle$ outcome and a $|11\rangle$ outcome. Specifically, in the first column, the bottom-left cell represents the probability for a $|00\rangle$ and $|11\rangle$ outcome, and the upper-right cell represents the probability of a $|11\rangle$ and $|00rangle$ outcome.

These off-diagonal elements are only possible in a state created from a superposition, which happens via the process of entanglement in preparing the initial state of the system.

There are two key contrasting cases which do _not_ produce the uniquely quantum non-locality result. The first one is a **pure product state** where for example the two qubits are both in the pure 0 state:

$$
|\psi_*\rangle = |0\rangle_A \otimes |0\rangle_B = \begin{pmatrix}1\\0\end{pmatrix}\otimes\begin{pmatrix}1\\0\end{pmatrix} = \begin{pmatrix}1\\0\\0\\0\end{pmatrix}
$$

The determinant of the corresponding 2x2 matrix is zero:

$$
C_* = \begin{pmatrix}1 & 0\\0&0\end{pmatrix}, \qquad \det C = 0
$$

and the density matrix has no off-diagonal elements:

$$
\rho_* = |\psi_*\rangle \langle \psi_* | = \begin{pmatrix} 0.5 & 0 & 0 & 0.0 \\ 0 & 0 & 0 & 0 \\ 0 & 0 & 0 & 0 \\ 0.0 & 0 & 0 & 0.5 \end{pmatrix}
$$

The more interesting foil corresponds to a local hidden variable case, where half the time a $|00\rangle$ state is prepared, and the other half it is $|11\rangle$, as compared to the _single_ state that is in a state of _superposition_. This contrast highlights precisely what it is about the superposition state that differs from a _standard_ probabilistic situation. This then provides the essential insight about entanglement.

$$
\rho_+ = \tfrac12 |00\rangle \langle00| + \tfrac12|11\rangle\langle11|
$$

$$
\rho_+ = |\psi_+\rangle \langle \psi_+ | = \begin{pmatrix} 0.5 & 0 & 0 & 0.0 \\ 0 & 0 & 0 & 0 \\ 0 & 0 & 0 & 0 \\ 0.0 & 0 & 0 & 0.5 \end{pmatrix}
$$

Critically, this also has no off-diagonal elements, and thus only produces classical-like probability results.

Thus, there is an obvious, absolute conclusion from this comparison:

* **superposition (entanglement) is an essential ingredient** in producing non-local quantum effects.

Any local hidden variable strategy that does not result in a wave state that is in a state of supervision is not going to capture the quantum effects, because, as you can see in the above math, the resulting probabilities for mixing the different states are purely classical, acting as simple weighted-average multipliers on the resulting density operator. 

Interestingly, there is indeed a literature on the effects of superposition in classical waves, i.e., in electromagnetic waves, which reproduce many of the quantum-like effects of superposition ([[@Spreeuw98]]; [[@QianEberly11]]; [[@QianLittleHowellEtAl15]]; [[@KarimiBoyd15]]; [[@Khrennikov06]]). However, these experiments do not capture the non-locality present in the Bell's test experiments, which follows naturally from the tensor-product nature of the configuration space used in the above Hilbert-space formulation.

The main question in this context is thus, how could Nature actually exhibit the corresponding non-locality in terms of a [[universal basis]] space representation that doesn't just build it in as an assumption by using a tensor space.

## The measurement observables

Because you cannot directly observe the density matrix values shown above, actual experiments involve two different detectors, A and B (typically called Alice and Bob) that each measure a spin component in the $x$–$z$ plane, at angles $\alpha$ and $\beta$. The observable is $\hat n(\theta)\cdot\vec\sigma$ with $\hat n = (\sin\theta, 0, \cos\theta)$:

$$
A(\alpha) = \sin\alpha\,\sigma_x + \cos\alpha\,\sigma_z = \begin{pmatrix}\cos\alpha & \sin\alpha \\ \sin\alpha & -\cos\alpha\end{pmatrix}
$$

and identically $B(\beta)$. Each has eigenvalues $\pm 1$ (since $A^2 = I$, $\operatorname{tr}A = 0$), with eigenvectors

$$
|{+}_\alpha\rangle = \begin{pmatrix}\cos\frac{\alpha}{2}\\[2pt] \sin\frac{\alpha}{2}\end{pmatrix},\qquad |{-}_\alpha\rangle = \begin{pmatrix}-\sin\frac{\alpha}{2}\\[2pt] \cos\frac{\alpha}{2}\end{pmatrix}
$$

Check: $A(\alpha)|{+}_\alpha\rangle = \binom{\cos\alpha\cos\frac\alpha2 + \sin\alpha\sin\frac\alpha2}{\sin\alpha\cos\frac\alpha2 - \cos\alpha\sin\frac\alpha2} = \binom{\cos(\alpha/2)}{\sin(\alpha/2)} = +|{+}_\alpha\rangle$ ✓

The spectral projectors are

$$
P^A_\pm(\alpha) = |{\pm}_\alpha\rangle\langle{\pm}_\alpha| = \tfrac12\begin{pmatrix}1\pm\cos\alpha & \pm\sin\alpha\\ \pm\sin\alpha & 1\mp\cos\alpha\end{pmatrix}, \qquad A(\alpha) = P^A_+ - P^A_-
$$

Alice measuring $\alpha$ **and** Bob measuring $\beta$ is a single four-outcome projective measurement on the joint space, with rank-1 projectors

$$
\Pi_{s_A s_B} = P^A_{s_A}(\alpha) \otimes P^B_{s_B}(\beta), \qquad s_A, s_B \in \{+,-\}
$$

These are four orthogonal rank-1 projectors summing to $I_4$. Outcome probabilities by the Born rule:

$$
P(s_A, s_B \mid \alpha,\beta) = \langle \psi_\otimes|\,\Pi_{s_A s_B}\,|\psi_\otimes\rangle = \left|\left(\langle s_A{}_\alpha| \otimes \langle s_B{}_\beta|\right)|\psi_\otimes\rangle\right|^2
$$

Using $\langle {+}_\alpha|0\rangle = \cos\frac\alpha2$, $\langle{+}_\alpha|1\rangle = \sin\frac\alpha2$, $\langle{-}_\alpha|0\rangle = -\sin\frac\alpha2$, $\langle{-}_\alpha|1\rangle = \cos\frac\alpha2$, and $|\psi_\otimes\rangle = \frac{1}{\sqrt2}(|00\rangle + |11\rangle)$:

$$
\langle {+}_\alpha {+}_\beta|\psi_\otimes\rangle = \tfrac{1}{\sqrt2}\left[\cos\tfrac\alpha2\cos\tfrac\beta2 + \sin\tfrac\alpha2\sin\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\cos\tfrac{\alpha-\beta}{2}
$$

$$
\langle {+}_\alpha {-}_\beta|\psi_\otimes\rangle = \tfrac{1}{\sqrt2}\left[-\cos\tfrac\alpha2\sin\tfrac\beta2 + \sin\tfrac\alpha2\cos\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\sin\tfrac{\alpha-\beta}{2}
$$

$$
\langle {-}_\alpha {+}_\beta|\psi_\otimes\rangle = \tfrac{1}{\sqrt2}\left[-\sin\tfrac\alpha2\cos\tfrac\beta2 + \cos\tfrac\alpha2\sin\tfrac\beta2\right] = -\tfrac{1}{\sqrt2}\sin\tfrac{\alpha-\beta}{2}
$$

$$
\langle {-}_\alpha {-}_\beta|\psi_\otimes\rangle = \tfrac{1}{\sqrt2}\left[\sin\tfrac\alpha2\sin\tfrac\beta2 + \cos\tfrac\alpha2\cos\tfrac\beta2\right] = \tfrac{1}{\sqrt2}\cos\tfrac{\alpha-\beta}{2}
$$

Writing $\delta \equiv \alpha - \beta$, the joint probabilities are

$$
P(\pm,\pm) = \tfrac12\cos^2\tfrac{\delta}{2}, \qquad P(\pm,\mp) = \tfrac12\sin^2\tfrac\delta2
$$

**Normalization:** $\cos^2\frac\delta2 + \sin^2\frac\delta2 = 1$ ✓

**Marginals (no-signalling):**

$$
P(A = +) = P(+,+) + P(+,-) = \tfrac12\cos^2\tfrac\delta2 + \tfrac12\sin^2\tfrac\delta2 = \tfrac12
$$

Alice's marginal is $\tfrac12$ *regardless of $\beta$* — Bob's setting choice is invisible to Alice. All the structure is in the correlations, none in the marginals. This is why entanglement doesn't permit signalling.

## The correlation function

$$
E(\alpha,\beta) = \langle \psi_\otimes|A(\alpha)\otimes B(\beta)|\psi_\otimes\rangle = P(+,+) + P(-,-) - P(+,-) - P(-,+)
$$

$$
= \cos^2\tfrac\delta2 - \sin^2\tfrac\delta2 = \boxed{\cos(\alpha - \beta)}
$$

A cross-check without projectors: for $|\psi_\otimes\rangle = \frac{1}{\sqrt2}\sum_i|ii\rangle$ one has $\langle\psi_\otimes|M\otimes N|\psi_\otimes\rangle = \frac12\operatorname{tr}(MN^T)$, and $\operatorname{tr}\!\big(A(\alpha)B(\beta)\big) = 2\cos(\alpha-\beta)$. ✓

# 6. CHSH and the classical bound

With two settings each ($a, a'$ for Alice; $b, b'$ for Bob):

$$
S = E(a,b) - E(a,b') + E(a',b) + E(a',b')
$$

**Local hidden variable bound.** If each run carries predetermined values $A_a, A_{a'}, B_b, B_{b'} \in \{-1,+1\}$ fixed by some $\lambda$, then for each $\lambda$:

$$
S(\lambda) = A_a(B_b - B_{b'}) + A_{a'}(B_b + B_{b'})
$$

Since $B_b, B_{b'} = \pm1$: either $B_b = B_{b'}$, making the first bracket $0$ and the second $\pm2$; or $B_b = -B_{b'}$, making the second $0$ and the first $\pm2$. Either way $|S(\lambda)| \le 2$, and averaging over $\lambda$ preserves it:

$$
\boxed{|S_{\text{LHV}}| \le 2}
$$

**Quantum prediction.** With $E = \cos(\alpha-\beta)$:

$$
S = \cos(a-b) - \cos(a-b') + \cos(a'-b) + \cos(a'-b')
$$

Choosing $a = 0,\ a' = \frac\pi2,\ b = \frac\pi4,\ b' = \frac{3\pi}{4}$ makes all four terms contribute $+\frac{\sqrt2}{2}$:

$$
S = 4\cdot\tfrac{\sqrt2}{2} = 2\sqrt2 \approx 2.8284
$$

# 7. Now with explicit numbers

## The four observables

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

## The four 4×4 tensor products

Shown multiplied by $\sqrt2$ so the entries are integers. Acting on $|\psi_\otimes\rangle = \frac{1}{\sqrt2}(1,0,0,1)^T$:

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

$$
Mv = \tfrac{1}{\sqrt2}\cdot\tfrac{1}{\sqrt2}\begin{pmatrix}1\cdot1 + 1\cdot 0\\ 1\cdot1 - 1\cdot0 \\ -1\cdot 0 - 1\cdot 1\\ -1\cdot0 + 1\cdot1\end{pmatrix} = \tfrac12\begin{pmatrix}1\\1\\-1\\1\end{pmatrix}
$$

$$
v^T(Mv) = \tfrac{1}{\sqrt2}(1,0,0,1)\cdot\tfrac12(1,1,-1,1)^T = \tfrac{1}{2\sqrt2}(1 + 1) = \tfrac{1}{\sqrt2} = 0.707107 \;\checkmark
$$

## All joint probabilities

| settings | $\delta = \alpha-\beta$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | sum | $E$ |
|---|---|---|---|---|---|---|---|
| $(a,b) = (0°,45°)$ | $-45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |
| $(a,b') = (0°,135°)$ | $-135°$ | 0.073223 | 0.426777 | 0.426777 | 0.073223 | 1.000000 | $-0.707107$ |
| $(a',b) = (90°,45°)$ | $+45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |
| $(a',b') = (90°,135°)$ | $-45°$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 1.000000 | $+0.707107$ |

Every marginal is exactly $0.500000$ for both parties in all four rows.

Sample check of row 1: $\frac12\cos^2(-22.5°) = \frac12(0.923880)^2 = \frac12(0.853553) = 0.426777$ ✓, and $\frac12\sin^2(-22.5°) = \frac12(0.382683)^2 = 0.073223$ ✓.

## Assembling $S$

$$
S = (+0.707107) - (-0.707107) + (0.707107) + (0.707107)
$$

$$
= 0.707107 \times 4 = \mathbf{2.828427}
$$

$$
2\sqrt2 = 2.828427\ \checkmark \qquad\text{versus the classical bound } 2
$$

Violation ratio: $2\sqrt2 / 2 = \sqrt2 \approx 1.414$.

Form the CHSH operator on $\mathbb{C}^4$:

$$
\hat S = A\otimes B - A\otimes B' + A'\otimes B + A'\otimes B'
$$

Applying it to the Bell state numerically gives $\hat S|\psi_\otimes\rangle = 2\sqrt2\,|\psi_\otimes\rangle$ — **the Bell state is an exact eigenvector of the CHSH operator at these settings**, with eigenvalue $2\sqrt2$. That's why the violation is maximal: you can't do better than an eigenvalue, and Tsirelson's bound says no eigenvalue is larger. The one-line reason (Landau's identity):

$$
\hat S^2 = 4\,I \;+\; [A,A']\otimes[B,B']
$$

Since $\|A\| = \|A'\| = 1$, $\|[A,A']\| \le 2$, so $\|\hat S^2\| \le 4 + 4 = 8$ and $\|\hat S\| \le 2\sqrt2$. If $A$ and $A'$ commuted — the classical case, where both have simultaneous definite values — the commutator term vanishes and you recover $|S| \le 2$ exactly. **The violation is precisely the non-commutativity of the two measurement settings, amplified through the tensor product structure.**

## The correlation factorizes — symbolically

This is the whole difference, in one line. Because the state is a product, the expectation of any product observable splits:

$$
E(\alpha,\beta) = \langle 0|\otimes\langle 0|\;A(\alpha)\otimes B(\beta)\;|0\rangle\otimes|0\rangle = \underbrace{\langle 0|A(\alpha)|0\rangle}_{\text{Alice alone}}\cdot\underbrace{\langle 0|B(\beta)|0\rangle}_{\text{Bob alone}}
$$

With $\langle 0|A(\alpha)|0\rangle = A(\alpha)_{00} = \cos\alpha$:

$$
\boxed{E_{\text{prod}}(\alpha,\beta) = \cos\alpha\,\cos\beta} \qquad\text{versus}\qquad E_{\psi_\otimes}(\alpha,\beta) = \cos(\alpha-\beta)
$$

**This is the structural difference.** $\cos\alpha\cos\beta$ is a *separable* function — a product of a function of $\alpha$ alone and a function of $\beta$ alone. $\cos(\alpha-\beta)$ is not; it cannot be factored, because $\cos(\alpha-\beta) = \cos\alpha\cos\beta + \sin\alpha\sin\beta$ has two terms. That second term $\sin\alpha\sin\beta$ is the entire Bell violation.

## The joint probabilities factorize too

$$
P(s_A,s_B) = \left|\langle s_A{}_\alpha|0\rangle\right|^2 \cdot \left|\langle s_B{}_\beta|0\rangle\right|^2 = P(s_A)\cdot P(s_B)
$$

with $P(A{=}+) = \cos^2\frac\alpha2$, $P(B{=}+) = \cos^2\frac\beta2$. The outcomes are **statistically independent** — the joint distribution is literally the product of the marginals, for every setting pair. There is nothing to correlate.

## Numbers at the Bell-optimal angles

| settings | $E$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | $P(A{=}{+})$ | $P(B{=}{+})$ |
|---|---|---|---|---|---|---|---|
| $(a,b)=(0°,45°)$ | $+0.707107$ | 0.853553 | 0.146447 | 0.000000 | 0.000000 | 1.0000 | 0.8536 |
| $(a,b')=(0°,135°)$ | $-0.707107$ | 0.146447 | 0.853553 | 0.000000 | 0.000000 | 1.0000 | 0.1464 |
| $(a',b)=(90°,45°)$ | $\;\;0.000000$ | 0.426777 | 0.073223 | 0.426777 | 0.073223 | 0.5000 | 0.8536 |
| $(a',b')=(90°,135°)$ | $\;\;0.000000$ | 0.073223 | 0.426777 | 0.073223 | 0.426777 | 0.5000 | 0.1464 |

Factorization check on row 3: $P(+,+) = 0.5 \times 0.853553 = 0.426777$ ✓ exactly.

$$
S = 0.707107 - (-0.707107) + 0 + 0 = \mathbf{1.414214} = \sqrt2
$$

Note that Alice's marginal here *does* depend on her own setting ($1.0$ vs $0.5$) — unlike the Bell state, where it was always $0.5$. That's a giveaway that this isn't a fair comparison. Hence:

So no local measurement on either qubit alone can tell them apart. Only the *correlations across settings* can.

## The correlation function

$$
E_{\text{mix}}(\alpha,\beta) = \operatorname{Tr}\!\left[\rho_{\text{mix}}\,A(\alpha)\otimes B(\beta)\right] = \tfrac12\langle00|A\otimes B|00\rangle + \tfrac12\langle11|A\otimes B|11\rangle
$$

$$
= \tfrac12\big[\cos\alpha\cos\beta\big] + \tfrac12\big[(-\cos\alpha)(-\cos\beta)\big] = \boxed{\cos\alpha\,\cos\beta}
$$

Same separable form as the pure product state.

## Numbers

| settings | $E_{\text{mix}}$ | $P(+,+)$ | $P(+,-)$ | $P(-,+)$ | $P(-,-)$ | $P(A{=}{+})$ |
|---|---|---|---|---|---|---|
| $(a,b)=(0°,45°)$ | $+0.707107$ | 0.426777 | 0.073223 | 0.073223 | 0.426777 | 0.5000 |
| $(a,b')=(0°,135°)$ | $-0.707107$ | 0.073223 | 0.426777 | 0.426777 | 0.073223 | 0.5000 |
| $(a',b)=(90°,45°)$ | $\;\;0.000000$ | 0.250000 | 0.250000 | 0.250000 | 0.250000 | 0.5000 |
| $(a',b')=(90°,135°)$ | $\;\;0.000000$ | 0.250000 | 0.250000 | 0.250000 | 0.250000 | 0.5000 |

$$
S = 0.707107 - (-0.707107) + 0 + 0 = \mathbf{1.414214} = \sqrt2
$$

## The diagnostic: where exactly they diverge

Put the two states' correlations side by side:

| settings | $E$ for $|\psi_\otimes\rangle$ | $E$ for $\rho_{\text{mix}}$ | same? |
|---|---|---|---|
| $(a,b) = (0°,45°)$ | $+0.707107$ | $+0.707107$ | ✅ **identical** |
| $(a,b') = (0°,135°)$ | $-0.707107$ | $-0.707107$ | ✅ **identical** |
| $(a',b) = (90°,45°)$ | $+0.707107$ | $\;\;0.000000$ | ❌ |
| $(a',b') = (90°,135°)$ | $+0.707107$ | $\;\;0.000000$ | ❌ |
| | $S = 2.828427$ | $S = 1.414214$ | |

Look at the first two rows: for Alice's setting $a = 0°$ (i.e. $A = \sigma_z$), the mixture reproduces the Bell state's joint probabilities **exactly, to every decimal place** — not just the correlation but all four $P(s_A,s_B)$. The classical mixture is a perfect impostor as long as Alice measures $\sigma_z$.

It fails the moment Alice switches to $a' = 90°$, where $A' = \sigma_x$. $\sigma_x$ is off-diagonal, so its expectation reads the *coherences* — precisely the matrix elements that $\rho_{\text{mix}}$ lacks. Both of those rows collapse to $E = 0$, killing half the CHSH sum.

**The single cleanest statement of the difference.** Measure both parties along the *same* in-plane angle $\theta$:

$$
E_{\psi_\otimes}(\theta,\theta) = \cos(\theta-\theta) = 1 \quad\text{for every }\theta$$
$$E_{\text{mix}}(\theta,\theta) = \cos^2\theta \quad\Rightarrow\quad 1 \text{ at } 0°,\; 0.5 \text{ at } 45°,\; 0 \text{ at } 90°
$$

The Bell state is perfectly correlated in *every* basis. The mixture is perfectly correlated in *one* basis and uncorrelated in the conjugate one. That basis-independence is what has no classical analogue, and CHSH is the device that detects it by querying two incompatible bases at once.

## Why no separable state can ever reach 2

For $\rho_{\text{mix}}$ the local hidden variable model is not hypothetical — write it down:

- $\lambda \in \{0, 1\}$, each with probability $\tfrac12$ (which pair was emitted).
- Given $\lambda$, Alice outputs $+1$ with probability $|\langle{+}_\alpha|\lambda\rangle|^2$, using only $\alpha$ and $\lambda$.
- Given $\lambda$, Bob independently outputs $+1$ with probability $|\langle{+}_\beta|\lambda\rangle|^2$, using only $\beta$ and $\lambda$.

Then $E(\alpha,\beta) = \frac12[\cos\alpha\cos\beta] + \frac12[(-\cos\alpha)(-\cos\beta)] = \cos\alpha\cos\beta$ ✓ — reproduced exactly, with no communication. Since an LHV model exists, the $|S| \le 2$ proof from before applies verbatim.

Maximizing over *all* settings for $E = \cos\alpha\cos\beta$:

$$
S = \cos a\,(\cos b - \cos b') + \cos a'\,(\cos b + \cos b')
$$

$$
\max S = |\cos b - \cos b'| + |\cos b + \cos b'| = 2\max(|\cos b|, |\cos b'|) \le 2
$$

The bound $2$ is reachable (e.g. $a = 0°, b = 0°, b' = 180°$) — but only by the degenerate choice $B' = -B$, where Bob isn't really using two independent settings. Saturated, never exceeded.

## Summary

| | $|\psi_\otimes\rangle$ | $|0\rangle|0\rangle$ | $\rho_{\text{mix}}$ |
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


