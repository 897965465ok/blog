<template>
  <div class="container  w-full  px-4 py-4  shadow    bg-white ">
    <div
      class=" flex flex-col   h-24   shadow-inner  my-4  "
      v-for="item in comments"
      :key="item.uuid"
    >
      <div class="flex   h-full  justify-start text-left my-4 py-3  ">
        <div class=" h-full flex justify-center  content-center">
          <img class="block h-full " src="@/assets/jerry.png" />
        </div>
        <div class=" flex mx-3  flex-grow  justify-between  flex-col  ">
          <span class="text-gray-500  text-xl  font-serif ">{{
            item.User[0].name
          }}</span>
          <p class=" text-gray-600  text-2xl">
            {{ item.Content }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "replay",
  porps: {
    comments: {
      default: [],
      type: []
    }
  },
  data() {
    return {
      comments: []
    };
  },
  async created() {
    let articleId = this.$route.query.uuid;
    console.log(articleId);
    let { data } = await this.$getComments(articleId);
    console.log(data);
    if (data.code == 200) {
      this.comments = data.result;
    }
  }
};
</script>
<style lang="scss" scoped>
// .comment-list {
//   display: flex;
//   flex-direction: column;
//   box-sizing: border-box;
//   background: #ffffff;
//   margin-top: 10px;
//   border-bottom: black solid 1px;
//   .comment-body {
//     height: 100%;
//     display: flex;
//     justify-content: center;
//     align-items: center;
//     .image-wrapper {
//       height: 100%;
//       width: 120px;
//       margin: 0px 16px;
//       img {
//         height: 100%;
//         width: 100%;
//         border-radius: 100px;
//       }
//     }
//     .content {
//       height: 100%;
//       .content-text {
//         text-align: left;
//       }
//     }
//   }
//   .icon-wrapper {
//     height: 20px;
//   }
// }
</style>
