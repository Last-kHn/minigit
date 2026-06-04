# MiniGit

## Descripción

MiniGit es una herramienta ligera de control de versiones desarrollada en Go con fines educativos. El proyecto busca implementar los conceptos fundamentales de sistemas como Git mediante una arquitectura simple basada en terminal.

---

## Estado del Proyecto

**Avance actual: 40%**

### Funcionalidades implementadas

* [x] Inicialización de repositorios (`init`)
* [x] Área de staging (`add`)
* [x] Creación de commits (`commit`)
* [x] Historial de commits (`log`)
* [x] Consulta de estado (`status`)
* [x] Validación de commits vacíos
* [x] Limpieza automática del staging

### Funcionalidades pendientes

* [ ] Gestión de ramas internas
* [ ] Cambio entre ramas
* [ ] Comparación de versiones
* [ ] Restauración de commits
* [ ] Optimización de almacenamiento

---

## Requisitos

* Go 1.22+
* Linux o Windows

---

## Compilación

```bash
go build -o minigit
```

---

## Ejecución

```bash
./minigit
```

---

## Comandos Disponibles

### Inicializar repositorio

```bash
./minigit init
```

### Agregar archivo

```bash
./minigit add archivo.txt
```

### Consultar estado

```bash
./minigit status
```

### Crear commit

```bash
./minigit commit "Primer commit"
```

### Ver historial

```bash
./minigit log
```

---

## Estructura del Proyecto

```text
Minigit/
├── main.go
├── repository.go
├── staging.go
├── commit.go
├── status.go
└── go.mod
```

---

## Flujo de Trabajo

```text
init
 ↓
add
 ↓
status
 ↓
commit
 ↓
log
```

---

## Rama de Desarrollo

El proyecto utiliza dos ramas principales:

* `main`: versión estable.
* `develop`: desarrollo activo e integración de nuevas funcionalidades.

---

## Próximas Mejoras

Durante el siguiente avance se implementarán mecanismos de gestión de ramas y mejoras en el sistema de control de versiones para acercar el comportamiento de MiniGit a herramientas profesionales.
