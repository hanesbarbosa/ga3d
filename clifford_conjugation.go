package ga3d

import (
	"math/big"
)

// CliffordConjugation gives the involution of a multivector.
func CliffordConjugation(m Multivector) Multivector {
	// pf = positive factor
	pf := big.NewRat(1, 1)
	// nf = negative factor
	nf := big.NewRat(-1, 1)

	m.E0.Mul(pf, m.E0)
	m.E1.Mul(nf, m.E1)
	m.E2.Mul(nf, m.E2)
	m.E3.Mul(nf, m.E3)
	m.E12.Mul(nf, m.E12)
	m.E13.Mul(nf, m.E13)
	m.E23.Mul(nf, m.E23)
	m.E123.Mul(pf, m.E123)

	return m
}
