package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var (
		str string
		flag int
		temp int
		num [100]int
		tempf int32
		judge [100]int
		flag2 int
		new string
	)
	fmt.Scan(&str)
	lcc1:=[]rune(str)
	xs:=lcc1
	for i,v:=range lcc1{
		if v=='=' {
			flag=i
		}
	}
	var n = 0
	for i,v:= range xs{
		if (i==0&&(v>47&&v<58))||(i==flag+1&&(v>47&&v<58)) {
			tempf='+'
			temp1,_:=strconv.Atoi(string(v))
			temp+=temp1
		}else{
			if v>47&&v<58{
				temp1,_:=strconv.Atoi(string(v))
				temp=temp*10+temp1
				if i==len(xs)-1 {
					if tempf=='+' {
						num[n]=temp
					}else{
						num[n]=-temp
					}
				}
			}else {
				if i!=0 {
					if v>96&&v<123  {
						if xs[i-1]<48||xs[i-1]>57 {
							temp =1
						}
						if tempf=='=' {
							temp=1
						}
					}
					if tempf=='+' {
						num[n]=temp
					}else{
						if tempf=='-' {
							num[n]=-temp
						}
						if xs[i-1]=='=' {
							num[n]=temp
						}
					}
					if v>96&&v<123 {
						new= string(v)
						judge[n]=1
					}
					if v=='='{
						flag2=n
					}
					temp=0
					tempf=0
					n++
				}
				if v=='-'||v=='+'||v=='=' {
					tempf=v
				}
			}
		}
	}
	sum:=0
	sum2:=0
	for n,_:=range num{
		if n <= flag2&&judge[n]==0 {
			sum+= -num[n]
		}
		if n >flag2&&judge[n]==0 {
			sum+=num[n]
		}
		if  n < flag2&&judge[n]==1 {
			sum2+=num[n]
		}
		if  n > flag2&&judge[n]==1 {
			sum2+= -num[n]
		}
	}
	if float64(sum)/float64(sum2)<0.000000001&&float64(sum)/float64(sum2)>-0.00000000001 {
		fmt.Printf("%s=%.3f",new,0.000)
	}else {
		fmt.Printf("%s=%.3f",new,float64(sum)/float64(sum2))
	}
}
