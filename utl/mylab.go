// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utl

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

// IntFill fills a slice of integers
func IntFill(s []int, val int) {
	for i := 0; i < len(s); i++ {
		s[i] = val
	}
}

// IntVals allocates a slice of integers with size==n, filled with val
func IntVals(n int, val int) (s []int) {
	s = make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = val
	}
	return
}

// StrVals allocates a slice of strings with size==n, filled with val
func StrVals(n int, val string) (s []string) {
	s = make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = val
	}
	return
}

// IntsAlloc allocates a matrix of integers
func IntsAlloc(m, n int) (mat [][]int) {
	mat = make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	return
}

// IntRange generates a slice of integers from 0 to n-1
func IntRange(n int) (res []int) {
	if n <= 0 {
		return
	}
	res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return
}

// IntRange2 generates slice of integers from start to stop (but not stop)
func IntRange2(start, stop int) []int {
	return IntRange3(start, stop, 1)
}

// IntRange3 generates a slice of integers from start to stop (but not stop), afer each 'step'
func IntRange3(start, stop, step int) (res []int) {
	switch {
	case stop == start:
		return
	case stop > start:
		n := (stop - start) / step
		res = make([]int, n)
		for i, v := 0, start; v < stop; i, v = i+1, v+step {
			res[i] = v
		}
	case stop < start:
		if step > 0 {
			return
		}
		n := (stop - start) / step
		res = make([]int, n)
		for i, v := 0, start; v > stop; i, v = i+1, v+step {
			res[i] = v
		}
	}
	return
}

// IntAddScalar adds a scalar to all values in a slice of integers
func IntAddScalar(a []int, s int) (res []int) {
	res = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] + s
	}
	return
}

// IntUnique returns a unique and sorted slice of integers
func IntUnique(slices ...[]int) (res []int) {
	if len(slices) == 0 {
		return
	}
	nn := 0
	for i := 0; i < len(slices); i++ {
		nn += len(slices[i])
	}
	res = make([]int, 0, nn)
	for i := 0; i < len(slices); i++ {
		a := make([]int, len(slices[i]))
		copy(a, slices[i])
		sort.Ints(a)
		for j := 0; j < len(a); j++ {
			idx := sort.SearchInts(res, a[j])
			if idx < len(res) && res[idx] == a[j] {
				continue // found
			} else {
				if idx == len(res) { // append
					res = append(res, a[j])
				} else { // insert
					res = append(res[:idx], append([]int{a[j]}, res[idx:]...)...)
				}
			}
		}
	}
	return
}

// IntPy returns a Python string representing a slice of integers
func IntPy(a []int) (res string) {
	res = "["
	for i := 0; i < len(a); i++ {
		res += strconv.Itoa(a[i])
		if i < len(a)-1 {
			res += ", "
		}
	}
	res += "]"
	return
}

// DblOnes generates a slice of double precision '1s'
func DblOnes(n int) (res []float64) {
	res = make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = 1.0
	}
	return
}

// DblVals generates a slice of double precision values
func DblVals(n int, v float64) (res []float64) {
	res = make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = v
	}
	return
}

// IntMinMax returns the maximum and minimum elements in v
//  NOTE: this is not efficient and should be used for small slices only
func IntMinMax(v []int) (mi, ma int) {
	mi, ma = v[0], v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < mi {
			mi = v[i]
		}
		if v[i] > ma {
			ma = v[i]
		}
	}
	return
}

// DblMinMax returns the maximum and minimum elements in v
//  NOTE: this is not efficient and should be used for small slices only
func DblMinMax(v []float64) (mi, ma float64) {
	mi, ma = v[0], v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < mi {
			mi = v[i]
		}
		if v[i] > ma {
			ma = v[i]
		}
	}
	return
}

// DblSum sums all items in v
//  NOTE: this is not efficient and should be used for small slices only
func DblSum(v []float64) (sum float64) {
	for i := 0; i < len(v); i++ {
		sum += v[i]
	}
	return
}

// DurSum sums all seconds in v
//  NOTE: this is not efficient and should be used for small slices only
func DurSum(v []time.Duration) (seconds float64) {
	for _, t := range v {
		seconds += t.Seconds()
	}
	return
}

// StrIndexSmall finds the index of an item in a slice of strings
//  NOTE: this function is not efficient and should be used with small slices only; say smaller than 20
func StrIndexSmall(a []string, val string) int {
	for idx, str := range a {
		if str == val {
			return idx
		}
	}
	return -1 // not found
}

// IntIndexSmall finds the index of an item in a slice of ints
//  NOTE: this function is not efficient and should be used with small slices only; say smaller than 20
func IntIndexSmall(a []int, val int) int {
	for idx, item := range a {
		if item == val {
			return idx
		}
	}
	return -1 // not found
}

// IntFilter filters out components in slice
//  NOTE: this function is not efficient and should be used with small slices only
func IntFilter(a []int, out func(idx int) bool) (res []int) {
	for i := 0; i < len(a); i++ {
		if out(i) {
			continue
		}
		res = append(res, a[i])
	}
	return
}

// IntNegOut filters out negative components in slice
//  NOTE: this function is not efficient and should be used with small slices only
func IntNegOut(a []int) (res []int) {
	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			continue
		}
		res = append(res, a[i])
	}
	return
}

// LinSpace returns evenly spaced numbers over a specified closed interval.
func LinSpace(start, stop float64, num int) (res []float64) {
	if num <= 0 {
		return []float64{}
	}
	if num == 1 {
		return []float64{start}
	}
	step := (stop - start) / float64(num-1)
	res = make([]float64, num)
	for i := 0; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	res[num-1] = stop
	return
}

// LinSpaceOpen returns evenly spaced numbers over a specified open interval.
func LinSpaceOpen(start, stop float64, num int) (res []float64) {
	if num <= 0 {
		return []float64{}
	}
	step := (stop - start) / float64(num)
	res = make([]float64, num)
	for i := 0; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	return
}

// LocGlobMaps returns two maps of integers where g2l contains the indices corresponding
// to each item in allkeys, except that excluded keys are indicated with -1.
// l2g contains only the indices of of keys not excluded
//  Example:
//                         0    1    2    3    4    5
//   allkeys  := []string{"a", "b", "A", "B", "α", "β"}
//   excluded := []string{          "A",           "β"}
//   g2l       = []int   { 0,   1,  -1,   2,   3,  -1 }
//   l2g       = []int   { 0,   1,        3,   4      }
func LocGlobMaps(allkeys, excludedkeys []string) (g2l, l2g []int) {
	N := len(allkeys)           // number of global (all) keys
	n := N - len(excludedkeys)  // number of local keys
	g2l = make([]int, N)        // global => local
	l2g = make([]int, n)        // local => global
	k2I := make(map[string]int) // key => global index I (-1 means excluded)
	for I, k := range allkeys {
		k2I[k] = I
		g2l[I] = -1
	}
	for _, k := range excludedkeys {
		_, ok := k2I[k]
		if !ok {
			Panic("mylab.go: LocGlobMaps: cannot find excluded key '%s' within allkeys=%v (excludedkeys=%v)", k, allkeys, excludedkeys)
		}
		k2I[k] = -1 // exclude key k
	}
	i := 0
	for I, k := range allkeys {
		if k2I[k] >= 0 { // not excluded
			g2l[I] = i
			l2g[i] = I
			i += 1
		}
	}
	return
}

// LocGlobMapsM is the same as LocGlobMaps but using a map for excludedkeys
func LocGlobMapsM(allkeys []string, excludedkeysM map[string]bool) (g2l, l2g []int) {
	nexcl := 0 // number of excluded keys
	for _, val := range excludedkeysM {
		if val {
			nexcl += 1
		}
	}
	excludedkeys, k := make([]string, nexcl), 0
	for key, val := range excludedkeysM {
		if val {
			excludedkeys[k] = key
			k += 1
		}
	}
	return LocGlobMaps(allkeys, excludedkeys)
}

// Dbl2Str converts a slice of doubles (float64) to a slice of strings
func Dbl2Str(v []float64, format string) (s []string) {
	s = make([]string, len(v))
	for i := 0; i < len(v); i++ {
		s[i] = Sf(format, v[i])
	}
	return
}

// Str2Dbl converts a slice of strings to a slice of doubles (float64)
func Str2Dbl(s []string) (v []float64) {
	v = make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		v[i] = Atof(s[i])
	}
	return
}

// DblSplit splits a string into floats
func DblSplit(s string) (r []float64) {
	ss := strings.Fields(s)
	r = make([]float64, len(ss))
	for i, v := range ss {
		r[i] = Atof(v)
	}
	return
}

// Atan2p implements a positive version of atan2, in such a way that: 0 ≤ α ≤ 2π
func Atan2p(y, x float64) (αrad float64) {
	αrad = math.Atan2(y, x)
	if αrad < 0.0 {
		αrad += 2.0 * math.Pi
	}
	return
}

// Atan2pDeg implements a positive version of atan2, in such a way that: 0 ≤ α ≤ 360
func Atan2pDeg(y, x float64) (αdeg float64) {
	αdeg = math.Atan2(y, x) * 180.0 / math.Pi
	if αdeg < 0.0 {
		αdeg += 360.0
	}
	return
}

// Ramp function => MacAulay brackets
func Ramp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	}
	return x
}

// Heaviside step function (== derivative of Ramp(x))
func Heav(x float64) float64 {
	if x < 0.0 {
		return 0.0
	}
	if x > 0.0 {
		return 1.0
	}
	return 0.5
}

// Sign function
func Sign(x float64) float64 {
	if x < 0.0 {
		return -1.0
	}
	if x > 0.0 {
		return 1.0
	}
	return 0.0
}

// Sramp implements a smooth ramp function. Ramp
func Sramp(x, β float64) float64 {
	if -β*x > 500.0 {
		return 0.0
	}
	return x + math.Log(1.0+math.Exp(-β*x))/β
}

// SrampD1 returns the first derivative of Sramp
func SrampD1(x, β float64) float64 {
	if -β*x > 500.0 {
		return 0.0
	}
	return 1.0 / (1.0 + math.Exp(-β*x))
}

// SrampD2 returns the second derivative of Sramp
func SrampD2(x, β float64) float64 {
	if β*x > 500.0 {
		return 0.0
	}
	return β * math.Exp(β*x) / math.Pow(math.Exp(β*x)+1.0, 2.0)
}
