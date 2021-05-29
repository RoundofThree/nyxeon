<template>
  <div class="container mx-auto px-4 py-4 text-center">
    <!--  Heatmap -->
    <template v-if="ready">
      <heat-map :entriesIn="quests"></heat-map>
    </template>
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
      ready: false,
      quests: []  // array of {_id, content, categories, date}
    }
  },
  async created() {
    const setLoggedIn = getState().setLoggedIn
    axios.get(API, axiosConfig).then(res => {
      setLoggedIn(true)
      this.quests = res.data 
      this.ready = true
      console.log(res.data)
    }).catch(err => {
      console.log(err) 
      setLoggedIn(false)
      this.$router.push({name: "Home"})
    })
  }
}

</script>