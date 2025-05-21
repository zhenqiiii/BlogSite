<script setup>

    // alert($route.params.id);
</script>
<template>
    <el-container class="sub1">
    <!-- 文章内容 -->
    <el-main class="main">
    <div>

    </div>

    </el-main>

    <!-- ---------------------------------------------------------------------------------------------------- -->

    <!-- 右侧边栏 -->
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
                <span>Category</span>
              </div>
            </template>
            <div class="tag_container">
              <el-row>
                <el-col :span="6">
                  <el-tag size="large">Vue3</el-tag>
                </el-col>
                <el-col :span="6">
                  <el-tag size="large">React</el-tag>
                </el-col>
                <el-col :span="12">
                  <el-tag size="large">ElementPlus</el-tag>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="10">
                  <el-tag size="large">JavaScript</el-tag>
                </el-col>
                <el-col :span="4">
                  <el-tag size="large">Go</el-tag>
                </el-col>
                <el-col :span="6">
                  <el-tag size="large">Gin</el-tag>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="6">
                  <el-tag size="large">auth</el-tag>
                </el-col>
                <el-col :span="6">
                  <el-tag size="large">Git</el-tag>
                </el-col>
                <el-col :span="10">
                  <el-tag size="large">Docker</el-tag>
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
</template>