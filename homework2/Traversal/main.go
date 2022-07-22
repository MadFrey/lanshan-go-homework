/**
 * @Author: lrc
 * @Date: 2022/7/22-11:42
 * @Desc:
 **/

package main

import "fmt"

var (
	arr   [400][400]int
	P     [400]int
	D     [400][400]int
	final [400][400]int
)

func main() {
	n := 0
	m := 0
	x := 0
	y := 0
	dx := [8]int{-1, -2, -2, -1, 1, 2, 2, 1}
	dy := [8]int{2, 1, -1, -2, 2, 1, -1, -2}
	fmt.Scan(&n, &m, &x, &y)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			D[i][j] = 9999999
		}
	}

	D[x-1][y-1] = 0
	final[x-1][y-1] = 1
	for x = 0; x < n; x++ {
		for y = 0; y < m; y++ {
			for j := 0; j < 8; j++ {
				if (x+dx[j] > -1 && x+dx[j] < n) && (y+dy[j] > -1 && y+dy[j] < m) {
					if D[x][y]+1 < D[x+dx[j]][y+dy[j]] {
						D[x+dx[j]][y+dy[j]] = D[x][y] + 1
						final[x+dx[j]][y+dy[j]] = 1
					}
				}
			}
			for a := 0; a < n; a++ {
				for b := 0; b < m; b++ {
					if final[a][b] == 1 {
						for j := 0; j < 8; j++ {
							if (a+dx[j] > -1 && a+dx[j] < n) && (b+dy[j] > -1 && b+dy[j] < m) {
								if D[a][b]+1 < D[a+dx[j]][b+dy[j]] {
									D[a+dx[j]][b+dy[j]] = D[a][b] + 1
									final[a+dx[j]][b+dy[j]] = 1
								}
							}
						}
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if D[i][j] == 9999999 {
				D[i][j] = -1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j != m-1 {
				fmt.Printf("%-5d", D[i][j])
			} else {
				fmt.Printf("%-5d\n", D[i][j])
			}
		}
	}
}
