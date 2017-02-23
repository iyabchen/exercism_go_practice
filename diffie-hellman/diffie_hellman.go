package diffiehellman

// An application for math/big
import (
	"crypto/rand" // should use crypto/rand rather than math/rand to generate
	"math/big"
)

const testVersion = 1

// Alice and Bob use Diffie-Hellman key exchange to share secrets.  They
// start with prime numbers, pick private keys, generate and share public
// keys, and then generate a shared secret key.
// Step 1:   PrivateKey(p *big.Int) *big.Int
// Step 2:   PublicKey(private, p *big.Int, g int64) *big.Int
// Step 2.1: NewPair(p *big.Int, g int64) (private, public *big.Int)
// Step 3:   SecretKey(private1, public2, p *big.Int) *big.Int
//
// Private keys should be generated randomly.

// Pick a number > 1 and < p => p!=2
// Given a prime number p, return a number >1 && <p
// Should not generate same key no matter how many times invoked when possible
func PrivateKey(p *big.Int) *big.Int {
	// rand [0,p-2) + 2 => [2, p) i.e. (1,p)
	two := big.NewInt(2)
	x := big.NewInt(1)
	x = x.Sub(p, two)
	r, _ := rand.Int(rand.Reader, x) // rand.Reader is a cryptographically strong pseudo-random generator
	r.Add(r, two)
	return r

}

// Given private key a, and prime number p, g,
// calculate public key A = g**a mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	z := big.NewInt(1)
	return z.Exp(big.NewInt(g), private, p)
}

// Given public key B, private key a, and prime p
// calculate secretkey s = B**a mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	z := big.NewInt(1)
	return z.Exp(public2, private1, p)
}

// Given prime number p, g, return a pair of private and public keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return private, public
}
