<template>
  <div id="app">
    <transition name="fade">
      <router-view :key="$route.path" />
    </transition>
    <Music></Music>
  </div>
</template>
<script>
export default {
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

body ,html {
  height: 100%;
  width: 100%;
  // background-color: #e9ecef;
}
</style>
