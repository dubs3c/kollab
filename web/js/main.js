

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

    if (headers == null || headers == undefined || headers == "") {
        console.log("headers is fucked");
        headers = null;
    }

    if (body == null || body == undefined) {
        console.log("body is fucked");
    }

    const d = {
        "path":path,
        "verb": verb,
        "headers": headers,
        "body": btoa(body),
    }
    addPath(d);
}


function addPath(data) {

    fetch('http://127.0.0.1:8080/api/defaulthttp', {
        method: 'POST',
        body: JSON.stringify(data),
        credentials: 'same-origin', 
      })
      .then((response) => 

        response.json().then(data => ({
            data: data,
            status: response.status,
            ok: response.ok
        })

        ))
        .then(res => {

            if(res.ok) {
                showToast("Path added...")
                console.log('Success:', res);
            } else {
                showToast(res.data.error,true);
                console.log('Error:', res);
            }
        })
       .catch((error) => {
            showToast(error, true)
            console.error('Network Error:', error);
       });
}


function showToast(body, error) {
    let thetoast = "";
    let toast  = "";
    if (error) {
        thetoast = document.getElementById('toastError');
        toast = new bootstrap.Toast(thetoast);
        toast._element.lastElementChild.innerText = body;
    } else {
        thetoast = document.getElementById('toastSuccess');
        toast = new bootstrap.Toast(thetoast);
        toast._element.childNodes[1].childNodes[1].innerText = body;
    }

    toast.show();
}