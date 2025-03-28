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

🚀 Instrucciones para ejecutar el backend
1️⃣ Clonar el repositorio
git clone https://github.com/tu-usuario/tu-repositorio.git
cd tu-repositorio
2️⃣ Construir la imagen de Docker
Ejecuta el siguiente comando en la raíz del proyecto para crear la imagen del backend:
docker build -t laliga-backend .

3️⃣ Ejecutar el contenedor
Para correr el contenedor en segundo plano y exponer el puerto 8080:

docker run -d -p 8080:8080 --name laliga-container laliga-backend

4️⃣ Verificar que el contenedor está corriendo
docker ps

Evidencia del funcionamineto
![crear partido](https://github.com/user-attachments/assets/766cb4ff-93f0-4aff-8ba2-73ba5a72410a)
![1](https://github.com/user-attachments/assets/1310c7f9-cc5e-4438-890a-18d48d12c7cf)
![4](https://github.com/user-attachments/assets/feb2a678-1a9f-4904-98a1-935998cdca7d)
![3](https://github.com/user-attachments/assets/64036e20-544e-4fef-ad70-234977a0543d)
![2](https://github.com/user-attachments/assets/dc73cd88-9bf2-46a3-87eb-97edc3640361)







