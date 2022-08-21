# Go-TinyCC

## Description

Go-TinyCC is a small wraper library that allows **Dynamic** runtime C code compiling and linking using Golang.

It used [tinyCC](https://bellard.org/tcc/) compiler to compile/link runtime C sources. 
So there is some benefits over gcc or clang compilers. 
> from tinyCC documentation

- SMALL! You can compile and execute C code everywhere, for example on
  rescue disks. (This also helpful when you need compile code on fly) 
- FAST! tcc generates optimized x86 code. No byte code
  overhead. Compile, assemble and link about 7 times faster than 'gcc
  -O0'.
- UNLIMITED! Any C dynamic library can be used directly. TCC is
  heading toward full ISOC99 compliance. TCC can of course compile
  itself.  
- SAFE! tcc includes an optional memory and bound checker. Bound
  checked code can be mixed freely with standard code.
- Compile and execute C source directly. No linking or assembly
  necessary. Full C preprocessor included.
- C script supported : just add '#!/usr/local/bin/tcc -run' at the first
  line of your C source, and execute it directly from the command
  line.
