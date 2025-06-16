<script setup>
import axios from 'axios';
import { useRoute } from 'vue-router';
// axios测试：2025/5/22
// 现在假设我已经拿到文章id，需要完成的是向后端发送请求并渲染响应
const route = useRoute();
const id = route.params.id;
const getArticle = (id) => {
  return 
}


</script>
<template>
  <el-card class="article_content">
    <template #header>
      <div class="content_header">
        文章卡片头部
      </div>
      <!-- 内容部分 -->
    </template>

  </el-card>
</template>

<style scoped>
.article_content{
  height: 100%;
  width: 100%;
}

</style>