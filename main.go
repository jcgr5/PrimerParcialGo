package main

func main() {
	//semaforos.Semaf("./InstanciasConcurrencia/01_5.input")
	//coleccion.Col("./InstanciasConcurrencia/01_5.input")
	//dataRace.Race("./InstanciasConcurrencia/01_5.input")
	//unHilo.St("./InstanciasConcurrencia/01_5.input")
	/*fmt.Println("Experimentación ST vs MT")
	// instancia := "./InstanciasConcurrencia/01_5.input"
	instancia := "./InstanciasConcurrencia/08_2500.input"
	var wg sync.WaitGroup
	//Versión SingleThreaded
	fmt.Println("-> Ejecutando versión de un solo hilo")
	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	var suma float32
	var mu sync.Mutex
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

	//Versión MultiThreaded - Data Race
	/*for i := 0; i < len(coleccionOperaciones); i++ {
		wg.Add(1)
		go func(op calculadora.Operacion) {
			op.Operar()
			suma += op.Resultado
			wg.Done()
		}(coleccionOperaciones[i])
	}*/
	//Versión MultiThreaded coleccion - Data Race
	/*suma := make([]float32, len(coleccionOperaciones))
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
	}*/

	//Versión MultiThreaded - Semáforos
	//.
	//.
	//.

}
