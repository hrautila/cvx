
// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/cvx package. It is free software, distributed
// under the terms of GNU Lesser General Public License Version 3, or any later
// version. See the COPYING tile included in this archive.

/*
Package cvx is a package for solving convex optimization problems.

It is a straightforward translation of parts of the CVXOPT python package for convex
optimization. Spefically it provides interfaces for solving linear and quadratic cone
programs and convex programs with non-linear objectives.

Package cvx depends on column order matrix implementation and access to BLAS and
LAPACK linear algebra libraries. 

Solvers

Following solvers are provided:

   ConeLp		Linear Cone programs
   ConeQp		Quadratic Cone programs
   Lp		Linear programs
   Qp		Quadratic programs
   Socp		Second-Order Cone programs
   Sdp		Semidefinite programs
   Cpl		Convex programs with linear objectives
   Cp		Convex programs with non-linear objectives
   Gp		Geometric programs

Main solvers for Cone Programs are ConeLp and ConeQp which provide interfaces for advanced
usage with custom solvers.

The non-linear convex optimization solvers are Cp and Cpl. They provide interfaces for advanced
usage with solvers.

The Lp, Qp, Socp, Sdp and Gp solvers provide only the standard matrix interface without any
customization options.


Output

All solvers return 

Extended Usage

Package support three levels of advanced usage to allow exploiting the problem structure.

On the first level a custom KKT solver can provide to solve the KKT equations.

W is a scaling matrix, a block diagonal mapping

    W*u = ( W0*u_0, ..., W_{N+M}*u_{N+M} )

defined as follows.

For the 'l' block (W_0):

    W_0 = diag(d),

with d a positive vector of length ml.

For the 'q' blocks (W_{k+1}, k = 0, ..., N-1):

    W_{k+1} = beta_k * ( 2 * v_k * v_k' - J )

where beta_k is a positive scalar, v_k is a vector in R^mq[k]
with v_k[0] > 0 and v_k'*J*v_k = 1, and J = [1, 0; 0, -I].

For the 's' blocks (W_{k+N}, k = 0, ..., M-1):

    W_k * u = vec(r_k' * mat(u) * r_k)

where r_k is a nonsingular matrix of order ms[k], and mat(x) is
the inverse of the vec operation.

The optional argument kktsolver is a function that will be
called as g = kktsolver(W).

 W is a FloatMatrixSet that contains the parameters of the scaling:

 W.At("d")    is a positive  matrix of size (ml,1). (array of size one)
 W.At("di")   is a positive  matrix matrix with the elementwise inverse of W.At("d").
 W.At("beta") is a matrix [ beta_0, ..., beta_{N-1} ]
 W.At("v")    is an array [ v_0, ..., v_{N-1} ] of float matrices.
 W.At("r")    is an array [ r_0, ..., r_{M-1} ] of matrices
 W.At("rti")  is an array [ rti_0, ..., rti_{M-1} ],
              with rti_k the inverse of the transpose of r_k of W.At("rti")

The call g = kktsolver(W) should return a function g that solves 
the KKT system by g(x, y, z).  On entry, x, y, z contain the 
righthand side bx, by, bz.  On exit, they contain the solution,
with uz scaled, the argument z contains W*uz.  In other words, 
on exit x, y, z are the solution of

   [ P   A'  G'*W^{-1} ] [ ux ]   [ bx ]
   [ A   0   0         ] [ uy ] = [ by ].
   [ G   0   -W'       ] [ uz ]   [ bz ]


Cvxopt User's Guide

For more detailed discussion on using solvers see

   http://abel.ee.ucla.edu/cvxopt/userguide/index.html

 */
package cvx
