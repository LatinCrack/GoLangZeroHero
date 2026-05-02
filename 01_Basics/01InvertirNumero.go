package main

import "fmt"

func main() {
	// OBJETIVO: Invertir un número entero manejando valores positivos, negativos y validación de entrada.

	var n int
	invertido := 0

	fmt.Print("Ingresa un número entero: ")

	// MEJORA 1: Manejo de errores. Capturamos si el usuario ingresa texto (como "popop").
	// Scanln devuelve un error si la entrada no coincide con el tipo 'int'.
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error: ¡Debes ingresar un número entero válido!")
		return
	}

	original := n

	// MEJORA 2: Condición de bucle. Cambiamos 'n > 0' por 'n != 0'.
	// Esto permite que el algoritmo procese números negativos (como -952).
	for n != 0 {

		// PASO A: Extraer el último dígito.
		// En Go, el residuo (%) de un negativo sigue siendo negativo (-952 % 10 = -2).
		digito := n % 10

		// PASO B: Acumular. El signo se mantiene en la operación aritmética.
		invertido = (invertido * 10) + digito

		// PASO C: Eliminar el último dígito usando división entera.
		n = n / 10
	}

	// Resultados finales
	fmt.Printf("El numero original es: %d\n", original)
	fmt.Printf("El numero invertido es: %d\n", invertido)
}
