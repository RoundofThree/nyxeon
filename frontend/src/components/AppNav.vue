<template>
  <v-toolbar>
    <v-toolbar-title>
      <router-link to="/" tag="span" style="cursor: pointer">{{ appTitle }}</router-link>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items class="hidden-xs-only">
      <v-btn flat v-for="item in menuItems" :key="item.title" :to="item.path">
        {{ item.title }}
      </v-btn>
    </v-toolbar-items>
  </v-toolbar>
</template>

<script>
export default {
  name: "AppNav",
  data() {
    return {
      appTitle: "Nyxeon",
      menuItems: [
        { title: "Home", path: "/" },
        { title: "About", path: "/about" }
      ]  
    }
  }
}
</script>