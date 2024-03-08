package precompiles

import (
	bls "github.com/drand/kyber-bls12381"
	"github.com/sirupsen/logrus"
)

type ArbPairing struct {
	Address addr // 0x23
}


func (con ArbPairing) Pairing(c ctx, evm mech, a [96]byte, b [48]byte) ([]byte, error) {
	logrus.Info("----------------------------------------", a,b)
	s := bls.NewBLS12381Suite()
	ap := s.G2().Point()
	ap.UnmarshalBinary(a[:])
	bp := s.G1().Point()
	bp.UnmarshalBinary(b[:])
	p := s.Pair(bp, ap)
	return p.MarshalBinary()
}
