# Projet_GO_Reservation
 Développer un système de réservation de salles de classe en Go, en commençant avec une interface en ligne de commande pour gérer les réservations, les salles et les utilisateurs.

 - Alexandre Helleux    : AlexandreDjazz
 - Marc Lecomte         : Spatulox


 ## Fonctionnalités

 ### Discussion avec la base de donnée - Spatulox
 Dans un premier temps un package pouvant réaliser des requêtes en base de donnée rapidement à été dévelopé en utilisant une structure avec des fonctions dedans.
 Ce choix de créer une sorte de "classe" permet de l'instancier pour éviter de se perdre avec des fonctions reliées à rien.

 ### Package de débug - Spatulox
 Dans un deuxième temps, un "package de débug" à lui aussi été créé sous forme de "classe", permettant d'avoir des messages d'informations, de debugs et d'erreurs en couleur.
 Il avait aussi pour but d'écrire dans un fichier texte afin d'enregister des logs, ce qui a été abandonné.

 ### Constantes - Spatulox
 Un package réunissant toutes les constantes a été crée, permettant l'utilisation :
 - De la structure LogHelper (instanciée en "Log") sans devoir l'instancier dans chaque fichier
 - De constantes définnissant les noms des Tables SQL pour avoir une autocomplétion sans erreurs
 - Une constante NullString

 ### Création des structures Salles & Réservations - Spatulox
 Fichiers avec la définition des structures Salles et Réservation.
 Ces structures ont été placées dans un package séparé appelé "Model" pour éviter les imports récursif si ces structures avaient été placée dans la package reservation

 ### Package Réservation - AlexandreDjazz / Spatulox
 Ce package regroupe toutes les fonctionnalités autour des réservations et des salles.
 C'est le package principal du projet regroupant salle.go et reservation.go
 Salle.go : AlexandreDjazz
 Reservation.go : Spatulox

 ### Interface WEB - AlexandreDjazz / Spatulox
 La création d'une interface graphique WEB à été très rapide car pensée dès le début lors de la création des fonctions principales.



 ## Séparation du travail

 AlexandreDjazz :
 - Création de la Base de donnée
 - Menu CLI
 - Gestion des salles
 - Import / Export JSON (Interface Web)
 - Interface web gestion des salles

 - Easter Egg

 Spatulox :
 - Création du docker (??)
 - Discussion avec la base de donnée
 - Création des Constantes
 - Gestion des réservations
 - Import / Export JSON (Fonctions)
 - Mise en place du serveur Web
 - Interface web gestion des reservations
 
 - Easter Egg

 ## Choix de conception

 ### Base de donnée
 Nous avons choisi de créer des fonctions de Base de données permettant d'éviter de copier-coller du code, et afin de gagner du temps lors de la discussion avec la base de donnée. De plus c'est quelque chose mis en place dans tous les projets commun, impliquant une base de donnée, entre AlexandreDjazz et Spatulox 👀.
 
 Ces fonctions prennent un paramètre de "débug" permettant de visualiser la requête réalisée pour un meilleur débuguage.
 De plus ces fonctions, une fois bien construite, évitent les problèmes de synthaxe de code, et donc du débuguage inutile.

 ### Log
 Ces fonctions permettent un suivi plus simple et durable si le projet doit tourner en arrière plan sur une machine à distance. Elles enregistrent (normalement, mais pas là) dans un fichier texte de log, pour savoir où le programme passe et peut créer des erreurs.
 Elle permettent aussi une meilleur aide pour le débugguage en mettant des couleurs pour les messages.

 ### Constantes
 La création de constantes à permis d'éviter des problèmes de synthaxe pour les noms des tables, et a simplifié la mise en place des fonctions de Log

 ### Package principal - Fonctions
 Le choix de laisser la possiblité d'appeler les fonctions avec des paramètres null ou non, a simplifé l'implémentation d'une interface graphique et de leur appelation dans les Handlers.
 Lorsque les fonctions ont des paramètres à "nil", la fonction demande de rentrer des informations en ligne de commande.
 Si les fonctions recoivents des paramètres, elles vérifient le bon format de ces données, puis elles exécutent les lignes de codes restantes.
