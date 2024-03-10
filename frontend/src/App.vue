<template>
	<router-view></router-view>
</template>

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

.input {
	width: 300px;
	margin: 0 auto;
	border: 1px solid #DEE4EA;
	border-radius: 5px;
	padding: 30px;
	box-sizing: border-box;
}

label {
	font-size: 14px;
}

input,
select {
	display: block;
	width: 100%;
	margin: 2px auto 10px auto;
	font-size: 18px;
	background: #DEE4EA;
	color: #161A1D;
	padding: 10px;
	border: none;
	outline: none;
	border-radius: 5px;
	box-sizing: border-box;
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
	border-radius: 5px;
	cursor: pointer;
	transition: all .25s ease;

	&:hover {
		transform: scale(1.1);
	}
}
</style>