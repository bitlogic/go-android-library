# go-android-library

Package de ejemplo para generar una librería que pueda integrarse con Android.

### 💻 Ambiente de desarrollo necesario
Para configurar el ambiente, se utilizaron dos imágenes de Docker distintas, una que provee los SDK de Android y el lenguaje Go en sí ([ci/Dockerfile.android](https://github.com/bitlogic/go-android-library/blob/main/ci/Dockerfile.android)), y otra que toma esta configuración y ajusta las variables de entorno necesarias para la compilación ([ci/Dockerfile.library](https://github.com/bitlogic/go-android-library/blob/main/ci/Dockerfile.library)).
Originalmente, la primera imagen se tomó de [aquí](https://github.com/paulaolmedo/go4droid "go4droid"), y se adaptó a las necesidades del desarrollo (por ej., el Dockerfile aquí visto, no contiene gradle)

Si se quiere desplegar de manera automática este ambiente de desarrollo, el usuario debe contar con VSCode (y Docker desktop). Al abrir este proyecto en dicho IDE, éste detectará la presencia de la carpeta [.devcontainer](https://github.com/bitlogic/go-android-library/tree/main/.devcontainer), en la cual se encuentra un archivo de configuración **.json**, que permite iniciar un contenedor con los Dockerfiles mencionados.

#### IMPORTANTE ❗
Además, si finalmente se quiere generar una imagen de 🐳 que contenga solamente el paquete generado, en ([ci/Dockerfile.library](https://github.com/bitlogic/go-android-library/blob/main/ci/Dockerfile.library)) se especifica como hacerlo (ver líneas comentadas!)

### 📦 Contenido del package 
La carpeta [library](https://github.com/bitlogic/go-android-library/tree/main/library) contiene todo el código necesario. Se notará que se encuentra todo en una sóla carpeta, y _no_ porque sea un ejemplo, si no porque actualmente _gomobile_ no soporta del todo la importación de distintos paquetes (ver [aquí](https://github.com/golang/go/issues/39735) el issue asociado).

En cuanto al código en sí, existen un par de aclaraciones que es necesario realizar:

* **_Tipos_**: _gomobile_ admite sólo [éstos](https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions)
* **_Parámetros_**: se deben pasar por referencia.
Es decir que si normalmente se tendría una función:

```golang
func AddUser(userInfo User){
...
}
```

Para poder tener su "equivalente" en la librería, deberá escribirse como:
```golang
func AddUser(userInfo *User){
...
}
```
* **_Retornos_**: aplica la misma condición que para los parámetros. Además, es muy imporante tener en cuenta que *no* permite el retorno de _slices_ de _structs_. Para solucionar esto, *sí* admite el retorno de _slices_ de _bytes_, como puede observarse en [mobile.go](https://github.com/bitlogic/go-android-library/blob/bea5b5c78e707497b080200b2c6017e89bfb3f17/library/mobile.go#L29)
* **_Uso general_**: si bien el compilador genera setters y getters para _structs_, no lo hará para _structs_ de _structs_, como es el caso de [model.go](https://github.com/bitlogic/go-android-library/blob/main/library/model.go). Es por esto, que para poder manipular este tipo de objeto, se deberá agregar un método adicional para generar dicha asociación.

