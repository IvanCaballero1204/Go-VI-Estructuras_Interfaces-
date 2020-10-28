package multimedia

import "strconv"

type Multimedia interface {
	Mostrar() string
}

type ContenidoWeb struct {
	ContenidoMultimedia []Multimedia
}

func (cw *ContenidoWeb) Mostrar() string {
	var auxWeb string
	for _, contenido := range cw.ContenidoMultimedia {
		auxWeb += contenido.Mostrar() + "\n"
	}

	return auxWeb
}

type Imagen struct {
	Titulo, Formato string 
	Canal float64
}

func (i *Imagen) Mostrar() string {
	auxCadena := ""

	auxCadena += "Image Name: " + i.Titulo + "." + i.Formato + ", Channels: " + strconv.FormatFloat(i.Canal, 'f', 6, 64)
	
	return auxCadena
}

type Audio struct {
	Titulo, Formato string 
	Duracion int
}

func (a *Audio) Mostrar() string {
	auxCadena := ""

	auxCadena += "Track: " + a.Titulo + "." + a.Formato + ", Duration: " + strconv.Itoa(a.Duracion) + " segs"
	
	return auxCadena
}

type Video struct {
	Titulo, Formato string
	Frames int
}

func (v *Video) Mostrar() string {
	auxCadena := ""

	auxCadena += "Video: " + v.Titulo + "." + v.Formato + ", Frames: " + strconv.Itoa(v.Frames)
	
	return auxCadena
}