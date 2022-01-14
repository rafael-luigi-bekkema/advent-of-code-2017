package main

import "fmt"

type C3 struct {
	x, y, z int
}
type Particle struct {
	p, v, a C3
	del     bool
}

func (p *Particle) dist() int {
	return Abs(p.p.x) + Abs(p.p.y) + Abs(p.p.z)
}

func (p *Particle) step() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func (p1 *Particle) cmp(p2 *Particle) bool {
	return p1.p.x == p2.p.x && p1.p.y == p2.p.y && p1.p.z == p2.p.z
}

func day20parse(lines []string) []Particle {
	particles := make([]Particle, len(lines))
	for i, line := range lines {
		p := &particles[i]
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.p.x, &p.p.y, &p.p.z, &p.v.x, &p.v.y, &p.v.z, &p.a.x, &p.a.y, &p.a.z)
	}
	return particles
}

func day20a(lines []string) int {
	particles := day20parse(lines)
	var mindist, minpart int
	for n := 0; n < 500; n++ {
		for i := range particles {
			p := &particles[i]
			p.step()
			if d := p.dist(); i == 0 || d < mindist {
				mindist = d
				minpart = i
			}
		}
	}
	return minpart
}

func day20b(lines []string) int {
	particles := day20parse(lines)
	for n := 0; n < 200; n++ {
		var deletes []int
		for i := range particles {
			p := &particles[i]
			if p.del {
				continue
			}
			p.step()
			for j := i - 1; j >= 0; j-- {
				if p2 := &particles[j]; p2.cmp(p) {
					if p2.del {
						continue
					}
					deletes = append(deletes, i, j)
				}
			}
		}
		for _, d := range deletes {
			particles[d].del = true
		}
	}
	var count int
	for _, p := range particles {
		if !p.del {
			count++
		}
	}
	return count
}
