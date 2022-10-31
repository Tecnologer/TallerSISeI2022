# Factory

Factory carga la lista de servicios desde un archivo JSON y los inicializa para que puedan ser utilizados en la aplicacion.

## Agregar un servicio

El archivo JSON esta formado por un arreglo de objetos, cada objeto es un servicio y cada servicio debe de contener su token y su tipo.

Para agregar un nuevo servicio, basta con crear/modificar el objeto correspondiente a dicho servicio.

### Ejemplo

```json
[
    {
        "type": "telegram",
        "token": "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
    },
    {
        "type": "slack",
        "token": "xoxb-123445677-2353253636236j-akanjndjbkbd7318721"
    }
]
```
