# Setup

Necesitamos tener instalado `docker` y `docker-compose`
* Ejecutar el comando `docker-compose up`.
* Con esto tendremos creado dos containers, uno el servicio en si y otro para redis

* El servicio corre en el puerto 8080

# Hay 3 rutas disponibles

POST http://localhost:8080/:userId/tweet
- Crea un nuevo tweet para el usuario especificado, con el tweet en el body de la request
```json 
    {
        "content": "new twwet for testing"
    }
```
- Autenticación: Requerida

POST http://localhost:8080/:userId/follow
- Permite al usuario userId seguir a otro usario especificado en el body de la request
```json 
    {
        "userId": "2"
    }
```
- Autenticación: Requerida

GET http://localhost:8080/:userId/timeline
- Permite al usuario userId obtener los tweets de los usarios que sigue
- Autenticación: Requerida
