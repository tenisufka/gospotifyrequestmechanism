# GoSpotify Request Mechanism

Implementacja w języku Go oryginalnego projektu **Spotify Request Mechanism**.

Aplikacja umożliwia uczniom zgłaszanie utworów do kolejki Spotify, a następnie automatycznie moderuje zgłoszenia przed dodaniem ich do odtwarzania. Projekt powstał w celu zautomatyzowania działania szkolnego radiowęzła, a dokładniej procesu dodawania piosenek przez uczniów.

Projekt został stworzony przez grupę uczniów Technikum Mechaniczno-Elektrycznego im. Nikoli Tesli w Chorzowie.

> [!NOTE]
> Aplikacja w żaden sposób nie jest powiązana z Technikum Mechaniczno-Elektrycznym im. Nikoli Tesli w Chorzowie ani jego marką. Jest to niezależny projekt stworzony przez uczniów szkoły.

> [!IMPORTANT]
> Jest to projekt stworzony w Polsce 🇵🇱.

---

## 🚀 Funkcjonalności

* **Logowanie hosta** przez Spotify OAuth
* **Wyszukiwanie utworów** w bazie Spotify
* **Dodawanie utworów** do kolejki odtwarzania
* **Pobieranie tekstów piosenek** z zewnętrznego API LRCLIB
* **Automatyczna moderacja** zgłoszeń
* **Interfejs webowy** oparty o czysty HTML i CSS

### 🛡️ Moderacja utworów
Utwór zostanie automatycznie odrzucony, jeśli:
1. Jest oznaczony jako **Explicit** (wulgaryzmy/nieodpowiednie treści).
2. Trwa dłużej niż **5 minut**.
3. Zawiera słowa znajdujące się na **liście zabronionych wyrażeń**.

---

## 🛠️ Technologie

* **Język:** Go
* **Framework Web:** Gin Gonic
* **Integracje:** * Spotify Web API & Spotify OAuth
    * LRCLIB API

---

## ⚙️ Konfiguracja

Przed uruchomieniem projektu należy przygotować plik konfiguracyjny (np. `.env`) z następującymi zmiennymi środowiskowymi:

```env
SPOTIFY_ID=your_client_id
SPOTIFY_SECRET=your_client_secret
REDIRECT_URL=http://localhost:8080/callback
SESSION_SECRET=change_me

## 🏃 Uruchomienie

Aby uruchomić aplikację lokalnie, wykonaj poniższe polecenie w terminalu:

```bash
go run ./cmd/server

## Autorzy

### Spotify Request Mechanism

**Michał Christ**  
Python Developer  
Discord: `merfirdi`

**Piotr Billik**  
Host Manager  
Discord: `lisu5367`

**Łukasz Winkler**  
Web Developer

### GoSpotify Request Mechanism

**Alan Nicpoń**  
Go Developer  
Discord: `69jk`