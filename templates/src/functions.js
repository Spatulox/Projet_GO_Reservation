function showPopup(message) {
    console.log(message)
    document.getElementById("errorMessage").textContent = message;
    document.getElementById("errorPopup").classList.add("active");

    setTimeout(()=>{
        closePopup()
    }, 5000)
}

function closePopup() {
    document.getElementById("errorPopup").classList.remove("active");
}



function redirectToMainMenu(){
    window.location.href = `/`;
}

function redirectToMainList() {
    window.location.href = `/reservation`;
}

function redirectToCreateReserv() {
    window.location.href = `/reservation/create`;
}

function redirectToRoomList() {
    const idRoom = document.querySelector('input[name="idRoom"]').value;
    window.location.href = `/reservation/list?idRoom=${idRoom}`;
}

function redirectToDateList() {
    const idRoom = document.querySelector('input[name="idDate"]').value;
    window.location.href = `/reservation/list?idDate=${idRoom}`;
}

function redirectToIdList(id) {
    window.location.href = `/reservation/list?idReserv=${id}`;
}

function redirectDelete(id){
    fetch(`/reservation/cancel?idReserv=${id}`, {
        method: 'GET'
    })
    .then(response => {
        if (response.ok) {
            response.text().then(message => {
                showPopup(message);
            });
        } else {
            response.text().then(errorMessage => {
                showPopup(errorMessage);
            });
        }
    })
    .catch(error => {
        showPopup('Erreur lors de la requête :', error);
    });
}













//-----------------------------------------------------------

const urlParams = new URLSearchParams(window.location.search);

if (urlParams.has('message')) {
    // Le paramètre "message" est présent
    const messageValue = urlParams.get('message');
    showPopup(messageValue)
}
