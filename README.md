# Lab_6_web
LaLiga Tracker - Backend
Este es el backend para LaLiga Tracker, una API REST para gestionar partidos de La Liga. Fue desarrollado utilizando Go y se ejecuta dentro de un contenedor Docker.

📌 Características
API REST para gestionar información sobre partidos de La Liga.

Implementación en Go.

Contenerización con Docker para fácil despliegue.

🛠️ Requisitos
Asegúrate de tener instalado:

Docker
## Métodos y Endpoints

### 1️⃣ **GET /api/matches
   **Descripción**: Obtiene todos los partidos.

### 2️⃣ **GET /api/matches/:id
   **Descripción**: Obtiene un partido específico por su ID.
   - **Parámetros**: 
     - `:id` - ID del partido.

### 3️⃣ **POST /api/matches
   **Descripción**: Crea un nuevo partido.
   - **Cuerpo de la solicitud: Información del partido a crear.

### 4️⃣ **PUT /api/matches/:id
   **Descripción**: Actualiza un partido existente.
   - **Parámetros**:
     - `:id` - ID del partido.
   - **Cuerpo de la solicitud: Datos actualizados del partido.

### 5️⃣ **DELETE /api/matches/:id
   **Descripción**: Elimina un partido específico por ID.
   - **Parámetros**:
     - `:id` - ID del partido.

### 6️⃣ **PATCH /api/matches/:id/goals
   **Descripción**: Registra un gol en el partido.
   - **Parámetros**:
     - `:id` - ID del partido.
   - **Cuerpo de la solicitud**: Información del gol (jugador, minuto, etc.).

### 7️⃣ **PATCH /api/matches/:id/yellowcards
   **Descripción**: Registra una tarjeta amarilla en el partido.
   - **Parámetros**:
     - `:id` - ID del partido.
   - **Cuerpo de la solicitud**: Información del jugador y minuto de la tarjeta amarilla.

### 8️⃣ **PATCH /api/matches/:id/redcards
   **Descripción**: Registra una tarjeta roja en el partido.
   - **Parámetros**:
     - `:id` - ID del partido.
   - **Cuerpo de la solicitud**: Información del jugador y minuto de la tarjeta roja.

### 9️⃣ **PATCH /api/matches/:id/extratime
   **Descripción**: Registra el tiempo extra del partido.
   - **Parámetros**:
     - `:id` - ID del partido.
   - **Cuerpo de la solicitud**: Tiempo adicional del partido.

# 🚀 Instrucciones para ejecutar el backend

1️⃣ **Clonar el repositorio**  
   Clona el repositorio y entra en el directorio del proyecto:
   git clone https://github.com/tu-usuario/tu-repositorio.git
   cd tu-repositorio
2️⃣ Construir la imagen de Docker

  Ejecuta el siguiente comando en la raíz del proyecto para crear la imagen del backend:
docker build -t laliga-backend .

## 3️⃣ Ejecutar el contenedor

Para correr el contenedor en segundo plano y exponer el puerto 8080, ejecuta el siguiente comando:

docker run -d -p 8080:8080 --name laliga-container laliga-backend
  

Evidencia del funcionamineto
![crear partido](https://github.com/user-attachments/assets/766cb4ff-93f0-4aff-8ba2-73ba5a72410a)
![1](https://github.com/user-attachments/assets/1310c7f9-cc5e-4438-890a-18d48d12c7cf)
![4](https://github.com/user-attachments/assets/feb2a678-1a9f-4904-98a1-935998cdca7d)
![3](https://github.com/user-attachments/assets/64036e20-544e-4fef-ad70-234977a0543d)
![2](https://github.com/user-attachments/assets/dc73cd88-9bf2-46a3-87eb-97edc3640361)


 ```bash




