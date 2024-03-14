package semaforos

import (
	"fmt"
	"github.com/jcgr5/PrimerParcialGo/calculadora"
	"github.com/jcgr5/PrimerParcialGo/gestorArchivos"
	"sync"
	"time"
)

func Semaf(inst string) {

	fmt.Println("-> Ejecutando versión de multi hilo con semaforo")
	instancia := inst
	var wg sync.WaitGroup
	var suma float32
	var mu sync.Mutex
	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)

	tiempoInicial := time.Now()
	for i := 0; i < len(coleccionOperaciones); i++ {
		wg.Add(1)
		go func(op calculadora.Operacion, m *sync.Mutex) {
			op.Operar()
			mu.Lock()
			suma += op.Resultado
			mu.Unlock()
			wg.Done()
		}(coleccionOperaciones[i], &mu)
	}
	wg.Wait()
	fmt.Println("sumatoria: ", suma)
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())
}
