package nodejs

import (
	"os"
	"reflect"
	"testing"
)

func TestPackageLockJsonParsing(t *testing.T) {
	data, err := os.ReadFile("../../testdata/package-lock.json")
	if err != nil {
		panic(err)
	}
	got, _ := NpmDeps(data)
	want := map[string] string {
		"js-tokens": "4.0.0",
		"loose-envify": "1.4.0",
		"object-assign": "4.1.1",
		"react": "17.0.2",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
