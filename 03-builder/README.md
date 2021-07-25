# Builder Pattern

When construction gets a little too complicated.

## Motivation

Some object are simple to construct (simple constructor call)

Other objects require several steps to create it and have a factory function with 10 arguments is not productive.

Builder provides an API for constructing an object step-by-step. 

## Goal

When piecewise object construction is complicated, provide an API for doing it succinctly.

- Un builder est un composant séparé utilisé pour construire un objet
- Il permet également de faire un builder fluent qui retourner le receiver (chaine)
- Différents facette d'un objet peuvent être construit avec différents builders travaillant en tandem avec différentes structs



