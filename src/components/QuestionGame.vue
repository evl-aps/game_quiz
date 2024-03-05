<template>
	<h2>Привет, {{ $store.getters.getPlayerName }}</h2>
	
	<div class="questionPlace">
		<h3>Вопрос № {{ index + 1 }}</h3>
		<p>{{ questions[index].question }}</p>

		<div class="radio" v-for="answer in questions[index].answers" :key="answer.id">
			<input type="radio" :id="answer.id" :value="answer.value" v-model="questions[index].userAnswer">
			<label :for="answer.id">{{ answer.text }}</label>
		</div>

		<button class="greenBtn" @click="nextQuestion">
			{{ index + 1 == questions.length ? 'FINISH' : 'NEXT'}}
		</button>
	</div>
</template>

<script>
export default {
	props: {
		index: Number,
		questions: Array,
		nextQuestion: Function,
	}
}
</script>

<style lang="scss" scoped>
.questionPlace {
	padding: 30px;
	width: 70%;
	margin: 20px auto;
	border: 1px solid #DEE4EA;
	border-radius: 10px;

	h3 {
		margin-bottom: 10px;
	}

	p {
		margin-bottom: 10px;
	}

	.radio {
		margin-bottom: 10px;
		cursor: pointer;

		label {
			margin-left: 8px;
			cursor: pointer;
		}
	}

	.greenBtn {
		margin-right: 0;
		margin-left: auto;
	}
}
</style>