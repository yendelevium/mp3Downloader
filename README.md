---

# MP3 Downloader  
**A web scraper that downloads MP3 tracks locally.**  

⚠ **Disclaimer:**  
This project is intended for **educational purposes only**. While it can access copyrighted material, it is not meant for any illegal or unethical use. The creator does not condone or encourage using this tool to download copyrighted material.  

---

## Features  
- Scrapes the website [hydr0.org](https://hydr0.org/) to locate and download MP3 files.  
- Supports single and simultaneous downloads using Go routines.  
- Customizable download location through `.env` configuration.  

---

## Installation  

### 1. Clone the repository  
```bash
git clone https://github.com/yendelevium/mp3Downloader.git
cd mp3Downloader
```  

### 2. Optional: Set up a `.env` file  

#### Create the `.env` file  
For Linux/macOS:  
```bash
touch .env
```  
For Windows:  
```bash
type nul > .env
```  

#### Configure the `.env` file  
- **Windows:**  
  ```plaintext
  downloadLocationWindows = "<your download location>"
  ```  
- **Unix/Linux:**  
  ```plaintext
  downloadLocationUnix = "<your download location>"
  ```  
If you're using both WSL and Windows, include both configurations.  

If no `.env` file is provided, songs will be downloaded to the default path:  
`cmd/mp3Downloader/`.  

### 3. Install dependencies  
```bash
go mod vendor
```  

---

## Usage  

### 1. Navigate to the project directory containing `main.go`  
For Linux/macOS:  
```bash
cd mp3Downloader/cmd/mp3Downloader
```  
For Windows:  
```bash
cd mp3Downloader\cmd\mp3Downloader\
```  

### 2. Build and run the project  
For Linux/macOS:  
```bash
go build . && ./mp3Downloader
```  
For Windows:  
```bash
go build . ; .\mp3Downloader.exe
```  

### 3. Search and download songs  
- Enter the name of the song, artist, or movie.  
- Select the desired track(s) by entering the corresponding index (`idx`).  

#### Example:  
```plaintext
Enter the name of the song/artist  
Beautiful Mistakes  

0 Track Name: Beautiful Mistakes  
  Artist Name: Franky Perez  

1 Track Name: Beautiful Mistake  
  Artist Name: Lesh  

2 Track Name: Maroon 5 ft. Megan Thee Stallion (Acoustic Cover) (192 kbps)  
  Artist Name: Beautiful Mistakes  

3 Track Name: Beautiful Mistakes  
  Artist Name: Hard EDM Workout  

4 Track Name: Beautiful Mistakes (Leakim Remix)  
  Artist Name: Maroon 5 feat. Megan Thee Stallion  

5 Track Name: Beautiful Mistakes ft. Megan Thee Stallion  
  Artist Name: Maroon 5  

Enter the idx of the song you want to download (if multiple, separate by spaces):  
```  

#### Single download:  
```plaintext
3
```  

#### Simultaneous downloads:  
```plaintext
1 2 5
```  

---

## How It Works  
- The scraper parses the HTML of [hydr0.org](https://hydr0.org/) to identify song information.  
- You select tracks by their index, and the scraper downloads them to the specified directory.  

---

### Upcoming Updates  
- Add a user-friendly interface and deploy it as a web application.  
- Transform it into a feature-rich command-line tool.  

---

## Legal and Ethical Considerations  
This project is not designed for commercial purposes. Always ensure you have the right to download and use the material before proceeding.  

--- 
