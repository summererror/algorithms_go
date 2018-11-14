package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanf("%d %d", &a, &b) //obtain the input value
	if (a <= 0) || (b <= 0) {
		println(0) // one of a, b is zero or lower than zero, set gcd(a, b)=0
	} else {
		if a%b == 0 {
			println(b) // if a%b == 0, then gcd(a, b)=b
		} else {
			tmpB := b
			tmpA := a
			for tmpB != 0 { // calculate, until a%b==0, save in tmpB
				tmpA, tmpB = tmpB, tmpA%tmpB // calculate once, get the gcd(b , a%b)'s input value, i.e. b, a%b
				//println(tmpA, tmpB)
			}
			print(tmpA)
		}
	}

	//println("\n",a,b)

}
