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
- 06/18/2023 God Dangmit. They Updated the Course
- 06/24/2023 Lets finish this thing before the month ends!!!
- 07/22/2023 Gues what? I couldn't :C But we can make it
- 07/23/2023 Common, we can do it
- 07/24/2023 Let's go!!!

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

con const, aveces se necesita setear diferentes variables al mismo tiempo. 
Para hacer esto, simplemente vamos a tener que usar const cómo si fuera una function. 
pasando los valores con un ().

Desde ahí podemos definir de forma normal las variables. O podemos usar IOTA.

## IOTA

IOTA es un tipo de operador que nos permite definir consts de una forma ciclica.
el formato de uso es:

const (
	foo int = IOTA // 0
	bar  // 1
	foo1 // 2
	bar1 // 3
)

IOTA incrementa en 1, desde 0, siendo un int esto se puede modificar para cualquier op de número.
Algunos ejemplos robados, pero curiosos que ví:

Con corridos de memoria, se define en bin los permisos UNIX.
1 READ
2 WRITE
4 EXECUTE
const (
	read   = 1 \<\<> iota \/\/ 00000001 = 1
	write              \/\/ 00000010 = 2
	remove             \/\/ 00000100 = 4

	\/\/ admin will have all of the permissions
	admin = read \| write \| remove
)

Desde que se define IOTA, cada valor consecuente va a ser definido según cómo indique el paso del ciclo.
Para evitar o saltarse un valor del ciclo, se usa el operador _ entre las variables

const (
	foo int = IOTA // 0
	_ // 1
	bar // 2
)

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

byte 8bits: se pueden usar para almacenar también solo un char
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
rune: se puede usar para guardar valores unicode
bool

## Conversión de Tipos de Datos

### Numbers

La forma más sencillas de conversiones son de números, simplemente vamos a usar 
las variables definidas de los tipos de datos, que es el nombre del tipo y pasamos el 
número cómo tal a convertir. Utíl si vamos a operar diferentes tipos de ints.

int8 + int32 = ERROR
Salta error siempre y cuando el resultado sea mayor a int8

### Strings

Para strings se ocupa un módulo completo de GO llamado strconv el cual tiene cómo functions.
Los cuales las más sencillas son 

Atoi str -> int
y
Itoa int -> str

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

Nosotros podemos intentar obtener cualquier valor de los maps usando sus keys. Pero si esta no existe nos puede dar un error.
Este valor es pasado junto al valor de la llave, si tiene, cómo un return multiple

foo, fooExists := anyMap['fooKey']

fooExists == true Si sí existe un valor con la llave fooKey
false si no.

Esto se puede conjugar con las definiciones de scope en los Ifs definiendo el valor y el verificador, para 
despues validar el verificador

Lo más probable es que se pueda usar en todos los tipos de estructuras de datos, pero en Maps.
para eliminar un valor según su key, vamos a usar delete(map, key)

Se puede iterar por todo el map usando for y pasando al range el map, el cual va a 
retornar cómo variable la key y su val.

for key, val := range map {
	...
}

## Structures 

Types o Structures son formas en las que se puede generar un tipo de obj 'generico' donde 
se definen las propiedades del object cómo si fuera un interfaz de un tipo de obj.

En el ejemplo dado solo se define un tipo con datos sencillos, pero creo que se puede con estructuras de datos
y functions.

para definirlo se sigue la syntaxis

type Nombre struct {
	property type
	... ...
}

Para usarlo, vamos a definir una variable con el tipo de Nombre, o el nombre del struct.
Podemos acceder a las propiedades usando

var foo Nombre
foo.property
Donde las podemos definir y llamar.

Pero si no queremos usar var, podemos hacerlo con el := Llamandolo cómo si se fuera a definir los valores cómo tal,
se pueden definir por orden o por nombre las propiedades

foo := Nombre{val}

o

foo := Nombre{property: val}

### Methods & Pointers

Primero los pointers, los pointers son formas en las que en la memoria del pc se guardan los datos. 
Generalmente se trabajan apartir de valores No-globales, o modificando la memoria de las variables cuando se definen algunas 
cosas cómo estas variables

x := 10
y := x

Si yo modifico x, y va a cambiar. Pero si modifico y, x no va a cambiar, porqué? Por qué y es una copia de x. No tiene el mismo valor o tiene la misma
memoria. Para hacer que sea el mismo valor, simplemente le ponemos 10, o una copia con copy().

Pero si necesitamos que si modificamos y, se modifique x. Debemos usar la referencia de la memoria de x y guardarla en y.

Esto se define con la syntaxis

var x int = 10
var y *int = &x

Diciendo que *int es una referencia de un int y que se asigna la referencia en memoria de x.
Si modificamos y en cualquier scope, cómo se esta modificando la memoria cual apunta x, x cambiara también.

Esto se puede aplicar pasando en functions que reciban parametros que sean referencias con *int o *type.
Y usandolo internamente con *parameter, para que se defina que es un pointer


Ahora los methods, los methods son una forma en la que podemos definir metodos para nuestras propiedades de los structs, 
no se definen internamente desde el struct, pero creamos una function la cual tenga un receptor el cual es una variable que es un pointer 
de una instancia del struct. Añadiendo una nueva syntaxis al func

type Person struct {
	property string
}

func (person *Person) getProperty () {
	ftm.Printf("property: %s", person.property)
}

Estos metodos pueden modificar y llamar las propiedades de los structs, pero tienen que ser llamados y usadas desde 
una instancia de person:

per := Person { ... }

p.getProperty() -> Funciona!
getProperty() -> NO funciona...

## Operadores de Booleans

No son operadores de Booleans, el tipo de dato, pero operadores lógicos, básicamente los mismo de toda la vida de C.
Siguen la secuencia de 'resolución' (?) de PEMDAS

Parentesiss
Exponentes
Multiplicación
División
Suma
Resta

## If - Else

El if esle tiene un uso relativamente pytonica, pero tiene un truco cómo si fuera un for loop de C.

Se define cómo 

if condicional {

} else if condificonal {

} else {

}

No es necesario tener la condicional con (), pero dentro de esta se pueden 
definir variables que se pueden usar dentro del scope del If. Definiendolo y 
continuando la condicional con un ;

if foo := 'val'; foo != 'val' {
	...
	fmt.Println(foo) -> 'val'
}

## Switch Makes a Comeback

Switch tiene la syntaxis normal de C.
Solo que se puede declarar sin (). Y agregando el extra del settear variables al nivel del scope.
Algo extra es que no se tiene que declarar el break del case. Simplemente se define un default o 
otro case nuevo.

Pero ademas de esto, podemos dejarlo con condicional inicial, o solo con la variable de scope.
Y podemos poner condicionales en cada case del switch, cosa que no se puede hacer en la gran mayoria de 
lenguajes.
La variable de scope se puede usar en los condicionales de los cases.

## For Loop

Sigue la tendencia de las estructuras anteriores. Misma syntaxis que la de C, pero sin parentesis.
Y hace uso de break para romper todo el Loop, o continue para saltarse o pasar a la siguiente iteración.

Se puede crear un for loop infinito, simplemente no definiendo nada, y solo los {}.

## Functions

Tiene un revoltijo de ni el Hps

Se declaran con la palabra reservada func. Nombre parametros entre (), posible return y scope {}.
En los parametros, se tienen que declarar el tipo de estos. con:

func foo (bar string) {}

Pero si se tienen varios parametros con el mismo tipo, se pueden definir todos con el mismo tipo

(foo, bar, a string). Y todos seran de tipo string

En los datos de posible return, se puede definir que se van a devolver más de un valor separando el tipo 
del valor con comas

func foo () int, int {
	return 1, 0
}

Pero este return, puede definir también variables dentro del scope de la function.
Dejando dentro de un () el nombre de las varaibles y su tipo

func foo () (bar string, see string) {
	fmt.Println(bar, see)
	return bar, see
}

Haciendo que la function siempre tenga un return. Y no haya problemas al definir el tipo del return (?)
Sinceramente es pura mierda, pero bueno.

Si los valores a retornar definidos también tienen los mismos tipos, se pueden definir para varios

func foo () (bar, foo string) {
	return bar, foo
}

Estas functions que no retornan un Array o Slicen completo, se pueden obtener los returns 
definiendo dos variables cuando se guarde y se llame la function.

foo, bar = foo() -> foo, bar

Pero si no necesitamos un valor, NO podemos dejarlo volando, por eso lo mejor es dejarlo cómo una 
variable con nombre \_

\_, bar = foo() => foo, bar ... solo se guarda, o se usa, bar
No puede dar error al no usar _ (?)

## Matrices

Son las mismas cosas que hemos visto de un Array de Go

Definido con

var foo [numValores]tipo

Pero algo extra es la posibilidad de iterar atravez de los valores de los Arrays usando el for loop

foo = [5]int{1, 4, 7, 9, 10}
for i, val := range foo {
	fmt.Printf('i: %d, val: %d', i, val)
}

Tomando el index del valor en i y el valor en val

Se pueden formar Arrays multi dimensionales definiendo con dos [][] y al momento de definirlo,
hacerlo con {} dentro de {}, { {}, {}, {} }.


