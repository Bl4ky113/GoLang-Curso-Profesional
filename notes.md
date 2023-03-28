# Go Lang: Curso Profesional. 0 -> Master 2023

Start: 03/27/2023
End: 

Sessions:
- 03/27/2023

## Que es GO?

Es un lenguaje de programación de alto nivel, concurrente y Compilado.
Inspirado en la syntaxis de C y con un rendimiento parecido.

Tiene un tipado tanto dinamico cómo estatico, pero principalmente es estatico.

Tiene parte de compilado, y interpertado, para que sea posible compilarlo y 
usalo en diferentes OS.

Tiene una systexis ya definida cómo Python, y es multiparadigma aunque no maneja cómo 
tal OOP clasico.

No cuenta con errores u exceptions desde la base, pero las podemos aplicar.

Lo fuerte de GO es la concurrencia, que inicia diferentes funcionalidades del programa evitando que este 

## Install y Uso de GO

Hijueputa, que cosa tan dificil de instalar. Pero bueno.

Para installar, obviamente desde pacman en Arch, y todos deben tener su pkg desde sus respectivos 
reps.

Pero no solo eso, ademas cómo es compilado, en nvim, vamos a usar vim-go que es una herramienta que nos 
permite usar diferentes utilidades de go dentro de nvim. Tiene que installar en g:go_bin_path los bin que 
GO va a usar para ejecutarse y demas.
Esto me confunde ya que puede no tenga un control de paquetes u lib cómo python con venv u js - node - npm con packages.json.

Pero ademas de esto vamos a tener que setear en nuestro bashrc el path de los mismos bin exportanto $GOPATH

Desde ahí es para usarlo o compilarlo y ejecutarlo cómo un bin cualquiera.
O simplemente correrlo con $go run script.go

Pero ojo con los archivos, ya que GO toma cómo un dir cómo un paquete de GO, es decir que va a importar y revisar cada archivo .go
al momento de ejecutar y verificar difenrentes cosas. Teniendo errores si esos archivos no deberian estar conectados u 
relacionados.

Usa comentarios de JS

## Variables

Se pueden definir de dos formas, una tipada y una auto tipada
la tipada se usa var, el nombre y el tipo. Mientras que con el auto tipado, solo se usa un operador de asignación

var name string = ""
name := ""

Este automaticamente entiende que es un char.

Algunos tipos son:
- string: " "
- int: 0
- float32: 0
- float64: 0 
- bool: false

Estos valores no tienen null, undefined o none. Solo un init, que son los valores defaults.
Para definirlo con default, solo vamos a usar
var name string

se puede definir variable o constante con 
var lowerCase
y 
const UpperCase
respectivamente.

## Operadores

Los de siempre los aritmeticos
- suma +
- resta - 
- multiplicacion *
- division /
- Modulo %


## Que carajo es fmt

FMT es una implementación de printf y scanf de C. Es decir una forma de I/O para que podamos usar en nuetros paquetes de GO. 
Se pueden usar lo más sencillo cómo prints simples o con ln
fmt.Print()
fmt.Println()
pasando strings u variables para que las impriman.

O se puede usar lo verdaderamente poderoso de fmt que es el Printf y Sprintf, que es imprimir un String 
Formateado, no solo pasando valores, pero implementando formas de imprimirlos  cómo tal.

este formato es

"texto %s variable", variablePasada

Pero estos formatos no solo son formatos sencillos, estos son algunos
- %s strings o slices
- %x base16 lowercase
- %X ... UPPERCASE
- %b base2 binario
- %d base10 ints
- %o base8
- %O base8 con prefijo 0o
- %U unicode 
- %f float decimal normal
- %e notacion cientifica
- %b notacion cientifica sin coma

Con algunos formatos se puede pasar diferentes valores, generalmente con los de floats

La diferencia entre Printf y Sprintf, es que Sprint genera un string apartir del formato para que 
pueda ser usado de otras formas.

Aunque esto solo es la parte de OUTPUT de fmt

## Funciones

tienen la siguiente sintaxis:

func nombreFunction (parametro string, a, b int "mismo tipo para a y b") "tipo al retornar" string {
	return valueString
}

Omitir los "" de la syntaxis

## Tipos de Datos, bien explicados

En general se dividen en varios tipos de datos, cómo C manda.
Algunos solo varian por el tamaño y uso

### Ints
uint 8 - 64 - 2^n bits: Non-signed ints
int 8 - 64 - 2^n bits: signed ints

byte 8bits
rune 32 bits
uint 32 u 64 bits non-signed
int 32 u 64 bits signed

### Floats

float32
float64

### Num muy grandes
complex64
complex128

### Otros Primitivos
string
bool

## GO to read the Docs

En general, que los Docs de GO son bastante complejos, pero que tienen una buena 
cantidad de contenido, ejemplos y demás para aprender de ellos.
