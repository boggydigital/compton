const months = ["JAN","FEB","MAR","APR","MAY","JUN","JUL","AUG","SEP","OCT","NOV","DEC"]
document.addEventListener("DOMContentLoaded", () => {
    let todaySymbol = document.querySelector("symbol#today")
    if (todaySymbol) {
        let today = new Date()
        let todayMonth = todaySymbol.querySelector("#today-month")
        let todayDay = todaySymbol.querySelector("#today-day")
        if (todayMonth !== null && todayDay !== null) {
            todayMonth.textContent = months[today.getMonth()]
            todayDay.textContent = today.getDate()
        }
    }
})