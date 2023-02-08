

function onAddPathClick() {
    let path = document.getElementById("http-path").value;
    let verb = document.getElementById("http-path-verb").value;
    let headers = document.getElementById("http-path-headers").value;
    let body = document.getElementById("http-path-body").value;

    if (path == null || path == undefined) {
        console.log("Path is fucked");
    }

    if (verb == null || verb == undefined) {
        console.log("verb is fucked");
    }

    if (headers == null || headers == undefined) {
        console.log("headers is fucked");
    }

    if (body == null || body == undefined) {
        console.log("body is fucked");
    }

    addPath(path, verb, headers, body);
}


function addPath(path, verb, headers, body) {

    fetch('http://127.0.0.1:8080/api/defaulthttp', {
        method: 'POST',
        body: btoa(body),
        credentials: 'same-origin', 
      })
      .then((response) => response.json())
      .then((result) => {
            console.log('Success:', result);
       })
       .catch((error) => {
            console.error('Error:', error);
       });
}