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

# IPTC Editor
A simple web application to view and edit IPTC metadata in images, written in Go (backend) and Svelte + TypeScript (frontend).

## Structure

### Backend (`/backend`)
- **Language:** Go, Gin Framework
- **Purpose:** Provides a REST API for reading and editing IPTC metadata in images. Handles image file management and metadata persistence (e.g., via SQLite).
- **Key folders:**
  - `assets/`: Where the actual files are stored
  - `db/`: Database models and queries
  - `router/`: API routing and endpoints
  - `migrate/` & `migrations/`: Database migrations
  - `errors/`: Error types

### Frontend (`/frontend`)
- **Technologies:** Svelte, TypeScript, Vite
- **Purpose:** User interface for viewing and editing IPTC metadata. Communicates with the backend via HTTP.
- **Key folders and files:**
  - `src/components/`: UI components (e.g., metadata editor, file list, bulk editor)
  - `src/routes/`: SvelteKit routes for the page structure (unused as there is only one page)
  - `src/utils.ts`: Utility functions (e.g., API calls, date formatting)
  - `static/`: Static public files (i.e., Favicon and IPTC tag definitions)

### Communication
- The frontend communicates with the backend REST API to load and save metadata.
- Metadata is transferred as JSON between frontend and backend.

### Data Model
- Each image has a unique ID and a set of IPTC metadata fields.
- Files and folders are stored in the filesystem as well as the database.
- Metadata is read directly from image files using `exiftool` library.

### Development & Deployment
- **Docker:** Docker containers are provided for local development and deployment.
- **Migrations:** Database schema is versioned using migrations.

## Running the project

### Using Docker (preferred)
1. Install Docker
2. Execute the following command in the root directory of the project:
   ```bash
   docker-compose up
   ```
   or, for just starting the backend server (if you want to run the frontend manually):
    ```bash
    docker-compose up backend
    ```
3. The backend will be available at `http://localhost:8080` and the frontend at `http://localhost:4173`
4. Run migrations:
   ```bash
   docker exec -it <backend-container-name> go run main.go --migrate
   ```

### Manual Setup
#### Backend (using Go)
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
   
#### Frontend (Bun or NPM)
1. Install Bun
2. Go to "frontend" folder
3. Install dependencies
```bash
   bun install
```
4. Copy the `.env.example` file to `.env` and adjust the variables if needed
   ```bash
   cp .env.example .env
   ```
5. Run the frontend
   ```bash
   bun dev
   ```
