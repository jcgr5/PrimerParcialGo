# PrimerParcialGo

Este proyecto contiene 4 procesos distintos para procesar datos, para llamar estas funciones realizar lo siguiente:

1-ubicarse en el archivo main.go  
  
2-escribir la siguiente linea de codigo (eliminando los // del inicio de cada linea) dependiendo del proceso que desee utilizar:  
  
-para utilizar metodo de multi hilo con semaforos: //semaforos.Semaf("./InstanciasConcurrencia/01_5.input")  
-para utilizar metodo de multi hilo con coleccion reservada: //coleccion.Col("./InstanciasConcurrencia/01_5.input")  
-para utilizar metodo de multi hilo con una variable: //dataRace.Race("./InstanciasConcurrencia/01_5.input")  
-para utilizar metodo de un solo hilo: //unHilo.St("./InstanciasConcurrencia/01_5.input")  

3-Para cambiar las instancias de concurrencia, modifique 01_5.input por una de las siguientes opciones:  
  
01_5.input  
02_20.input  
03_40.input  
04_100.input  
05_200.input  
06_500.input  
07_1000.input  
08_2500.input  
09_10000.input  
