<template>
<div class="w-full flex h-screen bg-gray-200 justify-center">
  <div class="grid bg-white rounded-lg shadow-xl w-11/12 md:w-9/12 lg:w-1/2 py-4 my-10">

    <div class="flex justify-center">
      <div class="flex">
        <h1 class="text-gray-600 font-bold md:text-2xl text-xl">Register your quest</h1>
      </div>
    </div>

    <div class="grid grid-cols-1 mt-5 mx-7">
      <label class="uppercase md:text-sm text-xs text-gray-500 text-light font-semibold">Categories</label>
      <input class="py-2 px-3 rounded-lg border-2 border-purple-300 mt-1 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" type="text" placeholder="Insert the categories separated by a single comma" />
    </div>

    <div class="grid grid-cols-1 mt-5 mx-7">
      <label class="uppercase md:text-sm text-xs text-gray-500 text-light font-semibold">Content</label>
      <textarea style="resize: none; height:auto;" rows="8" class="form-textarea h-24 py-2 px-3 rounded-lg border-2 border-purple-300 mt-1 focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" type="text" placeholder="Summarize what you have learnt/done." />
    </div>

    <div class='flex items-center justify-center  md:gap-8 gap-4 pt-5 pb-5'>
      <button @click="submitQuest" class='w-auto bg-gray-500 hover:bg-gray-700 rounded-lg shadow-xl font-medium text-white px-4 py-2'>Create</button>
    </div>

  </div>
</div>
</template>
<script>
// if the fetchQuests request fails, then redirect to / 
import axios from "axios"
const API = "http://localhost:8080/quests/new"
const axiosConfig = {
  headers: {
    "Content-Type": "application/json"
  },
  withCredentials: true
}

const parseForm = () => {
    // SET data from form input 
}
export default {
  data() {
    return {
      content: "",
      categories: []
    }
  },
  methods: {
    submitQuest: () => {
      parseForm()
      axios.post(API, {}, axiosConfig).then(res => {
        // show flash 
        console.log(res)
      }).catch(e => {
        console.log(e)
      })
    }
  }
}

</script>