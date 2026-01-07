# 🎵 Groupie-Tracker 

## 🚀 Présentation   :
Bienvenue dans le dépôt Github du **Projet Groupie Tracker**, un site web utilisant une API externe codé en Go qui affiche des informations détaillées sur des artistes, groupes musicaux.
Ce projet est développé dans le cadre d'un module à **STRASBOURG Ynov Campus**.

## 📖 Description

**Music Tracker** est une application web développée en Go qui permet de :
- 🔍 Rechercher des artistes et groupes musicaux
- 📊 Visualiser leurs informations (membres, date de création, premier album)
- 🎤 Consulter leurs dates et lieux de concerts
- 🌍 Explorer une base de données complète d'artistes internationaux

Le projet utilise l'API [Groupie Trackers](https://groupietrackers.herokuapp.com/api) pour récupérer les données en temps réel.

## ✨ Fonctionnalités

- ✅ **Page d'accueil** avec liste complète des artistes
- ✅ **Recherche dynamique** par nom d'artiste
- ✅ **Page détaillée** pour chaque artiste avec :
  - Photo de l'artiste/groupe
  - Liste des membres
  - Date de création
  - Premier album sorti
  - Dates et lieux de concerts
- ✅ **Design moderne** avec interface responsive et compacte 
- ✅ **Architecture propre** avec séparation des fonctionnalités 

## 🛠️ Installation et exécution :
### 1. Cloner le dépôt
```bash
git clone https://github.com/loschoe/Groupie-Tracker-SCHOEPF_Camara.git
```
### 2. Installer les dépendances Go
```bash
go mod tidy
```
### 3. Lancer le serveur
```bash
go run .
```
### 4. Ouvrir la page 
Ouvrez votre navigateur et allez sur ```http://localhost:8080```.

<img width="1887" height="659" alt="image" src="https://github.com/user-attachments/assets/502a4ce3-0a1a-4969-babb-69bf28d05feb" />

## 📁 Arborescence des dossiers/fichiers 
```
groupie-tracker/
│
├── main.go            # Point d'entrée de l'application
├── go.mod             # Fichier de module GO
├── .gitignore         # Ignore des fichiers stockés 
├── README.md          # Fichier de présentation du porjet 
│
├── handlers/          # Gestionnaire de routes HTTP
│   ├── home.go       
│   ├── artist.go
│   └── about.go
│
├── models/            # Structure de données 
│   ├── artist.go
│   └── relation.go
│
├── services/          # Utilisation de l'API
│   └── api.go
│
├── utils/            # Utilitaires (formatage) 
│   └── formatter.go
│
├── templates/        # Templates HTML
│   ├── home.html
│   ├── artist.html
│   └── aboutUS.html
│
├── static/           # Fichiers statiques (styles & images)
     ├── css/
     └── img/
```

## 🛠️ Technologies Utilisées

- **Backend** : Go (Golang)
- **Frontend** : HTML, CSS
- **Templates** : `html/template`
- **API** : [Groupie Trackers API](https://groupietrackers.herokuapp.com/api)

## 📸 Captures d'Écran

### Page Artiste
<img width="1918" height="858" alt="image" src="https://github.com/user-attachments/assets/a09db26d-db63-4c47-bd4f-0655e3450af6" />

## 🤝 Contribution

Les contributions sont les bienvenues ! Pour contribuer :

1. **Fork** le projet
2. Créez une **branche** pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. **Committez** vos changements (`git commit -m 'Add some AmazingFeature'`)
4. **Push** vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrez une **Pull Request**

## 👥 Auteurs

- **Loschoe** - [GitHub](https://github.com/loschoe)
- **Timcmr** - [[GitHub](https://github.com/timcmr)
---

⭐ **Si vous aimez ce projet, n'oubliez pas de lui donner une étoile !** ⭐
