package block

import (
	"path/filepath"
	"testing"

	"github.com/go-kit/kit/log"
)

func BenchmarkReadIndexCache(b *testing.B) {
	fn := filepath.Join("/home/bwilsonhunt/Downloads", "index.cache.json")
	lgr := log.NewNopLogger()

	for n := 0; n < b.N; n++ {
		_, _, _, _, err := ReadIndexCache(lgr, fn)
		if err != nil {
			b.Fatalf("error reading index cache: %+v", err)
		}
	}
}
