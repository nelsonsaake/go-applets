g_url = "/users/"

function url(id) {
    return g_url + id;
}

function email(id) {
    var email 
    if (id != "") email = document.getElementById("email" + id).value;
    else email = document.getElementById("email").value;
    return "email="+email;
} 

function pwd(id) {
    pwd = document.getElementById("pwd"+ id).value;
    return "password="+pwd;
}

function users_post(){
    var xhttp = new XMLHttpRequest();
    xhttp.open("POST", g_url);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

    xhttp.onreadystatechange = function() {
        window.location.href = g_url;
    };

    email = document.getElementById("email").value;
    pwd = document.getElementById("pwd").value;
    data = "email="+email+"&&password="+pwd;
    xhttp.send(data);
}

function users_put(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("PUT", url(id), true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

    xhttp.onreadystatechange = function () {
        window.location.href = g_url;
    }

    var data = email(id) + "&&" + pwd(id);
    xhttp.send(data); 
}

function users_delete(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("DELETE", url(id), true);

    xhttp.onreadystatechange = function () {
        window.location.href = g_url;
    };
	
    xhttp.send();
}