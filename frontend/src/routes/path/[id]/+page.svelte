<script lang="ts">
	import { AddDefaultPath } from "../../../actions/DefaultPathAction.svelte";
    import type {Path} from "../../../models"
    /** @type {import('./$types').PageData} */  
    export let data: Path;




    let httpPath: string = data.Path;
    let httpVerb: string = data.Verb;
    let httpHeaders: string = data.Headers.join("\n");
    let httpBody: string = atob(data.Body);

    $: httpPath = httpPath.replace(" ", "-");

    $: if(httpPath != "" && !httpPath.startsWith("/")) {
        httpPath = "/" + httpPath;
    }

    async function onAddPathClick(){
        const splitHeaders = httpHeaders.split("\n");
        const p: Path = {
            Id: data.Id,
            Path: httpPath,
            Verb: httpVerb,
            Headers: splitHeaders,
            Body: btoa(httpBody),
        }
        let success = await AddDefaultPath(p, "/api/defaulthttp/" + data.Id, "PUT");
        console.log(success);   
    }

    async function DeletePathClick() {
        const splitHeaders = httpHeaders.split("\n");
        const p: Path = {
            Id: data.Id,
            Path: httpPath,
            Verb: httpVerb,
            Headers: splitHeaders,
            Body: btoa(httpBody),
        }
        let success = await AddDefaultPath(p, "/api/defaulthttp/" + data.Id, "DELETE");
        console.log(success);   
    }

</script>
  
<div class="row">
    <div class="col-md">
        <a href="/"><button class="btn btn-sm btn-dark"><strong>Back</strong></button></a>
        <br />
        <br />
    </div>
</div>
<div class="row">
    <div class="col-md-6">
        <h4>Endpoint</h4>
            <form>
                <div class="mb-3">
                    <label for="http-path">Path</label>
                    <input type="text" bind:value={httpPath} placeholder="/some/path/payload.js" class="form-control" id="http-path">
                </div>
                
                <div class="mb-3">
                    <label for="floatingSelect">Select HTTP Verb</label>
                    <select required bind:value={httpVerb} class="form-select" id="http-path-verb" aria-label="Select HTTP verb">
                        <option selected></option>
                        <option value="GET">GET</option>
                        <option value="POST">POST</option>
                        <option value="PUT">PUT</option>
                        <option value="PATCH">PATCH</option>
                        <option value="DELETE">DELETE</option>
                        <option value="HEAD">HEAD</option>
                        <option value="OPTIONS">OPTIONS</option>
                    </select>
                </div>

                <div class="mb-3">
                    <label for="http-path-headers">HTTP Headers</label>
                    <textarea class="form-control" bind:value={httpHeaders} id="http-path-headers" style="height: 10em;"></textarea>
                </div>

                <div class="mb-3">
                    <label for="http-path-body">Response Body</label>
                    <textarea bind:value={httpBody} class="form-control" id="http-path-body" placeholder="&lt;script&gt;alert(1)&lt;/script&gt;" style="height: 15em;"></textarea>
                </div>
            </form>
      
            <div class="">
                <button type="button" on:click={DeletePathClick} class="btn btn-md btn-danger">Delete</button>
                <button type="button" on:click={onAddPathClick} class="btn btn-success">Save changes</button>
            </div>
    </div>
    <div class="col-md-6">
        <h4>Event Log</h4>
    </div>
</div>