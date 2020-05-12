# HackDalton: Secret Agent (Writeup)

> Warning! There are spoilers ahead

The image is encoded with least significant bit encoding (LSB), a form of [steganography](https://en.wikipedia.org/wiki/Steganography)

From the wikipedia page for steganography:
> For example, a 24-bit bitmap uses 8 bits to represent each of the three color values (red, green, and blue) of each pixel. The blue alone has 28 different levels of blue intensity. The difference between 11111111 and 11111110 in the value for blue intensity is likely to be undetectable by the human eye. Therefore, the least significant bit can be used more or less undetectably for something else other than color information. If that is repeated for the green and the red elements of each pixel as well, it is possible to encode one letter of ASCII text for every three pixels. 

A script to encode and decode images can be found at [main.go](./main.go).