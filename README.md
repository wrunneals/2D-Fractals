# 2D Fractals
A fractal rendering program that is written in GO. Purpose of this project is to make a multi-threaded renderer that can make high resolution images for making wallpapers/posters/etc. \

<img src="/assets/Image1.png" alt="example" width=400 />

### Formulas
Inside the function **iteratePixel** is where the main iteration is found
```
for ; abs(z) < 2 && i < maxIter; i ++{
        z = z * z + c
}
```
The Mandelbrot iteration is pictured above is running the formula 
<img src="https://latex.codecogs.com/svg.image?z&space;=&space;z^2&space;&plus;&space;c&space;" title="z = z^2 + c " />. 
Though it's possible to change this formula to produce different fractals.

Mandelbrot can be extended to polynomials in the form 
<img src="https://latex.codecogs.com/svg.image?z=az^n&space;&plus;bz^{n-1}...&space;&plus;&space;c" title="z=az^n +bz^{n-1}... + c" /> <br>
For instance the power 8 Mandelbrot would be the formula 
<img src="https://latex.codecogs.com/svg.image?z=z^8&plus;c" title="z=z^8+c" />

The burrning ship fractal is another variation: 
<img src="https://latex.codecogs.com/svg.image?z&space;=&space;(|Real(z)|&space;&plus;&space;i|Imag(z)|)^2&space;&plus;&space;c" title="z = (|Real(z)| + i|Imag(z)|)^2 + c" /> <br>



### Plans
- [x] Multithreading
- [] Smooth(continuous) fractal iteration values
- [] Multi-sampling