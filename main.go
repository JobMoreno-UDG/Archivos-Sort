package main

import ("fmt"
		"os"
		"sort"
	)

type Nombre []string

func (a Nombre) Len() int           { return len(a) }
func (a Nombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nombre) Less(i, j int) bool { return a[i] < a[j] }

func main(){
	var cadenas []string
	var cont int
	fmt.Println("Cuantos Strings va a Introducir ?")
	fmt.Scan(&cont)
	x := 0
	var cad string
	for x < cont{
		fmt.Scan(&cad)
		cadenas = append(cadenas,cad)
		x += 1
	}
	file,err:=os.Create("asecendente.txt")
	sort.Sort(Nombre(cadenas))
	fmt.Println(cadenas)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _,v := range cadenas{
		file.WriteString(v+"\n")
	}
	defer file.Close()
	
	file2,err:=os.Create("descendente.txt")
	sort.Sort(sort.Reverse(Nombre(cadenas)))
	fmt.Println(cadenas)
	for _,v := range cadenas{
		file2.WriteString(v+"\n")
	}
	defer file2.Close()

}