# 

<h1 align="center">
    <a href="https://github.com/stephenSLI/image-processing">
		Image Processing with Go
    </a>
    <br/>
</h1>

<h4 align="center">A small project used to test and learn different image processing ideas.</h4>

### Blur

A blur can be applied to a given image that is a jpeg/jpg or a png image. 
The kernel and iterations can be adjusted to change the amount of blur while
sigma can be changed to adjust the blur for the gaussian blur only.

```bash
NAME:
   main blur - Perform a blur on a given image.

USAGE:
   main blur [command options] [arguments...]

OPTIONS:
   --file value, -f value        The path to the file being blurred.
   --type value, -t value        The type of blur to apply. (default: "mean", "gaussian")
   --iterations value, -i value  The number of iterations to apply to the blur. (default: 3)
   --kernel value, -k value      The size of the kernel used on the blur. (default: 31)
   --sigma value, -s value       Sets the sigma value if used in the blur, e.g Gaussian blur. (default: 40)
   --help, -h                    show help (default: false)
```
