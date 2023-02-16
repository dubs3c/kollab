

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
        toast._element.childNodes[0].childNodes[0].innerText = body;
    }

    toast.show();
}