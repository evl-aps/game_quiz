<template>
	<h1>QUIZ</h1>

	<StartGame v-if="!$store.getters.getStartGame"/>
	<QuestionGame v-else />
</template>

<script>
import StartGame from './components/StartGame.vue'
import QuestionGame from './components/QuestionGame.vue'

export default {
	components: { StartGame, QuestionGame },

	data() {
		return {
		}
	},
}
</script>

<style lang="scss">
* {
	font-family: Arial, Helvetica, sans-serif;
	margin: 0; 
	padding: 0;
	color: #DEE4EA;
}

body {
	background: #161A1D;
}

h1 {
	text-align: center;
	margin: 20px 0;
}
</style>