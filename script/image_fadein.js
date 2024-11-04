document.addEventListener("DOMContentLoaded", () => {
    let images = document.querySelectorAll("issa-image>img")
    images.forEach(img => {
        if (img.complete) {
            img.classList.remove("loading")
        } else {
            img.addEventListener("load", (e) => {
                e.target.classList.remove("loading")
            });
        }
        img.addEventListener("error", (e) => {
            e.target.style.display = "none"
        });
        img.addEventListener("click", (e) => {
            e.target.classList.toggle("loading")
        })
    })
});