+++
Name = "Hilbert space"
Categories = ["Standard Model"]
bibfile = "mechphys.json"
+++

By far the most widely-used calculational framework in standard QM is the algebraic **matrix mechanics** approach, pioneered by Heisbenberg, Dirac, Hilbert, von Neumann and others in the mid 1920s. It involves _state vector_ representations of the state of a system, encoded via complex-valued vectors representing _probability amplitudes_ (i.e., a '**Hilbert space**). This state vector is a specific way of encoding the [[configuration space]] of the entire set of relevant variables, and is thus manifestly [[non-locality|non-local]], and represents the entire state a given point in time, in a way that is thus incompatible with the principles of relativity.

This state vector evolves under _unitary_ transformations (rotations in the complex vector space), which preserve the overall magnitudes of the vectors, even as they rotate around in the space. The unitary nature of the rotation transformations represents the behavior of the system when it is being governed by the [[Schrodinger]] wave dynamics under the Copenhagen dualistic framework, which perfectly preserves the overall underlying probability space as long as nobody "looks at it the wrong way" (i.e., makes a measurement). Then, at the end, a "measurement" is made by collapsing the probability space down to a single discrete outcome (i.e., along an eigenvector of the resulting state). 

This matrix formalism is equivalent to a self-consistent form of probability theory, which can be derived from abstract axioms having nothing to do with quantum physics ([[@Gleason75]]; [[@Jaynes90]]; [[@CavesFuchsSchack02]]; [[@FuchsMerminSchack14]]; [[@Mermin18]]). Indeed, this framework is so general that its only real physical commitment is that quantum physics obeys strict [[conservation]] laws, due to the use of unitary state evolution operators.

Thus, the claim that standard QM is such a successful framework must be understood within this context: yes, it is accurate in capturing this basic fact of conservation, but it really isn't going very far out on a limb here: nothing wagered, nothing lost; but also perhaps not so much gained.

## Bra-Ket notation and vector spaces

See [[bell]] for a concrete application of all this to the Bell's inequality test of [[non-locality]].

The basic element of the Hilbert space is an _abstract_ vector, which basically means that it has a _dimensionality_ but, unlike a more concrete vector defined in 3D coordinate space, _the basis vectors of the space are arbitrary._ The only thing that _is_ defined is a rule for computing the **inner product** (also known as the **dot product**), which, notably, is invariant with respect to any arbitrary set of orthogonal, normal-length basis vectors. For real-valued, two-dimensional vectors $\mathbf{x}$ and $\mathbf{y}$, the inner product is just the sum of the products of the values:

$$
\langle \mathbf{x}, \mathbf{y} \rangle = x_1 y_1 + x_2 y_2
$$

This inner product defines a _distance metric_ for the space, where the distance between the two vectors is the square root of this inner product.

In quantum applications, the elements of the vectors are [[complex number]]s, and full understanding of these is required before proceeding. In this case, the inner product is defined as multiplication by the complex conjugate, e.g., for $z = (z_1, z_2) = ({z_1}_r + i{z_1}_i, {z_2}_r + i{z_2}_i)$ where the _r_ and _i_ subscripts denote the real and imaginary components of each value, and likewise for $w$:

$$
\langle z, w \rangle \equiv z^* w = z^*_1 w_1 + z^*_2 w_2 = {z_1}_r {w_1}_r + {z_1}_i {w_1}_i + {z_2}_r {w_2}_r + {z_2}_i {w_2}_i
$$

This is equivalent to the distance you'd get in a 4D space, so the complex values just provide extra "internal" dimensions to the space, in effect.

The _bra-ket_ notation introduced by Paul Dirac provides a compact way of denoting one of these abstract vectors, which typically goes in the "ket" (right-hand side of the word "bra[c]ket") position, along with an **operator** that performs some kind of operation on the vector, which is the "bra" and it goes on the left:

$$
\mbox{bra} | \mbox{ket} \rangle
$$

(note that the left angle bracket $\langle$ is omitted here -- that is reserved for computing the inner product length measure).

For example, the wave function state $\psi$ might be a _ket_ vector: 

$$
| \psi \rangle
$$

The time-evolution of that wave state according to the Schrodinger equation would be indicated by putting this _unitary_ function that transforms the state from one point in time to another, $U(t,t_0)$ into the _bra_ position:

$$ 
U(t,t_0) | \psi \rangle
$$

We'll see how this actually works in [[#time evolution]] below.

You can also just put a constant into the _bra_ position, and you can use arbitrary symbols for the _kets_:

$$
\frac{1}{\sqrt{2}} | + \rangle + \frac{1}{\sqrt{2}} | - \rangle
$$

which indicates a state that is an equal mix of the $+$ and $-$ vector states. Typically this notation is used to indicate the [[spin]] up ($+$) and down ($-$) quantum states, where the above expression thus indicates a state that is an even _mix_ or **superposition** of these two basis states.

Mathematically, $+$ is typically defined as $(1,0)$ and $-$ as $(0,1)$ in the spin coordinate space -- these are the **pure** states, which in general you only get after you make a measurement, which, due to the [[contextual]] nature of quantum phenomena, rotates whatever starting state you might have into one of these pure **eigenstates**. 

In general, a _ket_ acts like a column vector:

$$
| w \rangle = \begin{pmatrix} w_1 \\ w_2 \end{pmatrix}
$$

while a _bra_ acts like a row vector:

$$
\langle z | = (z^*_1, z^*_2)
$$

which works out for the dot product notation:

$$
\langle z | w \rangle = (z^*_1, z^*_2) \begin{pmatrix} w_1 \\ w_2 \end{pmatrix} = z^*_1 w_1 + z^*_2 w_2
$$

As you can see, to get this dot product correct for the complex number case, the _bra_ version is subject to the complex conjugate.

You can transform a _ket_ into a _bra_, which requires taking the **conjugate transpose** of the _ket_ (also known as a **Hermitian transpose**):

$$
| z \rangle \rightarrow \langle z | = z^\dagger = \bar{z}^T = (z^*_1, z^*_2)
$$

This allows one to construct the **tensor product** or **outer product** of two _ket_ states, which is how you  construct the [[configuration space]] for two otherwise independent particle states:

$$
| a \rangle \otimes | b \rangle = | a \rangle \langle b | = \begin{pmatrix} a_1 \\ a_2 \end{pmatrix} (b_1, b_2) = \begin{pmatrix} a_1 b_1 & a_1 b_2 \\ a_2 b_1 & a_2 b_2 \end{pmatrix} 
$$

Here you can see that the number of elements in the state goes up as an exponential function of the number of times such an outer product is performed, i.e., $2^N$ for N 2D spin states.

## Measurement

A measurement amounts to **projecting** the quantum state onto a **measurement basis** state. This is written as:

$$
p(m) = \langle \psi | \Pi_m | \psi \rangle
$$

for a measurement _m_ parameterized by the **projectors** $\Pi_m$, which is an _N x N_ dimensional matrix that selects a specific subspace of the full _N_ dimensional tensor product space.

The key step in this computation is to compute the **density operator** $\rho$, which is the _outer product_ of the quantum state with itself:

$$
\rho = | \psi \rangle \langle \psi |
$$

This is effectively the full **correlation matrix** for every element of the state against every other element. The projector matrix then selects a subspace to project the correlations onto, and the final probability is computed as the trace (sum of the diagonal elements) of the resulting matrix. See below for the full equation, on the more general ensembles case. These diagonal elements of the projection matrix represent the eigenvalues of the projected matrix.

### Ensembles

A more general version of the overall framework includes the possibility of multiple different quantum states prepared with different probabilities $p_j$, in which case the aggregate density matrix can be computed as a weighted function of the probabilities, and then the projection matrix can be applied to that aggregate.

$$
p(m) = tr \left[ \Pi_m \left( \sum_j p_j \rho_j \right) \right] = tr \left[ \Pi_m \left( \sum_j p_j \langle \psi | \psi \rangle \right) \right] 
$$

The critical point here is that _the density matrix $\rho$ contains all the relevant information about all possible measurement outcomes._ Therefore, in most cases, that is the single most informative thing to compute. Of course, the actual projections are still needed to predict the outcomes of specific experiments.

## Time evolution

TODO:

