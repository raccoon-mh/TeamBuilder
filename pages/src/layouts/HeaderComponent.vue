<template>
  <div class="sticky-top">
    <header class="navbar navbar-expand-md sticky-top d-print-none">
      <div class="container-xl">
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbar-menu"
          aria-controls="navbar-menu"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <h1 class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
          <a href="index.html">Simple LoL Team Builder</a>
        </h1>

        <div class="navbar-nav flex-row order-md-last">
          <div class="d-none d-md-flex">
            <a href="?theme=dark" class="nav-link px-0 hide-theme-dark" title="Enable dark mode">
              <!-- Dark mode icon -->
            </a>
            <a href="?theme=light" class="nav-link px-0 hide-theme-light" title="Enable light mode">
              <!-- Light mode icon -->
            </a>
          </div>
        </div>
      </div>
    </header>

    <header class="navbar-expand-md">
      <div class="collapse navbar-collapse" id="navbar-menu">
        <div class="navbar">
          <div class="container-xl">
            <ul class="navbar-nav">
              <li class="nav-item">
                <a class="nav-link" href="./">
                  <span class="nav-link-title">Home</span>
                </a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </header>
  </div>
</template>

<script>
export default {
  name: "HeaderComponent",
};
</script>

<style>
  @import '~@tabler/core/dist/css/tabler.min.css?1692870487';
  @import '~@tabler/core/dist/css/tabler-flags.min.css?1692870487';
  @import '~@tabler/core/dist/css/tabler-payments.min.css?1692870487';
  @import '~@tabler/core/dist/css/tabler-vendors.min.css?1692870487';
  @import '~@tabler/core/dist/css/demo.min.css?1692870487';
</style>
