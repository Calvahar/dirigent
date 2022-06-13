# Proftaak - Dirigent Server - Versie 1.1.6

De Dirigent server is verantwoordelijk voor het aansturen van diverse clients. Iedere client die verbinding maakt met de Dirigent server - met de juiste WebSocket upgrade headers -, wordt in een pool gezet en krijgt vervolgens commando's toegestuurd via een WebSocket verbinding.

## Aan de slag

Deze instructies dienen als hulpmiddel voor het opzetten van een omgeving voor de Dirigent server. Deze omgeving kan gebruikt worden voor ontwikkeling en testing. Tevens dient deze README ter documentatie van de ontwikkeling en voortgang van de Dirigent server [van Proftaakgroep B, groep 4].

### Vereisten

Om de Dirigent server werkend te krijgen, zijn enkele vereisten: 
- [Go](https://go.dev/doc/install)
- [Git](https://git-scm.com/downloads)

### Installatie

Clone de repositorie en `cd` naar de directory;
```
git clone https://github.com/Proftaak-Semester-2/dirigent.git && cd dirigent
```

Installeer de benodigde dependencies;
```
go install
```

Vervolgens kan de server opgestart worden;
```
go run main.go
```
Terminal output: 
![Terminal Output](https://github.com/Proftaak-Semester-2/dirigent/blob/main/assets/test_output.png?raw=true)

### Configuratie
In de root staat een `.env.example` bestand. Verander de naam van dit bestand naar `.env` en laat het in de root van de dirigent directory staan. In de .env (environmental variables) kunnen een aantal instellingen ingesteld worden.

```shell
# Server
SERVER_BIND=0.0.0.0
SERVER_PORT=3000

# Config
CONNECT=true 
DISCONNECT=true
ERROR=true
```

## Testen
### API
---
- Voor de static endpoint is er een automatische test geschreven. Deze kan uitgevoerd worden met het commando:
```
go test -v main_test.go
```
- Het resultaat van deze test zou er ongeveer als volgt uit moeten zien (tijden kunnen verschillen):
```
=== RUN   Test_Dirigent
--- PASS: Test_Dirigent (0.00s)
PASS
ok      command-line-arguments  0.388s
```

### WebSocket
---
- Deze test kan tot nu toe alleen handmatig worden uitgevoerd door de dirigent server op te starten en naar http://localhost:3000/connect te navigeren in een browser. Wanneer er op de knop wordt geklikt en deze groen wordt, werkt de WebSocket.

## Gemaakt met

- [Go](https://go.dev/) - Basis voor de Dirigent server
- [Fiber](https://gofiber.io/) - API framework voor Go 
- [ikisocket](https://github.com/antoniodipinto/ikisocket) - WebSocket wrapper voor Fiber WebSocket

## Bijdragen

Een Pull-request kan altijd gemaakt worden en zal bekeken worden voordat deze geaccepteerd wordt.
