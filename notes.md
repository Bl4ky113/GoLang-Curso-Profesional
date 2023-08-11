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
- 07/25/2023 Still going strong. 4 days in a row
- 07/26/2023 Common, we can do it!!! 5 days in a row!!! Let's get the full week
- 08/01/2023 FUCK. Well anyways, before I enter the Nacho in a proper manner, I WANT THIS DONE
- 08/03/2023 IT'S NOT OVER YET I PROMISE
- 08/04/2023 Common... It's not over yet
- 08/05/2023 Let's go. I'm running out of stuff to say here
- 08/07/2023 Sorry, didn't had my pc at hand :p
- 08/08/2023 First day, it wasn't that rough
- 08/10/2023 Yeah, it's pretty simple all of this stuff.

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

Los Structs se les puede poner un alias (?)
para cuando se tiene que codificar o decodificar las propiedades, y usar otro tipo de syntaxis, Full-CamelCase en GO.

Para esto, alfrente del tipo, vamos a poner en \`\` el tipo de codificación y el nombre despúes de unos :

type Nombre struct {
    Property string \`json:property\`
}

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
No puede dar error al no usar _ (?) No, no da error

## Advanced Functions (multiple parameters)

para tener varios parametros en una function, vamos a agregar en los parametros ... y el tipo del valor.
Esto nos va a dar un Array con los datos qué se le pasen a la function, sin tener que definir cuantos valores y
a cada uno un nombre de variable.

Pero si queremos no definir el tipo de valor que vayamos a manejar, podemos simplemente agregar en el tipo del
valor

...interface{}

Y ahí podra recibir cualquier tipo de valor.

Si necesitamos si o si tener unos parametros, es necesario pasarlos antes de la lista de parametros.

## Advanced Functions (recursive funtion)

Las functions que son recursivas son complejas de por sí, pero la forma
de definirlas no tanto.

Simplemente vamos a llamar en un return la misma function, asegurandonos qué la function va a tener un fin o
qué se cierre un loop, retornando un valor donde no se llame a la function.

## Advanced functions ( anonymous functions )

Se pueden hacer anonymous functions generando una function dentro de otra, pero despues de los {} 
agregar los () con los parametros necesarios, para qué una vez se creé, se ejecuté.

func () {

}()

Esto también se puede definir sin los () a una variable, donde la podemos usar cómo si fuera una function normal.
Aprovechando esta funcionalidad, se puede también pasar a otras functions cómo parametros, donde podemos definir que tipos de parametros
tienen que obtener y también que return necesitan. Supongo que si necesitan una ref de un interface, también es posible de definir.

## Advanced functions ( high order functions )

Son simplemente functions que usan otras functions para su funcionalidad usando la funcionalidad de go de 
poder definir un parametro cómo una function, que parametros recibe y que return hace.

## Advanced functions ( closure function )

Las closure functions son más que todo functions que definen un ambiente de variables con valores, para crear y retornar
una nueva anonymous function, qué generalmente se va a ejecutar mediante una variable.

Las instancias de la anonymous functions van a tener el closure de datos y variables qué haya definido la function.
Podiendo ejecutarse varias veces y compartir los mismos datos.

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

## Error triggering - handle in GO

Los errores no se pueden hacer un trigger, o un throw, u algo cómo la gran mayoria de lenguajes de programación.
Pero son totalmente manejados por el código cómo tal. Ya hemos visto unos casos.
Y la forma de mandar y mostrar errores es que junto al return de las funciones, si ocurre un error, vamos a 
retornar el valor base del tipo de return y un error.

Qué NO es un tipo de dato, pero un a interfaz que simplemente retorna el str del error cómo tal. De pronto se puede 
modificar para que tenga cosas cómo traces, codes, y demás. Pero quien sabe.

La forma sencilla de crear un error custom, es usando la lib integrada de errors. 
Donde si queremos crear un nuevo error, simplemente vamos a usar su method New y pasar el str del error.

Desde el return, donde se implemente la function, vamos a hacer un handle apropiado del error en cuestión.

## defer

defer es una forma única en la que podemos 'atrasar' la ejecución de un código, que es más oportuno para 
cosas cómo cerrar y guardar las cosas apropiadamente. Cómo al editar un archivo, debemos cerrarlo una vez 
lo hayamos usado. Esto para evitar que no tenga los cambios hechos u ocurra otro error.

Este funciona agregando diferentes instrucciónes a un stack, donde el primero en entrar, es el primero en salir.
No creo que sea todo lo que se pueda hacer con esto, más que un stack de acciones a realizar el código una vez termine 
de ejecutar todas las otras acciones.

## PANIC at the disco

Lo anterior era una forma de generar errores y manejarlos hasta cierto nivel,
pero con panic(). Es un clasico error que rompe el código, 
donde nos muestra un str de error, el cual se pasa cómo parametro, y el goroutine de cómo ocurrio y 
donde el error.

Esto rompe el código y evita su ejecución, pero podemos agregar un catch o un execption, dentro de nuestro código junto a defer. 
Para esto vamos a crear una anonymous function, donde vamos a usar recover(). Qué es una function interna que nos permite hacer
catch de cualquier error que panic lanza, sin romper la ejecución del código, más que usar o realizar el handle de la anonymous function.

PERO OJO. Esto no es recomendado hacerlo en el código de verdad, en general, usar los metodos anteriores. Y usar solo estos para 
manejos más globales y no tan especificos, o usarlos cómo sería un catch en cualquier otro lenguaje de programación.

## Logging stuff

En Go, podemos hacer logs avanzados, para evitar usar solo fmt.
Esto se hace con la lib de log, que nos permite hacer diferentes logs, configurarlos y 
inclusive interactuar con el runtime del programa ejecutando un Fatal o un Panic

los logs normales son lo de Print, Printf y Println, que imprimen la fecha y hora junto al mensaje 
que hayamos pasado.

Se puede interactuar usando Fatal, qué hace que el programa haga exit con code 1. Es decir se detenga todo.
Y se puede hacer un Panic, qué es lo mismo pero con code 2, mostrando el runtime y datos de donde se llamo el 
Panic.

Se pueden configurar, por ejemplo, agregando un prefifo cuando se generen los logs dentro de un módulo, esto con: SetPrefix(str). 
Se puede destinar los logs a una instancia de un Archivo Abierto usando: SetOutput(file)

La fecha y hora de los logs se pueden des/activar usando flags, que son la información que se va a pasar. 
No se especifico mucho, pero pasando a SetFlags(0), se quitan la fecha y hora.

## File System y Operating System

Para crear y abrir archivos debemos usar al lib de OS, y usar el method de OpenFile
este recibe el nombre del archivo, las flags de uso, cómo: create, append, read - write only, etc.
y los permisos qué se setear en el archivo.

Pero sí o sí, para guardar los cambios y el archivo cómo tal. Debemos despúes de usarlo usar el method Close de la 
instancia del archivo. Una forma sencilla es ponerlo con un defer justo al lado de la creación del archivo.

Al momento de crear o abrir el archivo, también puede retornar un error al estilo de handle básico de errores.

Se pueden crear archivos solo usando Create y pasando el nombre del archivo

## JSON

Con el módulo Json de GO se puede crear un encoder o decoder de JSON apartir de una instancia de un archivo.
Este con el method NewDecoder o NewEncoder. 

Y deade ahí se puede usar la instancia del decoder o encode para leer o escribir un archivo JSON.

Decode recib ela referencia de un slice o Array que va a recibir los datos.
Mientras que Encode recibe un slice o Array qué es lo que va a escribir en el JSON.

Estos generalmente van a tener que ser structures, pero cómo se puede definir que las propiedades tengan un 
alias cuando sean decodificadas o codificadas de o a JSON. No hay tanto problema al momento de 
manipular los datos.

En general el manejo de los archivos de-codificadores. pueden retornar error, por eso es mejor tener un fallback
listo para casos así.

## Modules: Create, Install, Upload, and Share

Para Crear un modulo, vamos a ir desde la terminal:

$ go mod init $MODULEPATH

Y ahí se va a crear un archivo go.mod
este va a definir diferentes configuraciones del module cómo tal. Cómo los 
requirements que se necesitan para poder ejecutar el módulo. 

En el archivo main, vamos a definir una linea al inicio llamada package que va a ser desde 
donde se va a usar cómo tal el module, cómo si fuera de una interfaz.


Estos se pueden crear redirecciones para modules que tenemos instalados localmente.

Para esto vamos a usar :

$ go mod -replace url/requirement/module=../dir/requirement/module

Y para definirlo cómo una dependencia, vamos a usar 

$ go mod tidy
Qué nos permite instalar las depencias que no tengamos y eliminar que no necesitamos.

Una vez instalada el module, lo podemos usar según el nombre del package definido, y accediendo a 
sus function cómo si fuera una interfaz

## Testing on Go

Con el módulo testing podemos correr tests en Go.
vamos a primero crear un archivo con el mismo nombre de nuestro module, definir el mismo nombre de package, 
pero le agregamos al nombre del archivo \_test.go para que go lo reconozca cómo un test.

desde ahí vamos a crear functions que inicien con Test y que reciban cómo parametro un pointer de
testing.T. Que nos va a permitir interactuar, hacer logs, mandar errores y demás con el testing de go.

Desde ahí podemos hacer más que todo testcases que sean correctos, cómo que salten errores.
Y usando el parametro de testing.T podemos llamar Fatalln, Fatalf, o Fatal que informa al test suite que 
ha ocurrido un error.

Lo que no termino de entender, es que en el Fatal, manda un string con backquote o \`\`. Pero supongo que 
es para poder hacer un formato correcto de un Regex.

## Regex

Existe un módulo integrado de Go para manejar Regex. Nada raro, pero no se explica nada a detalle.
Además de que al momento de crearlo se debe usar backquotes o \`\`.

## Compilar y Installar

Simplemente una vez instalado con

$ go compile file.go

vamos a poder usar el script cómo un bin. El cual generalmente se puede dejar en ~/.local/shared/bin o ~/.go/bin 
para mejor acceso y uso.

Sobretodo ~/.go/bin para uso de librerias y uso de otros scripts.

## Subir a Github

Simplemente lo de siempre, repo, coger el remote, commit, y push. Podemos ponerle un README donde se explique qué es 
cómo se usa. Docs generales. 

Ya para instalar vamos a usar go get -u y la url parcial del repo. Tipo

github o dominio / usuario / repo

$ go get -u github.com/bl4ky/something

Y ahí se añadira de una vez al go.mod cómo una dependencia. Y se creara un archivo go.sum, pero no se especifica para qué es.
En el import dentro de nuestros scripts, vamos a usar el mismo formato de url, si es que no lo redireccionamos.

## i'm GO to POO

GO cómo tal no esta diseñado para ser usado totalmente cómo un lenguaje de POO. 
Pero tiene varias estructuras que nos puede permitir emular estas acciones. 

## POO Classes

las clases cómo tal no existen en Go, para esto vamos a tomar las interfaces y los types de Go. 
Donde vamos a definir propiedades y usando functions que necesitan una referencia de un type. Podemos agregarle 
methods a estas 'classes'.

Pero si necesitamos 'simular' un constructor. Vamos a tener que si o si usar una function que retorne una referencia del type o interface 
de la clase.

Y qué desde ahí reciba los parametros inciales, y realice sus operaciones de constructor cómo tal. 
Obviamente si tiene condicionales qué hace que no pueda funcionar, retornar también un error.

## Encapsulización 

En Go, solo podemos hacer las propiedades y methods de las interfaces privadas y públicas.
Es bastante sencillo, simeplemente usamos minuscula o mayuscula, respectivamente para cada caso.

## Interfaces, de verdad

Las interfaces de struct funcionan especialmente para cosas cómo propiedades, pero cuando se trata de methods, 
toca hacer vueltas raras con functions y pointers.

Para eso se puede implementar una interface. La cual no hace todo más fácil sinceramente, pero se puede simplificar el 
proceso de tener qué usar el mismo method con diferentes o tipos relacionados de objs.

Con la siguiente syntaxis

type Nombre interface {
	method(params)
}

No se define cómo tal el method, este DEBE ser hecho por fuera del interface. Pero este method puede ser uno que se repita, pero teniendo diferentes 
refs de pointers, de diferentes tipos de clases.

## Composition y Polymorphism

Básicamente, se pueden llamar y crear diferentes functions con el mismo nombre, pero con diferente ref pointer, o hacerlas con un
tipo de obj diferente.

Los conceptos de esto es que se puede crear un nuevo tipo de obj apartir de otro, para el cambio o añadido de 
methods, es simplemente crear una nueva function con un refpointer diferente.

Pero para properties, debemos hacer más vueltas, pero simples. 
Vamos a agregar las nuevas properies, o las modificadas cómo cualquier otra. Y 
vamos a 'agregar' cómo si fuera una el tipo de obj de donde queremos hacer inheritance.
Y cuando vayamos a hacer un init cómo tal del obj. Pasarle las propiedades al tipo de obj padre cómo 
si fuera un init normal de esté pero sin constructor. 

type Name struct {
	ParentType
	...
}

Pero al momento de acceder a las propiedades, podemos acceder normalmente a estas. 
O mejor dicho sin tener que acceder cómo si fuera un subnivel

NO. Name.ParentType.ParentProperty
YES. Name.ParentProperty

## Concurency

La concurrencia en Go es practimante implementada de la forma más sencilla y simple posible.
Permitiendonos ejecutar código de una forma paralela con diferentes núcleos de nuestro procesador
de forma dinamica. 

Para esto, simeplemente vamos a usar la palabra reservada: go 
y listo!

Generalmente, y más utíl cuando se tienen que hacer varias acciones con una function apartir de una lista de datos 
cómo un Array.

Pero esto puede qué se ejecute tan rápido que el programa termine sin dejar prints, logs o demás acciones.

## Channels

Los channels son una forma en la que podemos obtener y usar datos async. Tienen un formato diferente de generación y 
uso frente a las variables normales. Pero es extremadamente utíl cuando se usan junto a la concurency.

Para definir un nuevo channel, vamos a usar make y el tipo de dato qué se espera que vaya a tener.

ch := make(chan string)
Desde ahí podemos usar un tipo de operador especial (- *pero con un \< \> qué esta jodiendo mucho el syntax

Qué vamos a usar a la derecha del channel para guardar o agregar datos a este, y a la izquierda cuando vayamos a usar estos datos.

Esto nos va a arreglar el error anterior de la concunrency. Qué no se hicieran bien los logs.

Pero algo qué no me cuadra es que al momento de hacer un print de los datos del channel. El hp del curso lo hace en un for loop aparte
Voy a probar hacerlo en el mismo loop y fuera. Para ver si cambia algo

Dentro -> misma cosa, pero creo que esta un poco lento
Fuera -> Mierda, no son la misma cosa, se pueden hacer. Pero es mucho mucho más rápido.

## Generics in New Go 1.18

Go has fallen...

Con el ejemplo anterior de avanced functions. El multiple parameter con un tipo de dato raro usando 
interface{}. Acepta cualquier tipo de parametro. Esto lo podemos cambiar por simeplemente 
un any. 

Funcionara de la misma manera.

## Restrictions on Any

Se pueden crear Restrictions en las functions para usar dentro de los parametros que se van a recibir 

la syntaxis es:

func foo [NewType int] (param NewType) {
	...
}

Con el NewType podemos definir que tenga varios tipos, pero qué se define solo uno cuando se usa la function

[NewType int | float32]

Y si necesitamos o queremos usar valores que extienden a los tipos de valores del NewType, nos dara error ya que no son 
exactamente el mismo tipo, pero podemos hacer que este se aproxime usando ~ al inicio de cada tipo de NewType

[NewType ~int | ~float32]

Podemos agregar varios tipos al restriction, pero esto haría que nuestra function tubiera un nombre extremadamente largo
para evitar esto, y poder hacer restrictions más complejos podemos usar type interfaces. 

Siendo la syntaxis

type constraints interface {
	type | type2 | ~typ3
}

De ahí podemos usar otros constraints hechos con Interfaces, haciendo cada vez más complejos estos mismos.
Algunos se encuentran disponibles en el packet manager de Go.

Estos constraints una vez definido el tipo del valor, se pueden usar operadores de comparación cómo si fuera el mismo valor.
Pasando el tipo de valor final a callbacks para tambien poder ser usado cómo el tipo de valor definido.

Las structs tampoco se salvan. 

Se puede agregar, despúes del nombre, un [] con el nombre y tipos de datos para un dato generico.
Donde lo podremos usar el tipo generico en las interfaces.


