package coleccion

import (
	"comparativosConcurrencia/calculadora"
	"comparativosConcurrencia/gestorArchivos"
	"fmt"
	"sync"
	"time"
)

func Col(inst string) {
	fmt.Println("-> Ejecutando versión de multi hilos con coleccion reservada")
	instancia := inst
	var wg sync.WaitGroup
	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	suma := make([]float32, len(coleccionOperaciones))
	var sumas float32
	tiempoInicial := time.Now()
	for i := 0; i < len(coleccionOperaciones); i++ {
		wg.Add(1)
		go func(op calculadora.Operacion, ii int) {
			op.Operar()
			suma[ii] += op.Resultado
			wg.Done()
		}(coleccionOperaciones[i], i)
	}
	wg.Wait()
	for j, _ := range suma {
		sumas += suma[j]
	}
	fmt.Println("sumatoria: ", sumas)
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())
}
