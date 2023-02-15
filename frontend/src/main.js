

export function ShowToast(body, error) {
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