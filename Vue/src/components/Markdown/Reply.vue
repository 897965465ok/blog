<template>
   <el-row class="comment-list">
      <el-row class="icon-wrapper">
        <span>#Jiang</span>
      </el-row>
      <el-row class="comment-body">
        <el-col class="image-wrapper" :span="4">
          <img src="@/assets/jerry.png" alt="" />
        </el-col>
        <el-col class="content" :span="20">
          <p class="content-text">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Cumque
            tenetur nihil consequatur numquam mollitia perspiciatis possimus
            minus ipsa est dolorum? Expedita voluptate vel quisquam, aperiam
            laudantium sint atque ex officia.
          </p>
        </el-col>
      </el-row>
    </el-row>
</template>
<script>
export default {
  name:"replay",
  data(){
    return {
 
    }
  }
}
</script>
<style lang="scss" scoped>
.comment-list {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  background: #ffffff;
  .comment-body {
    height: 100%;
      display: flex;
        justify-content: center;
        align-items: center;
    .image-wrapper {
      height: 100%;
      width: 120px;
      margin: 0px 16px;
      img {
        height: 100%;
        width: 100%;
        border-radius: 100px;
      }
    }
    .content {
      height: 100%;
      .content-text {
        text-align: left;
        
      }
    }
  }
  .icon-wrapper {
    height: 20px;
  }
}
</style>