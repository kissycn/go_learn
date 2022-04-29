package chapter12_benchmark

import (
	"fmt"
	"strings"
	"testing"
)

var s1 = []string{
	"PikePikePikePikePikePikePike",
	"KobeKobeKobeKobeKobeKobeKobeKobe",
	"JamesJamesJamesJamesJamesJamesJames",
}

func contactStringByOperator(s []string) string {
	var res string
	for _, v := range s {
		res += v
	}
	return res
}

func contactBySprintf(s []string) string {
	var res string
	for _, v := range s {
		res = fmt.Sprintf("%s%s", s, v)
	}
	return res
}

func contactByJoin(s []string) []string {
	var res []string

	for _, v := range s {
		strings.Join(res, v)
	}
	return res
}

func BenchmarkContactStrByOp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		contactStringByOperator(s1)
	}
}

func BenchmarkContactStrBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		contactBySprintf(s1)
	}
}

func BenchmarkContactStrByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		contactByJoin(s1)
	}
}
