# Proftaak - Dirigent Server - Versie 1.1.1

De Dirigent server is verantwoordelijk voor het aansturen van diverse clients. Iedere client die verbinding maakt met de Dirigent server - met de juiste WebSocket upgrade headers -, wordt in een pool gezet en krijgt vervolgens commando's toegestuurd via een WebSocket verbinding.

## Aan de slag

Deze instructies dienen als hulpmiddel voor het opzetten van een omgeving voor de Dirigent server. Deze omgeving kan gebruikt worden voor ontwikkeling en testing. Tevens dient deze README ter documentatie van de ontwikkeling en voortgang van de Dirigent server [van Proftaakgroep B, groep 4].

### Vereisten

Om de Dirigent server werkend te krijgen, zijn enkele vereisten: 
- [Go](https://go.dev/doc/install)
- [Git](https://git-scm.com/downloads)

### Installatie

Clone de repositorie en `cd` naar de directory
```
git clone https://github.com/Proftaak-Semester-2/dirigent.git && cd dirigent
```

Installeer de benodigde dependencies
```
go install
```

Vervolgens kan je de server opstarten
```
go run main.go && go run colors.go
```
Terminal output: 
![Terminal Output](https://github.com/Proftaak-Semester-2/dirigent/blob/main/assets/test_output.png?raw=true)

### Configuratie
In de root staat een `.env` bestand. Hierin kunnen een aantal instellingen ingesteld worden.

```shell
# Server
SERVER_PORT=3000
```

## Test uitvoeren

Op dit moment zijn er nog geen automated tests. Er kan echter wel een handmatige test uitgevoerd worden.

### Voorbeeld tests
Test de API

- Navigeer in een browser naar http://127.0.0.1:3000/test
- De tekst "*Test endpoint werkend!*" wordt zichtbaar in de browser. Daarnaast wordt er in de terminal informatie over de connectie geschreven.

Test de WebSocket

- Navigeer naar de [client](/tests/client.html) en open deze in een browser.

## Gemaakt met

- [Go](https://go.dev/) - Basis voor de Dirigent server
- [Fiber](https://gofiber.io/) - API framework voor Go 
- [ikisocket](https://github.com/antoniodipinto/ikisocket) - WebSocket wrapper voor Fiber

## Bijdragen

Een Pull-request kan altijd gemaakt worden en zal bekeken worden voordat deze geaccepteerd wordt.
