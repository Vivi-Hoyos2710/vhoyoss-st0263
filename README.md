# info de la materia: ST0263-TOPICOS ESPECIALES EN TELEMATICA

# Estudiante(s): Viviana Hoyos Sierra, vhoyoss@eafit.edu.co

# Profesor: Edwin Nelson Montoya Munera, emontoya@eafit.edu.co


# Nombre del proyecto: Reto 1 y 2

# 1. Breve descripción de la actividad

Red P2P No estructurada basada en servidor central con microservicios e implementacion de comunicación con gRPC y REST.

## 1.1. Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)
### Requisitos funcionales

- Autenticación de peers (login, logout, tokens jwt)
- Consulta de recursos (query, sendIndex)
- Servicio de transferencia de archivos (download, upload)
- Servicio de cambio de moneda (currency exchange)
- Rotación de las responsabilidades (Round Robin).
- Servicio de Bootstrap o inicialización de los nuevos peers a la red P2P (Archivos setup.py y los .env).
- Implementación de comunicación basada en gRPC y API REST.
- Uso de Docker para la creacion y despliegue de contenedores.
- Despliegue del reto en AWS Academy.

### Requisitos no funcionales

- Mantenibilidad
- Disponibilidad
- Escalabilidad
- Rendimiento

## 1.2. Que aspectos NO cumplió o desarrolló de la actividad propuesta por el profesor (requerimientos funcionales y no funcionales)

No se hace la transferencia periódica del índice de archivos de manera independiente, sino que se transfiere cada que un peer realiza alguna acción dentro de la red, es decir, cuando hay alguna actualización.
No se implementó MOM (El profesor recomendo no utilizarlo)


# 2. información general de diseño de alto nivel, arquitectura, patrones, mejores prácticas utilizadas.
### Diagrama de arquitectura

![ArquitecturaP2P](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/79054e84-29a8-4943-a08d-90bb0e7be5dd)

### Patrones de diseño

- DRY
- KISS

### Mejores Prácticas

- Distribución modular: El sistema se organizó de tal manera que todos los servicios y módulos de clientes y server grpc, tuvieran su propio archivo y carpeta, para que sea más fácil localizar los archivos a la hora de hacer cambios.
- Escalabilidad horizontal: La estructura del codigo y contenedores nos permite agregar muchos mas peers al sistema de manera facil sin consumir demasiados recursos.
- Gestión de configuraciones: No se quemaron variables como direcciones ips o puertos directamente en el codigo sino que se utilizaron archivos como los .env para gestionar las variables de manera correcta.
- Tolerancia a fallos: Se implementaron metodos para el manejo de errores dentro del fujo del programa.

# 3. Descripción del ambiente de desarrollo y técnico: lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

## Lenguajes de programación

Se utilizaron Python y Go. el primero para la construcción de los peers y el segundo para el servidor central.

## Principales librerias utilizadas y sus versiones

- grpcio==1.62.0
- grpcio-tools==1.62.0
- protobuf==4.25.3
- python-dotenv==1.0.1
- requests==2.28.1

## Como se compila y ejecuta.

Primero se deben crear las instancias correpondientes en AWS, una para el central-server y luego una mas para cada peer que se desee utilizar. se recomientda implementar el uso de ip's elásticas para evitar problemas a la hora de ejeutar el programa.

Una vez ya creadas, se puede utilizar el terminal directamente desde la pagina o se puede acceder a ellas remotamente mediante SSH.
Posteriormente se debe clona el repositorio con el proyecto en cada una de las instancias.
Tambien sera necesario instalar docker para luego poder correr los contenedores, este tutorial podria resultarle util en caso de que necesite ayuda:

`https://docs.docker.com/compose/install/linux/`

Luego se deberan realizar algunos cambios en el archivo _docker-compose.yml_ de las instancias de los peers, dado el caso de que se desee cambiar el puerto en el que se ejecutaran (el valor por defecto es 9000). Tambien se deberan hacer cambios en los archivos .env para espcificar las ip's correspondientes a las instancias en las que se ejecutan el pservidor y el servidor central. Por ultimo, tambien se debera modificar el archivo .env del servidor central, donde se pogra modificar el puerto en caso de que lo desee (el valor por defecto es 8000) y agregar una palabra "SECRET" que se utilizara en la encriptacion.

# Alguno comandos utiles:
Para editar el docker-compose.yml en las instancias de peer1 y peer2

```
vi docker-compose.yml
```

Para editar el archivo .env (en central_server, peer1 y peer2)
```
touch .env
vi .env
```

Luego de crear los archivos de configuración, solo faltaria hacer _build_ para construir los contenedores y luego ejecutando el codigo que le corresponde a cada instancia.

# Comandos para construir y ejecutar los contenedores:
central-server:
```
docker compose build 
docker compose up
```

PServer de peer1 y peer2: 
```
docker compose build
docker compose up pserver
```

PClient de peer1 y peer2:
```
docker compose run -i pclient
```

Una vez ejecutados los comandos ya se estaria ejecutando el servidor central y los pservers y pclients de cada peers, por lo que el programa estaria listo para ser utilizado por el usuario. 


## Detalles del desarrollo.

El servidor central fue desarrollado en Golang y tanto el PCliente como el PServidor fueron desarrollados en Python. Se desplegaron tres instancias EC2 de AWS, una para el servidor central (central-server), y una para peers (peer1 y peer2).

## Detalles técnicos

Endpoints:
- /login
- /logout
- /sendIndex
- /indexTable
- /query
- /getPeerUploading

## Descripción y como se configura los parámetros del proyecto (ej: ip, puertos, conexión a bases de datos, variables de ambiente, parámetros, etc)

Como se menciono anteriormente, para configurar las variables de entorno y setear las ip's y puertos se debe acceder y modificar los archivos _docker-compose.yml_ y  _.env_, En ambos, peers y server. Tambien cabe resaltar que no se utilizaron bases de datos dentro del del proyecto. 


## Opcional - detalles de la organización del código por carpetas o descripción de algún archivo. (ESTRUCTURA DE DIRECTORIOS Y ARCHIVOS IMPORTANTE DEL PROYECTO, comando 'tree' de linux)
# Estructura de directorios
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/a8c34b27-1b3d-4333-a5a1-f75955083187)

# 4. Descripción del ambiente de EJECUCIÓN (en producción) lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

# IP o nombres de dominio en nube o en la máquina servidor.
Se crearon 3 IPs elasticas en aws, una para cada instancia en la que se ejecutaria el servidor o un peer:
Servidor central: 18.213.144.86
Peer 1: 3.227.110.227
Peer 2: 52.201.60.240

## Como se lanza el servidor.
Para ejecutar el servidor dentro de la instancia se deben ejecutar los siguientes comandos en sucesión:
```
git clone https://github.com/DamianDuque/daduquel-st0263.git
cd daduquel-st0263
docker compose build
docker compose up
```

## Una mini guia de como un usuario utilizaría el software o la aplicación
Para una guia sobre como ejecutar el archivo remitase a la sección 3.
Una vez se hayan creado y ejecutado los contenedores en docker el usuario tendra acceso a todos los servicios disponibles en el cliente, a continuación se incluye una imagen del menu principal del CLI:
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/f3407de6-e62f-4cab-8e9e-16f09958e119)
Para utilizar cualquier servicio bastara con que que el usuario ingrese por terminal el numero que le corresponde.
Una vez selecionado el servicio, el programa le preguntara por mas informacion adicional en caso de ser requerido (ej. Archivo a descargar).
Una vez se complete la operacion dentro del servicio el programa volvera a su menu principal en la que el ususario podra elegir su propia operación.

## Opcionalmente - si quiere mostrar resultados o pantallazos 

# Menu CLI
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/06865590-8674-4d37-a746-7dd2a3b23266)

# Query
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/1b583c18-5113-441d-b869-3f0eb026e2c9)

# Download
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/95bc2f0a-d87b-4a78-a3e8-5b804be5af94)
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/f9684ff7-6c48-4b76-9a4e-185e8a78a43e)
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/a3400e60-6dde-4b95-8368-8696e872e525)

# Upload
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/041ba6a9-8aeb-4c1b-b0c5-9c98613aec62)

# Currency Exchange
![image](https://github.com/DamianDuque/daduquel-st0263/assets/94024545/20420acf-3f3f-469a-bfa0-c9ead150ae8b)


# 5. Otra información que considere relevante para esta actividad.

# referencias:
MissCoding. (2022, 9 mayo). Python gRPC Tutorial - Create a gRPC Client and Server in Python with Various Types of gRPC Calls [Vídeo]. YouTube. 
«Install Docker Engine on Ubuntu». (2024, 31 enero). Docker Documentation. 