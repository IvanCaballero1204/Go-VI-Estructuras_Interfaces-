package main

import (
	"fmt"
	"./multimedia"
)

func main(){
	var stop bool = false
	var opc int
	imagen := new(multimedia.Imagen)
	audio := new(multimedia.Audio)
	video := new(multimedia.Video)
	contenidoWeb := new(multimedia.ContenidoWeb)

	for !stop {
		fmt.Println(":::MENU:::")
		fmt.Println("1) AGREGAR IMAGEN")
		fmt.Println("2) AGREGAR AUDIO")
		fmt.Println("3) AGREGAR VIDEO")
		fmt.Println("4) MOSTRAR")
		fmt.Println("5) SALIR")
		fmt.Print("OPCION (1-5): ")
		fmt.Scan(&opc)
	
		switch opc {
			case 1:
				input := ""
				auxFloat := 0.0
		
				fmt.Print("Titulo: ")
				fmt.Scan(&input)
				imagen.Titulo = input
				fmt.Print("Formato (jpg, gif, png): ")
				fmt.Scan(&input)
				imagen.Formato = input
				fmt.Print("Canal: ")
				fmt.Scan(&auxFloat)
				imagen.Canal = auxFloat
				contenidoWeb.ContenidoMultimedia = append(contenidoWeb.ContenidoMultimedia, imagen)
			case 2:
				input := ""
				auxInt := 0
		
				fmt.Print("Titulo: ")
				fmt.Scan(&input)
				audio.Titulo = input
				fmt.Print("Formato (mp3, wma, wav): ")
				fmt.Scan(&input)
				audio.Formato = input
				fmt.Print("Duracion (segs): ")
				fmt.Scan(&auxInt)
				audio.Duracion = auxInt
				contenidoWeb.ContenidoMultimedia = append(contenidoWeb.ContenidoMultimedia, audio)
			case 3:
				input := ""
				auxInt := 0
		
				fmt.Print("Titulo: ")
				fmt.Scan(&input)
				video.Titulo = input
				fmt.Print("Formato (mp4, wmv, mov): ")
				fmt.Scan(&input)
				video.Formato = input
				fmt.Print("Frames: ")
				fmt.Scan(&auxInt)
				video.Frames = auxInt
				contenidoWeb.ContenidoMultimedia = append(contenidoWeb.ContenidoMultimedia, video)
			case 4:
				fmt.Print(contenidoWeb.Mostrar())
			case 5:
				stop = true
			default:
				fmt.Println("ERROR OPCION NO VALIDA")
		}
	}
}