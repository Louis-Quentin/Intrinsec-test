# Intrinsec-test

Description
Ce projet est un script Go qui récupère les articles d'un blog, les analyse, et les enregistre dans un fichier JSON. Le processus comprend plusieurs étapes : la récupération des données HTML, le parsing des articles via des expressions régulières, puis la sauvegarde des résultats dans un fichier JSON.

Structure du Projet

1. fetchData
Cette fonction est responsable de la récupération des données HTML depuis une page web spécifiée par l'URL. Elle effectue une requête HTTP, lit la réponse et renvoie le corps de la réponse sous forme de chaîne de caractères.

2. parseData
Cette fonction analyse le corps HTML pour extraire les liens et les titres des articles en utilisant des expressions régulières.

3. saveAsJSON
Cette fonction permet de sauvegarder les articles extraits dans un fichier JSON.

4. saveResults
Cette fonction prend les résultats extraits par parseData, crée des objets Article et les passe à saveAsJSON pour les sauvegarder.


Comment lancer le projet

Prérequis
Go doit être installé sur votre machine. Vous pouvez vérifier cela avec la commande suivante :
```go version```

Cloner ou télécharger le projet. Si vous utilisez Git, vous pouvez cloner le projet avec la commande suivante :
```git clone git@github.com:Louis-Quentin/Intrinsec-test.git```
```cd Intrinsec-test ```

Construire et exécuter le projet :
Pour exécuter le projet, vous pouvez simplement utiliser la commande suivante :
```go run main.go```

Après avoir exécuté le programme, vous devriez trouver un fichier Articles.json dans le même répertoire. Ce fichier contiendra les titres et liens des articles extraits du blog.
