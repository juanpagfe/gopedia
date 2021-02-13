#Gopedia

Very small Wikipedia scrapper to query short concepts written in Go. This is a project example I have developed to start working on the language.

##Compile
Assuming you have already [GIT](https://git-scm.com/) and [Golang](https://golang.org/) on your system.

```bash
git clone https://github.com/juanpagfe/gopedia.git
cd gopedia
go build
```

##Usage

```bash
./gopedia Golang
#Go is a statically typed, compiled programming language designed at Google[14] by Robert Griesemer, Rob Pike, and Ken Thompson.[12] Go is syntactically similar to C, but with memory safety, garbage collection, structural typing,[6] and CSP-style concurrency.[15] The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.[16]
#There are two major implementations:
#A third-party transpiler GopherJS[21] compiles Go to JavaScript for front-end web development.
```

To query on a different language (ISO-639-1 compatible)

```bash
./gopedia -l es Golang
#Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico como Python y con el rendimiento de C o C++. Ha sido desarrollado por Google[9]​ y sus diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson. [10]​ Actualmente está disponible en formato binario para los sistemas operativos Windows, GNU/Linux, FreeBSD  y Mac OS X, pudiendo también ser instalado en estos y en otros sistemas mediante el código fuente.[11]​[12]​ Go es un lenguaje de programación compilado, concurrente, imperativo, estructurado, orientado a objetos y con recolector de basura que de momento es soportado en diferentes tipos de sistemas UNIX, incluidos Linux, FreeBSD, Mac OS X y Plan 9 (puesto que parte del compilador está basado en un trabajo previo sobre el sistema operativo Inferno). Las arquitecturas soportadas son i386, amd64 y ARM.
#El día de la publicación del lenguaje Go, Francis McCabe, desarrollador del lenguaje de programación Go! (anteriormente llamado Go), solicitó que se le cambiase el nombre al lenguaje de Google para evitar confusiones con su lenguaje.[13]​ McCabe creó Go! en el año 2003; sin embargo, aún no ha registrado el nombre.[14]​
#Go es un nuevo lenguaje de programación para sistemas lanzado por Google en noviembre de 2009. Aunque empezó a ser desarrollado en septiembre de 2007 por Robert Griesemer, Rob Pike y Ken Thompson.
#Go no utiliza excepciones.[21]​ Los creadores del lenguaje han dado varios motivos para que esto sea así. La principal es que añadir una capa de excepciones agrega una complejidad innecesaria al lenguaje y al entorno de ejecución. Por definición las excepciones deberían ser excepcionales pero al final se acaban usando como controladores del flujo de la aplicación y dejan de tener la finalidad de excepcionalidad. Según los creadores, las excepciones tienen que ser realmente excepcionales y el uso que se le da mayoritariamente no justifica su existencia.
```

If you want less information (or more) specify the amount of parragraph you want to read (Default 3).

```bash
./gopedia -l es -p 1 Golang
#Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico como Python y con el rendimiento de C o C++. Ha sido desarrollado por Google[9]​ y sus diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson. [10]​ Actualmente está disponible en formato binario para los sistemas operativos Windows, GNU/Linux, FreeBSD  y Mac OS X, pudiendo también ser instalado en estos y en otros sistemas mediante el código fuente.[11]​[12]​ Go es un lenguaje de programación compilado, concurrente, imperativo, estructurado, orientado a objetos y con recolector de basura que de momento es soportado en diferentes tipos de sistemas UNIX, incluidos Linux, FreeBSD, Mac OS X y Plan 9 (puesto que parte del compilador está basado en un trabajo previo sobre el sistema operativo Inferno). Las arquitecturas soportadas son i386, amd64 y ARM.
```

To print the help.

```bash
./gopedia -help
#Usage of ./gopedia:
#  -h	Print help
#  -l string
#    	ISO-639-1 Language code. (default "en")
#  -p int
#    	Max number of paragraphs. (default 3)
```