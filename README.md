# api-blog

El API provee la gestión completa del blog del gimnasio, incluyendo categorías, noticias, secciones, nutrición, dieta, comentarios y respuestas.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
- Golang
- Gin Framework
- PostgreSQL
- Git

### Variables de Entorno

| Variable | Descripción |
|----------|-------------|
| `DB_HOST` | Dirección del servidor de base de datos |
| `DB_PORT` | Puerto de conexión con la base de datos |
| `DB_USER` | Usuario con acceso a la base de datos |
| `DB_PASS` | Password del usuario |
| `DB_NAME` | Nombre de la base de datos |

**NOTA:** Las variables se deben configurar en un archivo `.env` en la raíz del proyecto.

## Estructura del Proyecto
api-blog/
├── config/
│ └── db.go # Configuración y conexión a BD
├── controllers/
│ ├── categoria_controller.go
│ ├── noticia_controller.go
│ ├── articulo_seccion_controller.go
│ ├── nutricion_controller.go
│ ├── dieta_comida_controller.go
│ ├── comentario_comunidad_controller.go
│ └── respuesta_comentario_controller.go
├── models/
│ ├── categoria.go
│ ├── noticia.go
│ ├── articulo_seccion.go
│ ├── nutricion.go
│ ├── dieta_comida.go
│ ├── comentario_comunidad.go
│ └── respuesta_comentario.go
├── routes/
│ ├── categoria_routes.go
│ ├── noticia_routes.go
│ ├── articulo_seccion_routes.go
│ ├── nutricion_routes.go
│ ├── dieta_comida_routes.go
│ ├── comentario_comunidad_routes.go
│ └── respuesta_comentario_routes.go
├── main.go
├── go.mod
├── go.sum
├── .env
└── .gitignore

text

## Endpoints de la API

### Categorías (`/api/v1/categorias`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/categorias` | Obtener todas las categorías |
| GET | `/api/v1/categorias/{id}` | Obtener una categoría por ID |
| POST | `/api/v1/categorias` | Crear una nueva categoría |
| PUT | `/api/v1/categorias/{id}` | Actualizar una categoría |
| DELETE | `/api/v1/categorias/{id}` | Eliminar una categoría |

### Noticias (`/api/v1/noticias`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/noticias` | Obtener todas las noticias |
| GET | `/api/v1/noticias/{id}` | Obtener una noticia por ID |
| POST | `/api/v1/noticias` | Crear una nueva noticia |
| PUT | `/api/v1/noticias/{id}` | Actualizar una noticia |
| DELETE | `/api/v1/noticias/{id}` | Eliminar una noticia |

### Artículos Secciones (`/api/v1/articulos-secciones`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/articulos-secciones` | Obtener todas las secciones |
| GET | `/api/v1/articulos-secciones/{id}` | Obtener una sección por ID |
| POST | `/api/v1/articulos-secciones` | Crear una nueva sección |
| PUT | `/api/v1/articulos-secciones/{id}` | Actualizar una sección |
| DELETE | `/api/v1/articulos-secciones/{id}` | Eliminar una sección |

### Nutrición (`/api/v1/nutricion`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/nutricion` | Obtener todos los planes de nutrición |
| GET | `/api/v1/nutricion/{id}` | Obtener un plan por ID |
| POST | `/api/v1/nutricion` | Crear un nuevo plan de nutrición |
| PUT | `/api/v1/nutricion/{id}` | Actualizar un plan de nutrición |
| DELETE | `/api/v1/nutricion/{id}` | Eliminar un plan de nutrición |

### Dieta Comidas (`/api/v1/dieta-comidas`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/dieta-comidas` | Obtener todas las comidas de dieta |
| GET | `/api/v1/dieta-comidas/{id}` | Obtener una comida por ID |
| POST | `/api/v1/dieta-comidas` | Crear una nueva comida |
| PUT | `/api/v1/dieta-comidas/{id}` | Actualizar una comida |
| DELETE | `/api/v1/dieta-comidas/{id}` | Eliminar una comida |

### Comentarios Comunidad (`/api/v1/comentarios-comunidad`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/comentarios-comunidad` | Obtener todos los comentarios |
| GET | `/api/v1/comentarios-comunidad/{id}` | Obtener un comentario por ID |
| POST | `/api/v1/comentarios-comunidad` | Crear un nuevo comentario |
| PUT | `/api/v1/comentarios-comunidad/{id}` | Actualizar un comentario |
| DELETE | `/api/v1/comentarios-comunidad/{id}` | Eliminar un comentario |

### Respuestas Comentario (`/api/v1/respuestas-comentario`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/respuestas-comentario` | Obtener todas las respuestas |
| GET | `/api/v1/respuestas-comentario/{id}` | Obtener una respuesta por ID |
| POST | `/api/v1/respuestas-comentario` | Crear una nueva respuesta |
| PUT | `/api/v1/respuestas-comentario/{id}` | Actualizar una respuesta |
| DELETE | `/api/v1/respuestas-comentario/{id}` | Eliminar una respuesta |

## Ejecución del Proyecto

### 1. Clonar el repositorio

```bash
git clone https://github.com/cristiangaitan17/api-blog.git
cd api-blog
2. Configurar variables de entorno
Crear un archivo .env en la raíz del proyecto:

env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=tu_contraseña
DB_NAME=gimnasio_db
3. Instalar dependencias
bash
go mod tidy
4. Ejecutar el proyecto
bash
go run main.go
El servidor correrá en http://localhost:8080

Ejemplo de petición POST
bash
curl -X POST http://localhost:8080/api/v1/categorias \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Nutrición Deportiva",
    "seccion_lugar": "blog_principal",
    "descripcion": "Artículos sobre alimentación para deportistas",
    "activo": true
  }'
Dependencias
Paquete	Propósito
github.com/gin-gonic/gin	Framework web
github.com/lib/pq	Driver para PostgreSQL
github.com/joho/godotenv	Manejo de variables de entorno
Estado del Proyecto
✅ Completo - API con CRUD completo para las 7 tablas:

categorias

noticias

articulos_secciones

nutricion

dieta_comidas

comentarios_comunidad

respuestas_comentario

Autor
Cristian Gaitán
