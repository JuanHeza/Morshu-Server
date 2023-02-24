package main

import (
	"fmt"
	"time"
)

type Disponibilidad int64
type Limpieza int64

const (
	Disponible Disponibilidad = iota
	Pedido
	NoDisponible
	Retirado
	Ilimitado

	cargaRapida Limpieza = iota
	sinAdministrativo
)

type ImagenesProducto struct {
	VarianteProducto string `bson:"variante_producto,omitempty"`
	Imagen           string `bson:"imagen,omitempty"`
}
type Product struct {
	Id                string             `bson:"_id,omitempty"`
	Nombre            string             `bson:"nombre,omitempty"`
	Categoria         string             `bson:"categoria,omitempty"`
	FechaCreacion     int64              `bson:"fecha_creacion,omitempty"`
	FechaModificacion int64              `bson:"fecha_modifico,omitempty"`
	UsuarioCreacion   string             `bson:"usuario_creacion,omitempty"`
	UsuarioModifico   string             `bson:"usuario_modifico,omitempty"`
	Disponibilidad    Disponibilidad     `bson:"disponibilidad,omitempty"`
	Cantidad          int                `bson:"cantidad,omitempty"`
	PrecioBase        float32            `bson:"precio_base,omitempty"`
	precioVariable    bool               `bson:"precio_variable,omitempty"`
	Oferta            float32            `bson:"oferta,omitempty"`
	PrecioFinal       float32            `bson:"precio_final,omitempty"`
	OfertaActiva      bool               `bson:"oferta_ctiva,omitempty"`
	Eliminado         bool               `bson:"eliminado,omitempty"`
	Tags              []string           `bson:"tags,omitempty"`
	Imagen            []ImagenesProducto `bson:"imagen,omitempty"`
}

var (
	imagenesProducto ImagenesProducto
	producto         Product
)

func New(nombre, categoria string, cantidad int, ofertaActiva bool, precioBase, oferta float32, tags []string, imagen []ImagenesProducto) (this *Product) {
	this = &Product{
		Nombre:            nombre,
		Categoria:         categoria,
		FechaCreacion:     time.Now().Unix(),
		FechaModificacion: time.Now().Unix(),
		UsuarioCreacion:   "",
		UsuarioModifico:   "",
		Disponibilidad:    Disponible,
		Cantidad:          cantidad,
		PrecioBase:        precioBase,
		Oferta:            oferta,
		OfertaActiva:      ofertaActiva,
		Eliminado:         false,
		Tags:              tags,
		Imagen:            imagen,
	}
	if ofertaActiva {
		this.PrecioFinal = precioBase * ((100 - oferta) / 100)
	} else {
		this.PrecioFinal = precioBase
	}
	return
}

func (this *Product) compare(that *Product) (result bool) {
	result = this.Nombre == that.Nombre && this.Categoria == that.Categoria && this.FechaCreacion == that.FechaCreacion && this.FechaModificacion == that.FechaModificacion && this.Disponibilidad == that.Disponibilidad && this.Cantidad == that.Cantidad && this.PrecioBase == that.PrecioBase && this.Oferta == that.Oferta && this.PrecioFinal == that.PrecioFinal && this.OfertaActiva == that.OfertaActiva && this.Eliminado == that.Eliminado
	if result {
		result = producto.compareTags(this.Tags, that.Tags) && imagenesProducto.compare(this.Imagen, that.Imagen)
	}
	return
}

func (pr Product) compareTags(this, that []string) (result bool) {
	return
}

func (imPR ImagenesProducto) compare(this, that []ImagenesProducto) (result bool) {
	return
}
func (this *Product) Print() (printed string) {
	format := `{
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
        %-16s    %v
}`

	printed = fmt.Sprintf(format, "Id", this.Id, "Nombre", this.Nombre, "Categoria", this.Categoria, "FechaCreacion", this.FechaCreacion, "FechaModifico", this.FechaModificacion, "UsuarioCreacion", this.UsuarioCreacion, "UsuarioModifico", this.UsuarioModifico, "Disponibilidad", this.Disponibilidad, "Cantidad", this.Cantidad, "PrecioBase", this.PrecioBase, "Oferta", this.Oferta, "PrecioFinal", this.PrecioFinal, "OfertaActiva", this.OfertaActiva, "Tags", this.Tags, "Imagen", this.Imagen)
	fmt.Println(printed)
	return
}
func (this *Product) Clean(tipo Limpieza) (output interface{}) {
	switch tipo {
	case cargaRapida:
		output = struct {
			Id             string         `json:"_id,omitempty"`
			Nombre         string         `json:"nombre,omitempty"`
			Categoria      string         `json:"categoria,omitempty"`
			FechaCreacion  int64          `json:"fecha_creacion,omitempty"`
			Disponibilidad Disponibilidad `json:"disponibilidad,omitempty"`
			Cantidad       int            `json:"cantidad,omitempty"`
			PrecioFinal    float32        `json:"precio_final,omitempty"`
		}{
			Id:             this.Id,
			Nombre:         this.Nombre,
			Categoria:      this.Categoria,
			FechaCreacion:  this.FechaCreacion,
			Disponibilidad: this.Disponibilidad,
			Cantidad:       this.Cantidad,
			PrecioFinal:    this.PrecioFinal,
		}
	default:
		output = this
	}
	return
}
func (this *Product) Validate() (valid bool) {
	return
}
func (this *Product) Create() {}
func (this *Product) Read()   {}
func (this *Product) Update() {}
func (this *Product) Delete() {}
func CreateMany()             {}
func ReadCriteria()           {}
func UpdateMany()             {}
func DeleteMany()             {}
