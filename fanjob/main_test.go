package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	const strCount = 100
	items := genStrings(strCount)

	start := time.Now()
	r1 := doJobs(items)
	fmt.Printf("%v\n", time.Now().Sub(start))

	start = time.Now()
	r2 := doJobsFanout(items)
	fmt.Printf("%v\n", time.Now().Sub(start))

	assert.ElementsMatch(t, r1, r2)
}
