# Go Lang: Curso Profesional. 0 -> Master 2023

Start: 03/27/2023
End: 

Sessions:
- 03/27/2023
- 04/02/2023
- 04/14/2023
- 05/09/2023 Hey there
- 05/10/2023
- 05/15/2023 Still Standing, I swear

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

## FMT Input

Para usar Insputs del usuario por la terminal, vamos a usar de fmt Scan, Scanln, Scanf. Con las 
mismas funcionalidades de Print. 
Pero tiene un cambio,
Solo funciona para que el usuario entre los datos. No imprime un label, para eso demobe hacer un 
print antes del Scan. 
Simplemente vamos a tener que pasar la variable que va a recibir el input. Pero si este no es 
pasado casí inmediatamente se va a pasar por el valor default. 
Para solucionar esto, vamos a poner un & antes de la variable para que GO espere hasta que el usuario 
pase un \n u un enter para recibir los datos.

Estos datos los procesa cada uno cómo un string con espacios entre cada valor. Si pasamos más de un valor para el input,
cada espacio que tenga el input, va a ser un valor nuevo.

Las variantes f es para procesar el input cómo standard input o con %d y cosas, ln simplemente con un ln se termina el read 
cuando pasan un \n o EOF

## Arrays 

Los arrays son de tamaño inmutable, y de un solo tipo, definido desde el inicio.

Definición con largo, tipo y valores 
var foo [5]string{ 'a', 'b', 'c', 'd', 'e' }

Se pueden definir sin los valores ya definidos, 
siendo los valores el valor default del tipo de dato:
var foo [5]string;
[ '', '', '', '', '' ]

El Valor ya definido
var foo := [5]string{ 'a', 'b', 'c', 'd' }

Sin el número de valores definidos, se cuenta desde la creación y no puede ser modificado
var foo := [...]string{ 'a', 'b', 'c', 'd' } //len(foo) => 4

Con valores con un indice diferentes:
var foo := [...]string{ 0: 'a', 2:'b', 4: 'c' }
['a', '', 'b', '', 'c'] // len(foo) => 5
Se generan valores defaults en los indices que no estan definidos

Slice de Python en los arrays
var foo [3]int{1, 2, 3, 4, 5}
foo[1:3] => [2, 3, 4]

foo[:3] => [1, 2, 3, 4]
for[3:] => [4, 5]

## Slicen

Los Slicen son practimante pointers a Arrays que son modificables, y que necesariamente no
tenemos que definir su tamaño inicial.

Eso los podemos definir:
cómo

var foo = int[]
foo := []int

O sí usamos el slice de python desde un Array u otro Slice

var a = [4]int
var foo = a[:1]
foo => [0, 0, 0]

Estos Slices tienen, junto a los Arrays, 3 valores
sus valores, [0, 0, 0]
su largo 3, número de valores
y su capacidad, el valor predefinido del array, el de los -> [number]int

se pueden ver estos valores con:
foo
len(foo)
cap(foo)

Podemos agregar valores a un Slice usando 
append(slice, value)

Pero cuando nos pasemos del tamaño normal, este se va a duplicar, haciendo que el 
pointer cambie a otro array, y se re aloque a un slice con el doble de capacidad. 
Cambiando varias cosas sobre este Slice.

Si no necesitamos, o queremos, que cuando creamos un slice sea un pointer y no su 
propio valor. Podemos usar:
copy(copyFoo, foo)

## Make

Make es una función que nos permite crear arrays y slices, teniendo en cuenta sus valores, 
largo y capacidad.

make([]values, len, cap) -> slice

## Maps

Los Mapas son los dicts o Objs literales de GO, los definimos con make,
pasando make en lugar de los valores, y a este le pasamos el tipo de 
valor de las llaves, y el tipo de valor.

foo := make(map[key]value)

los valores se obtienen de la forma general

foo[key] -> value en [key]

En los values, puede almacenar arrs y slices. Para acceder a estos con doble indice foo[key][index]
