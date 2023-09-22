

function addPathToTable(path, verb) {
    console.log(path);
    /*
    <tr>
        <td>/x.js</td>
        <td>GET</td>
        <td><button type="button" class="btn btn-danger btn-sm">Delete</button>
            <button type="button" class="btn btn-primary btn-sm">View</button>
        </td>
    </tr>
    */

    let tr = document.createElement("tr");
    let tdPath = document.createElement("td");
    let tdVerb = document.createElement("td");
    let tdButtons = document.createElement("td");
    tdPath.innerText = path 
    tdVerb.innerText = verb
    tdButtons.innerHTML = `<button type="button" class="btn btn-danger btn-sm">Delete</button>
    <button type="button" class="btn btn-primary btn-sm">View</button>`
    tr.append(tdPath);
    tr.append(tdVerb);
    tr.append(tdButtons);

    let pathTable = document.getElementById("path-table");
    pathTable.append(tr);
}


function onAddPathClick() {
    console.log("what")
    let path = document.getElementById("http-path").value;
    let verb = document.getElementById("http-path-verb").value;
    let headers = document.getElementById("http-path-headers").value;
    let body = document.getElementById("http-path-body").value;

    if (path == null || path == undefined || path == "") {
        console.log("Path is fucked");
        showToast("Must specify path...",true);
        return;
    }

    if (verb == null || verb == undefined) {
        console.log("verb is fucked");
    }

    if (headers == undefined || headers == "") {
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

    const lol = data;

    fetch('/api/defaulthttp', {
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
                showToast("Path added...");
                addPathToTable(lol.path, lol.verb);
                console.log('Success:', res);
            } else {
                showToast(res.data.error, true);
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