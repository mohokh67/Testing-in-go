package unique_id

import (
	"testing"
)

func BenchmarkUUIDGenerator(b *testing.B) {
	// b.Log("We are testing UUID generator by Google")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		UUIDGenerator()
	}

	b.StopTimer()

}

func BenchmarkKSUIDGenerator(b *testing.B) {
	// b.Log("We are testing KSUID generator")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		KSUIDGenerator()
	}

	b.StopTimer()
}
