<template>
  <aplayer autoplay :music="first" :list="list" />
  <!-- <iframe
      frameborder="no"
      border="0"
      marginwidth="0"
      marginheight="0"
      width="330"
      height="450"
      src="//music.163.com/outchain/player?type=0&id=6744975757&auto=0&height=430"
    >
  </iframe>-->
</template>
<script>
import Aplayer from 'vue-aplayer'
export default {
  name: "music",
  components: {
    Aplayer
  },
  data() {
    return {
      list: [
        {
          title: '叹',
          artist: 'Silent Siren',
          src: 'https://ws.stream.qqmusic.qq.com/C400000cFbj12hpdo1.m4a?guid=210790619&vkey=B84C63339D2D3FD8957C42CE3416FE41726A1FC5024B2DD747E8197CB180339D5FF4E8D83694DC9888B372BBC5E0650D90BB779A84B1A212&uin=&fromtag=66',
          // pic: 'https://moeplayer.b0.upaiyun.com/aplayer/secretbase.jpg'
        },
        {
          title: '叹',
          artist: 'Silent Siren',
          src: 'https://ws.stream.qqmusic.qq.com/C400000cFbj12hpdo1.m4a?guid=210790619&vkey=B84C63339D2D3FD8957C42CE3416FE41726A1FC5024B2DD747E8197CB180339D5FF4E8D83694DC9888B372BBC5E0650D90BB779A84B1A212&uin=&fromtag=66',
          // pic: 'https://moeplayer.b0.upaiyun.com/aplayer/secretbase.jpg'
        },
      ],
      first: {
          title: '叹',
          artist: 'Silent Siren',
          src: 'https://ws.stream.qqmusic.qq.com/C400000cFbj12hpdo1.m4a?guid=210790619&vkey=B84C63339D2D3FD8957C42CE3416FE41726A1FC5024B2DD747E8197CB180339D5FF4E8D83694DC9888B372BBC5E0650D90BB779A84B1A212&uin=&fromtag=66',
          // pic: 'https://moeplayer.b0.upaiyun.com/aplayer/secretbase.jpg'
        },
    }
  }
};
</script>