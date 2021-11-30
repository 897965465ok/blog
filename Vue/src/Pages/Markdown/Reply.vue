<template>
   <el-row>
   <el-row v-for="item in  comments" :key="item.UserID"  class="comment-list" >
      <el-row class="icon-wrapper">
        <span>{{item.UserName}}</span>
      </el-row>
      <el-row class="comment-body">
        <el-col class="image-wrapper" :span="4">
          <img src="@/assets/jerry.png" alt="" />
        </el-col>
        <el-col class="content" :span="20">
          <p class="content-text">
         {{item.Content}}
          </p>
        </el-col>
      </el-row>
    </el-row>
    </el-row>
</template>
<script>
export default {
  name:"replay",
  porps:{
     comments:{
       default:[],
       type:[]
     }
  },
  data(){
    return {
    comments:[]
    }
  },
 async created(){
    let articleId = this.$route.query.uuid
     let {data}  = await this.$getComments(articleId)
     console.log(data)
      if(data.code == 200){
        this.comments = data.result
      }
  },

}
</script>
<style lang="scss" scoped>
.comment-list {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  background: #ffffff;
  margin-top: 10px;
  border-bottom: black solid 1px;
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