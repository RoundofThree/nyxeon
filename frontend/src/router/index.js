import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

/*
const requiresAuthentication = (to, from, next) => {
  // check session token in cookie with backend 
  console.log("Checking session token...")
  let success = document.cookie.match(/^(.*;)?\s*nyx_sess_id\s*=\s*[^;]+(.*)?$/)
  console.log(document.cookie)
  // if cookie not set
  if (success) {
    return next() 
  }
  // failure 
  next('/')
}
*/

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    // beforeEnter: requiresAuthentication,
    component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue')
  },
  {
    path: '/quests/new',
    name: 'NewQuest',
    component: () => import(/* webpackChunkName: "newquest" */ '../views/NewQuest.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
