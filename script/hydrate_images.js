document.addEventListener("DOMContentLoaded", () => {
    let dehydratedImages = document.querySelectorAll("[data-dehydrated]")
    dehydratedImages.forEach(di => {
        let poster = di.parentNode.querySelector("[data-src]")
        dehydratedData = di.getAttribute("data-dehydrated")
        if (dehydratedData) {
            hydratedSrc = hydrateColor(dehydratedData)
            di.removeAttribute("data-dehydrated")
            di.src = hydratedSrc
        }
        registerImageLoadingEvents(di, poster)
    })
});
registerImageLoadingEvents = (img, next) => {
    if (img.complete) {
        img.classList.remove("loading")
        if (next) {
            loadDataSrcImage(next)
        }
    } else {
        img.addEventListener("load", (e) => {
            e.target.classList.remove("loading")
            if (next) {
                loadDataSrcImage(next)
            }
        });
    }
    img.addEventListener("error", (e) => {
        e.target.style.display = "none"
    });
    img.addEventListener("click", (e) => {
        if (e.target.classList.has("poster")) {
            e.target.classList.toggle("loading")
        }
    })
}
loadDataSrcImage = (img) => {
    img.src = img.getAttribute("data-src")
    img.removeAttribute("data-src")
    registerImageLoadingEvents(img, null)
}