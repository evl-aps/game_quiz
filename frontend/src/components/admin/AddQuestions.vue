<template>
<div class="adminContent">
	<div class="adminContentItem">
		<h2>Добавить вопрос</h2>

		<label for="question">Вопрос</label>
		<input type="text" id="question">

		<button class="greenBtn" @click="addQuestion">Добавить вопрос</button>
	</div>

	<div class="adminContentItem">
		<h2>Добавить ответы</h2>

		<label for="quest">Выбрать вопрос</label>
		<select id="quest">
			<option value="" selected disabled hidden></option>
			<option value="1">test</option>
			<option value="2">test 2</option>
			<option value="3">test 3</option>
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
}
</script>