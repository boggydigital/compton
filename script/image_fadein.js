document.addEventListener("DOMContentLoaded", () => {
    let posters = document.querySelectorAll("img.poster")
    posters.forEach(poster => {
        if (poster.complete) {
            poster.classList.remove("loading")
        } else {
            poster.addEventListener("load", (e) => {
                poster.classList.remove("loading")
            });
        }
        poster.addEventListener("error", (e) => {
            e.target.style.display = "none"
        });
        poster.addEventListener("click", (e) => {
            e.target.classList.toggle("loading")
        })
    })
    let placeholders = document.querySelectorAll("img.placeholder")
    placeholders.forEach(placeholder => {
        placeholder.addEventListener("error", (e) => {
            e.target.style.display = "none"
        });
    })
});