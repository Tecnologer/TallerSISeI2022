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
        "token": "blablablabal2765678903"
    },
    {
        "type": "slack",
        "token": "slakc-token-1238632765678903"
    }
]
```
