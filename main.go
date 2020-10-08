package main

import "fmt"

type barbell interface {
	getWeight() int
}
type bigBarbell struct{

}
func (b *bigBarbell) getWeight() int{
	return 20
}
type kg5 struct{
	barbell barbell
}
func (k *kg5) getWeight() int{
	barbellWeight :=k.barbell.getWeight()
	return barbellWeight + 10
}
type kg10 struct{
	barbell barbell
}
func (k *kg10) getWeight() int{
	barbellWeight :=k.barbell.getWeight()
	return barbellWeight + 20
}
type kg15 struct{
	barbell barbell
}
func (k *kg15) getWeight() int{
	barbellWeight :=k.barbell.getWeight()
	return barbellWeight + 30
}
func main(){
	barbell:=&bigBarbell{}
	barbellKg15:=&kg15{
		barbell:barbell,
	}
	barbellKg10:=&kg10{
		barbell:barbellKg15,

	}
	barbellKg5:=&kg5{
		barbell:barbellKg10,

	}
	fmt.Printf("Weight is %d\n", barbellKg5.getWeight())
}
