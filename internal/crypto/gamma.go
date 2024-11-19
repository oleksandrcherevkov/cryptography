package crypto

import (
	"math"
	"strconv"
)

type Gamma struct {
	random *linearCongruentialGenerator
}

const (
	m          = math.MaxUint
	windowSize = strconv.IntSize
)

func NewGamma(a, c, b uint) *Gamma {
	random := newLinear(a, c, b)
	return &Gamma{
		random: random,
	}
}

func (g *Gamma) Encrypt(s *Stream) error {
	buffer := make([]byte, windowSize/8)
	transform := func(bytes []byte) {
		gamma := g.random.Next()
		applyGamma(bytes, gamma)
	}
	s.Transform(transform, buffer)
	return nil
}

func (g *Gamma) Decrypt(s *Stream) error {
	return g.Encrypt(s)
}

func applyGamma(bytes []byte, gamma uint) {
	for i, block := range bytes {
		gammaByte := byte(gamma >> (i * 8))
		bytes[i] = block ^ gammaByte
	}
}

type linearCongruentialGenerator struct {
	next uint
	a    uint
	c    uint
}

func newLinear(a, c, next uint) *linearCongruentialGenerator {
	return &linearCongruentialGenerator{a, c, next}
}

func (g *linearCongruentialGenerator) Next() uint {
	g.next = (g.a*g.next + g.c) % m
	return g.next
}
