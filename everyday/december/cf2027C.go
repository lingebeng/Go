package december

import (
	. "fmt"
	"io"
)

func Cf2027C(in io.Reader, out io.Writer) {
	var T, n, v int
	Fscan(in, &T)
	for ; T > 0; T-- {
		Fscan(in, &n, &v)
		g := map[int][]int{}
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			v += i
			g[v] = append(g[v], v+i)
		}
		ans := n
		vis := map[int]bool{}
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			ans = max(ans, v)
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		dfs(n)
		Fprintln(out, ans)
	}
}
