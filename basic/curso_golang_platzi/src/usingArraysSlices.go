package main

import "fmt"

func mainIX() {
	//Arrays : Este elemento es INMUTABLE
	var array [4]int
	fmt.Println(array)

	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	//Len, nos dice la cantidad de elementos en el Array
	fmt.Println("Numero de elementos:", len(array))

	//Cap, nos indica la capacidad maxima del array
	fmt.Println("Capacidad máxima:", cap(array))

	//Slice: Este elemento es MUTABLE - Es similar a un array, pero se le pueden incluir mas elementos
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice: %v Numero de elementos: %d Capacidad máxima: %d\n", slice, len(slice), cap(slice))

	//¿Cómo se interactua con cada uno de los elementos de un ARRAY, SLICE o LIST?
	//Métodos en el slice
	//NOTA: La siguiente notación funciona así: [3:4], donde 3 es INCLUSIVO y 4 es EXCLUSIVO
	fmt.Println(slice[0])   //Imprime el elemento que se encuentra en el indice 0
	fmt.Println(slice[:3])  //Imprime los elementos hasta el indice 3
	fmt.Println(slice[2:4]) //Imprime los elementos entre el indice 2 y el indice 4
	fmt.Println(slice[4:])  //Imprime desde el elemento 4 hasta el final

	//Append
	slice = append(slice, 7)
	fmt.Printf("Slice: %v Numero de elementos: %d Capacidad máxima: %d\n", slice, len(slice), cap(slice))

	//Append, para agregar otra lista al slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Printf("NewSlice: %v Numero de elementos: %d Capacidad máxima: %d\n", newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("Slice: %v Numero de elementos: %d Capacidad máxima: %d\n", slice, len(slice), cap(slice))

	slice = append(slice, 11, 12, 13, 14)
	fmt.Printf("Slice: %v Numero de elementos: %d Capacidad máxima: %d\n", slice, len(slice), cap(slice))
}
