[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/cBV3pX2A)
# Web Technologien // begleitendes Projekt Sommersemester 2025
Zum Modul Web Technologien gibt es ein begleitendes Projekt. Im Rahmen dieses Projekts werden wir von Veranstaltung zu Veranstaltung ein Projekt sukzessive weiter entwickeln und uns im Rahmen der Veranstaltung den Fortschritt anschauen, Code Reviews machen und Entwicklungsschritte vorstellen und diskutieren.

Als organisatorischen Rahmen für das Projekt nutzen wir GitHub Classroom. Inhaltlich befassen wir uns mit einer  Client-Server Anwendung mit deren Hilfe IPTC Metadaten von Bildern schnell und einfach gepflegt werden können.

Dokumentieren Sie in diesem Beibootprojekt Ihre Entscheidungen gewissenhaft unter Zuhilfenahme von [Architectual Decision Records](https://adr.github.io) (ADR).

Hier ein paar ADR Beispiele aus dem letzten Semestern:
- https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2022-Moosgloeckchen/tree/main/docs/decisions
- https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2022-mweiershaeuser/tree/main/adr
- https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2022-twobiers/tree/main/adr

Halten Sie die Anwendung, gerade in der Anfangsphase möglichst einfach, schlank und leichtgewichtig (KISS).

___

## Running the project

### Backend (using Go)
1. Install go (ver. 1.24)
2. Go to "backend" folder
3. Install dependencies
   ```bash
   go mod download
   ```
4. Run the server
   ```bash
    go run main.go
    ```
   
### Using Docker
1. Install Docker
2. Execute the following command in the root directory of the project:
   ```bash
   docker-compose up
   ```