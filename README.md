# Push-Swap

Ce projet a ete realise dans le cadre du cursus Ynov/Ytrack. Il consiste a trier une liste d'entiers en utilisant deux piles (stacks) et un ensemble d'instructions specifiques, avec le moins d'operations possible.

Le projet contient deux programmes ecrits en Go :
- push-swap : Calcule et affiche la liste d'instructions pour trier les nombres recus en argument.
- checker : Lit des instructions sur l'entree standard et verifie si elles permettent de trier correctement la pile.

## Prerequis

- Go (version 1.21 ou superieure)

## Installation et Compilation

Clonez le depot et placez-vous a la racine du projet :

git clone https://github.com/TON_NOM_UTILISATEUR/push-swap.git
cd push-swap

Compilez les deux programmes :

go build -o push-swap .
go build -o checker/checker ./checker

## Utilisation

### Le programme push-swap

Il prend une liste d'entiers en argument et affiche les instructions pour les trier.

./push-swap "2 1 3 6 5 8"

En cas d'erreur (arguments non entiers, doublons), le programme affiche Error sur la sortie d'erreur.

### Le programme checker

Il prend une liste d'entiers en argument et lit les instructions sur l'entree standard. Si la pile est triee a la fin, il affiche OK, sinon KO.

echo -e "rra\npb\nsa\nrra\npa" | ./checker/checker "3 2 1 0"

Vous pouvez verifier que votre push-swap fonctionne bien en combinant les deux :

ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker/checker "$ARG"

## Algorithme utilise

- Pour les petites listes (3 a 5 elements) : Un tri specifique tres rapide qui utilise les instructions de base.
- Pour les grandes listes (plus de 5 elements) : Un algorithme de tri par selection optimise. Le programme cherche le plus petit element et calcule le chemin le plus court pour l'envoyer dans la pile B, puis remet tout dans l'ordre.

## Tests

Des tests unitaires sont inclus pour verifier le bon fonctionnement des piles :

go test ./...

## Auteur

Khaldoune Sherine - Etudiant a Ynov
