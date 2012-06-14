
// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/linalg package. It is free software, distributed
// under the terms of GNU Lesser General Public License Version 3, or any later
// version. See the COPYING tile included in this archive.

package lapack

import (
	"linalg"
	"matrix"
	"errors"
)

/*
 Eigenvalue decomposition of a real symmetric matrix
 (divide-and-conquer driver).

 Syevd(A, W, jobz=PJboNo, uplo=PLower, n=A.Rows, 
 ldA = max(1,A.Rows), offsetA=0, offsetW=0)

 PURPOSE

 Returns  eigenvalues/vectors of a real symmetric nxn matrix A.
 On exit, W contains the eigenvalues in ascending order.
 If jobz is PJobV, the (normalized) eigenvectors are also computed
 and returned in A.  If jobz is PJobNo, only the eigenvalues are
 computed, and the content of A is destroyed.

 ARGUMENTS
  A         float matrix
  W         float matrix of length at least n.  On exit, contains
            the computed eigenvalues in ascending order.

 OPTIONS
  jobz      PJobNo or PJobV
  uplo      PLower or PUpper
  n         integer.  If negative, the default value is used.
  ldA       nonnegative integer.  ldA >= max(1,n).  If zero, the
            default value is used.
  offsetA   nonnegative integer
  offsetB   nonnegative integer;
 */
func Syevd(A, W matrix.Matrix, opts ...linalg.Opt) error {
	pars, err := linalg.GetParameters(opts...)
	if err != nil {
		return err
	}
	ind := linalg.GetIndexOpts(opts...)
	if ind.N < 0 {
		ind.N = A.Rows()
		if ind.N != A.Cols() {
			return errors.New("A not square")
		}
	}
	if ind.N == 0 {
		return nil
	}
	if ind.LDa == 0 {
		ind.LDa = max(1, A.Rows())
	}
	if ind.LDa < max(1, ind.N) {
		return errors.New("lda")
	}
	if ind.OffsetA < 0 {
		return errors.New("offsetA")
	}
	sizeA := A.NumElements()
	if sizeA < ind.OffsetA+(ind.N-1)*ind.LDa+ind.N {
		return errors.New("sizeA")
	}
	// B is the W matrix
	if ind.OffsetW < 0 {
		return errors.New("offsetW")
	}
	sizeW := W.NumElements()
	if sizeW < ind.OffsetW + ind.N {
		return errors.New("sizeW")
	}

	var info int
	switch A.(type) {
	case *matrix.FloatMatrix:
		jobz := linalg.ParamString(pars.Jobz)
		uplo := linalg.ParamString(pars.Uplo)
		Aa := A.FloatArray()
		Wa := W.FloatArray()
		info = dsyevd(jobz, uplo, ind.N, Aa[ind.OffsetA:], ind.LDa, Wa[ind.OffsetW:])
	case *matrix.ComplexMatrix:
		return errors.New("Not a complex function")
	}
	if info != 0 {
		return errors.New("Syevd call error")
	}
	return nil
}


// Local Variables:
// tab-width: 4
// End:
