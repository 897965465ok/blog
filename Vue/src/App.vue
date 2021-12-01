<template>
  <div id="app">
    <transition name="fade">
      <router-view :key="$route.path" />
    </transition>
    <!-- <Music></Music> -->
    <!-- src/components/HelloWorld.vue -->    
<h1 class="font-semibold text-4xl text-purple-700">你好</h1>
  </div>
</template>
<script>
import Music from "./Pages/Music/Music";
export default {
  components: {
    Music
  },
  name: "App"
};
</script>
<style lang="scss">

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
// * {
//   margin: 0px;
//   padding: 0px;
//   font-family: "Microsoft Yahei";
// }
body ,html {
  height: 100%;
  background-color: #f7f4f4;
}
</style>
