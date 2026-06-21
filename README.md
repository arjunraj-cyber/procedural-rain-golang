Procedural Rain Visualizer (Go) 🌧️

A lightweight canvas generator written in Go that procedurally builds a dark, stormy sky layout and saves the output as a PNG graphic. 

This project explores image manipulation at the raw pixel level and implements a concurrent worker pool pattern to handle rendering.

Project Structure & Architecture 🛠️

The program's execution flow is broken down into structured processing steps:
Canvas Initialization: Allocates a custom `1080x1350` (4:5 aspect ratio) RGBA image matrix initialized with a dark, slate-blue background color.
Concurrent Rendering (Worker Pool): Spawns 4 background worker goroutines that concurrently process coordinate inputs from a buffered channel, plotting 800 individual raindrops pixel-by-pixel.
Text Overlay Processing: Utilizes the extended `font.Drawer` interface to map fixed-point fractions and render dynamic typography directly over the pixel grid.
4. Encoding & Export: Standardizes resource stream closure with `defer` and encodes the finished data matrix into a local `rainy_day.png` file.

Current Project Status: Environment Setup
The codebase logic is structurally complete. Running the project locally requires fetching the official extended Go graphics sub-modules to resolve local compiler dependencies:

```bash
go mod init rain-visualizer
go get golang.org/x/image
go run Learning.go

Core Dependencies
golang.org/x/image/font
golang.org/x/image/font/basicfont
golang.org/x/image/math/fixed
