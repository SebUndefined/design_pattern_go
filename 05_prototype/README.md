# Prototype Pattern

Quand il est plus facile de copier un objet existant pour initialiser complètement un nouveau.

## Motivation

- Objets compliqués (ex: voitures), ne sont pas élaboré en recommençant à zéro.
  - Ils héritent de design précédents
- Un design existant (partiel ou complet) est un Prototype
- Nous faisons une copie du prototype et on le customise
  - Requiert une copie complète (deep copy)
- Nous faisons un clone "facile" (via factory)

Prototype : Un objet partiellement ou complètement initialisé que l'on copie (clone) et que l'on utilise. 

## Conclusion

- Implémenter un prototype, construire un objet à partir de ce dernier et le stocker
- Deep copy du proto
- Customiser l'instance qui en résultat

