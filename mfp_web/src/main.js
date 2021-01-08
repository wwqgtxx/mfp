// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
//import echarts from 'echarts'
import ElementUI from 'element-ui' //新添加
import ECharts from 'vue-echarts'
import echarts from 'echarts' 
import 'echarts/lib/chart/bar'
import 'echarts/lib/component/tooltip'
import 'element-ui/lib/theme-chalk/index.css' //新添加，避免后期打包样式不同，要放在import App from './App';之前
import Vue from 'vue'
import App from './App'
import router from './router'

Vue.prototype.$echarts = echarts
Vue.component('v-chart', ECharts)
Vue.use(router) // 引入路由
Vue.use(ElementUI)   //新添加
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h=> h(App)
})
