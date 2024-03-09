<template>
<div>
	<h2>Привет, {{ $store.getters.getPlayerName }}</h2>
	
	<div class="questionPlace">
		<h3>Вопрос № {{ index + 1 }}</h3>
		<p>{{ $store.getters.getQuestions[index].question }}</p>

		<div class="radio" v-for="answer in $store.getters.getQuestions[index].answers" :key="answer.id">
			<div class="answer" @click="clickToAnswer(answer, index)">
				<div class="circle"><div :class="{ 'choisen': answer.choisen }"></div></div>
				<span>{{ answer.text }}</span>
			</div>
		</div>

		<button class="greenBtn" @click="nextQuestion">
			{{ index + 1 == $store.getters.getQuestions.length ? 'FINISH' : 'NEXT'}}
		</button>
	</div>
</div>
</template>

<script>
export default {
	props: {
		index: Number,
		clickToAnswer: Function,
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
	border-radius: 5px;
	font-size: 18px;

	h3 {
		margin-bottom: 10px;
	}

	p {
		margin-bottom: 10px;
	}

	.radio {
		margin-bottom: 10px;
		cursor: pointer;

		.answer {
			display: flex;
			align-items: center;

			.circle {
				width: 20px;
				height: 20px;
				border: 1px solid #DEE4EA;
				border-radius: 50%;
				position: relative;

				div {
					position: absolute;
					width: 14px;
					height: 14px;
					top: 50%;
					left: 50%;
					transform: translate(-50%, -50%);
					background: #DEE4EA;
					border-radius: 50%;
					display: none;
				}

				div.choisen {
					display: block;
				}
			}

			span {
				display: block;
				margin-left: 10px;
				font-size: 18px;
			}
		}

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