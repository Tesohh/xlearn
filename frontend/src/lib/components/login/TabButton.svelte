<script lang="ts">
	import { tabState } from '$lib/writables';

	export let title: 'Login' | 'Signup';
</script>

<button
	on:click={() => tabState.set(title)}
	class="flex items-center justify-center text-center w-full h-full"
>
	<div
		class="flex items-center justify-center rounded-t-md text-center h-full w-full {$tabState ==
		title
			? 'bg-primary'
			: 'bg-secondary'}"
	>
		<h1>{title}</h1>
	</div>
</button>
