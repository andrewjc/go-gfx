# go-gfx
A game/graphics engine written in golang

An experiment in game engines in Go.

Features:
* NDC (Normalized Device Coordinate) screen space rendering - ensuring that rendering is perspective correct on ANY monitor size eg very ultra wide.
* Entity system
* Physics force system - gravity, springs, electromagnetism (repell charged particles etc).
* Obj file loader

WIP: Physically based material system.


Requirements:
* GCC for compiling opengl libraries. If on windows use MingW (easiest install via choco: choco install mingw).
* Go compiler and tools
