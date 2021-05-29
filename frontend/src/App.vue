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
import getState from './store/session'

const state = getState().loggedIn

export default {
  data() {
    return {
      loggedIn: state
    }
  },
  components: {
    AppNav
  }
}
</script>

<style>
</style>