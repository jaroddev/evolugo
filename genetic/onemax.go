package genetic

import (
	. "evolugo/chromosomes"
	"math/rand"
)

type OneMax struct {
	InitPopLength int
	AlleleLength  int

	MaxCycle      int
	NotUpdatedFor int
}

func (o *OneMax) Init() Population {
	pop := make(Population, o.InitPopLength)

	for index := range pop {
		pop[index].Alleles = make([]bool, o.AlleleLength)

		for locus := range pop[index].Alleles {
			if rand.Intn(2) == 1 {
				pop[index].Alleles[locus] = true
			} else {
				pop[index].Alleles[locus] = false
			}
		}
	}

	return pop
}

func (o *OneMax) Fit(c *Chromosome) {
	c.Fitness = 0
	for _, allele := range c.Alleles {
		if allele {
			c.Fitness++
		}
	}
}

func (o *OneMax) Continue(algorithm *GA) bool {
	// Return true if the algorithm should continue
	// if return false then the algorithm stop

	// Stop if cycle number is higher than 80
	return algorithm.Cycle < o.MaxCycle &&
		// or if there were no update for at least 20 cycles
		// algorithm.Cycle-algorithm.LastUpdate < o.NotUpdatedFor &&
		// or if
		algorithm.Best.Fitness < float64(len(algorithm.Best.Alleles))
}