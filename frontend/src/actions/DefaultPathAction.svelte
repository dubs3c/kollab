<script context="module" lang="ts">

import {ShowToast} from "./../main"
import type {Path, LogEvent} from "../models";

let baseURL = import.meta.env.VITE_BASE_URL;
if(baseURL == undefined){
    baseURL = "";
}

export async function AddDefaultPath(path: Path, url: string = baseURL + "/api/defaulthttp", method: string = "POST"): Promise<Path> {
    return await fetch(url, {
		method: method,
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
            return JSON.parse(JSON.stringify(res.data))
        } else {
            ShowToast(res.data.error, true);
            console.log('Error:', res);
            return {}
        }
    })
    .catch((error) => {
        ShowToast(error, true)
        console.error('Network Error:', error);
        return {}
    });
    
}


export async function GetDefaultPaths(): Promise<Path[]> {
    return await fetch(baseURL + "/api/defaulthttp", {
		method: 'GET'
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
            let p: Path[] = JSON.parse(JSON.stringify(res.data));
            return p;
        } else {
            let p: Path[] = []
            return p;
        }
    })
}


export async function GetEvents(pathId: string = ""): Promise<LogEvent[]> {
    let url = baseURL + "/api/events"
    if(pathId != "" && pathId != undefined && pathId != null) {
        url = baseURL + "/api/events/" + pathId
    }

    return await fetch(url, {
		method: 'GET'
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
            let p: LogEvent[] = JSON.parse(JSON.stringify(res.data));
            return p;
        } else {
            let p: LogEvent[] = []
            return p;
        }
    })
}


export async function GetPath(id: string): Promise<Path> {
    return await fetch(baseURL + "/api/defaulthttp/" + id, {
		method: 'GET'
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
            let p: Path = JSON.parse(JSON.stringify(res.data));
            return p;
        } else {
            let p: Path = {Headers:[],Body:"",Id:"",Verb:"",Path:""}
            return p;
        }
    })
}

</script>