<template>
  <div>
    <HeaderComponent :title="headerTitle" />
    <main>
      <slot />
    </main>
    <FooterComponent :message="footerMessage" />
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
      footerMessage: '© 2024 Your Company'
    };
  }
}
</script>

<style scoped>
main {
  padding: 20px;
}
</style>
