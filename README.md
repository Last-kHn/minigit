#MiniGit
##Descripción del Proyecto

MiniGit es un sistema simplificado de control de versiones desarrollado en Go con fines educativos. 
El proyecto busca representar el funcionamiento básico de herramientas como Git mediante una implementación ligera basada en terminal.

Actualmente el sistema permite:

* Inicializar repositorios
* Agregar archivos al área de staging
* Realizar commits
* Visualizar historial de commits
  
# Tecnologías Utilizadas

* Go
* Git
* GitHub
* Linux Terminal

# Estado del Proyecto

Proyecto en desarrollo – Avance 30%

Funcionalidades implementadas:

* Repositorio `.minigit`
* Área de staging
* Sistema de commits
* Historial básico
* Uso de ramas Git (`main` y `develop`)

# Estructura del Proyecto

```text
Minigit/
├── main.go
├── repository.go
├── staging.go
├── commit.go
├── log.go
├── go.mod
└── README.md
```

---

# Ejecución del Proyecto

## Compilar

```bash
go build -o minigit
```

## Ejecutar

```bash
./minigit
```

---

# Comandos Implementados

## Inicializar repositorio

```bash
./minigit init
```

## Agregar archivo

```bash
./minigit add archivo.txt
```

## Crear commit

```bash
./minigit commit "Primer commit"
```

## Ver historial

```bash
./minigit log
```

---

# Diagrama de Casos de Uso (30%)

Casos de uso implementados:

* Inicializar repositorio
* Agregar archivos
* Realizar commits
* Visualizar historial

---

# Flujo de Trabajo Git

El proyecto utiliza:

* Rama `main` para versión estable
* Rama `develop` para desarrollo y pruebas

---

# Integrantes y Avances

| Integrante   | Avance                  |
| ------------ | ----------------------- |
| Integrante 1 | Análisis y BPMN         |
| Integrante 2 | Modelo entidad-relación |
| Integrante 3 | Implementación en Go    |
| Integrante 4 | Documentación y GitHub  |

---

# Repositorio

Repositorio oficial del proyecto:
https://github.com/Last-kHn/minigit

---

# Próximas Mejoras

* Manejo de ramas internas
* Validación de cambios
* Comparación de versiones
* Optimización de commits
* Mejor manejo de archivos
