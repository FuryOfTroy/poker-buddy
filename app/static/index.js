function card_onclick(card) {
    if(card.className != "selected") {
        card.className = "selected"
    } else {
        card.className = null
    }
}

function evaluatehand_onclick() {
    let cards = document.getElementsByClassName("selected")
    let cardNames = Array.from(cards).map((card) => card.name)
    let body = JSON.stringify(cardNames)
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("result").innerHTML = this.responseText;
        }
    };
    xhttp.open("POST", "api/cards/evaluate", true);
    xhttp.send(body);
}

function calculateodds_onclick() {
    let cards = document.getElementsByClassName("selected")
    let cardNames = Array.from(cards).map((card) => card.name)
    let body = JSON.stringify(cardNames)
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("result").innerHTML = this.responseText;
        }
    };
    xhttp.open("POST", "api/cards/calculateodds", true);
    xhttp.send(body);
}