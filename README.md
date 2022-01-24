# 2D Fractals

A fractal rendering program that is written in GO. Purpose of this project is to make a multi-threaded renderer that can make high resolution images for making wallpapers/posters/etc.

<img src="/res/Image1.png" alt="example" width=400 /> <br>
Image center (-0.761574 - 0.0847596i)

### Fractal Formulas

Inside the function **iteratePixel** is where the main iteration is found that defines the fractal
```
for ; abs(z) < 2 && i < maxIter; i ++{
        z = z * z + c
}
```
This formula can be changed to produce different fractal images. \
For instance the power 4 Mandelbrot can be expressed as ```z = z * z * z * z + c```.

### Coloring Algorithm

Program currently uses a histogram coloring approach to normalize the image. Escape time coloring can also be swapped in fairly easily

<img src="/res/Image1.png" alt="example" width=400 /> <br>

###### Other variations: 
- The Burning Ship Fractal 
```
z = complex(math.Abs(real(z)), math.Abs(imag(z)))
z = z * z + c 
```
<img src="/res/Image2.png" alt="example" width=400 /> <br>
Image center (-1.762 - 0.028i), Scale 0.1