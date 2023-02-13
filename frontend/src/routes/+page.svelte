<script lang="ts">

    function onAddPathClick(){

    }

</script>

<svelte:head>
	<title>Kollab | Dashboard </title>
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
                        <input type="text" placeholder="/some/path/payload.js" class="form-control" id="http-path">
                    </div>
                    
                    <div class="mb-3">
                        <label for="floatingSelect">Select HTTP Verb</label>
                        <select class="form-select" id="http-path-verb" aria-label="Select HTTP verb">
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
                        <textarea class="form-control" id="http-path-headers" placeholder="X-API: 12345
Set-Cookie: loggedin=true; Domain=example.com; Path=/" style="height: 10em;"></textarea>
                    </div>

                    <div class="mb-3">
                        <label for="http-path-body">Response Body</label>
                        <textarea class="form-control" id="http-path-body" placeholder="&lt;script&gt;alert(1)&lt;/script&gt;" style="height: 15em;"></textarea>
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
                  <tr>
                    <td>/x.js</td>
                    <td>GET</td>
                    <td><button type="button" class="btn btn-danger btn-sm">Delete</button>
                        <button type="button" class="btn btn-primary btn-sm">View</button>
                    </td>
                  </tr>
                  <tr>
                    <td>/Authenticate/login</td>
                    <td>POST</td>
                    <td><button type="button" class="btn btn-danger btn-sm">Delete</button>
                        <button type="button" class="btn btn-primary btn-sm">View</button>
                    </td>
                  </tr>
                </tbody>
              </table>
        </div>
    </div>

    <br />
    <br />


    <div class="row">
        <div class="col">
            <h3>HTTP Servers <button type="button" class="btn btn-light btn-sm" aria-label="New" data-bs-toggle="modal" data-bs-target="#exampleModal">New HTTP Server</button></h3>
            <hr />
        </div>

        <div class="col">
            <h3>TCP Servers <button type="button" class="btn btn-light btn-sm" aria-label="New" data-bs-toggle="modal" data-bs-target="#exampleModal">New TCP Server</button></h3>
            <hr />
        </div>

        <div class="col">
            <h3>UDP Servers <button type="button" class="btn btn-light btn-sm" aria-label="New" data-bs-toggle="modal" data-bs-target="#exampleModal">New UDP Server</button></h3>
            <hr />
        </div>

    </div>

