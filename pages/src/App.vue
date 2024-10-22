<template>
  <component :is="layoutComponent">
    <router-view />
  </component>
</template>

<script>
import BlankLayout from './layouts/BlankLayout.vue';
import MainLayout from './layouts/MainLayout.vue';


export default {
  computed: {
    layoutComponent() {
      const layout = this.$route.meta.layout || 'MainLayout';
      return layout === 'BlankLayout' ? BlankLayout : MainLayout;
    }
  }
}
</script>
