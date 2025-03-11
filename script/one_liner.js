document.addEventListener("DOMContentLoaded", () => {
    let oneLiners = document.querySelectorAll("one-liner")
    oneLiners.forEach(ol => {
        ol.addEventListener("click", (e) => {
            if (e.target.parentElement) {
                e.target.parentElement.classList.toggle("expanded")
            }
        })
    })
})