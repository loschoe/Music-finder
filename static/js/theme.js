const btn = document.getElementById("theme-toggle");
const icon = document.getElementById("theme-icon");
const root = document.documentElement;

// Fonction pour mettre à jour l'icône selon le thème
function updateIcon(theme) {
  if (theme === "light") {
    icon.className = "fas fa-moon"; // clair
  } else {
    icon.className = "fas fa-lightbulb";  // sombre
  }
}

// 1) Lecture du thème enregistré
let savedTheme = localStorage.getItem("theme");

// 2) Sombre par défaut si rien n'est enregistré
if (!savedTheme) {
  savedTheme = "dark";
  localStorage.setItem("theme", "dark");
}

// 3) Application du thème au chargement
root.setAttribute("data-theme", savedTheme);
updateIcon(savedTheme);

// 4) Toggle au clic
btn.addEventListener("click", (e) => {
  e.preventDefault();

  const current = root.getAttribute("data-theme");
  const newTheme = current === "light" ? "dark" : "light";

  root.setAttribute("data-theme", newTheme);
  localStorage.setItem("theme", newTheme);

  updateIcon(newTheme);
});