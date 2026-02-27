package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rafi")
	}
}
func BenchmarkHelloWorldRafi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zimraan")
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot open this laptop")
	}

	result := HelloWorld("Rafi")
	require.Equal(t, "Hii Rafi", result, "Result must be 'Hii Rafi'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("zimraan")
	require.Equal(t, "Hii zimraan", result, "Result must be 'Hii zimraan'")
	fmt.Println("TestHelloWorldRequire is done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rafi")
	assert.Equal(t, "Hii Rafi", result, "Result must be 'Hii Rafi'")
	fmt.Println("TestHelloWorldAssert is done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Rafi")

	if result != "Hii Rafi" {
		// result
		t.Error("Result must be 'Hii Rafi' ")
	}
	fmt.Println("TestHelloWorld is done")
}

func TestHelloAmsterdam(t *testing.T) {
	result := HelloWorld("Zimraan")

	if result != "Hii Zimraan" {
		// result
		t.Fatal("Result must be 'Hii Zimraan'")
	}

	fmt.Println("TestHelloAmsterdam is done")
}
