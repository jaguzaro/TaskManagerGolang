package main

import (
	tc "TaskManager/controllers"
	"fmt"
)

func main() {
	taskData := tc.GetNewTasks()
	ViewMenu(taskData)
}

func ViewMenu(taskData *tc.TasksData) {
	var exit bool
	for !exit {
		fmt.Println("Menu de acciones")
		fmt.Println("1. Agregar")
		fmt.Println("2. Listar")
		fmt.Println("3. Eliminar")
		fmt.Println("4. Cerrar")

		var option int
		fmt.Print("Ingrese una opcion: ")
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			fmt.Println("Agregando Tarea")
			taskData.AddTask()
		case 2:
			fmt.Println("Listando Tareas")
			taskData.ListTasks()
		case 3:
			fmt.Println("Eliminando Tarea")
			taskData.DeleteTask()
		case 4:
			fmt.Println("Saliendo...")
			exit = true
		default:
			fmt.Println("Opcion incorrecta")
		}
	}
}
