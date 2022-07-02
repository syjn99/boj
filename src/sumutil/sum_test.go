package sumutil
import "testing"
func TestSum(t *testing.T) {
    cases := []struct {
        start, end, result int
    }{
            { 5, 6, 11 },
            { 1, 9, 45 }, // (1+9)+(2+8)+(3+7)+(4+6)+5=45
            { 0, 9, 45 }, 
            { 1, 3, 6 },
            { 10, 19, 145 }, // 10+(10+1)+(10+2)+...+(10+9)=
                         // 10*10+sum(1,9)=145
    }
    for _, c := range cases {
        if got := Sum(c.start, c.end); got != c.result {
            t.Errorf("Sum(%v,%v) got %v, want %v", c.start, c.end,got, c.result)
        }
    }
}
func BenchmarkSum(b *testing.B) {
    for i:=0; i<b.N; i++ {
        Sum(0,i)
    }
}