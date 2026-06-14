# GoSpotify Request Mechanism

Implementacja projektu Spotify Request Mechanism.

Aplikacja stanowi koncepcję systemu umożliwiającego uczniom zgłaszanie utworów do kolejki Spotify, a następnie automatyczne moderowanie zgłoszeń przed dodaniem ich do odtwarzania. Projekt został opracowany jako propozycja usprawnienia procesu zarządzania muzyką w szkolnym radiowęźle oraz automatyzacji obsługi zgłoszeń utworów.

Projekt został stworzony przez grupę uczniów Technikum Mechaniczno-Elektrycznego im. Nikoli Tesli w Chorzowie i ma charakter edukacyjny oraz demonstracyjny.

> [!NOTE]
> Projekt nie jest wykorzystywany przez Technikum Mechaniczno-Elektryczne im. Nikoli Tesli w Chorzowie ani z nim oficjalnie powiązany. Stanowi niezależną inicjatywę uczniowską prezentującą możliwe rozwiązanie organizacyjne i techniczne dla szkolnego radiowęzła.

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
```
## 🏃 Uruchomienie

Aby uruchomić aplikację lokalnie, wykonaj poniższe polecenie w terminalu:

```bash
go run ./cmd/server
```
## Autorzy

### GoSpotify Request Mechanism

**Alan Nicpoń**  
Go Developer  
Discord: `69jk`

### Spotify Request Mechanism

**Michał Christ**  
Python Developer  
Discord: `merfirdi`

**Piotr Billik**  
Host Manager  
Discord: `lisu5367`

**Łukasz Winkler**  
Web Developer