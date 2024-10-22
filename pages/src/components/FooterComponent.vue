<template>

  <footer class="footer footer-transparent d-print-none">
    <div class="container-xl">
      <div class="row text-center align-items-center flex-row-reverse">
        <div class="col-lg-auto ms-lg-auto">
          <ul class="list-inline list-inline-dots mb-0">
            <li class="list-inline-item"><router-link class="link-secondary" to="/license">License</router-link></li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
  
</template>

<script>
export default {
  name: "FooterComponent",
};
</script>
