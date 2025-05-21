<script setup>
import { Edit, HomeFilled, Message, Notebook, Pointer, UserFilled } from '@element-plus/icons-vue';
import { inputEmits, SIZE_INJECTION_KEY } from 'element-plus';
import {ref} from 'vue';
import '../assets/style/Home.css';

// =========================顶部导航栏=====================================================

// 文章卡片的间隔
const article_gap = ref(20);
// 页面变动时的处理函数
const handleCurrentChange = (val) => {
  console.log(`current page: ${val}`)
}
// todo：编写文章数组，完成文章显示逻辑---完成
// 文章包括：标题、摘要、标签、日期、阅读量、阅读时长
// 若设置每页十篇文章，那么文章数组大小为10，每次点击翻页按钮都更新文章数组
const article_list = ref([
  {
    id:1,
    title: 'Vue Router 4.x的使用',
    digest: 'Vue Router 4.x是Vue.js的官方路由管理器，支持动态路由、嵌套路由和路由守卫等功能。通过配置路由表，可以轻松实现单页面应用的路由跳转和组件渲染。',
    tags: ['Vue', 'Router'],
    date: '2025-05-20',
    read_count: 100,
    read_time: '5min',
  },
  {
    id:2,
    title: '深入理解JavaScript的闭包',
    digest: '闭包是JavaScript中的一个重要概念，它允许函数访问其外部作用域的变量。通过闭包，可以实现数据的封装和私有化，提高代码的可维护性和可读性。',
    tags: ['JavaScript', 'Closure'],
    date: '2025-05-20',
    read_count: 200,
    read_time: '10min',
  },
  {
    id:3,
    title: '使用Element Plus构建响应式布局',
    digest: 'Element Plus是Vue.js的UI组件库，提供了丰富的组件和样式。通过使用Element Plus，可以快速构建响应式布局，提高开发效率。',
    tags: ['Element Plus', 'Layout'],
    date: '2025-05-22',
    read_count: 150,
    read_time: '8min',
  },
  {
    id:4,
    title: 'Vue3的响应式原理',
    digest: 'Vue3的响应式原理是基于Proxy和Reflect实现的，使用了getter和setter来追踪数据的变化。通过依赖收集和触发更新机制，实现了高效的响应式数据绑定。',
    tags: ['Vue3', 'Reactive'],
    date: '2025-05-17',
    read_count: 100,
    read_time: '5min',
  },
  {
    id:5,
    title: '深入理解JavaScript的异步编程',
    digest: 'JavaScript的异步编程模型是基于事件循环和回调函数实现的。通过使用Promise、async/await等语法，可以更方便地处理异步操作，提高代码的可读性和可维护性。',
    tags: ['JavaScript', 'Async'],
    date: '2025-05-18',
    read_count: 120,
    read_time: '6min',
  },
  {
    id:6,
    title: '使用Vuex管理Vue3的状态',
    digest: 'Vuex是Vue.js的状态管理库，提供了集中式的状态管理和数据流动机制。通过使用Vuex，可以更方便地管理应用的状态，提高代码的可维护性和可读性。',
    tags: ['Vuex', 'State Management'],
    date: '2025-05-19',
    read_count: 80,
    read_time: '4min',
  },
  {
    id:7,
    title: '使用Axios进行HTTP请求',
    digest: 'Axios是一个基于Promise的HTTP客户端，可以用于浏览器和Node.js。通过使用Axios，可以方便地发送HTTP请求和处理响应，提高代码的可读性和可维护性。',
    tags: ['Axios', 'HTTP'],
    date: '2025-05-21',
    read_count: 90,
    read_time: '4min',
  },
  {
    id:8,
    title: '使用Vue Router实现路由懒加载',
    digest: 'Vue Router支持路由懒加载，可以通过动态导入的方式实现按需加载，提高应用的性能和用户体验。',
    tags: ['Vue Router', 'Lazy Loading'],
    date: '2025-05-23',
    read_count: 110,
    read_time: '5min',
  },
  {
    id:9,
    title: '使用Element Plus实现表单验证',
    digest: 'Element Plus提供了丰富的表单组件和验证规则，可以方便地实现表单的输入验证和数据校验，提高用户体验和数据的准确性。',
    tags: ['Element Plus', 'Form Validation'],
    date: '2025-05-24',
    read_count: 130,
    read_time: '7min',
  },
  {
    id:10,
    title: '使用Vue3实现组件间通信',
    digest: 'Vue3提供了多种组件间通信的方式，包括props、emit、provide/inject等。通过合理地使用这些方式，可以提高组件的复用性和可维护性。',
    tags: ['Vue3', 'Component Communication'],
    date: '2025-05-25',
    read_count: 140,
    read_time: '6min',
  },
  {
    id:11,
    title: '使用Vue3实现自定义指令',
    digest: 'Vue3支持自定义指令，可以通过定义指令的钩子函数，实现对DOM元素的操作和行为的扩展，提高代码的复用性和可维护性。',
    tags: ['Vue3', 'Custom Directive'],
    date: '2025-05-26',
    read_count: 150,
    read_time: '7min',
  },
  {
    id:12,
    title: '使用Vue3实现动画效果',
    digest: 'Vue3提供了丰富的过渡和动画效果，可以通过使用transition组件和CSS动画，实现页面元素的动态效果，提高用户体验。',
    tags: ['Vue3', 'Animation'],
    date: '2025-05-27',
    read_count: 160,
    read_time: '8min',
  },
    
]);


</script>

<template>
   

<el-space direction="vertical" class="articles_container" :size="article_gap" fill="true">
  <!-- 卡片上需要包括:标题(链接)/分类(链接)/上传日期(简单显示)/配图 -->
    <!-- 后续通过从后端获取一个文章内容数组来渲染卡片内容 -->
     <!-- todo：点击标题跳转文章阅读界面 -->
  <el-card v-for="article in article_list" :key="article.id" class="article_card">
    <template #header>
      <div>
        <el-row class="article_card_header" justify="space-between">
          <!-- 标题 -->
          <el-col class="article_title" span="12">
            <el-link href="" class="title" target="_self">
              {{ article.title }}
            </el-link>
          </el-col>
          <!-- 日期 -->
          <el-col class="date" span="12">
            <el-icon><Calendar /></el-icon>{{ article.date }}
          </el-col>
        </el-row>
        
      </div>
    </template>

    <div class="article_card_body">
      <p class="digest">
        {{article.digest}}
      </p>
    </div>

    <template #footer>
      <div>
        <el-row class="article_card_footer" justify="space-between">
          <!-- 标签 -->
          <el-col class="tag" span="12">
            <el-space>
              <el-icon><CollectionTag /></el-icon>
              <el-tag v-for="tag in article.tags" :key="tag">{{ tag }}</el-tag>
            </el-space>
          </el-col>
          <!-- 阅读时长 -->
          <el-col class="words_timespend" span="12">
            <el-icon><Clock /></el-icon>{{article.read_time}}
          </el-col>
          <!-- 阅读量 -->
          <el-col class="read_count" span="12">
            <el-icon><Pointer /></el-icon>{{ article.read_count }}次阅读
          </el-col>
        </el-row>
      </div>
    </template>

  </el-card>

</el-space>
    

    <!-- 分页切换按钮部分 -->
<div>
  <el-card class="page_change_card">
    <!-- page-size:每页内容数 pager-count:显示的页面按钮数 total:总内容条数-->
    <!-- todo: 通过设置page-change事件来翻页（切换文章数组） -->
    <!-- todo: 后面从数据库读到所有文章数后传给total -->
    <el-pagination 
    class="page_change"
    background 
    layout="prev, pager, next, ->, total, jumper" 
    :total="100"
    :page-size="10"
    :pager-count="11"
    size="small"
    @current-change="handleCurrentChange"
    />
  </el-card>
</div>



</template>

