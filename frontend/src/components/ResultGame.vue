<template>
	<h2>Ваши результаты, {{ $store.getters.getPlayerName }}</h2>

	<div class="resultPlace">
		10/{{ $store.getters.getQuestions.length }}
	</div>

	<button class="greenBtn" @click="reloadPage">Repeat</button>
</template>

<script>
export default {
	methods: {
		reloadPage() {
			this.$store.commit('updatePlayerName', '')
			this.$store.commit('updateStartGame', false)
			this.$store.commit('updateFinishGame', false)
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
	border-radius: 5px;
}
</style>