<template>
  
    <HeaderComponent :title="headerTitle" />

    <div class="page-wrapper">

      <div class="page-body">
        <div class="container-xl">
          <slot></slot>
        </div>
      </div>

      <FooterComponent/>
    </div>

</template>

<script>
import FooterComponent from '../layouts/FooterComponent.vue';
import HeaderComponent from '../layouts/HeaderComponent.vue';

export default {
  name: 'MainLayout',
  components: {
    HeaderComponent,
    FooterComponent
  },
  data() {
    return {
      headerTitle: 'My Vue Application',
    };
  }
}
</script>

