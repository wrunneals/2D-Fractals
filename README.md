# 2D Fractals

A fractal rendering program that is written in GO. Purpose of this project is to make a multi-threaded renderer that can make high resolution images for making wallpapers/posters/etc.

<img src="/assets/Image1.png" alt="example" width=400 /> <br>

### Fractal Formulas

Inside the function **iteratePixel** is where the main iteration is found that defines the fractal
```
for ; abs(z) < 2 && i < maxIter; i ++{
        z = z * z + c
}
```
This formula can be changed to produce different fractal images. \
For instance the power 4 Mandelbrot can be expressed as ```z = z * z * z * z + c```.

Other variations: 
- The Burning Ship Fractal ```z = complex(math.abs(real(z)), math.abs(imag(z))) + c ```

### Palette Colors

Inside the function **getPaletteColor** there is an array of colors which is the color palette. This is a range of colors that we will lerp through based on the fractal escape count. Planning to have an additional tool to help with palette creation/testing.