<template>
  <div class="container   h-auto w-full  px-4    shadow    bg-white ">
    <div
      class=" flex flex-col    shadow-inner  my-4  "
      v-for="(item, index) in comments"
      :key="item.uuid"
    >
      <div class="flex  h-24  justify-start text-left my-4 py-3  ">
        <div class=" h-full flex justify-center  content-center">
          <img class="block h-full " src="@/assets/jerry.png" />
        </div>
        <div class=" flex mx-3  flex-grow  justify-between  flex-col  ">
          <span class="text-gray-500  text-xl  font-serif ">{{
            item.User.name
          }}</span>
          <p class=" text-gray-600  text-2xl">
            {{ item.Content }}
          </p>
        </div>
        <div class="flex  flex-col h-full   cursor-pointer  justify-end ">
          <!-- <p
            @click.stop="openInput(item)"
            class=" text-gray-500 hover:text-purple-500 "
          >
            回复
          </p> -->
        </div>
      </div>
      <!-- <Comment :comments="item.Replys"></Comment> -->
      <!-- <InputBox :reply="item" v-if="item.uuid == InputBoxID"></InputBox> -->
    </div>
  </div>
</template>
<script>
import InputBox from "./InputBox.vue";
export default {
  name: "Comment",
  components: { InputBox },
  props: {
    comments: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      InputBoxID: "",
      comments: []
    };
  },
  mounted() {
     console.log(this.comments)
  },
  methods: {
    openInput(user) {
      this.InputBoxID = user.uuid;
    }
  }
};
</script>
