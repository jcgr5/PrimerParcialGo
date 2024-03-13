package unHilo

import (
	"comparativosConcurrencia/gestorArchivos"
	"fmt"
	"time"
)

func St(inst string) {
	instancia := inst
	var suma float32
	fmt.Println("-> Ejecutando versión de un solo hilo")

	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	tiempoInicial := time.Now()
	for i := 0; i < len(coleccionOperaciones); i++ {
		coleccionOperaciones[i].Operar()
		suma += coleccionOperaciones[i].Resultado
	}
	fmt.Println("sumatoria: ", suma)
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())
}
