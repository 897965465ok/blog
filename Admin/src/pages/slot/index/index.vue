<template>
	<view class="content">
		{{item}}
		<view class="text-area">
			<Next>
				<!-- 一个不带 name 的 <slot> 出口会带有隐含的名字“default”。 -->
				<template v-slot:header="solotProps">
					{{solotProps.user.name}}
				</template>
				<template v-slot:main="solotProps">
					{{solotProps.user.name}}
				</template>
				<template v-slot:footer="solotProps">
					{{solotProps.user.name}}
				</template>
				<template v-slot:default="solotProps">
					{{solotProps.user.name}}
				</template>
			</Next>
			<button>点击就送</button>
		</view>
	</view>
</template>
<script>
	import Next from '../Next/Next.vue'
	export default {
		data() {
			return {
              item:"狗都不写代码"
			}
		},
		components: {
			Next
		},
	
		methods: {
		
		}
	}
</script>
