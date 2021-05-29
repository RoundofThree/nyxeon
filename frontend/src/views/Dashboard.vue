<template>
  <div class="container mx-auto px-4 py-4 text-center">
    <h1>You are logged in using OAuth2!</h1>
    <!--  Heatmap -->
    <heat-map></heat-map>
    <ul>
      <li v-for="quest in quests" :key="quest._id">
        <span>{{ quest._id }}</span>
        <span>{{ quest.content }}</span>
      </li>
    </ul>
  </div>
</template>
<script>
// if the fetchQuests request fails, then redirect to / 
import axios from "axios"
import HeatMap from '../components/HeatMap.vue'
import getState from '../store/session'

const API = "http://localhost:8080/quests/all"
const axiosConfig = {
  headers: {
    "Content-Type": "application/json"
  },
  withCredentials: true
}
export default {
  components: { HeatMap },
  data() {
    return {
      quests: []
    }
  },
  async created() {
    const setLoggedIn = getState().setLoggedIn
    axios.get(API, axiosConfig).then(res => {
      setLoggedIn(true)
      this.quests = res.data 
    }).catch(err => {
      console.log(err) 
      setLoggedIn(false)
      this.$router.push({name: "Home"})
    })
  }
}

</script>