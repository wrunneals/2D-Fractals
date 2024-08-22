# 2D Fractal Renderer in GO

A work in progress project for generating high resolution 2D fractal images using GO. Aim was to create a program that can efficiently render high resolution images that I can use to make wallpapers or posters while still allowing me the freedom to test new code/fractals without much effort. I decided to use GO for this project for it's powerful concurrency and native image package.

<img src="/res/RedWhite.png" alt="example" width=800 /> <br>

## Getting Started
Program is currently just three files:
- main.go - Main file that handles image IO and calls to renderer.go
- renderer.go - Rendering logic for fractal generation and handles dispatching go workers.
- pallet.go - Coloring logic with some testing code for visualizing pallets.

Inside renderer.go is the function **iteratePixel** where the main iteration is found that defines the fractal
```
for ; abs(z) < 2 && i < maxIter; i ++{
        z = z * z + c
}
```
This formula can be changed to produce different types of fractals (See 'Other Variations').

## Coloring Algorithm

Program currently uses a histogram coloring algorithm to normalize the image coloring at various scales/depths. Escape time coloring can also be used for testing but is not ideal for creating interesting images.

<img src="/res/sunbrot.png" alt="example" width=800 /> <br>

## Other variations: 
#### The Burning Ship Fractal 
Formula:
```
z = complex(math.Abs(real(z)), math.Abs(imag(z)))
z = z * z + c 
```
<img src="/res/Image2.png" alt="example" width=800 /> <br>
