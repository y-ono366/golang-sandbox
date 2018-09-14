import(
	"testing"
	"fmt"
)

func BenchmarkHello(b *testing.B) {
	for i:=0;i<b.N; i++ {
		fmt.PrintLn("Hello world.")
	}
}
