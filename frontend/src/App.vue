<template>
  <div>
    <!-- Navbar -->
    <app-nav></app-nav>
    <!-- Main -->
    <div class="flex w-full">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import AppNav from './components/AppNav.vue'

export default {
  components: {
    AppNav
  }
}
</script>

<style>
</style>