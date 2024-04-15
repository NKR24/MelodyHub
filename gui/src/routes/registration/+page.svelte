<script lang="ts">
	import { createUser, getUserId } from '../../users/users';
	import type { UserID } from '../../types';
	import StartScreen from '../../components/StartScreen/StartScreen.svelte';

	let id: UserID[] = $state([]);
	let username = $state('');
	let email = $state('');
	let password = $state('');

	async function postUser() {
		createUser({ username, email, password }).then(() => {
			username = '';
			email = '';
			password = '';
		});
	}
</script>

<main>
	<div class="container">
		<h1>CREATE ACCOUNT</h1>
		<input type="text" bind:value={username} placeholder="Username"/>
		<input type="email" bind:value={email} placeholder="Email" />
		<input type="password" bind:value={password} placeholder="Password" />
		<button onclick={postUser}>SUBMIT</button>
	</div>
</main>

<style>
	h1 {
		font-size: 50px;
		color: black;
	}

	button {
		color: black;
	}

	main {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
		height: 100%;
	}

	.container {
		row-gap: 30px;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
	}

	input {
		all: unset;
		border-bottom: 1px solid black;
		width: 300px;
		height: 50px;
	}
</style>
