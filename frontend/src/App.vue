<template>
  <v-app>
    <v-app-bar app absolute>
      <app-nav></app-nav>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>  
    </v-main>
    <v-footer app>
      Copyright RoundofThree
    </v-footer>
  </v-app>
</template>

<script>
import AppNav from "@/components/AppNav.vue"

export default {
  components: {
    AppNav  
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>