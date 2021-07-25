# Factory Pattern

Une façon de controller comment un objet est construit.

## Motivation

- La logique de création d'un objet devient trop compliquée
- Un exemple est une struct avec beaucoup de fields et qui doivent être initialisés correctement.
- Il s'agit d'une création complète d'un objet. Pas comme le builder qui le créé une pièce après l'autre). Il peut être extrait dans : 
  - une fonction séparée (factory function a.k.a. Constructor)
  - une structure séparée
    
Un composant responsable à lui tout seul de la création complète d'un objet. 

## Types

- Factory function: (constructeur) est une fonction permettant de construire un objet (tout simplement)
- Factory est une entité qui peut s'occuper de la création d'un objet
- Cela peut être une function ou une structure dédiée