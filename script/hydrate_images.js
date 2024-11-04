document.addEventListener("DOMContentLoaded", () => {
    let dehydratedImages = document.querySelectorAll("[data-dehydrated]")
    dehydratedImages.forEach(di => {
        di.addEventListener("error", (e) => {
            e.target.style.display = "none"
        });
        dehydratedData = di.getAttribute("data-dehydrated")
        hydratedSrc = hydrateColor(dehydratedData)
        di.removeAttribute("data-dehydrated")
        di.src = hydratedSrc
        console.log(di.classList)
    })
});