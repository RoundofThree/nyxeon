<template>
<div class="w-full flex h-screen bg-gray-200 justify-center">
  <div class="grid bg-white rounded-lg shadow-xl w-11/12 md:w-9/12 lg:w-1/2 py-4 my-10">

    <div class="flex justify-center">
      <div class="flex">
        <h1 class="text-gray-600 font-bold md:text-2xl text-xl">Register your quest</h1>
      </div>
    </div>
    <form v-on:submit.prevent="submitQuest">
    <div class="grid grid-cols-1 mt-5 mx-7">
      <label for="categories" class="uppercase md:text-sm text-xs text-gray-500 text-light font-semibold">Categories</label>
      <vue-tags-input class="py-2 px-3 rounded-lg border-2 border-purple-300 mt-1 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" element-id="categories" v-model="form.categories" :existing-tags="[
        { key: 1, name: 'Artificial Intelligence'},
        { key: 2, name: 'Microservice'}
      ]" id-field="key" text-field="name" @tag-added="onTagAdded" @tag-removed="onTagRemoved" :typeahead="true" />
    </div>

    <div class="grid grid-cols-1 mt-5 mx-7">
      <label for="content" class="uppercase md:text-sm text-xs text-gray-500 text-light font-semibold">Content</label>
      <textarea id="content" v-model="form.content" style="resize: none; height:auto;" rows="8" class="form-textarea h-24 py-2 px-3 rounded-lg border-2 border-purple-300 mt-1 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" type="text" placeholder="Summarize what you have learnt/done." />
    </div>

    <div class='flex items-center justify-center  md:gap-8 gap-4 pt-5 pb-5'>
      <button class='w-auto bg-gray-500 hover:bg-gray-700 rounded-lg shadow-xl font-medium text-white px-4 py-2'>Create</button>
    </div>
    </form>
  </div>
</div>
</template>
<script>
// if the fetchQuests request fails, then redirect to / 
import axios from "axios"
import VoerroTagsInput from '@voerro/vue-tagsinput'
import getState from '../store/session'
const API = "http://localhost:8080/quests/new"
// const API = "http://localhost:4444"
const axiosConfig = {
  headers: {
    "Content-Type": 'application/json;charset=UTF-8',
    accept: "application/json"
  },
  data: {},
  withCredentials: true
}
export default {
  components: {
    'vue-tags-input': VoerroTagsInput
  },
  data() {
    return {
      form: {
        content: "",
        categories: []
      }
    }
  },
  async created() {
    const loggedIn = getState().loggedIn
    if (!loggedIn) {
      this.$router.push({name: "Home"})
    }
  },
  methods: {
    submitQuest() {
      let that = this 
      if (this.form.content == "" || this.form.categories.length == 0) {
        alert("Invalid input")
        return 
      }
      const data = new FormData()
      data.append('content', this.form.content)
      this.form.categories.forEach(tag => data.append('categories[]', tag))
      axios.post(API, data, axiosConfig).then(res => {
        res.data 
        // clear form 
        that.form.content = ""
        that.form.categories = []
        this.$router.push({name: "Dashboard"})
      }).catch(e => {
        console.log(e)
      })
    },
    onTagAdded(tag) {
      this.form.categories.push(tag.name)
    },
    onTagRemoved(tag) {
      this.form.categories = this.form.categories.filter(e => e !== tag.name)
    }
  }
}

</script>

<style scoped>
vue-tags-input {
  border-color: white !important;
}
</style>