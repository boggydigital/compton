document.addEventListener("DOMContentLoaded", () => {
    let submitNavLinks = document.querySelectorAll("nav-links > ul > li > a.submit")
    submitNavLinks.forEach(snl => {
        snl.addEventListener("click", e => {
            let form = e.currentTarget.closest("form")
            if (form && form.submit) {
                form.submit()
            }
        })
    })
});