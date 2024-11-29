# mp3Downloader
Scrapes the web and downloads the mp3 of your song locally

## Installation
    
### Clone the repository
    git clone https://github.com/yendelevium/mp3Downloader.git
    cd mp3Downloader

## Optional
### Create a .env file in the root directory
#### For Linux/macOS
    touch .env

#### For windows
    type nul > .env
    (Note that you actually have to type "type")

### Populate the .env file
#### For windows
    downloadLocationWindows = "<your download location>"

#### For unix/linux
    downloadLocationUnix = "<your download location>"

Add both if you're planning to use WSL and Windows    
If you don't add the .env files or the download path, the songs will be downloaded by default in the cmd/mp3Downloader directory

#### Install dependencies 
    go mod vendor

## Working
The scraper scrapes the website https://hydr0.org/ to download the songs. It loactes the relevent details from the HTML and extracts it. This project is only for educational purposes, and is NOT TO BE USED FOR COMMERCIAL PURPOSES

### Change Directory to where the main.go file is located
#### For Linux/macOS
    cd mp3Downloader/cmd/mp3Downloader 

#### For windows
    cd mp3Downloader\cmd\mp3Downloader\

### Build and run the project
#### For Linux/macOS
    go build . && ./mp3Downloader

#### For windows
    go build . ; .\mp3Downloader.exe


### Enter the song name/artist name/movie name
    Enter the name of the song/artist

The web scraper will then search the website for the specified track.

### Enter the idx corresponding to the track you want to download
Here's an example. There will be more tracks to choose from when you actually run it
    Enter the name of the song/artist
    Beautiful Mistakes

    0 Track Name: Beautiful Mistakes
    Artist Name:Franky Perez

    1 Track Name: Beautiful Mistake
    Artist Name:Lesh

    2 Track Name: Maroon 5 ft. Megan Thee Stallion (Acoustic Cover) (192 kbps)
    Artist Name:Beautiful Mistakes

    3 Track Name: Beautiful Mistakes
    Artist Name:Hard EDM Workout

    4 Track Name: Beautiful Mistakes (Leakim Remix)
    Artist Name:Maroon 5 feat. Megan Thee Stallion

    5 Track Name: Beautiful Mistakes ft. Megan Thee Stallion
    Artist Name:Maroon 5

    Enter the idx of the song you want to download(if multiple, separate by spaces)

You can either enter one song, eg:

    3

And this will download the track corresponding to that idx

Or multiple, eg

    1 2 5

And all these songs will be downloaded SIMULTANEOUSLY(by using goroutines)

# DISCLAIMER
While this downloader can be used to access copyrighted material, this project is only for educational purposes, and I do not encourage anyone to use it for accessing copyrighted material either.