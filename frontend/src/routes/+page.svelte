<script lang="ts">
    import { base } from '$app/paths'
    import type {Path, LogEvent} from "./../models"
    import {AddDefaultPath, GetDefaultPaths, GetEvents} from "../actions/DefaultPathAction.svelte"
    import { onMount, onDestroy } from "svelte";
    import {defaultPaths} from "./../store"
    import Events from "../components/Events/Events.svelte";


    let path: string = "";
    let httpVerb: string = "GET"; // default http verb
    let httpHeaders: string = "";
    let httpBody: string = "";

    $: path = path.replace(" ", "-");

    $: if(path != "" && !path.startsWith("/")) {
        path = "/" + path;
    }

    onMount(async () => {
        let paths: Path[] = await GetDefaultPaths();
        $defaultPaths.DefaultPaths = paths;
    })


    function prepHeaders(): Path {
        const splitHeaders = httpHeaders.split("\n");
        return {
            Id: "",
            Path: path,
            Verb: httpVerb,
            Headers: splitHeaders,
            Body: btoa(httpBody),
        }
    }

    async function onAddPathClick(){
        const p = prepHeaders()
        let newPath = await AddDefaultPath(p);
        if(Object.keys(newPath).length > 0){
            $defaultPaths.DefaultPaths = [...$defaultPaths.DefaultPaths, newPath]
            console.log($defaultPaths.DefaultPaths);
        }
    }

    function deletePathFromState(p: Path) {
        // this does not trigger re-render for some reason
        // const index = $defaultPaths.DefaultPaths.indexOf(p, 0)
        // delete $defaultPaths.DefaultPaths[index]
        $defaultPaths.DefaultPaths = $defaultPaths.DefaultPaths.filter(path => path != p)

    }

    async function DeletePathClick(p: Path) {
        let result = await AddDefaultPath(p, "/api/defaulthttp/" + p.Id, "DELETE");
        if(Object.keys(result).length > 0){
            deletePathFromState(p)
        }
    }

</script>

<svelte:head>
	<title>KOLLAB | Dashboard</title>
</svelte:head>

    <!-- Modal -->
    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-xl">
        <div class="modal-content">
            <div class="modal-header">
            <h1 class="modal-title fs-5" id="exampleModalLabel">New Path</h1>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
                <form>
                    <div class="mb-3">
                        <label for="http-path">Path</label>
                        <input type="text" bind:value={path} placeholder="/some/path/payload.js" class="form-control" id="http-path">
                    </div>
                    
                    <div class="mb-3">
                        <label for="floatingSelect">Select HTTP Verb</label>
                        <select required bind:value={httpVerb} class="form-select" id="http-path-verb" aria-label="Select HTTP verb">
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
                        <textarea bind:value={httpHeaders} class="form-control" id="http-path-headers" placeholder="X-API: 12345
Set-Cookie: loggedin=true; Domain=example.com; Path=/" style="height: 10em;"></textarea>
                    </div>

                    <div class="mb-3">
                        <label for="http-path-body">Response Body</label>
                        <textarea bind:value={httpBody} class="form-control" id="http-path-body" placeholder="&lt;script&gt;alert(1)&lt;/script&gt;" style="height: 15em;"></textarea>
                    </div>
                </form>
                </div>
                <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                <button type="button" on:click={onAddPathClick} class="btn btn-primary">Save changes</button>
            </div>
        </div>
        </div>
    </div>
  

    <div class="row">
        <div class="col-md">
            <h3>Paths <button type="button" class="btn btn-light btn-sm" aria-label="New" data-bs-toggle="modal" data-bs-target="#exampleModal">New Path</button></h3>
            
            <table class="table">
                <thead>
                  <tr>
                    <th scope="col">Endpoint</th>
                    <th scope="col">Verb</th>
                    <th scope="col">Action</th>
                  </tr>
                </thead>
                <tbody id="path-table">
                    {#each $defaultPaths.DefaultPaths as item, i}
                        <tr>
                            <td>{item.Path}</td>
                            <td>{item.Verb}</td>
                            <td><a href="{base}/path/{item.Id}"><button type="button" class="btn btn-primary btn-sm">View</button></a>
                                <button type="button" on:click={() => DeletePathClick(item)} class="btn btn-sm btn-danger">Delete</button>
                            </td>
                        </tr>
                    {/each}
                </tbody>
              </table>
        </div>

    </div>

    <br />
    <br />
    <br />
    <br />


    <div class="row">
        <div class="col">
            <h3>Event Log</h3>
            <Events/>
        </div>
    </div>

