document.addEventListener("DOMContentLoaded", function () {
    const toggleButton = document.getElementById("dark-mode-toggle");

    const updateButtonText = (isDark) => {
        if (isDark) {
            toggleButton.textContent = "☾ ?";
        } else {
            toggleButton.textContent = "☼ ?";
        }
    }

    const setDarkMode = (isDark) => {
        if (isDark) {
            document.documentElement.classList.add("dark");
            localStorage.setItem("theme", "dark");
        } else {
            document.documentElement.classList.remove("dark");
            localStorage.setItem("theme", "light");
        }
        updateButtonText(isDark);
    }

    const savedTheme = localStorage.getItem("theme");
    if (savedTheme) {
        setDarkMode(savedTheme === "dark");
    } else if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        setDarkMode(true);
    }

    toggleButton.addEventListener("click", () => {
        const isDark = document.documentElement.classList.toggle("dark");
        setDarkMode(isDark);
    })
})