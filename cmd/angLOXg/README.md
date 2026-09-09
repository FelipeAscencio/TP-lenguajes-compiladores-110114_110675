# Uso Local de angLOXg

Esta sección detalla cómo compilar y probar el intérprete Lox directamente en tu terminal. El motor está construido 100% en Go, sin dependencias externas, asegurando portabilidad y fácil ejecución.

## Prerrequisitos

* Go 1.22 o superior instalado en el sistema.

## Instalación

Cloná el repositorio y navegá hasta este directorio:

```bash

git clone [https://github.com/FelipeAscencio/angLOXg.git](https://github.com/FelipeAscencio/angLOXg.git)

cd angLOXg/cmd/angLOXg

```

## Modos de ejecución

### 1. Consola interactiva (REPL)

Si ejecutás el programa sin argumentos, se abrirá la consola interactiva donde podés escribir sentencias de Lox línea por línea:

```bash

go run main.go

```

### 2. Ejecución de scripts

Para ejecutar un archivo con código Lox, pasá la ruta del archivo como argumento:

```bash

go run main.go script.lox

```

## Compilación

Si preferís generar el binario ejecutable para no depender del comando `go run`:

```bash

go build -o angloxg

./angloxg [ruta_al_script.lox]

```
