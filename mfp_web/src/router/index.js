import Vue from 'vue'
import Router from 'vue-router'
import Queue from '../components/Queue'
import Keywords from '../components/Keywords'
import Url from '../components/Url'
import Time from '../components/Time'

Vue.use(Router)

export default new Router({
  model: 'history',
  routes: [
    {
      path: '/',
      redirect: {
        path: '/Queue'
      }
    },
    {
      path: '/Queue',
      component: Queue,
    },
    {
      path: '/Keywords',
      component: Keywords,
    },

    {
      path: '/Url',
      component: Url,
    },
    {
      path: '/Time',
      component: Time,
    },
  ]
})
