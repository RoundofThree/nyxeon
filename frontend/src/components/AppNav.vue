<template>
  <nav
    class="z-50 w-full flex flex-wrap items-center justify-between px-2 py-3 bg-black"
  >
    <div
      class="container px-4 mx-auto flex flex-wrap items-center justify-between"
    >
      <div
        class="w-full relative flex justify-between lg:w-auto lg:static lg:block lg:justify-start"
      >
        <router-link
          class="text-sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap uppercase text-white"
          to="/"
          >Nyxeon</router-link
        >
      </div>
      <div
        class="lg:flex flex-grow items-center bg-white lg:bg-transparent lg:shadow-none"
        v-bind:class="{'hidden': !showMenu, 'block': showMenu}"
      >
        <ul class="flex flex-col lg:flex-row list-none lg:ml-auto">
          <li class="flex items-center">
            <router-link
              class="lg:text-white lg:hover:text-gray-300 text-gray-800 px-3 py-4 lg:py-2 flex items-center text-xs uppercase font-bold"
              to="/about"
              ><span class="inline-block ml-2">About</span></router-link
            >
          </li>
          <li class="flex items-center">
            <router-link
              class="lg:text-white lg:hover:text-gray-300 text-gray-800 px-3 py-4 lg:py-2 flex items-center text-xs uppercase font-bold"
              to="/dashboard"
              ><span class="inline-block ml-2">Dashboard (test only)</span></router-link
            >
          </li>
          <li class="flex items-center">
            <router-link
              class="lg:text-white lg:hover:text-gray-300 text-gray-800 px-3 py-4 lg:py-2 flex items-center text-xs uppercase font-bold"
              to="/quests/new"
              ><span class="inline-block ml-2">Register quest</span></router-link
            >
          </li>
          <li class="flex items-center">
            <button
              class="lg:text-white lg:hover:text-gray-300 text-gray-800 px-3 py-4 lg:py-2 flex items-center text-xs uppercase font-bold"
              @click="logout"
              ><span class="inline-block ml-2">Logout</span></button
            >
          </li>
          <li class="flex items-center">
            <button
              class="bg-white text-gray-800 active:bg-gray-100 text-xs font-bold uppercase px-4 py-2 rounded shadow hover:shadow-md outline-none focus:outline-none lg:mr-1 lg:mb-0 ml-3 mb-3"
              type="button"
              style="transition: all 0.15s ease 0s;"
            >
              View source code
            </button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script>
import axios from "axios"
const API = "http://localhost:8080/oauth/logout"
const axiosConfig = {
  headers: {
    "Content-Type": "application/json"
  },
  withCredentials: true
}

export default {
  data() {
    return {
      showMenu: false
    }
  },
  methods: {
    toggleNavbar: function(){
      this.showMenu = !this.showMenu;
    },
    logout: function() {
      axios.delete(API, axiosConfig).then(res => {
        // server must delete session cookie
        // redirect to / 
        console.log(res.data)
        this.$router.push({name: "Home"})
      })
    }
  }
}
</script>
