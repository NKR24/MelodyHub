<script>
	let email = $state('');
	let password = $state('');
</script>

<main>
	<div class="container">
		<h1>CREATE ACCOUNT</h1>
		<input type="email" bind:value={email} placeholder="Email" />
		<input type="password" bind:value={password} placeholder="Password" />
		<button>SUBMIT</button>
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
