<template>
  
    <HeaderComponent/>

    <div class="page-wrapper">

      <slot></slot>

      <FooterComponent/>
      
    </div>

</template>

<script>
import FooterComponent from '../components/FooterComponent.vue';
import HeaderComponent from '../components/HeaderComponent.vue';

export default {
  name: 'MainLayout',
  components: {
    HeaderComponent,
    FooterComponent
  }
}
</script>

