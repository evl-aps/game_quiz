<template>
	<h1>QUIZ</h1>

	<StartGame v-if="!$store.getters.getStartGame"/>
	<QuestionGame
		:index="index"
		:questions="questions"
		:nextQuestion="nextQuestion" 
		v-else-if="$store.getters.getStartGame && !$store.getters.getFinishGame"
	/>

	<ResultGame v-else :questions="questions"/>
</template>

<script>
import StartGame from './components/StartGame.vue'
import QuestionGame from './components/QuestionGame.vue'
import ResultGame from './components/ResultGame.vue'

export default {
	components: { StartGame, QuestionGame, ResultGame },

	data() {
		return {
			index: 0,
			questions: [
				{
					question: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
								Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`,
					answers: [
						{
							id: 1,
							text: 'Вариант 1',
							value: 1,
						},
						{
							id: 2,
							text: 'Вариант 2',
							value: 2,
						},
						{
							id: 3,
							text: 'Вариант 3',
							value: 3,
						},
						{
							id: 4,
							text: 'Вариант 4',
							value: 4,
						}
					],
					correctAnswer: 1,
					userAnswer: null
				},
				{
					question: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua???`,
					answers: [
						{
							id: 1,
							text: 'Вариант 14',
							value: 1,
						},
						{
							id: 2,
							text: 'Вариант 21',
							value: 2,
						},
						{
							id: 3,
							text: 'Вариант 32',
							value: 3,
						},
						{
							id: 4,
							text: 'Вариант 43',
							value: 4,
						}
					],
					correctAnswer: 3,
					userAnswer: null
				},
			]
		}
	},

	methods: {
		nextQuestion() {
			if(this.index + 1 != this.questions.length) {
				this.index++
			} else {
				this.$store.commit('updateFinishGame', true)
			}
		}
	}
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

h2 {
	text-align: center;
}

.greenBtn {
	font-size: 22px;
	font-weight: 600;
	padding: 10px 30px;
	outline: none;
	margin: 0 auto;
	display: block;
	background: #1F845A;
	border: none;
	border-radius: 10px;
	cursor: pointer;
	transition: all .25s ease;

	&:hover {
		transform: scale(1.1);
	}
}
</style>