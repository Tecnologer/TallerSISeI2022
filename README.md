# Taller SISeI 2022 - Dos Bots, Una Aplicacion

El objetivo de este taller es familiarizarse con el uso de API de mensajarias, en este ejemplo estaremos usando telegram y slack.

La aplicacion sera capas de implementar ambas APIs para permitirle al estudiante comunicarse con ambas mensajerias al mismo tiempo, es decir, un mensaje que se envie desde telegram podra ser recibido en slack y viceversa.

## Requisitos

1. [Go 1.11 o superior][2]
1. Editor de texto de su preferencia puede ser V[isual Studio Code][3] o si prefiere un IDE puede usar [GoLand][4].
1. [Git][5] instalado en el equipo
1. Cuenta de [github][6]

## Configuracion

### Go

Para mejor comodidad puede configurar la variable de entorno ([Windows][9], [Linux][10]) [GOPATH][1]. Esta variable le permitira a Go identificar en que folder trabajar, dicho folder sera donde se descarguen las dependencias.

### Git

Configurar el nombre de usuario y correo con el que se crearan los commits.

1. configurar [nombre de usuario][7]: `git config --global user.name "Mona Lisa"`
1. configurar [correo electronico][8]: `git config --global user.email "mona@lisa.com"`

> La bandera `--global` indica que dicha configuracion se hara para todos los repositorios.

>NOTA: En caso de obtener un error de que `git` (o `go`) no se reconoce como un comando, hay dos posibles causas:
>1. El directorio de instalacion no esta en la variable de entorno `$PATH`.
>2. Si si esta en `$PATH` solo tienes que reiniciar el editor o la terminal para que la referencia se actualice.


## Clonar repositorio

Para evitar problemas con referencias de paquetes, recomiendo clonarlo en `$GOPATH/src/github.com/tecnologer/TallerSISeI2022`.

Ejecutando el comando:

```bash
git clone https://github.com/Tecnologer/TallerSISeI2022.git $GOPATH/src/github.com/tecnologer/TallerSISeI2022
```

`$GOPATH` es una variable de entorno, si estas en windows el formato puede cambiar, es decir, para CMD `%GOPATH%` mientras que para PowerShell seria `$Env:GOPATH`.

Una vez clonado, hay que asegurarse de tener todas las dependencias del proyecto, para esto habra que ejecutar lo siguiente:

```bash
cd $GOPATH/src/github.com/tecnologer/TallerSISeI2022
go mod tidy
```

[1]: https://www.digitalocean.com/community/tutorials/understanding-the-gopath-es
[2]: https://go.dev/dl/
[3]: https://code.visualstudio.com/
[4]: https://www.jetbrains.com/go/
[5]: https://git-scm.com/downloads
[6]: https://github.com
[7]: https://docs.github.com/es/get-started/getting-started-with-git/setting-your-username-in-git?platform=windows
[8]: https://docs.github.com/es/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-email-preferences/setting-your-commit-email-address
[9]: https://geekflare.com/es/system-environment-variables-in-windows/
[10]: https://tecnofaq.com/como-configuro-las-variables-de-entorno-en-ubuntu/