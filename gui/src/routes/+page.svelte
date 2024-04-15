<script>
	import { onMount } from 'svelte';
	import Navbar from '../components/Navbar/Navbar.svelte';
	import StartScreen from '../components/StartScreen/StartScreen.svelte';
	import '../app.scss';

	let isLoaded = $state(false);
	onMount(() => {
		if (!localStorage.getItem('loaded')) {
			localStorage.setItem('loaded', 'true');
			isLoaded = true;
			setTimeout(() => {
				isLoaded = false;
			}, 5000);
		}
	});
</script>

{#if isLoaded}
	<StartScreen />
{:else}
	<Navbar />
	<slot />
{/if}
