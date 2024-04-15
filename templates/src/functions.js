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
                window.location.href = `/reservation?message=${message}`;
                //showPopup(message);
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

redirectUpdate

function redirectUpdate(id){

    // Récupérer l'élément select
    var selectElement = document.querySelector('select[name="etat"]');

    // Récupérer la valeur sélectionnée
    var etat = selectElement.value;

    fetch(`/reservation/update?idReserv=${id}?etat=${etat}`, {
        method: 'GET'
    })
    .then(response => {
        if (response.ok) {
            response.text().then(message => {
                window.location.href = `/reservation/list?idReserv=${id}`;
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


async function getAllRoomAvailable() {

    const horaire_start_date = document.getElementById("horaire_start_date").value
    const horaire_start_time = document.getElementById("horaire_start_time").value

    const horaire_end_date = document.getElementById("horaire_end_date").value
    const horaire_end_time = document.getElementById("horaire_end_time").value

    const ulCreateReservation = document.getElementById("ulCreateReservation")

    if (horaire_start_date == null || horaire_start_time == null || horaire_end_date == null || horaire_end_time == null){
        return
    }

     // Vérifier le format de la date
    let dateRegex = /^\d{4}-\d{2}-\d{2}$/;
    if (!dateRegex.test(horaire_start_date) || !dateRegex.test(horaire_end_date)) {
        showPopup("La date doit être au format AAAA-MM-JJ");
        return;
    }

    // Vérifier le format de l'heure
    let timeRegex = /^\d{2}:\d{2}$/;
    if (!timeRegex.test(horaire_start_time) || !timeRegex.test(horaire_end_time)) {
        showPopup("L'heure doit être au format HH:MM");
        return;
    }

    startDateTime = horaire_start_date + " " + horaire_start_time
    endDateTime = horaire_end_date + " " + horaire_end_time


    const data = {
        startDateTime: startDateTime,
        endDateTime: endDateTime
    };

    try {
        const response = await fetch('/salle/getAllAvail', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });


        if (response.ok) {
            const data = await response.json();

            const ulCreateReservation = document.getElementById('ulCreateReservation');
            ulCreateReservation.innerHTML = ""

            if(data.length > 0){
                data.forEach(salle => {
                    let li = document.createElement("li");
                    let tmp = `${salle.IdSalle} : ${salle.NomSalle} (${salle.PlaceSalle} Places)`;
                    li.textContent = tmp;
                    ulCreateReservation.appendChild(li);
                    console.log(tmp);
                });
            }
            else{
                ulCreateReservation.innerHTML = "Veuillez choisir une date et heure de départ et de fin pour voir les salles disponibles"
            }



            
        } else {
            const errorMessage = await response.text();
            showPopup(errorMessage);
        }
    } catch (error) {
        showPopup('Erreur lors de la requête :', error);
    }
}







//-----------------------------------------------------------

const urlParams = new URLSearchParams(window.location.search);

if (urlParams.has('message')) {
    // Le paramètre "message" est présent
    const messageValue = urlParams.get('message');
    showPopup(messageValue)
}
