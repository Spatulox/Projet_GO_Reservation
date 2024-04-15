function showPopup(message) {
    console.log(message)
    document.getElementById("errorMessage").textContent = message;
    document.getElementById("errorPopup").classList.add("active");
}

function closePopup() {
    document.getElementById("errorPopup").classList.remove("active");
}

const urlParams = new URLSearchParams(window.location.search);

if (urlParams.has('message')) {
    // Le paramètre "message" est présent
    const messageValue = urlParams.get('message');
    showPopup(messageValue)

    setTimeout(()=>{
        closePopup()
    }, 5000)
}