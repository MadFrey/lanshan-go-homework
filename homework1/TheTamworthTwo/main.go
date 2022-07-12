package main

import (
	"fmt"
)

type Fame struct {
	x int
	y int
	d int
}

type cow struct {
	x int
	y int
	d int
}
// 0,1,2,3 北东南西
type direction struct {
	tx int
	ty int
}

var oMap [10][10]string

func main()  {
	var (
		str string
		C cow
		F Fame
		dir [4]direction
		count int
		mp [160005]int
	)
	dir[0].tx=-1
	dir[0].ty=0
	dir[1].tx=0
	dir[1].ty=1
	dir[2].tx=1
	dir[2].ty=0
	dir[3].tx=0
	dir[3].ty=-1
	for i := 0; i <10; i++ {
		fmt.Scan(&str)
		  temp := []rune(str)
		for j,v:=range temp{
			a:=string(v)
			oMap[i][j]=a
			if a=="F"{
				F.x=i
				F.y=j
				F.d=0
			}
			if a=="C"{
				C.x=i
				C.y=j
				C.d=0
			}
		}
	}
	count=0
	for  {
		if C.x==F.x&&C.y==F.y {
			break
		}
		nt:=F.x+F.y*10+C.x*100+C.y*1000+F.d*10000+C.d*40000
		if mp[nt]==1 {
			count=0
			break
		}
		mp[nt]=1
		if F.x+dir[F.d].tx>9||F.x+dir[F.d].tx<0||F.y+dir[F.d].ty>9||F.y+dir[F.d].ty<0||oMap[F.x+dir[F.d].tx][F.y+dir[F.d].ty]=="*" {
			F.d=(F.d+1)%4
		}else{
				F.x+=dir[F.d].tx
				F.y+=dir[F.d].ty
		}
		if C.x + dir[C.d].tx>9||C.x + dir[C.d].tx<0||C.y+dir[C.d].ty>9||C.y+dir[C.d].ty<0||oMap[C.x+dir[C.d].tx][C.y+dir[C.d].ty]=="*"  {
			C.d=(C.d+1)%4
		}else {
			C.x+=dir[C.d].tx
			C.y+=dir[C.d].ty
		}
		count++
	}
	fmt.Println(count)
}
