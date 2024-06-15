function card_onclick(card) {
    if(card.className != "selected") {
        card.className = "selected"
    } else {
        card.className = null
    }
}

function evaluatehand_onclick() {
    call_api("api/cards/evaluate");
}

function calculateodds_onclick() {
    call_api("api/cards/calculateodds");
}

function call_api(url) {
    let cards = document.getElementsByClassName("selected")
    let cardNames = Array.from(cards).map((card) => card.name).join("")
    let body = JSON.stringify(cardNames)
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("result").innerHTML = this.responseText;
        }
    };
    xhttp.open("POST", url, true);
    xhttp.send(body);
}