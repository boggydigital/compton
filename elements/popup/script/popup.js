document.addEventListener("DOMContentLoaded", () => {
    let popupActors = document.querySelectorAll("[data-popup-target]")
    popupActors.forEach(pa => {
        pa.addEventListener("click", (e) => {
            let targetId = e.target.getAttribute("data-popup-target");
            let targetElement = document.querySelector("#"+targetId)
            if (targetElement) {
                let popupValue = targetElement.getAttribute("data-popup")
                if (popupValue == "hide") {
                    targetElement.setAttribute("data-popup", "show")
                } else {
                    targetElement.setAttribute("data-popup", "hide")
                }
            }
        })
    })
});
