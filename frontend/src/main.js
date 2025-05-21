import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
// 给index.html使用的样式表
import './style.css'
import'./assets/style/global.css'
import App from './App.vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router'


const app = createApp(App)
//  全局注册elementplus的图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.use(ElementPlus)
app.use(router).mount('#app')
