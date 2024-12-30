
# Interviews Service

**Interviews Service** es un servicio desarrollado en Go que permite gestionar entrevistas para candidatos y empleos.
Incluye endpoints para crear entrevistas, listar todas las entrevistas y verificar el estado del servicio.

## Estructura del Proyecto

```plaintext
interviews-service/
├── .github/              # Configuración para GitHub Actions
├── cmd/
│   └── main.go           # Punto de entrada principal de la aplicación
├── docs/                 # Documentación Swagger generada
├── internal/
│   ├── domain/           # Definiciones de modelos y estructuras
│   ├── repository/       # Interacción con la base de datos
│   ├── service/          # Lógica de negocio
│   └── transport/        # Handlers de HTTP (controladores)
├── pkg/
│   ├── config/           # Configuración de la aplicación
│   ├── db/               # Conexión a la base de datos
│   ├── logger/           # Configuración de logging
│   └── utils/            # Funciones utilitarias
├── .env                  # Variables de entorno (no incluir en producción)
├── .gitignore            # Archivos y carpetas ignoradas por Git
├── coverage.out          # Archivo de cobertura de pruebas
├── Dockerfile            # Archivo para construir la imagen de Docker
├── go.mod                # Gestión de dependencias de Go
├── Makefile              # Tareas comunes (compilación, pruebas, Swagger, etc.)
```

---

## Requisitos Previos

- **Go** 1.20 o superior
- **Docker** (opcional, para contenedores)
- **MySQL** como base de datos

---

## Cómo Probar en Local

### 1. Clonar el repositorio

```bash
git clone https://github.com/poolcamacho/interviews-service.git
cd interviews-service
```

### 2. Configurar las variables de entorno

Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
DATABASE_URL=admin_db:password@tcp(localhost:3306)/talent_management_db
JWT_SECRET_KEY=tu-secreto-jwt
PORT=3000
```

### 3. Ejecutar la aplicación localmente

Ejecuta el siguiente comando para iniciar el servicio:

```bash
make run
```

Accede al servicio en `http://localhost:3000`.

### 4. Generar la documentación Swagger

```bash
make swagger
```

La documentación estará disponible en `http://localhost:3000/swagger/index.html`.

### 5. Ejecutar pruebas

```bash
make test
```

Esto generará un archivo `coverage.out` con la cobertura de pruebas.

---

## Uso de Docker

### 1. Construir la imagen de Docker

```bash
docker build -t interviews-service .
```

### 2. Ejecutar el contenedor

```bash
docker run -d --name interviews-service   -e DATABASE_URL=admin_db:password@tcp(localhost:3306)/auth_service_db   -e JWT_SECRET_KEY=tu-secreto-jwt   -p 3000:3000 interviews-service
```

El servicio estará disponible en `http://localhost:3000`.

---

## Endpoints

### 1. **Health Check**

**Descripción**: Verifica el estado del servicio.

**Endpoint**: `GET /health`

**Ejemplo de respuesta**:

```json
{
  "status": "healthy"
}
```

---

### 2. **Registro de Entrevista**

**Descripción**: Crea una nueva entrevista.

**Endpoint**: `POST /interviews`

**Cuerpo de la Solicitud**:

```json
{
  "candidate_id": 101,
  "job_id": 202,
  "interview_date": "2024-12-30T15:00:00Z",
  "feedback": "Good technical skills."
}
```

**Ejemplo de Respuesta Exitosa**:

```json
{
  "message": "interview created successfully"
}
```

**Ejemplo de Respuesta de Error**:

```json
{
  "error": "candidate_id, job_id, and interview_date are required"
}
```

---

### 3. **Listar Entrevistas**

**Descripción**: Obtiene una lista de todas las entrevistas.

**Endpoint**: `GET /interviews`

**Ejemplo de Respuesta Exitosa**:

```json
[
  {
    "id": 1,
    "candidate_id": 101,
    "job_id": 202,
    "interview_date": "2024-12-30T15:00:00Z",
    "feedback": "Good technical skills."
  }
]
```
---

