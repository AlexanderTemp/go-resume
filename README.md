## 1. Lo básico

Para trabajar con paqueteria `go.mod` ejecutar comando:

```bash
go mod init  [aquí iría el modulo] example/hola_mundo
```

El nombre del **módulo** el que va en go.mod debe coincidir con la **ubicación real del repositorio** donde está el código

Si se subiese a github _/alexander/myproject_ el módulo en go.mod debe declararse como:

```bash
module github.com/alexander/myproject
```

De esa manera cuando otra persona quiera usar será:

```
got get github.com/alexander/myproject
```

> [!tldr] TLDR
>
> - **Código privado/local** Si solo lo usarás tu, el path puede ser inventado `example.com/mimodulo` o algo más corto, no pasa nada
> - **Código público/distribuible** Aquí si importa el path que sea real y accesible, sobre github, gitlab, etc.

### 1.1. Estructura base

- [Paquetería standard](https://pkg.go.dev/std)

```go
package main # paquete main (agrupar funciones)

import "fmt" # contiene funciones de formateo de texto revisar librería standard

func main() {
	fmt.Println("Hola mundo")
}

```

```
go run .
go help
```

### 1.2. Instalación paquetería externa

- [pkg.go.dev](https://pkg.go.dev/search?q=quote)

![[Pasted image 20250828095829.png]]

Se importa en el codigo el copiado

```go
package main

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

Luego se ejecuta `go mod tidy` para actualizar la paquetería

## 2. Crear un módulo
