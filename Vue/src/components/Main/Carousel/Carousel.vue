<template>
  <el-carousel
    class="carousel"
    indicator-position="none"
    :interval="3000"
    height="300px"
    :type="type"
  >
    <el-carousel-item v-for="(item, index) in pictures" :key="index">
      <el-image :src="item.large" height="300px"></el-image>
    </el-carousel-item>
  </el-carousel>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "Carousel",
  data() {
    return {
      carouselUrl: [],
    };
  },
  computed: mapState({
    pictures: state =>{ 
      let random  = Vue.prototype.$random(0,195)
      return state.pictures.length ? state.pictures.slice(random,random+5) : []
    }})
};
</script>
<style>
.carousel {
  width: 100%;
  margin:10px 0px;  
  border-radius: 5px;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}
.el-image {
  height: 100%;
  width: 100%;
}
.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>