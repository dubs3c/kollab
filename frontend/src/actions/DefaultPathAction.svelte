<script context="module" lang="ts">

import {baseURL} from "../config.dev.js";
import {ShowToast} from "./../main"
import type {Path} from "../models";


export async function AddDefaultPath(path: Path): Promise<boolean> {
    return await fetch(baseURL + "/api/defaulthttp", {
		method: 'POST',
		body: JSON.stringify({...path})
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
            ShowToast("Path added...");
            console.log('Success:', res);
            return true;
        } else {
            ShowToast(res.data.error, true);
            console.log('Error:', res);
            return false;
        }
    })
    .catch((error) => {
        ShowToast(error, true)
        console.error('Network Error:', error);
        return false;
    });
    
}


export async function GetDefaultPaths() {
    await fetch(baseURL + "/api/defaulthttp", {
		method: 'GET'
	})
    .then((response) => 

        response.json().then(data => ({
            data: data,
            status: response.status,
            ok: response.ok
        })
    ))
}

</script>