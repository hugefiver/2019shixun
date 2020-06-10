import Vue from 'vue'
import Router from 'vue-router'
//import Home from './views/Home.vue'

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    /*{
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('./views/About.vue') //webpackChunkName: "about"
    }*/
    {
      path: "/me",
      name: "me",
      component: () => import('./views/PersonInfo.vue')
    },
    {
      path: "/mochines",
      name: "mochines",
      component: () => import('./views/MochineCheck.vue')
    },
    {
      path: "/borrow",
      name: "borrow",
      component: () => import('./views/MochineBorrow.vue')
    },
    {
      path: "/admin/mochines",
      name: "mochine_admin",
      component: () => import('./views/MochineAdmin.vue')
    },
    {
      path: "*",
      redirect: "/me"
    }
  ]
})
