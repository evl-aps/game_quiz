<template>
	<div class="alert">
		<span v-for="(err, idx) in errors" :key="idx">{{ err }}</span>
	</div>
</template>

<script>
export default {
	props: {
		errors: Array
	}
}
</script>

<style lang="scss" scoped>
.alert {
	position: fixed;
	top: 20px;
	left: 50%;
	transform: translateX(-50%);
	padding: 20px;
	border-top: 8px solid #ff0000;
	background: #a70000;
	width: 90%;
	border-radius: 5px;
}
</style>