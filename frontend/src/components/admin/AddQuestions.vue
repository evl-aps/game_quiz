<template>
<div class="adminContent">
	<div class="adminContentItem">
		<h2>Добавить вопрос</h2>

		<label for="question">Вопрос</label>
		<input type="text" id="question" v-model="computedQuestion">

		<button class="greenBtn" @click="addQuestion">Добавить вопрос</button>
	</div>

	<div class="adminContentItem">
		<h2>Добавить ответы</h2>

		<label for="quest">Выбрать вопрос</label>
		<select id="quest">
			<option value="" selected disabled hidden></option>
			<option 
				v-for="question in $store.getters.getQuestionsList" 
				:key="question.ID"
				:value="question.ID"
			>
				{{ question.Question }}
			</option>
		</select>

		<label for="">Ответ 1</label>
		<input type="text">

		<label for="">Ответ 2</label>
		<input type="text">

		<label for="">Ответ 3</label>
		<input type="text">

		<label for="">Ответ 4</label>
		<input type="text">

		<button class="greenBtn">Добавить ответы</button>
	</div>
</div>
</template>

<script>
export default {
	props: {
		question: String,
		addQuestion: Function,
	},

	computed: {
		computedQuestion: {
			get() {
				return this.$store.getters.getQuestion
			},
			set(value) {
				this.$store.commit('updateQuestion', value)
				// return this.$store.getQuestion()
			}
		}
	}
}
</script>