<script setup>
import { Edit, HomeFilled, Message, Notebook, UserFilled } from '@element-plus/icons-vue';
import { inputEmits, SIZE_INJECTION_KEY } from 'element-plus';
import {ref} from 'vue';
import './style.css';

// 主页默认激活页面的index
const activeIndex = ref('1');

// 搜索框输入值
const input = ref('');



</script>
<template>
    <el-container class="first">
      <!-- 顶部横幅 -->
      <el-header class="header">
        <el-row class="header_row"  gutter="16">
          <!-- 导航栏选项 -->
          <el-col class="menu_col" :span="16">
          <el-menu 
          :default-active="activeIndex" 
          mode="horizontal" 
          class="menu" 
          :ellipsis="false"
          >
            <!-- 常规选项 -->
            <el-menu-item index="1"><router-link to="/home" class="menu_link"><el-icon><House /></el-icon>Home</router-link></el-menu-item>
            <el-menu-item index="2"><router-link to="/history" class="menu_link"><el-icon><Notebook /></el-icon>Log</router-link></el-menu-item>
            <el-menu-item index="3"><a href="https://github.com/zhenqiiii/BlogSite" target="_blank" class="menu_link"><el-icon><Link /></el-icon>Github</a></el-menu-item>
          </el-menu>
          </el-col>
          <!-- 搜索框 -->
          <el-col class="search_col" :span="8">
          <el-input v-model="input" placeholder="知无不言 言无不尽" clearable class="searchbox">
            <template #prefix>
              <el-icon class="el-input__icon"><search /></el-icon>
            </template>
          </el-input>
          </el-col>
        </el-row>
      </el-header>
      <el-container class="sub1">
        <!-- 主界面 -->
        <el-main class="main">
          <router-view></router-view>
        </el-main>
        <!-- 侧边栏 -->
        <el-aside class="aside" >
          <!-- 使用Layout布局组件 -->
          <!-- 个人介绍卡片 -->
          <el-row justify="center">
            <el-col>
              <el-card shadow="always">
                <el-avatar shape="square" :size="160" fit="fill" src="https://avatars.githubusercontent.com/u/148515933?v=4"/>

      
                <p class="name">zhenqiiii</p>
                <!-- 自己的链接 -->
                  <el-space size="large" spacer="|">
                  <el-button
                  type="default"
                  tag="a"
                  href="https://github.com/zhenqiiii"
                  target="_blank"
                  ><el-icon><Link /></el-icon></el-button>

                  <el-button><el-icon></el-icon></el-button>

                  <el-button
                  type="default"
                  tag="a"
                  href="mailto:zhangwenyang2004@foxmail.com"
                  target="_blank"
                  ><el-icon><Message /></el-icon></el-button>
                  </el-space>
              </el-card> 
            </el-col>
          </el-row>

          <!-- 爱好卡片 -->
            <el-row>
            <el-col>
              <el-card>
                <template #header>
                  <div>
                    <span>Hobby</span>
                  </div>
                </template>
                <el-row justify="space-around">
                  <el-col :span="8">
                    <el-icon><Edit /></el-icon>写代码
                  </el-col>
                  <el-col :span="8">
                    <el-icon><UserFilled /></el-icon>打游戏
                  </el-col>
                  <el-col :span="8">
                    <el-icon><Message /></el-icon>看电影
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </el-row>

            <!-- 分类 -->
          <el-row>
            <el-col>
              <el-card class="category_card">
                <template #header>
                  <div>
                    <span>Tags</span>
                  </div>
                </template>
                <div class="tag_container">
                  <el-row>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/vue3" class="tag_link">Vue3</router-link></el-tag>
                    </el-col>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/react" class="tag_link">React</router-link></el-tag>
                    </el-col>
                    <el-col :span="12">
                      <el-tag size="large"><router-link to="/tag/elementplus" class="tag_link">ElementPlus</router-link></el-tag>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="10">
                      <el-tag size="large"><router-link to="/tag/js" class="tag_link">JavaScript</router-link></el-tag>
                    </el-col>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/go" class="tag_link">Golang</router-link></el-tag>
                    </el-col>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/gin" class="tag_link">Gin</router-link></el-tag>
                    </el-col>
                  </el-row>
                  <el-row>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/auth" class="tag_link">Auth</router-link></el-tag>
                    </el-col>
                    <el-col :span="6">
                      <el-tag size="large"><router-link to="/tag/git" class="tag_link">Git</router-link></el-tag>
                    </el-col>
                    <el-col :span="10">
                      <el-tag size="large"><router-link to="/tag/docker" class="tag_link">Docker</router-link></el-tag>
                    </el-col>
                  </el-row>
                </div>
                
              </el-card>
            </el-col>
          </el-row>

          <!-- 友链 -->
          <el-row>
            <el-col>
              <el-card>
                <template #header>
                  <div>
                    <span>My Fellows</span>
                  </div>
                </template>
              </el-card>
            </el-col>
          </el-row>

          <!-- toys -->  
          <el-row>
            <el-col>
              <el-card>
                <div>
                  <el-link underline="never" type="primary" href="https://www.bilibili.com" target="_blank">
                    <el-icon class="is-loading"><ToiletPaper /></el-icon>小网站-欲罢不能，每天都想看
                  </el-link>
                </div>
            </el-card>
            </el-col>
          </el-row>

          <!-- 信息 -->
            <el-row>
            <el-col>
              <el-card shadow="always">
                <div>备案号：xxxxxx</div>
                <div>github链接</div>
              </el-card>
            </el-col>
            </el-row>
          
        </el-aside>
      </el-container>

      

    </el-container>





</template>

<style scoped>

.tag_link{
  text-decoration: none;
  color:inherit;
}


.menu_link{
  text-decoration: none;
  width: 100%;
}
.header{
  /* 导航栏居中效果由padding实现 */
  padding: 0 20% 20% 20%;
  line-height: 60px;
  font-size: large;
}
/* header内部组件 */
.header_row{
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px);
  border-radius: 0 0 20px 20px;
  padding: 0 5px;
  /* opacity: 0.9; */
}
/* 导航栏 */
.menu_col{
  height: 100%;
  /* margin-right: ; */
  /* border: 1px solid gray; */
}
.menu{
  height: 100%;
  width: 100%;
  background-color: rgba(255, 255, 255, 0);
}
.menu .el-menu-item{
  width: 33%;
  padding: 0%;
}
/* 搜索框 */
.search_col{
  /* border: 1px solid gray; */
  height: 100%;
}
.searchbox{
  width: 100%;
}

/* container 容器样式*/
.sub1{
  justify-content: center;
  height: 160vh;
}

/* container 内部组件样式 */
.aside{
  width: 25%;
  min-width: 240px;
  max-width: 400px;
  border-radius: 8px;
  margin: 10px;
  padding: 20px 20px;
}
.main{
  text-align: center;
  border-radius: 8px;
  margin: 10px;
  width: 70%;
  min-width: 800px;
  max-width: 1000px;
}
.main, .aside{
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px);
}



.aside .el-row{
  margin-bottom: 10px;
}
.aside .el-card{
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(10px);
}

.name{
  font-size: 20px;
  text-align: center;
  margin-top: 10px;
  font-weight: bold;
  /* border: 1px solid gray; */
}

/* 侧边栏卡片 */
/* 分类 */
.tag_container{
    font-size: clamp(10px, 1vw, 16px);
}
.tag_container el-tag{
    width: 100%;
   

}
</style>