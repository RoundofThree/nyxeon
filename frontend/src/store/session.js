import { reactive, toRefs } from "vue"

const state = reactive({
  loggedIn: false 
})

export default function getState() {
  // set logged status 
  const setLoggedIn = (val) => {
    state.loggedIn = val 
  }
  return {
    ...toRefs(state), // convert to refs when returning
    setLoggedIn
  }
}