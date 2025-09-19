document.addEventListener("DOMContentLoaded", () => {
    let copyToClipboardElements = document.querySelectorAll("copy-to-clipboard")
    copyToClipboardElements.forEach(ctc => {
        ctc.addEventListener("click", e => {
            let cta = e.target.parentNode.querySelector(".copy-to-clipboard-cta")

            navigator.clipboard.writeText(e.currentTarget.getAttribute("data-value")).then(
            res => {
                let success = e.target.parentNode.querySelector(".copy-to-clipboard-success")
                if (success) {
                    success.style.display = "initial"
                    if (cta) {
                        cta.style.display = "none"
                    }
                }
            },
            err => {
                let error = e.target.parentNode.querySelector(".copy-to-clipboard-error")
                if (error) {
                    error.style.display = "initial"
                    if (cta) {
                        cta.style.display = "none"
                    }
                }
            })
        })
    })
});