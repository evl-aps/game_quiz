<template>
	<h2>Ваши результаты, {{ $store.getters.getPlayerName }}</h2>

	<div class="resultPlace">
		10/{{ questions.length }}
	</div>

	<button class="greenBtn" @click="reloadPage">Repeat</button>
</template>

<script>
export default {
	props: {
		questions: Array,
	},

	methods: {
		reloadPage() {
			location.reload()
		}
	}
}
</script>

<style lang="scss" scoped>
.resultPlace {
	padding: 30px;
	width: 70%;
	margin: 20px auto;
	border: 1px solid #DEE4EA;
	border-radius: 10px;
}
</style>