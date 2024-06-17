var g_url = "/notes/";

function url(id) {
    return g_url + id.toString();
}

function display_notes_get_response(responseText) {
    display_response(responseText)
    document.getElementById("notes-data").innerHTML = responseText;
}

function display_response(responseText) {
    var outputElementId = "console";
    document.getElementById(outputElementId).innerHTML = responseText;
    console.log(xhttp.responseText);
}

function notes_get(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("GET", url(id));

    xhttp.onreadystatechange = display_notes_get_response(xhttp.responseText);
    xhttp.send()
}

function  notes_post() {
    var xhttp = new XMLHttpRequest();
    xhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhttp.open("POST", g_url);

    xhttp.onreadystatechange = display_response(xhttp.responseText);

    var data = "title=" + document.getElementById("title") + "&&body=" + document.getElementById("body");
    xhttp.send(data);
}

function  notes_put(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xhttp.open("PUT", url(id));

    xhttp.onreadystatechange = display_response(xhttp.responseText);

    var title = document.getElementById("title"+ id.toString()); 
    var body = document.getElementById("body"+ id.toString());
    var data = "title=" + title + "&&body=" + body;
    xhttp.send(data);
}

function  notes_delete(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("DELETE", url(id));

    xhttp.onreadystatechange = display_response(xhttp.responseText);
    xhttp.send();
}