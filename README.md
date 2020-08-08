# Hashcracker - written in GO
A GO script to generate and (fast) brute force MD5/SHA1/SHA256 hashes.

How to install:

1) git clone https://github.com/r1me75/Gohashcracker.git
2) cd Gohashcracker
3) go run Gohashcracker

If you don't have GO, you can install it here -> https://golang.org/dl/

You can easily brute force MD5, SHA-1 and SHA-256 hashes.
You can also generate those for your own passwords.
I have tested it with rockyou.txt, it takes exactly 5 seconds to crack the last password of rockyou.txt(14.000.000) while it takes 17 seconds with my Python script.
