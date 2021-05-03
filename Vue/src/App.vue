<template>
  <div id="app">
    <transition name="fade">
      <router-view :key="$route.path" />
    </transition>
  </div>
</template>
<script>
export default {
  name: "App",
};
</script>
<style lang="scss">
.fade-enter-active, .fade-leave-active {
  transition: opacity .3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
* {
  margin: 0px;
  padding: 0px;
  font-family: "Microsoft Yahei";
}
body {
  background-color: #f7f4f4;
}
</style>
