<template>
  
    <HeaderComponent :title="headerTitle" />

    <div class="page-wrapper">

      <slot></slot>

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
    return 
  }
}
</script>

