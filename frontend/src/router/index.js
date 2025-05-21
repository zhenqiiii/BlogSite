import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Article from '../views/Article.vue'
import History from '../views/History.vue'
import Tag from '../views/Tag.vue'
import Admin from '../views/Admin.vue'
import Management from '../views/Management.vue'
import Upload from '../views/Upload.vue'

// 路由实例
const router = createRouter(
  {
    // 历史模式
    history: createWebHistory(),
    // 路由规则
    routes: [
      {path: '/', redirect: '/home'},
      {path: '/home', component: Home},
      {path: '/history', component: History},
      {path: '/article/:id', component: Article},
      {path: '/tag/:tag', component: Tag},
      {
        path: '/admin', 
        component: Admin,
        children: [
          {
            // 文章管理
            path: 'management',
            component: Management,
          },
          {
            // 上传
            path: 'upload',
            component: Upload,
          },
        ],
      },
    ]
  }
)

export default router
// 这里是路由的配置文件，主要是定义路由的规则和路由的组件