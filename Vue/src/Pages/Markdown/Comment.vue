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
          <p
            @click.stop="openInput(index)"
            class=" text-gray-500 hover:text-purple-500"
          >
            回复
          </p>
        </div>
      </div>
      <!-- <Comment :comments="item.Replys"></Comment> -->
      <InputBox
        @close="close"
        :reply="item"
        v-if="index == InputBoxID"
      ></InputBox>
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
    },
    InputBoxID: {
      type: Number,
      default: Infinity
    }
  },
  data() {
    return {
      InputBoxID: Infinity,
      comments: []
    };
  },
  mounted() {
    console.log(this.comments);
  },
  methods: {
    openInput(index) {
      this.InputBoxID = index;
    },
    close() {
      this.InputBoxID = Infinity;
    }
  }
};
</script>
