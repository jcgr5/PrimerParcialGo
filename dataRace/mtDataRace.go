package dataRace

import (
	"comparativosConcurrencia/calculadora"
	"comparativosConcurrencia/gestorArchivos"
	"fmt"
	"sync"
	"time"
)

func Race(inst string) {
	instancia := inst
	var wg sync.WaitGroup
	var suma float32
	fmt.Println("-> Ejecutando versión multi hilo con data race")
	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	tiempoInicial := time.Now()
	for i := 0; i < len(coleccionOperaciones); i++ {
		wg.Add(1)
		go func(op calculadora.Operacion) {
			op.Operar()
			suma += op.Resultado
			wg.Done()
		}(coleccionOperaciones[i])
	}
	wg.Wait()
	fmt.Println("sumatoria: ", suma)
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())
}
