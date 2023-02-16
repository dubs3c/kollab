<script>

</script>
<div class="container">
    <header class="d-flex flex-wrap justify-content-center py-3 mb-4 border-bottom">
      <a href="/" class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-dark text-decoration-none">
        <span class="fs-4"><strong>KOLLABORATÖR</strong></span>
      </a>

      <ul class="nav nav-pills">
        <li class="nav-item"><a href="#" class="nav-link active" aria-current="page">Home</a></li>
        <li class="nav-item"><a href="#" class="nav-link">Settings</a></li>
        <li class="nav-item"><a href="#" class="nav-link">Logout</a></li>
      </ul>
    </header>


    <div class="row">
        <!-- Begin Toast wrapper -->
        <div aria-live="polite" aria-atomic="true">
            <div class="toast-container top-0 end-0 p-3">
                <div id="toastError" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
                <div class="toast-header text-bg-danger border-0">
                    <strong class="me-auto">ERROR</strong>
                    <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
                </div>
                <div class="toast-body">
                    This is a cool error message for you
                </div>
                </div>
            </div>

            <div class="toast-container top-0 end-0 p-3">
                <div id="toastSuccess" class="toast text-bg-success border-0" role="alert" aria-live="assertive" aria-atomic="true">
                    <div class="d-flex">
                    <div class="toast-body">
                        Hello, world! This is a toast message.
                    </div>
                    <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button>
                    </div>
                </div>
            </div>

        </div>
        <!-- End Toast wrapper -->
    </div>
   
    <div class="page-contents">
        <slot></slot>
    </div>

    <div class="row">
        <footer class="py-3 my-4">
            <div class="nav justify-content-center border-bottom pb-3 mb-3">
              <p></p>
            </div>
            <p class="text-center text-muted">KOLLABORATÖR | Kapten Svartskägg</p>
          </footer>
    </div>
</div>
