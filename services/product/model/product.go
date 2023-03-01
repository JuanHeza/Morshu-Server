package model
/* 
import (
	"EvilPanda/database"
	db "EvilPanda/database"
	dt "EvilPanda/util/dataType"
	"context"
	"fmt"
	_ "fmt"
	"log"
	"net/http"
	"time"
)

type Disponibilidad int64
type Limpieza int64

const (
	_ Disponibilidad = iota
	Disponible
	Pedido
	NoDisponible
	Retirado
	Ilimitado

	cargaRapida Limpieza = iota
	sinAdministrativo
)

type ImagenesProducto struct {
	VarianteProducto string `json:"variante_producto,omitempty" bson:"variante_producto,omitempty"`
	Imagen           string `json:"imagen,omitempty" bson:"imagen,omitempty"`
}
type Product struct {
	DataType dt.DataType               `json:"data_type,omitempty" bson:"data_type,omitempty"`
	Gnrl     generalProductData        `bson:",inline"`
	Admin    administrativeProductData `bson:",inline"`
	Aux      auxiliarProductData       `bson:",inline"`
}
type generalProductData struct {
	Id             string         `json:"_id" bson:"_id"`
	Nombre         string         `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Categoria      string         `json:"categoria,omitempty" bson:"categoria,omitempty"`
	Disponibilidad Disponibilidad `json:"disponibilidad,omitempty" bson:"disponibilidad,omitempty"`
	Cantidad       int            `json:"cantidad,omitempty" bson:"cantidad,omitempty"`
	PrecioBase     float32        `json:"precio_base,omitempty" bson:"precio_base,omitempty"`
	PrecioVariable bool           `json:"precio_variable,omitempty" bson:"precio_variable,omitempty"`
	Oferta         float32        `json:"oferta,omitempty" bson:"oferta,omitempty"`
	PrecioFinal    float32        `json:"precio_final,omitempty" bson:"precio_final,omitempty"`
	OfertaActiva   bool           `json:"oferta_ctiva,omitempty" bson:"oferta_ctiva,omitempty"`
}
type auxiliarProductData struct {
	Tags   []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Imagen []ImagenesProducto `json:"imagen,omitempty" bson:"imagen,omitempty"`
}
type administrativeProductData struct {
	FechaCreacion     int64  `json:"fecha_creacion" bson:"fecha_creacion"`
	FechaModificacion int64  `json:"fecha_modifico" bson:"fecha_modifico"`
	UsuarioCreacion   string `json:"usuario_creacion" bson:"usuario_creacion"`
	UsuarioModifico   string `json:"usuario_modifico" bson:"usuario_modifico"`
	Eliminado         bool   `json:"eliminado" bson:"eliminado"`
}

var (
	imagenesProducto ImagenesProducto
	producto         Product
)

func main_restriction(criteria []dt.Criteria) []dt.Criteria {
	criteria = append(criteria, dt.Criteria{Field: "data_type", Value: dt.Producto_type})
	criteria = append(criteria, dt.Criteria{Field: "eliminado", Value: false})
	return criteria
}
func New(nombre, categoria string, cantidad int, ofertaActiva bool, precioBase, oferta float32, tags []string, imagen []ImagenesProducto) (this *Product) {
	pr := new(Product)
	pr.Gnrl.Nombre = nombre
	pr.DataType = dt.Producto_type
	pr.Gnrl.Categoria = categoria
	pr.Admin.FechaCreacion = time.Now().Unix()
	pr.Admin.FechaModificacion = time.Now().Unix()
	pr.Admin.UsuarioCreacion = "user"
	pr.Admin.UsuarioModifico = "user"
	pr.Gnrl.Disponibilidad = Disponible
	pr.Gnrl.Cantidad = cantidad
	pr.Gnrl.PrecioBase = precioBase
	pr.Gnrl.PrecioVariable = false
	pr.Gnrl.Oferta = oferta
	pr.Gnrl.OfertaActiva = ofertaActiva
	pr.Admin.Eliminado = false
	pr.Aux.Tags = tags
	pr.Aux.Imagen = imagen
	if ofertaActiva {
		this.Gnrl.PrecioFinal = precioBase * ((100 - oferta) / 100)
	} else {
		this.Gnrl.PrecioFinal = precioBase
	}
	return
}


// func (this *Product) Print() (printed string) {
// 	format := `{
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
//         %-16s    %v
// }`

// 	printed = fmt.Sprintf(format, "Id", this.Id, "Nombre", this.Nombre, "Categoria", this.Categoria, "FechaCreacion", this.FechaCreacion, "FechaModifico", this.FechaModificacion, "UsuarioCreacion", this.UsuarioCreacion, "UsuarioModifico", this.UsuarioModifico, "Disponibilidad", this.Disponibilidad, "Cantidad", this.Cantidad, "PrecioBase", this.PrecioBase, "Oferta", this.Oferta, "PrecioFinal", this.PrecioFinal, "OfertaActiva", this.OfertaActiva, "Tags", this.Tags, "Imagen", this.Imagen)
// 	fmt.Println(printed)
// 	return
// }

func (pr *Product) Clean(tipo Limpieza) (output interface{}) {
	switch tipo {
	case cargaRapida:
		output = pr.Gnrl
	default:
		output = pr
	}
	return
}
func (pr *Product) Validate() (valid bool) {
	return
}
func (pr *Product) Create() {}
func (pr *Product) Read()   {}
func (pr *Product) Update() {}
func (pr *Product) Delete() {}
func CreateMany()           {}
func ReadCriteria(r *http.Request, criteria []dt.Criteria) (list []Product) {
	coll := db.GetCollection(r)
	criteria = main_restriction(criteria)
	buildedCriteria := db.BuildCriteria(criteria)
	fmt.Println(buildedCriteria)
	cursor, err := coll.Find(context.TODO(), buildedCriteria)
	if err != nil {
		return
	}
	if err = cursor.All(context.TODO(), &list); err != nil {
		log.Fatal(err)
		return
	}
	database.Disconnect()
	return
}
func UpdateMany() {}
func DeleteMany() {}
 */