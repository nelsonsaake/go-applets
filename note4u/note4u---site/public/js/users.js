// note4u manager

// the url
// all the requests will use this files
var g_url = "/users/";

// /users/1 points to the user in the users table with id=1
// this function take the id, appends it to the url
// url is expected to point to some resource 
// the the id is empty or incorrect it will return all users
function url(id) {
    return g_url + id.toString();
}

// the post and put method are required to make a request to the url
// providing an email and password
// the put method is used with the aim of modifying an existing resource
// the post method is expected to create a resource
// => the input using the id "email" takes the data for post request
// and the input using the id "email"<n> takes the data for a particular put request 
// 
// this function cre ates a name/key, value pair for the email field for post and put
// if it's called from put the id will be provided in the argument 
// if it is called from post the id is not expected 
// either way, this function will return a name, value pair

function email(id) {

    // we get the email value from the element with that particular id
    // that is email with an id; eg. "email1" for put
    // or jut "email" for post
    var email = document.getElementById("email"+ id).value; 
    return "email="+email; 
}

// pwd works just like the email
function pwd(id) {
    var pwd = document.getElementById("pwd" + id).value;
    return "password="+pwd;
}

// 
function users_post() {
    var xhttp = new XMLHttpRequest();
    xhttp.open("POST", g_url);
    xhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");

    xhttp.onreadystatechange = function () {
		window.location.href = g_url;
	};

    email = document.getElementById("email").value;
    pwd = document.getElementById("pwd").value;
    var data = "email="+email+"&&password="+pwd;
    xhttp.send(data);
}

// 
function users_put(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("PUT", url(id));
    xhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");

    xhttp.onreadystatechange = function () {
		window.location.href = g_url;
	};
    
    var data = email(id) + "&&" + pwd(id);
    xhttp.send(data);
}

// 
function users_delete(id) {
    var xhttp = new XMLHttpRequest();
    xhttp.open("DELETE", url(id));

    xhttp.onreadystatechange = function () {
		window.location.href = g_url;
	};
    xhttp.send();
}