<template>
	<view>
		<slot :user="header"></slot>
		<slot name="header" :user="header"></slot>
		<slot name="main" :user="user"></slot>
		<slot name="footer" :user="footer"></slot>
	</view>
</template>
<script>
export default {
	name: "Next",
	data() {
		return {
			user: {
				name: ""
			},
			header: {
				name: ""
			},
			footer: {
				name: ""
			},
			footer: {
				name: ""
			}
		};
	}
}
</script>

